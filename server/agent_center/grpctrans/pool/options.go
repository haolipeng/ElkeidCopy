package pool

const (
	defaultPoolLength = 1024
)

type Config struct {
	PoolLength int //池子大小
}

func NewConfig() *Config {
	c := &Config{PoolLength: defaultPoolLength}
	return c
}
