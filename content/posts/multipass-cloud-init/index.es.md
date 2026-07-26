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

Para evitar un entorno local lleno de instalaciones manuales y conflictos de versiones en backend y DevOps, utilizo una solución ligera y limpia: una máquina virtual (VM) de Ubuntu en **Multipass**, configurada desde el primer arranque con **cloud-init**. 

Esto me permite levantar un entorno declarativo, reproducible y desechable en minutos. Si algo falla, destruyo la VM y la recreo sin afectar mi máquina host.

> Repositorio en GitHub: [afr-dt/cloud-init-config](https://github.com/afr-dt/cloud-init-config)

---

## ¿Qué problema quiero resolver?

Trabajar con múltiples herramientas (Terraform, Kubernetes, CLIs de cloud y runtimes) directamente en la máquina host genera problemas clásicos:
- Conflictos entre versiones de lenguajes (Python, Node.js, Go) y paquetes del sistema.
- Dificultad para actualizar CLIs de AWS/GCP de forma consistente.
- Falta de documentación sobre cómo se instalaron las herramientas.
- Poca reproducibilidad al cambiar de computadora o compartir configuraciones.

El objetivo es lograr un entorno **aislado**, definido de forma **declarativa** (control de versiones en Git) y fácil de **destruir y recrear** en segundos.

---

## ¿Por qué Multipass + cloud-init?

Aunque los contenedores son muy útiles, a veces se necesita un entorno de desarrollo completo con un sistema operativo real.

**Multipass** permite crear VMs ligeras de Ubuntu de forma rápida y sencilla. **cloud-init** automatiza la instalación y configuración inicial mediante un archivo declarativo (YAML), reemplazando las largas guías manuales de instalación. En resumen: Multipass proporciona la máquina y cloud-init define su estado.

---

## Arquitectura del enfoque

```text
macOS / Linux host -> Multipass -> Ubuntu VM -> cloud-init -> Herramientas y Runtimes
```

Este flujo mantiene la máquina host limpia y el workspace aislado.

---

## ¿Qué incluye esta instancia?

La VM se aprovisiona automáticamente en su primer inicio con las siguientes herramientas:
- **Lenguajes y Runtimes**: Go, Python (gestión con `uv`), Node.js (gestión con `nvm`) y `build-essential`.
- **Cloud e Infraestructura**: Terraform, AWS CLI v2 y Google Cloud CLI (`gcloud`).
- **Kubernetes**: `kubectl`, `k9s`, `kubectx`, `kubens` y `kubecolor`.
- **Entorno y Utilidades**: Zsh con Starship, Neovim con LazyVim, `jq`, `fzf`, `htop`, `tmux`, `curl` y `git`.

---

## Lanzamiento y gestión de la instancia

Con Multipass instalado, clona el repositorio y ejecuta:

```bash
# Lanzar VM con la configuración declarativa
multipass launch lts --name devops-instance --cpus 2 --memory 4G --disk 28G --cloud-init ./cloud-config.yaml
```
![cloud-init](img/multipass-launch.png)

```bash
# Acceder a la shell de la VM
multipass shell devops-instance
```
![cloud-init](img/multipass-shell.png)

```bash
# Eliminar y purgar la VM al terminar
multipass delete --purge devops-instance
```

---

## Experiencia del primer arranque

Al ingresar a la VM por primera vez, el usuario `ubuntu` ya cuenta con privilegios `sudo` sin contraseña, la shell configurada con Zsh/Starship y el editor Neovim listo. 
Esto demuestra el valor de `cloud-init`: en lugar de ejecutar scripts manuales, todo el estado del sistema se define de forma declarativa, reproducible y fácil de compartir en un solo archivo YAML.

---

## Validación del aprovisionamiento

Una vez dentro de la VM, puedes verificar que las herramientas principales estén disponibles ejecutando:

```bash
terraform version && aws --version && gcloud version && kubectl version --client && go version && node -v && python --version
```
![cloud-init](img/multipass-ubuntu.png)

---

## Casos de uso comunes

Esta VM es ideal para:
- Desarrollar e implementar configuraciones de Terraform de forma aislada.
- Autenticar y ejecutar tareas en AWS y GCP sin persistir credenciales en el host.
- Experimentar con Kubernetes (`kubectl`, `k9s`) de forma segura.
- Probar runtimes de Go, Python o Node.js sin alterar las versiones locales de tu máquina.

---

## Trade-offs y límites

Este enfoque no reemplaza todos los casos de uso locales:
* Para tareas puntuales o de un solo runtime, un contenedor Docker suele ser más rápido y ligero.
* Si requieres interfaces gráficas (GUI) o integraciones directas con el hardware del host, es mejor trabajar de forma nativa.

Sin embargo, para tareas de DevOps e infraestructura, la reproducibilidad y el aislamiento de un entorno real basado en código superan estas limitaciones.
