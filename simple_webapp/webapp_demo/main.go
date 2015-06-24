package main

import "fmt"
import "html/template"
import "io/ioutil"
import "log"
import "net/http"
import "strings"

type Details struct {
	Data []string
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Listening on 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		log.Fatal(err)
	}
	details := Details{Data: detectSleRelease()}
	err = t.Execute(w, details)
	if err != nil {
		log.Fatal(err)
	}
}

func detectSleRelease() []string {
	content, err := ioutil.ReadFile("/etc/SuSE-release")
	if err != nil {
		log.Fatal(err)
	}
	lines := []string{}
	for _, l := range strings.Split(string(content), "\n") {
		if !strings.HasPrefix(l, "#") {
			lines = append(lines, l)
		}
	}
	return lines
}
