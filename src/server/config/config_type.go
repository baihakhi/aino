package config

import "time"

type DefaultConfig struct {
	Apps Apps
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Datasource struct {
	Url               string        `mapstructure:"url"`
	Port              string        `mapstructure:"port"`
	DatabaseName      string        `mapstructure:"databaseName"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Schema            string        `mapstructure:"schema"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
	MaxIdleConnection int           `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int           `mapstructure:"maxIdleConnection"`
	DebugMode         bool          `mapstructure:"debugMode"`
}

type Service struct {
	Gcp GcpService `mapstructure:"gcp"`
}

type GcpService struct {
	CredentialPath string `mapstructure:"credentialPath"`
	ProjectId      string `mapstructure:"projectId"`
	BucketName     string `mapstructure:"bucketName"`
}

type Role struct {
	Name       string   `mapstructure:"name"`
	Status     string   `mapstructure:"status"`
	ColumnName []string `mapstructure:"column_name"`
}
