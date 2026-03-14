package main

import (
	"errors"
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

	// Now lets just compare our found idiot is same as expected idiot.
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
	if err == nil{
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
func TestAddUserEmptyUserName(t *testing.T){
	testManager := NewManager()

	testFirstName := ""
	testLastName := "Userman"
	testEmail,err := mail.ParseAddress("VivekChaurasia01@gmail.com")
	if err != nil{
		t.Fatalf("error while parsing the email : %v",err)
	}

	// Lets add user with this bad Email...
	err = testManager.AddUser(testFirstName,testLastName,testEmail.String())
	if err == nil{
		t.Error("no error returned for invalid firstname")
	}else{
		expectedErr := "invalid first name:\"\""
		if err.Error() != expectedErr {   // error is interface : This converts the error object into its textual message.
			t.Errorf("bad error text,wanted:%s,got:%s",expectedErr,err)
		}
	}
	if len(testManager.users) >0 {
		t.Errorf("bad test manager user count, wanted :%d,got :%d",0,len(testManager.users))
	}
}

func TestAddUserEmptyLastName(t *testing.T){
	testManager := NewManager()

	testFirstName := "test"
	testLastName := ""
	testEmail,err := mail.ParseAddress("VivekChaurasia01@gmail.com")
	if err != nil{
		t.Fatalf("error parsing test email address : %v",err)
	}

	// Lets add user with this bad Email...
	err = testManager.AddUser(testFirstName,testLastName,testEmail.String())
	if err == nil{    //AddUser accepted invalid lastname.
		t.Error("no error returned for invalid lastname")
	}else{ 
		expectedErr := "invalid last name:\"\""
		if err.Error() != expectedErr {   // error is interface : This converts the error object into its textual message.
			t.Errorf("bad error text,wanted:%s,got:%s",expectedErr,err)
		}
	}
	if len(testManager.users) >0 {
		t.Errorf("bad test manager user count, wanted :%d,got :%d",0,len(testManager.users))
	}
}
func TestAddUserDuplicateName(t *testing.T){
	testManager := NewManager()

	testFirstName := "test"
	testLastName := "Userman"
	testEmail,err := mail.ParseAddress("VivekChaurasia01@gmail.com")

	if err != nil{
		t.Fatalf("error parsing test email address : %v",err)
	}

	// Lets add user with this bad Email...
	err = testManager.AddUser(testFirstName,testLastName,testEmail.String())
	if err != nil{
		t.Fatalf("error creating user : %v",err)
	}

	err = testManager.AddUser(testFirstName,testLastName,testEmail.String())
	if err == nil{
		t.Error("no error returned for the duplicate user")
	}else{
		expectedErr := "User with this name already exist:"
		if err.Error() != expectedErr {
			t.Errorf("bad error text,wanted :%s,got :%s",expectedErr,err)
		}
	}

	if len(testManager.users) != 1 {
		t.Errorf("bad test manager user count, wanted :%d,got :%d",1,len(testManager.users))
	}
}
func TestGetUserByName(t *testing.T){
	testManager := NewManager()
	
	err := testManager.AddUser("vivek","chaurasia","vivek@gmial.com")
	if err != nil{
		t.Fatalf("error while adding test user : %v",err)
	}

	err = testManager.AddUser("Aman","Patel","vivek@gmial.com")
	if err != nil{
		t.Fatalf("error while adding test user : %v",err)
	}

	err = testManager.AddUser("Bablu","chaurasia","vivek@gmial.com")
	if err != nil{
		t.Fatalf("error while adding test user : %v",err)
	}

	err = testManager.AddUser("Luther","Boggarapu","vivek@gmial.com")
	if err != nil{
		t.Fatalf("error while adding test user : %v",err)
	}

	err = testManager.AddUser("Shivam","Yadav","vivek@gmial.com")
	if err != nil{
		t.Fatalf("error while adding test user : %v",err)
	}

	tests := map[string]struct{
		first string
		last string
		expected *User
		expectedErr error
	}{
		"Simple LookUp":{
			first : "vivek",
			last : "chaurasia",
			expected: &testManager.users[0],
			expectedErr: nil,
		},
		"Last Element Lookup":{
			first : "Bablu",
			last : "chaurasia",
			expected: &testManager.users[2],
			expectedErr: nil,
		},
		"No Match LookUp":{
			first : "Zoro",
			last : "Luffy",
			expected: nil,
			expectedErr: ErrNoResultFound,
		},
		"Partial Match LookUp":{
			first : "vivek",
			last : "vivek",
			expected: nil,
			expectedErr: ErrNoResultFound,
		},
		"Empty First Name":{
			first : "",
			last : "chaurasia",
			expected: nil,
			expectedErr: ErrNoResultFound,
		},
		"Empty Last Name":{
			first : "shivam",
			last : "",
			expected: nil,
			expectedErr: ErrNoResultFound,
		},
	}
	for name,test := range tests {
		result,err := testManager.GetUserByName(test.first,test.last)
		if !reflect.DeepEqual(result,test.expected){
			t.Errorf("%s :invalid result\ngot : %+v\nWanted : %+v\n",name,result,test.expected)
		}
		if !errors.Is(err,test.expectedErr){
			t.Errorf("%s:Invalid error result \ngot :%v\nWanted: %v",name,result,test.expectedErr)
		}
	}
}

