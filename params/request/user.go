package request

/*
	{
	  "user":{
	    "email": "jake@jake.jake",
	    "password": "jakejake"
	  }
	}
*/
type UserAuthenticationRequest struct {
	User UserAuthenticationBody `json:"user"`
}

type UserAuthenticationBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
{
  "user":{
    "username": "Jacob",
    "email": "jake@jake.jake",
    "password": "jakejake"
  }
}
*/

type UserRegistrationRequest struct {
	User UserRegistrationBody `json:"user"`
}

type UserRegistrationBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
