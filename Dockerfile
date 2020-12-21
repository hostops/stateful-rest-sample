FROM golang:alpine AS build

WORKDIR /src/
COPY go.* /src/
RUN go mod download

COPY main.go /src/
ENV CGO_ENABLED=0
RUN go build -o /bin/stateful-rest-sample
RUN chmod +x /bin/stateful-rest-sample

FROM scratch
COPY --from=build /bin/stateful-rest-sample /bin/stateful-rest-sample

ENTRYPOINT ["/bin/stateful-rest-sample"]
