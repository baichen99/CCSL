const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
  permission_routes: state => state.permission.routes,
  sign: state => state.signlang.sign,
  regions: state => state.signlang.regions,
  wordTypes: state => state.signlang.wordTypes,
  letters: state => state.signlang.letters
};
export default getters;
