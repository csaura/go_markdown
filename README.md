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

try:
```
curl https://infinite-chamber-4320.herokuapp.com/markdown -X POST -d 'text=CommonMark
        ==========

        CommonMark is a rationalized version of Markdown syntax,
        with a [spec][the spec] and BSD3-licensed reference
        implementations in C and JavaScript.

        [Try it now!](http://spec.commonmark.org/dingus.html)'
```
