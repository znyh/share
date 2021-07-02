package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/google/wire"
	mongodb "github.com/znyh/library/pkg/database/mongodb"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMongoDB)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)

	RedisDel(ctx context.Context, key string) (err error)
	RedisIsExist(ctx context.Context, key string) (isExist bool, err error)
	RedisLRem(ctx context.Context, key string, start int32, end string) (err error)
	RedisLRange(ctx context.Context, key string, start, end int32) (ret []string, err error)
	RedisLTrim(ctx context.Context, key string, start, end int32) (err error)
	RedisLPush(ctx context.Context, key, value string) (err error)
	RedisHKeys(ctx context.Context, key string) (ret []string, err error)
	RedisHExist(ctx context.Context, key string, filed int64) (isExist bool, err error)
	RedisHSet(ctx context.Context, key, file, value string) (err error)
	RedisHGet(ctx context.Context, key, file string) (ret string, err error)
	RedisHMSet(ctx context.Context, key string, m map[string]interface{}) (err error)
	RedisHGetAll(ctx context.Context, key string) (ret []string, err error)
	RedisHIncr(ctx context.Context, key, file, value string) (val int64, err error)
	RedisHDel(ctx context.Context, key, file string) (err error)
	RedisHLen(ctx context.Context, key string) (cnt int64, err error)
	RedisHGetAllStringMap(ctx context.Context, key string) (ret map[string]string, err error)
	RedisPublishServerMsg(ctx context.Context, serverChan string, msg []byte) (err error)
	RedisGetLock(ctx context.Context, key, value string, ExpireTime int64) (GetLock bool, err error)
	RedisPipeline() (p redis.Pipeliner)
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis *redis.Redis
	mdb   *mongodb.DB
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB, mdb *mongodb.DB) (d Dao, cf func(), err error) {
	d = &dao{
		redis: r,
		db:    db,
		mdb:   mdb,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	_ = d.redis.Close()
	_ = d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
