package pkg

import (
	"github.com/spf13/viper"
)

type Config struct {
	Verbose bool
	DB      *db
	Server  *server
}

type db struct {
	Driver   string
	User     string
	Password string
	Database string
	Address  string
	Port     int
}

type server struct {
	Address string
	Port    int
}

// ParseConfig TODO implement read config file
func ParseConfig() (*Config, error) {
	viper.SetConfigFile("./configs/config.toml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	db := &db{
		Driver:   viper.GetString("db_driver"),
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
		Database: viper.GetString("db_database"),
		Address:  viper.GetString("db_address"),
		Port:     viper.GetInt("db_port"),
	}

	server := &server{
		Address: viper.GetString("server_address"),
		Port:    viper.GetInt("server_port"),
	}

	config := &Config{
		Verbose: true,
		DB:      db,
		Server:  server,
	}
	return config, nil
}
