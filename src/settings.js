module.exports = {
  title: "中国手语及聋人研究中心",
  fixedHeader: false,
  sidebarLogo: true,
  errorLog: "production",
  videojsOptions: {
    options: {
      controls: true,
      language: "zh-CN",
      autoplay: false,
      fluid: true,
      controlBar: {
        volumeBar: false,
        volumePanel: false,
        currentTimeDisplay: false
      },
      languages: {
        "zh-CN": {
          "The media could not be loaded, either because the server or network failed or because the format is not supported.":
            "暂无数据"
        }
      }
    }
  }
};
