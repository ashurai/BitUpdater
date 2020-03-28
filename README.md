# BitUpdater
A go lang demo app for bit coin real time value updater, no database in use, only local storage in use

#### Just hit command 
```
go run main.go
```
OR download the exe for windows
```
./BitUpdater.exe
```
or on other machines download the codebase and run go build in project root to get the signle executable file
```
go build
```
Only one external package is in use to handle routing and parsing easily path params, with net/http also we can achieve this,
but retrieving value from request will add some more steps so i tried to use gorilla mux package, so that code will look bit more cleaner.
```
github.com/gorilla/mux
```

There are more ways which can improve the performance, gorutines are in use to handle future problems faster as well, 
for ex: if we have 15 or 100 or more than that type or symbols we can easily handle with high scalability.

In main.go file use ```PreSymbolsList``` global variable to add more symbols, if you will not add a valid symbol
it will handle it and will not give the response for those configured symbols. 
Might be with channels you can face some challenges but yes they are more improvable. 
I am trying to add things what we can do/implement with golang features

#### EndPoints which we can use / acess to get single currency update
##### localhost:8099/currency/{symbol(ex:BTCUSD)}

#### EndPoints which we can use / acess to get all configured currency update
##### localhost:8099/currency/all


