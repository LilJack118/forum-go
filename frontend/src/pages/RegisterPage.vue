<template>
    <h1>Sign Up</h1>

    <div v-if="errorMsg" class="w-100 bg-danger text-white mb-5 py-2 px-3 border-rounded">
        <div>
            <h4 class="fw-bold m-0 mb-2">Oops! Something went wrong</h4>
            <p class="fw-bold m-0">{{ errorMsg }}</p>
        </div>
    </div>

    <form v-on:submit.prevent="signUp" method="POST">
        <MDBInput type="text" label="First Name" id="form1FirstName" v-model="userData.first_name" wrapperClass="mb-4"
            required />
        <MDBInput type="text" label="Last Name" id="form1LastName" v-model="userData.last_name" wrapperClass="mb-4"
            required />
        <MDBInput type="email" label="Email address" id="form1Email" v-model="userData.email" wrapperClass="mb-4"
            required />
        <MDBInput type="password" label="Password" id="form1Password" v-model="userData.password" wrapperClass="mb-4"
            required />

        <MDBBtn color="primary" type="submit" block>Sign Up</MDBBtn>
    </form>

    <div class="mt-3">
        <p>
            <span>Already have an account? </span>
            <router-link to="/login" class="fw-bold text-yellow">
                <span>Login</span>
            </router-link>
        </p>
    </div>
</template>
  
<script>
import {
    MDBRow,
    MDBCol,
    MDBInput,
    MDBCheckbox,
    MDBBtn,
} from "mdb-vue-ui-kit";
import axios from 'axios';
import { reactive } from 'vue';

export default {
    name: 'RegisterPage',
    components: {
        MDBRow,
        MDBCol,
        MDBInput,
        MDBCheckbox,
        MDBBtn,
    },
    data() {
        return {
            userData: reactive({
                first_name: '',
                last_name: '',
                email: '',
                password: '',
            }),
            errorMsg: ''
        };
    },
    watch: {
        errorMsg(newVal, oldVal) {
            if (!newVal) return;

            setTimeout(() => {
                this.errorMsg = '';
            }, 10000);
        }
    },
    methods: {
        signUp() {
            axios.post('auth/register', this.userData)
                .then(res => {
                    this.handleAuthSuccess(res);
                })
                .catch(err => {
                    this.handleRequestError(err);
                    this.cleanUserData();
                })
        },
        handleAuthSuccess(res) {
            // set tokens and user data in local storage
            this.$store.dispatch("setAuthData", { "data": res.data });
            this.cleanUserData();
        },
        handleRequestError(err) {
            if (err.response.data.detail) {
                this.errorMsg = err.response.data.detail;
            } else {
                this.errorMsg = "Try again later";
            }
        },
        cleanUserData() {
            this.userData.email = '';
            this.userData.password = '';
        },
    }
}

</script>
  
<style scoped></style>