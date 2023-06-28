package go_redis

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	client := GetRedisClient()
	var key = "go"
	err := client.Set(context.Background(), key, "hello world", 100*time.Second).Err()
	if err != nil {
		fmt.Println("redis set error")
	}
	result, err := client.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("redis get error")
	}
	fmt.Println(result)
	client.Del(context.Background(), key)
	result1, err1 := client.Get(context.Background(), key).Result()
	if err1 != nil {
		fmt.Println("redis get error")
	}
	fmt.Println(result1)

	err = client.MSet(context.Background(), "1", "1", "2", "2", "3", "3").Err()
	if err != nil {
		fmt.Println("redis mset error")
	}
	i, err := client.MGet(context.Background(), "1", "3", "2").Result()
	if err != nil {
		fmt.Println("redis mget error")
	}
	for _, val := range i {
		fmt.Println(val)
	}
	client.Del(context.Background(), "1", "2", "3")
}

func TestTryLock(t *testing.T) {
	var lockKey = uuid.New().String()
	var wg = sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tryLock, redisLock := TryLock(lockKey)
			if !tryLock {
				fmt.Println("try lock failed")
				return
			}
			defer UnLock(redisLock)
			time.Sleep(50 * time.Second)
		}()
	}
	wg.Wait()
}
