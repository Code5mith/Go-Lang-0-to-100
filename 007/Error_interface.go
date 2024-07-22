package main

import "fmt"


// MARK: this is the default implimentation of the error interface 
/* type error interface {
	Error() string
}
 */


 type userError struct {
	name string
 }

 func (e userError) Error() string {
	return fmt.Sprintf("%s has an problem on there account, contact customer service.", e.name)
 }

 // here look that we have made the return type of the func to an error which is an actual error type built into go 
 func login(userinfo, userError) error {
	// assume validate_user checks user creds.
	validate_status bool = validate_user(userinfo)

	if (valid_user_info){
		// we just return an instance of that struct. Go will treat is as any other error.
		
		return userError{
			name: userinfo.name;
		}
	}

 }
func main() {

	// MARK: In go errors are nullable strings and they are definded as an interface internally

	// Lets overwrite the error interface and add our own custom pieces. 


}
