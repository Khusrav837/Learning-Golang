package main

import (
	"Golang/Golang_html_template_2/github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `<h1>{{.TotalCount}} тем:</h1>
<table>
<tr style="text-align: left">
<th>#</th><th>State</th><th>User</th><th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='.HTMLURL'>{{.Title}}</a></td>
<tr>
{{end}}
</table>`

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
