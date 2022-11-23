export type product = {
    id: number
    name: string
    description: string
    tags: string
}

export type variation = {
    id: number
    type: number
    unit: string
    price: number
}

export type productStock = {
    id: number
    name: string
    location: string
    amount: number
}

export type stock = {
    id: number
    name: string
    location: string
}

export type variationWithStock = {
    id: number
    type: number
    unit: string
    price: number
    stock_availability: Array<productStock>
}

export type productWithVariation = {
    product: product
    variation_list: Array<variationWithStock>
}

export type salesReport = {
    id: number
    name: string
    type: number
    unit: string
    stock: string
    soldDate: Date
    amount: number
}
