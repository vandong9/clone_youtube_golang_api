# clone_youtube_golang_api

reference:
Gin example: https://github.com/restuwahyu13/go-rest-api


API:

- Main: handle init server and route
- Route: cordinate the path to handler by Transport layer
- Transport:  
    - Parse context to get input
    - call usecase to handle
    - receive the reponse from usecase, process and reponse back to request context
- usecase: 
    + This is business layer
    + Receive the input model and process business layer 
    + Using storage to CRUD to database
    + Return the reponse model to transport
- storage:
    + used by use-case
    + connect to database
    + query database
    + process the update to database
- Cache:
