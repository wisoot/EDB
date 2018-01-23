# Event Driven Banking

This is the backend application where you can register a bank account then deposit and withdraw funds from your account.

You can access all functionalities via API. The entire system treat money in cent, thus if you want to deposit $10, you will need to pass 1000 into an amount field.

### Version Control

This project is managed by git version control, but the fact that you are reading this file means you have already known this.

### Database Engine

This project require MySQL to run at this point.

### Getting Started

#### Building Project

1. modify config.yaml file with appropriate settings
2. run `dep ensure` to install dependencies (read more about this underneath)
3. execute all sql files under database/migrations

#### Running Application

By compiling and run `main.go` you should have access to the API via http://localhost:8080.

##### *Dependencies Management*

In case that `dep ensure` doesn't work as it is still a prototype as stated on the github, alternatively you can run these following commands instead:

1. `go get github.com/gin-gonic/gin`
2. `go get github.com/gin-contrib/sessions`
3. `go get github.com/go-sql-driver/mysql`
4. `go get github.com/spf13/viper`
5. `go get golang.org/x/crypto`
6. `go get github.com/Masterminds/squirrel`
