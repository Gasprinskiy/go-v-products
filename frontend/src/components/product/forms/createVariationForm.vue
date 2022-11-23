<template>
    <div class="add-variation-form">
        <column-group :gap="20">
            <column-group>
                <column-group>
                    <p>Форма</p>
                    <el-input-number
                        :precision="2"
                        :controls="false"
                        v-model="formValues.type"
                        :class="{error: isTypeValid}"
                        placeholder="Форма"
                        size="large"
                    />
                </column-group>
                <column-group>
                    <p>Ед. изм</p>
                    <el-input
                        placeholder="Единица измерения"
                        :controls="false"
                        size="large"
                        v-model="formValues.unit"
                        :class="{error: isUnitValid}"
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
    type: null,
    unit: null,
})
///////////////////////

// computed
const isTypeValid = computed(() => v$.value.type.$dirty && v$.value.type.$invalid)
const isUnitValid = computed(() => v$.value.unit.$dirty && v$.value.unit.$invalid)
///////////////

// validator
const validators = {
    type: {required},
    unit: {required},
}
const v$ = useVuelidate(validators, formValues)
///////////////////////

// emits
const emits = defineEmits(['create-variation'])

const validateAndEmit = () => {
    v$.value.$validate()
    if(!v$.value.$invalid) {
        emits('create-variation', formValues.value)
    }
}
///////////////////////
</script>