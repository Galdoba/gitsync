package setup

import (
	"fmt"
	"path/filepath"

	"github.com/Galdoba/appcontext/configmanager"
	"github.com/Galdoba/appcontext/logmanager"
	"github.com/Galdoba/appcontext/pathspec"
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
	cfgPath := xdg.Location(xdg.ForConfig(), xdg.WithProgramName(constants.APP_NAME), xdg.WithFileName("config.toml"))
	logPath := xdg.Location(xdg.ForState(), xdg.WithProgramName(constants.APP_NAME), xdg.WithSubDir([]string{"logs"}), xdg.WithFileName(constants.APP_NAME+".log"))
	layout, err := pathspec.NewLayout(constants.APP_NAME, []pathspec.Path{
		pathspec.NewCustomPath(pathspec.ConfigFileTemplate, pathspec.WithName(filepath.Base(cfgPath))),
		pathspec.NewCustomPath(pathspec.LogFileTemplate, pathspec.WithName(filepath.Base(logPath))),
	})
	if err := layout.Generate(); err != nil {
		return nil, fmt.Errorf("failed to generate program file layout: %v", err)
	}
	issues, err := layout.Assess()
	for _, issue := range issues {
		fmt.Printf("layout issue detected: %v\n", issue)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to asses program file layout: %v", err)
	}

	cm, err := configmanager.New(constants.APP_NAME, config.Default(), configmanager.ForcePath(cfgPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create config manager: %v", err)
	}
	if err := cm.Load(); err != nil {
		if err := cm.Save(); err != nil {
			return nil, fmt.Errorf("failed to save new config")
		}
		if err := cm.Load(); err != nil {
			return nil, fmt.Errorf("failed to load new config")
		}
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
