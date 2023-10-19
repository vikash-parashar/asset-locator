package render

import (
	"html/template"
	"io"
	"log"
)

func RenderTemplate(w io.Writer, tmpName string, data any) {
	tpl, err := template.ParseFiles("./templates/" + tmpName + ".html")
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
