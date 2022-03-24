package details

import (
	"bufio"
	"fmt"
	"os"
	"pikeys/utils/mnemonic"
	"strconv"
	"strings"
)

func GetDetails() {
	var mn string
	var count int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n\nEnter Mnemonic")
	fmt.Println("------------------------------------------")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Println("------------------------------------------\n\n")

	// remove the delimeter from the string
	mn = strings.TrimSuffix(input, "\n")

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter How Many Accounts You Want To Visualize")
	fmt.Println("------------------------------------------")
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Println("------------------------------------------\n\n")

	// remove the delimeter from the string
	sCount := strings.TrimSuffix(input, "\n")
	intVar, err := strconv.Atoi(sCount)
	if err != nil {
		fmt.Println("An error occured while converting Accounts to integer. Please try again", err)
		return
	}
	count = intVar
	mnemonic.GetAccounts(mn, count)

}
