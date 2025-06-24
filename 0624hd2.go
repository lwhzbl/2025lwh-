package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    // 创建两个通道用于同步
    letterCh := make(chan struct{})
    numberCh := make(chan struct{})

    // 启动字母打印协程
    go func() {
        defer wg.Done()
        for c := 'A'; c <= 'Z'; c++ {
            // 等待信号
            <-letterCh
            fmt.Printf("%c", c)
            // 发送信号给数字协程
            numberCh <- struct{}{}
        }
    }()

    // 启动数字打印协程
    go func() {
        defer wg.Done()
        for i := 1; i <= 26; i++ {
            // 等待信号
            <-numberCh
            fmt.Printf("%d", i)
            if i < 26 {
                // 发送信号给字母协程（最后一次不需要发送）
                letterCh <- struct{}{}
            }
        }
    }()

    // 启动第一个协程
    letterCh <- struct{}{}

    // 等待两个协程完成
    wg.Wait()
    fmt.Println()
}