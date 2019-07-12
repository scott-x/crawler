package  main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
)

func main(){
	resp, err :=http.Get("https://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status Code: %d", resp.StatusCode)
		return
	}
	e :=determinEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all, err :=ioutil.ReadAll(utf8Reader)
	if err!=nil {
		panic(err)
	}
	printCityList(all)


}

func determinEncoding(r io.Reader) encoding.Encoding{
  bytes, err :=bufio.NewReader(r).Peek(1024)
  if err!=nil{
     panic(err)
  }
  e, _, _ :=charset.DetermineEncoding(bytes,"")
  return  e
}

func printCityList(contents []byte){
	re :=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]+>([^<]+)</a>`)
	matches :=re.FindAllSubmatch(contents,-1)
	for _,m := range matches{
		fmt.Printf("City: %s, URL: %s \n",m[2],m[1])
	}
}