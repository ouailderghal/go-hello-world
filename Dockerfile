FROM golang:1.21-alpine as build
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM golang:1.21-alpine as app

COPY --from=build /app/main /main

CMD [ "/main" ]

