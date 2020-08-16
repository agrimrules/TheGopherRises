package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/panjf2000/ants"
)

func main() {
	defer ants.Release()
	var wg sync.WaitGroup
	initial, _ := makeRequest("https://xkcd.com/info.0.json")
	var total = 0
	if initial["num"] != nil {
		total = int(initial["num"].(float64))
	}

	p, _ := ants.NewPoolWithFunc(runtime.NumCPU()*2, func(i interface{}) {
		result, err := makeRequest(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))
		if err != nil {
			log.Fatal(err)
		}
		if result["img"] != nil {
			imgURL := result["img"].(string)
			downloadFile(imgURL, path.Base(imgURL))
		}
		wg.Done()
	})

	for i := 0; i <= total; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
}

func makeRequest(URL string) (map[string]interface{}, error) {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer response.Body.Close()
	var result map[string]interface{}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &result)
	return result, nil
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
	}
	defer response.Body.Close()

	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
