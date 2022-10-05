package utils

import (
	"context"
	"encoding/json"

	//"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/whitedevil31/go-mongo-api/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type Claims struct {
	ID primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}
type Student struct {
	ID       primitive.ObjectID 	`bson:"_id" json:"_id" `
	Name  string         `bson:"name" json:"name" validate:"required,alpha" `
	Password string         `bson:"password" json:"password" validate:"required,min=8,max=15"`
	Email   string           	`bson:"email" json:"email" validate:"required,email"`
}

var c *mongo.Client
var collection *mongo.Collection

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}

func AuthMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("authorization")
		if authorization ==""{
			JSONError(w,"TOKEN_NOT_FOUND",GetCode("TOKEN_NOT_FOUND"))
			return 
		}
		token := strings.Split(authorization, " ")[1]
	if token==""{
		JSONError(w,"TOKEN_NOT_FOUND",GetCode("TOKEN_NOT_FOUND"))
		return 
	}
		claims := &Claims{}
		key := config.ViperEnvVariable("JWT_SECRET")
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
if err != nil{
	JSONError(w,"TOKEN_ERROR",GetCode("TOKEN_ERROR"))
	return
}	 
	var result Student
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("students")
	
	err=collection.FindOne(context.Background(),bson.D{{Key:"_id",Value:claims.ID}}).Decode(&result)
	
	 if err!=nil{
		JSONError(w,"USER_NOT_EXIST",GetCode("USER_NOT_EXIST"))
		return
	 }	
	ctx := context.WithValue(r.Context(),"id",result.ID)
	h.ServeHTTP(w,r.WithContext(ctx))
    })
}
func StaffMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("authorization")
		if authorization ==""{
			JSONError(w,"TOKEN_NOT_FOUND",GetCode("TOKEN_NOT_FOUND"))
			return 
		}
		token := strings.Split(authorization, " ")[1]
	if token==""{
		JSONError(w,"TOKEN_NOT_FOUND",GetCode("TOKEN_NOT_FOUND"))
		return 
	}
		claims := &Claims{}
		key := config.ViperEnvVariable("JWT_SECRET_STAFF")
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
if err != nil{
	JSONError(w,"TOKEN_ERROR",GetCode("TOKEN_ERROR"))
	return
}	 
	var result Student
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("staffs")
	
	err=collection.FindOne(context.Background(),bson.D{{Key:"_id",Value:claims.ID}}).Decode(&result)
	
	 if err!=nil{
		JSONError(w,"USER_NOT_EXIST",GetCode("USER_NOT_EXIST"))
		return
	 }	
	ctx := context.WithValue(r.Context(),"id",result.ID)
	h.ServeHTTP(w,r.WithContext(ctx))
    })
}
type ErrorFormat struct{
	Message string  `json:"message" bson:"message"`
}
func JSONError(w http.ResponseWriter, err string, code int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.WriteHeader(code)
	t :=ErrorFormat{Message:err}
    json.NewEncoder(w).Encode(t)
}
func GetCode(err string) int {
   switch err {
       case "EMAIL_IN_USE":
       return  http.StatusConflict
       case "VALIDATION_FAILED":
		return http.StatusBadRequest
	case "INTERNAL_SERVER_ERROR":
		return http.StatusInternalServerError
	case "SOMETHING_WENT_WRONG":
		return http.StatusInternalServerError
	case "USER_NOT_EXIST":
		return http.StatusNotFound
	case "INCORRECT_PASSWORD":
		return http.StatusNotFound
	case "UNAUTHORISED_ACCESS":
		return http.StatusUnauthorized
	case "INVALID_LOGIN":
		return http.StatusUnauthorized
	case "TOKEN_ERROR":
		return http.StatusFailedDependency	
	case "INVALID_ID":
		return http.StatusBadRequest
	case "RESULT_NOT_FOUND":
		return http.StatusNotFound
	case "STUDENT_NOT_FOUND":
		return http.StatusNotFound
	case "TOKEN_NOT_FOUND":
		return http.StatusNotAcceptable
	default: 
		return 500
   }  

}

type StudentLoginCredentials struct{
	Email string     `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}
type StaffLoginCredentials struct{
	StaffEmail string     `bson:"staffEmail" json:"staffEmail" validate:"required,email"`
	StaffPassword string `bson:"staffPassword" json:"staffPassword" validate:"required"`
}
type LoginResponse struct{
	Token string  `json:"token"`
}