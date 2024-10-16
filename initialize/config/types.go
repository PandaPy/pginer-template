package config

// Config 包含所有应用的配置项
type Config struct {
	Server    ServerConfig        `mapstructure:"SERVER"`
	Databases map[string]DBConfig `mapstructure:"DATABASES"`
	Redis     RedisConfig         `mapstructure:"REDIS"`
}

// APPConfig 定义数据库配置
type ServerConfig struct {
	Listen       int      `mapstructure:"LISTEN"`
	Mode         string   `mapstructure:"MODE"`
	RouerPrefix  string   `mapstructure:"ROUTER_PREFIX"`
	AllowedHosts []string `mapstructure:"ALLOWED_HOSTS"`
	SecretKey    string   `mapstructure:"SECRET_KEY"`
}

// DBConfig 定义数据库配置
type DBConfig struct {
	Name         string `mapstructure:"NAME"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Host         string `mapstructure:"HOST"`
	Port         string `mapstructure:"PORT"`
	Config       string `mapstructure:"CONFIG"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	AutoMigrate  bool   `mapstructure:"AutoMigrate"`
}

func (m *DBConfig) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Name + "?" + m.Config
}

// RedisConfig 定义 Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
}
