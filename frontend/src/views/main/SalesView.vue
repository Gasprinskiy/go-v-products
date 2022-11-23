<template>
    <div class="product-view-container router-inner">
        <views-template 
            top-title="Продажи"
        >
            <template #top-right>
                <el-button 
                    type="primary" 
                    :icon="Plus"
                    @click="showDialog = true" 
                />
            </template>
            <template #bar>
                <sales-report-search-bar
                    :stock-options="stockOptions"
                    @submit-search="handleSearchValues"
                    @reset-list="handleSearchValues"
                    @get-stock-options="getStockOptions"
                />
            </template>
            <template #body>
                <sales-report-table
                    @get-details="getProductDetails"
                    :loadingflag="listLoading"
                    :list="list"
                />
            </template>
            <template #bottom>
                <pagination-bar 
                    :list-total-page="listTotalPage"
                    :pagination-params="filterParams"
                    :limit-options="limitOptions"
                    @offset-change="getListByOffset"
                    @limit-change="resetAndGetList"
                />
            </template>
        </views-template>
    </div>
    <el-dialog
        v-model="showDialog"
        title="Проада товара"
    >
       <add-sale-form
            v-if="showDialog"
            :product-options="productOptions"
            :variation-options="variationOptions"
            :stock-options="stockOptions"
            @search-product-options="findProductOptions"
            @search-variation-options="findVariationOptions"
            @get-stock-options="getStockOptions"
            @sell-product="addNewBuy"
       />
    </el-dialog>
</template>
  
<script setup lang="ts">
// imports
import viewsTemplate from '../../components/templates/viewsTemplate.vue';
import paginationBar from '../../components/bars/paginationBar.vue';
import salesReportSearchBar from '../../components/sales/bars/salesReportSearchBar.vue';
import salesReportTable from '../../components/sales/tables/salesReportTable.vue'
import addSaleForm from '../../components/sales/forms/addSaleForm.vue';

import { Plus } from '@element-plus/icons-vue';
import { getSalesReport, getSalesList, getProductList, getProductInfo, getStockList, buyProduct } from '../../apiworker/internal'; 
import { onBeforeMount, computed, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { notifyError, notifySuccess } from '../../composables/notify';
import moment from 'moment';
////////////////////

// Vue use
const router = useRouter()
const route = useRoute()
///////////////////

// computed
const pageNum = computed(() => route.params.page ? Number(route.params.page) : 1)
///////////////////

// refs
const list = ref()
const listLoading = ref(false)
const searchMode = ref(false)
const showDialog = ref(false)
const listTotalPage = ref()
const filterParams = ref({
    start: "",
    end: "",
    stock: null,
    name: null,
    limit: 5,
    offset: pageNum.value
})

const productOptions = ref()
const variationOptions = ref()
const stockOptions = ref()
///////////////////

// constatns
const limitOptions = [5, 10, 50]
///////////////////

// methods
const getListBySearchMode = async () => {
    if (searchMode.value) {
        return await getReport()
    }
    return await getList()
}

const getList = async () => {
    listLoading.value = true
    try {
        const data = await getSalesList(filterParams.value)
        list.value = data.list
        listTotalPage.value = data.totalCount
    } catch(e) {
        notifyError(e)
    }
    listLoading.value = false
}

const getReport = async () => {
    listLoading.value = true
    try {
        const data = await getSalesReport(filterParams.value)
        list.value = data.list
        listTotalPage.value = data.totalCount
    } catch(e) {
        notifyError(e)
    }
    listLoading.value = false
}

const handleSearchValues = async (value: any, mode: boolean) => {
    searchMode.value = mode
    filterParams.value.start = moment(value.start).format()
    filterParams.value.end = moment(value.end).format()
    filterParams.value.stock = value.stock ? value.stock : null 
    filterParams.value.name = value.name ? value.name : null

    await getListBySearchMode()
}

const resetAndGetList = async () => {
    filterParams.value.offset = 1
    router.push(`/sales-list/${filterParams.value.offset}`)
    await getListBySearchMode()
}

const getListByOffset = async (page: number | string) => {
    router.push(`/sales-list/${page}`)
    await getListBySearchMode()
}

const findProductOptions = async (querry: any) => {
    if (isNaN(querry)) {
        const param = {
            limit: 10,
            offset: 1,
            tags: querry
        }
        const data = await getProductList(param)
        productOptions.value = data.list
        return
    }
}

const findVariationOptions = async (id: number) => {
    const data = await getProductInfo(id)            
    variationOptions.value = data.variation_list
}

const getStockOptions = async () => {
    const data = await getStockList()
    stockOptions.value = data
}

const addNewBuy = async (params: any) => {
    listLoading.value = true
    try {
        await buyProduct(params)
        await getListBySearchMode()
        notifySuccess("Товар продан")
        showDialog.value = false
    } catch(e) {
        notifyError(e)
    }
    listLoading.value = false
}

const getProductDetails = async (value: any) => {
    router.push({
        name: "product-list",
        params: {
            page: 1,
            querry: value.name
        }
    })
}
///////////////////

// hooks
onBeforeMount(()=> getList())
///////////////////
</script>

  