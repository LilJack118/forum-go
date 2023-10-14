<template>
    <h2>Feed</h2>

    <hr>

    <MDBTabs v-model="activeTab">
        <!-- Tabs navs -->
        <MDBTabNav tabsClasses="mb-3" class="justify-content-center">
            <MDBTabItem tabId="all-posts" href="all-posts">All Posts</MDBTabItem>
            <MDBTabItem tabId="my-posts" href="my-posts">My Posts</MDBTabItem>
        </MDBTabNav>

        <!-- Tabs content -->
        <MDBTabContent class="mt-4">
            <MDBTabPane tabId="all-posts">
                <PostsList></PostsList>
            </MDBTabPane>

            <MDBTabPane tabId="my-posts">
                <PostsList></PostsList>
            </MDBTabPane>
        </MDBTabContent>
    </MDBTabs>
</template>
  
<script>
import {
    MDBTabs,
    MDBTabNav,
    MDBTabContent,
    MDBTabItem,
    MDBTabPane,
} from 'mdb-vue-ui-kit';
import PostsList from '../components/feed/PostsList.vue'
import { ref } from 'vue';

export default {
    name: 'FeedPage',
    components: {
        MDBTabs,
        MDBTabNav,
        MDBTabContent,
        MDBTabItem,
        MDBTabPane,
        PostsList
    },
    data() {
        return {
            activeTab: ref('all-posts'),
            allPostsPage: 1,
            myPostsPage: 1,
        }
    }
}

</script>
  
<style scoped></style>