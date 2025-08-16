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
		fmt.Println("Erro ao obter diretório do usuário.")
		panic(err)
	}
	mypvmFolder := filepath.Join(userHomeDir, ".mypvm")

	fmt.Printf("Definindo a versão %s como padrão...\n", version)

	// 1. Verificar se a versão está instalada localmente.
	versionFolder := filepath.Join(mypvmFolder, version)
	if _, err := os.Stat(versionFolder); os.IsNotExist(err) {
		return fmt.Errorf("a versão %s não está instalada. Use o comando 'install' primeiro", version)
	}

	// 2. Tentar remover o link simbólico existente, se houver.
	if _, err := os.Lstat(versionFolder); err == nil {
		fmt.Println("Removendo link simbólico anterior...")
		if err := os.Remove(versionFolder); err != nil {
			return fmt.Errorf("erro ao remover o link simbólico existente: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("erro ao verificar link simbólico existente: %v", err)
	}

	// 3. Criar o novo link simbólico.
	err = os.Symlink(versionFolder, mypvmFolder)
	if err != nil {
		return fmt.Errorf("erro ao criar o link simbólico para a versão %s: %v", version, err)
	}

	fmt.Printf("Link simbólico criado com sucesso. A versão %s agora é a padrão.\n", version)
	fmt.Println("\nPara usar esta versão no seu terminal, adicione o seguinte caminho ao seu PATH:")

	switch runtime.GOOS {
	case "windows":
		fmt.Printf("  %s\n", mypvmFolder)
		fmt.Println("\nVocê pode fazer isso temporariamente com: 'set PATH=%PATH%;%s'", mypvmFolder)
		fmt.Println("\nOu permanentemente, adicionando a pasta nas 'Variáveis de Ambiente' do Windows.")
	case "linux", "darwin":
		fmt.Printf("  %s\n", mypvmFolder)
		fmt.Println("\nAdicione a linha abaixo no seu arquivo de perfil de shell (ex: ~/.bashrc ou ~/.zshrc):")
		fmt.Println(`  export PATH="` + mypvmFolder + `:$PATH"`)
		fmt.Println("Em seguida, execute 'source ~/.bashrc' (ou o arquivo correspondente) ou reinicie o terminal.")
	}

	return nil
}
