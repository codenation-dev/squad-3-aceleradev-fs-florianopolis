package main
import (
	"time"
	"os"
	"path/filepath"
	"net/http"
	"bytes"
	"io"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
)

//WorkPath store the files path
type WorkPath struct{
	Directory string
	FileName string
	FullPath string
}

func getFileName() (*WorkPath, error)  {
	wp := new(WorkPath)
	currentTime := time.Now()
	dir, err := os.Getwd(); if err != nil{
		logs.Errorf("getFileName", "Error when tried to get directory: " + err.Error())
		return wp, err
	}
	formatTime := currentTime.Format("01-02-2006")
	dir += "/download/"
	wp.Directory = dir
	fileName := "Transparencia-" + formatTime + ".zip"
	wp.FileName = fileName
	wp.FullPath = filepath.Join(dir, fileName)
	return wp, err
}

// DownloadFile from url
func DownloadFile(URLFile, fileName string) (bool, error ){
	wasDownload := false
	logs.Info("DownloadFile", "Start Download File")
	_, erro := os.Stat(fileName)
	if os.IsNotExist(erro){
		client := http.Client{}
		requestBody := []byte("__EVENTTARGET=&__EVENTARGUMENT=&__LASTFOCUS=&__VIEWSTATE=%2FwEPDwULLTIwNDQzOTAyMzEPZBYCAgMPZBYMAgUPEA8WBh4ORGF0YVZhbHVlRmllbGQFCE9SR0FPX0lEHg1EYXRhVGV4dEZpZWxkBQpPUkdBT19ERVNDHgtfIURhdGFCb3VuZGdkEBVbBVRPRE9THUFETUlOSVNUUkFDQU8gR0VSQUwgRE8gRVNUQURPHUFHLk1FVFIuU09ST0NBQkEtQUdFTVNPUk9DQUJBKEFHRU5DSUEgTUVUUk9QT0xJVEFOQSBCQUlYQURBIFNBTlRJU1RBIEEoQUdFTkNJQSBNRVRST1BPTElUQU5BIERFIENBTVBJTkFTIC0gQUdFTShBR0VOQ0lBIE1FVFJPUE9MSVRBTkEgREUgQ0FNUElOQVMgQUdFTUNBKEFHRU5DSUEgTUVUUk9QT0xJVEFOQSBETyBWQUxFIERPIFBBUkFJQkEoQUdFTkNJQSBSRUdVTEFET1JBIFNBTkVBTUVOVE8gRU5FUkdJQSBFUyhBR0VOQ0lBIFJFR1VMQURPUkEgU0VSVi5QVUJMLkRFTC5UUkFOU1BPJENBSVhBIEJFTkVGSUNFTlRFIERBIFBPTElDSUEgTUlMSVRBUgpDQVNBIENJVklMJkNFTlRSTyBFU1QuREUgRURVQy5URUNOT0wuIFBBVUxBIFNPVVpBJ0NFVEVTQi1DT01QQU5ISUEgQU1CSUVOVEFMIERPIEVTVEFETyBERSNDSUEuIFBST0NFUy4gREFET1MgRVNULlMuUC4tUFJPREVTUCdDSUEuREVTRU5WLkhBQklULkUgVVJCQU5PIEVTVC5TLlAuLUNESFUoQ0lBLlNBTkVBTUVOVE8gQkFTSUNPIEVTVC5TLlBBVUxPLVNBQkVTUChDT01QQU5ISUEgREUgREVTRU5WT0xWSU1FTlRPIEFHUklDT0xBIERFJENPTVBBTkhJQSBERSBTRUdVUk9TIEVTVC5TLlAuLUNPU0VTUCBDT01QQU5ISUEgRE8gTUVUUk9QT0xJVEFOTy1NRVRSTyBDT01QQU5ISUEgRE9DQVMgREUgU0FPIFNFQkFTVElBTyhDT01QQU5ISUEgUEFVTC5UUkVOUyBNRVRST1BPTElUQU5PUy1DUFRNKENPTVBBTkhJQSBQQVVMSVNUQSBERSBPQlJBUyBFIFNFUlZJQ09TIC0lQ09NUEFOSElBIFBBVUxJU1RBIERFIFBBUkNFUklBUyAtIENQUCZDT01QQU5ISUEgUEFVTElTVEEgU0VDVVJJVElaQUNBTy1DUFNFQyZERVBBUlRBTUVOVE8gQUVST1ZJQVJJTyBFU1QuUy5QLi1EQUVTUChERVBBUlRBTUVOVE8gQUdVQVMgRU5FUkdJQSBFTEVUUklDQS1EQUVFJ0RFUEFSVEFNRU5UTyBERSBFU1RSQURBUyBERSBST0RBR0VNLURFUihERVBBUlRBTUVOVE8gRVNUQURVQUwgREUgVFJBTlNJVE8tREVUUkFOKERFUEFSVEFNRU5UTyBFU1RBRFVBTCBUUkFOU0lUTyBERVRSQU4gU1AoREVTRU5WT0xWRSBTUCBBR0VOQ0lBIERFIEZPTUVOVE8gRE8gRVNUQSdERVNFTlZPTFZJTUVOVE8gUk9ET1ZJQVJJTyBTLkEuIC0gREVSU0EoRU1BRS1FTVBSRVNBIE1FVFJPUE9MSVRBTkEgREUgQUdVQVMgRSBFTiZFTVAuTUVUUi5UUkFOU1AuVVJCQU5PUyBTUC5TL0EgRU1UVS1TUChFTVAuUEFVTElTVEEgUExBTkVKLk1FVFJPUExJVEFOTyBTLkEtRU1QIEZBQ1VMREFERSBERSBNRURJQ0lOQSBERSBNQVJJTElBKEZBQ1VMREFERSBNRURJQ0lOQSBTQU8gSk9TRSBETyBSSU8gUFJFVE8oRlVORC5DLkFURU5ELlNPQy1FRC5BRE9MRVNDLUZVTkQuQ0FTQS1TUCdGVU5ELlBFLkFOQ0hJRVRBLUMuUC5SQURJTyBUVi5FRFVDQVRJVkEmRlVOREFDQU8gQU1QQVJPIFBFU1FVSVNBIEVTVCBTQU8gUEFVTE8oRlVOREFDQU8gQ09OU0VSVi5QUk9ELkZMT1JFU1RBTCBFU1QuUy5QLihGVU5EQUNBTyBERVNFTlZPTFZJTUVOVE8gREEgRURVQ0FDQU8tRkRFKEZVTkRBQ0FPIElOU1RJVFVUTyBERSBURVJSQVMgRE8gRVNULlMuUC4jRlVOREFDQU8gTUVNT1JJQUwgREEgQU1FUklDQSBMQVRJTkEgRlVOREFDQU8gT05DT0NFTlRSTyBERSBTQU8gUEFVTE8kRlVOREFDQU8gUEFSQSBPIFJFTUVESU8gUE9QVUxBUi1GVVJQJkZVTkRBQ0FPIFBBUlFVRSBaT09MT0dJQ08gREUgU0FPIFBBVUxPKEZVTkRBQ0FPIFBSRVZJREVOQ0lBIENPTVBMRU1FTlRBUiBFU1RBRE8oRlVOREFDQU8gUFJPLVNBTkdVRS1IRU1PQ0VOVFJPIFNBTyBQQVVMTyhGVU5EQUNBTyBQUk9GLkQuTUFOT0VMIFAuUElNRU5URUwgLUZVTkFQKEZVTkRBQ0FPIFBST1RFQy5ERUZFU0EgQ09OU1VNSURPUi1QUk9DT04oRlVOREFDQU8gU0lTVEVNQSBFU1QuQU5BTElTRSBEQURPUy1TRUFERSdGVU5EQUNBTyBVTklWRVJTSURBREUgVklSVFVBTCBETyBFU1RBRE8WR0FCSU5FVEUgRE8gR09WRVJOQURPUihIT1NQLkNMSU5JQ0FTIEZBQy5NRURJQ0lOQSBSSUIuUFJFVE8tVVNQKEhPU1BJVEFMIENMSU5JQ0FTIEZBQ1VMREFERSBNRURJQ0lOQS1VU1AmSE9TUElUQUwgREFTIENMSU5JQ0FTIEZBQy5NRUQuQk9UVUNBVFUnSU1QUkVOU0EgT0ZJQ0lBTCBETyBFU1RBRE8gUy5BLiAgIElNRVNQJklOU1QgTUVESUNJTkEgU09DIEUgQ1JJTUlOIFMuUC4gLUlNRVNDKElOU1QuUEVTT1MgTUVESURBUyBFU1QuU0FPIFBBVUxPLUlQRU0vU1AmSU5TVC5QRVNRVS5URUNOT0xPR0lDQVMgRVNULlNQLVMvQS1JUFQoSU5TVElUVVRPIEFTU0lTVEVOQ0lBIE1FRElDQSBBTyBTRVJWSURPUihJTlNUSVRVVE8gREUgQVNTSVNURU5DSUEgTUVESUNBIFNFUlZJRE9SKElOU1RJVFVUTyBQQUdBTUVOVE9TIEVTUEVDSUFJUyBTUC0gSVBFU1AoSlVOVEEgQ09NRVJDSUFMIERPIEVTVEFETyBERSBTQU8gUEFVTE8tSiBQT0xJQ0lBIE1JTElUQVIgRVNUQURPIFNBTyBQQVVMTxxQUk9DVVJBRE9SSUEgR0VSQUwgRE8gRVNUQURPHlNBTyBQQVVMTyBQUkVWSURFTkNJQSAtIFNQUFJFViZTRUNSRVRBUklBIEFETUlOSVNUUkFDQU8gUEVOSVRFTkNJQVJJQShTRUNSRVRBUklBIERBIENVTFRVUkEgRSBFQ09OT01JQSBDUklBVElWFlNFQ1JFVEFSSUEgREEgRURVQ0FDQU8kU0VDUkVUQVJJQSBEQSBGQVpFTkRBIEUgUExBTkVKQU1FTlRPF1NFQ1JFVEFSSUEgREEgSEFCSVRBQ0FPE1NFQ1JFVEFSSUEgREEgU0FVREUfU0VDUkVUQVJJQSBEQSBTRUdVUkFOQ0EgUFVCTElDQShTRUNSRVRBUklBIERFIEFHUklDVUxUVVJBIEUgQUJBU1RFQ0lNRU5UJ1NFQ1JFVEFSSUEgREUgREVTRU5WT0xWSU1FTlRPIEVDT05PTUlDTyZTRUNSRVRBUklBIERFIERFU0VOVk9MVklNRU5UTyBSRUdJT05BTCRTRUNSRVRBUklBIERFIERFU0VOVk9MVklNRU5UTyBTT0NJQUwhU0VDUkVUQVJJQSBERSBFTkVSR0lBIEUgTUlORVJBQ0FPFlNFQ1JFVEFSSUEgREUgRVNQT1JURVMVU0VDUkVUQVJJQSBERSBHT1ZFUk5PKFNFQ1JFVEFSSUEgREUgSU5GUkFFU1RSVVRVUkEgRSBNRUlPIEFNQkklU0VDUkVUQVJJQSBERSBMT0dJU1RJQ0EgRSBUUkFOU1BPUlRFUyhTRUNSRVRBUklBIERFIFNBTkVBTUVOVE8gRSBSRUNVUlNPUyBISURSFVNFQ1JFVEFSSUEgREUgVFVSSVNNTyZTRUNSRVRBUklBIEVNUFJFR08gRSBSRUxBQ09FUyBUUkFCQUxITyhTRUNSRVRBUklBIEVTVC5ESVJFSVRPUyBQRVNTT0EgREVGSUNJRU5DHlNFQ1JFVEFSSUEgSlVTVElDQSBFIENJREFEQU5JQSVTRUNSRVRBUklBIFRSQU5TUE9SVEVTIE1FVFJPUE9MSVRBTk9TIFNQLkNMSU4uRkFDLk1FRC5NQVJJTElBLUhDRkFNRU1BKFNVUEVSSU5URU5ERU5DSUEgREUgQ09OVFJPTEUgREUgRU5ERU1JQVMVWwItMQExATIBMwE0ATUBNgE3ATgBOQIxMAIxMQIxMgIxMwIxNAIxNQIxNgIxNwIxOAIxOQIyMAIyMQIyMgIyMwIyNAIyNQIyNgIyNwIyOAIyOQIzMAIzMQIzMgIzMwIzNAIzNQIzNgIzNwIzOAIzOQI0MAI0MQI0MgI0MwI0NAI0NQI0NgI0NwI0OAI0OQI1MAI1MQI1MgI1MwI1NAI1NQI1NgI1NwI1OAI1OQI2MAI2MQI2MgI2MwI2NAI2NQI2NgI2NwI2OAI2OQI3MAI3MQI3MgI3MwI3NAI3NQI3NgI3NwI3OAI3OQI4MAI4MQI4MgI4MwI4NAI4NQI4NgI4NwI4OAI4OQI5MBQrA1tnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnZ2dnFgFmZAIHDxBkEBUBBVRPRE9TFQECLTEUKwMBZxYBZmQCCQ8QDxYGHwAFC1NJVFVBQ0FPX0lEHwEFDVNJVFVBQ0FPX0RFU0MfAmdkEBUEBVRPRE9TC0FQT1NFTlRBRE9TBkFUSVZPUwxQRU5TSU9OSVNUQVMVBAItMQExATIBMxQrAwRnZ2dnFgFmZAILDw9kFgIeCm9uS2V5UHJlc3MFJ3JldHVybiBNYXNjYXJhTW9lZGEodGhpcywnLicsJywnLGV2ZW50KWQCDQ8PZBYCHwMFJ3JldHVybiBNYXNjYXJhTW9lZGEodGhpcywnLicsJywnLGV2ZW50KWQCFQ9kFgJmD2QWBAIBDxYCHgdWaXNpYmxlaGQCAw8PFgIfBGhkFgICAw88KwARAgEQFgAWABYADBQrAABkGAIFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYBBQxpbWdFeHBvcnRUeHQFBGdyaWQPZ2QU3O1EpUneDTDR0La3eya%2BksNURyHibL%2FtP%2BcgqSjJ9Q%3D%3D&__VIEWSTATEGENERATOR=E42B1F40&__EVENTVALIDATION=%2FwEdAG6cMAHUcsWpjV%2B00oDIGub5ha8fMqpdVfgdiIcywQp19AS0oC9%2BkRn5wokBQj%2BYmSdj%2FRE4%2FVY2xVooDbyNylWSFXsupcqZ9EYohXUHrvyuvszqcPgWZLCNPbx1As5K6XI8YfiXwzc6jdd6doCEWNMhfUq2YkY3rbVwieJI30sGRBiYwU43rbtypsxax6Lexvr9tn%2FppXosAOoaLiPglbLZDQ4AHCggkRiV1y9R5Jk3hxzIBiDVeBd4ex%2FDPERS7Y3hxS83fVJEzO6I%2BsKPdRPTZbKZKzZ%2FiI%2Fo2LERffiPWbY0qpjFHBt23vPUuehVkAOA1ngNB93rbK%2Bu0E54XcLAmWLN%2Fl%2Bz5m0ApRDNS4L3FwTfILDr1aT4Crd1%2F2X2tGTSlHv5v4gI%2B%2F4UxQdVOOXcJIWT3hhEHPLkfTczdhS%2BJPFzCLQyhLlM%2FTIkVLdCEWiXz8XDG1%2BqV0wHjm1sFCkHt5aLy6yjxTyv1FFML9B%2Fo0JBJO%2By%2B74vfDQlvwQWQHtswD%2Bjri2Ja0FbYTVaHetzL3nIpMtKnzHrJejZWNnngPadPS2744kvbqzTJQaAdqOeYy%2FXyO581zGaQB16a5HkpT5jddxT22MOtOJS9%2BOuUHRXp8dj268DwFDqeWohT0vm1b0FOlCVjyi8V9MKHPYPpHgZ%2F2GzcT5zaEXX3Wa7dGMCaXmo3KMrfSTIEMtzpixzPEyfillVBjlMq8fiaJmavKW63uZc65AHMJEgzJBWOOnY33pftn93IOwZzZWV8DBA7v%2F9aPpqFJWx65SrmQqSjTKR9Q8znWzwmOcZE4%2FSuTP7i%2BXb7NoOWr4anBMJ9L8iQIpPyUdRVhTh0dqpW9mg677VkTJzeFDr78YgZsAwP%2FX%2BdTV%2FINjSEi5I3GKGi7myZ7%2BjeKd7PDtAjn8O4hLTJfL4LFg4Nvwdmd%2F53R8Jw4b9e%2FlLobx4zXIq3GAuywAjOQvHY8AEnfNd%2FlXdKYxyzc%2FwfpCNJupjNVpUse2VJD4oS1BuBPCBdQ5aaErF4JFlItPtLQCYFzs0jfHra3vGXa5DUmVxUHX61STePVHIx%2Bb2IzWzaVJbMWnr0ySeyyy%2FZ1AEi%2FGyAY4VRi7gupaG4KIpRnL0PqiHkB0m%2BFOAGOzlYyAzkRO1hwDnOQf3fkyzTk8GPsW4ORs6zPd%2BeDosaOUhW1MEtWA%2BSqsohtmqkoKbjumKVbQvus3TM3adBbzpeRPEjnLNywu7OwRAhFtyU0gmtXU9am1kuUbvzTaW93G%2FXW5pJhxIEGLJ46ijUCocW5ypp1AUfwUVaLtxxktia9eKFUCg16rKs9CfE8mQS1sJL8sXrl1kCYgl357rWaG95jfZ509s%2Bm2fA%2BOt0aP8OyaOU4R1ht8FAaoUaukJi9ac%2B52YAhiIATqgCuAVAUaz6iVZ30v9i3l79pG%2FQjT0yzItrPhgpeaj5FDDRNwFWQfE5v7dhuWXa0fqNuT0%2F3rHd8yAI%2FR31smXtVMpuDg4uNPHIl%2B2FxKOozxg%2Fv%2B%2BE9d%2FZoPPgEhC0wqwEcy5cuqQMsS7I2iwe1Xfp9TBV2uBNFpR3V1ws1NcSb0O892YPaDPsxrja2GQM7SzAShZDNlCOSW7Tt%2Fu0g%2BeirEQ%2FlwLvd%2FyO3h%2FPXkp4oZAfoeCSWuKxs7UkSXX7piPjdZRkxS8%2B1Tv52TtsW%2F%2FarETeAIdqgWD21SCG%2F%2BSG%2FyFJtRwUalOOSCKwgXmjHLagrrOpyOVvrzcda9t4I8AvfZJNBX4HCyHl%2F8v7zlaXsN6v3xdx7SBYcgTu1GewkDpUJSUGbiJpTFb9FwFesoo5ATV8LN38tAuINPU8rfSikTUmdlp8CARYKFn95WsBdjs1x8c6lK59jnQ%2FQHi2nKDMKfdQRVhcvnFwvt6SokCFQDX7AEtmU9OC%2Fkwe5SIcBU04jVZdwLiKogB2pPql%2FnA4CHA7mEf3AIr0wLOnRAQ0xjhC3PXHrIjjpV2suu3zMJ7LscXSxIToHr95TxJTzSEj9C7XyN%2FGMISH%2FTKb%2FPRxrbwGTEZF3x922wvTvFKuuxNUJFB79U3ZPxLws5iIazIlee0zV3InWYYPP26JIa5R0Em8ORb%2B%2FoUDlJKcdv6NoWV%2F5WtCyREa2Rxke5ZukLmT7xiWinv8jrwbnAz1AUaMm8xKsc4G6dNWu2jHrgAaNFlmOLZIeG0OTsyPhh%2B%2F0WQdOTAD9zAblcx6VvMEe43r2g9sGn75bO7ZW6nZ7hGBjKUqSH4S7Qy5ngR%2FiduIfdzD0oNgNO6zlZmgx%2BPVHfpxvG%2B1lXBZBLAe6JyY9%2FwY3j6%2BMGuruxn5MX0jsPeyBXK401Kwjl8g4KbJ6y3JnlYwpVFE%2BxaAmkM3Axtd6lTTj6rtggdQ6BBteBDf6J8rxtPlESdxIsw&txtNome=&orgao=-1&cargo=-1&situacao=-1&txtDe=&txtAte=&hdInicio=&hdFinal=&hdPaginaAtual=&hdTotal=&imgExportTxt.x=18&imgExportTxt.y=20")
		request, erro := http.NewRequest("POST", URLFile, bytes.NewBuffer(requestBody))
		request.Header.Set("Host","www.transparencia.sp.gov.br")
		request.Header.Set("User-Agent","Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
		request.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Set("Accept-Language","en-US,en;q=0.5")
		request.Header.Set("Accept-Encoding","gzip, deflate")
		request.Header.Set("Content-Type","application/x-www-form-urlencoded")
		request.Header.Set("Content-Length","8441")
		request.Header.Set("Connection","keep-alive")
		request.Header.Set("Referer","http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx")
		request.Header.Set("Upgrade-Insecure-Requests","1")
		logs.Info("DownloadFile", "POST Request was sent...")
		response, erro := client.Do(request);if erro != nil{
			logs.Errorf("DownloadFile", "Error when tried to get Response: " + erro.Error())
			return wasDownload, erro
		}
		logs.Info("DownloadFile", "Response from POST Request was receveid")
		defer response.Body.Close()
		file, erro := os.Create(fileName);if erro != nil{
			logs.Errorf("DownloadFile", "Error when tried to create file: " + erro.Error())
			return wasDownload, erro
		}
		logs.Info("DownloadFile", "Create new file into: "+ fileName)

		defer file.Close()
		_, erro = io.Copy(file, response.Body);	if erro != nil{
			logs.Errorf("DownloadFile", erro.Error())
			return wasDownload, erro
		}
		logs.Info("DownloadFile", "Writing Response File into New Empty File")
		wasDownload = true
	}else{
		logs.Info("DownloadFile", "The file already exists. Nothing was done")
	}
	return wasDownload, nil

}