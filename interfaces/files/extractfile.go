package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
)

// ExtractFile from zip path
func ExtractFile(pathFile string) {
	reader, erro := zip.OpenReader(pathFile)
	if erro != nil {
		logs.Errorf("ExtractFile", "Error to OpenReader: "+erro.Error())
	}
	logs.Info("ExtractFile", "Open reader to file: "+pathFile)
	defer reader.Close()
	for _, file := range reader.Reader.File {
		logs.Info("ExtractFile", "Open zipfile...")
		zipfile, erro := file.Open()
		if erro != nil {
			logs.Errorf("ExtractFile", "Error to open zipfile: "+erro.Error())
		}
		defer zipfile.Close()
		dir, erro := os.Getwd()
		if erro != nil {
			logs.Errorf("ExtractFile", "Error when tried to get directory: "+erro.Error())
		}
		dir = dir + "/download/"
		// get the individual file name and extract
		path := filepath.Join(dir, file.Name)
		logs.Info("ExtractFile", "Files will be extract to: "+path)
		if file.FileInfo().IsDir() {
			logs.Info("ExtractFile", "Creating directory: "+path)
			erro := os.MkdirAll(path, file.Mode())
			if erro != nil {
				logs.Errorf("ExtractFile", "Error when tried to create directory"+erro.Error())
			}

		} else {
			logs.Info("ExtractFile", "Extracting files, please wait...")
			write, erro := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, file.Mode())
			if erro != nil {
				logs.Errorf("ExtractFile", erro.Error())
			}
			_, erro = io.Copy(write, zipfile)
			if erro != nil {
				logs.Errorf("ExtractFile", erro.Error())
			}
			logs.Info("ExtractFile", "Extract complete!")

		}
	}

}
