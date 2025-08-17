# GoShop - 基于GoFrame的电商系统

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![GoFrame](https://img.shields.io/badge/GoFrame-v2.7.4-blue)](https://goframe.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

一个基于 [GoFrame](https://goframe.org/) 框架开发的现代化电商系统，支持前后端分离架构，具备完整的商品管理、订单处理、用户管理、优惠券系统等功能。

## ✨ 主要特性

### 🏗️ 技术架构
- **框架**: GoFrame v2.7.4 - 企业级Go应用开发框架
- **数据库**: MySQL 8.0+ - 主数据存储
- **缓存**: Redis 6.0+ - 缓存、分布式锁、消息队列
- **认证**: gtoken - JWT令牌认证，支持管理端和用户端双重认证
- **API文档**: 内置OpenAPI/Swagger支持
- **部署**: Docker + Kubernetes支持

### 🛍️ 业务功能

#### 商城核心功能
- **商品管理**: 商品分类、商品信息、商品规格、库存管理
- **订单系统**: 订单创建、支付处理、订单状态跟踪
- **购物车**: 商品添加、数量调整、批量操作
- **用户系统**: 用户注册登录、个人信息管理、收货地址
- **优惠券系统**: 普通优惠券、秒杀优惠券、优惠券分发

#### 内容管理
- **文章系统**: 种草推荐、用户发布、内容审核
- **评论系统**: 商品评论、文章评论、互动功能
- **点赞收藏**: 商品收藏、文章收藏、点赞功能
- **轮播图**: 首页轮播、活动推广

#### 秒杀系统
- **高并发处理**: 基于Redis+Lua脚本的原子操作
- **库存控制**: 实时库存扣减、超卖防护
- **限购机制**: 用户重复购买检查
- **异步处理**: Redis Stream消息队列处理优惠券分发

### 👨‍💼 管理后台
- **管理员系统**: 角色权限、管理员管理
- **数据大屏**: 业务数据统计和展示
- **权限控制**: 细粒度权限管理
- **文件管理**: 本地存储和七牛云存储支持
- **OAuth登录**: GitHub OAuth第三方登录

## 🚀 快速开始

### 环境要求

- Go 1.22+
- MySQL 8.0+
- Redis 6.0+
- 可选: Docker & Kubernetes

### 安装步骤

1. **克隆项目**
```bash
git clone https://github.com/jialechen7/goshop.git
cd goshop
```

2. **安装依赖**
```bash
go mod tidy
```

3. **安装GoFrame CLI工具**
```bash
# 如果你没有安装gf命令行工具
go install github.com/gogf/gf/cmd/gf/v2@latest
```

4. **配置数据库**
```bash
# 创建数据库并导入表结构
mysql -u root -p < hack/goshop.sql

# 导入初始数据
mysql -u root -p goshop < hack/init_data.sql
```

5. **配置Redis**
```bash
# 创建用户优惠券消息队列
redis-cli
XGROUP CREATE stream.user_coupon g1 0 MKSTREAM
```

6. **修改配置文件**
```bash
# 复制配置文件
cp manifest/config/config.yaml hack/config.yaml

# 编辑配置文件，修改数据库和Redis连接信息
vim hack/config.yaml
```

7. **启动应用**
```bash
# 开发模式运行
gf run

# 或使用make命令
make build && ./bin/goshop
```

8. **访问应用**
- 前端API（用户端）: http://localhost:8000/frontend
- 后端API（管理员端）: http://localhost:8000/backend  
- API文档: http://localhost:8000/swagger

## 📖 项目结构

```
goshop/
├── api/                    # API定义层
│   ├── backend/           # 后台管理API
│   └── frontend/          # 前台用户API
├── internal/              # 内部代码
│   ├── cmd/              # 启动命令
│   ├── controller/       # 控制器层
│   ├── logic/            # 业务逻辑层
│   ├── service/          # 服务层
│   ├── dao/              # 数据访问层
│   └── model/            # 数据模型
├── hack/                  # 开发工具和配置
│   ├── config.yaml       # 应用配置
│   ├── goshop.sql        # 数据库表结构
│   ├── init_data.sql     # 初始数据
│   └── seckill.lua       # 秒杀Lua脚本
├── manifest/              # 部署清单
│   ├── config/           # 配置文件
│   ├── docker/           # Docker配置
│   └── deploy/           # Kubernetes配置
└── utility/              # 工具包
    ├── captcha/          # 验证码
    ├── redis_lock/       # Redis分布式锁
    └── response/         # 响应处理
```

## 🔧 开发指南

### 常用命令

```bash
# 构建应用
make build

# 代码生成
make dao        # 生成DAO/DO/Entity
make ctrl       # 生成控制器
make service    # 生成服务接口

# 更新框架
make up

# Docker镜像
make image      # 构建镜像
make image.push # 构建并推送镜像

# 部署
make deploy     # 部署到Kubernetes
```

### API接口

#### 后台管理接口 (`/backend`)

**需要认证的接口**:
- `GET /backend/admin` - 管理员管理
- `GET /backend/dashboard` - 数据大屏
- `GET /backend/goods/list` - 商品列表
- `POST /backend/goods/add` - 添加商品
- `PUT /backend/goods/update` - 更新商品
- `DELETE /backend/goods/delete` - 删除商品
- `GET /backend/order/list` - 订单列表
- `GET /backend/coupon/list` - 优惠券列表
- `POST /backend/seckill_coupon/add` - 添加秒杀优惠券

**无需认证的接口**:
- `GET /backend/captcha/get` - 获取验证码
- `POST /backend/oauth/github/login` - GitHub登录
- `GET /backend/oauth/github/receive_code` - GitHub回调

#### 前台用户接口 (`/frontend`)

**需要认证的接口**:
- `POST /frontend/user/create` - 用户注册
- `GET /frontend/user/info` - 用户信息
- `GET /frontend/order/list` - 订单列表
- `POST /frontend/order/add` - 创建订单
- `GET /frontend/cart/list` - 购物车列表
- `POST /frontend/cart/add` - 添加购物车
- `POST /frontend/seckill_coupon/kill` - 参与秒杀

**无需认证的接口**:
- `GET /frontend/goods/list` - 商品列表
- `GET /frontend/goods/detail` - 商品详情
- `GET /frontend/category/list` - 分类列表
- `GET /frontend/rotation/list` - 轮播图列表

### 数据库设计

主要数据表:
- `admin_info` - 管理员信息
- `user_info` - 用户信息  
- `goods_info` - 商品信息
- `category_info` - 商品分类
- `order_info` - 订单信息
- `cart_info` - 购物车
- `coupon_info` - 优惠券信息
- `user_coupon_info` - 用户优惠券
- `rotation_info` - 轮播图
- `article_info` - 文章信息

### 配置说明

主要配置项 (`manifest/config/config.yaml`):

```yaml
server:
  address: ":8000"                    # 服务端口
  
database:
  default:
    link: "mysql:root:root@tcp(127.0.0.1:3306)/goshop"  # 数据库连接
    
redis:
  default:
    address: 127.0.0.1:6379          # Redis地址
    db: 1                            # Redis数据库
    pass: "123456"                   # Redis密码

qiniu:                               # 七牛云存储配置
  bucket: "your-bucket"
  accessKey: "your-access-key"
  secretKey: "your-secret-key"
```

## 🎯 核心功能实现

### 秒杀系统

采用Redis+Lua脚本实现高并发秒杀:

1. **原子性操作**: 使用Lua脚本确保库存检查和扣减的原子性
2. **防重复抢购**: Redis Set记录已参与用户
3. **异步处理**: Redis Stream处理优惠券分发
4. **分布式锁**: 使用redsync防止并发问题

### 认证授权

- **双端认证**: 管理端和用户端独立的JWT认证
- **权限控制**: 基于角色的权限管理系统  
- **中间件**: 统一的认证和权限检查中间件

### 消息队列

使用Redis Stream实现异步消息处理:
- 用户优惠券分发队列
- 支持消费者组和ACK机制
- 消息持久化和重试机制

## 🐳 Docker部署

### 构建镜像

```bash
# 构建应用
make build

# 构建Docker镜像
make image

# 运行容器
docker run -d -p 8000:8000 goshop:latest
```

### Docker Compose

```yaml
version: '3.8'
services:
  goshop:
    build: .
    ports:
      - "8000:8000"
    depends_on:
      - mysql
      - redis
    environment:
      - GF_DATABASE_LINK=mysql:root:password@tcp(mysql:3306)/goshop
      - GF_REDIS_ADDRESS=redis:6379
      
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: goshop
    volumes:
      - ./hack/goshop.sql:/docker-entrypoint-initdb.d/schema.sql
      
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
```

## ☁️ Kubernetes部署

项目包含完整的Kubernetes部署配置:

```bash
# 使用Kustomize部署
make deploy

# 或手动部署
kubectl apply -f manifest/deploy/kustomize/base/
```

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

- [GoFrame](https://goframe.org/) - 优秀的Go企业级开发框架
- [Redis](https://redis.io/) - 高性能内存数据库
- [MySQL](https://www.mysql.com/) - 可靠的关系型数据库

---

**注意**: 此项目仅供学习交流使用，请勿用于商业用途。