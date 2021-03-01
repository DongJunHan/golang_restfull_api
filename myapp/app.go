package myapp

import(
	"net/http"
	"fmt"
	//"strings"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Get UserInfo by /user/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("path : %s\n",r.URL.Path)
	//slice := strings.Split(r.URL.Path,"/")
	vars := mux.Vars(r)
	fmt.Fprint(w,"User Id:"+vars["id"])
}



func NewHandler() http.Handler{
	mux := mux.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/",indexHandler)

	mux.HandleFunc("/users",usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}",getUserInfoHandler)
	return mux
}
