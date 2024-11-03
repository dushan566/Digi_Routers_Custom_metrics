package api

import(
  "github.com/joho/godotenv"
  "log"
)

func GetEnvVars(){
  err := godotenv.Load("secrets/credentials.env")
  if err != nil{
    log.Fatal("Error loading .env credential file")
  }
}
