package main

//O go depende de pacotes
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// define constantes igual no javascript =D
const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
		}
	}

}

//como devolver duas variáveis através de uma função
func devolveNomeIdade() (string, int) {
	nome := "Douglinhas"
	idade := 24

	return nome, idade
}

func exibeIntroducao() {
	nome := "Douglinhas"
	versao := 1.1

	fmt.Println("Olá , sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	//Declaração slice (abstração de array)
	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://www.google.com/",
	// 	"https://mail.google.com/"}

	sites := leSitesDoArquivo()

	// Posso escrever o for da forma tradicional
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	for i := 0; i < monitoramentos; i++ {
		for j, site := range sites {
			fmt.Println("Testando site", j,
				":", site)

			testaSite(site)

		}

		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, ", foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, ", não foi carregado com sucesso. Status:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)

	// fmt.Println("O endereço da variável comandoLido é:", &comandoLido)
	fmt.Println("O comandoLido selecionado foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + " - online:" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.OpenFile("sites.txt", os.O_CREATE, 0666)
	//Quando necessário ler todo o conteudo do arquivo
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		fmt.Println(linha)

		if err == io.EOF {
			fmt.Println("Ocorreu um erro:", err)
			break
		}

	}

	arquivo.Close()

	return sites
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

func exibeNomes() {
	// Definição de um slice. É uma abstração de array.
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println("O meu slice tem", len(nomes), "itens")
	//a funcao cap mostra a capacidade do array
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	//forma de adicionar mais itens no array.
	//Quando adiciona um item no slice, a capacidade do slice dobra.
	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
