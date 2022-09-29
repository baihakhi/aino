# AINO [backend]
### Skilltest - web auth

By default app will be running in<br/>
`localhost:8081` for frontend<br/>
`localhost:8080` for backend 
Change the configuration for database in `.env`

## Build Setup
Make sure backend app already running before starting the frontend

``` bash
# get into repo
cd auth

# install dependencies
go mod tidy

# serve at localhost:8080
go run main.go

```