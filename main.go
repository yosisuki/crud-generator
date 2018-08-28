package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/helper"
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


	GenerateRepository(className, packageName)
	GenerateFieldsConstant(className, packageName, stringsArray)
	GenerateController(className, packageName, classNameLowerFirst, classNameTitle, classNameConstant)
	GenerateServiceApi(className, packageName)
	GenerateServiceImpl(className, packageName, classNameLowerFirst)

}
func GenerateServiceImpl(className string, packageName string, classNameLowerFirst string) {
	b, err := ioutil.ReadFile("templates/service-impl/service_impl_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)
	output = strings.Replace(output,"[classNameLowerFirst]", classNameLowerFirst, -1)


	fmt.Println(output)
}
func GenerateController(className string, packageName string, classNameLowerFirst string, classNameTitle string, classNameConstant string) {
	b, err := ioutil.ReadFile("templates/rest-web/controller/controller_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)
	output = strings.Replace(output,"[classNameLowerFirst]", classNameLowerFirst, -1)
	output = strings.Replace(output,"[classNameTitle]", classNameTitle, -1)
	output = strings.Replace(output,"[classNameConstant]", classNameConstant, -1)


	fmt.Println(output)
}

func GenerateServiceApi(className string, packageName string) {
	b, err := ioutil.ReadFile("templates/service-api/service_api_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)


	fmt.Println(output)
}

func GenerateRepository(className string, packageName string) {
	b, err := ioutil.ReadFile("templates/dao/api/repository_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)

	fmt.Println(output)
}

func GenerateFieldsConstant(className string, packageName string, stringsArray []string) {
	b, err := ioutil.ReadFile("templates/entity/constants/fields/fields_constant_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)
	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)

	fieldConstants := ""

	for _, value := range stringsArray {
		fieldConstants += "public static final String "
		valuesArray := strings.Split(value, " ")
		fieldConstants += strings.ToUpper(helper.ToSnakeCase(valuesArray[1])) + " \n"
	}

	output = strings.Replace(output, "[fieldConstants]", fieldConstants, -1)

	fmt.Println(output)
}