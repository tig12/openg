package werr

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// ====== ajouts Thierry ======
// TODO supprimer, plus utile
func Bidon() error {
	return errors.New("Erreur bidon")
}

// SprintHTML returns traceback as string formatted for html display
func SprintHTML(err error) string {
	res := ""
	res += `
<style>
    table.werr{
        border-collapse:collapse;
        margin:.5rem;
    }
    table.werr td{
        padding:2px 4px;
        border:1px solid #a2a9b1;
    }
    table.werr tr:nth-child(even){
        background-color:#f2e6d9;
    }
    table.werr tr:nth-child(odd){
        background-color:lightblue;
    }
    div.werr{
        font-weight:bold;
    }
    div.werr .message{
        font-size:1.3rem;
        display:inline-block;
        background:lightyellow;
        border:1px solid grey;
        padding:1rem;
        margin:1rem;
    }
</style>
    `
	res += "<table class=\"werr\">\n"
	lines := strings.Split(SprintSkip(err, ""), "\n")
	// lines contains traceback separated by "|", displayed in a table
	// and the error message, stored in msg, displayed in a div
	var tmp []string
	msg := ""
	for _, line := range lines {
		tmp = strings.Split(line, "|")
		if len(tmp) == 2 {
			res += "    <tr>\n"
			res += "        <td>" + strings.TrimSpace(tmp[0]) + "</td><td>" + strings.TrimSpace(tmp[1]) + "</td>\n"
			res += "    </tr>\n"
		} else {
			if msg != "" {
				msg += "\n<br>"
			}
			msg += tmp[0]
		}
	}
	res += "</table>\n"
	msg = "<div class=\"werr\"><div class=\"message\">" + msg + "</div></div>"
	res += msg
	return res
}

// ====== fin ajouts Thierry ======

type Error struct {
	Err   error
	trace []uintptr
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Is(err error) bool {
	if err, ok := err.(Error); ok {
		return errors.Is(e.Err, err.Err)
	}
	return errors.Is(e.Err, err)
}

func (e Error) Error() string {
	return e.Err.Error()
}

func New(s string) error {
	return Stack(errors.New(s))
}

func (e Error) SprintSkip(skip string) string {
	trace := []runtime.Frame{}
	frames := runtime.CallersFrames(e.trace)
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "runtime") {
			trace = append(trace, frame)
		}
		if !more {
			break
		}
	}
	maxLenFunc := 0
	s := ""
	for i := len(trace) - 1; i >= 0; i-- {
		frame := trace[i]
		if skip != "" && strings.HasSuffix(frame.Function, skip) {
			continue
		}
		s := fmt.Sprintf("%s:%d", frame.File, frame.Line)
		if len(s) > maxLenFunc {
			maxLenFunc = len(s)
		}
	}

	for i := len(trace) - 1; i >= 0; i-- {
		frame := trace[i]
		if skip != "" && strings.HasSuffix(frame.Function, skip) {
			s = ""
			continue
		}
		st := fmt.Sprintf("%s:%d", frame.File, frame.Line)
		s += fmt.Sprintf("%-"+strconv.Itoa(maxLenFunc)+"s | %s\n", st, frame.Function)
	}
	return s
}

// FprintSkip write SprintSkip to writer
func FprintSkip(w io.Writer, err error, skip string) {
	w.Write([]byte(SprintSkip(err, skip)))
}

// PrintSkip write SprintSkip to stdout
func PrintSkip(err error, skip string) {
	FprintSkip(os.Stdout, err, skip)
}

// PrintSkip remove lines before skip in suffix
// ex : ServeHTTP
// return as string
func SprintSkip(err error, skip string) string {
	var e Error
	if errors.As(err, &e) {
		return e.SprintSkip(skip) + err.Error() + "\n"
	}
	return fmt.Sprintf("%+v\n", err) // pkg.errors stack
}

// Fprint write traceback in f
func Fprint(f io.Writer, err error) {
	f.Write([]byte(Sprint(err)))
}

// Sprint return traceback as string
func Sprint(err error) string {
	return SprintSkip(err, "")
}

// Print print traceback to stdout
func Print(err error) {
	Fprint(os.Stdout, err)
}

// Wrapf wrap error with stack only if not already
// error is wrapped with fmt.Errorf(msg + " : %w",err)
func Wrapf(err error, msg string, args ...interface{}) error {
	if err == nil {
		return err
	}
	s := fmt.Sprintf(msg, args...)
	err = fmt.Errorf(s+"\n%w", err)
	var es Error
	if errors.As(err, &es) {
		return err
	}
	stk := getStackTrace(3)
	return Error{Err: err, trace: stk}
}

// Stack add stack trace to an error if it's not
func Stack(e error) error {
	if e == nil {
		return e
	}
	var es Error
	if errors.As(e, &es) {
		return e
	}
	stk := getStackTrace(3)
	return Error{Err: e, trace: stk}
}

// getStackTrace return Frames after nb
func getStackTrace(nb int) []uintptr {
	stackBuf := make([]uintptr, 1024)
	length := runtime.Callers(nb, stackBuf[:])
	stack := stackBuf[:length]
	return stack
}
