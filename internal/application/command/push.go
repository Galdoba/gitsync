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

func Push(actx *setup.AppContext) *cli.Command {
	return &cli.Command{
		Name:                            "push",
		Aliases:                         []string{},
		Usage:                           "push changes to git",
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
		Action:                          push(actx),
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

func push(actx *setup.AppContext) cli.ActionFunc {
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
			switch err := s.Push(); err {
			case nil:
				logger.Infof("%v pushed to %v", repo, path)
			default:
				logger.Warnf("failed to push %v: %v", repo, err.Error())
			}
		}

		return nil
	}
}
