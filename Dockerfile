ARG APP_VERSION
FROM golang:1.23 AS build
WORKDIR /src
COPY . .
RUN go build -ldflags "-X github.com/MIcQo/maptonic/config.Version=v${APP_VERSION}" -o maptonic .

FROM gcr.io/distroless/base-debian12
COPY --from=build /src/maptonic /maptonic
ENTRYPOINT ["/maptonic", "serve"]