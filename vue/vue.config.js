const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  // 配置项目的基本信息
  transpileDependencies: true,
  publicPath: '/',
  lintOnSave: true,
});
