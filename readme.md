# Shurl - Short URL
A simple URL Shortening Tool
<br>
Find it here :
[Shurl](https://arsalan.live/shurl)

## Routes

1. **Create Short URL**
   - **Method:** POST
   - **Endpoint:** `/create-short-url`
   

2. **Short URL Redirect**
   - **Method:** GET
   - **Endpoint:** `/:shortUrl`
   

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

