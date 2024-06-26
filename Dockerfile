FROM golang:1.22 as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./server ./...

FROM scratch
COPY --from=builder /app/server /server
COPY --from=builder /app/scripts/init.sql /scripts/init.sql
ENTRYPOINT ["/server"]
