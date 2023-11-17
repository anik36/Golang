package infra

import (
	"context"
	db "go-framework/internal/pg"
	middleware "go-framework/internal/wscutils"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/remiges-tech/logharbour/logHarbour"
)

var (
	clientID      = "MCX"
	clientSecret  = "OFyNKbP4g6sYtR4nACtS4V30ILsruzY1"
	keycloakURL   = "http://0.0.0.0:8080/"
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 0
)

type AppConfig struct {
	DatabaseURL string `json:"database_url"`
	Port        int    `json:"port"`
	DBHost      string `json:"db_host"`
	DBPort      int    `json:"db_port"`
	DBUser      string `json:"db_user"`
	DBPassword  string `json:"db_password"`
	DBName      string `json:"db_name"`
}

// initInfraServices sets up required infrastructure services. All the database connections,
// logger, etc. are initialized here.
func InitInfraServices(config AppConfig) (*db.Provider, logHarbour.LogHandles, *redis.Client) {
	//pgConfig := db.Config{
	//	Host:     "localhost",
	//	Port:     5432,
	//	User:     "erptest",
	//	Password: "erptest",
	//	DBName:   "erptest",
	//}
	pgConfig := db.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DBName:   config.DBName,
	}
	dbProvider := db.NewProvider(pgConfig)
	lh := logHarbour.LogInit("app1", "module1", "system1")

	// Set up Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	return dbProvider, lh, rdb
}

func setupMiddleware(rdb *redis.Client) gin.HandlerFunc {
	// Set up Redis cache and Verifier
	redisCache := &middleware.RedisTokenCache{
		Client: rdb,
		Ctx:    context.Background(),
	}

	provider, _ := oidc.NewProvider(context.Background(), keycloakURL)
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})
	authMiddleware := middleware.NewAuthMiddleware(clientID, clientSecret, keycloakURL, verifier, redisCache)

	return authMiddleware.MiddlewareFunc()
}

// SetupRouter sets up the Gin router with middleware.
func SetupRouter(lh *logHarbour.LogHandles, rdb *redis.Client) *gin.Engine {
	authMiddleware := setupMiddleware(rdb)
	r := gin.Default()
	r.Use(authMiddleware)

	logger := &middleware.CustomLogger{}
	r.Use(middleware.CustomLoggerMiddleware(logger))

	return r
}
