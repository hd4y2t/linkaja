## Cara install

#required
- go version go1.18.1 
- mysql 10.4.28-MariaDB
- postman any version

#create database in your localhost:phpmyadmin
- create database linkaja
- import database dari folder go_linkaja yang bernama linkaja.sql

#how to use rest api
- buka folder yang telah di unduh
- lalu klik kanan lalu tekan open in terminal
- lalu ketikkan go run main.go 
- buka post man
- lalu request menggunakan method get ke link
    127.0.0.1:8001/api/account/555001
- buat request baru gunakan method post ke link
    127.0.0.1:8001/api/account/555001/transfer
    dengan body
    {
    "To_Account_Number":"555002",
    "Amount":5000
    }
