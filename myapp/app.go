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

func getUserInfo89Handler(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprint(w,"User Id:89")
}



func NewHandler() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/",indexHandler)

	mux.HandleFunc("/users",usersHandler)
	mux.HandleFunc("/users/89",getUserInfo89Handler)
	return mux
}
