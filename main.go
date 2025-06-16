package main

import (
	"fmt"
	"os"
)

func Render() string {
	return `<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="utf-8">
	<title>Go pages</title>
</head>
<body>
	<h1>Hi 👋, I'm Alejandro</h1>
	<p>A passionate Backend Developer from Mexico 🇲🇽</p>
	<p>🔭 I’m currently working on Google Cloud Technologies ...</p>
	<p>🌱 I’m currently learning GitOps and DevSecOps... ...</p>
</body>
</html>`
}

func main() {
	html := Render()

	if len(os.Args) > 1 && os.Args[1] == "--write" {
		_ = os.MkdirAll("dist", 0o755)
		f, err := os.Create("dist/index.html")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, _ = f.WriteString(html)
	} else {
		fmt.Print(html)
	}
}
