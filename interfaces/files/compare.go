package main

import (
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/funcpublico"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/historico"
	"strconv"
	"strings"
	"bytes"
	//"sync"
)

//Clients get clients names from csv file
type Clients struct {
	Nome string `json:"nome"`
}
type CSVFile struct{
	Name string
	FullPath string
	TotalOfLines int64
}

type LineInterval struct{
	Start int64
	End int64
}

func getInfoFromCSVFile() (*CSVFile, error) {
	workPath, err := getFileName() 
	csvFile := new(CSVFile)
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return csvFile, err
	}
	fileName := getLastFiles(workPath.Directory, 1, ".txt")
	fullPath := workPath.Directory + fileName[0]
	csvFile.FullPath = fullPath
	csvFile.Name = fileName[0]
	cmd := exec.Command("bash", "-c", "cat " + fullPath + " | wc -l")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return csvFile, err
	}
	
	numberStr := strings.Trim(out.String(), "\"")
	numberStr = strings.ReplaceAll(numberStr, "\n", "")
	numLines, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		return nil, err
	}
	csvFile.TotalOfLines = numLines -1
	return csvFile, nil
}

func ProcessMultiLinesCSVFile() error {
	var line LineInterval
	var lines []LineInterval
	csvFileInfo, err := getInfoFromCSVFile()
	var linesByGoRotine int64
	TotalGoRotine := 1
	linesByGoRotine = 100000
	line.Start = 0
	if csvFileInfo.TotalOfLines > csvFileInfo.TotalOfLines / linesByGoRotine{
		ToProcess := csvFileInfo.TotalOfLines / linesByGoRotine
		var i int64
		for i =0; i <= ToProcess; i++ {
			line.End = line.Start + linesByGoRotine - 1
			if line.End > csvFileInfo.TotalOfLines{
				line.End = csvFileInfo.TotalOfLines - line.Start
				line.End = line.End + line.Start
			}
			lines = append(lines,line)
			line.Start = line.Start + linesByGoRotine 
			TotalGoRotine++ 
		}
	}else{
		line.End = csvFileInfo.TotalOfLines
		lines = append(lines,line)
	}
	
	/*wg := &sync.WaitGroup{}
	wg.Add(TotalGoRotine)*/
	for i:=0; i < len(lines); i++{
	//	PersistLinesInDb(wg, lines[i].Start, lines[i].End)
		PersistLinesInDb(lines[i].Start, lines[i].End)
	}
	//wg.Wait()
	err = AfterProcess()
	return err
}

//func PersistLinesInDb(wg *sync.WaitGroup, pStart int64, pEnd int64) {
func PersistLinesInDb(pStart int64, pEnd int64) {
	CSVLines, err := openCSV()
	for j:= pStart; j < pEnd; j++{
			
		persistPessoa(CSVLines[j])
		if err != nil {
			logs.Errorf("PersistLinesInDb", err.Error())
		}
	}
	//wg.Done()
}

func openCSV() ([][]string, error) {
	workPath, err := getFileName()
	var lines [][]string
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return nil, err
	}
	fileName := getLastFiles(workPath.Directory, 1, ".txt")
	fullPath := workPath.Directory + fileName[0]
	csvfile, err := os.Open(fullPath)
	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return nil, err
	}
	defer csvfile.Close()
	logs.Info("openFileCSV", "Opening file: "+fullPath)
	reader := csv.NewReader(csvfile)
	reader.Comma = ';'
	logs.Info("openFileCSV", "Reading file...")
//	rawdata, err := reader.ReadAll()
	numLine := 0
	logs.Info("openFileCSV", "Updating data in DB...")
	for {
		numLine++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logs.Errorf("openFileCSV", err.Error())
		}
		if numLine > 1 {
			lines = append(lines, record)
		}
	}
	return lines, err
}

func AfterProcess() error {
	//Setar todos os TotalLiquido para 0 de todos os clientes que o update=false
	//Ou seja, todos os funcionarios públicos que deixaram de ser funcionários
	err := funcpublico.UpdateAllSetRemuneracaodoMes(0) //CHANGED from totalliquido to remuneracaodomes
	if err != nil {
		logs.Errorf("insertIntoPessoa", err.Error())
		return err
	}

	//seta updated = false em todos os clientes da tabela após o término do processamento*/
	err = funcpublico.UpdateAllSetFlagUpdated(false)
	if err != nil {
		logs.Errorf("insertIntoPessoa", err.Error())
		return err
	}
	hist, err := funcpublico.GetAllFuncPublico()
	if err != nil {
		logs.Errorf("GetAllFuncPublico", err.Error())
		return err
	}
	historico.Insert(hist)
	if err != nil {
		logs.Errorf("Insert Historico", err.Error())
		return err
	}
	return err
}

