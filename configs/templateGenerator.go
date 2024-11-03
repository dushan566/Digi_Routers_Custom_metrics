package configs

import (
    "fmt"
		"log"
		"os"
    "text/template"
    "bytes"

)

func TemplateGenerator(a string, b map[string]interface{}, c string){

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles(a))

	var buf = &bytes.Buffer{}
	fmt.Println(buf)
	err := tmpl.Execute(buf, b)
	if err != nil{
		log.Fatalln(err)
	}

	// Debug steps (Print output)
	var s = buf.String()
	fmt.Println(s)

	// Use os.Create to open the output file.
	f, err := os.Create(c)
	if err != nil {
	log.Println("create output file: ", err)
	return
	}
	// A file is an io.Writer. You can execute the template directly to the open file:
	err = tmpl.Execute(f, b)
	if err != nil {
	log.Print("execute: ", err)
	return
	}
	// Close file
	f.Close()
}
