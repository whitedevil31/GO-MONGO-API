package routes

import (
	"github.com/gorilla/mux"
	"github.com/whitedevil31/go-mongo-api/pkg/controllers/projectController"
	"github.com/whitedevil31/go-mongo-api/pkg/controllers/userController"
	"github.com/whitedevil31/go-mongo-api/pkg/utils"
)

var RegisterStudentRoutes = func(router *mux.Router){
	mainRouter := router.PathPrefix("/api").Subrouter()
	studentRoutes := mainRouter.PathPrefix("/student").Subrouter()    
	staffRoutes := mainRouter.PathPrefix("/staff").Subrouter()
	studentRoutes.Use(utils.AuthMiddleware)
	staffRoutes.Use(utils.StaffMiddleware)
	
	//no auth 
	mainRouter.HandleFunc("/student/login",userController.StudentLogin).Methods("POST")   //student login
	mainRouter.HandleFunc("/staff/signup",userController.StaffSignUp).Methods("POST")  //staff signup
	mainRouter.HandleFunc("/staff/login",userController.StaffLogin).Methods("POST")   //staff login

	staffRoutes.HandleFunc("/get/student/all",userController.GetAllStudents).Methods("GET") //staff access get all his/her students 
	staffRoutes.HandleFunc("/create/student",userController.StudentSignUp).Methods("POST")   
	staffRoutes.HandleFunc("/get/student/{studentId}",userController.GetStudent).Methods("GET")   //staff access get his/her student
	staffRoutes.HandleFunc("/delete/student/{studentId}",userController.DeleteStudent).Methods("DELETE")   //staff access to delete any student


	studentRoutes.HandleFunc("/get/my_profile",userController.GetStudentProfile).Methods("GET")     
	studentRoutes.HandleFunc("/add_project",projectController.AddProject).Methods("POST")
	studentRoutes.HandleFunc("/get_my_projects",projectController.GetProjects).Methods("GET")
	studentRoutes.HandleFunc("/get_project/{projectId}",projectController.GetProject).Methods("GET")
	studentRoutes.HandleFunc("/delete_project/{projectId}",projectController.DeleteProject).Methods("DELETE")
	studentRoutes.HandleFunc("/update_project/{projectId}",projectController.UpdateProject).Methods("PATCH")

}	