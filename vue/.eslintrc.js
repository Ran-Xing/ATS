module.exports = {
  root: true,

  env: {
    node: true,
  },

  parserOptions: {
    parser: '@babel/eslint-parser',
  },

  rules: {
    'no-console': 'off',
    'no-debugger': 'off',
    'vue/multi-word-component-names': 'off',
    'vue/html-quotes': 'off',
  },

  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
};
