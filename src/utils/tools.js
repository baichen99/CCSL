import defaultSettings from "@/settings";

const title = defaultSettings.title;
const TOKEN_KEY = "CCSL_JWT_TOKEN";
const USER_KEY = "CCSL_USER_ID";

export function getToken() {
  return localStorage.getItem(TOKEN_KEY);
}

export function setToken(token) {
  return localStorage.setItem(TOKEN_KEY, token);
}

export function removeToken() {
  return localStorage.removeItem(TOKEN_KEY);
}

export function getUser() {
  return localStorage.getItem(USER_KEY);
}

export function setUser(userID) {
  return localStorage.setItem(USER_KEY, userID);
}

export function removeUser() {
  return localStorage.removeItem(USER_KEY);
}

export function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`;
  }
  return `${title}`;
}

export function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role));
  } else {
    return true;
  }
}

export function filterRoutes(routes, roles) {
  const res = [];

  routes.forEach(route => {
    const tmp = { ...route };
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterRoutes(tmp.children, roles);
      }
      res.push(tmp);
    }
  });
  return res;
}
