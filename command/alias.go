package command

import (
	"fmt"

	"github.com/cli/cli/internal/config"
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

	alias := args[0]

	out := colorableOut(cmd)

	expansion := processArgs(args)

	fmt.Fprintf(out, "- Adding alias for %s = %s\n", utils.Bold(alias), utils.Bold(expansion))

	if aliasExists(cfg, alias) {
		return fmt.Errorf("alias %s already exists", alias)
	}

	if !validCommand(expansion) {
		return fmt.Errorf("could not create alias: %s does not correspond to a gh command", utils.Bold(expansion))
	}

	// TODO set the alias on disk, probably going through config

	fmt.Printf("%#v\n", ctx)

	return nil
}

func validCommand(expansion string) bool {
	// TODO
	return false
}

func aliasExists(cfg config.Config, alias string) bool {
	// TODO
	return true
}

func processArgs(args []string) string {
	// TODO
	return "lol pbbbbt"
}
