package loadtest

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

const format =  "---------------------"+"当前1s内的执行数"+"/"+"总的执行数"+"/"+"平均的执行数"+"/"+"执行的时间"+"/"+"最长执行时间"


type Runner struct {
	threamNum int
	lastTimes int64
	nowTimes *int64
	runnedSecond int64
	maxRunTime int
	runfunction func()
}


func (runner *Runner) SetThreadNum(threadNum int)(*Runner)  {
	runner.threamNum = threadNum
	return runner
}
func (runner *Runner) SetRunFunction(r func())(*Runner)  {
	runner.runfunction = r;
	return runner
}
func (runner *Runner) Start(){
	for i:=0;i<runner.threamNum ;i++  {
		go func() {
			for ; ;  {
				runner.runfunction()
				atomic.AddInt64(runner.nowTimes, 1)
			}
		}()
	}
	for ;;{
		runner.runnedSecond++
		time.Sleep(time.Second)
		var nowtmp int64= *runner.nowTimes
		fmt.Println(format)
		fmt.Println("---------------------%d/%d/%d/%d/%dms",(nowtmp - runner.lastTimes),nowtmp,(nowtmp/runner.runnedSecond),runner.runnedSecond,runner.maxRunTime);
		runner.lastTimes = nowtmp
	}
}
func New()*Runner{
	runner := new(Runner)
	runner.threamNum = runtime.NumCPU()
	runner.lastTimes = 0
	runner.nowTimes = new(int64)
	runner.runnedSecond = 0
	runner.maxRunTime = 0

	return runner
}
