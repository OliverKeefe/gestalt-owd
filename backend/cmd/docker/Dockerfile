FROM golang:1.25.3-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/owd ./cmd/owd

FROM gcr.io/distroless/static-debian12:nonroot AS runtime

COPY --from=build /out/owd /usr/local/bin/owd

EXPOSE 8081

ENTRYPOINT ["/usr/local/bin/owd"]