// imports
import axios from "axios"
import { AxiosPromise } from 'axios';
import 
{ 
    getDataParam,
    createProductParam,
    createVariationParam,
    priceCreationParam,
    productAccountingParam,
    salesReportParam,
    stockCreateParam
} 
from "../entities/external"
/////////////////////////////////

// constatns
const url = "http://localhost:8080"
/////////////////////////////////

// api methods
const getProductInfo = async (id: number) : AxiosPromise => {
    return await axios.get(`${url}/product/${id}`)
}

const getStockInfo = async (param: getDataParam) : AxiosPromise => {    
    return await axios.get(`${url}/stock?product_id=${param.productID}&limit=${param.limit}&offset=${param.offset}`)
}

const getStockList = async () : AxiosPromise => {
    return await axios.get(`${url}/stock_list`)
}

const getProductList = async (param: getDataParam) : AxiosPromise =>{
    return await axios.get(`${url}/product_list?limit=${param.limit}&offset=${param.offset}&tag=${param.tags}`)
}

const createProduct = async (param: createProductParam) : AxiosPromise => {
    return await axios.post(`${url}/product/add`, param)
}

const createVariation = async (param: createVariationParam) : AxiosPromise => {
    return await axios.post(`${url}/product/add/variation`, param)
}

const createPrice = async (param: priceCreationParam) : AxiosPromise => {
    return await axios.post(`${url}/product/price`, param)
}

const createStock = async(param: stockCreateParam) : AxiosPromise => {
    return await axios.post(`${url}/stock/add`, param)
}

const addProductToStock = async (param: productAccountingParam) : AxiosPromise => {
    return await axios.post(`${url}/product/add/stock`, param)
}

const buyProduct = async (param: productAccountingParam) : AxiosPromise => {
    return await axios.post(`${url}/buy`, param)
}

const getSalesList = async(param: getDataParam) : AxiosPromise => {
    return await axios.get(`${url}/sales_list?limit=${param.limit}&offset=${param.offset}`)
}

const getSalesReport = async (param: salesReportParam) : AxiosPromise => {
    return await axios.post(`${url}/sales`, param)
}
/////////////////////////////////

// export
export default {
    getProductInfo,
    getStockInfo,
    getStockList,
    getProductList,
    createProduct,
    createVariation,
    createPrice,
    createStock,
    addProductToStock,
    buyProduct,
    getSalesList,
    getSalesReport
}
/////////////////////////////////