package id_generator

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNextId 测试NextId方法生成全局唯一ID
func TestNextId(t *testing.T) {
	// 使用300个线程，每个线程调用NextId("test")创建全局ID，并打印输出
	var wg sync.WaitGroup
	threadCount := 300

	// 用于保证输出的顺序和线程安全
	var mu sync.Mutex
	var count sync.Map

	for i := 0; i < threadCount; i++ {
		wg.Add(1) // 在启动 goroutine 之前增加计数
		go func(threadID int) {
			defer wg.Done() // 确保任务完成时减少计数
			for j := 0; j < 100; j++ {
				id := NextId("test")
				count.Store(id, true)
				mu.Lock() // 加锁以确保打印的顺序性和安全性
				fmt.Printf("Thread %d - ID: %d\n", threadID, id)
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	length := 0
	count.Range(func(key, value interface{}) bool {
		length++
		return true // 继续遍历
	})

	require.Equal(t, threadCount*100, length)
}
