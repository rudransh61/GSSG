# GSSG: GoLang Static Site Generator

GSSG is a lightweight static site generator written in Go (Golang). It converts Markdown files into HTML, allowing you to easily create and maintain static websites. GSSG comes with a built-in development server that enables you to preview your site locally while making changes.

## Features

- Simple and easy-to-use.
- Converts Markdown to HTML using Blackfriday.
- Automatic regeneration on file changes.
- Basic web server for local preview.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your machine.

### Installation

1. Clone the GSSG repository:

    ```bash
    git clone https://github.com/rudransh61/GSSG.git
    cd GSSG
    ```

2. Install the required Go packages:

    ```bash
    go get -u github.com/radovskyb/watcher
    go get -u github.com/russross/blackfriday/v2
    ```

3. Build the GSSG executable:

    ```bash
    $ go run main.go
    ```

### Usage

1. **Create Content:**

    Place your Markdown files in the `content` directory.

2. **Start GSSG:**

    Run the GSSG:

    ```bash
    $ go run main.go
    ```

    This will start the file watcher and the web server.

3. **Access the Generated Site:**

    Open your browser and go to [http://localhost:8080](http://localhost:8080). You should see your generated site.

4. **Preview Changes:**

    Make changes to your Markdown files in the `content` directory. GSSG will automatically regenerate the site when you save changes. Refresh your browser to see the updates.

### Customize Styles

GSSG comes with a basic stylesheet. If you want to customize the styles, open the `styles/index.css` file and update the styles based on your preferences.

### Deploying Your Site

Once you are satisfied with your generated site, you can deploy it to a web hosting service. Copy the contents of the `output` directory to your desired hosting location.

## Contributing

If you encounter any issues or have suggestions for improvements, please check the [GitHub repository](https://github.com/rudransh61/GSSG) and contribute or create an issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
