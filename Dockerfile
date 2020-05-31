FROM golang:alpine AS builder
WORKDIR /build

# fetch dependencies first
COPY go.mod .
COPY go.sum .
RUN go mod download

# build the code
COPY . .
RUN go build -o main cmd/server/main.go

FROM golang:alpine
ARG PORT=8080
ENV PORT_E=${PORT}

COPY --from=builder /build/cmd/server/index.html .
COPY --from=builder /build/schema.graphql .
COPY --from=builder /build/main .

# run the service
EXPOSE ${PORT_E}
CMD ./main \
    -host "0.0.0.0" \
    -port ${PORT_E} \
    -html index.html \
    -allowedOrigins "*" \
    -allowedHeaders "*"

