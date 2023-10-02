import { createApp } from 'vue'
import 'mdb-vue-ui-kit/css/mdb.min.css';

import App from './App.vue'
import router from './router.js'
import store from './store'
import "./interceptors/axios.js";


createApp(App)
.use(store)
.use(router)
.mount('#app')