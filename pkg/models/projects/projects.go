package projects

import (
	"context"
	"errors"
	"fmt"

	//"reflect"

	"github.com/whitedevil31/go-mongo-api/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var collection *mongo.Collection
type Project struct {
	ID       primitive.ObjectID 	`bson:"_id" json:"_id" `
	ProjectName  string         `bson:"projectName" json:"projectName" validate:"required,min=5,max=25" `
	ProjectDescription string         `bson:"projectDescription" json:"projectDescription" validate:"required,min=10,max=50"`
	ProjectProgress   int           	`bson:"projectProgress" json:"projectProgress" validate:"required"`
	StudentId primitive.ObjectID  `bson:"studentId" json:"studentId" `

}

type MessageResponse struct{
	Message string `json:"message" bson:"message"`
}
type GetMyProjectsResponse struct{
	Result []primitive.M `bson:"results"  json:"results"`
}
type GetMyProjectResponse struct{
	Result primitive.M `bson:"result"  json:"result"`
}
func init(){
	client = config.GetDB()
	}

func AddProject(project *Project,id primitive.ObjectID) (MessageResponse,error){
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("projects")
	
	fmt.Println(collection)
	res:=MessageResponse{}
	_,err := collection.InsertOne(context.Background(),&Project{
	ID:        primitive.NewObjectID(),  
	ProjectName: project.ProjectName,
	ProjectDescription: project.ProjectDescription,
	ProjectProgress: project.ProjectProgress,
	StudentId: id,

	})
	
	if err!=nil{
		fmt.Println(err)
		return res,errors.New("SOMETHING_WENT_WRONG")
	}
	
	return MessageResponse{Message:"Project added successfully!"},nil
}

func GetProjects(id primitive.ObjectID) (GetMyProjectsResponse,error){
	
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("projects")

	cursor,getProjectsError:=collection.Find(context.TODO(),bson.D{{Key: "studentId",Value: id}})
	res:=GetMyProjectsResponse{}
	var result []primitive.M
	if getProjectsError!=nil{

		if getProjectsError==mongo.ErrNoDocuments{
			return res,errors.New("RESULT_NOT_FOUND")
		 }

		return res,errors.New("SOMETHING_WENT_WRONG")
	}else{
		 getProjectsError = cursor.All(context.TODO(),&result)

		 
		if getProjectsError!=nil{
			
return res,errors.New("SOMETHING_WENT_WRONG")
		}
	
		
	}
	res.Result=result
	return res,nil
}



func GetProject(id primitive.ObjectID,projectId primitive.ObjectID) (GetMyProjectResponse,error){
	
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("projects")
	var result bson.M
	res :=GetMyProjectResponse{}	
	
	getProjectError:=collection.FindOne(context.TODO(),bson.D{{Key: "studentId",Value: id},{Key: "_id",Value: projectId}}).Decode(&result)
	if getProjectError!=nil{	
		if getProjectError==mongo.ErrNoDocuments{
			return res,errors.New("RESULT_NOT_FOUND")
		}
	
		return res,errors.New("SOMETHING_WENT_WRONG")
	}

	res.Result=result
	return res,nil
}


func DeleteProject(id primitive.ObjectID,projectId primitive.ObjectID) (MessageResponse,error){
	
	c:=config.GetDB()
	collection := c.Database("go-api").Collection("projects") 
	res:=MessageResponse{}
	curr,deleteProjectError:=collection.DeleteOne(context.TODO(),bson.D{{Key: "studentId",Value: id},{Key: "_id",Value: projectId}})
	if deleteProjectError!=nil{	

		return res,errors.New("SOMETHING_WENT_WRONG")
	}
	if curr.DeletedCount==0{
		return res,errors.New("RESULT_NOT_FOUND")
	}

	res.Message="Project deleted successfully !"

	return res,nil
}


func UpdateProject(id primitive.ObjectID,projectId primitive.ObjectID,updateProject *Project) (MessageResponse,error){

	c:=config.GetDB()
	collection := c.Database("go-api").Collection("projects") 
	res:=MessageResponse{}
//	var getProject bson.M
 updateItem:= bson.M{}
//fmt.Println(updateProject)
if updateProject.ProjectDescription!=""{
	updateItem["projectDescription"]=updateProject.ProjectDescription
}
if updateProject.ProjectName!=""{
	updateItem["projectName"]=updateProject.ProjectName
}
if updateProject.ProjectProgress!=0{
	updateItem["projectProgress"]=updateProject.ProjectProgress
}
	update := bson.M{
		"$set": updateItem }
	fmt.Println(updateItem)
	//if result
	filter := bson.D{{Key: "studentId",Value: id},{Key: "_id",Value: projectId}}
	curr,updateProjectError:=collection.UpdateOne(context.TODO(),filter,update)
	if updateProjectError!=nil{	
fmt.Print(updateProjectError)
		return res,errors.New("SOMETHING_WENT_WRONG")
	}
	if curr.MatchedCount==0{
		return res,errors.New("RESULT_NOT_FOUND")
	}

	res.Message="Project Updated successfully !"

	return res,nil
}