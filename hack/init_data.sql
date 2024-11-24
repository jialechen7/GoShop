-- ----------------------------
-- data of admin_info
-- ----------------------------
INSERT INTO `admin_info` (`name`, `password`, `role_ids`, `user_salt`, `is_admin`, `created_at`, `updated_at`) VALUES
('jialechen', 'b23239a15dbae17bc78cc83cd3c1d071', '1,2', 'wUUjkJqFiM', 1, NOW(), NOW()),
('admin', '3e7665414906ba5c30d993dc12cec9ed', '1,2', 'CQggcu45qq', 1, NOW(), NOW());

-- ----------------------------
-- data of position_info
-- ----------------------------
INSERT INTO `position_info` (`pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`) VALUES
('https://example.com/img1.jpg', 'Product A', 'https://example.com/productA', 1, 101, NOW(), NOW()),
('https://example.com/img2.jpg', 'Product B', 'https://example.com/productB', 2, 102, NOW(), NOW());

-- ----------------------------
-- data of rotation_info
-- ----------------------------
INSERT INTO `rotation_info` (`pic_url`, `link`, `sort`, `created_at`, `updated_at`) VALUES
('https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', '8', 1, NOW(), NOW()),
('https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', '6', 2, NOW(), NOW());

-- ----------------------------
-- data of order_info
-- ----------------------------
INSERT INTO `order_info` (`number`, `user_id`, `pay_type`, `remark`, `pay_at`, `status`, `consignee_name`, `consignee_phone`, `consignee_address`, `price`, `coupon_price`, `actual_price`, `created_at`, `updated_at`) VALUES
('1659231316407832000111', 1, 1, '备注1', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 100, 9900, '2022-08-27 09:35:16', '2022-08-27 09:35:16'),
('1659231554317361000757', 1, 1, '备注2', NULL, 2, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 200, 9800, '2022-08-27 09:39:14', '2022-08-27 09:39:14'),
('1661603467832912000516', 1, 2, '', '2022-12-13 21:52:26', 3, '', '', '', 0, 0, 0, '2022-12-08 20:31:07', '2022-12-08 20:31:07'),
('1661603562656619000513', 1, 3, '放到快递柜', '2022-12-13 21:52:19', 4, '王先生', '13269477432', '北京丰台汽车博物馆', 0, 0, 0, '2022-12-09 20:32:42', '2022-12-09 20:32:42');

-- ----------------------------
-- data of role_info
-- ----------------------------
INSERT INTO `role_info` (`name`, `desc`, `created_at`, `updated_at`)
VALUES ('运营', '运营权限', '2022-09-25 10:35:52', '2022-09-25 10:35:52');

-- ----------------------------
-- data of permission_info
-- ----------------------------
INSERT INTO permission_info (name, path, created_at, updated_at) VALUES
('首页配置', '/homepageManager', '2022-09-25 15:03:01', '2022-09-25 15:03:43'),
('轮播图', '/homepageManager/bannerSwiper', NOW(), NOW());

-- ----------------------------
-- data of user_info
-- ----------------------------
INSERT INTO `user_info` (name, avatar, password, user_salt, sex, status, sign, secret_answer, created_at, updated_at, deleted_at)
VALUES
    ('jialechen', 'https://cdn.acwing.com/media/user/profile/photo/40815_lg_d5f4bcc813.JPG', '905d0c656e6f93d5e73fb00ad0702a41', 'JAUrpNobzs', 0, 1, '个性签名', '银河中学', '2022-07-28 17:19:42', '2022-07-31 19:25:01', NULL);


