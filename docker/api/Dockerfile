ARG GO_VERSION=1.21.0
ARG ALPINE_VERSION=3.18
# ===== build go binary =====
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api backend/cmd/main.go

# ==== build docker image ====
FROM alpine:${ALPINE_VERSION}

WORKDIR /opt/tyranno

COPY --from=go-builder /go/src/api api
# RUN chmod +x "/usr/src/tyranno/api"

ENTRYPOINT ["/opt/tyranno/api"]