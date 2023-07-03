import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    redirect: "/home",
  },
  {
    path: "/home",
    name: "home",
    component: () => import("@/pages/home/index.vue"),
  },
  {
    path: "/blog",
    name: "blog",
    component: () => import("@/pages/blog/index.vue"),
  },
  {
    path: "/admin",
    name: "admin",
    component: () => import("@/pages/admin/index.vue"),
  },
  {
    path: "/:catchAll(.*)",
    meta: {
      closeTab: true,
    },
    component: () => import("@/pages/error/index.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
