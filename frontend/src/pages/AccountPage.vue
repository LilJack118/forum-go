<template>
    <h1>Account</h1>
    <hr>

    <form v-on:submit.prevent="updateAccount" method="PATCH">
        <MDBInput type="text" label="First Name" id="first_name" v-model="user.first_name" wrapperClass="mb-4" />
        <MDBInput type="text" label="Last Name" id="last_name" v-model="user.last_name" wrapperClass="mb-4" />

        <!-- Message input -->
        <MDBInput type="email" label="Email" id="email" v-model="user.email" wrapperClass="mb-4" aria-disabled />

        <!-- Submit button -->
        <MDBBtn type="submit" color="primary" block class="mb-4"> Update </MDBBtn>
        <MDBBtn type="button" color="danger" block class="mb-4"> Delete </MDBBtn>
    </form>
</template>
  
<script>
import {
    MDBInput,
    MDBBtn,
} from "mdb-vue-ui-kit";
import axios from 'axios';
import { reactive } from 'vue';

export default {
    name: 'AccountPage',
    components: {
        MDBInput,
        MDBBtn,
    },
    data() {
        return {
            user: reactive({
                first_name: "",
                last_name: "",
                email: ""
            })
        }
    },
    mounted() {
        this.loadUserData();
    },
    methods: {
        async loadUserData() {
            try {
                let res = await axios.get('api/account');
                if (res.status != 200) {
                    window.alert("There was an error loading the user data")
                    console.log(res);
                } else {
                    this.$store.dispatch("setUser", { "user": res.data });
                    this.user.first_name = res.data.first_name;
                    this.user.last_name = res.data.last_name;
                    this.user.email = res.data.email;
                }
            } catch (error) {
                window.alert("There was an error loading the user data")
                console.log(error)
            }
        },
        updateAccount() {
            axios.patch('api/account', this.user)
                .then(res => {
                    this.handleUpdateSuccess(res);
                })
                .catch(err => {
                    this.handleRequestError(err);
                    this.cleanPostData();
                })
        },
        deleteAccount() {
            axios.delete('api/account')
                .then(res => {
                    window.alert("Successfully deleted account")
                    this.$store.dispatch("resetUserData")
                })
                .catch(err => {
                    this.handleRequestError(err);
                    this.cleanUserData();
                })
        },
        handleUpdateSuccess(res) {
            window.alert("Successfully updated account data!")
            this.loadUserData();
        },
        handleRequestError(err) {
            let errorMsg;

            if (err.response.data.detail) {
                errorMsg = err.response.data.detail;
            } else {
                errorMsg = "Try again later";
            }

            window.alert(errorMsg);
        },
        cleanUserData() {
            this.user.first_name = '';
            this.user.last_name = '';
            this.user.email = '';
        },
    }
}

</script>
  
<style scoped></style>