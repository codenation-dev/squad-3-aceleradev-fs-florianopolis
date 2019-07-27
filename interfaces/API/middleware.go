package api

import ("net/http"
"encoding/json"
"github.com/gorilla/context"
"io/ioutil"
"bytes")

func (a *App) restricted(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		data,_ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var TempT tokenSt
		_ = json.Unmarshal(data,&TempT)  
		token,err :=  a.tokenVerify(TempT)
		if (err == nil) {
			if (token.Valid) {
				r.Body  = ioutil.NopCloser(bytes.NewBuffer(data))
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
