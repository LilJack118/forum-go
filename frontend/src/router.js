import { createRouter, createWebHistory } from 'vue-router'
import store from './store/index.js';
import axios from 'axios';

// Import views
import FeedPage from './pages/FeedPage.vue';
import PostPage from './pages/PostPage.vue';
import EditPostPage from './pages/EditPostPage.vue';
import CreatePostPage from './pages/CreatePostPage.vue';
import LoginPage from './pages/LoginPage.vue';
import RegisterPage from './pages/RegisterPage.vue';
import NotFoundPage from './pages/NotFoundPage.vue';

const routes = [
  {
    path: '/',
    name: 'feed',
    component: FeedPage
  },
  {
    path: '/posts/:id',
    name: 'post-page',
    component: PostPage
  },
  {
    path: '/posts/:id/edit',
    name: 'post-page-edit',
    component: EditPostPage
  },
  {
    path: '/posts/create',
    name: 'post-page-create',
    component: CreatePostPage
  },
  {
    path: '/login',
    name: 'LoginPage',
    component: LoginPage
  },
  {
    path: '/register',
    name: 'RegisterPage',
    component: RegisterPage
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: NotFoundPage
  }
]

const router = createRouter({
  base: process.env.VUE_APP_BASE_URL,
  history: createWebHistory(process.env.BASE_URL),
  routes
})


export default router
