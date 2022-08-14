package fileops

import "io/ioutil"

func ReadFile(path string) string {
	s, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(s)
}
