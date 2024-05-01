const { SemicolonPreference } = require("typescript");

module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react-hooks/recommended',
  ],
  ignorePatterns: ['dist', '.eslintrc.cjs'],
  parser: '@typescript-eslint/parser',
  plugins: ['react-refresh'],
  rules: {
    'react-refresh/only-export-components': [
      'warn',
      { allowConstantExport: true},
    ],
    // 'no-restricted-syntax': [
    //   'error',
    //   {
    //     selector: 'JSXAttribute[name.name="className"][value.value=/gap/]',
    //     message: 'The use of "gap" is not allowed, instead use "margin", "padding", etc..',
    //   },
    // ],
  },
}
