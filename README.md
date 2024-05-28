# News

Go + Gin + TypeScript + React

Try it out at: https://snews.fly.dev/

## How to run
Using docker, first build the image:

```bash
docker build . -t "news:v1.0"
```

and then run the container:

```bash
docker run --rm -p 8080:8080 -e GIN_MODE=release -e USERNAME=<Your SMTP email> -e PASSWORD=<Your SMTP password> news:v1.0 
```

Note, your email and password should be a google app password, you can create one by following these steps: https://myaccount.google.com/apppasswords

The server will be available at `http://localhost:8080`
