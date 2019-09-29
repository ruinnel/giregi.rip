module.exports = {
  root: true,
  env: {
    browser: true,
  },
  extends: [
    'standard',
    'plugin:vue/essential',
    'plugin:vue/recommended',
  ],
  plugins: [
    'html',
    'standard',
    'vue',
  ],
  rules: {
    'semi': [2, 'always'],
    'comma-dangle': ['error', {
      'arrays': 'always-multiline',
      'objects': 'always-multiline',
      'imports': 'always-multiline',
      'exports': 'always-multiline',
      'functions': 'ignore',
    }],
    'vue/html-closing-bracket-newline': 0,
    'vue/max-attributes-per-line': 0,
    'vue/singleline-html-element-content-newline': 0,
    'space-before-function-paren': 0,
    'vue/name-property-casing': ['error', 'PascalCase'],
    "vue/component-name-in-template-casing": ["error", "kebab-case", {
      "registeredComponentsOnly": true,
      "ignores": []
    }],
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'vue/html-self-closing': 0,
  },
  parserOptions: {
    parser: 'babel-eslint',
    sourceType: 'module',
  },
  globals: {
    __DEV__: 'readonly',
    __USE_MOCK_API_CLIENT__: 'readonly',
  },
};
