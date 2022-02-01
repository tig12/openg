/******************************************************************************

    Logging

    @license    GPL
    @history    2021-07-13 23:50:21+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
    "crypto/md5"
	"fmt"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
	"openg.local/openg/generic/tiglib"
	"os"
	"strings"
	"time"
)

/* 
func LogError(err error) {
    werr.Print(err)
    return
}
*/
func LogError(theError error) {
    ctx := NewContext() // TODO find a way to avoid creating a new context here ?
    var err error
    now := time.Now().Format(time.RFC3339) // string
    if ctx.Config.Run.Mode == "dev" {
        fmt.Printf("\n======================= %s =======================\n", now)
        werr.Print(theError)
        return
    }
    logfile := ctx.Config.Paths.Log + string(os.PathSeparator) + "error.log"
    osfile, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0644)
    defer osfile.Close()
    if err != nil  && os.IsNotExist(err){
        osfile, err = os.Create(logfile)
        if err != nil {
            panic(err)
        }
    }
    _, err = fmt.Fprintf(osfile, "\n======================= %s =======================\n%s\n", now, err)
    if err != nil {
        panic(err)
    }
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := NewContext() // TODO find a way to avoid creating a new context here ?
		if !strings.HasPrefix(r.RequestURI, "/static/") && ctx.Config.Run.Mode == "prod" {
			now := time.Now().Format(time.RFC3339) // string
			year := now[0:4]
			month := now[0:7]
			day := now[0:10]
			sep := string(os.PathSeparator)
			// logfile looks like /path/to/logs/2022/2022-01/2022-01-14.log
			logdir := strings.Join([]string{ctx.Config.Paths.Log, year, month}, sep)
			logfile := logdir + sep + day + ".log"
			var osfile *os.File
			defer osfile.Close()
			if _, err = os.Stat(logfile); err != nil && os.IsNotExist(err) {
				os.MkdirAll(logdir, 0755)
				osfile, err = os.Create(logfile)
				if err != nil {
					panic(err)
				}
			}
			osfile, err = os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			// Compute RemoteAddr
			remoteAddr := ""
			if tmp, ok := r.Header["X-Forwarded-For"]; ok {
			    remoteAddr = tmp[0] // if the site is behind an Apache proxy
            } else {
			    remoteAddr = r.RemoteAddr // not behind an Apache proxy
			}
            // anonymize ip address
            remoteAddr = fmt.Sprintf("%x", md5.Sum([]byte(remoteAddr)))
            remoteAddr = tiglib.StrRandomShorten(remoteAddr, 8)
            
			_, err = fmt.Fprintf(osfile, "%s %s %s %s\n",
				now,
				remoteAddr,
				r.Referer(),
				r.RequestURI)
			if err != nil {
				panic(err)
			}
		}
		next.ServeHTTP(w, r) // Call the next handler
	})
}
