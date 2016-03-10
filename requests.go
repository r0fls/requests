package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Text   string
	Status string
}

//TODO
// Post

func Get(url string) Response {
	response, err := http.Get(url)
	if err != nil {
		ret := fmt.Sprintf("%s", err)
		return Response{ret, "Error"}
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			ret := fmt.Sprintf("%s", err)
			return Response{ret, response.Status}
		}
		return Response{string(contents), response.Status}
	}
}
