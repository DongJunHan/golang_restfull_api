package myapp

import(
	"net/http"
	"fmt"
	"time"
	"strconv"
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

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Get UserInfo by /user/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,err)
		return
	}
	user, ok := userMap[id]
	if !ok{
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"No User Id:",id)
		return
	}

		
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w,string(data))
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
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w,string(data))
}
func deleteUserHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_ , ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"No Delete User ID:",id)
		return
	}
	delete(userMap,id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted User ID : ",id)

}
//update를 하는 로직은 대표적으로 두가지정도 있음.
//1. update시도후, 업데이트해야할 정보가 없으면 create함
//2. update시도후, 업데이트해야할 정보가 없으면 error를 반환
func updateUserHandler(w http.ResponseWriter, r *http.Request){
	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,err)
		return
	}

	user, ok := userMap[updateUser.ID]
	if !ok{
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID : ",updateUser.ID)
		return
	}
	//해당 유효성 검사의 단점.
	//클라이언트가 실제 값을 지우고싶어서 빈문자열을 보냈는지
	//기존의 데이터를 유지하고싶어서 해당 key값을 보내지않아 
	//default값으로 빈문자열이 왔는지 알방법이 필요함.
	// 실무에서는 update용 구조체를 따로 만들어서 관리한다고한다.
	// 업데이트용 구조체는 각 key값에 대응되는 flag값을 보관할 수 있는
	// key를 새로 만듦.
	if updateUser.FirstName != ""{
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != ""{
		user.LastName = updateUser.LastName
	}	
	if updateUser.Email != ""{
		user.Email = updateUser.Email
	}
	//if updateUser.FirstName != ""{
	//	user.FirstName = updateUser.FirstName
	//}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))

}

func NewHandler() http.Handler{
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/",indexHandler)
	//gorilla mux 에서 지원하는 함수. Method. 같은 이름의 url일지라도 뒤에 
	//메소드에따라서 구분지어줄 수 있다.
	mux.HandleFunc("/users",usersHandler).Methods("GET")
	mux.HandleFunc("/users",createUserHandler).Methods("POST")	
	mux.HandleFunc("/users",updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}",getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}",deleteUserHandler).Methods("DELETE")

	return mux
}
