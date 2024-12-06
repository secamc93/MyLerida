# AUTH
Microservicio encargado de autenticacion, ususarios y roles 

---

## Pre requisitos ⚙️

Para usar este microservicio, necesitas:

1. Go 1.23 o superior.
2. Makefile instalado en tu sistema.
3. Docker instalado en tu sistema (para crear y ejecutar contenedores Docker).
4. Tener instaladas todas las dependencias del sistema operativo requeridas.

---

## Configuración

La configuración de la aplicación se realiza a través de variables de entorno. Consulta el archivo `.env.example` para ver las variables disponibles y sus usos.



## Generar mocks 🧑‍💻

Ejecuta el siguiente comando:

```shell
make gen-mocks
```

---

## Linter y guía de estilos ✅

Ejecuta el siguiente comando para check-linter:

```shell
make check-linter
```

---

## Docker 🐳

Este proyecto viene con un `Dockerfile` preparado para la construcción de un contenedor Docker. Para construir y ejecutar el contenedor Docker, sigue estos pasos:

1. Construye la imagen Docker con el siguiente comando:

```shell
docker build -t auth .
```

2. Ejecuta el contenedor Docker con el siguiente comando:

```shell
docker run -p 8080:8080 auth
```

## Documentación 📝

