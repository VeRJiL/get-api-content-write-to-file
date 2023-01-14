## this a simple application which you can get comics from xkcd.com API.

## How to run the app
There are 2 types of application in this project `get` and `search`
<br>
`get` itself has 2 algorithm `normal` and `concurrent`

- you can run the application for getting the comics using normal algorithm 
`go run main.go get normal 100` which number `100` specifies how many algo should it get

- you can run the application for getting the comics using concurrent algorithm 
`go run main.go get concurrent 100` which number `100` specifies how many algo should it get

- you can run the application for searching among the found comics after running 
the above commands using `go run main.go search normal sleep red hat` which `normal`
specifies the filename which obtained from `normal algo` or you can specify `concurrent` as file name to search
in the file which obtained from `concurrent algo`. and any args after filename will be considered as search terms.
obviously you have run the get app before search in order to get the data.


## TODO:
- add a failure safe condition for concurrent version of the `get` application