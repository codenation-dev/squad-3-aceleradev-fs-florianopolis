package main

import (
	"os"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	//"squad-3-aceleradev-fs-florianopolis/interfaces/files"
)

const (
	//URLService is the address to download the zip file
	URLService = "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx"
)

func main() {
	DownloadAndExtractFile()
	openFileCSV()
	CreateJSONfile()
}

//DownloadAndExtractFile from URLService
func DownloadAndExtractFile() {
	workPath, erro := getFileName()
	zipFileName := workPath.FullPath
	if erro == nil {
		wasDownload, _ := DownloadFile(URLService, zipFileName)
		if erro != nil {
			logs.Errorf("DownloadAndExtractFile", erro.Error())
		}
		filesName := getLastFiles(workPath.Directory, 2, ".zip")
		if wasDownload {
			switch {
			case len(filesName) < 2:
				ExtractFile(zipFileName)
			case len(filesName) == 2:
				hashNewFile, erro := getHashFromFile(workPath.Directory + filesName[0])
				if erro != nil {
					logs.Errorf("DownloadAndExtractFile", erro.Error())
				}
				hashExistFile, erro := getHashFromFile(workPath.Directory + filesName[1])
				if erro != nil {
					logs.Errorf("DownloadAndExtractFile", erro.Error())
				}
				if hashExistFile == hashNewFile {
					erro = os.Remove(workPath.Directory + filesName[0])
					if erro != nil {
						logs.Errorf("DownloadAndExtractFile", erro.Error())
					}
					logs.Info("DownloadAndExtractFile", "Files are the same. New file was removed.")
				} else {
					ExtractFile(zipFileName)
				}
			}
		} else {
			logs.Info("DownloadAndExtractFile", "New file was not downloaded. File was not extract.")
		}
	} else {
		logs.Errorf("DownloadAndExtractFile", erro.Error())
	}
}
