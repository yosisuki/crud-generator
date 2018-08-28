package generators

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func GenerateRepository(className string, packageName string) string{
	b, err := ioutil.ReadFile("templates/dao/api/repository_template.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	contents := string(b)

	output := strings.Replace(contents,"[className]", className, -1)
	output = strings.Replace(output,"[packageName]", packageName, -1)

	fmt.Println(output)

	return output
}
