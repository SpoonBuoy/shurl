# Shurl - Short URL Project

Shurl is a simple short URL project built using the Gin framework in Go. It provides two main routes for creating short URLs and handling short URL redirects.

## Routes

1. **Create Short URL**
   - **Method:** POST
   - **Endpoint:** `/create-short-url`
   - **Handler Function:** `handler.CreateShortUrl`
   - **Description:** This route allows you to create a short URL by providing the original URL in the request body.

2. **Short URL Redirect**
   - **Method:** GET
   - **Endpoint:** `/:shortUrl`
   - **Handler Function:** `handler.HandleShortUrlRedirect`
   - **Description:** Accessing this route with a short URL will redirect the user to the original URL associated with that short URL.

## Config
   Modify the `config.json` file <br>
   ###
         {
            //port on which backend runs
            "port": ":8080", 

            //host addr of backend
            "host": "localhost/", 

            //redis server
            "redis-addr": "localhost:6379",

            //redis credentials
            "redis-password": "password", 
            
            "redis-db": 0 
         }

## Usage

    Clone the repository.
    Install dependencies using go mod tidy.

    Start redis server 
    On docker use this ❯ docker run --name rdb -p 6379:6379 redis redis-server --requirepass your-password

    go build
    ./url-shortener

## Tools Used 🚀

- **Go (Golang)** 
- **Gin** 
- **Redis** 

