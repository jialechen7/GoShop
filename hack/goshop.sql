-- Database creation
DROP DATABASE IF EXISTS `goshop`;
CREATE DATABASE IF NOT EXISTS `goshop` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `goshop`;

-- Table creation for admin_info
CREATE TABLE IF NOT EXISTS `admin_info` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(30) NOT NULL COMMENT '用户名',
    `password` VARCHAR(50) NOT NULL COMMENT '密码',
    `role_ids` VARCHAR(50) NOT NULL COMMENT '角色ids',
    `user_salt` VARCHAR(10) NOT NULL COMMENT '加密盐',
    `is_admin` TINYINT(1) DEFAULT 0 NOT NULL COMMENT '是否超级管理员',
    `created_at` DATETIME NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL,
    UNIQUE (`name`)
) COMMENT='管理员表';

-- Sample data for admin_info
INSERT INTO `admin_info` (`name`, `password`, `role_ids`, `user_salt`, `is_admin`, `created_at`, `updated_at`) VALUES
('jialechen', 'b23239a15dbae17bc78cc83cd3c1d071', '1,2', 'wUUjkJqFiM', 1, NOW(), NOW()),
('admin', 'b23239a15dbae17bc78cc83cd3c1d071', '1,2', 'wUUjkJqFiM', 1, NOW(), NOW());

-- Table creation for position_info
CREATE TABLE IF NOT EXISTS `position_info` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `pic_url` VARCHAR(255) NOT NULL COMMENT '图片链接',
    `goods_name` VARCHAR(100) NOT NULL COMMENT '商品名称',
    `link` VARCHAR(200) NOT NULL COMMENT '跳转链接',
    `sort` TINYINT DEFAULT 0 NOT NULL COMMENT '排序值',
    `goods_id` INT DEFAULT 0 NOT NULL COMMENT '商品id',
    `created_at` DATETIME NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
) COMMENT='手工位表';

-- Sample data for position_info
INSERT INTO `position_info` (`pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`) VALUES
('https://example.com/img1.jpg', 'Product A', 'https://example.com/productA', 1, 101, NOW(), NOW()),
('https://example.com/img2.jpg', 'Product B', 'https://example.com/productB', 2, 102, NOW(), NOW());

-- Table creation for rotation_info
DROP TABLE IF EXISTS `rotation_info`;
CREATE TABLE IF NOT EXISTS `rotation_info` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `pic_url` VARCHAR(200) DEFAULT '' NOT NULL COMMENT '轮播图片',
    `link` VARCHAR(200) DEFAULT '' NOT NULL COMMENT '跳转链接',
    `sort` TINYINT(1) DEFAULT 0 NOT NULL COMMENT '排序字段',
    `created_at` DATETIME NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
) COMMENT='轮播图表';

-- Sample data for rotation_info
INSERT INTO `rotation_info` (`pic_url`, `link`, `sort`, `created_at`, `updated_at`) VALUES
('https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', '8', 1, NOW(), NOW()),
('https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', '6', 2, NOW(), NOW());

