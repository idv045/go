package main 

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Go is...")
}
func contacts_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "is contacts")
}

func handleRequerst(){
	http.HandleFunc("/", home_page)	
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleRequerst()
}