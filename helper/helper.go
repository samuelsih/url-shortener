package helper

import "net/http"

//MapHelper map any paths with their original URL
//return HandlerFunc with http.ResponseWriter and *http.Request
func MapHelper(urlContent map[string]string, fallBack http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		//if match to path, redirect to that path
		if destination, isOk := urlContent[path]; isOk {
			http.Redirect(rw, r, destination, http.StatusFound)
			return
		} 
		
		//if not match, fallback
		fallBack.ServeHTTP(rw, r)
	}
}