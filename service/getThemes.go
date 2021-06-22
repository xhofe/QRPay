package service

import "io/ioutil"

func GetThemes() []string {
	files, _ := ioutil.ReadDir("./views/themes")
	themes := make([]string,0)
	for _,f := range files {
		name := f.Name()
		themes = append(themes, name[:len(name)-5])
	}
	return themes
}