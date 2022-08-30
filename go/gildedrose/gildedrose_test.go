package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Dexterity_Vest(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"+5 Dexterity Vest", 9, 19},
		{"+5 Dexterity Vest", 8, 18},
		{"+5 Dexterity Vest", 7, 17},
		{"+5 Dexterity Vest", 6, 16},
		{"+5 Dexterity Vest", 5, 15},
		{"+5 Dexterity Vest", 4, 14},
		{"+5 Dexterity Vest", 3, 13},
		{"+5 Dexterity Vest", 2, 12},
		{"+5 Dexterity Vest", 1, 11},
		{"+5 Dexterity Vest", 0, 10},
		{"+5 Dexterity Vest", -1, 8},
		{"+5 Dexterity Vest", -2, 6},
		{"+5 Dexterity Vest", -3, 4},
		{"+5 Dexterity Vest", -4, 2},
		{"+5 Dexterity Vest", -5, 0},
		{"+5 Dexterity Vest", -6, 0},
		{"+5 Dexterity Vest", -7, 0},
		{"+5 Dexterity Vest", -8, 0},
		{"+5 Dexterity Vest", -9, 0},
	}

	var expected = []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"+5 Dexterity Vest", 8, 18},
		{"+5 Dexterity Vest", 7, 17},
		{"+5 Dexterity Vest", 6, 16},
		{"+5 Dexterity Vest", 5, 15},
		{"+5 Dexterity Vest", 4, 14},
		{"+5 Dexterity Vest", 3, 13},
		{"+5 Dexterity Vest", 2, 12},
		{"+5 Dexterity Vest", 1, 11},
		{"+5 Dexterity Vest", 0, 10},
		{"+5 Dexterity Vest", -1, 8},
		{"+5 Dexterity Vest", -2, 6},
		{"+5 Dexterity Vest", -3, 4},
		{"+5 Dexterity Vest", -4, 2},
		{"+5 Dexterity Vest", -5, 0},
		{"+5 Dexterity Vest", -6, 0},
		{"+5 Dexterity Vest", -7, 0},
		{"+5 Dexterity Vest", -8, 0},
		{"+5 Dexterity Vest", -9, 0},
		{"+5 Dexterity Vest", -10, 0},
	}

	gildedrose.UpdateQuality(items)
	var total = len(items)

	for i := 0; i < total; i++ {
		if items[i].Name != expected[i].Name {
			t.Errorf("Day %d Name: Expected %s but got %s ", i+1, expected[i].Name, items[i].Name)
		}

		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("Day %d SellIn: Expected %d but got %d ", i+1, expected[i].SellIn, items[i].SellIn)
		}

		if items[i].Quality != expected[i].Quality {
			t.Errorf("Day %d Quality: Expected %d but got %d ", i+1, expected[i].Quality, items[i].Quality)
		}
	}
}

func Test_Aged_Brie(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 2, 0},
		{"Aged Brie", 1, 1},
		{"Aged Brie", 0, 2},
		{"Aged Brie", -1, 4},
		{"Aged Brie", -2, 6},
		{"Aged Brie", -3, 8},
		{"Aged Brie", -4, 10},
		{"Aged Brie", -5, 12},
		{"Aged Brie", -6, 14},
		{"Aged Brie", -7, 16},
		{"Aged Brie", -8, 18},
		{"Aged Brie", -9, 20},
		{"Aged Brie", -10, 22},
		{"Aged Brie", -11, 24},
		{"Aged Brie", -12, 26},
		{"Aged Brie", -13, 28},
		{"Aged Brie", -14, 30},
		{"Aged Brie", -15, 32},
		{"Aged Brie", -16, 34},
		{"Aged Brie", -17, 36},
		{"Aged Brie", -18, 49},
		{"Aged Brie", -19, 50},
	}

	var expected = []*gildedrose.Item{
		{"Aged Brie", 1, 1},
		{"Aged Brie", 0, 2},
		{"Aged Brie", -1, 4},
		{"Aged Brie", -2, 6},
		{"Aged Brie", -3, 8},
		{"Aged Brie", -4, 10},
		{"Aged Brie", -5, 12},
		{"Aged Brie", -6, 14},
		{"Aged Brie", -7, 16},
		{"Aged Brie", -8, 18},
		{"Aged Brie", -9, 20},
		{"Aged Brie", -10, 22},
		{"Aged Brie", -11, 24},
		{"Aged Brie", -12, 26},
		{"Aged Brie", -13, 28},
		{"Aged Brie", -14, 30},
		{"Aged Brie", -15, 32},
		{"Aged Brie", -16, 34},
		{"Aged Brie", -17, 36},
		{"Aged Brie", -18, 38},
		{"Aged Brie", -19, 50},
		{"Aged Brie", -20, 50},
	}

	gildedrose.UpdateQuality(items)
	var total = len(items)

	for i := 0; i < total; i++ {
		if items[i].Name != expected[i].Name {
			t.Errorf("Day %d Name: Expected %s but got %s ", i+1, expected[i].Name, items[i].Name)
		}

		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("Day %d SellIn: Expected %d but got %d ", i+1, expected[i].SellIn, items[i].SellIn)
		}

		if items[i].Quality != expected[i].Quality {
			t.Errorf("Day %d Quality: Expected %d but got %d ", i+1, expected[i].Quality, items[i].Quality)
		}
	}
}

