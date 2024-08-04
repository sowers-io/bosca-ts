module.exports = {
  env: {
    es2021: true,
    node: true,
  },
  extends: [
    'standard',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 12,
    sourceType: 'module',
  },
  plugins: [
    '@typescript-eslint',
  ],
  rules: {
    'space-before-function-paren': ['error', {
      anonymous: 'always',
      named: 'never',
      asyncArrow: 'always',
    }],
    'comma-dangle': ['error', 'always-multiline'],
    indent: ['error', 2],
    'keyword-spacing': ['error', { before: true, after: true }],
  },
}
