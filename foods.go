package cook

type FoodName string // 食材

const (
	FoodPotato FoodName = "potato"
	FoodNil    FoodName = "nil"
)

var (
	_ Food = (*APotato)(nil)
	_ Food = (*AFoodNil)(nil)
)

type APotato struct{}

func (a *APotato) Name() string {
	return string(FoodPotato)
}

type AFoodNil struct{}

func (a *AFoodNil) Name() string {
	return string(FoodNil)
}
