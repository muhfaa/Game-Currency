# Game Currency

Game currency is a system used to convert the game currency owned, based on the amount and rate owned by two different currencies.

In this application I use the Hexagonal architecture with Golang as the programming language and MySQL as the database

In this application I created 4 endpoints, 2 endpoints on the currency side and 2 more endpoints on the conversion side. The following is the request structure of this application.

[Game Currency](https://github.com/muhfaa/Game-Currency/blob/main/index.md)

# Database diagram
![Database diagram](https://github.com/muhfaa/Game-Currency/blob/a13d95bc1bfe95040f185bd8dab01e36db1ba9a4/ATTN.png)

# SQL Query
[SQl query to create table](https://github.com/muhfaa/Game-Currency/blob/dc785ccb8ed4e75bd23eceed9392480e9c5787f5/migrations/1.create_tabel_currency_and_conversion.up.sql)
## Requirement

- go 1.16+
- Mysql
- Docker

## Config

For config file i use `json`

## Running App

`docker-compose up --build`

- After that open [url](http://localhost:8090/?server=golang-game-currency-db&db=game-currency) and `username: root` `password: root`
- Next execute this [SQl query to create table](https://github.com/muhfaa/Game-Currency/blob/dc785ccb8ed4e75bd23eceed9392480e9c5787f5/migrations/1.create_tabel_currency_and_conversion.up.sql) in that SQL query
