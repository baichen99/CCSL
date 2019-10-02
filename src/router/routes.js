import IndexLayout from "@/layout/IndexLayout.vue"; // Normal Layout
import AdminLayout from "@/layout/AdminLayout.vue"; // Admin Layout

/**
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

const routes = [
  {
    path: "/",
    component: IndexLayout,
    hidden: true,
    children: [
      {
        path: "/",
        name: "Home",
        component: () => import("@/views/home/Home.vue")
      },
      {
        path: "/team",
        component: () => import("@/views/about/Team"),
        name: "Team",
        meta: { title: "研究团队" }
      },
      {
        path: "/lexical-database",
        name: "LexicalDatabase",
        component: () => import("@/views/database/LexicalDatabase.vue"),
        meta: {
          roles: ["super", "admin", "user"],
          title: "国家通用手语比对语料库",
          auth: true
        }
      },
      {
        path: "/introduction",
        name: "Introduction",
        component: () => import("@/views/about/Introduction.vue"),
        meta: {
          title: "中心简介"
        }
      },
      {
        path: "/contact",
        name: "Contact",
        component: () => import("@/views/about/Contact.vue"),
        meta: {
          title: "联系我们"
        }
      },
      {
        path: "/learning-platform",
        name: "LearningPlatform",
        component: () => import("@/views/learning/LearningPlatform"),
        meta: {
          roles: ["super", "admin", "learner"],
          title: "国家通用手语学习平台",
          auth: true
        }
      }
    ]
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/login/Login"),
    hidden: true
  },
  {
    path: "/dashboard",
    component: AdminLayout,
    children: [
      {
        path: "/dashboard",
        name: "Dashboard",
        component: () => import("@/views/dashboard/Dashboard"),
        meta: {
          title: "控制中心",
          icon: "dashboard",
          auth: true,
          roles: ["super", "admin"]
        }
      }
    ]
  },
  {
    path: "/website",
    component: AdminLayout,
    redirect: "/website/carousel",
    name: "Website",
    meta: {
      title: "网站设置",
      icon: "website"
    },
    children: [
      {
        path: "carousel",
        name: "Carousel",
        component: () => import("@/views/dashboard/CarouselsSetting"),
        meta: {
          title: "轮播图片设置",
          icon: "pictures",
          auth: true,
          roles: ["super", "admin"]
        }
      },
      {
        path: "news",
        name: "News",
        component: () => import("@/views/dashboard/NewsSetting"),
        meta: {
          title: "网站新闻设置",
          icon: "news",
          auth: true,
          roles: ["super", "admin"]
        }
      }
    ]
  },
  {
    path: "/users",
    component: AdminLayout,
    children: [
      {
        path: "/users",
        name: "Users",
        component: () => import("@/views/dashboard/UserSetting"),
        meta: {
          title: "用户设置",
          icon: "peoples",
          auth: true,
          roles: ["super"]
        }
      }
    ]
  },
  {
    path: "/404",
    name: "404Error",
    component: () => import("@/views/error/404"),
    hidden: true
  },
  {
    path: "/401",
    name: "401Error",
    component: () => import("@/views/error/401"),
    hidden: true
  },
  { path: "*", redirect: "/404", hidden: true }
];

export default routes;
