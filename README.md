# Devcode golang starter with mux - Level 5


## Tools yang di perlukan

- Git
- Docker 

### Cara menginstall Tools

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Docker : 
    - [Windows](https://docs.docker.com/desktop/install/windows-install/)
    - [Mac](https://docs.docker.com/desktop/install/mac-install/)
    - [Linux](https://docs.docker.com/desktop/install/linux-install/)

## Package yang digunakan

- mux
- air
- godotenv
- gorm


## Docker

Jika anda sudah menginstall docker, anda bisa menjalankan perintah `docker-compose up -d` untuk menjalankan API <b>Contact Manager</b> dan juga database <b>Mysql</b>. Tetapi pastikan `environment` pada file .env yang telah kamu buat dari .env.example sesuai dengan `environment` pada file `docker-compose.yaml`.

Apabila ada perubahan pada file kodingan anda, anda bisa build ulang container dengan perintah :
```
docker-compose up -d --build --force-recreate
```

## Menjalankan projek

- copy `.env.example` to `.env`
- install air `go install github.com/cosmtrek/air@latest` untuk menjalankan golang dengan mode development
- install package `go mod download`
- jalankan projek dengan perintah `go run main.go` atau `air` untuk mode development


## Migration

Projek ini menggunakan package `gorm` untuk mengurus database.
Migrasi table berada di file `main.go`.

Anda dapat membuat table migrasi dengan menggunakan `struct`, contohnya ada pada bagian `type Contact struct {...}` di file `main.go`.

Migrasi dijalankan dengan AutoMigrate pada gorm, contoh `db.AutoMigrate(&Contact{})`.

Untuk lebih lengkap anda bisa mengunjungi [link ini](https://gorm.io/docs/models.html) .


# Menjalankan unit testing dengan docker

Pastikan environment database dan port API pada file `.env` sama dengan `file docker-compose.yaml`.
Dan pastikan anda telah menjalakan database dan api pada docker lokal, kalau belum jalankan perintah berikut  `docker-compose up -d` atau `docker-compose up -d --build --force-recreate` untuk build ulang image ketika ada perubahan pada file.

Jalankan perintah berikut untuk melakukan unit testing:
```
docker run --network="host" -e API_URL=http://localhost:3030 -e LEVEL=5 alfi08/hello-unit-testing
```

# Submit ke Devcode
## Build docker image
Jalankan perintah berikut untuk Build docker image  `docker build . -t {name}`

contoh :
```
docker build . -t golang-hello
```


## Push projek ke docker hub

Pastikan sudah memiliki akun docker hub, dan login akun docker anda di lokal dengan perintah `docker login`.

Setelah itu jalankan perintah berikut untuk push docker image lokal ke docker hub.

```
docker tag golang-hello {username docker}/golang-hello
docker push {username docker}/golang-hello
```

Setelah itu submit docker image ke Devcode.

```
{username docker}/golang-hello
```