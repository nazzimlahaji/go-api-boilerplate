package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var errParameterNotFound = errors.New("Unable to find required parameter")

const (
	keyAppPort = "APP_PORT"

	keySentryDSN = "SENTRY_DSN"

	keyMinioEndpoint = "MINIO_ENDPOINT"
	keyMinioAccess   = "MINIO_ACCESS_KEY"
	keyMinioSecret   = "MINIO_SECRET_KEY"
	keyMinioBucket   = "MINIO_BUCKET_KEY"
	keyMinioSSLMode  = "MINIO_SSL_MODE"

	keyDBHost     = "DB_HOST"
	keyDBUser     = "DB_USER"
	keyDBPassword = "DB_PASSWORD"
	keyDBName     = "DB_NAME"
	keyDBPort     = "DB_PORT"
	keyDBSSLMode  = "DB_SSLMODE"
	keyDBTimeZone = "DB_TIMEZONE"
)

var keys = []string{
	keyAppPort,
	keySentryDSN,
	keyMinioEndpoint,
	keyMinioAccess,
	keyMinioSecret,
	keyMinioBucket,
	keyMinioSSLMode,
	keyDBHost,
	keyDBUser,
	keyDBPassword,
	keyDBName,
	keyDBPort,
	keyDBSSLMode,
	keyDBTimeZone,
}

type Config struct {
	AppPort       string
	SentryDSN     string
	MinioEndpoint string
	MinioAccess   string
	MinioSecret   string
	MinioBucket   string
	MinioSSLMode  bool
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	DBSSLMode     string
	DBTimeZone    string
}

func GetConfig() (*Config, error) {
	// // Set the file name of the configurations file
	// viper.SetConfigName(".env")
	// // Set the path to look for the configurations file
	// viper.AddConfigPath(".")
	// // Enable VIPER to read Environment Variables
	// viper.AutomaticEnv()

	viper.SetConfigFile(".env")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	for _, k := range keys {
		err := viper.BindEnv(k)
		if err != nil {
			return nil, fmt.Errorf("Error binding environment variable: %w", err)
		}
	}

	config := Config{
		AppPort:       viper.GetString(keyAppPort),
		SentryDSN:     viper.GetString(keySentryDSN),
		MinioEndpoint: viper.GetString(keyMinioEndpoint),
		MinioAccess:   viper.GetString(keyMinioAccess),
		MinioSecret:   viper.GetString(keyMinioSecret),
		MinioBucket:   viper.GetString(keyMinioBucket),
		MinioSSLMode:  viper.GetBool(keyMinioSSLMode),
		DBHost:        viper.GetString(keyDBHost),
		DBUser:        viper.GetString(keyDBUser),
		DBPassword:    viper.GetString(keyDBPassword),
		DBName:        viper.GetString(keyDBName),
		DBPort:        viper.GetString(keyDBPort),
		DBSSLMode:     viper.GetString(keyDBSSLMode),
		DBTimeZone:    viper.GetString(keyDBTimeZone),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c Config) Validate() error {
	requiredFields := map[string]string{
		keyAppPort:       c.AppPort,
		keySentryDSN:     c.SentryDSN,
		keyMinioEndpoint: c.MinioEndpoint,
		keyMinioAccess:   c.MinioAccess,
		keyMinioSecret:   c.MinioSecret,
		keyMinioBucket:   c.MinioBucket,
		keyDBHost:        c.DBHost,
		keyDBUser:        c.DBUser,
		keyDBPassword:    c.DBPassword,
		keyDBName:        c.DBName,
		keyDBPort:        c.DBPort,
		keyDBTimeZone:    c.DBTimeZone,
	}

	for key, value := range requiredFields {
		if value == "" {
			return fmt.Errorf("%w: %s", errParameterNotFound, key)
		}
	}

	return nil
}