func Test_Elixir_of_the_Mongoose(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Elixir of the Mongoose", 5, 7},
		{"Elixir of the Mongoose", 4, 6},
		{"Elixir of the Mongoose", 3, 5},
		{"Elixir of the Mongoose", 2, 4},
		{"Elixir of the Mongoose", 1, 3},
		{"Elixir of the Mongoose", 0, 2},
		{"Elixir of the Mongoose", -1, 0},
		{"Elixir of the Mongoose", -2, 0},
		{"Elixir of the Mongoose", -3, 0},
		{"Elixir of the Mongoose", -4, 0},
		{"Elixir of the Mongoose", -5, 0},
		{"Elixir of the Mongoose", -6, 0},
		{"Elixir of the Mongoose", -7, 0},
		{"Elixir of the Mongoose", -8, 0},
		{"Elixir of the Mongoose", -9, 0},
		{"Elixir of the Mongoose", -10, 0},
		{"Elixir of the Mongoose", -11, 0},
		{"Elixir of the Mongoose", -12, 0},
		{"Elixir of the Mongoose", -13, 0},
		{"Elixir of the Mongoose", -14, 0},
	}

	var expected = []*gildedrose.Item{

		{"Elixir of the Mongoose", 4, 6},
		{"Elixir of the Mongoose", 3, 5},
		{"Elixir of the Mongoose", 2, 4},
		{"Elixir of the Mongoose", 1, 3},
		{"Elixir of the Mongoose", 0, 2},
		{"Elixir of the Mongoose", -1, 0},
		{"Elixir of the Mongoose", -2, 0},
		{"Elixir of the Mongoose", -3, 0},
		{"Elixir of the Mongoose", -4, 0},
		{"Elixir of the Mongoose", -5, 0},
		{"Elixir of the Mongoose", -6, 0},
		{"Elixir of the Mongoose", -7, 0},
		{"Elixir of the Mongoose", -8, 0},
		{"Elixir of the Mongoose", -9, 0},
		{"Elixir of the Mongoose", -10, 0},
		{"Elixir of the Mongoose", -11, 0},
		{"Elixir of the Mongoose", -12, 0},
		{"Elixir of the Mongoose", -13, 0},
		{"Elixir of the Mongoose", -14, 0},
		{"Elixir of the Mongoose", -15, 0},
	}

	gildedrose.UpdateQuality(items)
	var total = len(items)

	for i := 0; i < total; i++ {
		if items[i].Name != expected[i].Name {
			t.Errorf("Day %d Name: Expected %s but got %s ", i+1, expected[i].Name, items[i].Name)
		}

		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("Day %d SellIn: Expected %d but got %d ", i+1, expected[i].SellIn, items[i].SellIn)
		}

		if items[i].Quality != expected[i].Quality {
			t.Errorf("Day %d Quality: Expected %d but got %d ", i+1, expected[i].Quality, items[i].Quality)
		}
	}
}

func Test_Sulfuras_Hand_of_Ragnaros(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
	}

	var expected = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
	}

	gildedrose.UpdateQuality(items)
	var total = len(items)

	for i := 0; i < total; i++ {
		if items[i].Name != expected[i].Name {
			t.Errorf("Day %d Name: Expected %s but got %s ", i+1, expected[i].Name, items[i].Name)
		}

		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("Day %d SellIn: Expected %d but got %d ", i+1, expected[i].SellIn, items[i].SellIn)
		}

		if items[i].Quality != expected[i].Quality {
			t.Errorf("Day %d Quality: Expected %d but got %d ", i+1, expected[i].Quality, items[i].Quality)
		}
	}
}

func Test_Backstage_passes_to_a_TAFKAL80ETC_concert(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 24},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 27},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 29},
		{"Backstage passes to a TAFKAL80ETC concert", 7, 31},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 33},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 38},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 41},
		{"Backstage passes to a TAFKAL80ETC concert", 2, 44},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 47},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -3, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -4, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
	}

	var expected = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 24},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 27},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 29},
		{"Backstage passes to a TAFKAL80ETC concert", 7, 31},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 33},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 38},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 41},
		{"Backstage passes to a TAFKAL80ETC concert", 2, 44},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 47},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -3, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -4, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
	}

	gildedrose.UpdateQuality(items)
	var total = len(items)

	for i := 0; i < total; i++ {
		if items[i].Name != expected[i].Name {
			t.Errorf("Day %d Name: Expected %s but got %s ", i+1, expected[i].Name, items[i].Name)
		}

		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("Day %d SellIn: Expected %d but got %d ", i+1, expected[i].SellIn, items[i].SellIn)
		}

		if items[i].Quality != expected[i].Quality {
			t.Errorf("Day %d Quality: Expected %d but got %d ", i+1, expected[i].Quality, items[i].Quality)
		}
	}
}
func Test_Conjured_Mana_Cake(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 3, 6},
		{"Conjured Mana Cake", 2, 4},
		{"Conjured Mana Cake", 1, 2},
		{"Conjured Mana Cake", 0, 0},
		{"Conjured Mana Cake", -1, 0},
		{"Conjured Mana Cake", -2, 0},
	}

	var expected = []*gildedrose.Item{
		{"Conjured Mana Cake", 2, 4},
		{"Conjured Mana Cake", 1, 2},
		{"Conjured Mana Cake", 0, 0},
		{"Conjured Mana Cake", -1, 0},
		{"Conjured Mana Cake", -2, 0},
		{"Conjured Mana Cake", -3, 0},
	}

	gildedrose.UpdateQuality(items)
	var total = len(items)

	for i := 0; i < total; i++ {
		if items[i].Name != expected[i].Name {
			t.Errorf("Day %d Name: Expected %s but got %s ", i+1, expected[i].Name, items[i].Name)
		}

		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("Day %d SellIn: Expected %d but got %d ", i+1, expected[i].SellIn, items[i].SellIn)
		}

		if items[i].Quality != expected[i].Quality {
			t.Errorf("Day %d Quality: Expected %d but got %d ", i+1, expected[i].Quality, items[i].Quality)
		}
	}
}
