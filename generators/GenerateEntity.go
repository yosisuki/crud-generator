package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/helpers"
)

func GenerateEntity(className string, classNameConstant string, classNameLowerFirst string, packageName string, stringsArray []string) string {
	b, err := ioutil.ReadFile("templates/entity/dao/entity_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)
	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)
	output = strings.Replace(output,"[classNameConstant]", classNameConstant, -1)
	output = strings.Replace(output,"[classNameLowerFirst]", classNameLowerFirst, -1)

	classContentsEntity := ""

	for _, value := range stringsArray {
		valuesArray := strings.Split(value, " ")
		classContentsEntity += "@Field(value = "+className+"Fields."+ strings.ToUpper(helpers.ToSnakeCase(valuesArray[1])) +")\n"
		classContentsEntity += "private " + value + ";\n"
		classContentsEntity += "public void set" + helpers.UcFirst(valuesArray[1]) + "("+valuesArray[0] + " " + valuesArray[1]+") {this."+valuesArray[1]+" = "+valuesArray[1]+";} \n"
		classContentsEntity += "public "+valuesArray[0]+" get" + helpers.UcFirst(valuesArray[1]) + "() {return "+valuesArray[1]+";} \n\n"
	}

	output = strings.Replace(output, "[classContentsEntity]", classContentsEntity, -1)

	fmt.Println(output)

	return output

}