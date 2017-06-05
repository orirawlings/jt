// jt (JSON Templating), a simple templating tool that takes json data
// on standard in and renders it through a golang text/template.
//
// Example usage:
//
//    $ curl -s 'https://api.github.com/users/orirawlings/repos' | jt repos.tmpl  # render JSON data with repos.tmpl template
//    Repositories:
//
//    orirawlings/alexa-sites
//    	Fork?: false
//    	Language: Python
//    	Description: Simple script to scrape top site names from alexa.com
//
//    orirawlings/git
//    	Fork?: true
//    	Language: C
//    	Description: Git Source Code Mirror - This is a publish-only repository and all pull requests are ignored. Please follow Documentation/SubmittingPatches procedure for any of your improvements.
//
//    orirawlings/gopl.io
//    	Fork?: true
//    	Language: Go
//    	Description: Example programs from "The Go Programming Language"
//
//    orirawlings/parens
//    	Fork?: false
//    	Language: Python
//    	Description: Parsing paired parentheses
//
//    orirawlings/webcrawler
//    	Fork?: false
//    	Language: Go
//    	Description: A non-trivial Go program which crawls the world wide web to help me learn/teach golang
//
//    $ cat repos.tmpl  # the template file
//    Repositories:
//    {{range .}}
//    {{.full_name}}
//    	Fork?: {{.fork}}
//    	{{with .language -}}
//    	Language: {{.}}
//    	{{end -}}
//    	{{with .description -}}
//    	Description: {{.}}
//    	{{end -}}
//    {{- end}}
package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func must(err error) {
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func main() {
	t := template.Must(template.ParseFiles(os.Args[1:]...))
	var v interface{}
	must(json.NewDecoder(os.Stdin).Decode(&v))
	must(t.Execute(os.Stdout, v))
}
