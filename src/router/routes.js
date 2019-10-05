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

const superUser = "super";
const adminUser = "admin";
const learner = "learner";
const user = "user";

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
          roles: [superUser, adminUser, user],
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
          roles: [superUser, adminUser, learner],
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
          roles: [superUser, adminUser]
        }
      }
    ]
  },
  {
    path: "/users-setting",
    component: AdminLayout,
    children: [
      {
        path: "/users-setting",
        name: "Users",
        component: () => import("@/views/dashboard/AllUsers"),
        meta: {
          title: "用户管理",
          icon: "user",
          auth: true,
          roles: [superUser]
        }
      }
    ]
  },
  {
    path: "/website",
    component: AdminLayout,
    redirect: "/website/carousels-setting",
    name: "Website",
    meta: {
      title: "网站管理",
      icon: "website"
    },
    children: [
      {
        path: "carousels-setting",
        name: "Carousels",
        component: () => import("@/views/dashboard/AllCarousels"),
        meta: {
          title: "轮播图片管理",
          icon: "pictures",
          auth: true,
          roles: [superUser, adminUser]
        }
      },
      {
        path: "news-setting",
        name: "News",
        component: () => import("@/views/dashboard/AllNews"),
        meta: {
          title: "网站新闻管理",
          icon: "news",
          auth: true,
          roles: [superUser, adminUser]
        }
      }
    ]
  },
  {
    path: "/lexical-database-setting",
    component: AdminLayout,
    redirect: "/lexical-database/videos",
    name: "LexicalDatabaseSetting",
    meta: {
      title: "比对语料库管理",
      icon: "contrast"
    },
    children: [
      {
        path: "videos",
        name: "LexicalVideos",
        component: () => import("@/views/dashboard/AllLexicalVideos"),
        meta: {
          title: "视频管理",
          icon: "video",
          auth: true,
          roles: [superUser, adminUser]
        }
      },
      {
        path: "signs",
        name: "Signs",
        component: () => import("@/views/dashboard/AllSigns"),
        meta: {
          title: "手形管理",
          icon: "sign",
          auth: true,
          roles: [superUser, adminUser]
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
