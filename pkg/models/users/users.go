package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/whitedevil31/go-mongo-api/pkg/config"
	"github.com/whitedevil31/go-mongo-api/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


var client *mongo.Client
var collection *mongo.Collection
type Student struct {
	ID       primitive.ObjectID 	`bson:"_id" json:"_id" `
	Name  string         `bson:"name" json:"name" validate:"required,alpha" `
	Password string         `bson:"password" json:"password" validate:"required,min=8,max=15"`
	Email   string           	`bson:"email" json:"email" validate:"required,email"`
	CreatedBy primitive.ObjectID  `bson:"createdBy" json:"createdBy" `
}
type Staff struct{
	ID  primitive.ObjectID 	`bson:"_id" json:"_id" `
	StaffName    string   `bson:"staffName" json:"staffName" validate:"required,alpha" `
	StaffPassword  string `bson:"staffPassword" json:"staffPassword" validate:"required,min=8,max=15"`
	Department   string  `bson:"department" json:"department" validate:"required,alpha" `
	StaffEmail   string           	`bson:"staffEmail" json:"staffEmail" validate:"required,email"`
}
type Claims struct {
	ID primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}
type SignUpResponse struct {
	Token string `json:"token" bson:"token`  
}
type GetStudentResponse struct {
	Result primitive.M `bson:"result"  json:"result"`
}
type MessageResponse struct{
	Message string `json:"message" bson:"message"`
}
func init(){
client = config.GetDB()
}


func  StaffSignUp(staff *Staff) (SignUpResponse,error) {
	collection := client.Database("go-api").Collection("staffs")
	var res bson.M
	err:=collection.FindOne(context.TODO(),bson.D{{Key:"staffEmail",Value:staff.StaffEmail}}).Decode(&res)
	errorResponse:= SignUpResponse{}
	if res!=nil{
	return errorResponse,errors.New("EMAIL_IN_USE")

	}
	password := []byte(staff.StaffPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
       return errorResponse,errors.New("INTERNAL_SERVER_ERROR")
    }
	id := primitive.NewObjectID()
	//expiration time is 24 hourss
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{	
			ExpiresAt:expirationTime.Unix(), 
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.ViperEnvVariable("JWT_SECRET_STAFF"))
	tokenString, err := token.SignedString(key)
	if err!=nil{
		fmt.Println(err)
	}
	_,err = collection.InsertOne(context.Background(),&Staff{
		ID:       id,
		StaffName:staff.StaffName,
		StaffPassword: string(hashedPassword),
		Department:staff.Department,
		StaffEmail: staff.StaffEmail,
	})
    if err != nil {
		return errorResponse,errors.New("INTERNAL_SERVER_ERROR")
    } 
	 tokenResponse :=SignUpResponse{Token:tokenString}
		return tokenResponse,nil
}
func  StaffLogin(staff *utils.StaffLoginCredentials)  (utils.LoginResponse,error){
	collection := client.Database("go-api").Collection("staffs")
	errorResponse:= utils.LoginResponse{}
	var res Staff
	findUserError:=collection.FindOne(context.TODO(),bson.D{{Key:"staffEmail",Value:staff.StaffEmail}}).Decode(&res)
	fmt.Println("FIND USE ERROR")
	passNotMatch := bcrypt.CompareHashAndPassword([]byte(res.StaffPassword), []byte(staff.StaffPassword))
	if findUserError!=nil{
		return errorResponse,errors.New("USER_NOT_EXIST")
	}
	if passNotMatch != nil{
		return errorResponse,errors.New("INCORRECT_PASSWORD")
	}
	

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: res.ID,
		StandardClaims: jwt.StandardClaims{	
			ExpiresAt:expirationTime.Unix(), 
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.ViperEnvVariable("JWT_SECRET_STAFF"))
	tokenString, _ := token.SignedString(key)


	 tokenResponse :=utils.LoginResponse{Token:tokenString}
	 return tokenResponse,nil;
}
func  StudentSignUp(student *Student,staffId primitive.ObjectID) (SignUpResponse,error) {
	collection := client.Database("go-api").Collection("students")
	var res bson.M
	err:=collection.FindOne(context.TODO(),bson.D{{Key:"email",Value:student.Email}}).Decode(&res)
	errorResponse:= SignUpResponse{}
	if res!=nil{
	return errorResponse,errors.New("EMAIL_IN_USE")

	}
	password := []byte(student.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
       return errorResponse,errors.New("INTERNAL_SERVER_ERROR")
    }
	id := primitive.NewObjectID()
	//expiration time is 24 hourss
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{	
			ExpiresAt:expirationTime.Unix(), 
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.ViperEnvVariable("JWT_SECRET"))
	tokenString, err := token.SignedString(key)
	if err!=nil{
		fmt.Println(err)
	}
	_,err = collection.InsertOne(context.Background(),&Student{
		ID:       id,
		Name:student.Name,
		Password: string(hashedPassword),
		Email: student.Email,
		CreatedBy: staffId,
	})
    if err != nil {
		return errorResponse,errors.New("INTERNAL_SERVER_ERROR")
    } 
	 tokenResponse :=SignUpResponse{Token:tokenString}
		return tokenResponse,nil
}

