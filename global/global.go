package global

import (
	"server/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *zap.Logger
	DB     *gorm.DB
	Redis  redis.Client
	Es     *elasticsearch.TypedClient
)
