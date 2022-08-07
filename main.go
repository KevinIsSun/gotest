/*
Copyright 2022 Naive Systems Ltd.

This software contains information and intellectual property that is
confidential and proprietary to Naive Systems Ltd. and its affiliates.
*/

package main

import "errors"

func getError() (string, error) {
	return "msg", errors.New("some error")
}

func main() {
	msg, _ := getError()
	print(msg + "\n")
}
