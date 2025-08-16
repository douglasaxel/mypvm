package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"mypvm/utils"
	"net/http"
	"sort"
)

type Release struct {
	Version string `json:"version"`
	Date    string `json:"date"`
}

// List all major versions of PHP available online, then all minor versions
func ListOnlineVersions() {
	fmt.Println("Buscando versões de PHP disponíveis online...")

	majorVersionsUrl := "https://www.php.net/releases/index.php?json"
	res, err := http.Get(majorVersionsUrl)

	if err != nil {
		log.Fatalf("Erro ao buscar versões principais do PHP disponíveis online: %v", err)
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Erro ao buscar versões principais do PHP disponíveis online: %v", res.Status)
		panic(res.Status)
	}

	var majorReleases map[string]Release

	if err := json.NewDecoder(res.Body).Decode(&majorReleases); err != nil {
		log.Fatalf("Erro ao decodificar JSON das versões principais: %v", err)
		panic(err)
	}

	versions := []string{}
	for _, release := range majorReleases {
		versions = append(versions, release.Version)
	}

	allVersions := make(map[string]Release)

	for _, version := range versions {
		detailUrl := fmt.Sprintf("https://www.php.net/releases/index.php?json=1&version=%s&max=100", version[:1])

		resDetail, err := http.Get(detailUrl)
		if err != nil {
			log.Fatalf("Aviso: Erro ao buscar a versão %s: %v", version, err)
		}

		defer resDetail.Body.Close()

		if resDetail.StatusCode != http.StatusOK {
			log.Fatalf("Aviso: Falha na requisição para a versão %s. Status %s", version, resDetail.Status)
		}

		var releasesDetail map[string]Release
		if err := json.NewDecoder(resDetail.Body).Decode(&releasesDetail); err != nil {
			log.Fatalf("Aviso: Erro ao decodificar JSON da versão %s: %v", version, err)
		}

		maps.Copy(allVersions, releasesDetail)
	}

	sorted := make([]Release, 0, len(allVersions))
	for k, release := range allVersions {
		release.Version = k
		sorted = append(sorted, release)
	}

	sort.Slice(sorted, func(i, j int) bool {
		versionAFloat := utils.VersionToFloat(sorted[i].Version)
		versionBFloat := utils.VersionToFloat(sorted[j].Version)

		return versionAFloat > versionBFloat
	})

	fmt.Println("\nVersões do PHP disponíveis online:")
	for _, release := range sorted {
		fmt.Printf("- Versão %s (Lançamento: %s)\n", release.Version, release.Date)
	}
}
