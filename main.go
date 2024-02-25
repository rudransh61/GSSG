package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/russross/blackfriday/v2"
)

type Page struct {
	Title   string
	Content template.HTML
}

func main() {
	watchForChanges()
}

func watchForChanges() {
	w := watcher.New()

	go func() {
		for {
			select {
			case event := <-w.Event:
				if event.Op == watcher.Write || event.Op == watcher.Create || event.Op == watcher.Rename {
					generateSite()
				}
			case err := <-w.Error:
				log.Fatal(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive("content"); err != nil {
		log.Fatal(err)
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatal(err)
	}
}

func generateSite() {
	outputDir := "output"

	err := os.RemoveAll(outputDir)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(outputDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk("content", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".md" {
			generatePage(path, outputDir)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func generatePage(mdFilePath, outputDir string) {
	content, err := os.ReadFile(mdFilePath)
	if err != nil {
		log.Fatal(err)
	}

	htmlContent := blackfriday.Run(content)

	title := getTitle(mdFilePath)

	page := Page{
		Title:   title,
		Content: template.HTML(htmlContent),
	}

	outputFilePath := filepath.Join(outputDir, fmt.Sprintf("%s.html", title))
	generateHTML(outputFilePath, page)
}

func generateHTML(outputFilePath string, page Page) {
	templateStr := `<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <article>
        <h1>{{.Title}}</h1>
        {{.Content}}
    </article>
</body>
</html>`

	tmpl, err := template.New("page").Parse(templateStr)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, page)
	if err != nil {
		log.Fatal(err)
	}
}

func getTitle(filePath string) string {
	return filepath.Base(filePath[:len(filePath)-len(filepath.Ext(filePath))])
}
