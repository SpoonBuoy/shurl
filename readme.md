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

## Usage

    Clone the repository.
    Install dependencies using go mod tidy.
    Start Redis Server on port 6379 and change the credentials.
    Run the project using go run main.go.
    Server will listen on port 8080

## Tools Used 🚀

- **Go (Golang)** 
- **Gin** 
- **Redis** 
- **GORM** 

