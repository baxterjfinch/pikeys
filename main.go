package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	root "pikeys/cmd"
	"time"
)

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	newpath := filepath.Join(dirname, "mnemonics")
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("filename.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newpath := filepath.Join(dirname, "mnemonics")
	log.Printf(newpath)
	sec := time.Now().Unix()
	filename := fmt.Sprintf("/mnemonic_%d", sec)
	filePath, _ := filepath.Abs(newpath + filename)

	root.Execute()
}
