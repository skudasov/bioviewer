package server

import (
	"testing"
	"net/http"
	"fmt"
	"time"
	"os"
	"io/ioutil"

	"github.com/lulunevermind/bioviewer/back/util"
)

var host = "localhost"
var port = 8080
var idp = "esia"
var userid = 240631324

func TestMain(m *testing.M) {
	r := GetRouter()
	go r.Run()
	<-time.After(1 * time.Second)
	os.Exit(m.Run())
}

func TestGetProfile(t *testing.T) {
	url := fmt.Sprintf("http://%s:%d/profile/%s/%d", host, port, idp, userid)
	resp, err := http.Get(url)
	util.CheckErr(err)
	respBody, err := ioutil.ReadAll(resp.Body)
	util.CheckErr(err)
	resp.Body.Close()
	fmt.Printf("Response is %s", respBody)
}