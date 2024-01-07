package utils

import (
	"encoding/json"
	"io/ioutil" // io puede que sea
	"net/http"

)

//This function will help when creating books, will recieve some body 

func ParseBody(r * http.Request, x interface{}) {

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
	
}