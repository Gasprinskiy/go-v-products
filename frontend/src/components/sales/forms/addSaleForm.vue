<template>
    <div class="add-sale-form">
        <column-group :gap="20">
            <column-group>
                <column-group>
                    <p>Продукт</p>
                    <el-select
                        v-model="formValues.product"
                        placeholder="Поиск по тэгам"
                        size="large"
                        filterable
                        remote
                        @change="emitVariationSearch"
                        :remote-method="emitProductSearch"
                        :class="{error: isProductValid}"
                    >
                        <el-option
                            v-for="product in props.productOptions"
                            :key="product.id"
                            :label="product.name"
                            :value="product.id"
                        />
                    </el-select>
                </column-group>
                <column-group v-if="isProductChosen">
                    <p>Форма</p>
                    <el-select
                        v-model="formValues.variation"
                        placeholder="Выберите форму товара"
                        size="large"
                        filterable
                        remote
                        :remote-method="emitProductSearch"
                        :class="{error: isVariationValid}"
                    >
                        <el-option
                            v-for="variation in props.variationOptions"
                            :key="variation.id"
                            :label="`${variation.type} ${variation.unit}`"
                            :value="variation.id"
                        />
                    </el-select>
                </column-group>
                <column-group v-if="isVariationChosen">
                    <p>Склад</p>
                    <el-select
                        v-model="formValues.stock"
                        placeholder="Выберите склад"
                        size="large"
                        @focus="emitGetStockList"
                        :class="{error: isStockValid}"
                    >
                        <el-option
                            v-for="stock in props.stockOptions"
                            :key="stock.id"
                            :label="`${stock.name} на ${stock.location}`"
                            :value="stock.id"
                        />
                    </el-select>
                </column-group>
                <column-group v-if="isVariationChosen">
                    <p>Количество</p>
                    <el-input-number
                        placeholder="Количество"
                        :controls="false"
                        size="large"
                        v-model="formValues.amount"
                        :class="{error: isAmountValid}"
                    />
                </column-group>
            </column-group>
            <column-group v-if="isVariationChosen">
                <el-button
                    type="primary"
                    @click="validateAndEmit"
                >
                    Продажа
                </el-button>
            </column-group>   
        </column-group>
    </div>
</template>

<script setup lang="ts">
// imports
import columnGroup from '../../templates/columnGroup.vue';
import useVuelidate from '@vuelidate/core'
import { required } from '@vuelidate/validators'
import { ref, computed, defineEmits, defineProps } from 'vue';
///////////////////////

// props
const props = defineProps({
    productOptions: {
        type: Array<any>,
        default: []
    },
    variationOptions: {
        type: Array<any>,
        default: []
    },
    stockOptions: {
        type: Array<any>,
        default: []
    }
})

// refs
const formValues = ref({
    product: null,
    variation: null,
    stock: null,
    amount: null
})
///////////////////////

// computed
const isProductValid = computed(() => v$.value.product.$dirty && v$.value.product.$invalid)
const isVariationValid = computed(() => v$.value.variation.$dirty && v$.value.variation.$invalid)
const isStockValid = computed(() => v$.value.stock.$dirty && v$.value.stock.$invalid)
const isAmountValid = computed(() => v$.value.amount.$dirty && v$.value.amount.$invalid)

const isProductChosen = computed(() => formValues.value.product !== null)
const isVariationChosen = computed(() => formValues.value.variation !== null)
///////////////

// validator
const validators = {
    product: {required},
    variation: {required},
    stock: {required},
    amount: {required},
}
const v$ = useVuelidate(validators, formValues)
///////////////////////

// emits
const emits = defineEmits(['sell-product', 'search-product-options', 'search-variation-options', 'get-stock-options'])

const emitProductSearch = (querry: string) => {
    if(querry) {
        emits('search-product-options', querry)
    }
}

const emitVariationSearch = (value: any) => {
    emits('search-variation-options', value)
}

const emitGetStockList = () => {
    emits('get-stock-options')
}

const validateAndEmit = () => {    
    v$.value.$validate()
    if(!v$.value.$invalid) {
        emits('sell-product', formValues.value)
    }
}
///////////////////////
</script>