package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListLocalVersions() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user directory.")
		panic(err)
	}
	mypvmFolder := filepath.Join(userHomeDir, ".mypvm")

	fmt.Println("Listing PHP versions installed locally...")

	if _, err := os.Stat(mypvmFolder); os.IsNotExist(err) {
		fmt.Println("No versions installed locally.")
		return
	}

	directories, err := os.ReadDir(mypvmFolder)
	if err != nil {
		fmt.Println("Error listing locally installed versions.")
		return
	}

	isFoundVersion := false
	for _, dir := range directories {
		if dir.IsDir() {
			isFoundVersion = true
			fmt.Printf(" - Installed version: %s\n", dir.Name())
		}
	}

	if !isFoundVersion {
		fmt.Println("No versions installed locally.")
	}
}
