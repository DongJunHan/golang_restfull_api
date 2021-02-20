package main

import(
	"net/http"
	"WEB-INF/golang_restfull_api/myapp"
)

func main(){
	http.ListenAndServe(":3000",myapp.NewHandler())
}
