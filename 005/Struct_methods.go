package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// This Function is a method for the struct above since we have added an instance
// of the stuct within the declaration of the struct this function can be used as a method

// Struct method can access the properties of the struct.
func (authI authenticationInfo) getAuthInfo() string {
	return fmt.Sprintf(
		"User Info : %s,%s",
		authI.username,
		authI.password,
	)
}

func main() {
	user := authenticationInfo{
		username: "Jhon Doe",
		password: "Hello_there",
	}

	// This concept is similar to methods in classes.

	fmt.Println(user.getAuthInfo())
}
