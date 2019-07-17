package logs

/*

**** Need environment variable LOG_PATH **** 

Usage:

	Set the environment variable LOG_PATH with the path to log folder

	To use add on import

	Commands:

	- Errorf("SOURCE file or package","erro message") -> Log to error_yyyy_mm_dd.log
		Format - SOURCE | DESCRIPTION

	- Warning("SOURCE file or package","warn message") -> Log to warning_yyyy_mm_dd.log
		Format - SOURCE | DESCRIPTION

	- Request("method","SOURCE address","destination address","parameters","observation") -> Log to request_yyyy_mm_dd.log
		Format - METHOD | SOURCE | DESTINATION | PARAMETERS | OBS

	- Info("SOURCE file or package","info message") -> Log to info_yyyy_mm_dd.log
		Format - SOURCE | DESCRIPTION

*/


import ("os"
	"time"
	"fmt"
)

func getFileName(name string) string {
	LogPath := os.Getenv("LOG_PATH")
	t := time.Now()
	var filename string
	filename = fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,name,t.Year(),t.Month(),t.Day())
	return filename
}

// Errorf formated
func Errorf(source string, description string) {
	t := getNow()
	s := fmt.Sprintf("%s : [ERROR] : %s | %s",t,source,description)
	logIt(getFileName("error"),s)
}

// Warning formated
func Warning(source string, description string) {
	t := getNow()
	s := fmt.Sprintf("%s : [WARNING] : %s | %s",t,source,description)
	logIt(getFileName("warning"),s)
}

// Request formated
func Request(method string,source string,destination string,parameters string, observation string){
	t := getNow()
	s := fmt.Sprintf("%s : [REQUEST] : %s | %s | %s | %s | %s",t,method,source,destination,parameters,observation)
	logIt(getFileName("request"),s)
}

// SimpleRequest formated
func SimpleRequest(source string, description string){
	t := getNow()
	s := fmt.Sprintf("%s : [REQUEST] : %s | %s",t,source,description)
	logIt(getFileName("request"),s)
}

//Info formated
func Info(source string, description string){
	t := getNow()
	s := fmt.Sprintf("%s : [INFO] : %s | %s",t,source,description)
	logIt(getFileName("info"),s)
}

func getNow() string {
	t := time.Now()
	s := fmt.Sprintf("%d:%d:%d",t.Hour(),t.Minute(),t.Second())
	return s
}

func logIt(file string, message string){
	f,_ := os.OpenFile(file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0664)
	writeFile(message, f)
}

func writeFile(s string,f *os.File) {
	_,err := fmt.Fprintln(f,s)
	if(err!=nil){
		panic(err)
	}
	f.Close()
}