package myapp

import(
	"testing"
	"net/http/httptest"
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
	assert.Contains(string(data),"Get UserInfo")
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
