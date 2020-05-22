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
	Use: "set",
	// TODO HACK: even when inside of a single-quoted string, cobra was noticing and parsing any flags
	// used in an alias expansion string argument. Since this command needs no flags, I disabled their
	// parsing. If we ever want to add flags to alias set we'll have to figure this out. I haven't
	// checked if this is fixed in a new cobra.
	DisableFlagParsing: true,
	Short:              "Create a shortcut for a gh command",
	Long:               `TODO`,
	Args:               cobra.MinimumNArgs(2),
	RunE:               aliasSet,
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
	expansion := processArgs(args[1:])

	out := colorableOut(cmd)
	fmt.Fprintf(out, "- Adding alias for %s: %s\n", utils.Bold(alias), utils.Bold(expansion))

	if aliasCfg.Exists(alias) {
		return fmt.Errorf("alias %s already exists", alias)
	}

	if !validCommand(expansion) {
		return fmt.Errorf("could not create alias: %s does not correspond to a gh command", utils.Bold(expansion))
	}

	err = aliasCfg.Add(alias, expansion)
	if err != nil {
		return fmt.Errorf("could not create alias: %s", err)
	}

	fmt.Fprintf(out, "%s Added alias.\n", utils.Green("✓"))

	return nil
}

func validCommand(expansion string) bool {
	// TODO root.traverse?
	return true
}

func processArgs(args []string) string {
	if len(args) == 1 {
		return args[0]
	}

	return strings.Join(args, " ")
}
