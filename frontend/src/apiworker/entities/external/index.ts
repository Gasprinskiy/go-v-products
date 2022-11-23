export type priceCreationParam = {
    variation_id: number
    price: number
    active_from: Date
    active_till: Date
}

export type stockCreateParam = {
    stock_name: string
    location: string
}

export type createProductParam = {
    product_name: string
    description: string
    tags: string
    variation_type: number
    unit_type: string
}

export type createVariationParam = {
    product_id: number
    variation_type: number
    unit_type: string
}

export type productAccountingParam = {
    stock_id: number
    product_id: number
    variation_id: number
    amount: number
}

export type salesReportParam = {
    start_date: Date
    end_date: Date
    limit: number
    offset: number
    product_name: string
    storage_id: number
}

export type getDataParam = {
    limit: number
    offset: number
    productID: number | 0
    tags: string | ""
}
