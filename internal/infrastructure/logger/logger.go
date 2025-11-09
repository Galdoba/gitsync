package logger

type Logger interface {
	Fatalf(string, ...any)
	Errorf(string, ...any)
	Warnf(string, ...any)
	Infof(string, ...any)
	Debugf(string, ...any)
	Noticef(string, ...any)
}

// func NewCharmLogger() (*log.Logger, error) {
// 	path := xdg.Location(
// 		xdg.ForState(),
// 		xdg.WithProgramName(constants.APP_NAME),
// 		xdg.WithSubDir([]string{"logs"}),
// 		xdg.WithFileName(constants.APP_NAME+".log"),
// 	)
// 	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open log writer: %v", err)
// 	}
// 	defer f.Close()
// 	logger := log.New(f)
// 	logger.SetTimeFormat(time.DateTime)
// 	logger.Errorf("test message with string %v", "STRING")
// 	return logger, nil
// }
