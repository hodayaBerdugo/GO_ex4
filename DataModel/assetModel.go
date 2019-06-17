package DataModel

//import _ "container/list"

type Asset struct {
	Location string
	Meters   int
	Rooms    int
	Floor    int
	Price    int
	Valid    int
	Types    []int
	User     User
}

func NewAsset(rooms int, price int, valid int) *Asset {
	return &Asset{Rooms: rooms, Price: price, Valid: valid}
}

//update asset price
func UpdatePrice(asset Asset, price int) {
	asset.Price = price
}

//set asset as valid - means that it will show up in search results
func SetAssetAsValid(asset Asset) {
	asset.Valid = 1
}

//set asset as valid - means that it won't show up in search results
func SetAssetAsNotValid(asset Asset) {
	asset.Valid = 0
}
