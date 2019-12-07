package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)岁</div>`)
var sexRe = regexp.MustCompile(`"genderString":"([^"]+)"`)
var heightRe = regexp.MustCompile(`"heightString":"([0-9a-z]+)"`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+kg)</div>`)
var incomeRe = regexp.MustCompile(`</div><div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`"marriageString":"([^"]+)"`)
var nativeRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var workRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var childRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>是否想要孩子:([^<]+)</div>`)
var idRe = regexp.MustCompile(`href="http://m.zhenai.com/u/([\d]+)`)
func ParserProfile(content []byte,url string,name string)(result engine.ParserResult){
	profile := model.Profile{}

	profile.Age = extractString(content,ageRe)
	profile.Sex = extractString(content,sexRe)
	profile.Height = extractString(content,heightRe)
	profile.Weight = extractString(content,weightRe)
	profile.Income = extractString(content,incomeRe)
	profile.Marriage = extractString(content,marriageRe)
	profile.NativePlace = extractString(content,nativeRe)
	profile.WorkPlace = extractString(content,workRe)
	profile.Child = extractString(content,childRe)
	profile.Name = name

	result.Items = append(result.Items,engine.Item{
		Id:      extractString(content,idRe),
		Url:     url,
		Payload: profile,
	})

	return result
}

func extractString(content []byte,r *regexp.Regexp)string{
	match := r.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}