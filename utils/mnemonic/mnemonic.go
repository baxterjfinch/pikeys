package mnemonic

import (
	"fmt"
	gethhdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
)

func GetAccounts(mnemonic string, count int) {
	fmt.Printf("\nRetrieving %d accounts for mnemonic\n", count)
	fmt.Println("------------------------------------------\n")

	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := gethhdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		sprintPath := fmt.Sprintf("m/44'/60'/0'/0/%d", i)
		path := gethhdwallet.MustParseDerivationPath(sprintPath)

		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}
		pkey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Account %d: %s\n", i+1, account.Address.Hex())
		fmt.Printf("PKey %d: %s", i+1, pkey)
		fmt.Println("\n")
	}
	fmt.Println("------------------------------------------\n")

}
