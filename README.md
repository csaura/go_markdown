go_markdown
===========

Web service to send text in markdown and receive the parsed text

GET example:
```
curl 'http://192.168.128.47:8080/markdown?text=foo' -X GET
````

POST example:
````
curl http://192.168.128.47:8080/markdown -X POST -d 'text=_fooo_ asdfasdfasd'
````
