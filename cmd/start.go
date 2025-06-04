package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia o monitoramento de conectividade",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoramento iniciado... (implementação futura de ping)")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
