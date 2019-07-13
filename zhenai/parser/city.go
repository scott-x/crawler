package parser

import (
	"regexp"
	"github.com/scott-x/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>`
func CityParser(contents []byte) engine.ParseResult{
	re :=regexp.MustCompile(cityRe)
	matches :=re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches{
		result.Items = append(result.Items,"People:"+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:PeopleParser,
		})
		//fmt.Printf("City: %s, URL: %s \n",m[2],m[1])
	}
	return result
}