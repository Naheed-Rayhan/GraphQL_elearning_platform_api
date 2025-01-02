package resolver

import (
	"log"

	"github.com/Naheed-Rayhan/graphql-api/entities"
	"github.com/Naheed-Rayhan/graphql-api/infrastructure/database"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	repo *database.CourseRepository
}

func InitResolver(repo *database.CourseRepository) *Resolver {
	return &Resolver{repo: repo}
}

func GetUserByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		repo := &database.CourseRepository{}
		return repo.GetUserByID(id)		
	}
	return nil, nil
}


func GetAllUsers(p graphql.ResolveParams) (interface{}, error) {
	repo := &database.CourseRepository{}
	return repo.GetAllUsers()
}

func (r Resolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	

	user := entities.User{
        FirstName: p.Args["first_name"].(string),
		LastName: p.Args["last_name"].(string),
		Email: p.Args["email"].(string),
		Password: p.Args["password"].(string),
		Role: p.Args["role"].(string),
		Bio: p.Args["bio"].(string),
    }
    
	log.Println("from resolver ",user)


	
	createdUser, err := r.repo.CreateUser(user)
    if err != nil {
        return nil, err  // Return the error if something goes wrong
    }

    return createdUser, nil  // Return the created user and nil error
}