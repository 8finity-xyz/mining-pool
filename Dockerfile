FROM golang:1.24 AS build-stage

# build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o /pool

# run
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY --from=build-stage /pool /pool
USER nonroot:nonroot

ENTRYPOINT ["/pool"]