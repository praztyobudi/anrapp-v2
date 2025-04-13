FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/backend .
EXPOSE 8080
CMD ["/bin/backend"]
