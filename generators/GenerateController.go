package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func GenerateController(className string, packageName string, classNameLowerFirst string, classNameTitle string, classNameConstant string) string {
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

	return output;
}