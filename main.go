package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"spring-crud-generator/generators"
	"os"
	"log"
	"spring-crud-generator/helpers"
)

func main(){
	b, err := ioutil.ReadFile("GatewayEndPoint.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	stringsArray := strings.Split(str,"\n")

	className := "CrudGatewayEndPoint"
	packageName := "com.tiket.tix.gateway.dashboard"
	classNameLowerFirst := "crudGatewayEndPoint"
	classNameTitle := "CRUD Gateway EndPoint"
	classNameConstant := strings.ToUpper(helpers.ToSnakeCase(className))
	mainFolder := "/Users/yosisuki/Desktop/tix-api/TIX-GATEWAY-DASHBOARD/"
	packageFolder := strings.Replace(packageName, ".", "/", -1)+ "/"


	repository := generators.GenerateRepository(className, packageName)
	entity := generators.GenerateEntity(className, classNameConstant,classNameLowerFirst, packageName, stringsArray)
	fieldsConstant := generators.GenerateFieldsConstant(className, packageName, stringsArray)
	controller := generators.GenerateController(className, packageName, classNameLowerFirst, classNameTitle, classNameConstant)
	serviceApi := generators.GenerateServiceApi(className, packageName)
	serviceImpl := generators.GenerateServiceImpl(className, packageName, classNameLowerFirst)
	restWebModelRequest := generators.GenerateRequest(className, packageName, stringsArray)
	restWebModelResponse := generators.GenerateResponse(className, packageName, stringsArray)

	writeFile(mainFolder + "dao/src/main/java/" + packageFolder + "dao/api/", className + "Repository.java", repository)
	writeFile(mainFolder + "entity/src/main/java/" + packageFolder + "entity/dao/", className +".java", entity)
	writeFile(mainFolder + "entity/src/main/java/" + packageFolder + "entity/constant/fields/", className + "Fields.java", fieldsConstant)
	writeFile(mainFolder + "rest-web/src/main/java/"+ packageFolder + "rest/web/controller/", className  + "Controller.java", controller)
	writeFile(mainFolder + "service-api/src/main/java/" + packageFolder , "service/api/" + className + "Service.java", serviceApi)
	writeFile(mainFolder + "service-impl/src/main/java/" + packageFolder, "service/impl/" + className + "ServiceImpl.java", serviceImpl)
	writeFile(mainFolder + "rest-web-model/src/main/java/" + packageFolder + "rest/web/model/request/", className + "Request.java", restWebModelRequest)
	writeFile(mainFolder + "rest-web-model/src/main/java/" + packageFolder + "rest/web/model/response/", className + "Response.java", restWebModelResponse)
}
func writeFile(folder string, fileName string, content string) {

	os.MkdirAll(folder, os.ModePerm)

	file, err := os.Create(folder + fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, content)
}