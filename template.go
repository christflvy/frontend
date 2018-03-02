package main

import (
	"html/template"

	"bytes"
	"strconv"
	"strings"
)

var renderEngine *template.Template

func initRender() {
	funcMap := template.FuncMap{}
	funcMap["Currency"] = CurrencyFormat
	funcMap["CallTemplate"] = CallTemplate

	templates := "files/templates/*.html"
	renderEngine = template.Must(template.New("").Funcs(funcMap).ParseGlob(templates))
}

func CurrencyFormat(n int) string {
	nStr := strconv.Itoa(n)
	na := strings.Split(nStr, ".")
	r := []rune(na[0])
	b := []rune{}
	for i, j := len(r)-1, 3; i >= 0; i-- {
		if j == 0 && r[i] != '-' {
			b = append([]rune{'.'}, b...)
			j = 3
		}
		j--
		b = append([]rune{r[i]}, b...)
	}
	return string(b)
}

func CallTemplate(name string, data interface{}) (template.HTML, error) {
	buf := bytes.NewBuffer([]byte{})
	err := renderEngine.ExecuteTemplate(buf, name, data)
	ret := template.HTML(buf.String())

	return ret, err
}
