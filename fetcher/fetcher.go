package fetcher

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)


func Fetch(url string) ([]byte, error){
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err!=nil {
		return nil,err
	}

	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	resp, err :=http.DefaultClient.Do(request)
	if  err!=nil{
		return nil,err
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status Code: %d", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}
	e :=determinEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all, err :=ioutil.ReadAll(utf8Reader)
	if err!=nil {
		panic(err)
	}
	return  all,nil
}

func determinEncoding(r io.Reader) encoding.Encoding{
	bytes, err :=bufio.NewReader(r).Peek(1024)
	if err!=nil{
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e, _, _ :=charset.DetermineEncoding(bytes,"")
	return  e
}