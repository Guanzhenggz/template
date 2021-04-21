package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var tests = [] struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func errPanic(_ http.ResponseWriter,_ *http.Request) error{
	panic(123)
}

type testinguserError string

func (e testinguserError) Error() string {
	return e.Message()
}

func (e testinguserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,request *http.Request) error {
	return testinguserError("user error")
}

func errNotFound(_ http.ResponseWriter,_ *http.Request) error{
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter,_ *http.Request) error{
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter,_ *http.Request) error{
	return errors.New("Unknown error")
}

func noError(writer http.ResponseWriter,request *http.Request) error{
	fmt.Fprintln(writer,"no error")
	return nil
}

func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet,
			"http://www.baidu.com", nil)
		f(recorder, request)
		b, _ := ioutil.ReadAll(recorder.Body)
		if tt.message != strings.Trim(string(b), "\n") || tt.code != recorder.Code {
			t.Errorf("expect (%d , %s); "+
				"actual (%d , %s);", tt.code, tt.message, recorder.Code, recorder.Body)
		}
	}
}

	func TestErrWrapperInServer(t *testing.T) {
		for _,tt := range tests {
			f := errWrapper(tt.h)
			server := httptest.NewServer(http.HandlerFunc(f))
			resp, _ := http.Get(server.URL)
			b, _ := ioutil.ReadAll(resp.Body)
			if tt.message != strings.Trim(string(b), "\n") || tt.code != resp.StatusCode {
				t.Errorf("expect (%d , %s); "+
					"actual (%d , %s);", tt.code, tt.message, resp.StatusCode, resp.Body)
			}

		}
	}

