import { 
    Goods,
    Coin,
    Money
} from '@element-plus/icons-vue'

export const mainRoutes  = [
    {
        path: "/product_list",
        icon: Goods,
        title: "Продукты",
        name: "product-list"
    },
    {
        path: "/stock_list",
        icon: Coin,
        title: "Склады",
        name: "stock-list"
    },
    {
        path: "/sales-list",
        icon: Money,
        title: "Продажи",
        name: "sales-list"
    }
]

