FROM golang:1.19.3
COPY ../go.mod ../go.sum /app/
COPY ./benchmarks /app/benchmarks
WORKDIR /app
RUN go mod tidy

ARG FILENAME
ENV REPORT_NAME report_$FILENAME

RUN mkdir /app/out

CMD cd benchmarks && go test -bench=. > /app/out/$REPORT_NAME -benchmem -cpuprofile /app/out/profile_cpu.out