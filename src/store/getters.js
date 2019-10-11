const getters = {
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
  device: state => state.app.device,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
  wordPosTypes: state => state.sign.wordPosTypes,
  constructTypes: state => state.sign.constructTypes,
  newsColumns: state => state.sign.newsColumns,
  userTypes: state => state.sign.userTypes,
  genderTypes: state => state.sign.genderTypes
};

export default getters;
