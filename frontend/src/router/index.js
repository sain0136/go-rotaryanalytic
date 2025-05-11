import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import { useAuth } from "../composables/useAuth";

const { isLoggedIn } = useAuth();

const routes = [
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/",
    name: "home",
    component: Home,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  console.log("checking auth", isLoggedIn.value);
  if (to.meta.requiresAuth && !isLoggedIn.value) {
    next({ name: "login" });
  } else if (to.name === "login" && isLoggedIn.value) {
    next({ name: "home" });
  } else {
    next();
  }
});

export default router;
