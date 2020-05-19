package command

import "github.com/spf13/cobra"

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
	Use: "set",
}
