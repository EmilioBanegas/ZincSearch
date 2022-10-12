import { createRouter, createWebHistory } from "vue-router";
import Search from "../views/Search.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Search",
      component: Search,
    },
    {
      path: "/email/:id",
      name: "Email",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/ShowEmail.vue"),
    },
    {
      path: '/:pathMatch(.*)*', 
      name: 'not-found',
      redirect:"/"
    }
  ],
});

//Agregando transición entre componentes
router.afterEach((to, from) => {
  if(to.path== from.path) {
    to.meta.transitionName = null;
    return;
  }
  const toDepth = to.path.split('/').length;
  const fromDepth = from.path.split('/').length;
  
  to.meta.transitionName = fromDepth <= toDepth ? 'slide-left' : 'slide-right'; 
})

export default router;
