
# ASCII Art Web

## Description
```markdown
Ascii-art-web is an HTTP server written in Go that provides a web interface for generating ASCII art based on user input using different banners. This project extends the previous ascii-art project and allows users to choose one of the following banners:

- **shadow**
- **standard**
- **thinkertoy**

The server implements two primary HTTP endpoints:
- **GET /** – Returns an HTML page with a form for input.
- **POST /ascii-art** – Receives text and the selected banner, processes the input, and returns the generated ASCII art.

The server responds with appropriate HTTP status codes:
- **200 OK** – If the request is processed successfully.
- **404 Not Found** – If required templates or banner files are not found.
- **400 Bad Request** – For invalid requests.
- **500 Internal Server Error** – For any unexpected errors.

## Authors

- Damilya Amangeldykyzy
- Tamerlan Temirkhanov 
- Arnur Kabdylkak 

## Usage: How to Run

1. **Install Go**:  
   Ensure that Go is installed on your computer. You can download it from the [official Go website](https://golang.org/dl/).

2. **Clone the Repository**:
   ```bash
   git clone https://01.tomorrow-school.ai/git/damangel/ascii-art-web.git
   cd ascii-art-web
   ```

3. **Run the Server**:
   ```bash
   go run main.go
   ```
   By default, the server starts on port `8080`. You can change the port in the configuration if needed.

4. **Access the Application**:  
   Open your web browser and navigate to [http://localhost:8080](http://localhost:8080). You will see the main page where you can input text, select a banner, and generate the ASCII art.

## Implementation Details: Algorithm

1. **Template Rendering**:  
   The main page is rendered using Go's `html/template` package. All HTML templates are located in the `templates` directory at the root of the project.

2. **Form Handling**:  
   The form on the main page collects the following data:
   - **Text Input**: The string that will be converted into ASCII art.
   - **Banner Selection**: An option to choose one of the following banners: `shadow`, `standard`, or `thinkertoy`.

3. **Request Processing**:
   - When the form is submitted, a POST request is sent to the `/ascii-art` endpoint.
   - The server parses the form data and validates the input.
   - Depending on the selected banner, the corresponding ASCII template file is loaded.
   - Each character of the input text is converted to its ASCII art representation according to the banner file, and the lines are concatenated to form the final output.

4. **Response Formation**:
   - The generated ASCII art is sent back as part of an HTML response.
   - Proper HTTP status codes (200, 400, 404, 500) are used to indicate the result of the request processing.

## Instructions

- **Templates Directory**:  
  Ensure that the HTML templates are placed in the `templates` directory at the root of the project.

- **Banner Files**:  
  The banner files (`shadow.txt`, `standard.txt`, `thinkertoy.txt`) must be available to the server according to the logic implemented in the project.

- **Error Handling**:  
  The server is designed to handle errors gracefully, including missing templates, invalid form data, and any unexpected issues.

- **Code and Standards**:  
  The code follows best practices in Go and uses only the standard Go packages.

---
```