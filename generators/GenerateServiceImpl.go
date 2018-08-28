package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func GenerateServiceImpl(className string, packageName string, classNameLowerFirst string) string {
	b, err := ioutil.ReadFile("templates/service-impl/service_impl_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)
	output = strings.Replace(output,"[classNameLowerFirst]", classNameLowerFirst, -1)


	fmt.Println(output)

	return output
}
