FROM golang:1.23.2 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GO_ARCH=amd64 go build -o weather .

FROM scratch
WORKDIR /app
COPY --from=build /app/weather .
ENTRYPOINT ["./weather"]
