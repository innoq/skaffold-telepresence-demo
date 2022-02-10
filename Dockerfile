FROM golang:buster as builder
LABEL maintainer=dimitrij.drus@innoq.com

RUN apt-get update && apt-get install -y xz-utils

# UPX is GPL
ADD https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | \
    tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

ENV USER=appuser
ENV UID=10001

RUN useradd --system --user-group login-provider
RUN adduser \
    --disabled-login \
    --gecos "" \
    --home "/nonexistent" \
    --no-create-home \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"

WORKDIR /go/src/hello-go

COPY go.mod .
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"
RUN strip --strip-unneeded hello-go
RUN upx hello-go

# The actual image of the app
FROM scratch
LABEL maintainer=dimitrij.drus@innoq.com

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/src/hello-go/hello-go /opt/hello-go/hello-go

WORKDIR /opt/hello-go

USER appuser:appuser

ENV GIN_MODE=release
ENTRYPOINT ["/opt/hello-go/hello-go"]
