package ginapi

import (
	"fmt"
	"jun2/cur/config"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/price"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/tools/sqlnull"
	"jun2/cur/uimport"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinApi struct {
	conf *config.Config
	r    *gin.Engine
	ui   uimport.UsecaseImports
}

func NewGinApi(conf *config.Config, e *gin.Engine, ui uimport.UsecaseImports) *GinApi {
	api := GinApi{
		conf,
		e,
		ui,
	}

	api.r.GET("/product/:id", api.FindFullProductDataById)
	api.r.GET("/stock", api.FindStockInfo)
	api.r.GET("/product_list", api.FindProductList)
	api.r.GET("/stock_list", api.LoadStockList)
	api.r.GET("/sales_list", api.FindSalesList)
	api.r.POST("/stock/add", api.CreateStock)
	api.r.POST("/product/add", api.CreateProduct)
	api.r.POST("/product/add/variation", api.CreateVariation)
	api.r.POST("product/price", api.CreatePrice)
	api.r.POST("/product/add/stock", api.AddProductToStock)
	api.r.POST("/buy", api.AddProductSale)
	api.r.POST("/sales", api.FindSalesReportList)

	return &api
}

func (e *GinApi) StartServer() {
	e.r.Run(e.conf.ServerIP())
}

func (e *GinApi) CreateProduct(c *gin.Context) {
	param := product.AddProductParams{}
	if err := c.BindJSON(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if param.VariationType <= 0 {
		paramErr := global.ErrInvalidParam("variation_type")
		c.String(http.StatusBadRequest, paramErr.Error())
		return
	}

	createdID, err := e.ui.Usecase.Product.CreateProduct(param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(createdID)

	c.JSON(http.StatusOK, createdID)
}

func (e *GinApi) CreateVariation(c *gin.Context) {
	param := product.AddProductParams{}
	if err := c.BindJSON(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if param.VariationType <= 0 {
		paramErr := global.ErrInvalidParam("variation_type")
		c.String(http.StatusBadRequest, paramErr.Error())
		return
	}

	createdID, err := e.ui.Usecase.Variation.CreateVariation(param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, createdID)

}

func (e *GinApi) CreatePrice(c *gin.Context) {
	newPrice := price.Price{}

	if err := c.ShouldBindJSON(&newPrice); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if newPrice.Price <= 0 {
		paramErr := global.ErrInvalidParam("price")
		c.String(http.StatusBadRequest, paramErr.Error())
		return
	}

	priceID, err := e.ui.Usecase.Price.CreatePrice(newPrice)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, priceID)

}

func (e *GinApi) CreateStock(c *gin.Context) {
	newStock := stock.StockInfo{}

	if err := c.BindJSON(&newStock); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	stockID, err := e.ui.Usecase.Stock.CreateStock(newStock)
	fmt.Println(err)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, stockID)
}

func (e *GinApi) AddProductToStock(c *gin.Context) {
	param := stock.ProductStock{}
	if err := c.BindJSON(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if param.Amount <= 0 {
		paramErr := global.ErrInvalidParam("amount")
		c.String(http.StatusBadRequest, paramErr.Error())
		return
	}

	accountingID, err := e.ui.Usecase.Stock.AddProductToStock(param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, accountingID)
}

func (e *GinApi) AddProductSale(c *gin.Context) {
	param := sales.Sales{}
	if err := c.BindJSON(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if param.Amount <= 0 {
		paramErr := global.ErrInvalidParam("amount")
		c.String(http.StatusBadRequest, paramErr.Error())
		return
	}

	id, err := e.ui.Usecase.Sales.AddProductSale(param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (e *GinApi) FindFullProductDataById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	data, err := e.ui.Usecase.Product.FindProductByID(id)
	if err != nil {
		if err == global.ErrNoData && data.ID > 0 {
			c.JSON(http.StatusOK, data)
			return
		}
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func (e *GinApi) FindStockInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("product_id"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	offset = offset - 1
	offset = limit * offset

	data, err := e.ui.Usecase.Stock.FindStockInfo(id, limit, offset)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (e *GinApi) FindSalesList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	offset = offset - 1
	offset = limit * offset

	data, err := e.ui.Usecase.Sales.FindSalesList(limit, offset)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (e *GinApi) LoadStockList(c *gin.Context) {
	data, err := e.ui.Usecase.Stock.LoadStockList()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (e *GinApi) FindProductList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	tags := c.Query("tag")

	offset = offset - 1
	offset = limit * offset

	data, err := e.ui.Usecase.Product.FindProductList(tags, limit, offset)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func (e *GinApi) FindSalesReportList(c *gin.Context) {
	param := sales.SalesReportRequest{}
	if err := c.BindJSON(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if param.Offset.Valid {
		param.Offset = sqlnull.NewInt64(param.Offset.GetInt() - 1)
		param.Offset = sqlnull.NewInt64(param.Limit.GetInt() * param.Offset.GetInt())
	} else {
		param.Offset = sqlnull.NewInt64(global.DefaultOffset)
	}
	data, err := e.ui.Usecase.Sales.FindSalesReport(param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
