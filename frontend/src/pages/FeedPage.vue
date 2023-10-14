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
                All Posts
            </MDBTabPane>

            <MDBTabPane tabId="my-posts">
                My Posts
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
import { ref } from 'vue';

export default {
    name: 'FeedPage',
    components: {
        MDBTabs,
        MDBTabNav,
        MDBTabContent,
        MDBTabItem,
        MDBTabPane,
    },
    data() {
        return {
            activeTab: ref('all-posts')
        }
    }
}

</script>
  
<style scoped></style>