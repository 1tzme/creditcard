package features

import (
	"fmt"
	"os"
)

func Generate(args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter creditcard number.")
		os.Exit(1)
	}

	pick := false
	sample := ""
	
	if args[0] == "--pick" {
		pick = true

		if len(args) < 2 {
			fmt.Println("Please enter creditcard number.")
			os.Exit(1)
		}
		if len(args) > 2 {
			fmt.Println("Too many arguments")
			os.Exit(1)
		}
	}

	starsCount := strings.Count()
}