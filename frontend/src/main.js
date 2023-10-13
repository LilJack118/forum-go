import { createApp } from 'vue'
import 'mdb-vue-ui-kit/css/mdb.min.css';

import router from './router.js'
import store from './store'
import "./interceptors/axios.js";

import App from './App.vue'


createApp(App)
.use(store)
.use(router)
.mount('#app')