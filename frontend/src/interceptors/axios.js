import axios from 'axios';
import store from '../store/index.js';

axios.defaults.baseURL = process.env.VUE_APP_BASE_API_URL;
// if access token in local storage use it as default authorization header
if (localStorage.getItem("access_token")){
    axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("access_token")}`;
}

// this variable is used to prevent infinite loop if refresh request will
// return 401
let refresh = false;
// interceptors are used to handle responses or requests
// before they are handled by then or catch
axios.interceptors.response.use(resp => {
    // if next respose after refresh request is 200
    if (resp.status == 200){
        refresh = false;
    }
    return resp
}, async error =>{
    // after every request if error occurs check if status code is 401,
    // if so try to refresh access token
    if (error.response.status === 401 && !refresh && localStorage.getItem("refresh_token")){
        refresh = true;
        const {status, data}  = await axios.post('auth/token/refresh', {
            "refresh_token": localStorage.getItem("refresh_token") || "invalid"
        });
        
        if (status === 200){
            localStorage.setItem('access_token', data.access_token);
            // update default authorization header for all requests
            axios.defaults.headers.common["Authorization"] = `Bearer ${data.access_token}`;
            error.config.headers["Authorization"] = `Bearer ${data.access_token}`;
            // do previous request
            return axios(error.config);
        }
    }else if(error.response.status === 401){
        // if second two 401 after each other, refresh token is invalid or expired
        // so reset user data
        store.dispatch("resetUserData");
    };
    return Promise.reject(error);
})