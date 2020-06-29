# Splio's Velib Finder

This API allows Splio employees of Splio HQ to know the number of available velib of nearby stations. 

### Package description 

- dao : Responsible for getting data from various sources 
- handler : API structure and endpoint exposition
- model : all used data model related to application logic

### Launching

At folder source, please run : 

````go run main.go````

### Testing 

To test the health-check function, in folder /handler, please run : 
 
````go test````

### Possible amelioration 

Velib API also exposes the number of mechanical and electric Velib available per station, which can be distributed through the API. As we all know, electric ones are way better ;-) 