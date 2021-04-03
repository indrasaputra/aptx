# URL Shortener

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/url-shortener)](https://goreportcard.com/report/github.com/indrasaputra/url-shortener)
[![Workflow](https://github.com/indrasaputra/url-shortener/workflows/Test/badge.svg)](https://github.com/indrasaputra/url-shortener/actions)
[![codecov](https://codecov.io/gh/indrasaputra/url-shortener/branch/main/graph/badge.svg?token=VI4V05KUEO)](https://codecov.io/gh/indrasaputra/url-shortener)
[![Maintainability](https://api.codeclimate.com/v1/badges/e28a29089f4c66303cb0/maintainability)](https://codeclimate.com/github/indrasaputra/url-shortener/maintainability)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=indrasaputra_url-shortener&metric=alert_status)](https://sonarcloud.io/dashboard?id=indrasaputra_url-shortener)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/url-shortener.svg)](https://pkg.go.dev/github.com/indrasaputra/url-shortener)

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## API

### gRPC

The API can be seen in proto files (`*.proto`) in directory `proto/`.

### RESTful JSON

The API is automatically generated in OpenAPIv2 format when generating gRPC codes.
The generated files are stored in directory `openapiv2` in JSON format (`*.json`).
To see the RESTful API contract, do the following:
- Open the generated json file(s)
- Copy the content
- Open [https://editor.swagger.io/](https://editor.swagger.io/)
- Paste the content in [https://editor.swagger.io/](https://editor.swagger.io/)

## How to Run

- Read [Prerequisites](doc/PREREQUISITES.md)
- Then, read [How to Run](doc/HOW_TO_RUN.md)

## Development Guide

- Read [Prerequisites](doc/PREREQUISITES.md)
- Then, read [Development Guide](doc/DEVELOPMENT_GUIDE.md)

## Code Map

- Read [Code Map](doc/CODE_MAP.md)

## FAQs

- Read [FAQs](doc/FAQS.md)