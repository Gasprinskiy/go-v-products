// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/repository/interface.go

// Package repository is a generated GoMock package.
package repository

import (
	price "jun2/cur/internal/entity/price"
	product "jun2/cur/internal/entity/product"
	sales "jun2/cur/internal/entity/sales"
	stock "jun2/cur/internal/entity/stock"
	variation "jun2/cur/internal/entity/variation"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProduct) CreateProduct(newProduct product.Product) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", newProduct)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductMockRecorder) CreateProduct(newProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProduct)(nil).CreateProduct), newProduct)
}

// FindProductByID mocks base method.
func (m *MockProduct) FindProductByID(id int) (product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductByID", id)
	ret0, _ := ret[0].(product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductByID indicates an expected call of FindProductByID.
func (mr *MockProductMockRecorder) FindProductByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductByID", reflect.TypeOf((*MockProduct)(nil).FindProductByID), id)
}

// FindProductCountByTags mocks base method.
func (m *MockProduct) FindProductCountByTags(paramtags string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductCountByTags", paramtags)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductCountByTags indicates an expected call of FindProductCountByTags.
func (mr *MockProductMockRecorder) FindProductCountByTags(paramtags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductCountByTags", reflect.TypeOf((*MockProduct)(nil).FindProductCountByTags), paramtags)
}

// FindProductList mocks base method.
func (m *MockProduct) FindProductList(limit, offset int) ([]product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductList", limit, offset)
	ret0, _ := ret[0].([]product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductList indicates an expected call of FindProductList.
func (mr *MockProductMockRecorder) FindProductList(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductList", reflect.TypeOf((*MockProduct)(nil).FindProductList), limit, offset)
}

// FindProductListByTags mocks base method.
func (m *MockProduct) FindProductListByTags(limit, offset int, tags string) ([]product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductListByTags", limit, offset, tags)
	ret0, _ := ret[0].([]product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductListByTags indicates an expected call of FindProductListByTags.
func (mr *MockProductMockRecorder) FindProductListByTags(limit, offset, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductListByTags", reflect.TypeOf((*MockProduct)(nil).FindProductListByTags), limit, offset, tags)
}

// LoadProductCount mocks base method.
func (m *MockProduct) LoadProductCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadProductCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadProductCount indicates an expected call of LoadProductCount.
func (mr *MockProductMockRecorder) LoadProductCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadProductCount", reflect.TypeOf((*MockProduct)(nil).LoadProductCount))
}

// MockVariation is a mock of Variation interface.
type MockVariation struct {
	ctrl     *gomock.Controller
	recorder *MockVariationMockRecorder
}

// MockVariationMockRecorder is the mock recorder for MockVariation.
type MockVariationMockRecorder struct {
	mock *MockVariation
}

// NewMockVariation creates a new mock instance.
func NewMockVariation(ctrl *gomock.Controller) *MockVariation {
	mock := &MockVariation{ctrl: ctrl}
	mock.recorder = &MockVariationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariation) EXPECT() *MockVariationMockRecorder {
	return m.recorder
}

// CheckVariatonExistence mocks base method.
func (m *MockVariation) CheckVariatonExistence(param variation.Params) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVariatonExistence", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVariatonExistence indicates an expected call of CheckVariatonExistence.
func (mr *MockVariationMockRecorder) CheckVariatonExistence(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVariatonExistence", reflect.TypeOf((*MockVariation)(nil).CheckVariatonExistence), param)
}

// CreateVariation mocks base method.
func (m *MockVariation) CreateVariation(newVariation variation.Params) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVariation", newVariation)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVariation indicates an expected call of CreateVariation.
func (mr *MockVariationMockRecorder) CreateVariation(newVariation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVariation", reflect.TypeOf((*MockVariation)(nil).CreateVariation), newVariation)
}

// FindVariationInfoByID mocks base method.
func (m *MockVariation) FindVariationInfoByID(id int) ([]variation.Variation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVariationInfoByID", id)
	ret0, _ := ret[0].([]variation.Variation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVariationInfoByID indicates an expected call of FindVariationInfoByID.
func (mr *MockVariationMockRecorder) FindVariationInfoByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVariationInfoByID", reflect.TypeOf((*MockVariation)(nil).FindVariationInfoByID), id)
}

// MockStock is a mock of Stock interface.
type MockStock struct {
	ctrl     *gomock.Controller
	recorder *MockStockMockRecorder
}

// MockStockMockRecorder is the mock recorder for MockStock.
type MockStockMockRecorder struct {
	mock *MockStock
}

// NewMockStock creates a new mock instance.
func NewMockStock(ctrl *gomock.Controller) *MockStock {
	mock := &MockStock{ctrl: ctrl}
	mock.recorder = &MockStockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStock) EXPECT() *MockStockMockRecorder {
	return m.recorder
}

// AddProductToStock mocks base method.
func (m *MockStock) AddProductToStock(newAccounting stock.ProductStock) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductToStock", newAccounting)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProductToStock indicates an expected call of AddProductToStock.
func (mr *MockStockMockRecorder) AddProductToStock(newAccounting interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductToStock", reflect.TypeOf((*MockStock)(nil).AddProductToStock), newAccounting)
}

// CheckProductExistence mocks base method.
func (m *MockStock) CheckProductExistence(param stock.ProductStock) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckProductExistence", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckProductExistence indicates an expected call of CheckProductExistence.
func (mr *MockStockMockRecorder) CheckProductExistence(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckProductExistence", reflect.TypeOf((*MockStock)(nil).CheckProductExistence), param)
}

// CreateStock mocks base method.
func (m *MockStock) CreateStock(param stock.StockInfo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStock", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStock indicates an expected call of CreateStock.
func (mr *MockStockMockRecorder) CreateStock(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStock", reflect.TypeOf((*MockStock)(nil).CreateStock), param)
}

// DecreaseAccountingAmount mocks base method.
func (m *MockStock) DecreaseAccountingAmount(amount, acID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseAccountingAmount", amount, acID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecreaseAccountingAmount indicates an expected call of DecreaseAccountingAmount.
func (mr *MockStockMockRecorder) DecreaseAccountingAmount(amount, acID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseAccountingAmount", reflect.TypeOf((*MockStock)(nil).DecreaseAccountingAmount), amount, acID)
}

// FindProductAmount mocks base method.
func (m *MockStock) FindProductAmount(param sales.Sales) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductAmount", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductAmount indicates an expected call of FindProductAmount.
func (mr *MockStockMockRecorder) FindProductAmount(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductAmount", reflect.TypeOf((*MockStock)(nil).FindProductAmount), param)
}

// FindProductStockCount mocks base method.
func (m *MockStock) FindProductStockCount(productID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductStockCount", productID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductStockCount indicates an expected call of FindProductStockCount.
func (mr *MockStockMockRecorder) FindProductStockCount(productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductStockCount", reflect.TypeOf((*MockStock)(nil).FindProductStockCount), productID)
}

// FindProductStockInfo mocks base method.
func (m *MockStock) FindProductStockInfo(productID, limit, offset int) ([]stock.ProductStockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductStockInfo", productID, limit, offset)
	ret0, _ := ret[0].([]stock.ProductStockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductStockInfo indicates an expected call of FindProductStockInfo.
func (mr *MockStockMockRecorder) FindProductStockInfo(productID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductStockInfo", reflect.TypeOf((*MockStock)(nil).FindProductStockInfo), productID, limit, offset)
}

// FindStockInfo mocks base method.
func (m *MockStock) FindStockInfo(limit, offset int) ([]stock.ProductStockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStockInfo", limit, offset)
	ret0, _ := ret[0].([]stock.ProductStockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStockInfo indicates an expected call of FindStockInfo.
func (mr *MockStockMockRecorder) FindStockInfo(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStockInfo", reflect.TypeOf((*MockStock)(nil).FindStockInfo), limit, offset)
}

// FindVariationAvailability mocks base method.
func (m *MockStock) FindVariationAvailability(id int) ([]stock.ProductStock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVariationAvailability", id)
	ret0, _ := ret[0].([]stock.ProductStock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVariationAvailability indicates an expected call of FindVariationAvailability.
func (mr *MockStockMockRecorder) FindVariationAvailability(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVariationAvailability", reflect.TypeOf((*MockStock)(nil).FindVariationAvailability), id)
}

// IncreaseAccountingAmount mocks base method.
func (m *MockStock) IncreaseAccountingAmount(amount, acID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseAccountingAmount", amount, acID)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseAccountingAmount indicates an expected call of IncreaseAccountingAmount.
func (mr *MockStockMockRecorder) IncreaseAccountingAmount(amount, acID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseAccountingAmount", reflect.TypeOf((*MockStock)(nil).IncreaseAccountingAmount), amount, acID)
}

// LoadStockCount mocks base method.
func (m *MockStock) LoadStockCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadStockCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadStockCount indicates an expected call of LoadStockCount.
func (mr *MockStockMockRecorder) LoadStockCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadStockCount", reflect.TypeOf((*MockStock)(nil).LoadStockCount))
}

// LoadStockList mocks base method.
func (m *MockStock) LoadStockList() ([]stock.StockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadStockList")
	ret0, _ := ret[0].([]stock.StockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadStockList indicates an expected call of LoadStockList.
func (mr *MockStockMockRecorder) LoadStockList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadStockList", reflect.TypeOf((*MockStock)(nil).LoadStockList))
}

// MockPrice is a mock of Price interface.
type MockPrice struct {
	ctrl     *gomock.Controller
	recorder *MockPriceMockRecorder
}

// MockPriceMockRecorder is the mock recorder for MockPrice.
type MockPriceMockRecorder struct {
	mock *MockPrice
}

// NewMockPrice creates a new mock instance.
func NewMockPrice(ctrl *gomock.Controller) *MockPrice {
	mock := &MockPrice{ctrl: ctrl}
	mock.recorder = &MockPriceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrice) EXPECT() *MockPriceMockRecorder {
	return m.recorder
}

// CreatePrice mocks base method.
func (m *MockPrice) CreatePrice(param price.Price) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePrice", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePrice indicates an expected call of CreatePrice.
func (mr *MockPriceMockRecorder) CreatePrice(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrice", reflect.TypeOf((*MockPrice)(nil).CreatePrice), param)
}

// MockSales is a mock of Sales interface.
type MockSales struct {
	ctrl     *gomock.Controller
	recorder *MockSalesMockRecorder
}

// MockSalesMockRecorder is the mock recorder for MockSales.
type MockSalesMockRecorder struct {
	mock *MockSales
}

// NewMockSales creates a new mock instance.
func NewMockSales(ctrl *gomock.Controller) *MockSales {
	mock := &MockSales{ctrl: ctrl}
	mock.recorder = &MockSalesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSales) EXPECT() *MockSalesMockRecorder {
	return m.recorder
}

// AddProductSale mocks base method.
func (m *MockSales) AddProductSale(param sales.Sales) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductSale", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProductSale indicates an expected call of AddProductSale.
func (mr *MockSalesMockRecorder) AddProductSale(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductSale", reflect.TypeOf((*MockSales)(nil).AddProductSale), param)
}

// FindSalesList mocks base method.
func (m *MockSales) FindSalesList(limit, offset int) ([]sales.SalesReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSalesList", limit, offset)
	ret0, _ := ret[0].([]sales.SalesReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSalesList indicates an expected call of FindSalesList.
func (mr *MockSalesMockRecorder) FindSalesList(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSalesList", reflect.TypeOf((*MockSales)(nil).FindSalesList), limit, offset)
}

// FindSalesReport mocks base method.
func (m *MockSales) FindSalesReport(param sales.SalesReportRequest) ([]sales.SalesReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSalesReport", param)
	ret0, _ := ret[0].([]sales.SalesReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSalesReport indicates an expected call of FindSalesReport.
func (mr *MockSalesMockRecorder) FindSalesReport(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSalesReport", reflect.TypeOf((*MockSales)(nil).FindSalesReport), param)
}

// FindSalesReportCount mocks base method.
func (m *MockSales) FindSalesReportCount(param sales.SalesReportRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSalesReportCount", param)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSalesReportCount indicates an expected call of FindSalesReportCount.
func (mr *MockSalesMockRecorder) FindSalesReportCount(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSalesReportCount", reflect.TypeOf((*MockSales)(nil).FindSalesReportCount), param)
}

// LoadSalesListCount mocks base method.
func (m *MockSales) LoadSalesListCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSalesListCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadSalesListCount indicates an expected call of LoadSalesListCount.
func (mr *MockSalesMockRecorder) LoadSalesListCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSalesListCount", reflect.TypeOf((*MockSales)(nil).LoadSalesListCount))
}
