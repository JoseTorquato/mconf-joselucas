package main

import (
	"html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"

	"github.com/urfave/negroni"
   )
// Criação da var templates e apontando para o ponteiro
var templates *template.Template

// Member struct que retorna json 
type Member struct{
	Slug string `json: slug`
	Currency string	`json: currency`
	Image string `json: image`
	Balance uint `json: balance`
	YearlyIncome uint `json: yearlyIncome`
	BackersCount uint `json: backersCount`
	ContributorsCount uint `json: contributorsCount`
}

func api(w http.ResponseWriter, req *http.Request)  {
	response, err := http.Get("https://opencollective.com/seanlarkin.json")

	templates = template.Must(template.ParseGlob("*.html"))	
	
	if err != nil {
		log.Fatal(err)
	} else {
		api, _ := ioutil.ReadAll(response.Body)
		// Imprimi no console o resultado da api trazido
		fmt.Println(string(api))
	
	// criando um user com o type struct para armazenar os dados vindo em fitas de bytes 
	user := Member{}
	if erro := json.Unmarshal([]byte(api), &user); erro != nil {
		log.Fatal(erro)
	}
	
	// aparição no console de fácil leitura
	fmt.Println()
	fmt.Printf("Nome: %s\nMoeda da doação: %s\nImagem: %s\n", user.Slug, user.Currency, user.Image)
	fmt.Println()
	
		
	// executando o template e renderizando no api.html	
	templates.ExecuteTemplate(w, "api.html", user)
	
}
	
}
// func main responsavel pela chamada do programa
func main() {
	// Api via get http como fazer pesquisas com linha de comando ?
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", api)
	
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)
  
	http.ListenAndServe(":3000", n)
}
