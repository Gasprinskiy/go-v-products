<template>
    <div class="add-price-form">
        <column-group :gap="20">
            <column-group>
                <column-group>
                    <p>Склад</p>
                    <el-select
                        v-model="formValues.stock"
                        :class="{error: isStockValid}"
                        placeholder="Склад"
                        size="large"
                    >
                        <el-option
                            v-for="stock in props.stockOptions"
                            :key="stock.id"
                            :label="`${stock.name} на ${stock.location}`"
                            :value="stock.id"
                            size="large"
                        />
                    </el-select>
                </column-group>
                <column-group>
                    <p>Кол-во</p>
                    <el-input-number
                        placeholder="Количество"
                        :controls="false"
                        size="large"
                        v-model="formValues.amount"
                        :class="{error: isAmountValid}"
                    />
                </column-group>
            </column-group>
            <column-group>
                <el-button
                    type="primary"
                    @click="validateAndEmit"
                >
                    Добавить
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
    stockOptions: {
        type: Array<any>,
        default: []
    }
})
///////////////////////

// refs
const formValues = ref({
    stock: null,
    amount: null,
})
///////////////////////

// computed
const isStockValid = computed(() => v$.value.stock.$dirty && v$.value.stock.$invalid)
const isAmountValid = computed(() => v$.value.amount.$dirty && v$.value.amount.$invalid)
///////////////

// validator
const validators = {
    stock: {required},
    amount: {required},
}
const v$ = useVuelidate(validators, formValues)
///////////////////////

// emits
const emits = defineEmits(['add-to-stock'])

const validateAndEmit = () => {
    v$.value.$validate()
    if(!v$.value.$invalid) {
        emits('add-to-stock', formValues.value)
    }
}
///////////////////////
</script>