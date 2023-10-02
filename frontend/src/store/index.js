import { createStore } from 'vuex'

// import modules
import AuthModule from './modules/auth';

export default createStore({
  modules: {
    auth: AuthModule,
  }
})
