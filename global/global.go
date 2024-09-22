package global

import (
	"github.com/cc14514/go-geoip2"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"template/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Es     *elasticsearch.TypedClient
	Log    *zap.SugaredLogger
	Etcd   *clientv3.Client
	AddrDB *geoip2.DBReader
)
