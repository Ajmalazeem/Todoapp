package api

import (
	"testing"

	"github.com/magiconair/properties/assert"
)


func TestPostTodo(t *testing.T){
	assert := assert.new()

	// activate HTTP mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	


	
}



// func TestPostTodo(t *testing.T){
// 	for _, test := range PostTodoTests{
// 		if  Output := string("Id: " + test.id +","+"Todo: " + test.todo); Output!= test.expected{
// 			t.Errorf("output: %q not equal to expected: %q", Output,test.expected)
// 		}
		
// 	}
// }

