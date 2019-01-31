package main

import (
    "encoding/json"
    "io/ioutil"
    "luw/load"
    "luw/trie"
    "net/http"
    "path/filepath"
)

var tree = trie.NewTrie()

func init()  {
    items := load.Load()
    for _,item := range items {
        tree.Insert(item.Word)
    }
}

func main()  {
    http.HandleFunc("/word", func(writer http.ResponseWriter, request *http.Request) {
        prefix := request.FormValue("prefix")
        result,_ := tree.FindString(prefix)
        writer.Header().Set("Content-Type","application/json")
        b,err := json.Marshal(result)
        if err != nil {
            panic(err)
        }
        writer.Write(b)
    })

    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        c,err := load.CurrentDir()
        if err != nil {
            panic(err)
        }
        index := filepath.Join(c,"html/index.html")
        b,err := ioutil.ReadFile(index)
        if err != nil {
            panic(err)
        }
        writer.Header().Set("Content-Type","text/html; charset=UTF-8")
        writer.Write(b)
    })

    http.ListenAndServe(":9005",nil)
}
