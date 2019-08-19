package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() []byte {
	resp, err := http.Get("https://api.github.com/search/repositories?sort=stars&q=atoma")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := get()
		fmt.Println("aa")
		fmt.Printf("%q", b)
		fmt.Println("b")

		fmt.Fprintf(w, string(b))
	})

	http.ListenAndServe(":8080", nil)
}
