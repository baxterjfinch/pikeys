package mnemonic

import (
	"fmt"
	gethhdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/sync/errgroup"
	"log"
)

type PiKeysService struct {
	Mnemonic string
	Seed     []byte
}

func CreateNew() (*PiKeysService, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	//entropy, _ := bip39.NewEntropy(256)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//log.Printf(mnemonic)
	mnemonic := "canoe move tray sustain tank submit elder hidden joke slow rapid benefit lizard need laugh rally license violin hire cement brush fresh damp column"
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := gethhdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := gethhdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	pkey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	fmt.Println(pkey)

	path = gethhdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	pkey, err = wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
	fmt.Println(pkey)

	return &PiKeysService{
		Mnemonic: mnemonic,
		Seed:     seed,
	}, nil
}

func (pws *PiKeysService) Run(group *errgroup.Group) {
}
