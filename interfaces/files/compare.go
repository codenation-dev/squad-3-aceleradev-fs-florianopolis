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
	"encoding/csv"
	"squad-3-aceleradev-fs-florianopolis/entities"
	"strconv"
	"encoding/json"
	"strings"
)


func openFileCSV() error {
	workPath, err := getFileName(); if err != nil {
		log.Println(err.Error())
		return err
	}
	fileName := getLastFiles(workPath.Directory, 1 ,".txt")
	fullPath := workPath.Directory + fileName[0]
	csvfile, err := os.Open(fullPath); if err != nil {
		log.Println(erro.Error())
		return err
	}
	defer csvfile.Close()
	
	reader := csv.NewReader(csvfile)
	reader.Comma = ';'
	rawdata, err := reader.ReadAll(); if err != nil {
		log.Println(err.Error())
		return err
	}
}

func insertIntoPessoa(rawdata [][]string ) error{
	if len(rawdata) > 0{
		for i, column := range rawdata {
			if (i > 0){			
				Remuneracaodomes, err := strconv.ParseFloat(changeComma(column[3]), 64); if err != nil {
					log.Println(err.Error())
					return err
				}
				Redutorsalarial, err := strconv.ParseFloat(changeComma(column[8]),64);if err != nil {
					log.Println(err.Error())
					return err
				}
				Totalliquido, err := strconv.ParseFloat(changeComma(column[9]),64);if err != nil {
					log.Println(err.Error())
					return err
				}
				if Totalliquido > 20000{
					Pessoa.Nome  = column[0]
					Pessoa.Cargo = column[1]
					Pessoa.Orgao = column[2]
					Pessoa.Remuneracaodomes = Remuneracaodomes
					Pessoa.RedutorSalarial = Redutorsalarial
					Pessoa.TotalLiquido = Totalliquido
					
					if true{//verifica se é cliente do banco na api e pega o ID da pessoa
						Pessoa.Update = true
						Pessoa.ClientedoBanco = true//pega o valor do json do cliente
						jsonData, err := json.Marshal(Pessoa);if err != nil {
							log.Println(err.Error())
							return err
						}
						log.Println(string(jsonData))
						//insere no banco
					}else{
						Pessoa.ClientedoBanco = false
						Pessoa.TotalLiquido = 0
						//atualiza no banco
					}

					/*
					- Caso salário líquido > 20k, 
					- Verifica se nome já é cliente
					- Caso cliente, update dos dados e update = true
					- Caso não cliente, insere os dados e seta update = true
					- Caso update = false, set valorliquido = 0
					- Ao final, setar novamente update = false em todos os clientes da tabela*/
				}

			}
			
		}
	}
}

func changeComma(pFloatNumber string) string{
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pFloatNumber,",",";"),".",":"),";","."),":",",")
}

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

func getLastFiles(pathDir string, countFiles int, extension string) []string {
	var filesName []string
	files, erro := ioutil.ReadDir(pathDir)

	if erro != nil{
		log.Println(erro.Error())
	}
	sort.Slice(files, func(i,j int) bool{
    	return files[i].ModTime().Unix() > files[j].ModTime().Unix()
	})
	if (countFiles > 0 && extension != ""){
		for i, file := range files {
			if file.Mode().IsRegular() {
				if filepath.Ext(file.Name()) == extension {
					if len(filesName) < countFiles{
						filesName = append(filesName, files[i].Name()) 					
					}else{
						break
					}
				}
			}
		}
	}
    
	return filesName
}