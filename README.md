SIMPLE CASE MICROSERVICES BACKEND USE GOLANG

### SUMMARY

_auth_ dir -> contains simple logic API wich can consume GET, POST method <br />
_payment_ dir -> contains simple logic API wich can consume GET, POST method

### PREREQUISITE

1. GO >= 1.24
2. Visual Studio Code

### API DOCS

POST /payment <br />
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
