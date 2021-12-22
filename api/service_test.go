package api

import (
	"errors"
	"log"
	"testing"

	"github.com/Ajmalazeem/models"
	"github.com/Ajmalazeem/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type MockRepository struct{
	mock.Mock
	store.TodoStore
}
func (mock *MockRepository)PostTodo(models.PostTodoRequest) error{
	args := mock.Called()
	// result := args.Get(0)
	// result.(models.PostTodoRequest)
	return args.Error(0)
}

func (mock *MockRepository)GetTodo(models.GetTodoRequest)(*models.PostTodoRequest,error){
	args := mock.Called()
	result := args.Get(0)
	return result.(*models.PostTodoRequest),args.Error(1)
}

func (mock *MockRepository)PutTodo(models.PutTodoRequest) error{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository)DeleteTodo(models.DeleteTodoRequest)error{
	args := mock.Called()
	return args.Error(0)
}


func TestPostTodo(t *testing.T) {
	mockrepo := new(MockRepository)
	post := models.PostTodoRequest{Id: "1",Todo: "work"}
	mockrepo.On("PostTodo").Return(nil)
	testservice := NewTodoService(mockrepo)
	result := testservice.PostTodo(post)
	mockrepo.AssertExpectations(t)
	//assert.Equal(t,"1",result.Id)
	//assert.Equal(t,"work",result.Todo)
	assert.Nil(t,result)

}

func TestPostTodoError (t *testing.T){
	mockRepo := new(MockRepository)
	post := models.PostTodoRequest{Id:"1", Todo : "work"}
	err := errors.New("invalid")
	mockRepo.On("PostTodo").Return(err)
	testService := NewTodoService(mockRepo)
	result := testService.PostTodo(post)

	mockRepo.AssertExpectations(t)
	log.Println(err)
	assert.NotNil(t,result)
}

func TestGetTodo(t *testing.T){
	mockRepo := new(MockRepository)
	get := models.GetTodoRequest{Id: "1"}
	post := models.PostTodoRequest{Id: "1", Todo: "work"}
	var err error
	mockRepo.On("GetTodo").Return(&post,err)
	testService := NewTodoService(mockRepo)
	result, err := testService.GetTodo(get)

	mockRepo.AssertExpectations(t)
	log.Println(err)
	assert.Equal(t,"1",result.Id)
	assert.Equal(t,"work",result.Todo)
	assert.Nil(t,err)
}

func TestPutTodo(t *testing.T){
	mockRepo := new(MockRepository)
	Put := models.PutTodoRequest{Id: "1",Todo: "work"}
	var err error
	mockRepo.On("PutTodo").Return(err)
	testService := NewTodoService(mockRepo)
	result := testService.PutTodo(Put)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, result)
}

func TestDeleteTodo (t *testing.T){
	mockRepo := new(MockRepository)
	delete := models.DeleteTodoRequest{Id: "1"}
	var err error
	mockRepo.On("DeleteTodo").Return(err)
	testService := NewTodoService(mockRepo)
	result := testService.DeleteTodo(delete)
	mockRepo.AssertExpectations(t)
	assert.Nil(t,result)
}
