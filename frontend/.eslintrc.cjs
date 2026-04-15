module.exports = {
  extends: [
    'plugin:@typescript-eslint/recommended',
    'plugin:react/recommended',
    'prettier'
  ],
  parser: '@typescript-eslint/parser', // 使用TypeScript解析器
  // rules: {
  //   // 自定义规则（会覆盖extends的配置）
  //   'prettier/prettier': 'error',  // 将Prettier规则作为ESLint错误
  //   'no-console': 'warn',          // 示例：自定义ESLint规则
  // },
  // parserOptions: {
  //   ecmaVersion: 'latest',         // 支持最新ES语法
  //   sourceType: 'module'           // 使用ES模块
  // },
  // env: {
  //   browser: true,                 // 根据项目环境调整
  //   node: true
  // }
};
