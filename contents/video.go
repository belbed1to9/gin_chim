package contents

type Person struct {
	FirstName string `json:"firstname" biding:"required"`
	LastName  string `json:"lastname" biding:"required"`
	Age       int8   `json:"age" biding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

//  xml:"title" form:"title "validate":"email" binding:"required"
type Video struct {
	Title       string `json:"title" binding:"min=2,max=100" validate:"is-cool"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" biding:"required"`
}
