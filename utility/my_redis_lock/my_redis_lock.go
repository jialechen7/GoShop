package my_redis_lock

import (
	"context"
	"goshop/utility"
	"os"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gredis"
)

const (
	LockKeyPrefix = "lock:"
)

type ILock interface {
	TryLock(seconds int) bool
	Unlock()
}

type RedisLock struct {
	name         string
	redisClient  *gredis.Redis
	uuid         string
	unlockScript string
}

func New(name string, redisClient *gredis.Redis) *RedisLock {
	file, err := os.ReadFile("hack/unlock.lua")
	if err != nil {
		return nil
	}
	return &RedisLock{
		name:         name,
		redisClient:  redisClient,
		uuid:         utility.GenerateUUIDWithoutDash(),
		unlockScript: string(file),
	}
}

// TryLock 尝试获取锁
func (l *RedisLock) TryLock(seconds int) bool {
	// UUID拼接当前线程的标识，防止锁被其他线程误删
	goroutineId := utility.GetGoroutineID()
	value := l.uuid + "-" + gconv.String(goroutineId)
	gBool, err := l.redisClient.Do(context.TODO(), "SET", LockKeyPrefix+l.name, value, "EX", seconds, "NX")
	if err != nil {
		return false
	}
	return gBool.Bool()
}

//// Unlock 释放锁
//func (l *RedisLock) Unlock() {
//	goroutineId := utility.GetGoroutineID()
//	value := l.uuid + "-" + gconv.String(goroutineId)
//	gValue, err := l.redisClient.Do(context.BACKGROUND(), "GET", LockKeyPrefix+l.name)
//	if err != nil {
//		return
//	}
//	if gValue.String() != value {
//		return
//	}
//	// 假设判断通过了，但是由于某些原因在执行DEL之前阻塞了，导致锁过期，此时DEL会删除其他线程的锁，因此需要使用Lua脚本保证原子性
//	_, _ = l.redisClient.Do(context.BACKGROUND(), "DEL", LockKeyPrefix+l.name)
//}

// Unlock 释放锁
func (l *RedisLock) Unlock() {
	// 调用Lua脚本保证原子性
	goroutineId := utility.GetGoroutineID()
	value := l.uuid + "-" + gconv.String(goroutineId)
	_, err := l.redisClient.Eval(context.Background(), l.unlockScript, 1, []string{LockKeyPrefix + l.name}, []interface{}{value})
	if err != nil {
		return
	}
}
