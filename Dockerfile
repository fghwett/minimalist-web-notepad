FROM golang:1.15 as builder
ARG GOPROXY=direct
WORKDIR /go/src/fghwett/notepad
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -ldflags="-w -s -extldflags -static" ./main.go

FROM scratch
MAINTAINER fghwett fghwett@qq.com
WORKDIR /app
COPY --from=builder /go/src/fghwett/notepad/static /app/static
COPY --from=builder /go/src/fghwett/notepad/notepad /app/notepad
VOLUME /app/_tmp
EXPOSE 9099
ENTRYPOINT ["/app/notepad"]