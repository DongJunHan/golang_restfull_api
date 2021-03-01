package myapp

import(
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
)

type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

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

func createUserHandler(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,err)
		return
	}
	//Created User
	user.ID = 2
	user.CreatedAt = time.Now()
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w,string(data))
}


func NewHandler() http.Handler{
	mux := mux.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/",indexHandler)

	mux.HandleFunc("/users",usersHandler).Methods("GET")
	mux.HandleFunc("/users",createUserHandler).Methods("POST")	
	mux.HandleFunc("/users/{id:[0-9]+}",getUserInfoHandler)
	return mux
}
