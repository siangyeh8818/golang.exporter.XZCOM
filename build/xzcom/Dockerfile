FROM golang:1.12.7-stretch as builder

COPY go.mod /go/src/github.com/siangyeh8818/golang.exporter.alibabacloud/go.mod
COPY go.sum /go/src/github.com/siangyeh8818/golang.exporter.alibabacloud/go.sum

# Run golang at any directory, not neccessary $GOROOT, $GOPATH
ENV GO111MODULE=on
WORKDIR /go/src/github.com/siangyeh8818/golang.exporter.alibabacloud

# RUN go mod init github.com/pnetwork/sre.monitor.metrics
RUN go mod download
COPY cmd /go/src/github.com/siangyeh8818/golang.exporter.alibabacloud/cmd
COPY internal /go/src/github.com/siangyeh8818/golang.exporter.alibabacloud/internal
#COPY pkg /go/src/github.com/pnetwork/sre.monitor.metrics/pkg

# Build the Go app
RUN env GOOS=linux GOARCH=amd64 go build -o alibabacloud-exporter -v -ldflags "-s" github.com/siangyeh8818/golang.exporter.alibabacloud/cmd/

##### To reduce the final image size, start a new stage with alpine from scratch #####
FROM alpine:3.9
RUN apk --no-cache add ca-certificates libc6-compat

# Run as root
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/siangyeh8818/golang.exporter.alibabacloud/alibabacloud-exporter /usr/local/bin/alibabacloud-exporter

# EXPOSE 8081

ENTRYPOINT [ "alibabacloud-exporter" ] 
