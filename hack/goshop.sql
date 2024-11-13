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
('https://img10.360buyimg.com/n1/s450x450_jfs/t1/207741/15/47068/156692/67306ca8Fd87ffeeb/ad1b66065f2ed11e.jpg', 'a', 1, NOW(), NOW()),
('https://img13.360buyimg.com/n1/jfs/t1/240604/12/19894/141092/673161beFff0db991/3d9083445563844e.jpg.avif', 'b', 2, NOW(), NOW());

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- 插入 user_info 表的记录
-- ----------------------------
BEGIN;
INSERT INTO `user_info`
(`id`, `name`, `avatar`, `password`, `user_salt`, `sex`, `status`, `sign`, `secret_answer`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1, 'jialechen', 'https://cdn.acwing.com/media/user/profile/photo/40815_lg_d5f4bcc813.JPG', '905d0c656e6f93d5e73fb00ad0702a41', 'JAUrpNobzs', 1, 1, '个性签名', '银河中学', '2022-07-28 17:19:42', '2022-07-31 19:25:01', NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='轮播图表\n';

-- ----------------------------
-- Records of category_info
-- ----------------------------
BEGIN;
INSERT INTO `category_info` VALUES (1, 0, '家用电器', 'https://img10.360buyimg.com/n1/jfs/t1/211798/40/45226/70698/670e113cFbfebf347/7cb86ec50d195420.jpg.avif', NULL, NULL, NULL, 1, 1);
INSERT INTO `category_info` VALUES (2, 1, '电视', 'https://img11.360buyimg.com/n1/jfs/t1/107975/22/57903/111261/672ca34bFd0264716/832fe66f08b3bb76.jpg.avif', NULL, NULL, NULL, 2, 1);
INSERT INTO `category_info` VALUES (3, 2, '全面屏电视', 'https://img11.360buyimg.com/n1/jfs/t1/107975/22/57903/111261/672ca34bFd0264716/832fe66f08b3bb76.jpg.avif', NULL, NULL, NULL, 3, 1);
INSERT INTO `category_info` VALUES (4, 2, '教育电视', 'https://img11.360buyimg.com/n1/jfs/t1/107975/22/57903/111261/672ca34bFd0264716/832fe66f08b3bb76.jpg.avif', NULL, NULL, NULL, 3, 1);
INSERT INTO `category_info` VALUES (5, 1, '智慧屏电视', 'https://img11.360buyimg.com/n1/jfs/t1/107975/22/57903/111261/672ca34bFd0264716/832fe66f08b3bb76.jpg.avif', NULL, NULL, NULL, 3, 1);
INSERT INTO `category_info` VALUES (6, 0, '手机/数码', 'https://img14.360buyimg.com/n1/s450x450_jfs/t1/186627/35/51187/75651/672acd7eF1e7f61a8/9f6495558e53e4a9.jpg.avif', NULL, '2022-07-27 15:07:31', '2022-07-27 15:08:57', 1, 2);
INSERT INTO `category_info` VALUES (7, 6, '手机通讯', 'https://img14.360buyimg.com/n1/s450x450_jfs/t1/186627/35/51187/75651/672acd7eF1e7f61a8/9f6495558e53e4a9.jpg.avif', NULL, '2022-07-27 15:08:41', '2022-07-27 15:09:34', 2, 2);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章（种草）表';

-- ----------------------------
-- Records of article_info
-- ----------------------------
BEGIN;
INSERT INTO `article_info` VALUES (1, 0, '华凌空调真不错!', '京东买的，真的种草了', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 1, 0, '这里是文章正文', '2022-07-19 11:47:59', '2022-07-19 11:48:52', '2022-07-19 11:49:13');
INSERT INTO `article_info` VALUES (2, 2, '华凌空调真不错!', '京东买的，真的种草了', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 0, 0, '这里是文章正文', '2022-07-19 11:49:36', '2022-07-31 15:51:06', '2022-07-31 16:08:59');
INSERT INTO `article_info` VALUES (3, 2, '华凌空调真不错a', '京东买的，真的种草了a', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 0, 0, '这里是文章正文a', '2022-07-31 15:42:45', '2022-07-31 15:42:45', NULL);
INSERT INTO `article_info` VALUES (4, 1, '华凌空调真不错a', '京东买的，真的种草了a', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 0, 0, '这里是文章正文a', '2022-07-31 15:44:25', '2022-07-31 15:44:25', NULL);
INSERT INTO `article_info` VALUES (5, 1, '华凌空调真不错', '京东买的，真的种草了', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 1, 0, '这里是文章正文', '2022-07-31 19:06:59', '2022-07-31 19:06:59', NULL);
INSERT INTO `article_info` VALUES (6, 2, '华凌空调真不错', '京东买的，真的种草了', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 1, 0, '这里是文章正文', '2022-07-31 19:07:08', '2022-07-31 19:07:08', NULL);
INSERT INTO `article_info` VALUES (7, 1, '华凌空调真不错', '京东买的，真的种草了', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgfs17.gomein.net.cn%2FT108VWB4W_1RCvBVdK_800.jpg%3Fv%3D1&refer=http%3A%2F%2Fgfs17.gomein.net.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1660794257&t=795ee536d5af33788a249b08d0b28b6f', 1, 0, '这里是文章正文', '2022-07-31 19:08:03', '2022-07-31 19:08:03', NULL);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- -------------------------
-- Records of praise_info
-- ----------------------------
BEGIN;
INSERT INTO `praise_info` VALUES (8, 1 , 2, 4, '2023-01-19 12:18:07', '2023-01-19 12:18:07');
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of comment_info
-- ----------------------------
BEGIN;
INSERT INTO `comment_info` VALUES (4, 0, 1, 5, 2, '好评 下次还会买', '2022-07-31 17:23:48', '2022-07-31 17:23:48', NULL);
INSERT INTO `comment_info` VALUES (5, 0, 1, 5, 2, '来个评论', '2022-07-31 17:24:10', '2022-07-31 17:24:10', NULL);
INSERT INTO `comment_info` VALUES (7, 5, 1, 5, 2, '来个评论', '2022-07-31 17:24:59', '2022-07-31 17:24:59', NULL);
INSERT INTO `comment_info` VALUES (10, 1, 4, 5, 1, 'labore', '2023-01-19 14:25:24', '2023-01-19 14:25:24', NULL);
INSERT INTO `comment_info` VALUES (11, 1, 4, 5, 1, 'xxxxx', '2023-01-19 14:26:50', '2023-01-19 14:26:50', NULL);
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
INSERT INTO `consignee_info` VALUES (1, 1, 1, '王先生1', '13269477632', '北京', '北京市', '房山区', '拱辰街道', '大学城西', '2022-07-31 14:42:33', '2022-07-31 14:44:50', NULL);
COMMIT;