import { createRouter, createWebHistory } from 'vue-router'
import store from './store/index.js';
import axios from 'axios';

// Import views
import FeedPage from './pages/FeedPage.vue';
import PostPage from './pages/PostPage.vue';
import EditPostPage from './pages/EditPostPage.vue';
import CreatePostPage from './pages/CreatePostPage.vue';
import AccountPage from './pages/AccountPage.vue';
import LoginPage from './pages/LoginPage.vue';
import RegisterPage from './pages/RegisterPage.vue';
import NotFoundPage from './pages/NotFoundPage.vue';

const routes = [
  {
    path: '/',
    meta:{requiresAuthentication:true},
    name: 'feed',
    component: FeedPage
  },
  {
    path: '/posts/:id',
    meta:{requiresAuthentication:true},
    name: 'post-page',
    component: PostPage
  },
  {
    path: '/posts/:id/edit',
    meta:{requiresAuthentication:true},
    name: 'post-page-edit',
    component: EditPostPage
  },
  {
    path: '/posts/create',
    meta:{requiresAuthentication:true},
    name: 'post-page-create',
    component: CreatePostPage
  },
  {
    path: '/account',
    meta:{requiresAuthentication:true},
    name: 'AccountPage',
    component: AccountPage
  },
  {
    path: '/login',
    meta:{requiresGuest:true},
    name: 'LoginPage',
    component: LoginPage
  },
  {
    path: '/register',
    meta:{requiresGuest:true},
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


router.beforeEach((to, from, next) => {
  function checkPermissions(){
    const requiresGuest = to.matched.some((x) => x.meta.requiresGuest);
    const requiresAuthentication = to.matched.some((x) => x.meta.requiresAuthentication)
    const isLoggedin = store.getters["getUser"] !== null;

    if (requiresGuest && isLoggedin) {
      next("/");
    } else if(requiresAuthentication && !isLoggedin) {
      next("/login");
    } else {
      next();
    }
  };

  // if token is invalid and refresh token is expired, user
  // credentials will be set to null
  if (localStorage.getItem("access_token") || localStorage.getItem("refresh_token")){
    axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("access_token")}`;
    axios.get("auth/token/verify")
    .then(response => {
      checkPermissions();
    }).catch(error => {
      checkPermissions();
    })
  }else{
    checkPermissions();
  }

});



export default router
