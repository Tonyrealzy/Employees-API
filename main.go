package main

import (
	"context"
	"log"
	"my-crud-project/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading.env file")
	}
	log.Println("env file loaded successfully")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping error: ", err)
	}
	log.Println("connected to mongodb successfully")
}

func main() {
	// close the connection when the main function exits
	defer mongoClient.Disconnect(context.Background())

	collection := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	// create empoyee service
	employeeService := usecase.EmployeeService{MongoCollection: collection}

	r := mux.NewRouter()
	r.HandleFunc("/create-employee", employeeService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/get-employee/{id}", employeeService.GetEmployeeById).Methods(http.MethodGet)
	r.HandleFunc("/get-employees", employeeService.GetAllEmployees).Methods(http.MethodGet)
	r.HandleFunc("/update-employee/{id}", employeeService.UpdateEmployeeById).Methods(http.MethodPut)
	r.HandleFunc("/delete-employee/{id}", employeeService.DeleteEmployeeById).Methods(http.MethodDelete)
	r.HandleFunc("/delete-employees", employeeService.DeleteAllEmployees).Methods(http.MethodDelete)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", r)
}
