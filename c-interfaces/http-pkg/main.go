package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct {}

func main(){
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
	}

	//bs := make([]byte, 99999) /// the second number creates an space of that number of elements
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}

	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error){ // this guy is implementing the Write interface since it is following its structure
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}