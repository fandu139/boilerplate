package helper

import (
	"log"
	"strconv"
)

// Helpers ...
type Helpers struct{}

// HelpersHandler ...
func HelpersHandler() *Helpers {
	return &Helpers{}
}

// HelperInterface ...
type HelperInterface interface {
	StringToInt64(str string) int64
	StringToInt(str string) int
}

// StringToInt ..
func (handler *Helpers) StringToInt(str string) int {
	/** converting the str1 variable into an int using Atoi method */
	i1, err := strconv.Atoi(str)
	if err == nil {
		log.Println("Error: ", err)
	}
	return i1
}

// StringToInt64 ...
func (handler *Helpers) StringToInt64(str string) int64 {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		log.Println("Error: ", err)
	}
	return i64
}
