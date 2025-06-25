package main

import (
    "fmt"
    "log"
    "net/http"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Query().Get("title")
    fmt.Fprintf(w, "您正在查询图书: %s", title)
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
    // 假设这里简单处理POST请求，实际可能需要更多错误处理
    fmt.Fprintf(w, `{"message": "评论提交成功", "user": "小李", "comment": "这本书真棒!"}`)
}

func main() {
    http.HandleFunc("/book", bookHandler)
    http.HandleFunc("/comment", commentHandler)
    fmt.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}