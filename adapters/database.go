package adapters

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"

	"github.com/chrisenytc/ullli/config"
)

var pool *redis.Pool

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialURL(config.Get().RedisUrl)

			if err != nil {
				log.Errorf("Fatal error on database connection: %s", err)
				return nil, err
			}

			log.Info("Database connection established successfully.")

			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")

			return err
		},
	}
}

func LoadDatabase() {
	// Define connection pool
	pool = newPool()

	log.Info("The application is now using the database connection.")
}

func GetConnection() redis.Conn {
	return pool.Get()
}

func GetUrl(key string) (string, error) {
	conn := GetConnection()

	result, err := redis.String(conn.Do("HGET", key, "url"))

	conn.Close()

	return result, err
}

func GetUrlData(key string) ([]string, error) {
	conn := GetConnection()

	result, err := redis.Strings(conn.Do("HGETALL", key))

	conn.Close()

	return result, err
}

func CheckShortCode(key string) (bool, error) {
	conn := GetConnection()

	result, err := redis.Bool(conn.Do("EXISTS", key))

	conn.Close()

	return result, err
}

func SaveUrl(key string, url string) (string, error) {
	conn := GetConnection()

	result, err := redis.String(conn.Do("HMSET", key, "url", url, "clicks", 0, "created_at", time.Now().Format(time.RFC3339)))

	conn.Close()

	return result, err
}

func CountClick(key string) (int, error) {
	conn := GetConnection()

	result, err := redis.Int(conn.Do("HINCRBY", key, "clicks", 1))

	conn.Close()

	return result, err
}
