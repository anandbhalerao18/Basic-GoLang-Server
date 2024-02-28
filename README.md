# Basic-GoLang-Server
## Simple Golang HTTP Server

This project is a simple Golang HTTP server that:

* Serves static files from the `./static` directory.
* Handles form submissions from the `/form` path.
* Responds with "hello" to GET requests on the `/hello` path.

### Prerequisites

* Golang installed and configured ([https://go.dev/doc/install](https://go.dev/doc/install))

### Running the Server

1. Build the program:

```
go build
```

2. Run the server:

```
./<program_name>
```

This will start the server on port 8080.

### Usage

1. **Static Files:** Access any file within the `./static` directory by navigating to `http://localhost:8080/<filename>`.
2. **Form Submission:**
    * Navigate to `http://localhost:8080/form`.
    * Fill out the form with your name and address.
    * Submit the form.
    * The server will print the submitted values to the console.
3. **"Hello" Message:** Access `http://localhost:8080/hello` to see the "hello" message displayed.

### Code Breakdown

* **`formhandler` function:** Handles POST requests to the `/form` path.
    * Parses the form data.
    * Prints the values of "name" and "address" fields.
* **`hellohandler` function:** Handles GET requests to the `/hello` path.
    * Checks for the correct path and method.
    * Prints "hello" as the response.
* **`main` function:**
    * Creates a file server for serving static files.
    * Registers handlers for the `/`, `/form`, and `/hello` paths.
    * Starts the server on port 8080.

This README provides a basic understanding of the project. You can further explore and customize it based on your needs.
