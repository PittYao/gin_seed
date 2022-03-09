package redis

import (
	"flag"
	"fmt"
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/globalkey"
	"github.com/PittYao/gin_seed/app/common/zap"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
	"time"
)

// configFile path
var configFile = flag.String("c", "../../etc/config.yaml", "the config file")

func TestRedis(t *testing.T) {
	// load configFile
	flag.Parse()
	conf.MustLoad(*configFile, &global.CONFIG)

	// init zap logger
	global.LOG = zap.Zap()

	global.REDIS = Redis()

	// set
	err := global.REDIS.Set(globalkey.Ctx, "key", "value", 6*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// get
	val, err := global.REDIS.Get(globalkey.Ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// get error
	val2, err := global.REDIS.Get(globalkey.Ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

}
