package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func GenerateServiceApi(className string, packageName string) string {
	b, err := ioutil.ReadFile("templates/service-api/service_api_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)


	fmt.Println(output)

	return output
}