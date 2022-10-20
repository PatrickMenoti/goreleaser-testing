package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "azioncli",
	Long: "Root command",
}
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "azioncli hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!!!")
	},
}

var ByeCmd = &cobra.Command{
	Use:   "bye",
	Short: "azioncli goodbye",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bye!!!")
	},
}
var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	RootCmd.AddCommand(HelloCmd, ByeCmd, CompletionCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
