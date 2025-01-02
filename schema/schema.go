package schema

import (
	"context"
	"errors"
	"log"

	"github.com/Naheed-Rayhan/graphql-api/entities"
	"github.com/Naheed-Rayhan/graphql-api/usecases"
	"github.com/graphql-go/graphql"
)






var user = graphql.NewObject(
	graphql.ObjectConfig{
		Name : "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{ Type: graphql.Int},
			"first_name": &graphql.Field{ Type: graphql.String},
			"last_name": &graphql.Field{ Type: graphql.String},
			"email": &graphql.Field{ Type: graphql.String},
			"password": &graphql.Field{ Type: graphql.String},
			"role": &graphql.Field{ Type: graphql.String},
			"bio": &graphql.Field{ Type: graphql.String},
		},
	},
)


// var course = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name : "Course",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{ Type: graphql.Int},
// 			"title": &graphql.Field{ Type: graphql.String},
// 			"description": &graphql.Field{ Type: graphql.String},
// 			"duration": &graphql.Field{ Type: graphql.String},
// 			"price": &graphql.Field{ Type: graphql.Float},
// 			"instructor": &graphql.Field{ Type: graphql.String},
// 			"category": &graphql.Field{ Type: graphql.String},
// 		},
// 	},
// )

// var enrollment = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name : "Enrollment",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{ Type: graphql.Int},
// 			"user_id": &graphql.Field{ Type: graphql.Int},
// 			"course_id": &graphql.Field{ Type: graphql.Int},
// 			"completed": &graphql.Field{ Type: graphql.Boolean},
// 		},
// 	},
// )
	


// var lesson = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name : "Lesson",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{ Type: graphql.Int},
// 			"course_id": &graphql.Field{ Type: graphql.Int},
// 			"title": &graphql.Field{ Type: graphql.String},
// 			"content": &graphql.Field{ Type: graphql.String},
// 			"video_url": &graphql.Field{ Type: graphql.String},
// 			"order": &graphql.Field{ Type: graphql.Int},
// 		},
// 	},
// )


// var progress = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name : "Progress",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{ Type: graphql.Int},
// 			"enrollment_id": &graphql.Field{ Type: graphql.Int},
// 			"lesson_id": &graphql.Field{ Type: graphql.Int},
// 			"completed": &graphql.Field{ Type: graphql.Boolean},
// 		},
// 	},
// )

// var review = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name : "Review",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{ Type: graphql.Int},
// 			"course_id": &graphql.Field{ Type: graphql.Int},
// 			"user_id": &graphql.Field{ Type: graphql.Int},
// 			"rating": &graphql.Field{ Type: graphql.Int},
// 			"comment": &graphql.Field{ Type: graphql.String},
// 		},
// 	},
// )


var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: user,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{ Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Retrieve the CourseRepository from the context
    				courseRepository, ok := p.Context.Value("courseRepository").(*CourseRepository)
    				if !ok {
        				return nil, errors.New("unable to retrieve course repository from context")
    				}

					id, ok := p.Args["id"].(int)
					if ok {
						user, err := courseRepository.CourseRepo.GetUserByID(id)
						if err != nil {
							return nil, err  // Return the error if something goes wrong
						}
						return user, nil
					}
					return nil, nil
				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(user),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Retrieve the CourseRepository from the context
					courseRepository, ok := p.Context.Value("courseRepository").(*CourseRepository)
					if !ok {
						return nil, errors.New("unable to retrieve course repository from context")
					}
					
					users, err := courseRepository.CourseRepo.GetAllUsers()
					if err != nil {
						return nil, err  // Return the error if something goes wrong
					}
					return users, nil
				},
			},
			

		},
	},
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: user,
				Args: graphql.FieldConfigArgument{
					"first_name": &graphql.ArgumentConfig{ Type: graphql.NewNonNull(graphql.String)},
					"last_name": &graphql.ArgumentConfig{ Type: graphql.NewNonNull(graphql.String)},
					"email": &graphql.ArgumentConfig{ Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{ Type: graphql.NewNonNull(graphql.String)},
					"role": &graphql.ArgumentConfig{ Type: graphql.NewNonNull(graphql.String)},
					"bio": &graphql.ArgumentConfig{ Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Retrieve the CourseRepository from the context
    				courseRepository, ok := p.Context.Value("courseRepository").(*CourseRepository)
    				if !ok {
        				return nil, errors.New("unable to retrieve course repository from context")
    				}
					
					
					user := entities.User{
						FirstName: p.Args["first_name"].(string),
						LastName: p.Args["last_name"].(string),
						Email: p.Args["email"].(string),
						Password: p.Args["password"].(string),
						Role: p.Args["role"].(string),
						Bio: p.Args["bio"].(string),
					}
					
					createdUser, err := courseRepository.CourseRepo.CreateUser(user)
					if err != nil {
						return nil, err  // Return the error if something goes wrong
					}
				
					return createdUser, nil  // Return
				},
			},
		},
	},
)


var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: query,
		Mutation: mutation,

	},
)

type CourseRepository struct {
	CourseRepo *usecases.CourseUseCase
}


func ExecuteQuery(query string ,uc *usecases.CourseUseCase) *graphql.Result {
    log.Println("Executing query:", query)

	CourseRepository := &CourseRepository{CourseRepo: uc}

 
   // Create a context with the repository
   ctx := context.WithValue(context.Background(), "courseRepository", CourseRepository)

   result := graphql.Do(graphql.Params{
	   Schema:        Schema,
	   RequestString: query,
	   Context:       ctx,
   })
   if len(result.Errors) > 0 {
	   log.Printf("errors inside schema: %v", result.Errors)
   }
   return result
}