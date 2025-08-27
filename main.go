package main

import (
	"fmt"
	"os"
)

func Render() string {
	return `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="utf-8">
		<title>Go Pages</title>
		<style>
			body {
				font-family: sans-serif;
				line-height: 1.6;
				margin: 20px;
				transition: background-color 0.3s, color 0.3s;
			}
			body.dark-mode {
				background-color: #121212;
				color: #ffffff;
			}
			.theme-switcher {
				position: fixed;
				top: 20px;
				right: 20px;
				cursor: pointer;
				border: 1px solid #ccc;
				padding: 10px;
				border-radius: 5px;
			}
		</style>
	</head>
	<body>
		<div class="theme-switcher" onclick="toggleTheme()">🌙</div>

		<h1>Hi 👋, I'm Alejandro</h1>
		
		<section>
			<h2>Perfil (Español)</h2>
			<p>👨‍💻 Soy un Ingeniero Backend con experiencia diseñando soluciones modernas y escalables en la nube. Trabajo principalmente con Python, Node.js y Go (Golang), utilizando arquitecturas basadas en microservicios.</p>
			<p>☁️ Me especializo en Google Cloud Platform, donde he implementado servicios con Docker, Kubernetes, BigQuery y Cloud Storage.</p>
			<p>🔧 Manejo herramientas de infraestructura como Terraform y flujos de trabajo automatizados con CI/CD.</p>

			<h3>🛠️ Tecnologías que uso con frecuencia:</h3>
			<ul>
				<li>⚙️ Docker & Kubernetes</li>
				<li>📦 Terraform</li>
				<li>🚀 CI/CD</li>
				<li>🔧 Flask, FastAPI, Express.js</li>
				<li>🧪 Elasticsearch, BigQuery</li>
				<li>🐍 Python, 🟨 Node.js, 🦫 Go</li>
			</ul>

			<p>🎯 Certificado como Associate Cloud Engineer por Google Cloud.</p>
		</section>

		<hr>

		<section>
			<h2>Profile (English)</h2>
			<p>👨‍💻 I'm a Backend Engineer with experience designing scalable, cloud-native solutions. I work primarily with Python, Node.js, and Go, building microservices architectures.</p>
			<p>☁️ I'm specialized in Google Cloud Platform, where I’ve delivered services with Docker, Kubernetes, BigQuery, and Cloud Storage.</p>
			<p>🔧 I work with Terraform for infrastructure as code and implement automated workflows using CI/CD.</p>

			<h3>🛠️ Frequently used technologies:</h3>
			<ul>
				<li>⚙️ Docker & Kubernetes</li>
				<li>📦 Terraform</li>
				<li>🚀 CI/CD</li>
				<li>🔧 Flask, FastAPI, Express.js</li>
				<li>🧪 Elasticsearch, BigQuery</li>
				<li>🐍 Python, 🟨 Node.js, 🦫 Go</li>
			</ul>

			<p>🎯 Certified Google Cloud Associate Cloud Engineer.</p>
		</section>

		<script>
			function toggleTheme() {
				document.body.classList.toggle('dark-mode');
				const themeSwitcher = document.querySelector('.theme-switcher');
				if (document.body.classList.contains('dark-mode')) {
					themeSwitcher.textContent = '☀️';
				} else {
					themeSwitcher.textContent = '🌙';
				}
			}
		</script>
	</body>
	</html>
	`
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
