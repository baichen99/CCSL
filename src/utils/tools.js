import defaultSettings from "@/settings";

const title = defaultSettings.title;

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
