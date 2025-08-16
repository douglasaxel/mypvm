package main

import (
	"fmt"
	"log"
	"mypvm/functions"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Bem-vindo ao Gerenciador de versões do PHP!")
		fmt.Println("Uso: mypvm [comando] [argumentos]")
		fmt.Println("\n\nComandos disponíveis:")
		fmt.Println("\nlist   - Lista as versões de PHP disponíveis online")
		fmt.Println("list-local   - Lista as versões de PHP instaladas localmente")
		fmt.Println("install - Instala uma versão específica do PHP")
		fmt.Println("remove - Remove uma versão específica do PHP")
		fmt.Println("use    - Seleciona uma versão específica do PHP")

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
			fmt.Println("Por favor, informe a versão do PHP a ser instalada. Ex: 'gerenciador-php instalar 8.3.0'")
			os.Exit(1)
		}
		version := os.Args[2]
		if err := functions.InstallVersion(version); err != nil {
			log.Fatalf("Erro na instalação: %v", err)
		}
	case "remove":
		fmt.Println("Removendo uma versão específica do PHP...")
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("Por favor, informe a versão do PHP a ser usada. Ex: 'gerenciador-php use 8.3.0'")
			os.Exit(1)
		}
		version := os.Args[2]
		if err := functions.UseVersion(version); err != nil {
			log.Fatalf("Erro na seleção: %v", err)
		}
	default:
		fmt.Println("Comando inválido!")
		os.Exit(1)
	}
}
