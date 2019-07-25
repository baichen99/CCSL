import router from "@/router";
import store from "@/store";
import { Message } from "element-ui";
import NProgress from "nprogress"; // progress bar
import "nprogress/nprogress.css"; // progress bar style
// import { getToken, getPageTitle, hasPermission } from "@/utils/tools";
import { getToken, getPageTitle } from "@/utils/tools";

NProgress.configure({ showSpinner: false });

const LOGIN_PAGE = "Login";
const HOME_PAGE = "Home";

// const turnTo = (to, access, next) => {
//   if (canTurnTo(to.name, access, routes)) next();
//   else next({ replace: true, name: "404Error" });
// };

router.beforeEach(async (to, from, next) => {
  NProgress.start();
  document.title = getPageTitle(to.meta.title);
  const token = getToken() || store.getters.token;
  if (to.meta.auth) {
    // 需要登录才能访问
    if (!token && to.name !== LOGIN_PAGE) {
      // 未登录且要跳转的页面不是登录页
      next({
        name: LOGIN_PAGE,
        query: { redirect: to.fullPath }
      });
    } else if (token && to.name === LOGIN_PAGE) {
      // 已登录且要跳转的页面是登录页
      next({
        name: HOME_PAGE
      });
    } else {
      // 已登录，再检查用户权限
      const hasRoles = store.getters.roles && store.getters.roles.length > 0;
      const accessibleRoles = to.meta.roles && to.meta.roles.length > 0;
      if (hasRoles) {
        // 如果指定了用户权限，那么检查用户权限
        if (accessibleRoles) {
          const userRole = store.getters.roles[0];
          const allowRole = to.meta.roles;
          if (allowRole.includes(userRole)) {
            next();
          } else {
            next({ replace: true, name: "401Error" });
          }
        } else {
          // 否则所有登录用户都可以访问
          next();
        }
      } else {
        try {
          const data = await store.dispatch("user/getUserInfo");
          store.commit("user/SET_ROLES", data.userType);
          next({ ...to, replace: true });
        } catch (error) {
          await store.dispatch("user/logout");
          Message.error(error || "Error");
          next(`/login?redirect=${to.path}`);
          NProgress.done();
        }
      }
    }
  } else {
    // 不需要授权
    next();
  }
});

router.afterEach(() => {
  NProgress.done();
});
