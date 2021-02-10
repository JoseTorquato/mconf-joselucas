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

var templates *template.Template

 
type member struct{
	Slug string `json: slug`
	Currency string	`json: currency`
	Imagem string `json: image`
	Balance uint `json: balance`
	YearlyIncome uint `json: yearlyIncome`
	BackersCount uint `json: backersCount`
	ContributorsCount uint `json: contributorsCount`
}



func main() {
	response, err := http.Get("https://opencollective.com/seanlarkin.json")

	templates = template.Must(template.ParseGlob("*.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err != nil {
			log.Fatal(err)
		} else {
			api, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(api))
		
			user := member{}
		
			if erro := json.Unmarshal([]byte(api), &user); erro != nil {
				log.Fatal(erro)
			}
			fmt.Println()
			fmt.Printf("Nome: %s\nMoeda da doação: %s\n", user.Slug, user.Currency)
			fmt.Println()
			
		
			
			templates.ExecuteTemplate(w, "api.html", user)
	}
	})
	
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)
  
	http.ListenAndServe(":3000", n)
}

func api()  {
	
	
}