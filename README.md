# News

Go + TypeScript + React

## How to run
Using docker, first build the image:

```bash
docker build . -t "news:v1.0"
```

and then run the container:

```bash
docker run --rm -p 8080:8080 -e GIN_MODE=release news:v1.0
```

The server will be available at `http://localhost:8080`