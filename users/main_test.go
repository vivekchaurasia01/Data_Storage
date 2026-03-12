package main

import (
	"net/mail"
	"reflect"
	"testing"
)

func TestAddUser (t *testing.T){

	testManager := NewManager()
	
	// Simply we assign testparamter in a viriable for simplicity (helps in validation)
	testFirstName := "test"
	testLastName := "UserMan"
	testEmail,err := mail.ParseAddress("VivekChaurasia01@gmail.com")
	if err != nil{
		t.Fatalf("error while parsing the email : %v",err)
	}

	// Add idiots..
	err = testManager.AddUser(testFirstName,testLastName,testEmail.String())
	if err != nil{
		t.Fatalf("error while adding the user : %v",err)
	}

	// First lets make sure we have only one idiot at a time..
	if len(testManager.users) != 1{
		t.Errorf("bad test manager user count ,wanted:%d,got :%d",1,len(testManager.users))
		if len(testManager.users) < 1{
			t.Fatal()
		}
	}
	
	// Now we have to check expected user with our provided data..
	expectedUser := User{
		FirstName: testFirstName,
		LastName: testLastName,
		email: *testEmail,  // parseAddress return pointer to the email address we meed to dereference it ,cuz we are not storing pointers.
	}

	foundUser := testManager.users[0]

	// Now lets just comapre our found idiot is same as expected idiot.
	if !reflect.DeepEqual(expectedUser,foundUser){
		t.Errorf("added user data is not correct,wanted :%+v\n,got : %+v",expectedUser,foundUser) // %+v ---> Prints struct field names along with their values
	}

}

func TestAddUserInvalidEmail(t *testing.T){    //testing.T is the test context object.
	testManager := NewManager()

	testFirstName := "test"
	testLastName := "Userman"
	testEmail := "vivek"

	// Lets add user with this bad Email...
	err := testManager.AddUser(testFirstName,testLastName,testEmail)
	if err == nil{    //AddUser accepted invalid email.
		t.Error("no error returned for invalid email")
	}else{
		expectedErr := "invalid email: vivek"
		if err.Error() != expectedErr {   // error is interface : This converts the error object into its textual message.
			t.Errorf("bad error text,wanted :%s,got :%s",expectedErr,err)
		}
	}
	if len(testManager.users) >0 {
		t.Errorf("bad test manager user count, wanted :%d,got :%d",0,len(testManager.users))
	}
}
