package models

type Author struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Bio string `json:"bio"`
}
