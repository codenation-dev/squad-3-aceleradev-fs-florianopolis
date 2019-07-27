package api

import ("net/http"
"encoding/json"
"github.com/gorilla/context")

func (a *App) restricted(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		decode := json.NewDecoder(r.Body)
		var TempT tokenSt
		_ = decode.Decode(&TempT)  
		token,err :=  a.tokenVerify(TempT)
		
		if (err == nil) {
			if (token.Valid) {
				context.Set(r,"token",token)
				next(w,r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				responseCodeResult(w, Error, "Invalid Token")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			responseCodeResult(w,Error,"Unathorized Request")
		}
}
}
