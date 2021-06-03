package main

import (
	"io"
	"log"
	"net/http"
)

type LogWriter struct{}

func main() {
	resp, err := http.Get("http://ip.jsontest.com/")

	if err != nil {
		log.Printf("something went wrong while calling : %v \n", err)
	}

	log.Printf("Got Response : %v \n", resp)

	lw := LogWriter{}

	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	log.Printf("Consumed Data : %v", string(bs))
	return len(bs), nil
}

/*
	RUN:
	===
	raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/interfaces/custom_writer.go
	2021/06/03 10:19:35 Got Response : &{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Allow-Origin:[*] Cache-Control:[private] Content-Type:[application/json] Date:[Thu, 03 Jun 2021 02:19:35 GMT] Server:[Google Frontend] X-Cloud-Trace-Context:[4b651fe15c8c522144a15980d656081c]] 0xc0000ca1c0 -1 [] false true map[] 0xc0000fe000 <nil>}
	2021/06/03 10:19:35 Consumed Data : {"ip": "2404:e801:2000:b:60f4:b576:7d61:bbf8"}
	raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$
*/
