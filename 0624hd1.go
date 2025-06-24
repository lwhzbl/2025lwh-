package main

import (
    "fmt"
    "sync"
)

func main() {
    var (
        counter int
        mutex   sync.Mutex
        wg      sync.WaitGroup
    )

    // 设置需要启动的协程数量
    numWorkers := 50
    // 每个协程需要增加的次数
    incrementsPerWorker := 100

    wg.Add(numWorkers)

    // 启动50个协程
    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            defer wg.Done()

            for j := 0; j < incrementsPerWorker; j++ {
                // 加锁保证线程安全
                mutex.Lock()
                counter++
                mutex.Unlock()
            }

            fmt.Printf("Worker %d finished. Current counter: %d\n", workerID, counter)
        }(i)
    }

    // 等待所有协程完成
    wg.Wait()

    fmt.Printf("All workers finished. Final counter value: %d\n", counter)
}