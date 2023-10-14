<template>
    Posts List

    <div>
        <div v-for="post in posts" class="post card" v-on:click="openPost(post.id)">
            <div class="card-body text-start">
                <h5 class="card-title fw-bold">{{ post.title }}</h5>
                <p class="card-text">{{ post.content }}</p>
            </div>
        </div>
    </div>
</template>
  
<script>
import axios from 'axios';

export default {
    name: 'PostsList',
    properties: ["page", "limit"],
    components: {},
    data() {
        return {
            pageData: {},
        }
    },
    mounted() {
        this.loadPosts(this.page, this.limit);
    },
    computed: {
        posts() {
            if (this.pageData.posts) {
                return this.pageData.posts
            } else {
                return []
            }
        }
    },
    methods: {
        async loadPosts(page, limit) {
            let res = await axios.get(`api/posts?page=${page}&limit=${limit}`);
            if (res.status != 200) {
                console.log(res);
            } else {
                this.pageData = res.data;
            }
        },
        openPost(id) {
            this.$router.push({ name: 'post-page', params: { id: id } })
        }
    },
}
</script>

<style scoped>
.post {
    transition: all ease-in-out 0.2s;
}

.post:hover {
    cursor: pointer;
    transform: scale(1.01);
}
</style>