<template>
    Posts List

    <div>
        <div v-for="post in posts">
            {{ post }}
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
    methods: {
        async loadPosts(page, limit) {
            let res = await axios.get(`api/posts?page=${page}&limit=${limit}`);
            if (res.status != 200) {
                console.log(res);
            } else {
                this.pageData = res.data;
            }
        }
    },
    computed: {
        posts() {
            if (this.pageData.posts) {
                return this.pageData.posts
            } else {
                return []
            }
        }
    }
}
</script>

<style scoped></style>