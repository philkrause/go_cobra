package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use: "greet", 
	Short:"Prints a greet message",
	Long: "Print a long greet message"
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World" + string.Join(args))
	},

	func init(){
		rootCmd.AddCommand(greetCmd)
	}
}