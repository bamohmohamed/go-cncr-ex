package model

type Manga struct {
	Id            int
	Title, Author string
}

//func (Manga) NewManga(Id int, Title, Author string) Manga {
//	return Manga{Id: Id, Title: Title, Author: Author}
//}

var Mangas = [10]Manga{
	{
		Id:     1,
		Title:  "HxH",
		Author: "Yoshihiro Togashi",
	},
	{
		Id:     2,
		Title:  "One Piece",
		Author: "Eiichiro Oda",
	},
	{
		Id:     3,
		Title:  "Dragon Ball",
		Author: "Akira Toriyama",
	},
	{
		Id:     4,
		Title:  "Naruto",
		Author: "Naruto",
	},
	{
		Id:     5,
		Title:  "Bleach",
		Author: "Tite Kubo",
	},
	{
		Id:     6,
		Title:  "Slam Dunk",
		Author: "Takehiko Inoue",
	},
	{
		Id:     7,
		Title:  "Fairy Tail",
		Author: "Hiro Mashima",
	},
	{
		Id:     8,
		Title:  "City Hunter",
		Author: "Tsukasa Hojo",
	},
	{
		Id:     9,
		Title:  "Nana",
		Author: "Ai Yazawa",
	},
}
