package helpers

import (
	"errors"
	"net/http"
)

func RecoverWrap(cb func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("terjadi kesalahan")
				}
				sendMeMail(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		cb(w, r)
	}
}

func sendMeMail(err error) {
	// send mail
}
