#stage 1
FROM golang:alpine AS builder

RUN mkdir /builder

ADD . /build/

WORKDIR /build

RUN go build -o main .
#stage 2
FROM alpine

COPY . /app
#Copy from old container "/build/main" to "/app/"
COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 9010

CMD [ "./main" ]

