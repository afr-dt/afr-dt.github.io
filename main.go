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
	
	<p>
	👨‍💻 Soy un Ingeniero Backend con experiencia diseñando soluciones modernas y escalables en la nube. Trabajo principalmente con **Python**, **Node.js** y **Go (Golang)**, utilizando arquitecturas basadas en **microservicios**.

	☁️ Me especializo en **Google Cloud Platform**, donde he implementado servicios con **Docker**, **Kubernetes**, **BigQuery**, y **Cloud Storage**.

	🔧 Manejo herramientas de infraestructura como **Terraform** y flujos de trabajo automatizados con **CI/CD**.

	🛠️ Tecnologías que uso con frecuencia:
	- ⚙️ Docker & Kubernetes
	- 📦 Terraform
	- 🚀 CI/CD
	- 🔧 Flask FastAPI / Express.js
	- 🧪 Elasticsearch / BigQuery
	- 🐍 Python, 🟨 Node.js, 🦫 Go

	🎯 Certificado como **Associate Cloud Engineer** por Google Cloud.

	</p> 

	<p>
	👨‍💻 I'm a Backend Engineer with experience designing scalable, cloud-native solutions. I work primarily with **Python**, **Node.js**, and **Go**, building **microservices** architectures.

	☁️ I'm specialized in **Google Cloud Platform**, where I’ve delivered services with **Docker**, **Kubernetes**, **BigQuery**, and **Cloud Storage**.

	🔧 I work with **Terraform** for infrastructure as code and implement automated workflows using **CI/CD**.

	🛠️ Frequently used technologies:
	- ⚙️ Docker & Kubernetes
	- 📦 Terraform
	- 🚀 CI/CD
	- 🔧 Flask FastAPI / Express.js
	- 🧪 Elasticsearch / BigQuery
	- 🐍 Python, 🟨 Node.js, 🦫 Go

	🎯 Certified **Google Cloud Associate Cloud Engineer**
	</p>


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
