package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	mkDirExe()
	createFileExe()

}

func mkDirFn(path string, perm fs.FileMode) {
	err := os.MkdirAll(path, perm)
	if err != nil {
		log.Fatal(err)
	}
}

func mkDirExe() {
	pathList := []string{"./cmd/web", "./internal", "./ui/html/pages", "./ui/html/partials", "./ui/static/css", "./ui/static/img", "./ui/static/js"}

	var path string
	pathMem := &path
	for _, pathStr := range pathList {
		*pathMem = pathStr
		mkDirFn(pathStr, 0755)
	}
}

func createFileFn(name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func createFileExe() {
	mainGoPath := filepath.Join("cmd", "web", "main.go")
	mainCssPath := filepath.Join("ui", "static", "css", "main.css")
	mainJsPath := filepath.Join("ui", "static", "js", "main.js")
	indexHtmlPath := filepath.Join("ui", "html", "pages", "index.html")
	baseTmplPath := filepath.Join("ui", "html", "base.tmpl")

	pathLists := []string{mainGoPath, mainCssPath, mainJsPath, indexHtmlPath, baseTmplPath}

	for _, fileVal := range pathLists {
		fileMem := &fileVal
		*fileMem = fileVal
		file, err := os.Create(fileVal)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}
