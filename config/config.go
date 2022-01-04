package config

const (
	Environment        = "ENVIRONMENT"
	DBConnectionString = "DB_CONNECTION_STRING"
	JWTSecretKey       = "JWT_SECRET_KEY"
	DBName             = "DATABASE_NAME"
	DBCollection       = "DB_COLLECTION"
)

type Config struct{}

func Init() *Config {
	return &Config{}
}

func (c *Config) Environment() string {
	return getStringOrDefault(Environment, "development")
}

func (c *Config) DBConnectionString() string {
	return getStringOrDefault(DBConnectionString, "")
}

func (c *Config) JWTSecretKey() string {
	return getStringOrDefault(JWTSecretKey, "")
}

func (c *Config) DBName() string {
	return getStringOrDefault(DBName, "")
}

func (c *Config) DBCollection() string {
	return getStringOrDefault(DBCollection, "")
}
