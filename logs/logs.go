package logs

import ("os"
	"time"
	"fmt"
)

// Logger log structure
type Logger struct {
	LogPath string
	warningfile string
	errorfile string
	requestfile string
}

// NewLog Generate Logger
func NewLog(LogPath string) Logger {
	t := time.Now()
	errorfile := fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,"error",t.Year(),t.Month(),t.Day())
	warningfile := fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,"warning",t.Year(),t.Month(),t.Day())
	requestfile := fmt.Sprintf("%s%s_%d_%s_%d.log", LogPath,"html_interations",t.Year(),t.Month(),t.Day())
	NwLog := Logger{
		LogPath: LogPath,
		errorfile: errorfile,
		warningfile: warningfile,
		requestfile: requestfile}
	e,_ := os.Create(NwLog.errorfile)
	w,_ := os.Create(NwLog.warningfile)
	r,_ := os.Create(NwLog.requestfile)
	e.Close()
	w.Close()
	r.Close()
	return NwLog
}

// Errorf formated
func (L Logger) Errorf(errorString string) {
	s := fmt.Sprintf(" [ERROR] : %s",errorString)
	logIt(L.errorfile,s)
}

// Warning formated
func (L Logger) Warning(warnString string) {
	s := fmt.Sprintf(" [ERROR] : %s",warnString)
	logIt(L.warningfile,s)
}

// Request formated
func (L Logger) Request(requestString string){
	s := fmt.Sprintf(" [ERROR] : %s",requestString)
	logIt(L.requestfile,s)
}

func logIt(file string,message string){
	t := time.Now()
	message = fmt.Sprintf("%s : %s ",t,message)
	f,_ := os.OpenFile(file, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	writeFile(message, f)
}

func writeFile(s string,f *os.File) {
	_,err := fmt.Fprintln(f,s)
	if(err!=nil){
		panic(err)
	}
	f.Close()
}