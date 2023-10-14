<template>
    <h1 class="fw-bold">{{ post.title }}</h1>
    <p class="">Created {{ formatTimestamp(post.created_at) }}</p>
    <hr>
    <p>{{ post.content }}</p>
</template>
  
<script>
import axios from 'axios';
import moment from 'moment';

export default {
    name: 'PostPage',
    components: {},
    data() {
        return {
            post: {}
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
                    this.post = res.data;
                }
            } catch (error) {
                window.alert("There was an error loading the post")
                console.log(error)
            }

        },
        formatTimestamp(timestamp) {
            return moment(timestamp).fromNow()
        }
    }
}

</script>
  
<style scoped></style>