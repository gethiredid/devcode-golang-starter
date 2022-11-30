# Boilerplate golang

## Package yang digunakan

- mux
- air

## Menjalankan projek

- install air `go install github.com/cosmtrek/air@latest` untuk menjalankan golang dengan mode development
- install package `go mod download`
- set port di dengan perintah `set PORT=5000` untuk os windows atau `export PORT=5000` untuk os linux/macos
- jalankan projek dengan perintah `go run main.go` atau `air` untuk mode development

# Menjalankan unit testing dengan docker

Pastikan docker sudah terinstall di komputer anda.

## Build docker image
Jalankan perintah berikut untuk Build docker image  `docker build . -t {name}`

contoh :
```
docker build . -t golang-hello
```

## Jalankan docker image
Jalankan docker image dengan perintah `docker run -e PORT=5000 -p 5000:5000 {docker image}`

contoh: 
```
docker run -e PORT=5000 -p 5000:5000 golang-hello
```

### Jalankan unit testing

pastikan port ketika menjalankan docker image sama dengan `API_URL` ketika ingin menjalankan unit testing

```
docker run --network="host" -e API_URL=http://localhost:5000 alfi08/hello-unit-testing
```


## Upload projek ke docker hub
Pastikan sudah memiliki akun docker hub, dan login akun docker anda di lokal dengan perintah `docker login`.

Setelah itu jalankan perintah berikut untuk push docker image lokal ke docker hub.

```
docker tag golang-hello {username docker}/golang-hello
docker push {username docker}/golang-hello
```