FROM golang:1.21.4-alpine AS builder
ARG GITHUB_TOKEN GITHUB_USER
RUN apk update && apk add --no-cache git 

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOPRIVATE=github.com/turistikrota/service.shared

WORKDIR /

RUN echo "machine github.com login $GITHUB_USER password $GITHUB_TOKEN" > ~/.netrc

COPY go.* ./
RUN   --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . . 
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o main ./cmd/main.go

FROM scratch

ENV PORT 8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /main .
COPY --from=builder /locales ./locales
COPY --from=builder /assets ./assets
EXPOSE $PORT

CMD ["/main"]