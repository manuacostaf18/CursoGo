package main 

import (
	//"fmt"
	"encoding/json"
	"C:/Users/manua/OneDrive/Documentos/GO/Curso Udemy 1/Sección 4/structs/structs"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOk,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", nil)
}