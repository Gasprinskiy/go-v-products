package methodhelpers

import (
	"fmt"
	"jun2/structs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// createVariation создает вариацию для переданного продутка если definedPayload не пустой, если definedPayload пустой создает вариацию для существующего продукта если такой вариации
func CreateVariationByDefinedPayload(c *gin.Context, definedPayload structs.Variation, DB *sqlx.DB) (variationId int, err error, exist bool) {
	payload := definedPayload
	if definedPayload.ProductId <= 0 {
		err = c.BindJSON(&payload)
		if err != nil {
			return
		}
	}

	var existVariation structs.Variation
	err = DB.Get(&existVariation, "SELECT pv.variation_id FROM product$variations pv WHERE pv.product_id=$1 AND pv.variation_type=$2 AND unit_type=$3", payload.ProductId, payload.VariationType, payload.UnitType)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return
	}
	if existVariation.VariationId > 0 {
		variationId = existVariation.VariationId
		exist = true
		return
	}

	variationRow, err := DB.NamedQuery("INSERT INTO product$variations VALUES (DEFAULT, :product_id, :variation_type, :unit_type) RETURNING variation_id", payload)
	if err != nil {
		return
	}
	if variationRow.Next() {
		variationRow.Scan(&variationId)
	}

	return
}

// DefinePriceExistenceCheckMethod возвращает функцию для проверки сущетсвования цен по заданному диапазону
func DefinePriceExistenceCheckMethod(from, till time.Time, id int, DB *sqlx.DB) func() (error, int) {
	var existPrice structs.Price
	existRequest := "SELECT pr.price_id FROM product$price pr WHERE (pr.variation_id=$1) AND CURRENT_TIMESTAMP BETWEEN pr.active_from AND pr.active_till"
	switch {
	case till.IsZero():
		return func() (error, int) {
			err := DB.Get(&existPrice, existRequest+" AND $2 BETWEEN active_from AND active_till", id, from.Local())
			return err, existPrice.PriceId
		}
	default:
		return func() (error, int) {
			err := DB.Get(&existPrice, existRequest+" AND CURRENT_TIMESTAMP BETWEEN $2 AND $3", id, from.Local(), till.Local())
			return err, existPrice.PriceId
		}
	}
}

// DefinePriceCreationMethod определяет функцию создания цены
func DefinePriceCreationMethod(from, till time.Time, p structs.Price, DB *sqlx.DB) func() error {
	switch {
	case till.IsZero():
		return func() error {
			_, err := DB.Exec("INSERT INTO product$price VALUES (DEFAULT, $1, $2, DEFAULT, $3)", p.VariationId, p.Price, from.Local())
			return err
		}
	default:
		return func() error {
			_, err := DB.Exec("INSERT INTO product$price VALUES (DEFAULT, $1, $2, $3, $4)", p.VariationId, p.Price, till.Local(), from.Local())
			return err
		}
	}
}

// CheckForProductStockExistence проверяет есть ли товар на складе по заданным параметрам, возвращает ID учета (accounting_id)
func CheckForProductStockExistence(sp structs.StockProduct, DB *sqlx.DB) (int, error) {
	var existStockProduct structs.StockProduct
	err := DB.Get(&existStockProduct, "SELECT accounting_id FROM product$stocks ps WHERE ps.stock_id=$1 AND ps.variation_id=$2 AND ps.product_id=$3", sp.StockId, sp.VariationId, sp.ProductId)
	return existStockProduct.AccountingId, err
}

// GetVariationAndPriceById возвращает массив с вариациями и его ценами
func GetVariationAndPriceById(id int, DB *sqlx.DB) (res []structs.Variation, err error) {
	variationPriceSelect := `
	SELECT pv.variation_id, pv.product_id, pv.variation_type, pv.unit_type, pr.price
	FROM product$variations pv
        JOIN product$price pr ON (pv.variation_id = pr.variation_id)
	WHERE (pv.product_id = $1)
	AND CURRENT_TIMESTAMP BETWEEN pr.active_from AND COALESCE(NULL, pr.active_till, pr.active_from + interval '6 month')`

	err = DB.Select(&res, variationPriceSelect, id)

	return
}

// GetVariationStockAvailability возвращает измененный массив типа Product с информацией о вариациях и в каких складах эти вариации есть
func GetVariationStockAvailability(p structs.Product, DB *sqlx.DB) (res structs.Product, err error) {
	stockSelection := `
	SELECT ps.accounting_id, ps.stock_id, ps.product_id, ps.amount, ps.variation_id, st.stock_name, st.location
	FROM product$stocks ps
        JOIN stocks st ON (ps.stock_id = st.stock_id)
	WHERE (ps.variation_id=$1)`

	res = p
	var stockAv structs.StockProduct
	for key, value := range res.Variations {
		err = DB.Get(&stockAv, stockSelection, value.VariationId)
		if err != nil {
			break
		}
		if value.VariationId == stockAv.VariationId {
			res.Variations[key].StockAvailability = append(res.Variations[key].StockAvailability, stockAv)
		}
	}

	return
}

