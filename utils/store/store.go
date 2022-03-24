package store

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pikeys/internal/mnemonic"
)

func StoreMnemonic(newMnemonic *mnemonic.PiKeysService) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	newpath := filepath.Join(dirname, "pikey")
	filename := "/mnemonics.csv"
	filePath, _ := filepath.Abs(newpath + filename)
	_, err = os.Stat(filename)

	readfile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer readfile.Close()

	readdata, err := csv.NewReader(readfile).ReadAll()
	if err != nil {
		log.Println(err)
		return
	}
	rowCount := len(readdata)

	writefile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return
	}
	defer writefile.Close()

	w := csv.NewWriter(writefile)
	defer w.Flush()
	count := fmt.Sprintf("%d", rowCount)
	mn := fmt.Sprintf("%s", newMnemonic.Mnemonic)

	row := []string{count, mn}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	fmt.Println("File updated successfully", filename)
}
