#go mod tidy

FROM golang:1.17
# Setting environment variables
ENV GO111MODULE="on" \
    CGO_ENABLED="0" \
    GOOS="linux" \
    GOARCH="amd64" \
    GOPROXY="https://goproxy.cn,direct"
# Switch to workspace
WORKDIR /go/src/github.com/paynewu/awesomerevert/
# Load file
COPY . .
EXPOSE 9090
RUN go mod tidy
CMD go run example.go