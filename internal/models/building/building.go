package building

type RawBuilding struct {
	buildingURL string
}

type Building struct {
	Title         string
	BuildingType  string
	Neighbourhood string
	StreetName    string
	StreetNumber  string
	Area          string
	Price         string
}
