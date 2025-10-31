package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)    // channel chứa công việc
	results := make(chan int, 5) // channel chứa kết quả

	// Tạo 3 worker (goroutine)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Gửi 5 công việc vào channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // đóng channel jobs (không gửi thêm job nữa)
	fmt.Println("close channel")

	// Nhận kết quả từ 5 công việc
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Kết quả nhận được: %d\n", result)
	}

	fmt.Println("Hoàn tất tất cả công việc.")
}

// Hàm worker chạy ở goroutine riêng
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d đang xử lý job %d...\n", id, job)
		time.Sleep(time.Second) // giả lập xử lý tốn thời gian
		results <- job * 2      // gửi kết quả vào channel
	}
}
