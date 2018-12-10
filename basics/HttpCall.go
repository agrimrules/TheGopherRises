package main

import (
  "fmt"
  "net/http"
)

func contentType(url string)(string, error){
  resp, err := http.Get(url)
  if err != nil {
	 return "", err
  }

  defer resp.Body.Close()

  header := resp.Header.Get("Content-Type")
  return header, nil
}

func main() {
  op, err := contentType("https://google.com")
  fmt.Println(op)
  fmt.Println(err)
}
