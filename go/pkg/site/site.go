package site

type ID string

type Meta struct {
	Title string
}

type Site struct {
	ID   ID
	Meta Meta
}
