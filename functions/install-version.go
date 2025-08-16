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

// ProgressReader é um wrapper para io.Reader que mostra o progresso do download
type ProgressReader struct {
	r              io.Reader
	totalSize      int64
	downloadedSize int64
	lastPercent    int
}

// Read implementa a interface io.Reader e atualiza o progresso
func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.r.Read(p)
	pr.downloadedSize += int64(n)

	// Calcular e mostrar o progresso
	var percent int
	if pr.totalSize > 0 {
		percent = int((float64(pr.downloadedSize) / float64(pr.totalSize)) * 100)
	} else {
		// Se não temos o tamanho total, mostrar bytes baixados
		fmt.Printf("\rBaixando... %d bytes", pr.downloadedSize)
		return
	}

	// Atualizar a barra de progresso apenas quando a porcentagem mudar
	if percent != pr.lastPercent {
		pr.lastPercent = percent

		// Criar uma barra de progresso visual
		width := 50
		bar := make([]byte, width)
		for i := 0; i < width; i++ {
			if i < width*percent/100 {
				bar[i] = '='
			} else {
				bar[i] = ' '
			}
		}

		// Mostrar progresso em MB
		downloadedMB := float64(pr.downloadedSize) / 1024 / 1024
		totalMB := float64(pr.totalSize) / 1024 / 1024
		fmt.Printf("\r[%s] %d%% (%.2f MB / %.2f MB)", string(bar), percent, downloadedMB, totalMB)
	}

	return
}

func InstallVersion(version string) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Erro ao obter diretório do usuário.")
		panic(err)
	}
	mypvmFolder := filepath.Join(userHomeDir, ".mypvm")

	fmt.Printf("Iniciando a instalação da versão %s...\n", version)

	versionFolder := filepath.Join(mypvmFolder, version)
	if _, err := os.Stat(versionFolder); !os.IsNotExist(err) {
		fmt.Printf("A versão %s já está instalada.\n", version)
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
		return fmt.Errorf("sistema operacional não suportado: %s", runtime.GOOS)
	}

	fmt.Printf("Baixando de: %s\n", urlDownload)

	// Iniciar o request HTTP
	req, err := http.NewRequest("GET", urlDownload, nil)
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao baixar o arquivo: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao baixar. O servidor retornou: %s", resp.Status)
	}

	// Obter o tamanho total do arquivo
	totalSize := resp.ContentLength

	// Criar o arquivo temporário
	tempFile := version + fileType
	outFile, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo temporário: %v", err)
	}

	defer func() {
		if rErr := os.Remove(tempFile); rErr != nil {
			log.Printf("Aviso: Não foi possível remover o arquivo temporário %s: %v", tempFile, rErr)
		}
	}()

	// Criar um leitor com progresso
	progressReader := &ProgressReader{
		r:              resp.Body,
		totalSize:      totalSize,
		downloadedSize: 0,
		lastPercent:    -1,
	}

	// Copiar do leitor com progresso para o arquivo
	_, err = io.Copy(outFile, progressReader)
	if err != nil {
		outFile.Close()
		return fmt.Errorf("erro ao salvar o arquivo baixado: %v", err)
	}
	outFile.Close()

	// Garantir que a linha de progresso termine com uma nova linha
	fmt.Println()

	fmt.Println("Download concluído. Descompactando...")

	if err := utils.Decompress(tempFile, versionFolder); err != nil {
		return fmt.Errorf("erro ao descompactar: %v", err)
	}

	fmt.Printf("Instalação da versão %s concluída em '%s'.\n", version, versionFolder)
	return nil
}
