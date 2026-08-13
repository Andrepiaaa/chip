# chip

Una CLI interactiva escrita en Go para enseñar conceptos de sistemas (desde llamadas al sistema del kernel de Linux hasta arquitectura y diseño de sistemas distribuidos). chip está pensada para uso educativo: talleres, clases, demos y autoaprendizaje.

Este README ofrece instalación, uso, ejemplos, pruebas y guías para contribuir.

---

## Contenidos

- [Características](#características)  
- [Requisitos](#requisitos)  
- [Instalación](#instalación)  
- [Construir desde el código](#construir-desde-el-código)  
- [Ejecutar (modo interactivo)](#ejecutar-modo-interactivo)  
- [Ejemplos de uso](#ejemplos-de-uso)  
- [Testing](#testing)  
- [Formato y estilo de código](#formato-y-estilo-de-código)  
- [Contribuir](#contribuir)  
- [Roadmap / Ideas futuras](#roadmap--ideas-futuras)  
- [Licencia](#licencia)  
- [Contacto](#contacto)

---

## Características

- Interfaz de línea de comandos interactiva orientada a enseñanza.
- Módulos/ejercicios para:
  - Llamadas al sistema y conceptos del kernel.
  - Procesos, señales y concurrencia.
  - Redes y sockets básicos.
  - Arquitectura de sistemas distribuidos (coordinación, replicación, particionado).
  - Simulaciones guiadas y pequeñas prácticas.
- Diseñada para ser extensible: puedes añadir nuevos módulos y ejercicios.

---

## Requisitos

- Git (para clonar el repo)
- Go (recomendado 1.20+). chip está escrito en Go; para compilar o ejecutar necesitas tener Go instalado.
- (Opcional) sudo, si algunos ejercicios requieren permisos de sistema para demostraciones locales. Evita ejecutar comandos con privilegios sin revisar antes.

Comprueba la versión de Go:
```bash
go version
```

---

## Instalación

Clona el repositorio y compila la CLI:

```bash
git clone https://github.com/Andrepiaaa/chip.git
cd chip
```

Instalar usando `go install` (instala el binario en `$(go env GOPATH)/bin` o `$GOBIN` si está configurado):

```bash
go install ./...
```

Alternativa: compilar el binario localmente:

```bash
go build ./...
# o para crear un binario en el directorio actual
go build -o chip ./cmd/chip    # si el main está en cmd/chip
# o simplemente
go build -o chip .
```

Si `go install` coloca el binario en `$GOPATH/bin` o `$GOBIN`, asegúrate de que la carpeta esté en tu PATH:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Construir desde el código

Comandos útiles desde la raíz del repo:

- Compilar todo:
```bash
go build ./...
```

- Ejecutar sin compilar (modo desarrollo):
```bash
go run ./cmd/chip
# o, si el main está en la raíz del módulo:
go run .
```

---

## Ejecutar (modo interactivo)

Inicia la CLI:

```bash
chip
# o si ejecutas desde fuente:
go run ./cmd/chip
```

Al iniciar verás un menú/azúcar sintáctico con opciones interactivas (navegación por teclado, selección de módulos, etc.). chip está diseñada para guiar paso a paso.

Opciones de línea de comandos (ejemplo; ajusta según implementación real):

```text
chip --help
Usage: chip [comando] [flags]

Comandos:
  start     Inicia la experiencia interactiva
  list      Lista módulos disponibles
  run <id>  Ejecuta un módulo/ejercicio por id
  export    Exporta el estado/resultado de la sesión
  version   Muestra la versión del CLI
```

Ejemplo rápido (modo no interactivo, ejecutar un módulo específico):

```bash
chip list
chip run syscall-demo
```

> Nota: Si tu implementación actual usa subcomandos o argumentos distintos, ajusta los comandos anteriores. Si quieres, adapto las secciones de ejemplo al CLI real después de que me confirmes los subcomandos exactos.

---

## Ejemplos de uso

1) Iniciar una lección sobre llamadas al sistema:
```bash
chip start
# seleccionar "Llamadas al sistema" → seguir la guía interactiva
```

2) Ejecutar un módulo de red y ver output simulado:
```bash
chip run net-echo
```

3) Exportar el progreso de una sesión:
```bash
chip export --format=json --out=progreso.json
```

4) Automatizar una demo (script):
```bash
#!/usr/bin/env bash
chip run intro-kernel --non-interactive \
  && chip run syscall-examples --non-interactive
```

---

## Testing

Si el repo contiene tests (paquetes `*_test.go`), ejecutar:

```bash
go test ./...
```

Recomendado: añadir tests unitarios para la lógica no interactiva y tests de integración para los módulos que simulan escenarios.

---

## Formato y estilo de código

Sigue el estilo idiomático de Go:

- formatea con `gofmt`:
```bash
gofmt -w .
```
- revisa con `go vet`:
```bash
go vet ./...
```
- mantiene dependencias con `go mod`:
```bash
go mod tidy
```

Para contribuciones, pide que los commits pasen `gofmt` y `go vet`.

---

## Contribuir

Gracias por querer contribuir. Sugerimos este flujo:

1. Fork del repositorio.
2. Clonar tu fork y crear una rama descriptiva:
```bash
git clone git@github.com:TU_USUARIO/chip.git
cd chip
git checkout -b feat/nombre-de-la-caracteristica
```
3. Implementar la funcionalidad, añadir tests y formatear el código:
```bash
gofmt -w .
go test ./...
```
4. Commit con mensaje claro:
```bash
git add .
git commit -m "feat: añadir módulo de ejemplos de sockets"
git push origin feat/nombre-de-la-caracteristica
```
5. Abre un Pull Request describiendo:
   - Qué resuelve/cambia.
   - Cómo probarlo.
   - Si hay efectos colaterales (permisos, configuración del SO).

Guidelines para PRs:
- Pequeños commits temáticos.
- Tests que cubran la nueva lógica.
- No incluyas binarios o archivos generados en el commit.

Plantilla rápida de PR:
- Título: tipo(scope): descripción corta
- Descripción:
  - Resumen
  - Cambios
  - Cómo probar
  - Checklist (tests, formato)

---

## Roadmap / Ideas futuras

- Módulos interactivos adicionales: persistencia de ejercicios, entornos multiusuario.
- Integración con plataformas educativas (LMS).
- Modo web/GUI para presentar ejercicios desde un navegador.
- Contenedorizado (Docker) para demos reproducibles.
- Internacionalización (más idiomas).
- Ejercicios que no requieran permisos root (sandboxed).

Si tienes ideas o deseas priorizar alguna, abre un issue o discútelo en PR.

---

## Licencia

Incluye un archivo `LICENSE` en el repo. Si no hay licencia actualmente, recomendamos una licencia permisiva como MIT para facilitar uso educativo y colaboraciones.

Ejemplo (MIT):
```text
MIT License
Copyright (c) <Año> <Autor>
...
```

---

## Contacto

Si tienes preguntas, problemas o quieres trabajar en nuevas lecciones, abre un issue o contáctame vía GitHub.
