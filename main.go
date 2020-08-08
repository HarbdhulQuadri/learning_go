package main

import (
   "net/http"
   "fmt"
   "html/template"
)

//Create a struct that holds information to be displayed in our HTML file

type Fruits struct {
    Title string
}



type TodoPageData struct {
    PageTitle string
    Fruit     []Fruits
}

func main() {
   //Instantiate a Welcome struct object and pass in some random information. 
   //We shall get the name of the user as a query parameter from the URL
  
		tmpl := template.Must(template.ParseFiles("template/welcome.html"))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data := TodoPageData{
				PageTitle: "My Fruits",
				Fruit: []Fruits{
					{Title: "Apple"},
					{Title: "Avocado"},
					{Title: "Mango"},
					{Title: "Cashew"},
					{Title: "Orange"},
					{Title: "Lemon"},
					{Title: "Plantain"},
					{Title: "Cucumber"},
					{Title: "Apple"},
					{Title: "Grape"},

					
				},
			}
			tmpl.Execute(w, data)
		})

   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":8080", nil));
}