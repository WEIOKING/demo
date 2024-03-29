package go_redis

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var redisClient *redis.Client
var redisClient2 *redis.Client
var redisOnce sync.Once
var redisOnce2 sync.Once
var ctx = context.Background()

const unlockScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end`

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         "190.92.208.145:6379",
			Password:     "7MW0f7e@",
			DB:           9,
			ReadTimeout:  600 * time.Second,
			WriteTimeout: 600 * time.Second,
		})
		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			fmt.Println("redis init error")
		}
	})
	return redisClient
}
func GetRedisClient2() *redis.Client {
	redisOnce2.Do(func() {
		redisClient2 = redis.NewClient(&redis.Options{
			Addr:         "10.8.0.102:6379",
			Password:     "7cAwfXvnzJp8xz3Q59nzHeCKEen8N2FN",
			DB:           9,
			ReadTimeout:  600 * time.Second,
			WriteTimeout: 600 * time.Second,
		})
		_, err := redisClient2.Ping(ctx).Result()
		if err != nil {
			fmt.Println("redis init error")
		}
	})
	return redisClient2
}

type RedisLock struct {
	lockKey     string
	randomValue string
	expireTime  time.Duration
	watchDog    chan int
}

func TryLock(lockKey string) (bool, *RedisLock) {
	randomValue := uuid.New().String()
	nx := GetRedisClient().SetNX(context.Background(), lockKey, randomValue, 30*time.Second)
	b := nx.Val()
	if b {
		r := &RedisLock{lockKey: lockKey, randomValue: randomValue, expireTime: 30 * time.Second, watchDog: make(chan int, 1)}
		go startWatchDog(r)
		return b, r
	}
	return b, nil
}

func UnLock(redisLock *RedisLock) error {
	result, err := GetRedisClient().Eval(context.Background(), unlockScript, []string{redisLock.lockKey}, redisLock.randomValue).Result()
	if err != nil {
		fmt.Println("unLock error")
		return err
	}
	close(redisLock.watchDog)
	if result.(int64) == 0 {
		fmt.Printf("unlock failed, lockKey:%s, randomValue:%s \n", redisLock.lockKey, redisLock.randomValue)
		return nil
	}
	fmt.Printf("unlock successful, lockKey:%s, randomValue:%s \n", redisLock.lockKey, redisLock.randomValue)
	return nil
}

func startWatchDog(redisLock *RedisLock) {
	ticker := time.NewTicker(redisLock.expireTime / 3)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ctx, cancelFunc := context.WithTimeout(context.Background(), redisLock.expireTime/3*2)
			result, err := GetRedisClient().Expire(ctx, redisLock.lockKey, redisLock.expireTime).Result()
			cancelFunc()
			if err != nil || !result {
				fmt.Printf("watchDog 续期失败, lockKey:%s, randomValue:%s \n", redisLock.lockKey, redisLock.randomValue)
				return
			}
			fmt.Printf("watchDog 续期成功, lockKey:%s, randomValue:%s \n", redisLock.lockKey, redisLock.randomValue)
		case <-redisLock.watchDog:
			fmt.Printf("watchDog 关闭成功, lockKey:%s, randomValue:%s \n", redisLock.lockKey, redisLock.randomValue)
			return
		}
	}
}
