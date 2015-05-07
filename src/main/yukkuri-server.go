package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"os"
	"path/filepath"
)

func main() {
	useGlobalLogger()
	log.Println(fmt.Sprintf("Start"))

	root := "public"
	if _, err := os.Stat(root); err != nil {
		log.Println(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
	html := getFileListHtml(root)

	m := martini.Classic()
	m.Get("/", func() string {
		html = getFileListHtml(root)
		return html
	})
	m.Run()
}

func useGlobalLogger() {
	log.SetFlags(log.Ldate | log.Ltime)
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[yukkuri-server] ")
}

func getFileListHtml(root string) string {
	log.Println("File list html")
	html := "<html><head></head><body><ul>"
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			rel, err := filepath.Rel(root, path)
			log.Println(fmt.Sprintf("  %v", rel))
			html += "<li><a href=\"./" + rel + "\">" + rel + "</a></li>"
			return nil
		})
	if err != nil {
	}
	html += "</ul></body></html>"
	log.Println("File list html end")
	return html
}
