FROM golang:1.17 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 LD_FLAGS="-s -w" go build -a -o backend ./main.go

FROM scratch

COPY --from=builder /app/backend /backend

EXPOSE 8080

ENTRYPOINT ["/backend"]
