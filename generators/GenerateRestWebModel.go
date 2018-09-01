package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/helpers"
)

func GenerateResponse(className string, packageName string, stringsArray []string) string {
	b, err := ioutil.ReadFile("templates/rest-web-model/response/response_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)
	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)

	classContentsResponse := ""

	for _, value := range stringsArray {
		classContentsResponse += "private " + value + ";\n"
		valuesArray := strings.Split(value, " ")
		classContentsResponse += "public void set" + helpers.UcFirst(valuesArray[1]) + "("+valuesArray[0] + " " + valuesArray[1]+") {this."+valuesArray[1]+" = "+valuesArray[1]+";} \n"
		classContentsResponse += "public "+valuesArray[0]+" get" + helpers.UcFirst(valuesArray[1]) + "() {return "+valuesArray[1]+";} \n\n"
	}

	output = strings.Replace(output, "[classContentsResponse]", classContentsResponse, -1)

	fmt.Println(output)

	return output

}
func GenerateRequest(className string, packageName string, stringsArray []string) string{
	b, err := ioutil.ReadFile("templates/rest-web-model/request/request_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)
	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)

	classContentsRequest := ""

	for _, value := range stringsArray {
		classContentsRequest += "@NotBlank\n"
		classContentsRequest += "private " +value + ";\n"
		valuesArray := strings.Split(value, " ")
		classContentsRequest += "public void set" + helpers.UcFirst(valuesArray[1]) + "() {this."+valuesArray[1]+" = "+valuesArray[1]+";} \n"
		classContentsRequest += "public "+valuesArray[0]+" get" + helpers.UcFirst(valuesArray[1]) + "() {return "+valuesArray[1]+";} \n\n"
	}

	output = strings.Replace(output, "[classContentsRequest]", classContentsRequest, -1)

	fmt.Println(output)

	return output
}
