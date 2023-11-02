<div align="center">
      <h1><br/>Basic Trade API</h1>
</div>

# Description
A simple trade API project with CRUD operations, authentication using JWT, and powered by Golang, Gin, MySQL, and GORM.

# Features
This API is developed using Golang, Gin web framework, GORM for database operations, and JWT for authentication.

# Tech Used
![Golang](https://img.shields.io/badge/golang-%23F7DF1E.svg?style=for-the-badge&logo=go&logoColor=black)
![Gin](https://img.shields.io/badge/gin-%2361DAFB.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![GORM](https://img.shields.io/badge/gorm-%2300f.svg?style=for-the-badge&logo=go&logoColor=white)
![JWT](https://img.shields.io/badge/jwt-%2300f.svg?style=for-the-badge&logo=jwt&logoColor=white)

# Getting Started:
Before running the program, make sure you've installed the required dependencies.

- Install Golang: [Official Golang Installation Guide](https://golang.org/doc/install)
- Install Gin: `go get github.com/gin-gonic/gin`
- Install GORM: `go get gorm.io/gorm`
- Install MySQL driver for GORM: `go get gorm.io/driver/mysql`
- Install JWT library: `go get github.com/dgrijalva/jwt-go`

### Database setup:
Create your MySQL database and update the database configuration in the project.

### Run the program:
```shell
go run main.go
