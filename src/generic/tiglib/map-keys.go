package tiglib


// Returns a slice containing the keys of a map[string]int
func MapKeysStringInt(m map[string]int) (keys []string) {
    keys = make([]string, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    return keys
}


// Returns a slice containing the keys of a map
// TODO function OK, but unable to use it
func MapKeys(m map[interface{}]interface{}) (keys []interface{}) {
    keys = make([]interface{}, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    return keys
}
