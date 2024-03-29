package parser

import (
	"regexp"
	"github.com/scott-x/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]+>([^<]+)</a>`
func CityListParser(contents []byte) engine.ParseResult{
	re :=regexp.MustCompile(cityListRe)
	matches :=re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches{
		result.Items = append(result.Items,"City:"+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:CityParser,
		})
		//fmt.Printf("City: %s, URL: %s \n",m[2],m[1])
	}
	return result
}