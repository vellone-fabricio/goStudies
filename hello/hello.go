package main

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

const TIMER_DELAY = 5 * time.Second
const TEST_REPETITION = 3

func main() {
	for {
		command := mainMenu()
		readSiteTxt()
		switch command {
		case 1:
			initMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Saindo do programa!")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}

}

func mainMenu() int {
	fmt.Println("1 - Exibir monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var command int
	fmt.Scan(&command)

	return command
}

func initMonitoring() {
	fmt.Println("Exibindo monitoramento ...")
	sites := readSiteTxt()

	for i := 0; i < TEST_REPETITION; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(TIMER_DELAY)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		sitesLogger(site, true)
	} else {
		fmt.Println("Site:", site, "esta fora com com problemas. Status Code:", response.StatusCode)
		sitesLogger(site, false)
	}
}

func readSiteTxt() []string {
	var sites []string

	arq, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	reader := bufio.NewReader(arq)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}
	arq.Close()
	return sites
}

func sitesLogger(site string, status bool) {
	arq, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("erro", err)
	}
	arq.WriteString(time.Now().Format("02/01/2006 15/04/05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")
	arq.Close()
}

func printLogs() {
	fmt.Println("Exibindo logs...")

	arq, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("erro:", err)
	}

	fmt.Println(string(arq))
}
