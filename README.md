# Devcode golang starter with mux - Level 1

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

## Docker

Jika anda sudah menginstall docker, anda bisa menjalankan perintah `docker-compose up -d` untuk menjalankan API <b>Contact Manager</b> dan juga database <b>Mysql</b>. Tetapi pastikan `environment` pada file .env yang telah kamu buat dari .env.example sesuai dengan `environment` pada file `docker-compose.yaml`.

Apabila ada perubahan pada file kodingan anda, anda bisa build ulang container dengan perintah :
```
docker-compose up -d --force --recreate
``` 


## Menjalankan projek

- copy `.env.example` to `.env`
- install air `go install github.com/cosmtrek/air@latest` untuk menjalankan golang dengan mode development
- install package `go mod download`
- jalankan projek dengan perintah `go run main.go` atau `air` untuk mode development

# Menjalankan unit testing dengan docker

Pastikan docker sudah terinstall di komputer anda. Jika belum silahkan install terlebih dahulu mengikuti instruksi dari tutorial diatas.

## Build docker image
Jalankan perintah berikut untuk Build docker image  `docker build . -t {name}`

contoh :
```
docker build . -t golang-hello
```

## Jalankan docker image
Jalankan docker image dengan perintah `docker run -e PORT=3030 -p 3030:3030 {docker image}`

contoh: 
```
docker run -e PORT=3030 -p 3030:3030 golang-hello
```

### Jalankan unit testing

pastikan port ketika menjalankan docker image sama dengan `API_URL` ketika ingin menjalankan unit testing

```
docker run --network="host" -e API_URL=http://localhost:3030 -e LEVEL=2 alfi08/hello-unit-testing
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