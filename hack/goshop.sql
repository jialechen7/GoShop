-- Database creation
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
('https://example.com/rotation1.jpg', 'https://example.com/link1', 1, NOW(), NOW()),
('https://example.com/rotation2.jpg', 'https://example.com/link2', 2, NOW(), NOW());

-- Table creation for order_info
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
('1659231554317361000757', 1, 1, '备注2', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 200, 9800, '2022-08-27 09:39:14', '2022-08-27 09:39:14'),
('1661603467832912000516', 1, 0, '', '2022-12-13 21:52:26', 0, '', '', '', 0, 0, 0, '2022-12-08 20:31:07', '2022-12-08 20:31:07'),
('1661603562656619000513', 1, 1, '放到快递柜', '2022-12-13 21:52:19', 0, '王先生', '13269477432', '北京丰台汽车博物馆', 0, 0, 0, '2022-12-09 20:32:42', '2022-12-09 20:32:42');
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

-- 示例数据w
BEGIN;
INSERT INTO role_info (name, `desc`, created_at, updated_at) VALUES
 ('销售员', '负责管理订单和库存查看权限', NOW(), NOW()),
 ('客服人员', '负责查看订单及回复客户咨询权限', NOW(), NOW());
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
INSERT INTO permission_info (name, path, created_at, updated_at) VALUES
('文章1', 'admin.article.index', NOW(), NOW()),
('测试2', 'admin.test.index', NOW(), NOW());
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

BEGIN;
INSERT INTO role_permission_info (role_id, permission_id, created_at, updated_at) VALUES
(1, 1, NOW(), NOW()),
(1, 2, NOW(), NOW());
COMMIT;