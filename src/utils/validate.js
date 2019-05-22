/**
 * @param {string} path
 * @returns {Boolean}
 */
export function isExternal(path) {
  const reg = new RegExp(/^(https?:|mailto:|tel:)/);
  return reg.test(path);
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validUsername(username) {
  const shuReg = new RegExp(/\d{8}/);
  const emailReg = new RegExp(
    /^([A-Za-z0-9_\-\.\u4e00-\u9fa5])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$/
  );
  return shuReg.test(username) || emailReg.test(username);
}
