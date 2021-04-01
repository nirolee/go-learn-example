package main

import (
	"gin/errhandiling/filelistingserver/filelisting"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				http.Error(writer,
					http.StatusText(http.StatusNotFound),
					http.StatusNotFound)
			case os.IsPermission(err):
				http.Error(writer,
					http.StatusText(http.StatusForbidden),
					http.StatusForbidden)
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", filelisting.HandleFileList)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
