import axios from 'axios';

export default {
  state() {
    return {
        user: null,
    }
  },
  getters: {
    getUser(state) {
        // Check if user is in localStorage
        if (state.user !== null && localStorage.getItem("user") === null){
            state.user = null;
        }
        
        return state.user;
    }
  },
  mutations: {
    setUser(state, user) {
        state.user = user;
    },
  },
  actions: {
    resetUserData(context, payload){
        context.commit("setUser", {"user": null});
        // remove user and tokens from localStorage
        localStorage.removeItem("user");
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");
        // remove authentication header
        delete axios.defaults.headers.common["Authorization"]
    },
    setUser(context, payload){
        localStorage.setItem("user", JSON.stringify(payload.user));
        context.commit("setUser", payload.user);
    },
    setAuthData(context, payload){
        // set user data, access and refresh tokens
        context.commit("setUser", payload.data.user);
        localStorage.setItem("user", JSON.stringify(payload.data.user));
        localStorage.setItem("access_token", payload.data.access_token);
        localStorage.setItem("refresh_token", payload.data.refresh_token);
        // set default authorization header for all requests
        axios.defaults.headers.common["Authorization"] = `Bearer ${payload.data.access_token}`;
    },
  },
}
