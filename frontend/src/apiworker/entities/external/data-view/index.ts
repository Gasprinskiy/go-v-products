import { stockCreateParam } from './../index';
import { 
    createProductParam,
    createVariationParam,
    priceCreationParam,
    productAccountingParam,
    salesReportParam,
    getDataParam
} from '../'

const handleGetDataParam = (data: any) : getDataParam => {    
    const param : getDataParam = {
        limit: data.limit | 0,
        offset: data.offset | 0,
        productID: data?.productID,
        tags: data?.tags
    }
    return param
}

const hanldeCreateProductParams = (data: any) : createProductParam => {
    const param : createProductParam = {
        product_name: data.name,
        description: data.description,
        tags: data.tags,
        variation_type: data.type,
        unit_type: data.unit
    }
    return param
}

const handleCreateVariationParam = (data: any) : createVariationParam => {
    const param : createVariationParam = {
        product_id: data.id,
        variation_type: data.type,
        unit_type: data.unit
    }
    return param
}

const handlePriceCreationParam = (data: any) : priceCreationParam => {
    const param : priceCreationParam = {
        variation_id: data.id,
        price: data.price,
        active_from: data.from,
        active_till: data.till
    }
    return param
}

const handeStockCreationParam = (data: any) : stockCreateParam => {
    const param : stockCreateParam = {
        stock_name: data.name,
        location: data.location
    }
    return param
}

const handleProductAccountingParam = (data: any) : productAccountingParam => {
    const param : productAccountingParam = {
        stock_id: data.stock,
        product_id: data.product,
        variation_id: data.variation,
        amount: data.amount
    }
    return param
}

const handleSalesReportParam = (data: any) : salesReportParam => {
    const param : salesReportParam = {
        start_date: data.start,
        end_date: data.end,
        limit: data.limit,
        offset: data.offset,
        product_name: data.name,
        storage_id: data.stock
    }
    return param
}

export default {
    handleGetDataParam,
    hanldeCreateProductParams,
    handleCreateVariationParam,
    handlePriceCreationParam,
    handeStockCreationParam,
    handleProductAccountingParam,
    handleSalesReportParam
}