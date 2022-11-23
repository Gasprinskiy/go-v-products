package reqiestmethods

import (
	"fmt"
	"jun2/reqiestmethods/methodhelpers"
	"jun2/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// СreateProduct создает продукт с вариацией
func CreateProduct(c *gin.Context) {
	var productId int
	var payload structs.ProductWithVariations

	err := c.BindJSON(&payload)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	productPayload := structs.NewProduct(payload)
	productRow, err := DB.NamedQuery("INSERT INTO product VALUES (DEFAULT, :product_name, :description, :tags) RETURNING product_id", productPayload)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if productRow.Next() {
		productRow.Scan(&productId)
	}

	variationPayload := structs.NewVariation(payload, productId)
	variationId, err, _ := methodhelpers.CreateVariationByDefinedPayload(c, variationPayload, DB)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product_id":   productId,
		"variation_id": variationId,
	})
}

// СreateVariation создает вариацию для продутка по id продукта
func CreateVariation(c *gin.Context) {
	id, err, exist := methodhelpers.CreateVariationByDefinedPayload(c, structs.Variation{}, DB)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Variation exist",
			"id":      id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// СreateOrUpdatePrice проверяет цену в заданном диапазоне, если она есть то останавливает ее и записывает новую цену
func CreateOrUpdatePrice(c *gin.Context) {
	var price structs.Price
	err := c.BindJSON(&price)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	timeLayout := "2006-01-02 15:04:05"
	tFrom, err := time.Parse(timeLayout, price.ActiveFrom)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var tTill time.Time
	if price.ActiveTill != "" {
		tTill, err = time.Parse(timeLayout, price.ActiveTill)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}

	existenceCheck := methodhelpers.DefinePriceExistenceCheckMethod(tFrom, tTill, price.VariationId, DB)
	err, existPriceId := existenceCheck()
	if err != nil && err.Error() != "sql: no rows in result set" {
		c.String(http.StatusBadRequest, err.Error())
	}
	if existPriceId > 0 {
		_, err := DB.Exec("UPDATE product$price SET active_till=CURRENT_TIMESTAMP WHERE price_id=$1", existPriceId)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}

	createPrice := methodhelpers.DefinePriceCreationMethod(tFrom, tTill, price, DB)
	err = createPrice()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "OK")
}

// AddProductToStock добавляет продукт, его количество, типы формы и вариации в складской учет, если существует продукт с такими данными на складе то он просто добавляет количество
func AddProductToStock(c *gin.Context) {
	var stockProduct structs.StockProduct
	err := c.BindJSON(&stockProduct)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	existAccountingId, err := methodhelpers.CheckForProductStockExistence(stockProduct, DB)
	if err != nil && err.Error() != "sql: no rows in result set" {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if existAccountingId > 0 {
		_, err = DB.Exec("UPDATE product$stocks SET amount=amount+$1 WHERE accounting_id=$2", stockProduct.Amount, existAccountingId)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "EDITED")
		return
	}

	_, err = DB.NamedQuery("INSERT INTO product$stocks VALUES (DEFAULT, :stock_id, :product_id, :amount, :variation_id)", stockProduct)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}

// AddProductSale фиксирует продажу продукта на складе, обновляет количество и добавляет запись в таблицу с продажами
func AddProductSale(c *gin.Context) {
	var sale structs.Sales
	err := c.BindJSON(&sale)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var amountCheck structs.StockProduct
	err = DB.Get(&amountCheck, `SELECT ps.product_id, ps.stock_id, ps.variation_id, ps.amount
								FROM product$stocks ps
								WHERE (ps.product_id=$1 and ps.variation_id=$2 and ps.stock_id=$3)`,
		sale.ProductId, sale.VariationId, sale.StockId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if amountCheck.Amount < sale.Amount {
		c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Product amount: %d", amountCheck.Amount))
		return
	}

	_, err = DB.NamedQuery("INSERT INTO product$sales VALUES (DEFAULT, :product_id, :variation_id, :stock_id, DEFAULT, :amount)", sale)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	changedRow, err := DB.NamedQuery(`UPDATE product$stocks SET amount=amount-:amount
							WHERE stock_id=:stock_id AND product_id=:product_id AND variation_id=:variation_id
							RETURNING accounting_id`, sale)
	if !changedRow.Next() {
		c.String(http.StatusBadRequest, "No product in given stock")
		return
	}
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "OK")
}

// GetSalesReport возвращает отчет по продажам по переданным параметрам
func GetSalesReport(c *gin.Context) {
	var reportBody structs.SalesReportBody
	err := c.BindJSON(&reportBody)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if reportBody.StartDate == "" {
		c.String(http.StatusBadRequest, "start_date is required")
		return
	}
	if reportBody.EndDate == "" {
		c.String(http.StatusBadRequest, "end_date is required")
		return
	}

	getReport := methodhelpers.DefineGetSalesReportMethod(reportBody, DB)
	report, err := getReport()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if report == nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, report)
}
