package global

import (
	"gin_gorm/config"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Es     *elasticsearch.TypedClient
	Log    *zap.SugaredLogger
	Etcd   *clientv3.Client
)
