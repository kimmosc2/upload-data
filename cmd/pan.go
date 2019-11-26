package main

import (
    "log"
    "net/http"
    "os"
    "upload/heartbeat"
    "upload/locate"
    "upload/objects"
)

func main() {
    // log.SetFlags(log.Llongfile)
    go heartbeat.StartHeartBeat()
    go locate.StartLocate()
    http.HandleFunc("/objects/", objects.Handler)
    log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
