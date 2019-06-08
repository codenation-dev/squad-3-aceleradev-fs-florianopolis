package logs

import ("os"
	"time"
	"fmt"
)

func getFileName(name string) string {
	LogPath := os.Getenv("LOG_PATH")
	t := time.Now()
	var filename string
	switch name {
	case "error":
		filename = fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,"error",t.Year(),t.Month(),t.Day()) 
	case "warning":
		filename = fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,"warning",t.Year(),t.Month(),t.Day())
	case "request":
		filename = fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,"html_interations",t.Year(),t.Month(),t.Day())
	}
	return filename
}

// Errorf formated
func Errorf(errorString string) {
	s := fmt.Sprintf(" [ERROR] : %s",errorString)
	logIt(getFileName("error"),s)
}

// Warning formated
func Warning(warnString string) {
	s := fmt.Sprintf(" [ERROR] : %s",warnString)
	logIt(getFileName("warning"),s)
}

// Request formated
func Request(requestString string){
	s := fmt.Sprintf(" [ERROR] : %s",requestString)
	logIt(getFileName("request"),s)
}

func logIt(file string,message string){
	
	t := time.Now()
	message = fmt.Sprintf("%s : %s ",t,message)
	f,_ := os.OpenFile(file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, os.ModeAppend)
	writeFile(message, f)
}

func writeFile(s string,f *os.File) {
	_,err := fmt.Fprintln(f,s)
	if(err!=nil){
		panic(err)
	}
	f.Close()
}