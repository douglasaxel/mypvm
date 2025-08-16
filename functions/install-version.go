package functions

import (
	"fmt"
	"io"
	"log"
	"mypvm/utils"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// ProgressReader is a wrapper for io.Reader that shows download progress
type ProgressReader struct {
	r              io.Reader
	totalSize      int64
	downloadedSize int64
	lastPercent    int
}

// Read implements the io.Reader interface and updates progress
func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.r.Read(p)
	pr.downloadedSize += int64(n)

	// Calculate and show progress
	var percent int
	if pr.totalSize > 0 {
		percent = int((float64(pr.downloadedSize) / float64(pr.totalSize)) * 100)
	} else {
		// If we don't have the total size, show downloaded bytes
		fmt.Printf("\rDownloading... %d bytes", pr.downloadedSize)
		return
	}

	// Update the progress bar only when the percentage changes
	if percent != pr.lastPercent {
		pr.lastPercent = percent

		// Create a visual progress bar
		width := 50
		bar := make([]byte, width)
		for i := 0; i < width; i++ {
			if i < width*percent/100 {
				bar[i] = '='
			} else {
				bar[i] = ' '
			}
		}

		// Show progress in MB
		downloadedMB := float64(pr.downloadedSize) / 1024 / 1024
		totalMB := float64(pr.totalSize) / 1024 / 1024
		fmt.Printf("\r[%s] %d%% (%.2f MB / %.2f MB)", string(bar), percent, downloadedMB, totalMB)
	}

	return
}

func InstallVersion(version string) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user directory.")
		panic(err)
	}
	mypvmFolder := filepath.Join(userHomeDir, ".mypvm")

	fmt.Printf("Starting installation of version %s...\n", version)

	versionFolder := filepath.Join(mypvmFolder, version)
	if _, err := os.Stat(versionFolder); !os.IsNotExist(err) {
		fmt.Printf("Version %s is already installed.\n", version)
		return nil
	}

	var urlDownload, fileType string

	switch runtime.GOOS {
	case "windows":
		var arch string
		if runtime.GOARCH == "amd64" {
			arch = "x64"
		} else {
			arch = "x86"
		}
		var compiler string
		switch version[0] {
		case '7':
			compiler = "vc15"
		case '8':
			compiler = "vs16"
			if len(version) > 2 && version[2] == '4' {
				compiler = "vs17"
			}
		}

		urlDownload = fmt.Sprintf("https://windows.php.net/downloads/releases/php-%s-nts-Win32-%s-%s.zip", version, compiler, arch)
		fileType = ".zip"
	case "linux", "darwin": // macOS
		urlDownload = fmt.Sprintf("https://www.php.net/distributions/php-%s.tar.gz", version)
		fileType = ".tar.gz"
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	fmt.Printf("Downloading from: %s\n", urlDownload)

	// Iniciar o request HTTP
	req, err := http.NewRequest("GET", urlDownload, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download error. Server returned: %s", resp.Status)
	}

	// Get the total file size
	totalSize := resp.ContentLength

	// Create the temporary file
	tempFile := version + fileType
	outFile, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}

	defer func() {
		if rErr := os.Remove(tempFile); rErr != nil {
			log.Printf("Warning: Could not remove temporary file %s: %v", tempFile, rErr)
		}
	}()

	// Create a reader with progress
	progressReader := &ProgressReader{
		r:              resp.Body,
		totalSize:      totalSize,
		downloadedSize: 0,
		lastPercent:    -1,
	}

	// Copy from the progress reader to the file
	_, err = io.Copy(outFile, progressReader)
	if err != nil {
		outFile.Close()
		return fmt.Errorf("error saving downloaded file: %v", err)
	}
	outFile.Close()

	// Ensure that the progress line ends with a new line
	fmt.Println()

	fmt.Println("Download completed. Extracting...")

	if err := utils.Decompress(tempFile, versionFolder); err != nil {
		return fmt.Errorf("error extracting: %v", err)
	}

	fmt.Printf("Installation of version %s completed in '%s'.\n", version, versionFolder)
	return nil
}
