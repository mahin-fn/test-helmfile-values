FROM golang:1.20-alpine3.16 as build
WORKDIR /app 
COPY . . 
RUN go mod vendor
RUN go build -o main . 

FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/main .
ENTRYPOINT [ "/app/main" ]
EXPOSE 8080
CMD ["serve"]
