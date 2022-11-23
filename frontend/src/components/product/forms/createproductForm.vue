<template>
    <div class="product-add-form">
        <column-group :gap="35">
            <column-group :gap="20">
                <h3>Информация о продукте</h3>
                <column-group>
                    <column-group>
                        <p>Название</p>
                        <el-input 
                            style="width: 100%;" 
                            placeholder="Название" 
                            size="large" 
                            v-model="formValues.name"
                            :class="{ error: isNameValid }" 
                        />
                    </column-group>
                    <column-group>
                        <p>Описание</p>
                        <el-input 
                            placeholder="Описание" 
                            type="textarea" 
                            size="large"
                            v-model="formValues.description" 
                        />
                    </column-group>
                    <column-group>
                        <p>Теги</p>
                        <el-input 
                            placeholder="Теги" 
                            size="large" 
                            v-model="formValues.tags"
                            :class="{ error: isTagsValid }" 
                        />
                    </column-group>
                </column-group>
            </column-group>
            <column-group :gap="20">
                <h3>Информация о вариации</h3>
                <div class="variation-info">
                    <column-group>
                        <column-group>
                            <p>Форма выпуска</p>
                            <el-input-number 
                                :precision="2" 
                                placeholder="Форма выпуска" 
                                :controls="false" 
                                size="large"
                                v-model="formValues.type" 
                                :class="{ error: isTypeValid }" 
                            />
                        </column-group>
                        <column-group>
                            <p>Ед. изм</p>
                            <el-input 
                                placeholder="Единица измерения" 
                                size="large" 
                                v-model="formValues.unit"
                                :class="{ error: isUnitValid }" 
                            />
                        </column-group>
                    </column-group>
                </div>
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
import { required, minLength, maxLength } from '@vuelidate/validators'
import { ref, defineEmits } from 'vue';
import { computed } from '@vue/reactivity';
///////////////

// refs
const formValues = ref({
    name: null,
    description: null,
    tags: null,
    type: null,
    unit: null
})
///////////////

// computed
const isNameValid = computed(() => v$.value.name.$dirty && v$.value.name.$invalid)
const isTagsValid = computed(() => v$.value.tags.$dirty && v$.value.tags.$invalid)
const isTypeValid = computed(() => v$.value.type.$dirty && v$.value.type.$invalid)
const isUnitValid = computed(() => v$.value.unit.$dirty && v$.value.unit.$invalid)
///////////////

// validator
const validators = {
    name: { required, minLength: minLength(3), maxLength: maxLength(80) },
    tags: { required, minLength: minLength(3), maxLength: maxLength(100) },
    type: { required },
    unit: { required }
}
const v$ = useVuelidate(validators, formValues)
///////////////

// emits
const emits = defineEmits(['create-product'])
const validateAndEmit = () => {
    v$.value.$validate()
    if (!v$.value.$invalid) {
        emits('create-product', formValues.value)
    }
}

</script>
  
  
<style scoped>
/* .variation-info {
        display: flex;
     } */
</style>