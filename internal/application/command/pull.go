package command

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/gitsync/internal/application/setup"
	"github.com/Galdoba/gitsync/internal/domain/values/syncer"
	"github.com/urfave/cli/v3"
)

func Pull(actx *setup.AppContext) *cli.Command {
	return &cli.Command{
		Name:                            "pull",
		Aliases:                         []string{},
		Usage:                           "set local repos to pull changes",
		UsageText:                       "",
		ArgsUsage:                       "",
		Version:                         "",
		Description:                     "",
		DefaultCommand:                  "",
		Category:                        "",
		Commands:                        []*cli.Command{},
		Flags:                           []cli.Flag{},
		HideHelp:                        false,
		HideHelpCommand:                 false,
		HideVersion:                     false,
		EnableShellCompletion:           false,
		ShellCompletionCommandName:      "",
		ShellComplete:                   nil,
		ConfigureShellCompletionCommand: nil,
		Before:                          nil,
		After:                           nil,
		Action:                          pull(actx),
		CommandNotFound:                 nil,
		OnUsageError:                    nil,
		InvalidFlagAccessHandler:        nil,
		Hidden:                          false,
		Authors:                         []any{},
		Copyright:                       "",
		Reader:                          nil,
		Writer:                          nil,
		ErrWriter:                       nil,
		ExitErrHandler:                  nil,
		Metadata:                        map[string]interface{}{},
		ExtraInfo: func() map[string]string {
			panic("TODO")
		},
		CustomRootCommandHelpTemplate: "",
		SliceFlagSeparator:            "",
		DisableSliceFlagSeparator:     false,
		UseShortOptionHandling:        false,
		Suggest:                       false,
		AllowExtFlags:                 false,
		SkipFlagParsing:               false,
		CustomHelpTemplate:            "",
		PrefixMatchCommands:           false,
		SuggestCommandFunc:            nil,
		MutuallyExclusiveFlags:        []cli.MutuallyExclusiveFlags{},
		Arguments:                     []cli.Argument{},
		ReadArgsFromStdin:             false,
		StopOnNthArg:                  new(int),
	}
}

func pull(actx *setup.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		cfg := actx.Config
		logger := actx.Logger
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %v", err)
		}
		for repo, path := range cfg.TrackRepos {
			path = strings.ReplaceAll(path, "~", home)
			s := syncer.NewSync(repo, path)
			switch err := s.Pull(); err {
			case nil:
				logger.Infof("%v pulled to %v", repo, path)
			default:
				logger.Warnf("failed to pull %v: %v", repo, err.Error())
				fmt.Println(err)
			}
		}

		return nil
	}
}
