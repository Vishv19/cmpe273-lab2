package main 

import (
    "encoding/json" 
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"  
)

type jsonInput struct {
    Name string `json:"name"`
}

type jsonOutput struct {
    Greeting string `json:"greeting"`
}

func helloRequest(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {
    
    fmt.Fprintf(rw, "Hello, %s!\n", param.ByName("name"))

}   

func helloResponse(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

    input := jsonInput{}
    json.NewDecoder(req.Body).Decode(&input)

    output := jsonOutput{}
    output.Greeting = "Hello," + input.Name + "!"

    message , _ := json.Marshal(output)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", message)
}

func main() {
    router := httprouter.New()  
    router.GET("/hello/:name", helloRequest)
    router.POST("/hello", helloResponse)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: router,
    }
    server.ListenAndServe()
}
