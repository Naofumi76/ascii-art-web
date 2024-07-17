
# ASCII Art Web Application

This is a simple web application written in Go that converts user input text into ASCII art. The application provides a web interface where users can enter text, select an ASCII art style, and view the generated ASCII art. Additionally, users can download the generated ASCII art as a text file.

## Features

-   Convert input text to ASCII art in various styles (standard, shadow, thinkertoy).
-   Display the generated ASCII art on the web page.
-   Download the generated ASCII art as a text file.
-   Handle 404 errors for unknown routes.
-   Use HTML templates for rendering web pages.

## Requirements

-   Go 1.22.4 or later

## Installation

1.  Clone the repository:
```bash
git clone https://github.com/Naofumi76/ascii-art-web.git
cd ascii-art-web
```
2. Make sure you have Go installed, then build and run the application:
```bash
go run .
```

3. Open your web browser and navigate to `http://localhost:8080` to use the application.

## Usage

### Home Page

1.  Enter your text in the textarea provided.
2.  Select the desired ASCII art style from the dropdown menu.
3.  Click the "Submit" button to generate the ASCII art.

### Downloading ASCII Art

1.  After generating the ASCII art, a download link will appear.
2.  Click the "Download" button to download the ASCII art as a text file.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

