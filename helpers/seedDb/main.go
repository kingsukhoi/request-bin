package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {

	client := &http.Client{}

	methods := []string{"GET", "POST", "PUT", "DELETE"}

	for i := range 100 {
		body := map[string]interface{}{
			"req_number": i,
		}

		sendMe, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		randMethod := rand.IntN(len(methods))

		req, err := http.NewRequest(methods[randMethod], fmt.Sprintf("http://127.0.0.1:8080/bin?req_num=%d", i), bytes.NewBuffer(sendMe))
		if err != nil {
			panic(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			panic("bad status code")
		}
		_ = resp.Body.Close()
	}

}
