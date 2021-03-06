package myapp

import(
	"testing"
	"net/http/httptest"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World",string(data))
}

func TestUsers(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL+"/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(string(data),"No Users")
	//assert.Equal("Hello World",string(data))
}
func TestGetUserInfo(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL+"/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data),"No User Id:89")
}
func TestCreateUserInfo(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users","application/json",strings.NewReader(`{"first_name":"Han","last_name":"dong","email":"gamedokdok@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	
	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)

	assert.NotEqual(0,user.ID)
	id := user.ID
	res, err = http.Get(ts.URL+"/users/"+strconv.Itoa(id))	
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(res.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUserInfo(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	req, _ := http.NewRequest("DELETE",ts.URL+"/users/1",nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)
	
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data),"No Delete User ID:1")
	//insert	
	res, err = http.Post(ts.URL+"/users","application/json",strings.NewReader(`{"first_name":"Han","last_name":"dong","email":"gamedokdok@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	
	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)

	assert.NotEqual(0,user.ID)

	//delete
	req, _ = http.NewRequest("DELETE",ts.URL+"/users/1",nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)
	
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data),"Deleted User ID : 1")
	

}

func TestUpdateUser(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	req, _ := http.NewRequest("PUT",ts.URL+"/users",strings.NewReader(`{"id" : 1, "first_name":"updated", "last_name":"updated", "email" : "updated@naver.com"}`))

	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	//해당 데이터가 없을 경우.
	assert.Contains(string(data),"No User ID : 1")
	//post로 create를 요청함.
	res, err = http.Post(ts.URL+"/users","application/json",strings.NewReader(`{"first_name":"Han","last_name":"dong","email":"gamedokdok@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	
	//get data
	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0,user.ID)
	//update data
	//update시 업데이트하고싶은 데이터만 보내게 되면, 나머지 데이터는
	//default값으로 바꾸게 된다. string은 빈문자열로 보내게되고
	//기존의 정보를 빈문자열로 덮어쓰게 된다.
	//따라서 기존의 데이터는 그대로 가지고 가야한다.
	updateStr := fmt.Sprintf(`{"id" : %d, "first_name":"updated"}`,user.ID)
	req, _ = http.NewRequest("PUT",ts.URL+"/users",strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)
	//get data
	updateUser := new(User)
	err = json.NewDecoder(res.Body).Decode(updateUser)
	assert.NoError(err)

	assert.Equal(updateUser.ID,user.ID)
	assert.Equal("updated",updateUser.FirstName)
	assert.Equal(user.LastName,updateUser.LastName)
	assert.Equal(user.Email,updateUser.Email)
	
}

func TestUsers_WithUsersData(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	//insert first user
	res, err := http.Post(ts.URL+"/users","application/json",strings.NewReader(`{"first_name":"Han","last_name":"dong","email":"gamedokdok@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	
	//insert second user
	res, err = http.Post(ts.URL+"/users","application/json",strings.NewReader(`{"first_name":"json","last_name":"jjsn","email":"json@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	
	res, err = http.Get(ts.URL+"/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)

	userList := []*User{}
	err = json.NewDecoder(res.Body).Decode(&userList)
	assert.NoError(err)
	assert.Equal(2,len(userList))
}

