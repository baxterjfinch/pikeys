package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"pikeys/internal/details"
	"pikeys/internal/mnemonic"
	"pikeys/utils/store"
)

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(detailsCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates A New Mnemonic",
	Long:  `Creates A New Mnemonic And Stores it In ~/mnemonics`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n Generating and Saving A New Mnemonic\n")
		newMnemonic, err := mnemonic.CreateNew()
		store.StoreMnemonic(newMnemonic)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Provides Details of Mnemonics",
	Long:  `Type In Your Mnemonic and the number of keypairs you want to view`,
	Run: func(cmd *cobra.Command, args []string) {
		details.GetDetails()
	},
}
