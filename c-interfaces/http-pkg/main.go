package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
	}

	//bs := make([]byte, 99999) /// the second number creates an space of that number of elements
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
