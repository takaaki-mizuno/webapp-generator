package database_schema

type Name struct {
	Original string
	Singular NameForm
	Plural   NameForm
}

type NameForm struct {
	Camel string
	Title string
	Snake string
	Kebab string
}
