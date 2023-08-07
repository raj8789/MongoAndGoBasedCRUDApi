package main

import (
	"controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main(){
	r:=httprouter.New()
	uc:=controllers.NewUSerController(getSession())
	r.GET("/user/get/:id",uc.GetUser)
	r.POST("/user/post",uc.CreateUser)
	r.DELETE("/user/delete/:id",uc.DeleteUser)
	r.PUT("/user/update/:id",uc.UpdateUser)
	http.ListenAndServe("localhost:8080",r)
}
func getSession() (*mongo.Client,context.Context){
	uri := "mongodb://localhost:27017"
	// Specify the Stable API version in the ClientOptions object
	// Specify the Stable API version and append options in the ClientOptions object
	serverAPI := options.ServerAPI(options.ServerAPIVersion1).SetStrict(true).SetDeprecationErrors(true)

	// Pass in the URI and the ClientOptions to the Client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 700*time.Minute)
	return client,ctx

}