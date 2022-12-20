# Devcode Starter using Golang and Mux Level 2

## Hasil Akhir yang Diharapkan

Peserta dapat membuat dan menampilkan data kontak dengan menggunakan local variable

## Setup Environment

1. Download source code melalui link yang telah disediakan dari halaman assesment
2. Extract source code yang sudah terdownload pada perangkat anda
3. Buka source code yang sudah diextract menggunakan Code Editor, contoh Visual Studio Code
4. Salin isi dari file `.env.example` ke dalam file `.env`
5. install air `go install github.com/cosmtrek/air@latest` untuk menjalankan golang dengan mode development
6. Jalankan `go mod download` pada terminal untuk install packages
7. Jalankan `go run main.go` atau `air` untuk mode development pada terminal

## Instruksi Pengerjaan

1. Pastikan anda sudah meng-install tools yang diperlukan. Jika belum, silahkan ikuti langkah-langkahnya [disini](#menginstal-tools-yang-digunakan)
2. Jalankan API dan Database dengan Docker, silahkan ikuti langkah-langkahnya [disini](#menjalankan-api-dan-database-dengan-docker)
3. Sesuaikan request dan response pada route GET `/contacts` pada file `main.go` sesuai dengan [Dokumentasi API](https://documenter.getpostman.com/view/6584319/2s8Yt1rUtN) pada Postman
4. Sesuaikan request dan response pada route POST `/contacts` pada file `main.go` sesuai dengan [Dokumentasi API](https://documenter.getpostman.com/view/6584319/2s8Yt1rUtN) pada Postman
5. Lakukan unit testing pada local anda dengan menggunakan Docker, langkah-langkahnya dapat dilihat [disini](#menjalankan-unit-testing-dengan-Docker)
6. Push projek ke docker hub setelah semua test case berhasil dijalankan, langkah-langkahnya dapat dilihat [disini](#push-projek-ke-docker-hub)
7. Submit image docker yang telah dipush ke Docker Hub ke Halaman Submission Devcode, langkah-langkahnya dapat dilihat [disini](#push-projek-ke-docker-hub)

## Teknologi yang Digunakan

1. [Git](https://git-scm.com)
2. [GoLang](https://go.dev)
3. [Docker](https://www.docker.com)
4. [Mux](https://pkg.go.dev/github.com/gorilla/mux)
5. [Air](https://github.com/cosmtrek/air)
6. [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv#section-readme)

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
docker-compose up -d --force --recreate
```

## Menjalankan Unit Testing dengan Docker

Pastikan docker sudah terinstall di komputer anda. Jika belum silahkan install terlebih dahulu mengikuti instruksi dari tutorial diatas.

### Build Docker Image

Jalankan perintah berikut untuk Build docker image `docker build . -t {name}`

contoh :

```
docker build . -t golang-hello
```

### Jalankan Docker Image

Jalankan docker image dengan perintah `docker run -e PORT=3030 -p 3030:3030 {docker image}`

contoh:

```
docker run -e PORT=3030 -p 3030:3030 golang-hello
```

## Jalankan Unit Testing

pastikan port ketika menjalankan docker image sama dengan `API_URL` ketika ingin menjalankan unit testing

```
docker run --network="host" -e API_URL=http://localhost:3030 -e LEVEL=1 alfi08/hello-unit-testing
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
