FROM golang:1.16-alpine
WORKDIR /app
ENV JWT_SECRET <>
ENV MONGO_URI <>
ENV JWT_SECRET_STAFF <>
ENV PORT <>
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -mod=mod ./cmd/main
CMD ["./main"]