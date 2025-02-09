package repository

import (
	"context"
	"log"
	"my-crud-project/models"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://tonyrealzy:AxQQubTrrCOUzwtj@tonyrealzy.alipl.mongodb.net/?retryWrites=true&w=majority&appName=Tonyrealzy"))
	if err != nil {
		log.Fatal("an error occurred while connecting to mongodb", err)
	}
	log.Println("connected to mongodb successfully")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("ping successful")
	return mongoTestClient
}

func TestMongoDbOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	employee1 := uuid.New().String()
	employee2 := uuid.New().String()

	// connect to the collection
	collection := mongoTestClient.Database("CRUDdb").Collection("employee_test")
	employeeRepo := EmployeeRepo{MongoCollection: collection}

	//inserting dummy data
	t.Run("Insert employee 1", func(t *testing.T) {
		employee := models.Employee{
			EmployeeID: employee1,
			Name:       "Tony Stark",
			Department: "Engineering",
		}
		result, err := employeeRepo.InsertEmployee(&employee)
		if err != nil {
			t.Fatal("error inserting employee 1", err)
		}
		t.Log("inserted employee 1", result)
	})
	t.Run("Insert employee 2", func(t *testing.T) {
		employee := models.Employee{
			EmployeeID: employee2,
			Name:       "Berrise Obi",
			Department: "Arts",
		}
		result, err := employeeRepo.InsertEmployee(&employee)
		if err != nil {
			t.Fatal("error inserting employee 2", err)
		}
		t.Log("inserted employee 2", result)
	})
	t.Run("Get employee 1", func(t *testing.T) {
		result, err := employeeRepo.FindEmployeeById(employee1)
		if err != nil {
			t.Fatal("error getting employee 1", err)
		}
		t.Log("got employee 1", result)
	})
	t.Run("Get all employees", func(t *testing.T) {
		result, err := employeeRepo.FindAllEmployees()
		if err != nil {
			t.Fatal("error getting all employees", err)
		}
		t.Log("got all employees", result)
	})
	t.Run("Update employee 1", func(t *testing.T) {
		employee := models.Employee{
			EmployeeID: employee1,
			Name:       "Tony Stark a.k.a Iron Man",
			Department: "Engineering",
		}
		result, err := employeeRepo.UpdateEmployeeById(employee1, &employee)
		if err != nil {
			t.Fatal("error updating employee", err)
		}
		t.Log("updated employee", result)
	})
	t.Run("Get employee 1 after update", func(t *testing.T) {
		result, err := employeeRepo.FindEmployeeById(employee1)
		if err != nil {
			t.Fatal("error getting employee 1 after update", err)
		}
		t.Log("got employee 1 after update", result)
	})
	t.Run("Delete employee 1", func(t *testing.T) {
		result, err := employeeRepo.DeleteEmployeeById(employee1)
		if err != nil {
			t.Fatal("error deleting employee 1", err)
		}
		t.Log("deleted employee 1", result)
	})
	t.Run("Delete all employees", func(t *testing.T) {
		result, err := employeeRepo.DeleteAllEmployees()
		if err != nil {
			t.Fatal("error deleting all employees", err)
		}
		t.Log("deleted all employees", result)
	})
}
