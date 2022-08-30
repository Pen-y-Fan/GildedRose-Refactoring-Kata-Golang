package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {

	for _, item := range items {
		switch item.Name {

		case "Aged Brie":
			updateAgedBrie(item)

		case "Backstage passes to a TAFKAL80ETC concert":
			updateBackstagePasses(item)

		case "Sulfuras, Hand of Ragnaros":
			updateSulfurasHandOfRagnaros(item)

		case "Conjured Mana Cake":
			updateConjuredManaCake(item)

		default:
			updateNormal(item)
		}
	}
}

func updateConjuredManaCake(item *Item) {
	// "Conjured" items degrade in Quality twice as fast as normal items
	reduceQuality(item)
	reduceQuality(item)

	item.SellIn--

	if item.SellIn < 0 {
		reduceQuality(item)
		reduceQuality(item)
	}
}

func updateNormal(item *Item) {

	reduceQuality(item)

	item.SellIn--

	if item.SellIn < 0 {
		reduceQuality(item)
	}
}

func updateAgedBrie(item *Item) {
	increaseQuality(item)

	item.SellIn--

	if item.SellIn < 0 {
		increaseQuality(item)
	}
}

func updateBackstagePasses(item *Item) {
	increaseQuality(item)

	// Quality increases by 2 when there are 10 days or less
	if item.SellIn < 11 {
		increaseQuality(item)
	}
	// ... and by 3 when there are 5 days or less but
	if item.SellIn < 6 {
		increaseQuality(item)
	}

	item.SellIn--

	// Quality drops to 0 after the concert
	if item.SellIn < 0 {
		item.Quality = 0
	}
}

func updateSulfurasHandOfRagnaros(item *Item) {
	// "Sulfuras" is a legendary item and as such its Quality is 80 and it never alters.
}

func reduceQuality(item *Item) {
	// The Quality of an item is never negative
	if item.Quality > 0 {
		item.Quality--
	}
}

func increaseQuality(item *Item) {
	// The Quality of an item is never more than 50
	if item.Quality < 50 {
		item.Quality++
	}
}
