//Closing Channels
//Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.
package main

import "fmt"

func main() {
	//In this example we’ll use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine. When we have no more jobs for the worker we’ll close the jobs channel.
    jobs := make(chan int, 5)
    done := make(chan bool)

	//Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs. In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.
    go func() {
        for {
            j, more := <-jobs
			fmt.Println(j, more)
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}


// sent job 1
// sent job 2
// sent job 3
// sent all jobs
// 1 true
// received job 1
// 2 true
// received job 2
// 3 true
// received job 3
// 0 false
// received all jobs