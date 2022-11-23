<template>
    <div class="product-view-container router-inner">
        <views-template 
            top-title="Склады"
        >
            <template #top-right>
                <el-button 
                    type="primary" 
                    :icon="Plus"
                    @click="showDialog = true" 
                />
            </template>
            <template #bar>
                <stock-search-bar
                    :product-options="productOptions"
                    @submit-search="resetAndGetList"
                    @reset-list="resetAndGetList"
                    @update-query="updateProductID"
                    @search-product="getProductOptions"
                />
            </template>
            <template #body>
                <stock-table
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
        title="Создать склад"
    >
        <create-stock-form
            v-if="showDialog"
            @create-stock="createNewStock"
        />
    </el-dialog>
</template>
  
<script setup lang="ts">
// imports
import viewsTemplate from '../../components/templates/viewsTemplate.vue';
import paginationBar from '../../components/bars/paginationBar.vue';
import stockTable from '../../components/stock/tables/stockTable.vue'
import createStockForm from '../../components/stock/forms/createStockForm.vue';
import stockSearchBar from '../../components/stock/bars/stockSearchBar.vue';

import { Plus } from '@element-plus/icons-vue';
import { getStockInfo, createStock, getProductList } from '../../apiworker/internal'; 
import { onBeforeMount, computed, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { notifyError, notifySuccess } from '../../composables/notify';
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
const showDialog = ref(false)
const listTotalPage = ref()
const filterParams = ref({
    limit: 5,
    offset: pageNum.value,
    productID: 0
})
const productOptions = ref()
///////////////////

// constatns
const limitOptions = [5, 10, 50]
///////////////////

// methods
const getList = async () => {
    listLoading.value = true
    try {
        const data = await getStockInfo(filterParams.value)
        list.value = data.list
        listTotalPage.value = data.totalPage
    } catch(e) {
        notifyError(e)
    }
    listLoading.value = false
}

const createNewStock = async (val: any) => {
    try {
        await createStock(val)
        await getList()
        notifySuccess("Склад создан")
        showDialog.value = false
    } catch (e) {
        notifyError(e)
    }
}

const resetAndGetList = async () => {
    filterParams.value.offset = 1
    router.push(`/stock_list/${filterParams.value.offset}`)
    await getList()
}

const getListByOffset = async (page: number | string) => {
    router.push(`/stock_list/${page}`)
    await getList()
}

const updateProductID = (val: string) =>{    
    filterParams.value.productID = Number(val)    
}

const getProductOptions = async (val: string) => {
    const params = {
        limit: 10,
        offset: 0,
        tags: val
    }
    const data = await getProductList(params)
    productOptions.value = data.list
}
///////////////////

// hooks
onBeforeMount(() => getList())
///////////////////
</script>
  