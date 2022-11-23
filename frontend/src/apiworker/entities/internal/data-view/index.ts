import { salesReport } from './../index';

import { 
    productWithVariation,
    variationWithStock,
    variation,
    productStock,
    stock,
    product
} from ".."

const handleProductWithVariation = (data: any) : productWithVariation => {
    const variationList: Array<variationWithStock> = []

    data.variation_list.forEach((item : any) => {    
        variationList.push(handleVariationWithStock(item))
    })
    
    const result : productWithVariation = {
        product: {
            id: data.product_id,
            name: data.product_name,
            description: data.description,
            tags: data.tags,
        },
        variation_list: variationList, 
    }

    return result
}

const handleVariationWithStock = (data: any) : variationWithStock => {
    const stockAV : Array<productStock> = []
    data.stock_availability.forEach((stock : any) => {
        stockAV.push(handleproductStock(stock))
    });
    const result : variationWithStock = {
        id: data.variation_id,
        price: data.price,
        type: data.variation_type,
        unit: data.unit_type,
        stock_availability: stockAV
    }
    return result
}

const handleProductStockList = (data: any, limit: number) : {list: Array<productStock>, totalCount: number} => {
    const result : Array<productStock> = []
    data.product_stock_info.forEach((item : any) => {
        result.push(handleproductStock(item))
    });
    return {
        list: result,
        totalCount: handeTotalCount(data.total_count, limit)
    }
}

const handleStockList = (data: any) : Array<stock> => {
    const result : Array<stock> = []
    data.forEach((item : any) => {
        result.push(handlestockInfo(item))
    });
    return result
}

const handleProductList = (data: any, limit: number) : {list: Array<product>, totalCount: number} => {
    const result : Array<product> = []
    data.product_list.forEach((item: any) => {
        result.push(handleProduct(item))
    });
    return {
        list: result,
        totalCount: handeTotalCount(data.total_count, limit)
    }
}

const handleSalesReportList = (data: any, limit: number) : {list: Array<salesReport>, totalCount: number} => {
    const result : Array<salesReport> = []    
    data.sales_list.forEach((item: any) => {
        result.push(handleSalesReport(item))
    });
    return {
        list: result,
        totalCount: handeTotalCount(data.total_count, limit)
    }
}

const handleProduct = (data: any) : product => {
    const result : product = {
        id: data.product_id,
        name: data.product_name,
        description: data.description,
        tags: data.tags
    }
    return result
}

const handeTotalCount = (data: any, limit: number) => {
    return Math.ceil(data / limit)
}

const handlestockInfo = (data: any) : stock => {
    const result : stock = {
        id: data.stock_id,
        name: data.stock_name,
        location: data.location,
    }
    return result
}

const handleproductStock = (data: any) : productStock => {
    const result : productStock = {
        id: data.stock_id,
        name: data.stock_name,
        location: data.location,
        amount: data.amount | data.total_amount
    }
    return result
}

const handleSalesReport = (data: any) : salesReport => {
    const result: salesReport = {
        id: data.product_id,
        name: data.product_name,
        type: data.variation_type,
        unit: data.unit_type,
        stock: data.stock_name,
        soldDate: data.sold_date,
        amount: data.amount
    }
    return result
}

export default {
    handleProductWithVariation, 
    handleProductStockList,
    handleSalesReportList,
    handleProductList,
    handleStockList
}