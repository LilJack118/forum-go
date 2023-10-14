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
            posts_num: 0,
        }
    },
    computed: {
        async posts() {
            let res = await axios.get(`api/posts?page=${this.page}&limit=${this.limit}`);
            if (res.statusCode != 200) {
                console.log(res);
                return []
            }


            return []

        }
    }
}
</script>

<style scoped></style>