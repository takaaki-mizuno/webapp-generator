package openapispec

// Name ...
type Name struct {
	Original string
	Default  NameForm
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
