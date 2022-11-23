import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import router from './router'
import '@fontsource/open-sans';
import './styles/style.css'
import './styles/eplus-correction.css'
import 'element-plus/dist/index.css'

const app = createApp(App)

app
.use(router)
.use(ElementPlus)
.mount('#app')
