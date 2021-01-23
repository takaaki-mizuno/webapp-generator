package open_api_spec

type Name struct {
	Original string
	Default  NameForm
	Singular NameForm
	Plural   NameForm
}

type NameForm struct {
	Camel string
	Title string
	Snake string
	Kebab string
}
