package userController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/whitedevil31/go-mongo-api/pkg/models/users"
	"github.com/whitedevil31/go-mongo-api/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
var validate *validator.Validate

func StaffSignUp(w http.ResponseWriter, r *http.Request){

	validate = validator.New()
	staffSignUp:= &users.Staff{}
	utils.ParseBody(r, staffSignUp)
	validateErr := validate.Struct(staffSignUp)
	if validateErr!=nil{
		for _, e := range validateErr.(validator.ValidationErrors) {
		utils.JSONError(w,"validation failed for "+ e.Field()+" field",utils.GetCode("VALIDATION_FAILED"))
		return
	}
	}
	
	result,err:= users.StaffSignUp(staffSignUp)
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func StaffLogin(w http.ResponseWriter, r *http.Request){
	validate = validator.New()
	loginData := &utils.StaffLoginCredentials{}
	utils.ParseBody(r,loginData)
	validateErr := validate.Struct(loginData)
	if validateErr!=nil{
		for _, e := range validateErr.(validator.ValidationErrors) {
			fmt.Println(e)
		utils.JSONError(w,e.ActualTag()+" validation failed for "+ e.Field()+" field",utils.GetCode("VALIDATION_FAILED"))
		return
	}
	}
fmt.Println("Validate done!")
	token,err:= users.StaffLogin(loginData)
	fmt.Println("MODEL CONTROL DONE")
	if err!=nil{
		utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
		return   
	}
	res, _ :=json.Marshal(token)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	
}
func StudentSignUp(w http.ResponseWriter, r *http.Request){
	staffId :=  r.Context().Value("id")
	validate = validator.New()
	signUpStudent:= &users.Student{}
	utils.ParseBody(r, signUpStudent)
	validateErr := validate.Struct(signUpStudent)
	if validateErr!=nil{
		for _, e := range validateErr.(validator.ValidationErrors) {
		utils.JSONError(w,"validation failed for "+ e.Field()+" field",utils.GetCode("VALIDATION_FAILED"))
		return
	}
	}
	
	result,err:= users.StudentSignUp(signUpStudent,staffId.(primitive.ObjectID))
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func StudentLogin(w http.ResponseWriter, r *http.Request){
	validate = validator.New()
	loginData := &utils.StudentLoginCredentials{}
	utils.ParseBody(r,loginData)
	validateErr := validate.Struct(loginData)
	if validateErr!=nil{
		for _, e := range validateErr.(validator.ValidationErrors) {
			fmt.Println(e)
		utils.JSONError(w,e.ActualTag()+" validation failed for "+ e.Field()+" field",utils.GetCode("VALIDATION_FAILED"))
		return
	}
	}
fmt.Println("Validate done!")
	token,err:= users.StudentLogin(loginData)
	fmt.Println("MODEL CONTROL DONE")
	if err!=nil{
		utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
		return   
	}
	res, _ :=json.Marshal(token)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	
}

func GetStudentProfile(w http.ResponseWriter, r *http.Request){
	
	userId :=  r.Context().Value("id")
	userProfile,err:=users.GetStudentProfile(userId.(primitive.ObjectID))
	if err!=nil{
		utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	}
	res, _ :=json.Marshal(userProfile)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetStudent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	staffId :=  r.Context().Value("id")
	studentId,err :=  primitive.ObjectIDFromHex(vars["studentId"]) 
	if err!=nil{
		utils.JSONError(w,"INVALID_ID",utils.GetCode("INVALID_ID"))
		return 
	}
	result,err:= users.GetStudent(studentId,staffId.(primitive.ObjectID))
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetAllStudents(w http.ResponseWriter, r *http.Request){
	staffId :=  r.Context().Value("id")

	result,err:= users.GetAllStudents(staffId.(primitive.ObjectID))
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	staffId :=  r.Context().Value("id")
	studentId,err :=  primitive.ObjectIDFromHex(vars["studentId"]) 
	if err!=nil{
		utils.JSONError(w,"INVALID_ID",utils.GetCode("INVALID_ID"))
		return 
	}
	result,err:= users.DeleteStudent(studentId,staffId.(primitive.ObjectID))
	if err!=nil{
	utils.JSONError(w,err.Error(),utils.GetCode(err.Error()))
	return   
	}
	res, _ :=json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

