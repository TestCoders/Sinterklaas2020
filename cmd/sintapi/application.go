package main

import (
	"log"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	sources  map[string]SintClient
	cache    cache
}

func NewApplication() *Application {
	return &Application{
		errorLog: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		sources: map[string]SintClient{
			"bollie":    NewBollieClient(nil, "http://mockserver:3001"),
			"coolbere":  NewCoolbereClient(nil, "http://mockserver:3002"),
			"aliblabla": NewAliblablaClient(nil, "http://mockserver:3003"),
		},
		cache: cache{products: map[int][]Product{}},
	}
}
