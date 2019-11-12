const getters = {
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
  device: state => state.app.device,
  locale: state => state.app.locale,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
  wordPosTypes: state => state.sign.wordPosTypes,
  constructTypes: state => state.sign.constructTypes,
  userTypes: state => state.sign.userTypes,
  userState: state => state.sign.userState,
  genderTypes: state => state.sign.genderTypes,
  newsColumns: state => state.sign.newsColumns
};

export default getters;
