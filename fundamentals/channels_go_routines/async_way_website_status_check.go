package main

import (
	"log"
	"net/http"
	"time"
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

	var c chan string = make(chan string)

	for _, site := range sites {
		go checkStatus(&site, site, c)
	}

	/*for index := range sites {
		fmt.Printf("[%d] Response of Sites : %s \n", index, <-c)
	}*/

	// infinite loop
	/*for {
		go checkStatus(<-c, c)
	}*/

	/* for output := range c {
		fmt.Printf("Response of Sites : %s \n", output)
	} */

	/* for site := range c {
		//time.Sleep(2 * time.Second) // main - go routine sleeps for 2 seconds
		go checkStatus(site, c)
	} */

	for site := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			go checkStatus(&link, link, c) // WARNING : loop variable site captured by func literal
		}(site)
	}
}

func checkStatus(mainSite *string, site string, c chan string) {
	resp, err := http.Get(site)
	// time.Sleep(2 * time.Second) // child go routine will sleep for 2 seconds before sending message to channel
	if resp.StatusCode != 200 || err != nil {
		log.Printf("[main address : %p] [child address : %p] %s is DOWN!!!! (%d)", mainSite, &site, site, resp.StatusCode)
		//c <- fmt.Sprintf("%s is DOWN!!!! (%d)", site, resp.StatusCode)
		c <- site
		return
	}

	log.Printf("[main address : %p] [child address : %p] %s is UP!!!! (%d)", mainSite, &site, site, resp.StatusCode)
	//c <- fmt.Sprintf("%s is UP!!!! (%d)", site, resp.StatusCode)
	c <- site
}
