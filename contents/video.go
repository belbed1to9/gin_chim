package contents

type Person struct {
	FirstName string `db:"firstname" json:"firstname" biding:"required"`
	LastName  string `db:"lastname" json:"lastname" biding:"required"`
	Age       int8   `db:"age" json:"age" biding:"gte=1,lte=130"`
	Email     string `db:"email" json:"email" binding:"required,email"`
}

//  xml:"title" form:"title "validate":"email" binding:"required"
type Video struct {
	Title       string `db:"title" json:"title" binding:"min=2,max=100" validate:"is-cool"`
	Description string `db:"description" json:"description" binding:"max=200"`
	URL         string `db:"url" json:"url" binding:"required,url"`
	Author      Person `db:"author" json:"author" biding:"required"`
}
