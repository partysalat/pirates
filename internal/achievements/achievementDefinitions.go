package achievements

import (
	"justkile/magic-kingdom/internal/common"
	"strings"
	"time"
)

//func (broadcaster *Broadcaster) Add(c *websocket.Conn) {
//	broadcaster.lock.Lock()
//	broadcaster.connections = append(broadcaster.connections, c)
//	broadcaster.lock.Unlock()
//}

func exists(newsList []*common.News, predicate func(news2 *common.News) bool) bool {
	for _, newsItem := range newsList {
		if predicate(newsItem) {
			return true
		}
	}
	return false
}
func last(newsList []*common.News) *common.News {
	if len(newsList) == 0 {
		return nil
	}
	return newsList[len(newsList)-1]
}
func containsDrink(newsList []*common.News, drinkName string) bool {
	for _, newsItem := range newsList {
		if strings.Contains(getDrinkName(newsItem), drinkName) {
			return true
		}
	}
	return false
}
func getAtReverse(newsList []*common.News, index int) *common.News {
	if len(newsList) <= index {
		return nil
	}
	return newsList[len(newsList)-1-index]
}
func getDrinkName(newsItem *common.News) string {
	return newsItem.Payload["drink"].(map[string]interface{})["name"].(string)
}
func getDrinkType(newsItem *common.News) string {
	return newsItem.Payload["drink"].(map[string]interface{})["drinkType"].(string)
}
func getAmount(newsItem *common.News) int {
	return int(newsItem.Payload["amount"].(float64))
}
func countTypes(newsList []*common.News, drinkType string) int {
	counter := 0
	for _, newsItem := range newsList {
		if getDrinkType(newsItem) == drinkType {
			counter = counter + getAmount(newsItem)
		}
	}
	return counter
}

