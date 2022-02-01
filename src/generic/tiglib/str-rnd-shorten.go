package tiglib

import (
    "math/rand"
)

/**
    Shortens a string by randomly keeping a subset of its characters.
    Ex: StrRandomShorten("6022489e0260daa8ef4af916cf456f07", 8) returns "589a89f4"
    rand.Seed() is not called, so different executions for a given string return the same result.
    
    @param      str         String to shorten
    @param      finalLen    Length of the string after shortening
    
    @depends    InArrayInt, in-array.go
**/
func StrRandomShorten(str string, finalLen int) (string) {
    var res = make([]byte, finalLen)
    l := len(str) - 1
    for i := 0; i < finalLen; i++ {
	    res[i] = str[rand.Intn(l)]
	}
	return string(res)
}
