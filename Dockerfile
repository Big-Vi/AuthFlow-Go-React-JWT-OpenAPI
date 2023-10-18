## Dev
FROM golang:1.21.3-alpine

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install live reloader
RUN go install github.com/cosmtrek/air@latest

# Copy app files
COPY . .

EXPOSE 50052

CMD ["air", "serve"]