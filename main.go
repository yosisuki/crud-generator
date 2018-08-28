package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/generators"
)

func main(){
	b, err := ioutil.ReadFile("dataDiri.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	stringsArray := strings.Split(str,"\n")

	className := "DataDiri"
	packageName := "com.tiket.tix.flight.search"
	classNameLowerFirst := "dataDiri"
	classNameTitle := "Data Diri"
	classNameConstant := "DATA_DIRI"


	generators.GenerateRepository(className, packageName)
	generators.GenerateFieldsConstant(className, packageName, stringsArray)
	generators.GenerateController(className, packageName, classNameLowerFirst, classNameTitle, classNameConstant)
	generators.GenerateServiceApi(className, packageName)
	generators.GenerateServiceImpl(className, packageName, classNameLowerFirst)
	generators.GenerateRequest(className, packageName, stringsArray)
	generators.GenerateResponse(className, packageName, stringsArray)

}