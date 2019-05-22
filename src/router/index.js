import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);
import Index from "@/views/Index";
import Admin from "@/layout/Admin.vue"; // Admin Layout

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

export const routes = [
  {
    path: "/",
    component: Index,
    hidden: true,
    children: [
      {
        path: "/",
        redirect: "/dictionary",
        // name: "Index",
        // component: () => import("@/views/home/Home.vue"),
        hidden: true
      },
      {
        path: "/dictionary",
        name: "Dictionary",
        component: () => import("@/views/dictionary/Dictionary.vue"),
        hidden: true
      }
    ]
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/login/index"),
    hidden: true
  },
  {
    path: "/404",
    name: "404Error",
    component: () => import("@/views/404"),
    hidden: true
  },
  {
    path: "/dashboard",
    component: Admin,
    // redirect: "/dashboard",
    children: [
      {
        path: "/",
        name: "Dashboard",
        component: () => import("@/views/dashboard/index"),
        meta: { title: "Dashboard", icon: "dashboard" }
      }
    ]
  },
  {
    path: "/example",
    component: Admin,
    redirect: "/example/table",
    name: "Example",
    meta: { title: "Example", icon: "example", requireAuth: true },
    children: [
      {
        path: "table",
        name: "Table",
        component: () => import("@/views/table/index"),
        meta: { title: "Table", icon: "table" }
      },
      {
        path: "tree",
        name: "Tree",
        component: () => import("@/views/tree/index"),
        meta: { title: "Tree", icon: "tree" }
      }
    ]
  },
  { path: "*", redirect: "/404", hidden: true }
];

const router = new Router({
  mode: "history",
  scrollBehavior: () => ({ y: 0 }),
  routes: routes
});

export default router;