-- Table creation for order_info
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE IF NOT EXISTS `order_info` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `number` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
    `user_id` INT NOT NULL DEFAULT '0' COMMENT '用户id',
    `pay_type` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '支付方式 1微信 2支付宝 3云闪付',
    `remark` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
    `pay_at` DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '支付时间',
    `status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价',
    `consignee_name` VARCHAR(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `consignee_phone` VARCHAR(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人手机号',
    `consignee_address` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人详细地址',
    `price` INT NOT NULL DEFAULT '0' COMMENT '订单金额 单位分',
    `coupon_price` INT NOT NULL DEFAULT '0' COMMENT '优惠券金额 单位分',
    `actual_price` INT NOT NULL DEFAULT '0' COMMENT '实际支付金额 单位分',
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单信息表';

-- Sample data for order_info
BEGIN;
INSERT INTO `order_info` (`number`, `user_id`, `pay_type`, `remark`, `pay_at`, `status`, `consignee_name`, `consignee_phone`, `consignee_address`, `price`, `coupon_price`, `actual_price`, `created_at`, `updated_at`) VALUES
('1659231316407832000111', 1, 1, '备注1', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 100, 9900, '2022-08-27 09:35:16', '2022-08-27 09:35:16'),
('1659231554317361000757', 1, 1, '备注2', NULL, 2, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 200, 9800, '2022-08-27 09:39:14', '2022-08-27 09:39:14'),
('1661603467832912000516', 1, 2, '', '2022-12-13 21:52:26', 3, '', '', '', 0, 0, 0, '2022-12-08 20:31:07', '2022-12-08 20:31:07'),
('1661603562656619000513', 1, 3, '放到快递柜', '2022-12-13 21:52:19', 4, '王先生', '13269477432', '北京丰台汽车博物馆', 0, 0, 0, '2022-12-09 20:32:42', '2022-12-09 20:32:42');
COMMIT;


CREATE TABLE IF NOT EXISTS role_info (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '角色ID',
    name VARCHAR(50) NOT NULL COMMENT '角色名称',
    `desc` VARCHAR(255) COMMENT '角色描述',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME NULL COMMENT '删除时间',
    constraint role_info_pk
        unique (name)
) COMMENT='角色信息表';

BEGIN;
INSERT INTO `role_info` (`name`, `desc`, `created_at`, `updated_at`)
VALUES ('运营', '运营权限', '2022-09-25 10:35:52', '2022-09-25 10:35:52');
COMMIT;

CREATE TABLE IF NOT EXISTS permission_info (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '权限ID',
    name VARCHAR(50) NOT NULL COMMENT '权限名称',
    path VARCHAR(255) COMMENT '权限路径，指向具体的API或页面',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME NULL COMMENT '删除时间',
    constraint permission_info_pk
        unique (name)
) COMMENT='权限信息表';

BEGIN;
INSERT INTO permission_info (name, path, created_at, updated_at) VALUES
('首页配置', '/homepageManager', '2022-09-25 15:03:01', '2022-09-25 15:03:43'),
('轮播图', '/homepageManager/bannerSwiper', NOW(), NOW());
COMMIT;

CREATE TABLE IF NOT EXISTS role_permission_info (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    role_id INT NOT NULL COMMENT '角色ID',
    permission_id INT NOT NULL COMMENT '权限ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME NULL COMMENT '删除时间',
    constraint role_permission_info_permission_info_id_fk
        foreign key (permission_id) references permission_info (id)
            on update cascade on delete cascade,
    constraint role_permission_info_role_info_id_fk
        foreign key (role_id) references role_info (id)
            on update cascade on delete cascade
) COMMENT = '角色权限关联表';


CREATE TABLE file_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) NOT NULL COMMENT '图片名称',
    src VARCHAR(128) NOT NULL COMMENT '本地文件存储路径',
    url VARCHAR(255) NOT NULL COMMENT '文件地址',
    user_id INT NOT NULL COMMENT '用户id',
    created_at DATETIME NOT NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL
) COMMENT='文件信息表';

DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `name` VARCHAR(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar` VARCHAR(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像URL',
    `password` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密后的密码',
    `user_salt` VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐，用于生成密码',
    `sex` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '性别：1表示男，2表示女',
    `status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '状态：1表示正常，2表示拉黑冻结',
    `sign` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '个性签名',
    `secret_answer` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密保问题答案',
    `created_at` DATETIME DEFAULT NULL COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT NULL COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    constraint user_info_pk
        unique (name)
) ENGINE=InnoDB AUTO_INCREMENT DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- 插入 user_info 表的记录
-- ----------------------------
BEGIN;
INSERT INTO `user_info`
(`id`, `name`, `avatar`, `password`, `user_salt`, `sex`, `status`, `sign`, `secret_answer`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1, 'jialechen', 'https://cdn.acwing.com/media/user/profile/photo/40815_lg_d5f4bcc813.JPG', '905d0c656e6f93d5e73fb00ad0702a41', 'JAUrpNobzs', 0, 1, '个性签名', '银河中学', '2022-07-28 17:19:42', '2022-07-31 19:25:01', NULL);
COMMIT;


-- ----------------------------
-- Table structure for category_info
-- ----------------------------
DROP TABLE IF EXISTS `category_info`;
CREATE TABLE `category_info` (
    `id` int NOT NULL AUTO_INCREMENT,
    `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级id',
    `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'icon',
    `deleted_at` datetime DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `level` tinyint(1) NOT NULL DEFAULT '1' COMMENT '等级 默认1级分类',
    `sort` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='分类表\n';

-- ----------------------------
-- Records of category_info
-- ----------------------------
-- 插入一级分类: 电子产品
INSERT INTO category_info (name, parent_id, level, sort, created_at, updated_at)
VALUES
    ('电子产品', 0, 1, 1, NOW(), NOW());

-- 插入二级分类: 手机, 电脑, 配件
INSERT INTO category_info (name, parent_id, level, sort, created_at, updated_at)
VALUES
    ('手机', 1, 2, 1, NOW(), NOW()),
    ('电脑', 1, 2, 2, NOW(), NOW()),
    ('配件', 1, 2, 3, NOW(), NOW());

-- 插入三级分类: 智能手机, 手机配件, 笔记本电脑, 电竞电脑, 耳机
INSERT INTO category_info (name, parent_id, level, sort, created_at, updated_at)
VALUES
    ('智能手机', 2, 3, 1, NOW(), NOW()),
    ('手机配件', 2, 3, 2, NOW(), NOW()),
    ('笔记本电脑', 3, 3, 1, NOW(), NOW()),
    ('电竞电脑', 3, 3, 2, NOW(), NOW()),
    ('耳机', 4, 3, 1, NOW(), NOW());

-- 设置 AUTO_INCREMENT 从 10 开始
ALTER TABLE category_info AUTO_INCREMENT = 10;

-- 插入一级分类: 家居用品 (id 10)
INSERT INTO category_info (name, parent_id, level, sort, created_at, updated_at)
VALUES
    ('家居用品', 0, 1, 1, NOW(), NOW());

-- 插入二级分类: 家具, 家居饰品, 厨房用品 (id 11, 12, 13)
INSERT INTO category_info (name, parent_id, level, sort, created_at, updated_at)
VALUES
    ('家具', 10, 2, 1, NOW(), NOW()),
    ('家居饰品', 10, 2, 2, NOW(), NOW()),
    ('厨房用品', 10, 2, 3, NOW(), NOW());

-- 插入三级分类: 沙发, 餐桌, 窗帘, 地毯, 厨房小电器 (id 14, 15, 16, 17, 18)
INSERT INTO category_info (name, parent_id, level, sort, created_at, updated_at)
VALUES
    ('沙发', 11, 3, 1, NOW(), NOW()),
    ('餐桌', 11, 3, 2, NOW(), NOW()),
    ('窗帘', 12, 3, 1, NOW(), NOW()),
    ('地毯', 12, 3, 2, NOW(), NOW()),
    ('厨房小电器', 13, 3, 1, NOW(), NOW());


-- ----------------------------
-- Table structure for article_info
-- ----------------------------
DROP TABLE IF EXISTS `article_info`;
CREATE TABLE `article_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL DEFAULT '0' COMMENT '作者id',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '摘要',
  `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面图',
  `is_admin` tinyint(1) NOT NULL DEFAULT '2' COMMENT '1后台管理员发布 2前台用户发布',
  `praise` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '文章详情',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章（种草）表';

-- ----------------------------
-- Records of article_info
-- ----------------------------
INSERT INTO `article_info` (`user_id`, `title`, `desc`, `pic_url`, `is_admin`, `praise`, `detail`, `created_at`, `updated_at`)
VALUES (1, '探索MacBook Pro的性能', '这篇文章探讨了MacBook Pro的性能表现，适用于不同场景的使用分析', 'https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', 2, 10,
        'MacBook Pro 是苹果公司推出的高端笔记本电脑，以其卓越的性能和精致的设计深受专业人士和创作者的喜爱。无论是在处理大型图形设计任务，还是在进行视频编辑和开发工作，MacBook Pro 都能够提供无与伦比的性能支持。特别是其配备的 M1 或 M2 芯片，结合高分辨率的 Retina 屏幕，使得 MacBook Pro 成为当前市场上最受欢迎的高性能笔记本之一。',
        NOW(), NOW());
INSERT INTO `article_info` (`user_id`, `title`, `desc`, `pic_url`, `is_admin`, `praise`, `detail`, `created_at`, `updated_at`)
VALUES (1, 'iPhone 14 Pro的全面评测', '本文将对iPhone 14 Pro进行全面评测，探讨其性能、拍照功能及创新设计', 'https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', 2, 20,
        'iPhone 14 Pro 是苹果公司推出的最新一代旗舰智能手机，搭载了强大的 A16 Bionic 芯片，拥有极致的性能表现。其全新的动态岛设计不仅使得显示屏更加简洁美观，还增强了互动性。iPhone 14 Pro 配备了 48MP 主摄像头，支持更高质量的照片和视频拍摄，特别适合喜欢摄影的用户。此外，iPhone 14 Pro 的显示效果也极其出色，采用了 ProMotion 技术和 Always-On 屏幕，提供了更流畅的体验。',
        NOW(), NOW());


-- ----------------------------
-- Table structure for praise_info
-- ----------------------------
DROP TABLE IF EXISTS `praise_info`;
CREATE TABLE `praise_info` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '点赞表',
    `user_id` int NOT NULL,
    `type` tinyint(1) NOT NULL COMMENT '点赞类型 1商品 2文章',
    `object_id` int NOT NULL DEFAULT '0' COMMENT '点赞对象id 方便后期扩展',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_index` (`user_id`,`type`,`object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- -------------------------
-- Records of praise_info
-- ----------------------------
BEGIN;
INSERT INTO `praise_info` VALUES (1, 1 , 2, 2, '2023-01-19 12:18:07', '2023-01-19 12:18:07');
COMMIT;

-- ----------------------------
-- Table structure for comment_info
-- ----------------------------
DROP TABLE IF EXISTS `comment_info`;
CREATE TABLE `comment_info` (
    `id` int NOT NULL AUTO_INCREMENT,
    `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级评论id',
    `user_id` int NOT NULL DEFAULT '0',
    `object_id` int NOT NULL DEFAULT '0',
    `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论类型：1商品 2文章',
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论内容',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `unique_index` (`user_id`,`object_id`,`type`,`content`,`parent_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of comment_info
-- ----------------------------
BEGIN;
INSERT INTO `comment_info` VALUES (1, 0, 1, 2, 2, '好评 下次还会买', '2022-07-31 17:23:48', '2022-07-31 17:23:48', NULL);
INSERT INTO `comment_info` VALUES (2, 0, 1, 2, 2, '还行', '2022-07-31 17:24:10', '2022-07-31 17:24:10', NULL);
INSERT INTO `comment_info` VALUES (3, 2, 1, 2, 2, '真的还行吗?', '2022-07-31 17:24:59', '2022-07-31 17:24:59', NULL);
COMMIT;

-- ----------------------------
-- Table structure for consignee_info
-- ----------------------------
DROP TABLE IF EXISTS `consignee_info`;
CREATE TABLE `consignee_info` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '收货地址表',
    `user_id` int NOT NULL DEFAULT '0',
    `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认地址1  非默认0\n',
    `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `province` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `city` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `town` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '县区',
    `street` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '街道乡镇',
    `detail` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址详情',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of consignee_info
-- ----------------------------
BEGIN;
INSERT INTO `consignee_info` VALUES (1, 1, 1, '炸鸡', '13269477632', '北京', '北京市', '房山区', '拱辰街道', '大学城西', '2022-07-31 14:42:33', '2022-07-31 14:44:50', NULL);
COMMIT;

-- ----------------------------
-- Table structure for goods_info
-- ----------------------------
DROP TABLE IF EXISTS `goods_info`;
CREATE TABLE `goods_info` (
    `id` int NOT NULL AUTO_INCREMENT,
    `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
    `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
    `price` int NOT NULL DEFAULT '1' COMMENT '价格 单位分',
    `level1_category_id` int NOT NULL COMMENT '1级分类id',
    `level2_category_id` int NOT NULL DEFAULT '0' COMMENT '2级分类id',
    `level3_category_id` int NOT NULL DEFAULT '0' COMMENT '3级分类id',
    `brand` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌',
    `stock` int NOT NULL DEFAULT '0' COMMENT '库存',
    `sale` int NOT NULL DEFAULT '0' COMMENT '销量',
    `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
    `detail_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '商品详情',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品表';

-- ----------------------------
-- Records of goods_info
-- ----------------------------
-- 插入商品数据，分别对应每个分类
-- 家居用品 > 家具
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, created_at, updated_at)
VALUES
    ('沙发', 'https://img10.360buyimg.com/n1/jfs/t1/159795/9/49579/202482/672af451Ff0a61b34/f063016c41fa50de.jpg.avif', 1999, 10, 11, 0, NOW(), NOW()),
    ('餐桌', 'https://img12.360buyimg.com/n1/jfs/t1/160186/19/47269/98348/66c7e925F87921646/c7209966aa477011.jpg.avif', 899, 10, 11, 0, NOW(), NOW());

-- 家居用品 > 家居饰品
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, created_at, updated_at)
VALUES
    ('窗帘', 'https://img13.360buyimg.com/n7/jfs/t1/193417/1/50457/86358/67226adfF89704bc5/5844e97ccc356e60.jpg.avif', 199, 10, 12, 0, NOW(), NOW()),
    ('地毯', 'https://img13.360buyimg.com/n7/jfs/t1/225775/21/18743/102631/667034bdF12a1a520/3a5aee679a79711c.jpg.avif', 299, 10, 12, 0, NOW(), NOW());

-- 家居用品 > 厨房用品
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, created_at, updated_at)
VALUES
    ('厨房小电器', 'https://img12.360buyimg.com/n7/jfs/t1/238669/9/18316/146769/673420c4Fb68cb738/08e7812d64c1d3ed.png.avif', 599, 10, 13, 0, NOW(), NOW());

