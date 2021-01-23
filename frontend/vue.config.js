const fs = require('fs');
const path = require('path');

const srcRoot = './src';
module.exports = {
  publicPath: process.env.NODE_ENV === 'electron' ? './' : '/',
  chainWebpack: (config) => {
    config.plugin('define').tap((definitions) => {
      definitions[0].__DEV__ = process.env.NODE_ENV === 'development';
      definitions[0].__ELECTRON__ = process.env.NODE_ENV === 'electron';
      return definitions;
    });

    // add `src/*` to alias
    const files = fs.readdirSync(path.join(__dirname, srcRoot));
    files.filter((file) => fs.statSync(`${srcRoot}/${file}`).isDirectory())
      .forEach((file) => {
        config.resolve.alias.set(file, path.resolve(__dirname, `${srcRoot}/${file}`));
      });
  },
};
