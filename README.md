### GoLang Sample app for RSS Feed

[Tutorial](https://www.youtube.com/watch?v=dpXhDzgUSe4)

#### initialize project 
- go mod init projectname
#### build and run
- go build && ./goLang-app-rss-feed-scraper
#### added env config package
- go get github.com/joho/godotenv
#### create vendor dir
- go mod vendor
#### clean mod file 
- go mod tidy
### run query
- sqlc generate

#### install http module
- go mod tidy
- go mod vendor
- go get -u github.com/go-chi/chi/v5
- go get  github.com/go-chi/cors

- both are need for database 
    - go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    - go install github.com/pressly/goose/v3/cmd/goose@latest


