package main

import (
   "log"
   "github.com/ayush/secure-api/db"
   "github.com/ayush/secure-api/server"
   "github.com/ayush/secure-api/server/middleware/myJWT"
)

var host = "localhost"
var port = "8000"

func main() {
   db.InitDB()
   jwtError := myJWT.InitJWT()
   
   if jwtError != nil {
      log.Println("Error initializing JWT")
      log.Fatal(jwtError)
   }
   
   serverError := server.StartServer(host, port)
   if serverError != nil {
      log.Println("Error starting server!")
      log.Fatal(serverError)
   }
}