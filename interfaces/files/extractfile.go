package main
import(
	"archive/zip"
	"log"
	"os"
	"path/filepath"
	"io"
)

// ExtractFile from zip path
func ExtractFile(pathFile string)  {
	reader, erro := zip.OpenReader(pathFile)
	log.Println("Open reader to file:",pathFile)
	if erro != nil{
		log.Println("Error to OpenReader:", erro.Error());
	}
	defer reader.Close()
	for _, file := range reader.Reader.File{
		log.Println("Open zipfile...")
		zipfile, erro := file.Open()
		if erro != nil{
			log.Println("Error to open zipfile:", erro.Error())
		}
		defer zipfile.Close()
		dir, erro := os.Getwd()
		if erro != nil{
			log.Println("Error when tried to get directory:", erro.Error())
		}
		dir = dir + "/download/"
		// get the individual file name and extract
		path := filepath.Join(dir, file.Name)
		log.Println("Files will be extract to...", path)
		if file.FileInfo().IsDir(){
			log.Println("Creating directory", path)
			erro := os.MkdirAll(path, file.Mode())
			if erro != nil{
				log.Println("Error when tried to create directory" , erro.Error())
			}
			
		}else{
			log.Println("Extracting files, please wait...")
			write, erro := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, file.Mode())
			if erro != nil{
				log.Println(erro.Error())
			}
			_, erro = io.Copy(write, zipfile)
			log.Println("Extract complete!")
			if erro != nil{
				log.Println(erro.Error())
			}
		}
	}

}