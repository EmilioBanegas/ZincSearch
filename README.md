# Proyecto sobre búsqueda en base de correos de ZincSearch

Este es un proyecto que utiliza a ZincSearch como DB para la búsqueda de correos.

## Table of contents

- [Overview](#overview)
  - [The challenge](#the-challenge)
  - [Links](#links)
  - [Install](#install)
- [My process](#my-process)
  - [What I learned](#what-i-learned)

## Overview


### The challenge

- Indexar una base de datos en ZincSearch y luego crear un backend y frontend para consultarla.


### Links

- Solution URL: https://github.com/EmilioBanegas/ZincSearch.git

- Node.js https://nodejs.org/es/

- Vue 3: https://vuejs.org/

- Pinia: https://pinia.vuejs.org/

- Router Vue: https://router.vuejs.org/

- Quasar: https://quasar.dev/

- Go: https://go.dev/

- Chi: https://go-chi.io/#/

- Python 3: https://www.python.org/

- ZincSearch: https://zincsearch.com/



### Install
  El proceso para la instalación es el siguiente:
  1) Descargar e instalar las tecnologías usadas.

  2) Configurar la contraseña y usuario de ZincSearch

  3) Establecer los mismos datos del paso 2 en los archivos:
    /db/indexDB.py líneas 11 y 12.
    /backend/.env líneas 1 y 2.

  4) Descargar la base de datos ernor: http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz
  
  5) Agregar la carpeta ernor_mail_20110402 al directorio /db/

  6) Abrir el programa zinc.exe para levantar la base de datos.

  7) Ejecutar el archivo de python que se encuentra en /db/indexDB.py
    con el comando: python indexDB.py (el comando se ejecuta desde la carpeta /db/)
    para indexar la base de datos ernor a ZincSearch.

  8) Ejecutar el archivo de Go con el 
    comando: go run . (el comando se ejecuta desde la carpeta /backend/)

    En dado caso no tenga la librerías utilizadas, entonces utilice 
    el comando: go mod tidy
    dentro de las carpetas:
      - /backend/
      - /backend/helpers
      - /backend/routes



## My process
  
  Desarrollé el proyecto con los siguientes pasos:

  - Indexé la base de datos erno.
  - Desarrollé el backend.
  - Desarrollé el frontend.


### What I learned

- Aprendí a trabajar con Go, Chi, ZincSearch y poner en práctica los conocimientos que había adquirido 
  en el entorno de Vue 3.