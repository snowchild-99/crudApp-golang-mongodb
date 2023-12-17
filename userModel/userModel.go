package usermodel

/* Always write Your model Name and its Variabl in Captial otherwise it will give
struct field id has json tag but is not exported
which will map to json key
*/

type Users struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
	Age    string `json:"age"`
	Email  string `json:"email"`
}
