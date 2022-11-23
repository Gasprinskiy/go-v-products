package main

import (
	"jun2/helpers"
	"jun2/reqiestmethods"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	r := gin.Default()

	reqiestmethods.DB, err = sqlx.Open("postgres", helpers.GetConfig(&reqiestmethods.Config))

	if err != nil {
		log.Fatalln(err)
	}

	if err := reqiestmethods.DB.Ping(); err != nil {
		log.Fatalln(err)
	}

	r.GET("/product_list", reqiestmethods.GetProductList)
	r.GET("/product/:id", reqiestmethods.GetFullProductDataById)
	r.GET("/stock", reqiestmethods.GetStockInfo)
	r.POST("/product/add", reqiestmethods.CreateProduct)
	r.POST("/product/add/variation", reqiestmethods.CreateVariation)
	r.POST("/product/price", reqiestmethods.CreateOrUpdatePrice)
	r.POST("/product/add/stock", reqiestmethods.AddProductToStock)
	r.POST("/buy", reqiestmethods.AddProductSale)
	r.POST("/sales", reqiestmethods.GetSalesReport)

	r.Run()
}
