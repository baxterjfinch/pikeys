package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("in hugo")
	},
}

func Execute() {
	//ctx, cancel := context.WithCancel(context.Background())
	//group, ctx := errgroup.WithContext(ctx)
	//
	//service, err := mnemonic.NewPiKeysService(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//service.Run(group)
	//
	//eventChan := make(chan os.Signal, 1)
	//signal.Notify(eventChan, syscall.SIGINT, syscall.SIGTERM)
	//select {
	//case <-eventChan:
	//case <-ctx.Done():
	//}
	//cancel()
	//if err := group.Wait(); err != nil {
	//	panic(err)
	//}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
