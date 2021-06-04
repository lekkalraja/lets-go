package main

import (
	"log"
	"net/http"
)

func main() {
	sites := []string{
		"https://golang.org/",
		"http://openjdk.java.net/",
		"https://www.scala-lang.org/",
		"https://nodejs.org/en/",
		"https://www.python.org/",
		"https://aws.amazon.com/",
	}

	for _, site := range sites {
		checkStatus(site)
	}
}

func checkStatus(site string) {
	resp, err := http.Get(site)

	if resp.StatusCode != 200 || err != nil {
		log.Printf("%s is DOWN!!!! (%d)", site, resp.StatusCode)
		return
	}

	log.Printf("%s is UP!!!! (%d)", site, resp.StatusCode)
}
