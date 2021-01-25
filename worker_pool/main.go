package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id int
	RandNum int
}

type Result struct {
	Job *Job
	SumNum int
}

func main() {
	jobChan := make(chan *Job, 128)

	resChan := make(chan *Result, 128)
	createPoll(40, jobChan, resChan)

	go func(resChan chan *Result) {
		for res := range resChan {
			fmt.Println("job id:%v randnum:%v result:%d\n", res.Job.Id,
                res.Job.RandNum, res.SumNum)
		}
	}(resChan)

	var id = int(1000)
	for id<10000 {
		id++
		r := &Job{
			Id: id,
			RandNum: rand.Int(),
		}
		jobChan <- r
	}



}

func createPoll(num int, jobChan chan * Job, resChan chan * Result)  {

	for i:=1;i<num;i++ {
		go func(jobChan chan *Job, resChan chan * Result) {
			for job:=range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp :=r_num/10
					sum += tmp
					r_num /=10
				}
				r := &Result{
					Job: job,
					SumNum: sum,
				}
				resChan <- r
			}

		}(jobChan, resChan)
	}
}