+++
title = "🤖 Modelos de IA en Local: Introducción con Ollama"
date = 2026-05-23T10:00:00-06:00
description = "Una guía rápida para ejecutar modelos de lenguaje (LLMs) localmente en tu propia máquina usando Ollama."
tags = ["ai", "llm", "ollama", "local"]
categories = ["Machine Learning", "Tools"]
draft = true
+++

Ejecutar modelos de lenguaje en local se ha vuelto increíblemente sencillo gracias a herramientas como **Ollama**. En este post, veremos cómo empezar en pocos minutos.

## 🚀 ¿Por qué usar modelos locales?

- **Privacidad:** Tus datos no salen de tu máquina.
- **Sin Costo:** No pagas por tokens o suscripciones.
- **Sin Internet:** Funciona totalmente offline una vez descargado el modelo.

## ⚙️ Instalación

1. Descarga Ollama desde [ollama.com](https://ollama.com).
2. Instálalo siguiendo las instrucciones para tu sistema operativo (macOS, Linux o Windows).

## 📦 Ejecutando tu primer modelo

Una vez instalado, abre tu terminal y ejecuta:

```bash
ollama run llama3
```

Este comando descargará el modelo (si no lo tienes) y abrirá un chat interactivo.

## 🛠️ Comandos útiles

| Acción | Comando |
| :--- | :--- |
| Listar modelos | `ollama list` |
| Eliminar un modelo | `ollama rm <nombre>` |
| Actualizar Ollama | Solo descarga la nueva versión |

## 🔚 Conclusión

Ollama es la forma más fácil de experimentar con IA sin depender de la nube. ¿Qué modelo vas a probar primero?
