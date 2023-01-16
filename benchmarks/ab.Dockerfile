FROM golang:1.19.3
COPY ../main.go ../go.mod ../go.sum /app/
COPY ./benchmarks /app/benchmarks
WORKDIR /app
RUN go mod tidy

CMD go run ./main.go