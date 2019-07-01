package main

import (
	"log"
	"os"
)
const(
	URLService is the address to download the zip file
	URLService = "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx"
)

func main()  {
	//DownloadAndExtractFile()
	openFileCSV()
}

//DownloadAndExtractFile from URLService
func DownloadAndExtractFile()  {
	workPath, erro := getFileName()
	zipFileName := workPath.FullPath
	if erro == nil{
		wasDownload, _ := DownloadFile(URLService, zipFileName)
		if erro != nil{
			log.Println(erro.Error())
		}
		filesName := getLastFiles(workPath.Directory, 2 ,".zip")
		if wasDownload{
			switch {
			case len(filesName) < 2:
				ExtractFile(zipFileName)
			case len(filesName) == 2:
				hashNewFile, erro := getHashFromFile(workPath.Directory + filesName[0])
				if erro != nil{
					log.Println(erro.Error())
				}
				hashExistFile, erro := getHashFromFile(workPath.Directory + filesName[1])				
				if erro != nil{
					log.Println(erro.Error())
				}
				if hashExistFile == hashNewFile{
					erro = os.Remove(workPath.Directory + filesName[0])
					if erro != nil{
						log.Println(erro.Error())
					}
					log.Println("Files are the same. New file was removed.")
				}else{
					ExtractFile(zipFileName)
				}
			}
		}else {
			log.Println("New file was not downloaded. File was not extract.")
		}
	}else{
		log.Println(erro.Error())
	}
}