func  StudentLogin(student *utils.StudentLoginCredentials)  (utils.LoginResponse,error){
	collection := client.Database("go-api").Collection("students")
	errorResponse:= utils.LoginResponse{}
	var res Student
	findUserError:=collection.FindOne(context.TODO(),bson.D{{Key:"email",Value:student.Email}}).Decode(&res)
	fmt.Println("FIND USE ERROR")
	passNotMatch := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(student.Password))
	if findUserError!=nil{
		return errorResponse,errors.New("USER_NOT_EXIST")
	}
	if passNotMatch != nil{
		return errorResponse,errors.New("INCORRECT_PASSWORD")
	}
	

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: res.ID,
		StandardClaims: jwt.StandardClaims{	
			ExpiresAt:expirationTime.Unix(), 
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.ViperEnvVariable("JWT_SECRET"))
	tokenString, _ := token.SignedString(key)


	 tokenResponse :=utils.LoginResponse{Token:tokenString}
	 return tokenResponse,nil;
}

func GetStudentProfile(id primitive.ObjectID) (Student,error){
	collection := client.Database("go-api").Collection("students")
	var result Student
	err :=collection.FindOne(context.TODO(),bson.M{"_id": id}).Decode(&result)
   if err != nil{
	return result,errors.New("SOMETHING_WENT_WRONG")
   }
return result,nil
}
func GetStudent(id primitive.ObjectID,staffId primitive.ObjectID) (GetStudentResponse,error){
	
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("students")
	var result bson.M
	res :=GetStudentResponse{}	
	
	getProjectError:=collection.FindOne(context.TODO(),bson.D{{Key: "_id",Value: id},{Key: "createdBy",Value: staffId}}).Decode(&result)
	if getProjectError!=nil{	
		if getProjectError==mongo.ErrNoDocuments{
			return res,errors.New("RESULT_NOT_FOUND")
		}
	
		return res,errors.New("SOMETHING_WENT_WRONG")
	}

	res.Result=result
	return res,nil
}



func GetAllStudents(staffId primitive.ObjectID) ([]Student,error){
	var students []Student
	collection := client.Database("go-api").Collection("students")
	curr,err := collection.Find(context.Background(),bson.D{{Key: "createdBy",Value: staffId}})
    if err != nil {
   fmt.Println(err)
    } else {
      for curr.Next(context.Background()){
		var student Student

		err := curr.Decode(&student)
		
		if err!=nil{
			fmt.Println(err)
		}
	
		students = append(students, student)
	  }
    }

	return students,nil
}


func DeleteStudent(id primitive.ObjectID,staffId primitive.ObjectID) (MessageResponse,error){
	
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("students") 
	res:=MessageResponse{}
	curr,deleteStudent:=collection.DeleteOne(context.TODO(),bson.D{{Key: "_id",Value: id},{Key: "createdBy",Value: staffId}})
	if deleteStudent!=nil{	

		return res,errors.New("SOMETHING_WENT_WRONG")
	}
	if curr.DeletedCount==0{
		return res,errors.New("STUDENT_NOT_FOUND")
	}

	res.Message="Student account deleted successfully !"

	return res,nil
}
