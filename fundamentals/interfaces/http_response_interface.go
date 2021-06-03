package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://ip.jsontest.com/")

	if err != nil {
		log.Printf("something went wrong while calling : %v \n", err)
	}

	log.Printf("Got Response : %v \n", resp)

	/*bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Printf("The Response Body : %s \n", string(bs))*/

	io.Copy(os.Stdout, resp.Body)
}
