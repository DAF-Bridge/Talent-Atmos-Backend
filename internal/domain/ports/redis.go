package ports

type RedisClient interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}