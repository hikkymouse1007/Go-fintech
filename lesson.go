package main

import (
 "fmt"
 "io"
 "log"
 "os"
)

func loggingSettings(logFile string){
    logfile, _:= os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    multiLogFile := io.MultiWriter(os.Stdout, logfile)
    //log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
    log.SetOutput(multiLogFile)
}

func main() {
    loggingSettings("test.log")
 //file, error :=os.Open("hogehogheohg")
    _, err :=os.Open("hogehogheohg")
    if err != nil{
        log.Fatalln("exit", err)
    }

    log.Println("logging!")
    log.Printf("%T %v", "test", "test")

 //log.Fatalf("%T %v", "test", "test")

 //Fatalは実行されるとそこでコードが終了する
     log.Fatalln("error!")

     fmt.Println("ok!")
}




