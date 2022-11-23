<template>
    <div class="add-variation-form">
        <column-group :gap="20">
            <column-group>
                <column-group>
                    <p>Название</p>
                    <el-input
                        v-model="formValues.name"
                        :class="{error: isNameValid}"
                        placeholder="Название"
                        size="large"
                    />
                </column-group>
                <column-group>
                    <p>Адрес</p>
                    <el-input
                        placeholder="Адрес"
                        size="large"
                        v-model="formValues.location"
                        :class="{error: isLocationValid}"
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
    name: null,
    location: null,
})
///////////////////////

// computed
const isNameValid = computed(() => v$.value.name.$dirty && v$.value.name.$invalid)
const isLocationValid = computed(() => v$.value.location.$dirty && v$.value.location.$invalid)
///////////////

// validator
const validators = {
    name: {required},
    location: {required},
}
const v$ = useVuelidate(validators, formValues)
///////////////////////

// emits
const emits = defineEmits(['create-stock'])

const validateAndEmit = () => {
    v$.value.$validate()
    if(!v$.value.$invalid) {
        emits('create-stock', formValues.value)
    }
}
///////////////////////
</script>