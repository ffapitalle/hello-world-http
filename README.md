# Aplicación demo - hello-world

## Construcción de la imagen

```
$ docker build -t hello-world-http -f Dockerfile
```

### Construcción en etapas

```
$ docker build -t hello-world-http -f Dockerfile.multistage
```

## Docker Compose

```
$ docker-compose up
```