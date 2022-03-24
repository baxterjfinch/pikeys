package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	root "pikeys/cmd"
)

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	newpath := filepath.Join(dirname, "pikey")
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	filename := "/mnemonics.csv"
	filePath, _ := filepath.Abs(newpath + "/mnemonics.csv")
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("ERR: %s", err)
			return
		}
		w := csv.NewWriter(file)
		defer w.Flush()
		row := []string{"ID", "MNEMONIC"}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
		defer file.Close()
	} else {
		fmt.Println("File already exists!", filename)
		return
	}

	//file, err := os.Open("filename.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()

	//newpath := filepath.Join(dirname, "mnemonics")
	//log.Printf(newpath)
	//sec := time.Now().Unix()
	//filename := fmt.Sprintf("/mnemonic_%d", sec)
	//filePath, _ := filepath.Abs(newpath + filename)

	root.Execute()
}
