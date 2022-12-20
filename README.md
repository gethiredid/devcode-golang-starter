# Devcode Starter using Golang and Mux Level 3

## Hasil Akhir yang Diharapkan

Peserta dapat membuat dan menampilkan data kontak yang terkoneksi dengan database

## Setup Environment

1. Download source code melalui link yang telah disediakan dari halaman assesment
2. Extract source code yang sudah terdownload pada perangkat anda
3. Buka source code yang sudah diextract menggunakan Code Editor, contoh Visual Studio Code
4. Salin isi dari file `.env.example` ke dalam file `.env`
5. Lakukan migrasi database dengan mengikuti langkah-langkahnya yang bisa dilihat [disini](#migrasi-database)
6. install air `go install github.com/cosmtrek/air@latest` untuk menjalankan golang dengan mode development
7. Jalankan `go mod download` pada terminal untuk install packages
8. Jalankan `go run main.go` atau `air` untuk mode development pada terminal

## Instruksi Pengerjaan

1. Pastikan anda sudah meng-install tools yang diperlukan. Jika belum, silahkan ikuti langkah-langkahnya [disini](#menginstal-tools-yang-digunakan)
2. Jalankan API dan Database dengan Docker, silahkan ikuti langkah-langkahnya [disini](#menjalankan-api-dan-database-dengan-docker)
3. Sesuaikan request dan response pada route GET `/contacts` pada file `main.go` sesuai dengan [Dokumentasi API](https://documenter.getpostman.com/view/6584319/2s8Yt1rUtN) pada Postman
4. Sesuaikan request dan response pada route POST `/contacts` pada file `main.go` sesuai dengan [Dokumentasi API](https://documenter.getpostman.com/view/6584319/2s8Yt1rUtN) pada Postman
5. Lakukan unit testing pada local anda dengan menggunakan Docker, langkah-langkahnya dapat dilihat [disini](#menjalankan-unit-testing-dengan-Docker)
6. Push projek ke docker hub setelah semua test case berhasil dijalankan, langkah-langkahnya dapat dilihat [disini](#push-projek-ke-docker-hub)
7. Submit image docker yang telah dipush ke Docker Hub ke Halaman Submission Devcode, langkah-langkahnya dapat dilihat [disini](#push-projek-ke-docker-hub)

## Tools dan Packages yang Digunakan

1. [Git](https://git-scm.com)
2. [GoLang](https://go.dev)
3. [Docker](https://www.docker.com)
4. [Mux](https://pkg.go.dev/github.com/gorilla/mux)
5. [Air](https://github.com/cosmtrek/air)
6. [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv#section-readme)
7. [GORM](https://gorm.io)

## Menginstal Tools yang Digunakan

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- Docker
  - [Windows](https://docs.docker.com/desktop/install/windows-install/)
  - [Mac](https://docs.docker.com/desktop/install/mac-install/)
  - [Linux](https://docs.docker.com/desktop/install/linux-install/)

## Menjalankan API dan Database dengan Docker

Jika anda sudah menginstall docker, anda bisa menjalankan perintah `docker-compose up -d` untuk menjalankan API <b>Contact Manager</b> dan juga database <b>Mysql</b>. Tetapi pastikan `environment` pada file .env yang telah kamu buat dari .env.example sesuai dengan `environment` pada file `docker-compose.yaml`.

Apabila ada perubahan pada file kodingan anda, anda bisa build ulang container dengan perintah :

```
docker-compose up -d --build --force-recreate
```

## Migrasi Database

Projek ini menggunakan package `gorm` untuk mengurus database.
Migrasi table berada di file `main.go`.

Anda dapat membuat table migrasi dengan menggunakan `struct`, contohnya adal pada bagian `type Contact struct {...}` di file `main.go`.

Migrasi dijalankan dengan AutoMigrate pada gorm, contoh `db.AutoMigrate(&Contact{})`.

Untuk lebih lengkap anda bisa mengunjungi [link ini](https://gorm.io/docs/models.html) .

## Menjalankan Unit Testing dengan Docker

Pastikan environment database dan port API pada file `.env` sama dengan `file docker-compose.yaml`.
Dan pastikan anda telah menjalakan database dan api pada docker lokal, kalau belum jalankan perintah berikut `docker-compose up -d` atau `docker-compose up -d --build --force-recreate` untuk build ulang image ketika ada perubahan pada file.

Jalankan perintah berikut untuk melakukan unit testing:

```
docker run --network="host" -e API_URL=http://localhost:3030 -e LEVEL=3 alfi08/hello-unit-testing
```

## Submit ke Devcode

### Build Docker Image

Jalankan perintah berikut untuk Build docker image `docker build . -t {name}`

contoh :

```
docker build . -t golang-hello
```

### Push projek ke Docker Hub

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
