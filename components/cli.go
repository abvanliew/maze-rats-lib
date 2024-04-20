package components

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(spellCmd)
	rootCmd.AddCommand(testCmd)
}

var rootCmd = &cobra.Command{
	Use:              "maze-rats-cli",
	PersistentPreRun: startUp,
	Version:          VERSION,
	Short:            "build maze run elements",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var spellCmd = &cobra.Command{
	Use:     "spell",
	Aliases: []string{"spl", "sp", "spll"},
	Run:     spellExecute,
	Short:   "generate a spell name",
}

var testCmd = &cobra.Command{
	Use:   "test",
	Run:   testExecute,
	Short: "test reference tables",
}

func startUp(cmd *cobra.Command, args []string) {
	generateTables()
	openDiceBag()
}

func spellExecute(cmd *cobra.Command, args []string) {
	count := iterateCount(args)

	for i := 0; i < count; i++ {
		fmt.Println(Arcane.PickSpell())
	}
}

func testExecute(cmd *cobra.Command, args []string) {
	fmt.Println(tables["ReferenceTest"].pick())
}

func iterateCount(args []string) int {
	count := 1
	if len(args) > 0 {
		i, err := strconv.Atoi(args[0])
		if err == nil {
			count = i
		}
	}
	return count
}
