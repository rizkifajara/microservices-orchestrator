FROM golang:1.16-alpine
WORKDIR /app
COPY orchestrator .
CMD ["./orchestrator"]