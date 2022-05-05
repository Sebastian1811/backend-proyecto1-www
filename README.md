# Backend-proyecto1-www
Backend para proyecto 1 www. App backend desarrollado en go antes de configurar ambiente dev instalar [golang](https://go.dev/).

## Index

1. [Config ambiente dev](#configurar-ambiente-de-desarrollo)
2. [Arrancar server auto-reload](#arrancar-servidor-con-auto-reload)
3. [Arrancar sin auto-reload](#arrancar-el-servidor-sin-auto-reload)

## Configurar ambiente de desarrollo

1. Clonar este repositorio:

```
    git clone https://github.com/Sebastian1811/backend-proyecto1-www
```

2. Instalar dependencias proyecto:

```
    go get ./...
```

3. Instalar nodemon usando el gestor de paquetes npm:

```
    npm install nodemon
```

4. Si se usa windows se debe instalar el gestor de paquetes [Chocolatey](https://chocolatey.org/install) luego ejecutar el comando:

```
    choco install make
```

## Arrancar servidor con auto-reload

1. Ejecute el siguiente comando:

```
    make dev
```

## Arrancar el servidor sin auto-reload

1. Ejecute el siguiente comando:

```
    go run server.go
```

