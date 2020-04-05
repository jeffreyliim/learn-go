package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	results := make(chan string)
	// var values int64 = 0
	// values 2 and 7
	var a = []string{"http://www.mocky.io/v2/5e898dfa310000642bd39d55",
		"http://www.mocky.io/v2/5e89b1283100005600d39ddc",
		"http://www.mocky.io/v2/5e89b18a310000642bd39de0",
		"http://www.mocky.io/v2/5e89b1e93100000f00d39de1",
		"http://www.mocky.io/v2/5e898e583100004367d39d57"}
	for _, v := range a {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			makeReq(v, results)
		}(v)
	}

	go func() {
		defer close(results)
		wg.Wait()
	}()

	for res := range results {
		//i, _ := strconv.ParseInt(res, 10, 32)
		//values += i
		fmt.Println(res)
	}

	//fmt.Println(values)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func makeReq(url string, results chan string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	results <- string(bodyBytes)
	return nil
}
