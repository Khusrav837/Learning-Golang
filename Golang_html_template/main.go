package main

import (
	"Golang/Golang_html_template/github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} тем:
{{range .Items}}
-------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	resul, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, resul); err != nil {
		log.Fatal(err)
	}
}
