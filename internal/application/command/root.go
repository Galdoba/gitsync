package command

import (
	"fmt"

	"github.com/Galdoba/gitsync/internal/application/setup"
	"github.com/urfave/cli/v3"
)

func GitSync() (*cli.Command, error) {
	actx, err := setup.NewAppContext()
	if err != nil {
		return nil, fmt.Errorf("failed to setup infratsructure: %v", err)
	}
	cmd := cli.Command{
		Name:           "",
		Aliases:        []string{},
		Usage:          "sync repos to/with github",
		UsageText:      "",
		ArgsUsage:      "",
		Version:        "",
		Description:    "",
		DefaultCommand: "",
		Category:       "",
		Commands: []*cli.Command{
			Pull(actx),
		},
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
		Action:                          nil,
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

	return &cmd, nil
}
