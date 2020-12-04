package bollie

import (
	"encoding/json"
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
		errorLog: log.New(os.Stdout, "BOLLIE_ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "BOLLIE_INFO:\t", log.Ldate|log.Ltime),
		products: &inMemoryProductRepository{Products: map[int]product{
			1: {
				ID:    1,
				Name:  "Legeau",
				Price: 1.33,
			},
			2: {
				ID:    2,
				Name:  "Vroommobiel",
				Price: 23.59,
			},
			3: {
				ID:    3,
				Name:  "Iksboks Wan",
				Price: 299.95,
			},
			4: {
				ID:    4,
				Name:  "PleeSteeSjon",
				Price: 300.01,
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
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(errResp)
	return
}
