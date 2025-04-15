FROM golang:1.24

# Buat user dan grup
RUN useradd -ms /bin/bash praze

# Buat folder kerja dan beri hak akses ke user `praze`
WORKDIR /app
RUN chown -R praze:praze /app

# Switch ke user praze
USER praze

COPY --chown=praze:praze go.mod go.sum ./
RUN go mod download

COPY --chown=praze:praze . .

RUN go build -o /home/praze/backend .

EXPOSE 8080

CMD ["/home/praze/backend"]
