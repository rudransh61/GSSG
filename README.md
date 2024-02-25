## GSSG

A Static Site Generator with Markdown Support in GoLang

___

To Use it , 
write some `.md` files in `content` folder 
And Run 
```shell
$ go run main.go
``` 

Now this server will watch all the changes in content folder and Update the HTML in `output` folder


## Example :

```md
---
title: My First Page
date: 2024-02-25
---

# Welcome to My First Page

This is a sample Markdown file for your static site generator. You can include various elements such as headers, paragraphs, lists, and more.

## Features

- Markdown to HTML conversion
- Dynamic page generation
- File system monitoring

Feel free to customize this content as needed. The front matter (enclosed in `---`) contains metadata like the title and date, which your generator can extract for each page.

```