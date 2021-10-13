package databaseschema

// Name ...
type Name struct {
	Original string
	Singular NameForm
	Plural   NameForm
}

// NameForm ...
type NameForm struct {
	Camel string
	Title string
	Snake string
	Kebab string
}
