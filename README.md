# Sitio Estático en Go desplegado con GitHub Pages

Este repositorio contiene una página web estática generada con Go y desplegada automáticamente usando **GitHub Actions** y **GitHub Pages**.

## 🚀 ¿Qué hace este proyecto?

- Usa un programa en Go para generar un archivo HTML en la carpeta `dist/`.
- Compila y despliega automáticamente el contenido de `dist/` con GitHub Actions.
- Publica el sitio en: [https://<username>.github.io](https://<username>.github.io)

## 🛠️ Tecnologías

- [Go](https://golang.org/)
- [GitHub Actions](https://docs.github.com/en/actions)
- [GitHub Pages](https://pages.github.com/)

## 📂 Estructura

```
.
├── main.go         # Generador de HTML
├── dist/           # Carpeta generada automáticamente con index.html
└── .github/
└── workflows/
└── gh-pages.yml  # Workflow de despliegue
```