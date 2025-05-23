# ===== Step 1: Build React frontend =====
FROM node:18 AS frontend-builder

WORKDIR /app/client
COPY client/package*.json ./
RUN npm install
COPY client/ .
RUN npm run build


# ===== Step 2: Build Go backend =====
FROM golang:1.23 AS backend-builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
COPY --from=frontend-builder /app/client/dist ./static

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app . && chmod +x app






# ===== Step 3: Create minimal final image =====
FROM alpine:latest

WORKDIR /root/
COPY --from=backend-builder /app/app .
COPY --from=backend-builder /app/static ./static
COPY .env .env

EXPOSE 8080

CMD ["./app"]
