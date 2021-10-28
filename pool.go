package util

import (
	"fmt"
	"sync"
)

/**
一个简单的routine池
原理：生产者-消费者模式，使用方通过Submit方法生产数据（方法f），放入fChan，pool内部有poolSize个routine不断从fChan中取出方法执行
*/

type RoutinePool struct {
	poolSize int         //routine数
	chanSize int         //信道size
	fChan    chan func() //存放f的队列
	wg       sync.WaitGroup
}

/**
poolSize:内部运行的routine数
chanSize:存放方法f的信道size
*/
func NewRoutinePool(poolSize, chanSize int) *RoutinePool {
	p := RoutinePool{
		poolSize: poolSize,
		chanSize: chanSize,
		fChan:    make(chan func(), chanSize),
	}

	for i := 0; i < poolSize; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
		forLoop:
			for {
				select {
				case f, ok := <-p.fChan:
					if ok {
						f()
					} else {
						break forLoop
					}

				}
			}

		}()
	}

	return &p
}

/**
提交方法到pool，向已经Shutdown的pool Submit任务会出错
*/
func (p *RoutinePool) Submit(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("submit error:%v", r)
		}
	}()

	p.fChan <- f

	return
}

/**
阻塞直到所有提交的任务执行完毕
*/
func (p *RoutinePool) Wait() {
	p.wg.Wait()
}

/**
不再接受新的任务，已提交的任务会继续执行
*/
func (p *RoutinePool) ShutDown() {
	close(p.fChan)
}
