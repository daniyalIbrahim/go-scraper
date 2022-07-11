## go-scraper
A scraper built using golang and a scraping library go-rod. it is capable of taking screenshots of websites as a pdf or png format
    

#### Simple Website Screenshot or Website 2 Pdf

[![Docker](https://github.com/techonomylabs/go-scraper/actions/workflows/deploy-to-cloud-run.yml/badge.svg)](https://github.com/techonomylabs/go-scraper/actions/workflows/deploy-to-cloud-run.yml)
<img src="https://img.shields.io/github/workflow/status/techonomylabs/go-scraper/Docker?label=GCP%20CLOUD%20RUN"/>
<img src="https://img.shields.io/github/license/techonomylabs/go-scraper" />
<a href="https://github.com/techonomylabs/go-scraper/issues">
<img src="https://img.shields.io/github/issues/techonomylabs/go-scraper" />
</a>
<img src="https://img.shields.io/github/languages/count/techonomylabs/go-scraper?style=flat-square"/>


### Rest api with Go-rod for Webscraping
There are two http endpoints in this project.

    1. GET /img?url=<url>
    2. GET /pdf?url=<url>
Response is application/image or application/pdf depending on the request.

Some usual cli commands to get you going

    go fmt ./...

    go get -u ./...

    go mod tidy

    go mod vendor

    go build -o main.exe main.go   # Build the program


