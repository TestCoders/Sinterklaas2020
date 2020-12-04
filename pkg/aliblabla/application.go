package aliblabla

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	products IProductRepository
}

func (a *Application) Error(err error) {
	if err == nil {
		return
	}

	if e := a.errorLog.Output(1, err.Error()); e != nil {
		a.errorLog.Println(err.Error())
	}
}

func NewApplication() *Application {
	return &Application{
		errorLog: log.New(os.Stdout, "ALIBLABLA_ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "ALIBLABLA_INFO:\t", log.Ldate|log.Ltime),
		products: &inMemoryProductRepository{Products: map[int]product{
			1: {
				ID:    1,
				Name:  "Legeau",
				Price: 0.15,
			},
			2: {
				ID:    2,
				Name:  "Vroommobiel",
				Price: 22.00,
			},
			3: {
				ID:    3,
				Name:  "Iksboks Wan",
				Price: 310.13,
			},
			4: {
				ID:    4,
				Name:  "PleeSteeSjon",
				Price: 299.99,
			},
			5: {
				ID:    5,
				Name:  "Playdebiel",
				Price: 4.89,
			},
		}},
	}
}

func (a *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *Application) writeError(w http.ResponseWriter, status int, err error) {
	a.Error(err)
	errResp := &ResponseError{
		Error:       strconv.Itoa(status),
		Description: err.Error(),
	}
	w.WriteHeader(status)
	e := xml.NewEncoder(w)
	e.Indent("", "    ")
	e.Encode(errResp)
	return
}
