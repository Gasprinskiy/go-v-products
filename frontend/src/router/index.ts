import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/HomeView.vue'

import ProductListView from '../views/main/ProductListView.vue'
import StockView from '../views/main/StockView.vue'
import SalesView from '../views/main/SalesView.vue'

import ProductInfoView from '../views/child/ProductInfoView.vue'

const routes  = [
    {
        path: "/",
        component: HomeView,
        name: "home"
    },
    {
        path: "/product_list/:page?/:querry?",
        component: ProductListView,
        name: "product-list"
    },
    {
        path: "/product/:id",
        component: ProductInfoView,
        name: "product-info"
    },
    {
        path: "/stock_list/:page?",
        component: StockView,
        name: "stock-list"
    },
    {
        path: "/sales-list/:page?",
        component: SalesView,
        name: "sales-list"
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router