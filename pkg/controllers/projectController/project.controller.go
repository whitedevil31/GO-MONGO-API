package projectController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/whitedevil31/go-mongo-api/pkg/models/projects"
	"github.com/whitedevil31/go-mongo-api/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
var validate *validator.Validate


func AddProject(w http.ResponseWriter, r *http.Request){
	userId :=  r.Context().Value("id")
	fmt.Println(userId)
	validate = validator.New()
	addProject:= &projects.Project{}
	utils.ParseBody(r, addProject)
	validateErr := validate.Struct(addProject)
	fmt.Println(validateErr)
	if validateErr!=nil{
		for _, e := range validateErr.(validator.ValidationErrors) {
		utils.JSONError(w,"validation failed for "+ e.Field()+" field",utils.GetCode("VALIDATION_FAILED"))
		return
	}
	}
	fmt.Println("GOM")
	result,err:= projects.AddProject(addProject,userId.(primitive.ObjectID))
	fmt.Println(err)
	if err!=nil{

	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjects(w http.ResponseWriter, r *http.Request){
	userId :=  r.Context().Value("id")

	result,err:= projects.GetProjects(userId.(primitive.ObjectID))
	fmt.Println(err)
	if err!=nil{

	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetProject(w http.ResponseWriter, r *http.Request){
	userId :=  r.Context().Value("id")
	vars := mux.Vars(r)
	projectId,err :=  primitive.ObjectIDFromHex(vars["projectId"]) 
	if err!=nil{
		utils.JSONError(w,"INVALID_ID",utils.GetCode("INVALID_ID"))
		return 
	}
	result,err:= projects.GetProject(userId.(primitive.ObjectID),projectId)
	fmt.Println(err)
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteProject(w http.ResponseWriter, r *http.Request){
	userId :=  r.Context().Value("id")
	vars := mux.Vars(r)
	projectId,err :=  primitive.ObjectIDFromHex(vars["projectId"]) 
	if err!=nil{
		utils.JSONError(w,"INVALID_ID",utils.GetCode("INVALID_ID"))
		return 
	}
	result,err:= projects.DeleteProject(userId.(primitive.ObjectID),projectId)
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateProject(w http.ResponseWriter, r *http.Request){
	userId :=  r.Context().Value("id")
	vars := mux.Vars(r)
	var updateProject = &projects.Project{}
	utils.ParseBody(r,updateProject)
	projectId,err :=  primitive.ObjectIDFromHex(vars["projectId"]) 
	if err!=nil{
		utils.JSONError(w,"INVALID_ID",utils.GetCode("INVALID_ID"))
		return 
	}
	result,err:= projects.UpdateProject(userId.(primitive.ObjectID),projectId,updateProject)
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}