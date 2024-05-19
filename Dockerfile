FROM golang:1-bookworm as build
WORKDIR /go/src/
COPY go.mod go.sum ./
RUN go mod download
COPY . . ./
ARG BUILD_TIME=unknown
ARG BUILD_VERSION=unknown
RUN go build -ldflags="-s -w -X main.buildTime=${BUILD_TIME} -X main.buildVersion=${BUILD_VERSION}" -o app ./main.go

FROM gcr.io/distroless/static-debian12:nonroot as runtime
COPY --from=build /go/src/app /bin/app
ENTRYPOINT [ "/bin/app" ]
