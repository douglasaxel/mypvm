package main

import (
	"fmt"
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
	case "list-local":
		fmt.Println("Listando versões de PHP instaladas localmente...")
	case "install":
		fmt.Println("Instalando uma versão específica do PHP...")
	case "remove":
		fmt.Println("Removendo uma versão específica do PHP...")
	case "use":
		fmt.Println("Selecionando uma versão específica do PHP...")
	default:
		fmt.Println("Comando inválido!")
		os.Exit(1)
	}
}
