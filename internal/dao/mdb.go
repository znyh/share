package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	mg "github.com/znyh/middle-end/library/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (

	//database name
	shareDatabase = "share"

	//Collection name
	Guestshare = "Guestshare"
)

var (
	CollectionMap map[string]*mongo.Collection
)

func NewMongoDB() (db *mg.DB, err error) {
	var cfg struct {
		Client *mg.Config
	}

	if err = paladin.Get("mdb.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}

	db = mg.NewMongoDB(cfg.Client)
	InitCollectionMap(db)
	return
}

func InitCollectionMap(db *mg.DB) {
	CollectionMap = make(map[string]*mongo.Collection)
	database := db.Database(shareDatabase)
	CollectionMap[Guestshare] = database.Collection(Guestshare)
}

func (d *dao) MongoUpdate(ctx context.Context, table string, filter, update interface{}) (err error) {

	if v, ok := CollectionMap[table]; ok {
		opts := new(options.UpdateOptions)
		opts.SetUpsert(true)
		_, err = d.mdb.UpdateOne(ctx, v, filter, update, opts)
	}
	return
}
