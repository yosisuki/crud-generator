package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/generators"
)

func main(){
	b, err := ioutil.ReadFile("systemParameter.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	stringsArray := strings.Split(str,"\n")

	className := "SystemParameter"
	packageName := "com.tiket.tix.flight.core"
	classNameLowerFirst := "systemParameter"
	classNameTitle := "System Parameter"
	classNameConstant := "SYSTEM_PARAMETER"


	generators.GenerateRepository(className, packageName)
	generators.GenerateFieldsConstant(className, packageName, stringsArray)
	generators.GenerateController(className, packageName, classNameLowerFirst, classNameTitle, classNameConstant)
	generators.GenerateServiceApi(className, packageName)
	generators.GenerateServiceImpl(className, packageName, classNameLowerFirst)
	generators.GenerateRequest(className, packageName, stringsArray)
	generators.GenerateResponse(className, packageName, stringsArray)

}