-- 插入商品数据
-- 一级分类: 电子产品 > 二级分类: 手机 > 三级分类: 智能手机
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, brand, stock, sale, tags, created_at, updated_at)
VALUES
    ('iPhone 14 Pro', 'https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', 899900, 1, 2, 5, 'Apple', 100, 30, '旗舰,智能手机', NOW(), NOW());

-- 一级分类: 电子产品 > 二级分类: 手机 > 三级分类: 手机配件
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, brand, stock, sale, tags, created_at, updated_at)
VALUES
    ('原装苹果充电器', 'https://img11.360buyimg.com/n7/jfs/t1/120933/14/49345/20820/67331648F08cc5230/47180b9222541553.png.avif', 9900, 1, 2, 6, 'Apple', 200, 50, '充电器,配件', NOW(), NOW());

-- 一级分类: 电子产品 > 二级分类: 电脑 > 三级分类: 笔记本电脑
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, brand, stock, sale, tags, created_at, updated_at)
VALUES
    ('MacBook Pro 16', 'https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', 1799900, 1, 3, 7, 'Apple', 50, 20, '高端,笔记本电脑', NOW(), NOW());

-- 一级分类: 电子产品 > 二级分类: 电脑 > 三级分类: 电竞电脑
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, brand, stock, sale, tags, created_at, updated_at)
VALUES
    ('ROG 电竞笔记本', 'https://img10.360buyimg.com/n7/jfs/t1/137023/16/49666/119246/67336731F79e8f60a/ca255f329ceeec1a.jpg.avif', 1599900, 1, 3, 8, 'ROG', 30, 10, '电竞,笔记本电脑', NOW(), NOW());

INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, brand, stock, sale, tags, created_at, updated_at)
VALUES
    ('暗影精灵9', 'https://img12.360buyimg.com/n7/jfs/t1/217126/21/47609/138564/6731d8aaFb02f036d/f62ee8ad27b9f38c.jpg.avif', 1599900, 1, 3, 8, 'ROG', 30, 10, '电竞,笔记本电脑', NOW(), NOW());


-- 一级分类: 电子产品 > 二级分类: 配件 > 三级分类: 耳机
INSERT INTO goods_info (name, pic_url, price, level1_category_id, level2_category_id, level3_category_id, brand, stock, sale, tags, created_at, updated_at)
VALUES
    ('Sony WH-1000XM5', 'https://img14.360buyimg.com/n7/jfs/t1/193676/21/52428/127689/67347f8aF02cff604/fe6fad80b3d9ce03.jpg.avif', 299900, 1, 4, 9, 'Sony', 150, 40, '耳机,降噪', NOW(), NOW());

