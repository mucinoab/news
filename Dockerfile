# Cache project dependencies for node
FROM node:current-alpine as cacher_node
WORKDIR app/
COPY ./frontend/package*.json ./frontend/yarn.lock ./
RUN yarn install --immutable && yarn cache clean

# Build the React project
FROM node:current-alpine as builder_node
WORKDIR app/
COPY ./frontend/ .
# Copy over the cached dependencies from cacher_node
COPY --from=cacher_node /app/node_modules node_modules
RUN yarn run build

# Build the Go project
FROM golang:1.22.3-alpine as builder_go
WORKDIR app/
COPY go.mod go.sum main.go .
COPY newsletter newsletter
RUN go mod download && apk add g++ # Required for cgo, and cgo is required for sqlite3
RUN CGO_ENABLED=1 go build -o /app/news

# Build the final minimal image (less than 20 MB)
FROM alpine:3.20 as runtime
RUN mkdir -p frontend/public frontend/dist
COPY --from=builder_node /app/public/ frontend/public/
COPY --from=builder_node /app/dist/ frontend/dist/
COPY --from=builder_go /app/news .
COPY migrations migrations
ENTRYPOINT ["/news"]
