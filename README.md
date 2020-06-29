# Splio's Velib Finder

This API allows Splio employees of Splio HQ to know the number of available velib of nearby stations. 

### Package description 

- handler : API structure and endpoint exposition
- dao : Responsible for getting data from various sources 
- model : all used data model related to application logic

### Launching

At folder source, please run : 

````go run main.go````

### Testing 

To test the health-check function, in folder /handler, please run : 
 
````go test````

### Possible amelioration 

- Include a post endpoint for the user to add his favorites stations   