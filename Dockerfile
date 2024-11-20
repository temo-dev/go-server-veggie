FROM golang:1.23.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -o /out/main ./
ENTRYPOINT ["/out/main"]