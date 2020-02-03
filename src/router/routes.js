import IndexLayout from "@/layout/IndexLayout.vue"; // Normal Layout
import AdminLayout from "@/layout/AdminLayout.vue"; // Admin Layout

export const SuperUser = "super";
export const AdminUser = "admin";
export const DatabaseUser = "dbuser";
export const StudentUser = "student";
export const TeacherUser = "teacher";

const routes = [
  {
    path: "/",
    component: IndexLayout,
    hidden: true,
    children: [
      {
        path: "/",
        name: "Home",
        component: () => import("@/views/home/Home"),
        meta: {
          activeMenu: "/",
          title: "HomeMenu"
        }
      },
      {
        path: "/news-detail/:id",
        name: "NewsDetail",
        component: () => import("@/views/home/NewsDetail"),
        meta: { activeMenu: "/" }
      },
      {
        path: "/news-list/:column",
        name: "NewsList",
        component: () => import("@/views/home/NewsList"),
        meta: { activeMenu: "/" }
      },
      {
        path: "/database/lexical-database",
        name: "LexicalDatabase",
        component: () => import("@/views/database/LexicalDatabase"),
        meta: {
          roles: [DatabaseUser],
          auth: true
        }
      },
      {
        path: "/about/introduction",
        name: "Introduction",
        component: () => import("@/views/about/Introduction"),
        meta: { activeMenu: "/about/introduction" }
      },
      {
        path: "/about/team",
        component: () => import("@/views/about/Team"),
        name: "Team",
        meta: { activeMenu: "/about/team" }
      },
      {
        path: "/about/team-detail/:id",
        name: "TeamDetail",
        component: () => import("@/views/about/TeamDetail"),
        meta: { activeMenu: "/about/team" }
      },
      {
        path: "/about/contact",
        name: "Contact",
        component: () => import("@/views/about/Contact"),
        meta: { activeMenu: "/about/contact" }
      },
      {
        path: "/learning-platform",
        name: "LearningPlatform",
        component: () => import("@/views/learning/LearningPlatform"),
        meta: {
          roles: [StudentUser],
          auth: true,
          activeMenu: "/learning-platform"
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
          title: "ProfileCenter",
          icon: "profile",
          auth: true
        }
      }
    ]
  },
  {
    path: "/users-management",
    component: AdminLayout,
    children: [
      {
        path: "/users-management",
        name: "Users",
        component: () => import("@/views/dashboard/Users"),
        meta: {
          title: "UserManagement",
          icon: "user",
          auth: true,
          roles: [SuperUser]
        }
      }
    ]
  },

  {
    path: "/website-management",
    component: AdminLayout,
    redirect: "/website-management/introduction",
    name: "Website",
    meta: {
      title: "WebsiteManagement",
      icon: "website",
      roles: [SuperUser, AdminUser]
    },
    children: [
      {
        path: "introduction",
        name: "IntroductionManagement",
        component: () => import("@/views/dashboard/website/Introduction"),
        meta: {
          title: "IntroductionManagement",
          icon: "introduction",
          auth: true
        }
      },
      {
        path: "carousels",
        name: "Carousels",
        component: () => import("@/views/dashboard/website/Carousels"),
        meta: {
          title: "CarouselManagement",
          icon: "pictures",
          auth: true
        }
      },
      {
        path: "news",
        name: "News",
        component: () => import("@/views/dashboard/website/News"),
        meta: {
          title: "NewsManagement",
          icon: "news",
          auth: true
        }
      },
      {
        path: "members",
        name: "Members",
        component: () => import("@/views/dashboard/website/Members"),
        meta: {
          title: "MemberManagement",
          icon: "member",
          auth: true
        }
      },
      {
        path: "login-history",
        name: "LoginHistory",
        component: () => import("@/views/dashboard/website/LoginHistory"),
        meta: {
          title: "LoginHistory",
          icon: "history",
          auth: true,
          roles: [SuperUser]
        }
      },
      {
        path: "systems",
        name: "Systems",
        component: () => import("@/views/dashboard/website/Systems"),
        meta: {
          title: "SystemManagement",
          icon: "maintain",
          auth: true,
          roles: [SuperUser]
        }
      }
    ]
  },
  {
    path: "/database-management",
    component: AdminLayout,
    redirect: "/database-management/performers",
    name: "LexicalDatabaseSetting",
    meta: {
      title: "DatabaseManagement",
      icon: "contrast",
      roles: [SuperUser, AdminUser]
    },
    children: [
      {
        path: "performers",
        name: "Performers",
        component: () => import("@/views/dashboard/database/Performers"),
        meta: {
          title: "PerformerManagement",
          icon: "peoples",
          auth: true
        }
      },
      {
        path: "handshapes",
        name: "Handshapes",
        component: () => import("@/views/dashboard/database/Handshapes"),
        meta: {
          title: "HandshapeManagement",
          icon: "handshape",
          auth: true
        }
      },
      {
        path: "lexicon",
        name: "Lexicon",
        component: () => import("@/views/dashboard/database/Lexicons"),
        meta: {
          title: "LexiconManagement",
          icon: "words",
          auth: true
        }
      },
      {
        path: "videos",
        name: "LexicalVideos",
        component: () => import("@/views/dashboard/database/LexicalVideos"),
        meta: {
          title: "LexicalVideoManagement",
          icon: "video",
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
