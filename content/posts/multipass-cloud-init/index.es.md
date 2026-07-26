+++
title = "Multipass + cloud-init"
date = 2026-07-26T13:45:58-06:00
description = "Mi VM de Ubuntu reproducible para DevOps en un solo comando"
tags = ["devops", "cloud", "multipass", "ubuntu", "terraform", "kubernetes", "local-development"]
categories = ["DevOps", "Cloud"]
author = "Alejandro Flores"
draft = false
+++

# Multipass + cloud-init

**Mi VM de Ubuntu reproducible para DevOps en un solo comando**

Para evitar que el entorno local se llene de instalaciones manuales y conflictos de versiones en backend y DevOps, uso una combinación bastante limpia y ligera: una máquina virtual (VM) de Ubuntu en **Multipass**, configurada desde el primer arranque con **cloud-init**. 

Esto me permite levantar un entorno declarativo, reproducible y desechable en minutos. Si algo truena o se rompe, simplemente borro la VM y la vuelvo a crear en minutos, sin tocar mi laptop.

> Repositorio en GitHub: [afr-dt/cloud-init-config](https://github.com/afr-dt/cloud-init-config)

---

## ¿Qué problema quiero resolver?

Tener herramientas (Terraform, Kubernetes, CLIs de cloud y runtimes) instaladas directo en la máquina local suele traer varios dolores de cabeza:
- Conflictos entre versiones de lenguajes (Python, Node.js, Go) y paquetes del sistema.
- CLIs de AWS o GCP instaladas a mano que son difíciles de mantener al día de forma consistente.
- Falta de documentación sobre cómo se instalaron las herramientas.
- Cero reproducibilidad al cambiar de laptop o si quieres compartir tu entorno con alguien más.

El objetivo es lograr un entorno **aislado**, definido de forma **declarativa** (control de versiones en Git) y fácil de **destruir y recrear** en segundos.

---

## ¿Por qué Multipass + cloud-init?

Los contenedores resuelven mucho, pero a veces necesitas un entorno de desarrollo completo dentro de un sistema operativo real.

**Multipass** permite crear VMs ligeras de Ubuntu de forma rápida y sencilla. **cloud-init** automatiza la instalación y configuración inicial mediante un archivo declarativo (YAML), ahorrándonos el típico README eterno lleno de pasos manuales. En pocas palabras: Multipass nos da la máquina y cloud-init define su estado.

---

## Arquitectura del enfoque

```text
macOS / Linux host -> Multipass -> Ubuntu VM -> cloud-init -> Herramientas y Runtimes
```

La idea es que tu máquina local se quede limpia y el espacio de trabajo viva completamente aislado en la VM.

---

## ¿Qué incluye esta instancia?

Al arrancar por primera vez, la VM se configura sola con todo esto:
- **Lenguajes y Runtimes**: Go, Python (gestión con `uv`), Node.js (gestión con `nvm`) y `build-essential`.
- **Cloud e Infraestructura**: Terraform, AWS CLI v2 y Google Cloud CLI (`gcloud`).
- **Kubernetes**: `kubectl`, `k9s`, `kubectx`, `kubens` y `kubecolor`.
- **Entorno y Utilidades**: Zsh con Starship, Neovim con LazyVim, `jq`, `fzf`, `htop`, `tmux`, `curl` y `git`.

---

## Cómo levantar y controlar la VM

Con Multipass listo, clona el repositorio y corre estos comandos:

```bash
# Crear la VM usando nuestra configuración
multipass launch lts --name devops-instance --cpus 2 --memory 4G --disk 28G --cloud-init ./cloud-config.yaml
```
![cloud-init](img/multipass-launch.png)

```bash
# Entrar a la terminal de la VM
multipass shell devops-instance
```
![cloud-init](img/multipass-shell.png)

```bash
# Borrar y limpiar la VM por completo cuando termines
multipass delete --purge devops-instance
```

---

## Experiencia del primer arranque

La primera vez que entres a la VM, ya vas a tener al usuario `ubuntu` con permisos de `sudo` sin contraseña, Zsh con Starship listo y el editor Neovim configurado. 
Ahí es donde `cloud-init` brilla: en lugar de seguir guías eternas o correr scripts que a veces fallan, todo el estado de la máquina vive en un solo archivo YAML que se aplica de manera automática en el primer boot.

---

## Verificar que todo esté en orden

Ya dentro de la VM, puedes confirmar que todo se instaló bien corriendo este comando:

```bash
terraform version && aws --version && gcloud version && kubectl version --client && go version && node -v && python --version
```
![cloud-init](img/multipass-ubuntu.png)

---

## ¿Para qué la uso en el día a día?

Esta VM es ideal para:
- Probar configuraciones de Terraform sin llenar de archivos tu máquina.
- Trabajar con AWS y GCP (adiós a dejar llaves de API guardadas en tu host).
- Experimentar con Kubernetes (`kubectl`, `k9s`) de forma segura.
- Probar runtimes de Go, Python o Node.js sin desconfigurar las versiones que ya tienes en tu laptop.

---

## Trade-offs y límites

Claro, esto no resuelve todos los escenarios:
* Si solo necesitas una herramienta muy puntual, un simple contenedor de Docker puede ser más rápido y ligero.
* Si necesitas interfaces gráficas o interactuar directo con el hardware de tu laptop, es mejor trabajar de forma nativa.

Pero para tareas de DevOps e infraestructura, tener un Ubuntu real aislado y definido como código te va a ahorrar muchísimos dolores de cabeza.
