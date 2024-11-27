-- 1. 参数列表
-- 1.1 优惠券id
local coupon_id = ARGV[1]
-- 1.2 用户id
local user_id = ARGV[2]
-- 1.3 优惠券库存前缀
local stock_prefix = KEYS[1]
-- 1.4 已抢用户前缀
local user_prefix = KEYS[2]
-- 1.5 消息队列
local user_coupon_stream = KEYS[3]

-- 2. 数据key
-- 2.1 优惠券库存key
local stock_key = stock_prefix..coupon_id
-- 2.2 已抢用户key
local user_key = user_prefix..coupon_id

-- 3. 脚本业务
-- 3.1 判断库存是否充足，库存不足返回1
if tonumber(redis.call('get', stock_key)) <= 0 then
    return 1  -- 库存不足，返回1
end

-- 3.2 判断用户是否已抢过，已抢过返回2
if redis.call('sismember', user_key, user_id) == 1 then
    return 2  -- 用户已抢过，返回2
end

-- 3.3 库存减1
redis.call('decrby', stock_key, 1)  -- 用 decrby 减少库存，避免使用 incrby

-- 3.4 用户加入已抢名单
redis.call('sadd', user_key, user_id)

-- 3.5 发送消息到消息队列
-- XADD stream.user_coupon * k1 v1 k2 v2 ...
redis.call('XADD', user_coupon_stream, '*', 'coupon_id', coupon_id, 'user_id', user_id)
return 0  -- 成功，返回0
