FROM golang:1.19.3
COPY ./benchmarks /app
WORKDIR /app
ARG FILENAME
ENV REPORT_NAME report_$FILENAME

CMD go test -bench ... > $REPORT_NAME