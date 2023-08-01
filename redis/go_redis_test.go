package go_redis

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strconv"
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

func TestCopy(t *testing.T) {
	client := GetRedisClient()
	client2 := GetRedisClient2()
	keys := client.Keys(context.Background(), "*")
	for _, key := range keys.Val() {
		exists := client.Exists(context.Background(), key)
		if exists.Val() > 0 {
			ttl := client.TTL(context.Background(), key)
			dump := client.Dump(context.Background(), key)
			replace := client2.RestoreReplace(context.Background(), key, ttl.Val(), dump.Val())
			err := replace.Err()
			if err != nil {
				fmt.Println("key:(" + key + ")" + " RestoreReplace error:" + err.Error() + "          ttl:" + ttl.Val().String())
			}
		}
	}
}

func TestMultiCopy(t *testing.T) {
	client := GetRedisClient()
	client2 := GetRedisClient2()
	keys := client.Keys(context.Background(), "*")
	strings := keys.Val()
	length := len(strings)
	var wg = sync.WaitGroup{}
	mutex := sync.Mutex{}
	index := 0
	for i := 0; i < 100; i++ {
		i1 := i
		wg.Add(1)
		go func() {
			for true {
				mutex.Lock()
				if index >= length {
					mutex.Unlock()
					break
				}
				fmt.Println(index)
				key := strings[index]
				index++
				mutex.Unlock()
				exists := client.Exists(context.Background(), key)
				if exists.Val() > 0 {
					ttl := client.TTL(context.Background(), key)
					dump := client.Dump(context.Background(), key)
					replace := client2.RestoreReplace(context.Background(), key, ttl.Val(), dump.Val())
					err := replace.Err()
					if err != nil {
						fmt.Println("key:(" + key + ")" + " RestoreReplace error:" + err.Error() + "          ttl:" + ttl.Val().String())
					}
				}
			}
			fmt.Println("end " + strconv.Itoa(i1))
			wg.Done()
		}()
	}
	wg.Wait()
}
func TestRestore(t *testing.T) {
	client := GetRedisClient()
	client2 := GetRedisClient2()
	//key := "space:player:game:default:game:LIST"
	key := "space:player:game:default2:game:LIST"
	exists := client.Exists(context.Background(), key)
	if exists.Val() > 0 {
		ttl := client.TTL(context.Background(), key)
		dump := client.Dump(context.Background(), key)
		replace := client2.RestoreReplace(context.Background(), key, ttl.Val(), dump.Val())
		err := replace.Err()
		if err != nil {
			fmt.Println("key:(" + key + ")" + " RestoreReplace error:" + err.Error() + "          ttl:" + ttl.Val().String())
		}
	}
}