// DefineProductListRequestString возвращает строку запроса для получения списка продуктов по заданным параметрам
func DefineProductListRequestString(tags string) string {
	requeststring := "SELECT p.product_id, p.product_name, p.description, p.tags FROM product p"
	lastPart := "ORDER BY product_id DESC LIMIT $1"
	switch {
	case tags != "":
		stringTag := fmt.Sprintf("'%s%s'", tags, "%")
		return fmt.Sprintf("%s WHERE p.tags LIKE %s %s", requeststring, stringTag, lastPart)
	default:
		return fmt.Sprintf("%s %s", requeststring, lastPart)
	}
}

// DefineGetStockInfoMethod возвращает функцию получения информацию о складах и количества товара в них, или информацию о количестве товара на скаладх по productId
func DefineGetStockInfoMethod(productId int, DB *sqlx.DB) func() (res []structs.StockInfo, err error) {
	switch {
	case productId > 0:
		return func() (res []structs.StockInfo, err error) {
			err = DB.Select(&res, defineGetStockInfoRequestString(productId > 0), productId)
			return
		}
	default:
		return func() (res []structs.StockInfo, err error) {
			err = DB.Select(&res, defineGetStockInfoRequestString(productId > 0))
			return
		}
	}
}

// defineGetStockInfoRequestString определяет стррку запроса для метода DefineGetStockInfoMethod
func defineGetStockInfoRequestString(hasId bool) string {
	requeststring := `SELECT ps.stock_id, st.stock_name, st.location, SUM(ps.amount) AS total_amount
					  FROM product$stocks ps
					  JOIN stocks st ON (ps.stock_id = st.stock_id)`
	lastPart := "GROUP BY (ps.stock_id, st.stock_name, st.location)"
	filterPart := "WHERE (ps.product_id=$1)"
	switch {
	case hasId:
		return fmt.Sprintf("%s %s %s", requeststring, filterPart, lastPart)
	default:
		return fmt.Sprintf("%s %s", requeststring, lastPart)
	}
}

// DefineGetSalesReportMethod возвращает метод для запроса по отчету продаж
func DefineGetSalesReportMethod(rt structs.SalesReportBody, DB *sqlx.DB) func() (res []structs.SalesReport, err error) {
	var start, end time.Time
	limit := 3
	if rt.Limit != 0 {
		limit = rt.Limit
	}
	timeLayout := "2006-01-02 15:04:05"
	start, _ = time.Parse(timeLayout, rt.StartDate)
	end, _ = time.Parse(timeLayout, rt.EndDate)
	fmt.Println(start)
	switch {
	case rt.StorageId != 0 && rt.ProductName == "":
		return func() (res []structs.SalesReport, err error) {
			err = DB.Select(&res, defineGetSalesReportRequestStrin(rt), start, end, limit, rt.StorageId)
			return
		}
	case rt.ProductName != "" && rt.StorageId == 0:
		return func() (res []structs.SalesReport, err error) {
			err = DB.Select(&res, defineGetSalesReportRequestStrin(rt), start, end, limit)
			return
		}
	case rt.StorageId != 0 && rt.ProductName != "":
		return func() (res []structs.SalesReport, err error) {
			err = DB.Select(&res, defineGetSalesReportRequestStrin(rt), start, end, limit, rt.StorageId)
			return
		}
	default:
		return func() (res []structs.SalesReport, err error) {
			err = DB.Select(&res, defineGetSalesReportRequestStrin(rt), start, end, limit)
			return
		}
	}
}

// defineGetSalesReportRequestStrin определяет строку запроса для метода DefineGetSalesReportMethod
func defineGetSalesReportRequestStrin(rt structs.SalesReportBody) string {
	requestString := `SELECT ps.product_id, pr.product_name, pv.variation_type, pv.unit_type, st.stock_name, ps.sold_date, ps.amount
	FROM product$sales ps
			JOIN product pr ON (ps.product_id = pr.product_id)
			JOIN product$variations pv ON (ps.variation_id = pv.variation_id)
			JOIN stocks st ON (ps.stock_id = st.stock_id)
	WHERE ps.sold_date BETWEEN $1 AND $2
	`
	lastPart := "LIMIT $3"
	namePart := fmt.Sprintf("'%s%s%s'", "%", rt.ProductName, "%")
	switch {
	case rt.StorageId != 0 && rt.ProductName == "":
		return fmt.Sprintf("%s %s %s", requestString, "AND (ps.stock_id=$4)", lastPart)
	case rt.ProductName != "" && rt.StorageId == 0:
		return fmt.Sprintf("%s %s %s", requestString, "AND (pr.product_name LIKE "+namePart+")", lastPart)
	case rt.StorageId != 0 && rt.ProductName != "":
		return fmt.Sprintf("%s %s %s %s", requestString, "AND (ps.stock_id=$4) AND (pr.product_name LIKE", namePart+")", lastPart)
	default:
		return fmt.Sprintf("%s %s", requestString, lastPart)
	}
}
