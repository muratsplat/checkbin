package boot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/muratsplat/checkbin/client"
	"github.com/muratsplat/checkbin/helper"
	"github.com/muratsplat/checkbin/register/service"
)

var ()

const (
	NS       = "default"
	NAME     = "auth"
	VERSION  = "v0.0.1"
	ADDRESS  = ":8080"
	PROTOCOL = "websocket"
)

type Auth service.Register

func Run() {
	c := client.NewHttpClient()
	body := NewRegisterRequestJSON()
	bodyReader := bytes.NewReader(body)
	r, err := http.NewRequest("POST", "http://localhost:8080/register", bodyReader)
	res, err := c.Send(r)
	helper.Check(err)
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()
		color.Green("%s:%s was registered just now", NS, NAME)
	} else {
		defer res.Body.Close()
		color.Red("%s:%s could not registered just now", NS, NAME)

		bodyBuff, err := ioutil.ReadAll(res.Body)
		helper.Check(err)
		fmt.Println(string(bodyBuff))
	}

	// Todo: Adding connect to

	select {}

}

func NewRegisterRequestJSON() []byte {
	var auth Auth
	auth.NS = NS
	auth.Name = NAME
	auth.Version = VERSION
	auth.Address = ADDRESS
	auth.Protocol = PROTOCOL

	b, err := json.Marshal(auth)
	helper.Check(err)

	return b
}
