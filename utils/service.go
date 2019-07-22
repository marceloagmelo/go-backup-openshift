package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarServices lista dos services do openshift
func ListarServices(token string, url string) {
	// Listar os services
	fmt.Printf("Listando todos os services do ambiente %s\n\r", url)
	resultado, services := utils.ListService(token, url)
	if resultado == 0 {

		// Ler os dados dos services
		lerDadosDadosServices(token, services)
	} else {
		fmt.Println("[ListarServices] Services não encontrados")
	}
}

func lerDadosDadosServices(token string, services model.Services) {
	for i := 0; i < len(services.Items); i++ {
		nomeProjeto := services.Items[i].Metadata.Namespace
		nomeService := services.Items[i].Metadata.Name

		lerService(token, nomeProjeto, nomeService)
	}
}

func lerService(token string, nomeProjeto string, nomeService string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirService := dirProjeto + "/service"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirService, 0700)

	resultado, service := utils.GetServiceString(token, url, nomeProjeto, nomeService)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirService + "/" + nomeService + ".json"
		SalvarArquivoJSON(arquivo, service)
	} else {
		fmt.Printf("[lerService] Service %s não encontrado no projeto %s ambiente %s\n\r", nomeService, nomeProjeto, url)
	}
}