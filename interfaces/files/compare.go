package main

import (
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/funcpublico"
	"strconv"
	"strings"
)

type Clients struct {
	Nome string `json:"nome"`
}

func openFileCSV() error {
	workPath, err := getFileName()
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return err
	}
	fileName := getLastFiles(workPath.Directory, 1, ".txt")
	fullPath := workPath.Directory + fileName[0]
	csvfile, err := os.Open(fullPath)
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return err
	}
	defer csvfile.Close()
	logs.Info("openFileCSV", "Opening file: "+fullPath)
	reader := csv.NewReader(csvfile)
	reader.Comma = ';'
	logs.Info("openFileCSV", "Reading file...")
	rawdata, err := reader.ReadAll()
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return err
	}
	//dbi, _ := db.Init()
	//defer dbi.Database.Close()
	//err = insertIntoPessoa(rawdata, dbi)
	err = insertIntoPessoa(rawdata)
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return err
	}

	return err
}

//Check if person already exists in DB (by name)
func checkPersonInDB(name string) (bool, int) {
	Pessoa := new(entity.FuncPublico)
	alreadyInDB := false
	Pessoa, _ = funcpublico.GetByName(name)
	if Pessoa.Nome == name {
		alreadyInDB = true
	}
	return alreadyInDB, Pessoa.ID
}

//func to check if its a client
func isClient(name string) bool {
	isClient := false
	file, erro := ioutil.ReadFile("../api/clientlist-alterado.json")
	if erro != nil {
		logs.Errorf("isClient", erro.Error())
	}
	data := []Clients{}

	erro = json.Unmarshal(file, &data)
	if erro != nil {
		logs.Errorf("isClient", erro.Error())
	}

	for _, value := range data {
		if name == value.Nome {
			//fmt.Println(value.Nome + "=" + name)
			isClient = true
		}
	}
	return isClient
}

//func insertIntoPessoa(rawdata [][]string, dbi *db.MySQLDatabase) error {
func insertIntoPessoa(rawdata [][]string) error {
	if len(rawdata) > 0 {
		Pessoa := new(entity.FuncPublico)
		for i, column := range rawdata {
			if i > 0 {
				Remuneracaodomes, err := strconv.ParseFloat(changeComma(column[3]), 64)
				if err != nil {
					logs.Errorf("insertIntoPessoa", err.Error())
					return err
				}
				Redutorsalarial, err := strconv.ParseFloat(changeComma(column[8]), 64)
				if err != nil {
					logs.Errorf("insertIntoPessoa", err.Error())
					return err
				}
				Totalliquido, err := strconv.ParseFloat(changeComma(column[9]), 64)
				if err != nil {
					logs.Errorf("insertIntoPessoa", err.Error())
					return err
				}
				if Totalliquido > 20000 {

					Pessoa.Nome = column[0]
					Pessoa.Cargo = column[1]
					Pessoa.Orgao = column[2]
					Pessoa.Remuneracaodomes = Remuneracaodomes
					Pessoa.RedutorSalarial = Redutorsalarial
					Pessoa.TotalLiquido = Totalliquido

					//Funcao para procurar pelo nome no BD (Pessoa.Nome)
					alreadyExists, existingID := checkPersonInDB(Pessoa.Nome)

					if isClient(Pessoa.Nome) {
						Pessoa.ClientedoBanco = true
					} else {
						Pessoa.ClientedoBanco = false
					}

					//Verifica se nome já é cliente
					//Caso cliente, update dos dados e update = true
					if alreadyExists {
						Pessoa.ID = existingID
						Pessoa.Updated = true
						//Pessoa.ClientedoBanco = true
						//jsonData, err := json.Marshal(Pessoa)
						if err != nil {
							logs.Errorf("insertIntoPessoa", err.Error())
							return err
						}
						//log.Println(string(jsonData))
						//Atualiza no banco
						erro := funcpublico.Update(Pessoa)
						if erro != nil {
							logs.Errorf("insertIntoPessoa", erro.Error())
						}

						//Caso não cliente, insere os dados e seta update = true
					} else {
						//Pessoa.ClientedoBanco = false
						//Insere no banco
						erro := funcpublico.Insert(Pessoa)
						if erro != nil {
							logs.Errorf("insertIntoPessoa", erro.Error())
						}
						Pessoa.Updated = true
					}

					/*	PONTOS:
						- Caso salário líquido > 20k,
						- Verifica se nome já é cliente
						- Caso cliente, update dos dados e update = true
						- Caso não cliente, insere os dados e seta update = true
						- Caso update = false, set totalliquido = 0
						- Ao final, setar novamente update = false em todos os clientes da tabela*/
				}

			}

		}
		//Setar todos os TotalLiquido para 0 de todos os clientes que o update=false
		//Ou seja, todos os funcionarios públicos que deixaram de ser funcionários
		erro := funcpublico.UpdateAllSetTotalLiquido(0)
		if erro != nil {
			logs.Errorf("insertIntoPessoa", erro.Error())
			return erro
		}

		//seta updated = false em todos os clientes da tabela após o término do processamento*/
		erro = funcpublico.UpdateAllSetFlagUpdated(false)
		if erro != nil {
			logs.Errorf("insertIntoPessoa", erro.Error())
			return erro
		}

	} else {
		err := errors.New("the csv file is empty")
		logs.Errorf("insertIntoPessoa", err.Error())
		return err
	}

	return nil
}

func changeComma(pFloatNumber string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pFloatNumber, ",", ";"), ".", ":"), ";", "."), ":", ",")
}

func getHashFromFile(filePath string) (string, error) {
	var stringHashSHA1 string
	file, err := os.Open(filePath)
	if err != nil {
		logs.Errorf("getHashFromFile", err.Error())
		return stringHashSHA1, err
	}
	defer file.Close()
	hashSHA1 := sha1.New()
	_, err = io.Copy(hashSHA1, file)
	if err != nil {
		logs.Errorf("getHashFromFile", err.Error())
		return stringHashSHA1, err
	}
	stringHashSHA1 = hex.EncodeToString(hashSHA1.Sum(nil))
	return stringHashSHA1, err
}

func getLastFiles(pathDir string, countFiles int, extension string) []string {
	var filesName []string
	files, err := ioutil.ReadDir(pathDir)
	if err != nil {
		logs.Errorf("getHashFromFile", err.Error())
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Unix() > files[j].ModTime().Unix()
	})
	if countFiles > 0 && extension != "" {
		for i, file := range files {
			if file.Mode().IsRegular() {
				if filepath.Ext(file.Name()) == extension {
					if len(filesName) < countFiles {
						filesName = append(filesName, files[i].Name())
					} else {
						break
					}
				}
			}
		}
	}

	return filesName
}
