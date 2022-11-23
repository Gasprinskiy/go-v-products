package reqiestmethods

import (
	"fmt"
	"jun2/helpers"
	"jun2/reqiestmethods/methodhelpers"
	"jun2/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// возвращает полную информацию по продукту, его вариации и в каких складах есть такущая вариация
func GetFullProductDataById(c *gin.Context) {
	id, err := helpers.ParseId(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var product structs.Product
	err = DB.Get(&product, "SELECT p.product_id, p.product_name, p.description, p.tags FROM product p WHERE p.product_id=$1", id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	product.Variations, err = methodhelpers.GetVariationAndPriceById(id, DB)
	fmt.Println(err, "GetVariationAndPriceById")
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println("GetVariationAndPriceById")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	product, err = methodhelpers.GetVariationStockAvailability(product, DB)
	fmt.Println(err, "GetVariationStockAvailability")
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println("GetVariationStockAvailability")
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

// возвращает список продуктов, фильтрует по параметрам tags и limit если параметры переданы
func GetProductList(c *gin.Context) {
	productList := []structs.Product{}

	limit, err := helpers.ParseString(c.Query("limit"))
	if err != nil {
		limit = 3
	}
	tags := c.Query("tag")
	requestString := methodhelpers.DefineProductListRequestString(tags)

	err = DB.Select(&productList, requestString, limit)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, productList)
}

// возвращает сколько в общем товаров содержиться на складах, с фильтром по product_id возвращает на каких складах есть этот продукт и в каких количествах
func GetStockInfo(c *gin.Context) {
	productId, err := helpers.ParseString(c.Query("product_id"))
	getInfo := methodhelpers.DefineGetStockInfoMethod(productId, DB)

	stockList, err := getInfo()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if stockList == nil {
		c.String(http.StatusNotFound, "Not found")
		return
	}

	c.JSON(http.StatusOK, stockList)
}
