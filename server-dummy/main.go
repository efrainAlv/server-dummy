package main

import(
	"fmt"
	"net/http"
)


func defaultHanlder(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Software Avanzado - Tarea Práctica 5 - 201712350 - Helmut Najarro"))
}


func main(){

	srv := http.Server{
		Addr:":8080",
	}


	fmt.Println("Server on port 3000:8080")
	http.HandleFunc("/", defaultHanlder)
	srv.ListenAndServe()
}