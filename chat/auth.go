package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		//NOT authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		//ERROR
		panic(err.Error())
	} else {
		//SUCCESS
		h.next.ServeHTTP(w, r)
	}
 }

 func MustAuth(handler http.Handler) http.Handler {
	 return &authHandler{next: handler}
 }

 //handles third party login process
 //format: /auth/{action}/{third-party provider}
 func loginHandler(w http.ResponseWriter, r *http.Request) {
	 segs := strings.Split(r.URL.Path, "/")
	 action := segs[2]
	 provider := segs[3]
	 switch action {
	 case "login":
		 log.Println("TODO handle login for", provider)
	 default:
		 w.WriteHeader(http.StatusNotFound)
		 fmt.Fprintf(w, "Auth action %s not supported", action)
	 }
 }
