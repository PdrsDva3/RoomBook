package models

type TypeCreate struct {
	Type string `json:"type"`
}

type TagCreate struct {
	IDType int    `json:"id_type"`
	Tag    string `json:"tag"`
}

type TagBase struct {
	IDTag int    `json:"id_tag"`
	Tag   string `json:"tag"`
}

type TypeBase struct {
	IDType int    `json:"id_type"`
	Type   string `json:"type"`
}

type Tag struct {
	TagBase
	TypeBase
}

type TagHotel struct {
	IDTag   int `json:"id_tag"`
	IDHotel int `json:"id_hotel"`
}

type TagRoom struct {
	IDTag  int `json:"id_tag"`
	IDRoom int `json:"id_room"`
}

type TagsType struct {
	TypeBase
	Tags []TagBase `json:"tags"`
}