-- ----------------------------
-- data of category_info
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
-- data of article_info
-- ----------------------------
INSERT INTO `article_info` (`user_id`, `title`, `desc`, `pic_url`, `is_admin`, `praise`, `collection`, `detail`, `created_at`, `updated_at`)
VALUES (1, '探索MacBook Pro的性能', '这篇文章探讨了MacBook Pro的性能表现，适用于不同场景的使用分析', 'https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', 2, 10, 0,
        'MacBook Pro 是苹果公司推出的高端笔记本电脑，以其卓越的性能和精致的设计深受专业人士和创作者的喜爱。无论是在处理大型图形设计任务，还是在进行视频编辑和开发工作，MacBook Pro 都能够提供无与伦比的性能支持。特别是其配备的 M1 或 M2 芯片，结合高分辨率的 Retina 屏幕，使得 MacBook Pro 成为当前市场上最受欢迎的高性能笔记本之一。',
        NOW(), NOW()),
       (1, 'iPhone 14 Pro的全面评测', '本文将对iPhone 14 Pro进行全面评测，探讨其性能、拍照功能及创新设计', 'https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', 2, 20, 0,
        'iPhone 14 Pro 是苹果公司推出的最新一代旗舰智能手机，搭载了强大的 A16 Bionic 芯片，拥有极致的性能表现。其全新的动态岛设计不仅使得显示屏更加简洁美观，还增强了互动性。iPhone 14 Pro 配备了 48MP 主摄像头，支持更高质量的照片和视频拍摄，特别适合喜欢摄影的用户。此外，iPhone 14 Pro 的显示效果也极其出色，采用了 ProMotion 技术和 Always-On 屏幕，提供了更流畅的体验。',
        NOW(), NOW());

-- ----------------------------
-- data of comment_info
-- ----------------------------
INSERT INTO `comment_info` (`id`, `parent_id`, `user_id`, `object_id`, `type`, `content`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1, 0, 1, 2, 2, '好评 下次还会买', '2022-07-31 17:23:48', '2022-07-31 17:23:48', NULL),
    (2, 0, 1, 2, 2, '还行', '2022-07-31 17:24:10', '2022-07-31 17:24:10', NULL),
    (3, 2, 1, 2, 2, '真的还行吗?', '2022-07-31 17:24:59', '2022-07-31 17:24:59', NULL);

-- ----------------------------
-- data of consignee_info
-- ----------------------------
INSERT INTO `consignee_info`
VALUES
    (1, 1, 1, '炸鸡', '13269477632', '北京', '北京市', '房山区', '拱辰街道', '大学城西', '2022-07-31 14:42:33', '2022-07-31 14:44:50', NULL);

-- ----------------------------
-- data of goods_info
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
    ('原装苹果充电器', 'https://img11.360buyimg.com/n7/jfs/t1/120933/14/49345/20820/67331648F08cc5230/47180b9222541553.png.avif', 9900, 1, 4, 6, 'Apple', 200, 50, '充电器,配件', NOW(), NOW());

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

-- ----------------------------
-- data of goods_options_info
-- ----------------------------
INSERT INTO goods_options_info (goods_id, pic_url, name, price, stock, created_at, updated_at)
VALUES
    (1, 'https://img10.360buyimg.com/n1/jfs/t1/159795/9/49579/202482/672af451Ff0a61b34/f063016c41fa50de.jpg.avif', '大号沙发', 1999, 100, NOW(), NOW()),
    (2, 'https://img12.360buyimg.com/n1/jfs/t1/160186/19/47269/98348/66c7e925F87921646/c7209966aa477011.jpg.avif', '大号餐桌', 899, 100, NOW(), NOW()),
    (6, 'https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', 'iPhone 14Pro 256GB', 899900, 100, NOW(), NOW()),
    (6, 'https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', 'iPhone 14Pro 512GB', 1099900, 100, NOW(), NOW()),
    (6, 'https://img12.360buyimg.com/n7/jfs/t1/66219/21/25962/142412/66a4e56eF026a1b3f/0fafaee7a5def073.jpg.avif', 'iPhone 14Pro 1TB', 1299900, 100, NOW(), NOW()),
    (8, 'https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', 'MacBook Pro 16 512GB', 1799900, 50, NOW(), NOW()),
    (8, 'https://img12.360buyimg.com/n7/jfs/t1/224360/34/24095/23440/6726f94bFc20591e4/d25e4eafc3d38089.png.avif', 'MacBook Pro 16 1TB', 1999900, 50, NOW(), NOW());
-- ----------------------------
-- data of coupon_info
-- ----------------------------
INSERT INTO coupon_info (name, price, goods_ids, category_id, created_at, updated_at)
VALUES
    ('满1000减100', 10000, '8,9,10', 1, NOW(), NOW()),
    ('满2000减200', 20000, '8,9,10', 1, NOW(), NOW()),
    ('满3000减300', 30000, '8,9,10', 1, NOW(), NOW());
