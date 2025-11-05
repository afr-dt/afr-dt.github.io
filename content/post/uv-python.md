+++
title = "🚀 UV, el nuevo estándar para Python"
date = 2025-11-04T09:00:00-06:00
description = "UV reemplaza pip, virtualenv y pipx con una sola herramienta ultrarrápida. Aprende cómo usarlo y por qué migrar hoy."
tags = ["python", "devops", "virtualenv", "pip", "uv"]
categories = ["Python", "Tools"]
draft = false
+++

## 🧩 ¿Qué es UV?

[`uv`](https://astral.sh/uv) es una herramienta moderna creada por **Astral** (los de `ruff` 🦀) para reemplazar **pip**, **virtualenv**, **pip-tools** y **pipx** con un solo binario rápido y eficiente.

## ⚙️ Instalación

```bash
# En macOS & Linux.
curl -LsSf https://astral.sh/uv/install.sh | sh
# o con Homebrew
brew install uv

# En Windows.
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

## 📦 Crea tu entorno en segundos

Antes:

```bash
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

Ahora:

```bash
uv venv
uv pip install -r requirements.txt
```

Modo moderno (recomendado):

```bash
uv init myproject
uv add fastapi uvicorn
uv sync
```

uv crea .venv, pyproject.toml y uv.lock.

## ⚡ Ejecuta sin activar el entorno

```bash
uv run app.py
uv run pytest
```

## 🧪 Migrando desde requirements.txt

Compatibilidad:

```bash
uv pip install -r requirements.txt
```

Migración:

```bash
uv init
uv add -r requirements.txt
uv sync
```

Esto genera pyproject.toml y uv.lock

## 🧹 Limpieza y mantenimiento

| Acción                            | Comando                                        |
| --------------------------------- | ---------------------------------------------- |
| Limpiar caché                     | `uv cache clean`                               |
| Ver deps                          | `uv pip list`                                  |
| Actualizar desde requirements.txt | `uv pip install --upgrade -r requirements.txt` |

## 🔚 Conclusión

uv es simple, rápido, compatible y moderno.
Si pensabas que Python necesitaba un mejor gestor de entornos, este es el momento.
