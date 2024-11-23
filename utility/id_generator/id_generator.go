package id_generator

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	TimeBegin       = 1609430400   // 2021-01-01 00:00:00
	BITS_COUNTS     = 32           // 32位
	IncrementPrefix = "increment:" // redis自增key前缀
)

// NextId 基于Redis自增生成全局唯一ID
// 参数:
//   - key: 用于区分不同业务的标识
//
// 返回值:
//   - int64: 生成的全局唯一 ID
//
// ID 格式说明:
//   - 最高位为符号位，始终为 0（保证正数）。
//   - 时间戳部分占用 31 位，单位为秒，表示从固定时间起始的秒数。
//   - 自增序列部分占用 32 位，表示同一秒内的自增序列，确保在高并发情况下的唯一性。
func NextId(key string) int64 {
	// 1. 生成时间戳
	now := time.Now().Unix()
	now -= TimeBegin
	// 2. 生成序列号（雪花算法是机器号+序列号的模式）
	yy, mm, dd := time.Now().Date()
	ctx := context.TODO()
	gCount, err := g.Redis().Do(ctx, "INCR", IncrementPrefix+key+":"+fmt.Sprintf("%d:%02d:%02d", yy, mm, dd))
	if err != nil {
		return 0
	}
	count := gconv.Int64(gCount)
	// 3. 生成ID（位运算拼接ID）
	return now<<BITS_COUNTS | count
}
