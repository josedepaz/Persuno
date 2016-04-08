package main

import (
    "fmt"
        "net/http"

        "goji.io"
        "goji.io/pat"
        "golang.org/x/net/context"
)

func findAllPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Lista de posts")
}

func main() {
    mux := goji.NewMux()
    mux.HandleFuncC(pat.Get("/"), findAllPosts)

    http.ListenAndServe("localhost:8000", mux)
}