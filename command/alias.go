package command

import (
	"fmt"
	"strings"

	"github.com/cli/cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(aliasCmd)
	aliasCmd.AddCommand(aliasSetCmd)
}

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Create shortcuts for gh commands",
	Long:  `TODO`,
}

var aliasSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a shortcut for a gh command",
	Long:  `TODO`,
	Args:  cobra.MinimumNArgs(2),
	RunE:  aliasSet,
}

func aliasSet(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	cfg, err := ctx.Config()
	if err != nil {
		return err
	}

	aliasCfg, err := cfg.Aliases()
	if err != nil {
		return err
	}

	alias := args[0]

	out := colorableOut(cmd)

	expansion := processArgs(args[1:])

	fmt.Fprintf(out, "- Adding alias for %s = %s\n", utils.Bold(alias), utils.Bold(expansion))

	if aliasCfg.Exists(alias) {
		return fmt.Errorf("alias %s already exists", alias)
	}

	if !validCommand(expansion) {
		return fmt.Errorf("could not create alias: %s does not correspond to a gh command", utils.Bold(expansion))
	}

	// TODO set the alias on disk, probably going through config

	return nil
}

func validCommand(expansion string) bool {
	// TODO
	return true
}

func processArgs(args []string) string {
	if len(args) == 1 {
		return args[0]
	}

	return strings.Join(args, " ")
}
