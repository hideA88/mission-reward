package test

import (
	"fmt"
	"github.com/hideA88/mission-reward/pkg"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

func NewTestConfig() (*pkg.Config, error) {
	fp := filepath.Join(ProjectRoot(), "./configs/config.toml")
	return pkg.ParseTestConfig(fp)
}

func NewTestDbConn(config *pkg.Config, logger *zap.SugaredLogger) *sqlx.DB {
	//connect DB
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
		config.DB.User,
		config.DB.Password,
		config.DB.Address,
		config.DB.Port,
		config.DB.Database)
	db, err := sqlx.Open(config.DB.Driver, dataSource)
	logger.Info(dataSource)
	if err != nil {
		logger.Fatal("db connection error: ", err)
		defer db.Close()
	}
	return db
}

// ProjectRoot returns the project root directory abs path.
func ProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}
