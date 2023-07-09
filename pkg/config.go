package pkg

import (
	"github.com/spf13/viper"
)

type Config struct {
	General *general
	DB      *db
	Server  *server
	Test    *test_
}

type general struct {
	Verbose bool
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

type test_ struct {
	DbPort int
}

func ParseConfig(filePath string) (*Config, error) {
	err := readInit(filePath)
	if err != nil {
		return nil, err
	}
	return readConf()
}

func ParseTestConfig(filePath string) (*Config, error) {
	err := readInit(filePath)
	if err != nil {
		return nil, err
	}
	conf, err := readConf()
	if err != nil {
		return nil, err
	}
	conf.DB.Port = viper.GetInt("test.db_port")
	return conf, nil
}

func readInit(filePath string) error {
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func readConf() (*Config, error) {
	general := &general{
		Verbose: viper.GetBool("general.verbose"),
	}

	db := &db{
		Driver:   viper.GetString("db.driver"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
		Address:  viper.GetString("db.address"),
		Port:     viper.GetInt("db.port"),
	}

	server := &server{
		Address: viper.GetString("server.address"),
		Port:    viper.GetInt("server.port"),
	}

	config := &Config{
		General: general,
		DB:      db,
		Server:  server,
	}
	return config, nil
}
