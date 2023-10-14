<template>
    <h1>Edit Post</h1>
    <hr>

    <form v-on:submit.prevent="updatePost" method="PATCH">
        <!-- Name input -->
        <MDBInput type="text" label="title" id="title" v-model="post.title" wrapperClass="mb-4" />

        <!-- Message input -->
        <MDBTextarea label="Content" id="content" v-model="post.content" wrapperClass="mb-4" />
        <!-- Submit button -->
        <MDBBtn type="submit" color="primary" block class="mb-4"> Update </MDBBtn>
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
    name: 'EditPostPage',
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
    mounted() {
        this.loadPost();
    },
    methods: {
        async loadPost() {
            try {
                let res = await axios.get(`api/post/${this.$route.params.id}`);
                if (res.status != 200) {
                    console.log(res);
                } else {
                    this.post.title = res.data.title;
                    this.post.content = res.data.content;
                }
            } catch (error) {
                window.alert("There was an error loading the post")
                console.log(error)
            }

        },
        updatePost() {
            axios.patch(`api/post/${this.$route.params.id}`, this.post)
                .then(res => {
                    this.handleAuthSuccess(res);
                })
                .catch(err => {
                    this.handleRequestError(err);
                    this.cleanPostData();
                })
        },
        handleAuthSuccess(res) {
            window.alert("Successfully updated post!")
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