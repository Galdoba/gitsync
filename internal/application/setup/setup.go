package setup

import (
	"fmt"

	"github.com/Galdoba/appcontext/configmanager"
	"github.com/Galdoba/appcontext/logmanager"
	"github.com/Galdoba/appcontext/xdg"
	"github.com/Galdoba/gitsync/internal/application/constants"
	"github.com/Galdoba/gitsync/internal/infrastructure/config"
	"github.com/Galdoba/gitsync/internal/infrastructure/logger"
)

type AppContext struct {
	Config config.Config
	Logger logger.Logger
}

func NewAppContext() (*AppContext, error) {
	actx := AppContext{}
	cm, err := configmanager.New(constants.APP_NAME, config.Default())
	if err != nil {
		return nil, fmt.Errorf("failed to create config manager: %v", err)
	}
	if err := cm.Load(); err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}
	actx.Config = cm.Config()

	actx.Logger = logmanager.New(
		logmanager.WithHandlers(
			logmanager.NewHandler(
				xdg.Location(xdg.ForState(), xdg.WithProgramName(constants.APP_NAME), xdg.WithSubDir([]string{"logs"}), xdg.WithFileName(constants.APP_NAME+".log")),
				logmanager.LevelDebug,
				logmanager.NewTextFormatter(
					logmanager.WithLevelTag(true),
					logmanager.WithTimePrecision(3),
				),
			),
		),
		logmanager.WithLevel(logmanager.LevelDebug),
	)
	return &actx, nil
}
