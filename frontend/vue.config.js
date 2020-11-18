const fs = require('fs');
const path = require('path');

const srcRoot = './src';
module.exports = {
  publicPath: '/',
  chainWebpack: (config) => {
    config.plugin('define').tap((definitions) => {
      definitions[0].__DEV__ = process.env.NODE_ENV === 'development';
      return definitions;
    });

    // add `src/*` to alias
    const files = fs.readdirSync(`${__dirname}/${srcRoot}`);
    files.filter((file) => fs.statSync(`${srcRoot}/${file}`).isDirectory())
      .forEach((file) => {
        config.resolve.alias.set(file, path.resolve(__dirname, `${srcRoot}/${file}`));
      });
  },
};
