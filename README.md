# Simple App Stock

## Requirements

- MySQL
- Golang
- NodeJS

## Installation

- Clone this repository

```bash
git clone https://github.com/OkkyRoyDirgantara/fiber-vue.git
```

### Installation Backend Golang

- Go to folder backend

```bash
cd fiber-app
```

- Install dependencies

```bash
go mod download
```

- Create database

```bash
mysql -u root -p
```

```bash
CREATE DATABASE fiber-vue;
```

- Copy .env.example to .env

```bash
cp .env.example .env
```

- Edit .env file with your configuration

- Run backend on port 3000

```bash
go run app.go
```

### Installation Frontend Vue

- Go to folder frontend

```bash
cd vue-app
```

- Install dependencies

```bash
npm install
```

- Run frontend

```bash
npm run serve
```

## Postman API

- [Postman API](https://www.postman.com/restless-satellite-772023/workspace/fiber-vue)
