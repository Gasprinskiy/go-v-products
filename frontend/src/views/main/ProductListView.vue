<template>
    <div class="product-view-container router-inner">
        <views-template
            top-title="Продукты"
        >
            <template #top-right>
                <el-button
                    @click="showDialog = true"
                    type="primary"
                    :icon="Plus"
                />
            </template>
            <template #bar>
                <searc-bar
                    placeholder="Поиск"
                    :query="filterParams.tags"
                    @submit-search="getListByQuerryParams"
                    @reset-list="resetAndGetList"
                    @update-query="updateTags"
                />
            </template>
            <template #body>
                <product-table
                    :loadingflag="listLoading"
                    :list="list"
                    @row-clicked="goToProduct"
                />
            </template>
            <template #bottom>
                <pagination-bar
                    :limit-options="limitOptions"
                    :list-total-page="listTotalPage"
                    :pagination-params="filterParams"
                    @offset-change="getListByOffset"
                    @limit-change="resetAndGetList"
                />
            </template>
        </views-template>
    </div>
    <el-dialog
        v-model="showDialog"
        title="Создать продукт"
    >
        <create-product-form
            v-if="showDialog"
            @create-product="createNewProduct"
        />
    </el-dialog>
</template>
  
<script setup lang="ts">
// imports
import viewsTemplate from '../../components/templates/viewsTemplate.vue';
import paginationBar from '../../components/bars/paginationBar.vue';
import productTable from '../../components/product/tables/productTable.vue';
import searcBar from '../../components/product/bars/searcBar.vue';
import createProductForm from '../../components/product/forms/createproductForm.vue';

import { Plus } from '@element-plus/icons-vue';

import { notifyError, notifySuccess } from '../../composables/notify'
import { getProductList, createProduct } from '../../apiworker/internal';

import { onBeforeMount, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { computed } from '@vue/reactivity';
///////////////////

// Vue use
const router = useRouter()
const route = useRoute()
///////////////////

// computed
const pageNum = computed(() => route.params.page ? Number(route.params.page) : 1)
const routeQuerry = computed(() => route.params.querry ? route.params.querry : "")
///////////////////

// refs
const list = ref()
const listLoading = ref(false)
const showDialog = ref(false)
const listTotalPage = ref()
const filterParams = ref({
    limit: 5,
    offset: pageNum.value,
    tags: routeQuerry.value
})
///////////////////

// constatns
const limitOptions = [5, 10, 50]
///////////////////

// methods
const getList = async () => {
    listLoading.value = true
    try {        
        const data = await getProductList(filterParams.value)      
        console.log(data);
          
        list.value = data.list
        listTotalPage.value = data.totalCount
    } catch (e) {
        notifyError(e)
    }
    listLoading.value = false
}

const createNewProduct = async (val: any) => {
    try {
        await createProduct(val)
        await getList()
        notifySuccess("Продукт добавлен")
        showDialog.value = false
    } catch (e) {
        notifyError(e)
    }
}

const resetAndGetList = async () => {
    filterParams.value.offset = 1
    filterParams.value.tags = ""
    router.push(`/product_list/${filterParams.value.offset}`)
    await getList()
}

const getListByQuerryParams = async () => {
    filterParams.value.offset = 1
    router.push({
        name: "product-list",
        params: {
            page: filterParams.value.offset,
            querry: filterParams.value.tags
        }
    })
    await getList()
}

const getListByOffset = async (page: number | string) => {
    router.push(`/product_list/${page}`)
    await getList()
}

const updateTags = (val: string) =>{    
    filterParams.value.tags = val    
}

const goToProduct = (id: number) => {
    router.push(`/product/${id}`)
}
///////////////////

// hooks
onBeforeMount(()=> getList())
///////////////////
</script>

  