package api

import ("net/http"
"encoding/json")

func (a *App) restricted(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		decode := json.NewDecoder(r.Body)
		var TempT tokenSt
		_ = decode.Decode(&TempT)  
		token,err :=  a.tokenVerify(TempT)//
		
		if (err == nil) {
			if (token.Valid) {
				next(w,r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				unauth(w,r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			internalError(w,r)
		}
}
}