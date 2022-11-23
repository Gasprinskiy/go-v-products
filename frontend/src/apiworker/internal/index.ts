// imports
import { AxiosResponse } from "axios";
// import { type createProductParams } from '../entities/external';
import dataInternalView from '../entities/internal/data-view'
import dataExternalView from '../entities/external/data-view'
import external from "../external";
///////////////////

// response handler
const handleResponse = (method: (args: any, adit: any) => any, response: AxiosResponse, adit: any = null) : any => {
    if (response instanceof Error) {
        return response
    } 
    const data = method(response.data, adit)
    return data
}
////////////////////////////////////

// internal api methods
export const getProductInfo = async (id: number) => {
    const response = await external.getProductInfo(id)
    return await handleResponse(dataInternalView.handleProductWithVariation, response)      
}

export const getStockInfo = async (param: {limit : number, offset : number, productID: number}) => {
    const externalParam = dataExternalView.handleGetDataParam(param)
    const response = await external.getStockInfo(externalParam)    
    return handleResponse(dataInternalView.handleProductStockList, response, param.limit)  
}

export const getStockList = async () => {
    const response = await external.getStockList()
    return await handleResponse(dataInternalView.handleStockList, response) 
}

export const getProductList = async (param: {limit : number, offset : number, tags: string | string[]}) => {    
    const externalParam = dataExternalView.handleGetDataParam(param)    
    const response = await external.getProductList(externalParam)    
    return handleResponse(dataInternalView.handleProductList, response, param.limit)
}

export const createProduct = async (param: any) => {
    const externalParam = dataExternalView.hanldeCreateProductParams(param)
    const response = await external.createProduct(externalParam) 
    return response.data
}

export const createVariation = async (param: any) => {
    const externalParam = dataExternalView.handleCreateVariationParam(param)
    return await external.createVariation(externalParam)
}

export const createPrice = async (param: any) => {
    const externalParam = dataExternalView.handlePriceCreationParam(param)
    return await external.createPrice(externalParam)
}

export const createStock = async(param: any) => {
    const externalParam = dataExternalView.handeStockCreationParam(param)
    return await external.createStock(externalParam)
}

export const addProductToStock = async (param: any) => {
    const externalParam = dataExternalView.handleProductAccountingParam(param)
    return await external.addProductToStock(externalParam)
}

export const buyProduct = async (param: any) => {
    const externalParam = dataExternalView.handleProductAccountingParam(param)
    return await external.buyProduct(externalParam)
}

export const getSalesList = async(param: {limit : number, offset : number}) => {
    const externalParam = dataExternalView.handleGetDataParam(param)
    const response = await external.getSalesList(externalParam)
    return handleResponse(dataInternalView.handleSalesReportList, response, param.limit)  
}

export const getSalesReport = async (param: any) => {
    const externalParam = dataExternalView.handleSalesReportParam(param)
    const response = await external.getSalesReport(externalParam)
    return handleResponse(dataInternalView.handleSalesReportList, response, param.limit)
}
////////////////////////////////////

