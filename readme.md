
# ASCII Art Web

ASCII Art Web is an HTTP server written in Go that provides a web interface for generating ASCII art from user input using various banner styles. This project enhances the original ascii-art tool by offering multiple banner options: `shadow`, `standard`, and `thinkertoy`.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Implementation Details](#implementation-details)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [Authors](#authors)
- [License](#license)

---

## Overview

ASCII Art Web transforms plain text into visually appealing ASCII art using pre-defined banner templates. It leverages Go's powerful standard library for HTTP handling and template rendering, ensuring a responsive and robust user experience.

---

## Features

- **Dynamic ASCII Art Generation:** Converts user-provided text into ASCII art.
- **Multiple Banner Styles:** Supports three distinct banner styles:
  - `shadow`
  - `standard`
  - `thinkertoy`
- **User-Friendly Web Interface:** Easy-to-use HTML form for input submission.
- **Robust HTTP Status Handling:** 
  - `200 OK` for successful operations.
  - `400 Bad Request` for invalid inputs.
  - `404 Not Found` when templates or banner files are missing.
  - `500 Internal Server Error` for unforeseen errors.

---

## Technology Stack

- **Programming Language:** Go (Golang)  
  A statically typed, compiled language known for its simplicity, efficiency, and strong concurrency support.

- **HTTP Server:**  
  Built using Go's standard `net/http` package, ensuring a lightweight and efficient server implementation.

- **Template Engine:**  
  Utilizes Go's `html/template` package for secure and dynamic HTML rendering.

- **Project Structure:**  
  Organized into distinct directories:
  - `templates/` for HTML files.
  - Banner files (e.g., `shadow.txt`, `standard.txt`, `thinkertoy.txt`) managed as per the project's logic.

- **Error Handling:**  
  Leverages Go's native error handling for robust, predictable responses in various scenarios.

- **Dependencies:**  
  This project relies solely on Go's standard library, reducing external dependencies and simplifying deployment.

---

## Project Structure

```plaintext
ascii-art-web
├── ascii-art
│   ├── banners
│   │   ├── banners.go
│   │   ├── shadow.txt
│   │   ├── standard.txt
│   │   └── thinkertoy.txt
│   ├── cmd
│   │   └── main.go
│   └── pkg
│       ├── ascii
│       │   ├── template.go
│       │   └── template_test.go
│       ├── handlers
│       │   ├── ascii_art.go
│       │   ├── home.go
│       │   └── render.go
│       └── middleware
│           ├── logging.go
│           └── recovery.go
├── go.mod
├── main.go
├── readme.md
├── static
│   ├── images
│   │   ├── bezel.png
│   │   └── scanlines.png
│   └── index.js
└── templates
    ├── 400.html
    ├── 404.html
    ├── 500.html
    └── index.html
```

---

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)

### Steps

1. **Clone the Repository**
   ```bash
   git clone https://01.tomorrow-school.ai/git/damangel/ascii-art-web.git
   cd ascii-art-web
   ```

2. **Run the Server**
   ```bash
   go run main.go
   ```
   The server will start on port `8080` by default. Modify the configuration if a different port is required.

---

## Usage

1. Open your web browser and navigate to: [http://localhost:8080](http://localhost:8080)
2. Enter the text you wish to convert into ASCII art.
3. Select the desired banner style (`shadow`, `standard`, or `thinkertoy`).
4. Submit the form to generate and view the ASCII art.

---

## Implementation Details

### Template Rendering

- Utilizes Go's `html/template` package for rendering HTML pages.
- HTML templates are stored in the `templates` directory at the project root.

### Form Handling

- The main page form captures:
  - **Text Input:** The string to be converted.
  - **Banner Selection:** Choice of banner style for ASCII art generation.

### Request Processing

- **GET /**: Serves the main HTML page containing the input form.
- **POST /ascii-art**:
  - Validates the submitted form data.
  - Loads the corresponding banner file based on user selection.
  - Maps each character of the input to its ASCII art representation.
  - Concatenates the generated lines to form the final ASCII art output.
  
---

## Error Handling

- **200 OK:** Successfully processed request.
- **400 Bad Request:** Indicates invalid form data.
- **404 Not Found:** Triggered when required HTML templates or banner files are missing.
- **500 Internal Server Error:** For any unexpected errors during processing.

---

## Contributing

Contributions are welcome! If you'd like to enhance ASCII Art Web, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and open a pull request.

---

## Authors

- **Damilya Amangeldykyzy**
- **Tamerlan Temirkhanov**
- **Arnur Kabdylkak**

---

## License

This project is licensed under the [MIT License](LICENSE).

---

