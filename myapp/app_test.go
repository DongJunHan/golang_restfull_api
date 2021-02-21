package myapp

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"

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
	assert.Equal(string(data),"User Id:89")
}
