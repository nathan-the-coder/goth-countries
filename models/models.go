package models

type Country struct {
	Name					string
	Capital 			string
	Population		int
	Continent			string
	Description		string
}

var Countries = []Country{
	{
		Name:            "Japan",
		Capital:         "Tokyo",
		Population:      125700000,
		Continent:       "Asia",
		Description:     "An island country known for its unique culture and technology.",
	},
	{
		Name:            "Brazil",
		Capital:         "Brasilia",
		Population:      214300000,
		Continent:       "South America",
		Description:     "The largest country in South America, famous for its biodiversity.",
	},
	{
		Name:            "France",
		Capital:         "Paris",
		Population:      67390000,
		Continent:       "Europe",
		Description:     "Known for its art, culture and cuisine.",
	},
}


