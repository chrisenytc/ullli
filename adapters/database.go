package adapters

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"

	"github.com/chrisenytc/ullli/config"
)

var connection redis.Conn

func LoadDatabase() {
	// Define client
	conn, err := redis.DialURL(config.Get().RedisUrl)

	if err != nil {
		log.Panicf("Fatal error on database connection: %s", err)
	}

	connection = conn

	log.Info("Database connection established successfully.")
	log.Info("The application is now using the database connection.")
}

func GetConnection() redis.Conn {
	return connection
}

func GetUrl(key string) (string, error) {
	return redis.String(connection.Do("HGET", key, "url"))
}

func GetUrlData(key string) ([]string, error) {
	return redis.Strings(connection.Do("HGETALL", key))
}

func CheckShortCode(key string) (bool, error) {
	return redis.Bool(connection.Do("EXISTS", key))
}

func SaveUrl(key string, url string) (string, error) {
	return redis.String(connection.Do("HMSET", key, "url", url, "clicks", 0, "created_at", time.Now().Format(time.RFC3339)))
}

func CountClick(key string) (int, error) {
	return redis.Int(connection.Do("HINCRBY", key, "clicks", 1))
}
