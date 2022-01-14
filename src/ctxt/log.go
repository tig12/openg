/******************************************************************************

    Logging

    @license    GPL
    @history    2021-07-13 23:50:21+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"fmt"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
	"os"
	"strings"
	"time"
)

func LogError(err error) {
	werr.Print(err)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := NewContext() // TODO find a way to avoid creating a new context here
		if !strings.HasPrefix(r.RequestURI, "/static/") && ctx.Config.Run.Mode == "prod" {
			now := time.Now().Format(time.RFC3339) // string
			year := now[0:4]
			month := now[0:7]
			day := now[0:10]
			sep := string(os.PathSeparator)
			// logfile looks like /path/to/log/2022/2022-01/2022-01-14.log
			logdir := strings.Join([]string{ctx.Config.Paths.Log, year, month}, sep)
			logfile := logdir + sep + day + ".log"
			var file *os.File
			defer file.Close()
			if _, err = os.Stat(logfile); err != nil && os.IsNotExist(err) {
				os.MkdirAll(logdir, 0755)
				file, err = os.Create(logfile)
				if err != nil {
					panic(err)
				}
			}
			file, err = os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			_, err = fmt.Fprintf(file, "%s %s %s %s\n",
				now,
				r.RemoteAddr,
				r.RequestURI,
				r.Referer())
			if err != nil {
				panic(err)
			}
		}
		next.ServeHTTP(w, r) // Call the next handler
	})
}
