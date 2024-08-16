ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION} AS build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build -o /output/ipv4mix .

FROM scratch AS final
WORKDIR /app
COPY --from=build /output /app
ENTRYPOINT [ "/app/ipv4mix" ]