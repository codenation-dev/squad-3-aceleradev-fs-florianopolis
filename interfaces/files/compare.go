package main
import(
	"os"
	"log"
	"io"
	"sort"
	"io/ioutil"
	"encoding/hex"
	"crypto/sha1"
	"path/filepath"
)

func getHashFromFile(filePath string) (string, error) {
	var stringHashSHA1 string
	file, erro := os.Open(filePath)
	if erro != nil{
		log.Println(erro.Error())
		return stringHashSHA1, erro
	}
	defer file.Close()
	hashSHA1 := sha1.New()
	_, erro = io.Copy(hashSHA1, file)
	if erro != nil{
		log.Println(erro.Error())
		return stringHashSHA1, erro
	}
	stringHashSHA1 = hex.EncodeToString(hashSHA1.Sum(nil))
	return stringHashSHA1, erro
}

func getLastTwoFiles(pathDir string) []string {
	var filesName []string
	files, erro := ioutil.ReadDir(pathDir)

	if erro != nil{
		log.Println(erro.Error())
	}
	sort.Slice(files, func(i,j int) bool{
    	return files[i].ModTime().Unix() > files[j].ModTime().Unix()
	})
	
	for i, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".zip" {
				if len(filesName) < 2{
					filesName = append(filesName, files[i].Name()) 					
				}else{
					break
				}
			}
		}
	}
    
	return filesName
}