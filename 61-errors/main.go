package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	str := new(string)
	*str = "Hello World"
	rstr := new(string)
	var err error
	rstr, err = ReverseString(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*rstr)
	}

	rstr, err = ReverseString(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*rstr)
	}

	count, err := GetVovelCountFromFile("hello.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("number of vovels:", count)
	}

	count, err = GetVovelCountFromFile2("hello.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("number of vovels:", count)
	}

}

// if the string is nil or the string is empty what do you do
// the error must be the last return value. Technically error can also be the first or any return type.
// but idiomatic go , the error is always a last return type
func ReverseString(str *string) (*string, error) {
	if str == nil {
		// return "", errors.New("nil string")
		return nil, fmt.Errorf("nil string")
	}
	rstr := new(string)
	for _, v := range *str {
		*rstr = string(v) + *rstr
	}
	return rstr, nil
}

func GetVovelCountFromFile(filename string) (uint, error) {
	fd, err := os.Open(filename) // This is a function from standard package called os
	if err != nil {
		return 0, err
	}

	defer fd.Close()
	fi, err := fd.Stat()

	if fi.IsDir() {
		return 0, errors.New("cannot find vovels in a directory")
	}

	if err != nil {
		return 0, err
	}

	bytes := make([]byte, fi.Size())

	_, err = fd.Read(bytes) // This is not reading bytes.. This is fd reads the file and write it to bytes
	if err != nil {
		return 0, err
	}
	count := GetVovelCount(string(bytes))
	return count, nil
}

func GetVovelCountFromFile2(filename string) (uint, error) {
	fd, err := os.Open(filename) // This is a function from standard package called os
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	buf := make([]byte, 10)
	for {
		_, err := fd.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if err == io.EOF {
			// close the channel here
			break
		}
		// send data to channel
		// from other go rountine read from this channel and find vovel count , add that vovel count to the previous count
	}

	return GetVovelCount(string(buf)), nil

}

func GetVovelCount(str string) uint {
	var count uint
	for _, v := range str {
		if string(v) == "A" || string(v) == "a" || string(v) == "E" || string(v) == "e" || string(v) == "I" || string(v) == "i" || string(v) == "O" || string(v) == "o" || string(v) == "U" || string(v) == "u" {
			count++
		}
	}
	return count
}
