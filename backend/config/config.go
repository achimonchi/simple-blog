package config

type Config struct {
	Database *ConfigDB
	App      *ConfigApp
}

type ConfigApp struct {
	Port string
}

type ConfigDB struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string

	MaxIdleConns int
	MaxOpenConns int
	MaxLifeTime  int
	MaxIdleTime  int
}
