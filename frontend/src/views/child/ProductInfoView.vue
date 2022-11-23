<template>
    <div class="product-info-view router-inner" v-loading="loadingflag">
        <views-template
          :top-title="productTitle"
          :full-body="true"
        >
            <template #top-right>
                <el-button
                    type="primary"
                    :icon="Plus"
                    @click="openVariationCreation"
                >
                    Вариация
                </el-button>
            </template>   
            <template #body>
                <column-group :gap="35">
                    <product-info
                        :info="product.info"
                    />
                    <column-group>
                        <h2>Вариации</h2>
                        <product-variation-table
                            :variation-list="product.variationList"
                            @price-event="openPriceCreation"
                            @stock-event="openStockAddition"
                        />
                    </column-group>
                </column-group>
            </template>
        </views-template>
    </div>
    <el-dialog
        :title="chosenProductTitle"
        v-model="showDialog"
        width="25%"
        @closed="resetInner"
    >
        <create-price-form
            v-if="modalInner['addPriceForm']" 
            @create-price="createNewPrice"
        />
        <add-product-to-stock-form
            v-if="modalInner['addProductStockForm']"
            :stock-options="stockOptions"
            @add-to-stock="addToStock"
        />
        <create-variation-form
            v-if="modalInner['addVariationForm']"
            @create-variation="createNewVariation"
        />
    </el-dialog>
</template>
  
<script setup lang="ts">
// imports
import viewsTemplate from '../../components/templates/viewsTemplate.vue';
import columnGroup from '../../components/templates/columnGroup.vue';
import productInfo from '../../components/product/tables/productInfo.vue';

import productVariationTable from '../../components/product/tables/productVariationTable.vue';

import createPriceForm from '../../components/product/forms/createPriceForm.vue';
import addProductToStockForm from '../../components/product/forms/addProductToStockForm.vue';
import createVariationForm from '../../components/product/forms/createVariationForm.vue';

import { computed } from '@vue/reactivity';
import { useRoute } from 'vue-router';
import { getProductInfo, getStockList, createPrice, addProductToStock, createVariation } from '../../apiworker/internal';
import { notifyError, notifySuccess } from '../../composables/notify'
import { onBeforeMount, ref } from 'vue';
import { Plus } from '@element-plus/icons-vue';
import moment from 'moment';
//////////////////

// Vue use
const route = useRoute()
//////////////////

// computed
const productID = computed(() => Number(route.params.id))
const productTitle = computed(() => `Продукт #${productID.value}`)
const chosenProductTitle = computed(() => `${productTitle.value}` + (chosenRowID.value ? `, Вариация #${chosenRowID.value}` : ''))
//////////////////

// refs
const product = ref({
    info: {
        name: null,
        description: null,
        tags: null
    },
    variationList: []
})
const stockOptions = ref([])
const modalInner = ref({
    'addPriceForm': false,
    'addProductStockForm': false,
    'addVariationForm': false
})
const showDialog = ref(false)
const chosenRowID = ref()
const loadingflag = ref(false)
//////////////////

// methods
const getProduct = async () => {
    loadingflag.value = true
    try {
        const data = await getProductInfo(productID.value)
        product.value.info = data.product
        product.value.variationList = data.variation_list
    } catch (e) {
        notifyError(e)
    }
    loadingflag.value = false
}

const getStockOptions = async () => {
    loadingflag.value = true
    try {
        const data = await getStockList()
        stockOptions.value = data
    } catch (e) {
        notifyError(e)
    }
    loadingflag.value = false
}

const openVariationCreation = async () =>{
    chosenRowID.value = null
    modalInner.value['addVariationForm'] = true
    showDialog.value = true
}

const openPriceCreation = (id: number) => {
    modalInner.value['addPriceForm'] = true
    showDialog.value = true
    chosenRowID.value = id
}

const openStockAddition = async (id: number) => {
    if (stockOptions.value.length === 0) {
        await getStockOptions()
    }
    modalInner.value['addProductStockForm'] = true
    showDialog.value = true
    chosenRowID.value = id
}

const createNewVariation = async (values: any) => {
    loadingflag.value = true

    const newVariation = {
        id: productID.value,
        type: values.type,
        unit: values.unit
    }

    try {
        const response = await createVariation(newVariation)
        await getProduct()
        notifySuccess(`Вариация #${response.data} создана`)
        showDialog.value = false
    } catch (e) {
        notifyError(e)
    }

    loadingflag.value = false
}

const createNewPrice = async (values: any) => {
    loadingflag.value = true

    const newPrice = {
        id: chosenRowID.value,
        price: values.price,
        from: moment(values.from).format(),
        till: moment(values.till).format()
    }

    try {
        await createPrice(newPrice)
        await getProduct()
        notifySuccess("Цена создана")
        showDialog.value = false
    } catch (e) {
        notifyError(e)
    }

    loadingflag.value = false
}

const addToStock = async (values: any) => {
    loadingflag.value = true

    const newStockAccounting = {
        stock: values.stock,
        variation: chosenRowID.value,
        product: productID.value,
        amount: values.amount
    }

    try {
        await addProductToStock(newStockAccounting)
        await getProduct()
        notifySuccess("Продукт добавлен в склад")
        showDialog.value = false
    } catch (e) {
        notifyError(e)
    }

    loadingflag.value = false
}

const resetInner = () => {
    (Object.keys(modalInner.value) as (keyof typeof modalInner.value)[]).forEach((key) =>{
        modalInner.value[key] = false
    })
}
//////////////////

// hooks
onBeforeMount(() => getProduct())
//////////////////
</script>