//OpenAndProcessFileCSV open file csv and insert in DB
func OpenCSVAndInsertCSV(intLine int64, endLine int64) error {
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
	//rawdata, err := reader.ReadAll()
	numLine := 0
	logs.Info("openFileCSV", "Updating data in DB...")
	for {
		numLine++
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			logs.Errorf("openFileCSV", err.Error())
		}
		if numLine > 1 {
			err = persistPessoa(record)
			if err != nil {
				logs.Errorf("openFileCSV", err.Error())
			}
		}
	}

	logs.Info("openFileCSV", "Data stored in DB")


	if err != nil {
		logs.Errorf("openFileCSV", err.Error())
		return err
	}

	logs.Info("openFileCSV", "UpdateAllSetTotalLiquido")
	if numLine > 1 {
		//Setar todos os TotalLiquido para 0 de todos os clientes que o update=false
		//Ou seja, todos os funcionarios públicos que deixaram de ser funcionários
		err = funcpublico.UpdateAllSetRemuneracaodoMes(0) //CHANGED from totalliquido to remuneracaodomes
		if err != nil {
			logs.Errorf("insertIntoPessoa", err.Error())
			return err
		}

		//seta updated = false em todos os clientes da tabela após o término do processamento*/
		err = funcpublico.UpdateAllSetFlagUpdated(false)
		if err != nil {
			logs.Errorf("insertIntoPessoa", err.Error())
			return err
		}
		hist, err := funcpublico.GetAllFuncPublico()
		if err != nil {
			logs.Errorf("GetAllFuncPublico", err.Error())
			return err
		}
		historico.Insert(hist)
		if err != nil {
			logs.Errorf("Insert Historico", err.Error())
			return err
		}
	}


	return err
}

//Check if person already exists in DB (by name)
func checkPersonInDB(name string) (bool, int) {
	Pessoa := new(entity.FuncPublico)
	alreadyInDB := false
	Pessoa, erro := funcpublico.GetByName(name)
	if erro != nil {
		logs.Errorf("CheckpersonInDB", erro.Error())
	}
	if Pessoa.Nome == strings.Trim(name, " ") {
		alreadyInDB = true
	}
	return alreadyInDB, Pessoa.ID
}

//func to check if its a client
func isClient(name string) bool {
	isClient := false
	file, erro := ioutil.ReadFile("../API/Clientlist.json")
	if erro != nil {
		logs.Errorf("isClient", erro.Error())
	}
	data := []Clients{}

	erro = json.Unmarshal(file, &data)
	if erro != nil {
		logs.Errorf("isClient", erro.Error())
	}

	for _, value := range data {
		if strings.Trim(name, " ") == value.Nome {
			//fmt.Println(value.Nome + "=" + name)
			isClient = true
		}
	}
	return isClient
}

//func insertIntoPessoa(rawdata [][]string, dbi *db.MySQLDatabase) error {
func persistPessoa(column []string) error {
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
	/*	PONTOS:
		- Caso salário líquido > 20k,
		- Verifica se nome já é cliente
		- Caso cliente, update dos dados e update = true
		- Caso não cliente, insere os dados e seta update = true
		- Caso update = false, set totalliquido = 0
		- Ao final, setar novamente update = false em todos os clientes da tabela*/
	if Remuneracaodomes > 20000 { //CHANGED from totalliquido to remuneracaodomes
		Pessoa := new(entity.FuncPublico)
		Pessoa.Nome = column[0]
		Pessoa.Nome = strings.Replace(Pessoa.Nome, "'", "''", 1) //prevent from single quotes in names (Escape character)
		Pessoa.Cargo = column[1]
		Pessoa.Orgao = column[2]
		Pessoa.Remuneracaodomes = Remuneracaodomes
		Pessoa.RedutorSalarial = Redutorsalarial
		Pessoa.TotalLiquido = Totalliquido
		Pessoa.Updated = true
		//Funcao para procurar pelo nome no BD (Pessoa.Nome)
		alreadyExists, existingID := checkPersonInDB(Pessoa.Nome)

		if isClient(Pessoa.Nome) {
			Pessoa.ClientedoBanco = true
		} else {
			Pessoa.ClientedoBanco = false
		}

		//Verifica se nome já é cliente
		//Caso cliente, update dos dados
		if alreadyExists {
			Pessoa.ID = existingID
			//Atualiza no banco
			erro := funcpublico.Update(Pessoa)
			if erro != nil {
				logs.Errorf("insertIntoPessoa", erro.Error())
				return erro
			}

			//Caso não cliente, insere os dados
		} else {
			//Insere no banco
			erro := funcpublico.Insert(Pessoa)
			if erro != nil {
				logs.Errorf("insertIntoPessoa", erro.Error())
				return erro
			}
		}
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
