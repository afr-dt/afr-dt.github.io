+++
title = "🚀 UV, el nuevo estándar para Python"
date = 2025-11-04T09:00:00-06:00
description = "UV reemplaza pip, virtualenv y pipx con una sola herramienta ultrarrápida escrita en Rust. Aprende cómo dominarlo."
tags = ["python", "devops", "virtualenv", "pip", "uv", "rust"]
categories = ["Python", "Tools"]
draft = true
+++

## 🧩 ¿Qué es UV?

[`uv`](https://astral.sh/uv) es una herramienta moderna creada por **Astral** (los creadores de `ruff` 🦀) diseñada para ser el único administrador que necesites en el ecosistema Python. Escrito en **Rust**, es entre 10 y 100 veces más rápido que las herramientas tradicionales.

### ⚡ ¿Por qué es tan rápido?
- **Global Cache:** No vuelve a descargar lo que ya tienes.
- **Hardlinks:** Usa enlaces físicos para no duplicar archivos en tu disco.
- **Resolución Paralela:** Resuelve dependencias de forma masiva y eficiente.

## ⚙️ Instalación

```bash
# macOS & Linux
curl -LsSf https://astral.sh/uv/install.sh | sh

# Windows
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

---

## 🐍 1. Gestión de Versiones de Python

Olvídate de `pyenv` o de pelear con las versiones de tu sistema. `uv` descarga y gestiona versiones de Python por ti.

```bash
# Instalar una versión específica
uv python install 3.12

# Listar versiones instaladas y disponibles
uv python list

# Usar una versión específica para un proyecto
uv venv --python 3.11
```

---

## 📦 2. Flujo de Trabajo Moderno (Proyectos)

El modo recomendado es usar `uv` como gestor de proyectos (estilo `cargo` o `npm`).

```bash
# Inicializar un proyecto
uv init my-project
cd my-project

# Añadir dependencias (actualiza pyproject.toml y uv.lock)
uv add fastapi uvicorn

# Sincronizar el entorno (instala todo lo que falte)
uv sync

# Ver el árbol de dependencias
uv tree
```

---

## 🛠️ 3. Gestión de Herramientas (Pipx killer)

Puedes instalar y ejecutar herramientas de CLI de forma aislada sin ensuciar tu entorno global.

```bash
# Ejecutar una herramienta sin instalarla (on-the-fly)
uvx ruff check .

# Instalar una herramienta globalmente
uv tool install yt-dlp
uv tool install black

# Listar herramientas instaladas
uv tool list
```

---

## 🧪 4. Migración y Compatibilidad

Si vienes de un flujo basado en `requirements.txt`:

```bash
# Instalar desde un archivo existente
uv pip install -r requirements.txt

# Exportar tu proyecto a requirements.txt (con hashes)
uv export --format requirements-txt > requirements.txt
```

---

## 🧹 5. Comandos Útiles de Mantenimiento

| Acción | Comando |
| :--- | :--- |
| **Actualizar UV** | `uv self update` |
| **Limpiar Caché** | `uv cache clean` |
| **Generar Lockfile** | `uv lock` |
| **Ejecutar Script** | `uv run myscript.py` |
| **Listar paquetes** | `uv pip list` |

## 🔚 Conclusión

`uv` no es solo una mejora incremental; es un cambio de paradigma. Centraliza la gestión de versiones, entornos y herramientas en una sola interfaz rápida y moderna. Si aún no lo usas, hoy es el mejor día para migrar.
