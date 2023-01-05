package main

import (
	"io"
	"net/http"
)

func fun(res http.ResponseWriter, req *http.Request) {

	client := &http.Client{}

	request, _ := http.NewRequest(req.Method, req.RequestURI, req.Body)

	for k := range req.Header {

		if k != "Accept-Encoding" {

			request.Header.Add(k, req.Header.Get(k))

		}

	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	ans, _ := io.ReadAll(response.Body)

	res.Write(ans)

}

func main() {

	http.HandleFunc("/", fun)
	http.ListenAndServe(":8081", nil)

}
