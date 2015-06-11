package main

import (
    "fmt"
    "strings"
    "reflect"
)

// Convert []interface{} to []string{}
func arrIntfToStr(inp []interface{}) []string {
    outp := make([]string, len(inp))
    for i, val := range inp {
        outp[i] = fmt.Sprintf("%v", val)
    }
    return outp
}

// Convert []string{} to []interface{}
func arrStrToIntf(inp []string) []interface{} {
    outp := make([]interface{}, len(inp))
    for i, val := range inp {
        outp[i] = val
    }
    return outp
}

// Determine if index is the last element in the array
func arrLast(i int, inp interface{}) bool {
    return i == reflect.ValueOf(inp).Len() - 1
}

// Join elements in an array to a string
func arrJoin(sep string, inp []interface{}) string {
    return strings.Join(arrIntfToStr(inp), sep)
}

// Split string into an array
func strSplit(sep string, inp string) []interface{} {
    return arrStrToIntf(strings.Split(inp, sep))
}

// Repeat string x number of times
func strRepeat(rep int, inp string) string {
        return strings.Repeat(inp, rep)
}

// Get keys from interface{}
func intfKeys(inp interface{}) (interface{}, error) {
    if inp == nil {
        return nil, nil
    }

    val := reflect.ValueOf(inp)
    if val.Kind() != reflect.Map {
        return nil, fmt.Errorf("Cannot call keys on a non-map value: %v", inp)
    }

    vk := val.MapKeys()
    k := make([]interface{}, val.Len())
    for i, _ := range k {
        k[i] = vk[i].Interface()
    }

    return k, nil
}

// Get type (usefull for debugging templates)
func intfType(inp interface{}) string {
    return fmt.Sprintf("%v", reflect.TypeOf(inp))
}

// Test if type is a map i.e. not printable
func intfIsMap(inp interface{}) bool {
    return reflect.TypeOf(inp).Kind() == reflect.Map
}

// String replace
func strReplace(oldStr string, newStr string, str string) string {
    return strings.Replace(str, oldStr, newStr, -1)
}

// String trim
func strTrim(trim string, str string) string {
    return strings.Trim(str, trim)
}

// String trim left
func strTrimLeft(trim string, str string) string {
    return strings.TrimLeft(str, trim)
}

// String trim right
func strTrimRight(trim string, str string) string {
    return strings.TrimRight(str, trim)
}

// If no value is passed for the second arg. it returns the default
func intfDefault(def interface{}, inp_opt ...interface{}) interface{} {
    if len(inp_opt) > 0 {
        def = inp_opt[0]
    }
    return def
}

func strCenter(size int, str string) string {
    if size < len(str) {
        return str
    }

    pad := (size - len(str)) / 2
    lpad := pad
    rpad := size - len(str) - lpad

    return fmt.Sprintf("%s%s%s", strings.Repeat(" ", lpad), str, strings.Repeat(" ", rpad))
}

// Capitalize first character in string
// date/time func
