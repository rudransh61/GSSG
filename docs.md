# GSSG - GoLang Static Site Generator

## Introduction

GSSG is a minimal static site generator written in Go (Golang). It converts Markdown files into HTML and provides a basic web server for previewing your generated content. This guide will walk you through setting up and using GSSG.

## Prerequisites

- [Go](https://golang.org/) installed on your machine.

## Installation

1. Clone the GopherPages repository:

    ```bash
    git clone https://github.com/rudransh61/GSSG.git
    cd GSSG
    ```

2. Install the required Go packages:

    ```bash
    go get -u github.com/radovskyb/watcher
    go get -u github.com/russross/blackfriday/v2
    ```

3. Run the GSSG :

    ```bash
    go run main.go
    ```

## Getting Started

1. **Create Content:**

    Place your Markdown files in the `content` directory. For example, create a file named `example.md` with the following content:

    ```markdown
    <hr>
    title: Example Page
    date: 2024-02-25
    <hr>

    # Welcome to GopherPages

    This is an example page for your GSSG static site.
    ```

2. **Start GSSG:**

    Run the GSSG:

    ```bash
    go run main.go
    ```

    This will start the file watcher and the web server.

3. **Access the Generated Site:**

    Open your browser and go to [http://localhost:8080](http://localhost:8080). You should see your generated site.

4. **Preview Changes:**

    Make changes to your Markdown files in the `content` directory. GopherPages will automatically regenerate the site when you save changes. Refresh your browser to see the updates.

## Customize Styles

GSSG comes with a basic stylesheet. If you want to customize the styles, follow these steps:

1. Open the `styles/index.css` file.

2. Update the styles based on your preferences.

3. Save the changes.

4. GSSG will automatically detect the changes and apply the updated styles.

## Deploying Your Site

Once you are satisfied with your generated site, you can deploy it to a web hosting service. Simply copy the contents of the `output` directory to your desired hosting location.

## Conclusion

Congratulations! You've successfully set up and used GSSG to generate and preview your static site. Feel free to explore more features and customize the generator to suit your needs. If you encounter any issues or have suggestions for improvements, please check the [GitHub repository](https://github.com/rudransh61/GSSGs) and contribute or create an issue.