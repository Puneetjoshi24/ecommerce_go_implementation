package service

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"models"
)

type (
	// UserService represents the service for operating on the User resource
	UserService struct {
		session *mgo.Session
	}
)

//NewUserService creating a new UserService instance
func NewUserService(s *mgo.Session) *UserService {
	return &UserService{s}
}

// GetUser retrieves an individual user resource
func (us UserService) GetUser(id string) (models.User, bool) {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return models.User{}, true
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.User{}

	// Fetch user
	if err := us.session.DB("eCommerce").C("users").FindId(oid).One(&u); err != nil {
		return models.User{}, true
	}

	return u, false
}

// GetAllUsers retrieves all user resources
func (us UserService) GetAllUsers() ([]models.User, bool) {

	// Stub user
	//u := models.User{}
	var users []models.User
	// Fetch users
	if err := us.session.DB("eCommerce").C("users").Find(nil).All(&users); err != nil {
		return []models.User{}, true
	}

	return users, false
}

// CreateUser creates a new user resource
func (us UserService) CreateUser(user models.User) (models.User, bool) {

	// Add an Id
	user.ID = bson.NewObjectId()
	c := us.session.DB("eCommerce").C("users")
	// Write the user to mongo
	err := c.Insert(user)
	if err != nil {
		panic(err)
	}

	return user, false
}

// RemoveUser removes an existing user resource
func (us UserService) RemoveUser(id string) bool {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return true
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := us.session.DB("eCommerce").C("users").RemoveId(oid); err != nil {
		return true
	}

	return false
}
