package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"models"
	"net/http"
)
type UserController struct{
	client *mongo.Client
	context context.Context
}
func NewUSerController(client *mongo.Client, context context.Context) *UserController{
	uc:=&UserController{client,context}
	return uc
}

func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request ,param httprouter.Params){
	id,_:=strconv.Atoi(param.ByName("id"))
	filter :=bson.D{{"_id", id}}
	result,err:=userController.client.Database("mongo-database").Collection("user").Find(context.TODO(),filter)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound) 
		return
	}else{
		fmt.Println("Data Ftech From Database")
	}
	var results []models.User
	result.All(context.TODO(),&results)
	var ans string
	for _, result := range results {
		  res, _ := json.Marshal(result)
		  ans=string(res)
	}
	if ans==""{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"User is Not Present")
	}else{
		w.Header().Set("content-Type","application/json")
	    w.WriteHeader(http.StatusOK)
	    fmt.Fprintf(w,"%s\n",ans)
	}
}

 func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request ,_ httprouter.Params){
	user:=models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	result,err:=userController.client.Database("mongo-database").Collection("user").InsertOne(userController.context,user)
	if err!=nil{
		fmt.Println(user.Name)
		fmt.Println(err)
		panic(err)
	}else{
		fmt.Println(result.InsertedID)
	}
	userJson,err:=json.Marshal(user)
	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",userJson)
 }

 func(userController UserController) DeleteUser(w http.ResponseWriter, r *http.Request ,param httprouter.Params){
	id,_:=strconv.Atoi(param.ByName("id"))
	fmt.Printf("Id of user=%v",id)
	condition := bson.D{{"_id", id}}
	var singleResult bson.M
	err:=userController.client.Database("mongo-database").Collection("user").FindOneAndDelete(context.TODO(),condition).Decode(&singleResult)
	if err!=nil{
		fmt.Println("**********************")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"User is Not Present")
		return
	}else{
		fmt.Println(singleResult)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted User",singleResult,"\n")
 }
 func(userController UserController) UpdateUser(w http.ResponseWriter, r *http.Request ,param httprouter.Params){
	id,_:=strconv.Atoi(param.ByName("id"))
	fmt.Printf("Id of user=%v\n",id)
	newData := bson.D{}
	if  r.Header.Get("age") != ""{
		age,_:=strconv.Atoi(r.Header.Get("age"))
		newData = append(newData, bson.E{"age", age})
	}
	if  r.Header.Get("gender") != ""{
		gender:=r.Header.Get("gender")
		newData = append(newData, bson.E{"gender", gender})
	}
	if  r.Header.Get("name") != "" {
		name:=r.Header.Get("name")
		newData = append(newData, bson.E{"name", name})
	}
	filtercondition := bson.D{{"_id", id}}
    updatecondition :=  bson.D{{"$set", newData}}
	result, err := userController.client.Database("mongo-database").Collection("user").UpdateOne(context.TODO(),filtercondition,updatecondition)
	if err != nil {
		fmt.Fprintln(w,"Not ABle to Update Document In Collection")
	} else {
		if result.MatchedCount != 0 {
		  w.WriteHeader(http.StatusOK)
		  fmt.Fprintln(w,"Document Updated Successfully In Collection")
		} else{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w,"Document Not Found In Collection")
		}
	}
	fmt.Fprintln(w,"Updated User",result,"\n")
 }