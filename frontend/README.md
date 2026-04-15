# 极睿科技-引流宝 (React 重构版)

这是一个使用 React + Vite 重构的企业品牌官网项目。

## 项目结构

```
mint-react/
├── public/                 # 静态资源
│   ├── logo.svg           # 网站 Logo
│   ├── logo_*.png         # 客户 Logo
│   └── qrcode.png         # 公众号二维码
├── src/
│   ├── components/        # 公共组件
│   │   ├── Header/        # 导航栏组件
│   │   ├── Footer/        # 页脚组件
│   │   ├── HeroCarousel/  # 首页轮播组件
│   │   ├── Products/      # 产品展示组件
│   │   ├── Cases/         # 客户案例组件
│   │   ├── ContactForm/   # 联系表单组件
│   │   └── LoginModal/    # 登录弹窗组件
│   ├── pages/             # 页面组件
│   │   ├── Home/          # 首页
│   │   ├── Product/       # 产品页
│   │   ├── Solution/      # 解决方案页
│   │   ├── Price/         # 价格页
│   │   └── LoginPage/     # 登录页
│   ├── App.jsx            # 主应用组件
│   ├── main.jsx           # 应用入口
│   └── index.css          # 全局样式
├── index.html             # HTML 模板
├── package.json           # 项目依赖
└── vite.config.js         # Vite 配置
```

## 功能特性

- **响应式设计**: 支持 PC 和移动端适配
- **路由管理**: 使用 React Router 实现页面导航
- **登录系统**: 
  - 弹窗式登录（首页等页面）
  - 独立登录页面 (/login)
  - 登录状态持久化（localStorage）
  - 支持微信登录（预留接口）
- **页面模块**:
  - 首页: 轮播图、产品服务、客户案例、联系表单
  - 产品页: 引流宝、内容宝、直播宝介绍
  - 解决方案页: 服装、美妆、食品行业方案
  - 价格页: 三种价格方案展示

## 登录功能说明

项目已预留完整的登录功能接口：

1. **Header 组件**: 显示登录按钮或用户信息
2. **LoginModal 组件**: 弹窗式登录/注册
3. **LoginPage 页面**: 独立登录页面
4. **App 组件**: 管理登录状态和用户信息

### 后续接入真实登录 API 的步骤：

1. 在 `LoginModal.jsx` 和 `LoginPage.jsx` 中的 `handleSubmit` 函数中替换模拟请求为真实 API 调用
2. 在 `App.jsx` 中添加 token 管理和自动登录逻辑
3. 根据后端接口调整表单字段和验证规则

## 开发环境运行

```bash
# 安装依赖
pnpm install

# 启动开发服务器
pnpm dev

# 构建生产版本
pnpm build

# 预览生产版本
pnpm preview
```

## 技术栈

- React 18
- React Router 6
- Vite 5
- CSS3 (原生样式，无 UI 框架依赖)

## 浏览器支持

- Chrome (最新版)
- Firefox (最新版)
- Safari (最新版)
- Edge (最新版)

## 注意事项

1. 项目使用原生 CSS 进行样式管理，未引入 UI 框架，保持轻量
2. 图片资源需要放在 `public` 目录下才能直接通过 `/文件名` 引用
3. 登录功能目前为模拟实现，需要接入后端 API
