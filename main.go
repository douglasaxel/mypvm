package main

import (
	"fmt"
	"log"
	"mypvm/functions"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Welcome to PHP Version Manager!")
		fmt.Println("Usage: mypvm [command] [arguments]")
		fmt.Println("\n\nAvailable commands:")
		fmt.Println("\nlist   - Lists PHP versions available online")
		fmt.Println("list-local   - Lists PHP versions installed locally")
		fmt.Println("install - Installs a specific PHP version")
		fmt.Println("remove - Removes a specific PHP version")
		fmt.Println("use    - Selects a specific PHP version")

		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		functions.ListOnlineVersions()
	case "local":
		functions.ListLocalVersions()
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Please specify the PHP version to install. Example: 'mypvm install 8.3.0'")
			os.Exit(1)
		}
		version := os.Args[2]
		if err := functions.InstallVersion(version); err != nil {
			log.Fatalf("Installation error: %v", err)
		}
	case "remove":
		fmt.Println("Removing a specific PHP version...")
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("Please specify the PHP version to use. Example: 'mypvm use 8.3.0'")
			os.Exit(1)
		}
		version := os.Args[2]
		if err := functions.UseVersion(version); err != nil {
			log.Fatalf("Selection error: %v", err)
		}
	default:
		fmt.Println("Invalid command!")
		os.Exit(1)
	}
}
