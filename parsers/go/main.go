package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"sponsortext/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter sponsor text variables (or \"exit\" to stop): ")

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Failed to read input: %s\n", err)
			continue
		}

		if scanner.Text() == "exit" {
			break
		}

		parsed := internal.ParseSponsorTextVariables(scanner.Text())

		b, err := json.MarshalIndent(parsed, "", "  ")
		if err != nil {
			fmt.Printf("Failed to marshal parsed variables: %s\n", err)
			continue
		}

		fmt.Printf("Equivalent JSON: %s\n", string(b))
	}
}
