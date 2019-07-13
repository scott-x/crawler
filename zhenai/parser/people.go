package parser

import (
	"regexp"
	"github.com/scott-x/crawler/engine"
)

const peopleRe = `<div class="m-content-box m-des" data-v-bff6f798=""><span data-v-bff6f798="">(.*)</span><!----></div>`
func PeopleParser(contents []byte) engine.ParseResult{
	re :=regexp.MustCompile(peopleRe)
	matches :=re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches{
		result.Items = append(result.Items,"love voice"+string(m[1]))
		//result.Requests = append(result.Requests,engine.Request{
		//	Url:string(m[1]),
		//	ParseFunc:engine.NilParser,
		//})
		//fmt.Printf("City: %s, URL: %s \n",m[2],m[1])
	}
	return result
}