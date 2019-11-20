const getters = {
  settings: state => state.settings,
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
  device: state => state.app.device,
  locale: state => state.app.locale,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
  wordPosTypes: state => state.sign.wordPosTypes,
  wordInitial: state => state.sign.wordInitial,
  wordFormations: state => state.sign.wordFormations,
  userTypes: state => state.sign.userTypes,
  userState: state => state.sign.userState,
  genderTypes: state => state.sign.genderTypes,
  newsColumns: state => state.sign.newsColumns,
  newsState: state => state.sign.newsState,
  newsTypes: state => state.sign.newsTypes,
  languageTypes: state => state.sign.languageTypes
};

export default getters;
