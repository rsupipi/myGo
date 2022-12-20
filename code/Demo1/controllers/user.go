package controllers

import (
	"net/http"
	"regexp"
)

/*		**** Define the user controller ****	*/
type userController struct {
	userIDPattern *regexp.Regexp
}

/*
 * Here it specify the type that want to bind too. It is the magic that makes a function into a method
 * uc - works as the "this" key word here
 */

/*		**** Hander interface ****
since the implementation is same, it automatically implements this handler interface. We don't need to specify that we
are explicitly implementing that interface. I can just implement the method.
*/

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from pipi"))
}

/*		**** define a constructor ****
 * for define a constructor it uses the "new" word as "newUserController"
 * Then it returns the pointer of the controller object
 */
func newUserController() *userController {
	/*	**** &userController ****
		* it hasn't a named variable here, just constructing the object and using the address of operator
		* that is permissible when dealing with struct data types, which means, we can immediately take the address of it.

		If you are dealing with a numerical literal( like 42), you can't take the address of 42, because there is no memory
		location that specially holding that. you have to create the variable first.
		But with "Structs", we are allowed to do this.

		This is a local variable, we are constructing "userController" in this scope of the function and returning the
		address of it. Go will recognize if we are returning the address of a local variable, and it'll automatically
		promote that variable up to the level that it needs to be. So we are not loosing this "userController" variable
		that we are creating. We can return the address of that local variable. We are not going to free up that memory
		and override it.
	*/
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
