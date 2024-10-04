package main

import "fmt"

type notifier interface {
	notify() error
}

type sms struct {
	message string
}

type email struct {
	message string
}

func (s sms) notify() error {
	fmt.Println(s.message)
	return nil
}

func (e email) notify() error {
	fmt.Println(e.message)
	return nil
}

func printNotification(n notifier) {
	n.notify()
}
