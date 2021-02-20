package myapp

import(
	"net/http"
	"fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Get UserInfo by /user/{id}")
}


func NewHandler() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/",indexHandler)

	mux.HandleFunc("/users",usersHandler)
	return mux
}
