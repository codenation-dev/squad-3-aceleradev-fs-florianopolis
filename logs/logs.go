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
	t := time.Now()
	f, _ := os.OpenFile(L.errorfile, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	s := fmt.Sprintf("%s : [ERROR] : %s",t,errorString)
	writeFile(s,f)
}

// Warning formated
func (L Logger) Warning(warnString string) {
	t := time.Now()
	f, _ := os.OpenFile(L.warningfile, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	s := fmt.Sprintf("%s : [ERROR] : %s",t,warnString)
	writeFile(s,f)
}

// Request formated
func (L Logger) Request(htmlString string){}

func writeFile(s string,f *os.File) {
	_,err := fmt.Fprintln(f,s)
	if(err!=nil){
		panic(err)
	}
	f.Close()
}