package cmd

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
	db := &db{
		Driver:   "mysql",
		User:     "user",
		Password: "password",
		Database: "mission_reward",
		Address:  "localhost",

		Port: 3306,
	}

	server := &server{
		Address: "localhost",
		Port:    8080,
	}

	config := &Config{
		Verbose: true,
		DB:      db,
		Server:  server,
	}
	return config, nil
}
