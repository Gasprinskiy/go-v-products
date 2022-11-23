<template>
    <div class="add-price-form">
        <column-group :gap="20">
            <column-group>
                <column-group>
                    <p>Цена</p>
                    <el-input-number
                        v-model="formValues.price"
                        placeholder="Цена"
                        size="large"
                        :controls="false"
                        :precision="2"
                        :class="{error: isPriceValid}"
                    />
                </column-group>
                <column-group>
                    <p>Активен с</p>
                    <el-date-picker
                        v-model="formValues.from"
                        size="large"
                        type="datetime"
                        placeholder="Активен с"
                        :class="{error: isFromValid}"
                    />
                </column-group>
                <column-group>
                    <p>Активен до</p>
                    <el-date-picker
                        v-model="formValues.till"
                        size="large"
                        type="datetime"
                        placeholder="Активен до"
                        :class="{error: isTillValid}"
                    />
                </column-group>
            </column-group>
            <column-group>
                <el-button
                    type="primary"
                    @click="validateAndEmit"
                >
                    Создать
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
import { ref, computed, defineEmits } from 'vue';
///////////////////////

// refs
const formValues = ref({
    price: null,
    from: null,
    till: null
})
///////////////////////

// computed
const isPriceValid = computed(() => v$.value.price.$dirty && v$.value.price.$invalid)
const isFromValid = computed(() => v$.value.from.$dirty && v$.value.from.$invalid)
const isTillValid = computed(() => v$.value.till.$dirty && v$.value.till.$invalid)
///////////////

// validator
const validators = {
    price: {required},
    from: {required},
    till: {required},
}
const v$ = useVuelidate(validators, formValues)
///////////////////////

// emits
const emits = defineEmits(['create-price'])

const validateAndEmit = () => {
    v$.value.$validate()
    if(!v$.value.$invalid) {
        emits('create-price', formValues.value)
    }
}
///////////////////////
</script>