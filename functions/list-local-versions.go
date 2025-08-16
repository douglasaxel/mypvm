package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListLocalVersions() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Erro ao obter diretório do usuário.")
		panic(err)
	}
	mypvmFolder := filepath.Join(userHomeDir, ".mypvm")

	fmt.Println("Listando versões de PHP instaladas localmente...")

	if _, err := os.Stat(mypvmFolder); os.IsNotExist(err) {
		fmt.Println("Nenhuma versão instalada localmente.")
		return
	}

	directories, err := os.ReadDir(mypvmFolder)
	if err != nil {
		fmt.Println("Erro ao listar versões instaladas localmente.")
		return
	}

	isFoundVersion := false
	for _, dir := range directories {
		if dir.IsDir() {
			isFoundVersion = true
			fmt.Printf(" - Versão instalada: %s\n", dir.Name())
		}
	}

	if !isFoundVersion {
		fmt.Println("Nenhuma versão instalada localmente.")
	}
}
