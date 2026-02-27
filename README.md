SIMPLE CASE MICROSERVICES BACKEND USE GOLANG

### BRANCH DETAIL

Because this project is for personal RnD, this detail source code for each branch

- main -> simple usecase microservices, and deployed with docker compose
- kong -> add deployment for kubernetes, and use kong for gateway
- envoy -> add deployment for kubernetes, and use envoy for gateway
- envoy-keda -> add deployment for kubernetes, use envoy for gateway, and add KEDA for autoscaling

### SUMMARY

_auth_ dir -> contains simple logic API wich can consume GET, POST method <br />

_products_ dir -> contains simple logic API wich can consume GET, POST method

### PREREQUISITE

1. GO >= 1.24
2. Visual Studio Code

### API DOCS

POST /products <br />
[{
"id": 102,
"title": "Smartwatch XYZ",
"price": 99.99,
"description": "Smartwatch canggih dengan fitur detak jantung dan pelacakan tidur.",
"category": "Elektronik",
"image": "https://example.com/images/smartwatch.jpg",
"quantity": 50
}]

POST /auth <br />
{
"id": 1,
"name": "dimas",
"age": 20
}
