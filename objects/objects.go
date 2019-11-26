package objects

import (
    "io"
    "log"
    "net/http"
    "os"
    "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        get(w,r)
        return
    }
    if r.Method == http.MethodPut {
        put(w,r)
        return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
}

func get(w http.ResponseWriter, r *http.Request) {
    f, err := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    defer f.Close()
    io.Copy(w,f)
}

func put(w http.ResponseWriter, r *http.Request) {
    f, err := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
    if err != nil {
        log.Println(err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer f.Close()
    io.Copy(f,r.Body)
}
