package main

import (
	"github.com/TestCoders/Sinterklaas2020/pkg/aliblabla"
	"github.com/TestCoders/Sinterklaas2020/pkg/bollie"
	"github.com/TestCoders/Sinterklaas2020/pkg/coolbere"
	"sync"
)

func bollieServer() {
	app := bollie.NewApplication()
	app.StartServer()
}

func coolbereServer() {
	app := coolbere.NewApplication()
	app.StartServer()
}

func aliblablaServer() {
	app := aliblabla.NewApplication()
	app.StartServer()
}

func StartMockServer() {
	wg := new(sync.WaitGroup)
	wg.Add(3) // Two servers

	go func() {
		bollieServer()
		wg.Done()
	}()

	go func() {
		coolbereServer()
		wg.Done()
	}()

	go func() {
		aliblablaServer()
		wg.Done()
	}()

	wg.Wait()
}
