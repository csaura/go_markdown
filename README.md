go_markdown
===========

Web service to send text in markdown and receive the parsed text

To run the service:

```
PORT=8080 ./web
```


GET example:
```
curl 'localhost:8080/markdown?text=foo' -X GET
````

POST example:
````
curl localhost:8080/markdown -X POST -d 'text=_fooo_ asdfasdfasd'
````