var AchievementDefinitions = []*common.AchievementDefinition{
	{
		Achievement: common.Achievement{Name: "Matrosenschluck", Id: 1, Description: "Ein Bier bestellt", Image: "/images/achievements/beer1.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "BEER") >= 1
		},
	}, {
		Achievement: common.Achievement{Name: "Piratenpinte", Id: 2, Description: "5 Bier bestellt", Image: "/images/achievements/beer5.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "BEER") >= 5
		},
	}, {
		Achievement: common.Achievement{Name: "Schurkenstärkung", Id: 3, Description: "10 Bier bestellt", Image: "/images/achievements/beer10.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "BEER") >= 10
		},
	}, {
		Achievement: common.Achievement{Name: "Plünderer Pils", Id: 4, Description: "15 Bier bestellt", Image: "/images/achievements/beer15.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "BEER") >= 15
		},
	},
	{
		Achievement: common.Achievement{Name: "Legendäres Lager", Id: 5, Description: "20 Bier bestellt", Image: "/images/achievements/beer20.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "BEER") >= 20
		},
	},
	{
		Achievement: common.Achievement{Name: "Matrosenmischung", Id: 10, Description: "1 Cocktail bestellt", Image: "/images/achievements/cocktail1.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 1
		},
	},
	{
		Achievement: common.Achievement{Name: "Freibeutermix", Id: 11, Description: "5 Cocktails bestellt", Image: "/images/achievements/cocktail5.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 5
		},
	}, {
		Achievement: common.Achievement{Name: "Käptn's Cocktail", Id: 12, Description: "10 Cocktails bestellt", Image: "/images/achievements/cocktail10.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 10
		},
	},
	{
		Achievement: common.Achievement{Name: "Admiral's Bester", Id: 13, Description: "15 Cocktails bestellt", Image: "/images/achievements/cocktail15.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 15
		},
	},
	{
		Achievement: common.Achievement{Name: "Goldener Grogtail", Id: 14, Description: "20 Cocktails bestellt", Image: "/images/achievements/cocktail20.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 20
		},
	},
	{
		Achievement: common.Achievement{Name: "Matrosen-Limo", Id: 20, Description: "1 Softdrink bestellt", Image: "/images/achievements/softdrink1.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "SOFTDRINK") >= 1
		},
	},
	{
		Achievement: common.Achievement{Name: "Piratenplörre", Id: 21, Description: "5 Softdrinks bestellt", Image: "/images/achievements/softdrink5.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "SOFTDRINK") >= 5
		},
	},
	{
		Achievement: common.Achievement{Name: "Schurkensprudel", Id: 22, Description: "10 Softdrinks bestellt", Image: "/images/achievements/softdrink10.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "SOFTDRINK") >= 10
		},
	},
	{
		Achievement: common.Achievement{Name: "Wrackwasser", Id: 23, Description: "15 Softdrinks bestellt", Image: "/images/achievements/softdrink15.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "SOFTDRINK") >= 15
		},
	},
	{
		Achievement: common.Achievement{Name: "Sprudel Sultan", Id: 24, Description: "20 Softdrink bestellt", Image: "/images/achievements/softdrink20.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "SOFTDRINK") >= 20
		},
	},
	{
		Achievement: common.Achievement{Name: "Doppeldecker Pirat", Id: 25, Description: "1 Bier + 1 Cocktail bestellt", Image: "/images/achievements/double1.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 1 && countTypes(newsList, "BEER") >= 1
		},
	},
	{
		Achievement: common.Achievement{Name: "Doppelter Doppeldecker Pirat", Id: 26, Description: "5 Bier + 5 Cocktails bestellt", Image: "/images/achievements/double5.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 5 && countTypes(newsList, "BEER") >= 5
		},
	},
	{
		Achievement: common.Achievement{Name: "Doppelter Doppel-Doppeldecker Pirat", Id: 27, Description: "10 Bier + 10 Cocktails bestellt", Image: "/images/achievements/double10.webp"},
		Predicate: func(newsList []*common.News) bool {
			return countTypes(newsList, "COCKTAIL")+countTypes(newsList, "COCKTAIL_DISENCHANTED") >= 10 && countTypes(newsList, "BEER") >= 10
		},
	},
	{
		Achievement: common.Achievement{Name: "Kanonenmeister", Id: 30, Description: "10 Shots auf einmal bestellt", Image: "/images/achievements/10shots.webp"},
		Predicate: func(newsList []*common.News) bool {
			currentNews := last(newsList)
			return getDrinkType(currentNews) == "SHOT" && getAmount(currentNews) >= 10
		},
	},
	{
		Achievement: common.Achievement{Name: "Frühschoppen Freibeuter", Id: 31, Description: "Bier zwischen 8-12 Uhr morgens", Image: "/images/achievements/earlybird.webp"},
		Predicate: func(newsList []*common.News) bool {
			currentNews := last(newsList)
			now := time.Now()
			t1 := time.Date(now.Year(), now.Month(), now.Day(), 7, 0, 0, 0, now.Location())
			t2 := time.Date(now.Year(), now.Month(), now.Day(), 11, 0, 0, 0, now.Location())
			return getDrinkType(currentNews) == "BEER" && now.After(t1) && now.Before(t2)
		},
	},
	{
		Achievement: common.Achievement{Name: "Das selbe nochmal", Id: 32, Description: "3x nacheinander denselben Drink bestellt", Image: "/images/achievements/samedrink3.webp"},
		Predicate: func(newsList []*common.News) bool {
			currentNews := last(newsList)
			currentNews1 := getAtReverse(newsList, 1)
			currentNews2 := getAtReverse(newsList, 2)
			if currentNews2 == nil || currentNews1 == nil || currentNews == nil {
				return false
			}
			return getDrinkName(currentNews) == getDrinkName(currentNews1) && getDrinkName(currentNews2) == getDrinkName(currentNews1)
		},
	},

	// Trink spezifisch
	// Berliner Luft
	// ...
	{
		Achievement: common.Achievement{Name: "Sie nannten sie MUM", Id: 100, Description: "Kapitäneuse bestellt", Image: "/images/achievements/deinemudder.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Kapitäneuse")
		},
	},
	{
		Achievement: common.Achievement{Name: "Saurer Seebär", Id: 101, Description: "Seeräubers Saurer bestellt", Image: "/images/achievements/gin-sour.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Seeräubers Saurer")
		},
	},
	{
		Achievement: common.Achievement{Name: "Käpt'n Jack", Id: 102, Description: "Fluch der Karibik bestellt", Image: "/images/achievements/long-island-iced-tea.jpeg"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Fluch der Karibik")
		},
	},
	{
		Achievement: common.Achievement{Name: "Tiki Terror", Id: 103, Description: "Schatzinsel Zauber bestellt", Image: "/images/achievements/maitai.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Schatzinsel Zauber")
		},
	},
	{
		Achievement: common.Achievement{Name: "Rum Raider", Id: 104, Description: "Captain's Punch bestellt", Image: "/images/achievements/planters-punch.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Captain's Punch")
		},
	},
	{
		Achievement: common.Achievement{Name: "Havana Held", Id: 105, Description: "Blackbeards Liebster bestellt", Image: "/images/achievements/mojito.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Blackbeards Liebster")
		},
	},
	{
		Achievement: common.Achievement{Name: "Verfluchter Freibeuter", Id: 109, Description: "Einmal Kielholen bestellt", Image: "/images/achievements/zombie.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Einmal Kielholen")
		},
	},
	{
		Achievement: common.Achievement{Name: "Kapitän Tschunksegel und der Kampf um den goldenen Grogtail", Id: 115, Description: "Kapitän Tschunksegel bestellt", Image: "/images/achievements/tschunk.webp"},
		Predicate: func(newsList []*common.News) bool {
			return strings.Contains(getDrinkName(last(newsList)), "Kapitän Tschunksegel")
		},
	},
	{
		Achievement: common.Achievement{Name: "Greatest of all Time", Id: 115, Description: "Alle Cocktails mindestens 1x bestellt", Image: "/images/achievements/goat.webp"},
		Predicate: func(newsList []*common.News) bool {
			return containsDrink(newsList, "Kapitäneuse") && containsDrink(newsList, "Seeräubers Saurer") && containsDrink(newsList, "Fluch der Karibik") && containsDrink(newsList, "Schatzinsel Zauber") && containsDrink(newsList, "Captain's Punch") && containsDrink(newsList, "Blackbeards Liebster") && containsDrink(newsList, "Einmal Kielholen") && containsDrink(newsList, "Kapitän Tschunksegel")
		},
	},

	// are set manually
	{
		Achievement: common.Achievement{Name: "Käptn Tschunksegel und der Kampf um die tote Fichte", Id: 200, Description: "Turmverteidigung (leicht) gewonnen", Image: "/images/achievements/game_easy.webp"},
		Predicate: func(newsList []*common.News) bool {
			return false
		},
	},
	{
		Achievement: common.Achievement{Name: "Piraten der Toten Fichte, vereinigt euch!", Id: 201, Description: "Turmverteidigung (normal) gewonnen", Image: "/images/achievements/game_mid.webp"},
		Predicate: func(newsList []*common.News) bool {
			return false
		},
	},
	{
		Achievement: common.Achievement{Name: "Ein neues Boot tut not", Id: 202, Description: "Turmverteidigung (schwer) gewonnen", Image: "/images/achievements/game_hard.webp"},
		Predicate: func(newsList []*common.News) bool {
			return false
		},
	},
}
