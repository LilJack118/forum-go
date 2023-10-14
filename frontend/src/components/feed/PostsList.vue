<template>
    <div class="mt-2 mb-5">
        <div class="d-flex" v-for="post in posts">
            <div class="post my-3 card w-100" v-on:click="openPost(post.id)">
                <div class="card-body text-start">
                    <h5 class="card-title fw-bold">{{ post.title }}</h5>
                    <p class="card-text">{{ post.content }}</p>
                    <p class="p-0 m-0" style="font-size: 10px">{{ post.created_at }}</p>

                </div>
            </div>
            <div v-if="load == 'my'" class="my-auto">
                <button class="btn btn-sm bg-primary mb-2 text-white fw-bold">Edit</button>
                <button class="btn btn-sm bg-danger text-white fw-bold">Delete</button>
            </div>
        </div>
    </div>
</template>
  
<script>
import axios from 'axios';

export default {
    name: 'PostsList',
    props: ['load', 'active'],
    components: {},
    data() {
        return {
            page: 1,
            limit: 12,
            loadedAll: false,
            loading: false,
            posts: [],
        }
    },
    mounted() {
        this.loadPosts();
        window.addEventListener("scroll", this.handleScroll)
    },
    unmounted() {
        window.removeEventListener("scroll", this.handleScroll)
    },
    methods: {
        handleScroll(event) {
            if (!this.active || this.loading || this.loadedAll) return;

            let scrollHeight = Math.max(
                document.body.scrollHeight, document.documentElement.scrollHeight,
                document.body.offsetHeight, document.documentElement.offsetHeight,
                document.body.clientHeight, document.documentElement.clientHeight
            );
            let currentScroll = window.scrollY + window.innerHeight;

            let modifier = 300; // at 300 px from bottom start loading next page
            if (currentScroll + modifier > scrollHeight) {
                this.loadPosts();
            }
        },
        async loadPosts() {
            this.loading = true;

            try {
                let endpoint = this.load == 'all' ? 'api/posts' : 'api/posts/my';
                let res = await axios.get(`${endpoint}?page=${this.page}&limit=${this.limit}`);
                if (res.status != 200) {
                    console.log(res);
                } else {
                    if (res.data.posts != null) {
                        this.posts.push(...res.data.posts);
                        this.page += 1;
                    } else {
                        this.loadedAll = true;
                    }
                }
            } catch (error) {
                console.log(error)
            }

            this.loading = false;
        },
        openPost(id) {
            this.$router.push({ name: 'post-page', params: { id: id } })
        },
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