package store

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func StoreMnemonic() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	newpath := filepath.Join(dirname, "mnemonics")
	log.Printf(newpath)
	sec := time.Now().Unix()
	filename := fmt.Sprintf("/mnemonic_%d", sec)
	filePath, _ := filepath.Abs(newpath + filename)
	log.Printf(filePath)

	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("ERR: %s", err)
			return
		}
		defer file.Close()
	} else {
		fmt.Println("File already exists!", filename)
		return
	}

	fmt.Println("File created successfully", filename)
}
