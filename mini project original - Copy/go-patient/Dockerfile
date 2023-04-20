FROM golang:alpine AS builder

RUN mkdir /builder

ADD . /build/

WORKDIR /build

RUN go build -o main .

FROM alpine

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3010

CMD [ "./main" ]

