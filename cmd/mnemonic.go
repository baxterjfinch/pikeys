package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"pikeys/internal/mnemonic"
	"pikeys/utils/store"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates A New Mnemonic",
	Long:  `Creates A New Mnemonic And Stores it In ~/mnemonics`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n Generating and Saving A New Mnemonic\n")
		_, err := mnemonic.CreateNew()
		store.StoreMnemonic()
		if err != nil {
			log.Fatal(err)
		}
	},
}
