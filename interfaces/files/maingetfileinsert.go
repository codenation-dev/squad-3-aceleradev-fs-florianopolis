package main

import (
	"os"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"sync"

	"github.com/robfig/cron"
)

const (
	//URLService is the address to download the zip file
	URLService = "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx"
)

func main() {
	if true{
		logs.Info("Start App", "The application was Started")
		wg := &sync.WaitGroup{}
		wg.Add(1)
		cronJob := cron.New()
		cronJob.Start()
		logs.Info("Start App", "Application is Waiting until the time match...")
		cronJob.AddFunc("30 00 22 10 * ?", Execute) //dia 10 de cada mes as 22:00:30
		cronJob.AddFunc("30 00 22 25 * ?", Execute) //dia 25 de cada mes as 22:00:30
		wg.Wait()
		Execute()
	}else{
		Execute()
	}
}

//Execute when the time is match
func Execute() {
	if DownloadAndExtractFile() {
		OpenAndProcessFileCSV()
		CreateNotify()
	}
	logs.Info("Execute", "Finished running Execute")
}

//DownloadAndExtractFile from URLService
func DownloadAndExtractFile() bool {
	process := false
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
				process = true
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
					process = true
				}
			}
		} else {
			logs.Info("DownloadAndExtractFile", "New file was not downloaded. File was not extract.")
		}
	} else {
		logs.Errorf("DownloadAndExtractFile", erro.Error())
	}
	return process
}
