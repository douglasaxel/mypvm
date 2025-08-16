package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func UseVersion(version string) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user directory.")
		panic(err)
	}
	mypvmFolder := filepath.Join(userHomeDir, ".mypvm")

	fmt.Printf("Setting version %s as default...\n", version)

	// 1. Check if the version is installed locally.
	versionFolder := filepath.Join(mypvmFolder, version)
	if _, err := os.Stat(versionFolder); os.IsNotExist(err) {
		return fmt.Errorf("version %s is not installed. Use the 'install' command first", version)
	}

	// 2. Try to remove the existing symbolic link, if any.
	if _, err := os.Lstat(versionFolder); err == nil {
		fmt.Println("Removing previous symbolic link...")
		if err := os.Remove(versionFolder); err != nil {
			return fmt.Errorf("error removing existing symbolic link: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking existing symbolic link: %v", err)
	}

	// 3. Create the new symbolic link.
	err = os.Symlink(versionFolder, mypvmFolder)
	if err != nil {
		return fmt.Errorf("error creating symbolic link for version %s: %v", version, err)
	}

	fmt.Printf("Symbolic link created successfully. Version %s is now the default.\n", version)
	fmt.Println("\nTo use this version in your terminal, add the following path to your PATH:")

	switch runtime.GOOS {
	case "windows":
		fmt.Printf("  %s\n", mypvmFolder)
		fmt.Printf("\nYou can do this temporarily with: 'set PATH=%%PATH%%;%s'", mypvmFolder)
		fmt.Println("\nOr permanently by adding the folder to Windows 'Environment Variables'.")
	case "linux", "darwin":
		fmt.Printf("  %s\n", mypvmFolder)
		fmt.Println("\nAdd the line below to your shell profile file (e.g., ~/.bashrc or ~/.zshrc):")
		fmt.Println(`  export PATH="` + mypvmFolder + `:$PATH"`)
		fmt.Println("Then run 'source ~/.bashrc' (or the corresponding file) or restart your terminal.")
	}

	return nil
}
