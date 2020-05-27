package dao

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	SET_IF_NOT_EXIST     = "NX" // 不存在则设置
	SET_WITH_EXPIRE_TIME = "PX" // 过期时间(毫秒)
	SET_LOCK_SUCCESS     = "OK" // 成功
	UN_LOCK_SUCCESS      = 1    // 删除锁成功
	UN_LOCK_NON_EXISTENT = 0    // 删除锁时,锁不存在
)

// NewRedisPool
func NewRedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			conn, e =  redis.Dial("tcp","localhost:6379")
			if e != nil {
				panic(conn)
			}
			return
		},
		MaxIdle:         50,
		MaxActive:       2000,
		IdleTimeout:     180 * time.Second,
	}
}


/*
    设锁和解锁保证是一个客户端请求
 */
func (d *Dao) SetLock(key,requestId string,ex int) bool {
	conn := d.RdPool.Get()
	defer conn.Close()
	msg, err := redis.String(
		conn.Do("SET",key,requestId,SET_IF_NOT_EXIST,SET_WITH_EXPIRE_TIME,ex),
	)
	if err != redis.ErrNil && err != nil {
		d.Logger.Printf("d.SetLock key(%s) requestId(%s) ex(%d) err(%v)", key, requestId, ex, err)
	}
	if msg == SET_LOCK_SUCCESS {
		return true
	}
	return false
}

// 获得分布式锁值
func (d *Dao) GetLock(conn redis.Conn, key string) string {
	msg, err := redis.String(conn.Do("GET",key))
	if err != redis.ErrNil && err != nil {
		d.Logger.Printf("d.GetLock key(%s) err(%v)", key, err)
	}
	return msg
}

// 只能自己删除自己的锁 否则因为操作超时可能会删除别人的锁
func (d *Dao) UnLock(key ,requestId string) bool{
	conn := d.RdPool.Get()
	defer conn.Close()
	if d.GetLock(conn, key) == requestId {
		msg, err := redis.Int64(conn.Do("DEL",key))
		if err != redis.ErrNil && err != nil {
			d.Logger.Printf("d.UnLock key(%s) requestId(%s) err(%v)", key, requestId, err)
		}
		// 避免操作时间过长,自动过期时再删除返回结果为0
		if msg == UN_LOCK_SUCCESS || msg == UN_LOCK_NON_EXISTENT {
			return true
		}
		return false
	}
	return false
}