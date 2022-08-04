/*
Copyright 2022 Naive Systems Ltd.

This software contains information and intellectual property that is
confidential and proprietary to Naive Systems Ltd. and its affiliates.
*/

package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1) //@ diag(`should call wg.Add(1) before starting`)
		wg.Done()
	}()
	wg.Wait()
}
