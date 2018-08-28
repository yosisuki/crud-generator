package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/helpers"
)

func GenerateFieldsConstant(className string, packageName string, stringsArray []string) string {
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
		fieldConstants += strings.ToUpper(helpers.ToSnakeCase(valuesArray[1])) + "; \n"
	}

	output = strings.Replace(output, "[fieldConstants]", fieldConstants, -1)

	fmt.Println(output)

	return output
}
