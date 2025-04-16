FROM golang:1.24

# Buat user dan grup
RUN useradd -ms /bin/bash praze

# Buat folder kerja dan beri hak akses ke user `praze`
WORKDIR /app
RUN chown -R praze:praze /app

RUN mkdir -p tmp && chmod -R 777 tmp

# Switch ke user praze
USER praze

COPY --chown=praze:praze go.mod go.sum ./
RUN go mod download

COPY --chown=praze:praze . .

RUN go build -o /home/praze/backend .

EXPOSE 8080

CMD ["/home/praze/backend"]

#========================================================
# FROM golang:1.24

# # Buat user non-root
# RUN useradd -ms /bin/bash praze

# # Working directory
# WORKDIR /app
# RUN chown -R praze:praze /app

# # Ganti user
# USER praze

# # Copy go.mod dan go.sum dulu
# COPY --chown=praze:praze go.mod go.sum ./
# RUN go mod download

# # Copy seluruh source code
# COPY --chown=praze:praze . .

# # Build binary ke path eksplisit
# RUN go build -o /home/praze/backend .

# # Verifikasi file binary setelah build
# RUN ls -lah /app

# # Pastikan file executable
# RUN chmod +x /home/praze/backend

# EXPOSE 8080

# # Jalankan aplikasi
# CMD ["/home/praze/backend"]

