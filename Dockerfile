FROM golang:alpine3.19 AS build

RUN mkdir /var/app
WORKDIR /var/app
RUN apk add --no-cache git
COPY . /var/app/
RUN go mod download
RUN go build -o /api ./cmd/api/main.go

# production image
FROM alpine:3.19.1
RUN mkdir /app
WORKDIR /app
COPY --from=build /api /app/api
EXPOSE 8080
CMD [ "/app/api" ]
