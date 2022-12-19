# Devcode Starter using Golang and Mux Level 1

## Hasil akhir yang Diharapkan

Peserta dapat menampilkan message hello world dalam format JSON pada url http://localhost:3030/hello dan submit challenge di Devcode menggunakan Docker

## Setup Environment

1. Download source code melalui link yang telah disediakan dari halaman assesment
2. Extract source code yang sudah terdownload pada perangkat anda
3. Buka source code yang sudah diextract menggunakan Code Editor, contoh Visual Studio Code
4. Salin isi dari file `.env.example` ke dalam file `.env`
5. install air `go install github.com/cosmtrek/air@latest` untuk menjalankan golang dengan mode development
6. Jalankan `go mod download` pada terminal untuk install packages
7. Jalankan `go run main.go` atau `air` untuk mode development pada terminal

## Instruksi pengerjaan

1. Pastikan anda sudah meng-install tools yang diperlukan. Jika belum, silahkan ikuti langkah-langkahnya [disini](#menginstal-tools-yang-diperlukan)
2. Jalankan API dan Database dengan Docker, silahkan ikuti langkah-langkahnya [disini](#menjalankan-api-dan-database-dengan-docker)
3. Lakukan unit testing pada local anda dengan menggunakan Docker, langkah-langkahnya dapat dilihat [disini](#menjalankan-unit-testing-dengan-Docker)
4. Push projek ke docker hub setelah semua test case berhasil dijalankan, langkah-langkahnya dapat dilihat [disini](#push-projek-ke-docker-hub)
5. Submit image docker yang telah dipush ke Docker Hub ke Halaman Submission Devcode, langkah-langkahnya dapat dilihat [disini](#push-projek-ke-docker-hub)

## Tools dan Packages yang digunakan

1. [Git](https://git-scm.com)
2. [Docker](https://www.docker.com)
3. [Mux](https://pkg.go.dev/github.com/gorilla/mux)
4. [Air](https://github.com/cosmtrek/air)
5. [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv#section-readme)

## Menginstal Tools yang digunakan

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
