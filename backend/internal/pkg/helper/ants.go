package helper

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/panjf2000/ants/v2"
)

var (
	pool *ants.Pool
	once sync.Once
)

func GetAnts() *ants.Pool {
	once.Do(func() {
		p, err := ants.NewPool(15) // 设置最大协程数
		if err != nil {
			glog.Errorf(context.TODO(), "ants.NewPool failed, err: %v", err)
			panic(err)
		}

		pool = p
	})

	return pool
}
