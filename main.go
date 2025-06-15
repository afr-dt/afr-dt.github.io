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
	<h1>¡Alejandro Flores!</h1>
	<p>GitHub Actions</p>
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
