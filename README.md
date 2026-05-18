# Sitio estático con Hugo desplegado en GitHub Pages

Este repositorio contiene un sitio estático construido con **Hugo** (tema [hugo-coder](https://github.com/luizdepra/hugo-coder)) y desplegado automáticamente mediante **GitHub Actions** a **GitHub Pages**.

## ¿Qué hace este proyecto?

- Genera el sitio con Hugo (Extended) y compila SCSS del tema.
- Despliega automáticamente a GitHub Pages al hacer push en `main`.

## Requisitos

- Hugo Extended ≥ 0.146.0
- Dart Sass disponible en PATH (p. ej. `npm i -g sass@1.62.1`)
- Hugo Modules habilitado (el tema se resuelve como módulo)

## Estructura

```
.
├── hugo.toml            # Configuración del sitio
├── content/             # Contenido Markdown
│   ├── about.md
│   └── posts/           # Artículos del blog
├── archetypes/          # Plantillas para nuevo contenido
├── (tema vía Hugo Modules, no se necesita `themes/` en el repo)
└── .github/workflows/gh-pages.yml  # Workflow de despliegue
```

## Desarrollo local

```
hugo mod tidy
hugo server -D
```

Visita http://localhost:1313 para previsualizar. Usa `-D` para incluir borradores.

## Build manual

```
hugo --gc --minify
```

La salida se genera en `public/` (no se debe commitear).

## Deploy

El deploy se ejecuta vía workflow en `.github/workflows/gh-pages.yml` al hacer push a `main`. Asegura en Settings → Pages que la fuente sea "GitHub Actions".

## Crear contenido

```
hugo new posts/mi-articulo.md
```

Usa front matter y nombres en minúsculas con guiones.

Notas: mantén `baseURL` correcto en `hugo.toml`; no subas `public/` ni secretos. Si falla Sass, verifica que `sass` está disponible en el PATH.

## Dependencias de módulos

Commitea `go.mod` y `go.sum` para builds reproducibles. Usa `hugo mod tidy` cuando agregues o actualices el tema o módulos.

El tema está fijado al release v1.2 de hugo-coder (`b476d777`).
