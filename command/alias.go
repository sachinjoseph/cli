package command

import (
	"fmt"

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

	alias := args[0]

	// TODO process args

	// TODO check if overwriting
	// TODO check if actual command
	// TODO set the alias on disk, probably going through config

	fmt.Printf("%#v\n", ctx)

	return nil
}
