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
  wordPosTypes: state => state.data.wordPosTypes,
  wordInitial: state => state.data.wordInitial,
  wordFormations: state => state.data.wordFormations,
  userTypes: state => state.data.userTypes,
  userState: state => state.data.userState,
  genderTypes: state => state.data.genderTypes,
  newsColumns: state => state.data.newsColumns,
  newsState: state => state.data.newsState,
  newsTypes: state => state.data.newsTypes,
  languageTypes: state => state.data.languageTypes,
  memberTypes: state => state.data.memberTypes,
  memberDegrees: state => state.data.memberDegrees,
  signs: state => state.data.signs,
  lexicons: state => state.data.lexicons,
  performers: state => state.data.performers
};

export default getters;
