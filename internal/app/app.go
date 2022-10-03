package app

import (
	"encoding/json"
	"fmt"
	"github.com/fairytale5571/migrator/internal/models"
	"github.com/fairytale5571/migrator/pkg/database"
	"github.com/fairytale5571/migrator/pkg/logger"
	"os"
)

func Version() string {
	return "0.0.1"
}

type App struct {
	Logger *logger.Wrapper
	DB     *database.DB
	Config models.Config
}

var app *App

func NewApp() (*App, error) {
	a := &App{}
	a.Logger = logger.New("migrate_app")
	err := a.readConfig(&a.Config)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		a.Config.User,
		a.Config.Password,
		a.Config.Host,
		a.Config.Port,
		a.Config.Name)
	db, err := database.New(url)
	if err != nil {
		a.Logger.Fatalf("error create database: %v", err)
		return nil, err
	}
	a.DB = db
	return a, nil
}

func (a *App) readConfig(cfg *models.Config) error {
	file, err := os.ReadFile("@extensions/grc_config.json")
	if err != nil {
		a.Logger.Errorf("error read config file: %v", err)
		return err
	}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		a.Logger.Errorf("error unmarshal config file: %v", err)
		return err
	}
	return nil
}

func Migrate() string {
	var err error

	path := os.Getenv("MIGRATIONS_PATH")
	if path == "" {
		return "fail"
	}
	if app == nil {
		app, err = NewApp()
		if err != nil {
			return fmt.Sprintf("cant start app: %s", err.Error())
		}
	}
	app.DB.StartMigrate(path)
	return "Success"
}

// "migrate" callExtension "migrate"
