const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
  // permission_routes: state => state.permission.routes,
  sign: state => state.sign.sign,
  regions: state => state.sign.regions,
  wordTypes: state => state.sign.wordTypes,
  letters: state => state.sign.letters,
  words: state => state.sign.words
};

export default getters;
