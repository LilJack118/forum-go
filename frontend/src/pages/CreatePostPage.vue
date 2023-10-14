<template>
    <h1>Create Post</h1>
    <hr>

    <form v-on:submit.prevent="createPost" method="POST">
        <!-- Name input -->
        <MDBInput type="text" label="title" id="title" v-model="post.title" wrapperClass="mb-4" />

        <!-- Message input -->
        <MDBTextarea label="Content" id="content" v-model="post.content" wrapperClass="mb-4" />
        <!-- Submit button -->
        <MDBBtn type="submit" color="primary" block class="mb-4"> Create </MDBBtn>
    </form>
</template>
  
<script>
import {
    MDBInput,
    MDBBtn,
    MDBTextarea
} from "mdb-vue-ui-kit";
import axios from 'axios';
import { reactive } from 'vue';

export default {
    name: 'CreatePostPage',
    components: {
        MDBInput,
        MDBBtn,
        MDBTextarea
    },
    data() {
        return {
            post: reactive({
                title: "",
                content: "",
            })
        }
    },
    methods: {
        createPost() {
            axios.post('api/post', this.post)
                .then(res => {
                    this.handleAuthSuccess(res);
                })
                .catch(err => {
                    this.handleRequestError(err);
                    this.cleanPostData();
                })
        },
        handleAuthSuccess(res) {
            window.alert("Successfully created new post!")
            this.cleanPostData();
            this.$router.push({ name: "feed" })
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
        cleanPostData() {
            this.post.title = '';
            this.post.content = '';
        },
    }
}

</script>
  
<style scoped></style>