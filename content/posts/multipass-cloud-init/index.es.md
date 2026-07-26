+++
title = "Multipass + cloud-init"
date = 2026-07-26T13:45:58-06:00
description = "Mi VM de Ubuntu reproducible para DevOps en un solo comando"
tags = ["devops", "cloud", "multipass", "ubuntu", "terraform", "kubernetes", "local-development"]
categories = ["DevOps", "Cloud"]
author = "Alejandro Flores"
draft = true
+++

# Multipass + cloud-init

**Mi VM de Ubuntu reproducible para DevOps en un solo comando**

En backend, cloud y DevOps es muy fácil que el entorno local termine lleno de instalaciones manuales, versiones mezcladas y configuraciones difíciles de repetir. Después de probar distintas formas de aislar herramientas sin ensuciar mi host, terminé usando una combinación que me ha funcionado muy bien: una VM ligera de Ubuntu en Multipass, provisionada desde el primer arranque con `cloud-init`.

La idea es simple: levantar un entorno limpio, reproducible y desechable en minutos. Si algo se rompe, la VM se elimina y se vuelve a crear; si cambio de equipo, el setup se puede reconstruir con el mismo archivo declarativo.

> Repositorio en GitHub: [afr-dt/cloud-init-config](https://github.com/afr-dt/cloud-init-config)

---

## ¿Qué problema quiero resolver?

Si trabajas con Terraform, Kubernetes, CLIs de cloud y varios runtimes al mismo tiempo, tarde o temprano el sistema host se vuelve incómodo.

Algunos síntomas clásicos:

- Versiones distintas de Python, Node.js o Go mezcladas con las del sistema.
- CLIs de AWS o GCP instaladas a mano y difíciles de actualizar de forma consistente.
- Herramientas que funcionaban en una laptop, pero nadie recuerda cómo volver a instalarlas exactamente igual.
- Muy poca reproducibilidad al cambiar de máquina o compartir entorno con otra persona.

Lo que quería era tener tres cosas muy claras:

- Un entorno aislado del host.
- Una configuración declarativa y versionada en Git.
- La libertad de destruir y recrear la VM cuando quiera, sin miedo a romper nada importante.

---

## ¿Por qué Multipass + cloud-init?

Los contenedores resuelven muchos casos, pero no siempre son la mejor experiencia cuando lo que buscas es una caja de herramientas completa con shell, editor, runtimes, CLIs e instalaciones de sistema dentro de un Ubuntu real.

Ahí es donde Multipass encaja muy bien: permite lanzar VMs ligeras de Ubuntu de forma rápida, con buen equilibrio entre aislamiento y simplicidad. Y `cloud-init` completa la historia porque convierte todo el aprovisionamiento inicial en una definición declarativa, en lugar de una lista larga de pasos manuales en un README.

En otras palabras, Multipass me da la máquina; `cloud-init`, el estado deseado.

---

## Arquitectura rápida del enfoque

La idea general se puede resumir así:

```text
macOS / Linux host
        |
        v
    Multipass
        |
        v
   Ubuntu VM
        |
        v
   cloud-init
        |
        v
Runtimes + Cloud CLIs + Kubernetes tools + Shell setup
```

Ese aislamiento es justo lo que me interesa: el host se mantiene limpio y la VM se convierte en un workspace reproducible para tareas de DevOps y backend.

---

## ¿Qué incluye esta instancia DevOps?

La VM arranca con Ubuntu y durante el primer boot instala una base bastante práctica para trabajo diario.

### Lenguajes y runtimes

- Go
- Python gestionado con `uv`
- Node.js LTS gestionado con `nvm`
- `build-essential` para dependencias nativas

### Infraestructura y cloud

- Terraform
- AWS CLI v2
- Google Cloud CLI (`gcloud`)

### Herramientas de Kubernetes

- `kubectl`
- `k9s`
- `kubectx`
- `kubens`
- `kubecolor`

### Utilidades del entorno

- `jq`, `fzf`, `htop`, `tmux`, `curl`, `git`, `unzip`
- Zsh con Starship
- Neovim con LazyVim
- PATH personalizado para tener las herramientas disponibles desde el inicio

El resultado es que, al entrar a la VM, ya tienes un entorno funcional en lugar de una instalación vacía de Ubuntu que todavía requiere media hora de configuración manual.

---

## Cómo lanzar la instancia

Una vez que tienes Multipass instalado, el flujo es muy simple. Clonas el repositorio y lanzas la instancia con:

```bash
multipass launch lts \
  --name devops-instance \
  --cpus 2 \
  --memory 4G \
  --disk 28G \
  --cloud-init ./cloud-config.yaml
```

![cloud-init](img/multipass-launch.png)

Defino CPU, memoria y disco explícitamente para que la VM siga siendo útil sin consumir más recursos de los necesarios en el host.

Luego puedes entrar con:

```bash
multipass shell devops-instance
```

![cloud-init](img/multipass-shell.png)

Y cuando ya no la necesites, eliminarla es igual de directo:

```bash
multipass delete --purge devops-instance
```

---

## Experiencia del primer arranque

La intención es que el primer login ya se sienta como un entorno de trabajo, no como un servidor recién instalado.

Por defecto, la configuración deja listo lo siguiente:

- Usuario `ubuntu`
- `sudo` sin contraseña para ese usuario
- Zsh + Starship para tener un prompt más informativo
- Neovim con LazyVim
- Mejor ergonomía para Kubernetes con `k9s`, `kubectx`, `kubens` y `kubecolor`

Aquí es donde `cloud-init` realmente aporta valor: en vez de documentar pasos manuales o scripts difíciles de mantener, todo el estado deseado vive en un archivo YAML que se aplica automáticamente en el primer boot.

Esto hace que el entorno sea:

- Declarativo
- Reproducible
- Compartible
- Fácil de reconstruir

---

## Validación rápida después del provisioning

Una mejora útil para este flujo es verificar desde el inicio que las herramientas principales realmente quedaron disponibles.

Después de entrar a la VM, puedes validar con algo como esto:

```bash
terraform version
aws --version
gcloud version
kubectl version --client
go version
node -v
python --version
```

![cloud-init](img/multipass-ubuntu.png)

Si estos comandos responden correctamente, ya tienes una señal bastante confiable de que la instancia quedó lista para tareas comunes de infraestructura, cloud y desarrollo.

---

## ¿Para qué lo estoy usando?

Hay varios escenarios donde esta VM me resulta especialmente útil:

- Probar Terraform sin ensuciar mi host.
- Trabajar con AWS y GCP desde un entorno aislado.
- Experimentar con herramientas de Kubernetes sin tocar mi configuración local principal.
- Tener una caja de desarrollo con Go, Python y Node.js lista para usar.

Como todo vive dentro de la VM, puedo instalar cosas, romper configuraciones, probar ideas y volver a un estado limpio en pocos minutos simplemente recreando la instancia.

---

## Trade-offs y límites

Este enfoque no reemplaza todos los flujos locales.

Si solo necesitas ejecutar una herramienta muy puntual, un contenedor puede ser más ligero. Y si dependes mucho de integraciones directas con el sistema host o herramientas gráficas específicas, seguir trabajando fuera de la VM puede ser más práctico.

Aun así, para mi caso el beneficio principal sigue siendo la consistencia: un Ubuntu real, aislado, definido como código y fácil de reconstruir.

---
