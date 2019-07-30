package api


import ("net/http"
"github.com/gorilla/context"
"fmt"
"squad-3-aceleradev-fs-florianopolis/entities/logs")


func (a *App) restricted(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				logs.Info("REQUEST_HEADER",fmt.Sprintf("%s :: %s",name,value))
			}
		}
		ut := r.Header.Get("Access-Token")
		TempT := tokenSt{Token:ut}
		token,err := a.tokenVerify(TempT)
		if (err == nil) {
			if (token.Valid) {
				context.Set(r,"token",token)
				next(w,r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				responseCodeResult(w, Unauth, "Invalid Token")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			responseCodeResult(w, Unauth,"Unathorized Request")
		}
}
}
