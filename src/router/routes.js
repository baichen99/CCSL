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
        component: () => import("@/views/home/Home")
      },
      {
        path: "/news-detail/:id",
        name: "NewsDetail",
        component: () => import("@/views/home/NewsDetail"),
        meta: { title: "新闻详情" }
      },
      {
        path: "/news-list/:column",
        name: "NewsList",
        component: () => import("@/views/home/NewsList")
      },
      {
        path: "/database/lexical-database",
        name: "LexicalDatabase",
        component: () => import("@/views/database/LexicalDatabase"),
        meta: {
          roles: [superUser, adminUser, user],
          title: "国家通用手语比对语料库",
          auth: true
        }
      },
      {
        path: "/about/introduction",
        name: "Introduction",
        component: () => import("@/views/about/Introduction"),
        meta: {
          title: "中心简介"
        }
      },
      {
        path: "/about/team",
        component: () => import("@/views/about/Team"),
        name: "Team",
        meta: { title: "研究团队" }
      },
      {
        path: "/about/contact",
        name: "Contact",
        component: () => import("@/views/about/Contact"),
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
    path: "/profile",
    component: AdminLayout,
    children: [
      {
        path: "/profile",
        name: "Profile",
        component: () => import("@/views/dashboard/Profile"),
        meta: {
          title: "个人中心",
          icon: "profile",
          auth: true
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
        component: () => import("@/views/dashboard/Users"),
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
    path: "/performers-setting",
    component: AdminLayout,
    children: [
      {
        path: "/performers-setting",
        name: "Performers",
        component: () => import("@/views/dashboard/Performers"),
        meta: {
          title: "被试管理",
          icon: "peoples",
          auth: true,
          roles: [superUser, adminUser]
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
      icon: "website",
      roles: [superUser, adminUser]
    },
    children: [
      {
        path: "carousels-setting",
        name: "Carousels",
        component: () => import("@/views/dashboard/Carousels"),
        meta: {
          title: "轮播图片管理",
          icon: "pictures",
          auth: true
        }
      },
      {
        path: "news-setting",
        name: "News",
        component: () => import("@/views/dashboard/News"),
        meta: {
          title: "网站新闻管理",
          icon: "news",
          auth: true
        }
      },
      {
        path: "members-setting",
        name: "Members",
        component: () => import("@/views/dashboard/Members"),
        meta: {
          title: "研究团队管理",
          icon: "member",
          auth: true
        }
      }
    ]
  },
  {
    path: "/lexical-database",
    component: AdminLayout,
    redirect: "/lexical-database/videos",
    name: "LexicalDatabaseSetting",
    meta: {
      title: "比对语料库管理",
      icon: "contrast",
      roles: [superUser, adminUser]
    },
    children: [
      {
        path: "words",
        name: "LexicalWords",
        component: () => import("@/views/dashboard/LexicalWords"),
        meta: {
          title: "词表管理",
          icon: "words",
          auth: true
        }
      },
      {
        path: "videos",
        name: "LexicalVideos",
        component: () => import("@/views/dashboard/LexicalVideos"),
        meta: {
          title: "视频管理",
          icon: "video",
          auth: true
        }
      },
      {
        path: "signs",
        name: "Signs",
        component: () => import("@/views/dashboard/Signs"),
        meta: {
          title: "手形管理",
          icon: "sign",
          auth: true
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
