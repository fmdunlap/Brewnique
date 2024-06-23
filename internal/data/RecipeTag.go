package data

type RecipeTag string

const (
	// Beer
	RecipeTagAle        RecipeTag = "ale"
	RecipeTagLager      RecipeTag = "lager"
	RecipeTagHybrid     RecipeTag = "hybrid"
	RecipeTagIPA        RecipeTag = "ipa"
	RecipeTagPaleAle    RecipeTag = "pale_ale"
	RecipeTagStout      RecipeTag = "stout"
	RecipeTagPorter     RecipeTag = "porter"
	RecipeTagWheatBeer  RecipeTag = "wheat_beer"
	RecipeTagBelgian    RecipeTag = "belgian"
	RecipeTagPilsner    RecipeTag = "pilsner"
	RecipeTagBock       RecipeTag = "bock"
	RecipeTagMarzen     RecipeTag = "marzen"
	RecipeTagDunkel     RecipeTag = "dunkel"
	RecipeTagSour       RecipeTag = "sour"
	RecipeTagSaison     RecipeTag = "saison"
	RecipeTagBarleywine RecipeTag = "barleywine"
	RecipeTagAmberAle   RecipeTag = "amber_ale"
	RecipeTagBrownAle   RecipeTag = "brown_ale"
	RecipeTagKolsch     RecipeTag = "kolsch"
	RecipeTagHefeweizen RecipeTag = "hefeweizen"
	RecipeTagWitbier    RecipeTag = "witbier"
	RecipeTagDubbel     RecipeTag = "dubbel"
	RecipeTagTripel     RecipeTag = "tripel"
	RecipeTagQuad       RecipeTag = "quad"

	// Mead
	RecipeTagTraditionalMead RecipeTag = "traditional_mead"
	RecipeTagMelomel         RecipeTag = "melomel"
	RecipeTagCyser           RecipeTag = "cyser"
	RecipeTagPyment          RecipeTag = "pyment"
	RecipeTagMetheglin       RecipeTag = "metheglin"
	RecipeTagBraggot         RecipeTag = "braggot"
	RecipeTagAcerglyn        RecipeTag = "acerglyn"
	RecipeTagBochet          RecipeTag = "bochet"
	RecipeTagHydromel        RecipeTag = "hydromel"
	RecipeTagSweetMead       RecipeTag = "sweet_mead"
	RecipeTagDryMead         RecipeTag = "dry_mead"
	RecipeTagSessionMead     RecipeTag = "session_mead"

	// Wine
	RecipeTagRedWine       RecipeTag = "red_wine"
	RecipeTagWhiteWine     RecipeTag = "white_wine"
	RecipeTagRose          RecipeTag = "rose"
	RecipeTagFruitWine     RecipeTag = "fruit_wine"
	RecipeTagDessertWine   RecipeTag = "dessert_wine"
	RecipeTagSparklingWine RecipeTag = "sparkling_wine"

	// Cider
	RecipeTagHardCider   RecipeTag = "hard_cider"
	RecipeTagPerry       RecipeTag = "perry"
	RecipeTagDryCider    RecipeTag = "dry_cider"
	RecipeTagSweetCider  RecipeTag = "sweet_cider"
	RecipeTagSpicedCider RecipeTag = "spiced_cider"

	// Spirits
	RecipeTagWhiskey   RecipeTag = "whiskey"
	RecipeTagVodka     RecipeTag = "vodka"
	RecipeTagGin       RecipeTag = "gin"
	RecipeTagRum       RecipeTag = "rum"
	RecipeTagBrandy    RecipeTag = "brandy"
	RecipeTagMoonshine RecipeTag = "moonshine"

	// Other Fermented
	RecipeTagKombucha     RecipeTag = "kombucha"
	RecipeTagHardKombucha RecipeTag = "hard_kombucha"
	RecipeTagKvass        RecipeTag = "kvass"
	RecipeTagTepache      RecipeTag = "tepache"
	RecipeTagJun          RecipeTag = "jun"

	// Country Wines
	RecipeTagDandelionWine   RecipeTag = "dandelion_wine"
	RecipeTagElderflowerWine RecipeTag = "elderflower_wine"
	RecipeTagRhubarbWine     RecipeTag = "rhubarb_wine"

	// Ingredients
	RecipeTagHops   RecipeTag = "hops"
	RecipeTagMalts  RecipeTag = "malts"
	RecipeTagYeast  RecipeTag = "yeast"
	RecipeTagFruit  RecipeTag = "fruit"
	RecipeTagSpices RecipeTag = "spices"
	RecipeTagHoney  RecipeTag = "honey"
	RecipeTagGrains RecipeTag = "grains"
	RecipeTagHerbs  RecipeTag = "herbs"

	// Brewing Methods
	RecipeTagAllGrain         RecipeTag = "all_grain"
	RecipeTagExtract          RecipeTag = "extract"
	RecipeTagPartialMash      RecipeTag = "partial_mash"
	RecipeTagBIAB             RecipeTag = "biab"
	RecipeTagWildFermentation RecipeTag = "wild_fermentation"

	// Characteristics
	RecipeTagHighABV    RecipeTag = "high_abv"
	RecipeTagSession    RecipeTag = "session"
	RecipeTagHoppy      RecipeTag = "hoppy"
	RecipeTagMalty      RecipeTag = "malty"
	RecipeTagSweet      RecipeTag = "sweet"
	RecipeTagDry        RecipeTag = "dry"
	RecipeTagBitter     RecipeTag = "bitter"
	RecipeTagSmoky      RecipeTag = "smoky"
	RecipeTagBarrelAged RecipeTag = "barrel_aged"

	// Seasonality
	RecipeTagSummer  RecipeTag = "summer"
	RecipeTagWinter  RecipeTag = "winter"
	RecipeTagSpring  RecipeTag = "spring"
	RecipeTagFall    RecipeTag = "fall"
	RecipeTagHoliday RecipeTag = "holiday"

	// Difficulty
	RecipeTagBeginner     RecipeTag = "beginner"
	RecipeTagIntermediate RecipeTag = "intermediate"
	RecipeTagAdvanced     RecipeTag = "advanced"

	// Batch Size
	RecipeTagSmallBatch RecipeTag = "small_batch"
	RecipeTagLargeBatch RecipeTag = "large_batch"

	// Equipment
	RecipeTagKeg    RecipeTag = "keg"
	RecipeTagBottle RecipeTag = "bottle"
	RecipeTagCarboy RecipeTag = "carboy"
	RecipeTagBucket RecipeTag = "bucket"

	// Region
	RecipeTagAmerican      RecipeTag = "american"
	RecipeTagGerman        RecipeTag = "german"
	RecipeTagBelgianRegion RecipeTag = "belgian_region"
	RecipeTagBritish       RecipeTag = "british"
	RecipeTagIrish         RecipeTag = "irish"
	RecipeTagCzech         RecipeTag = "czech"
	RecipeTagFrench        RecipeTag = "french"
	RecipeTagAsian         RecipeTag = "asian"
)

type TagCategory string

const (
	// TagCategory constants
	TagCategoryBeer            TagCategory = "Beer"
	TagCategoryMead            TagCategory = "Mead"
	TagCategoryWine            TagCategory = "Wine"
	TagCategoryCider           TagCategory = "Cider"
	TagCategorySpirits         TagCategory = "Spirits"
	TagCategoryOtherFermented  TagCategory = "Other Fermented"
	TagCategoryCountryWines    TagCategory = "Country Wines"
	TagCategoryIngredients     TagCategory = "Ingredients"
	TagCategoryBrewingMethods  TagCategory = "Brewing Methods"
	TagCategoryCharacteristics TagCategory = "Characteristics"
	TagCategorySeasonality     TagCategory = "Seasonality"
	TagCategoryDifficulty      TagCategory = "Difficulty"
	TagCategoryBatchSize       TagCategory = "Batch Size"
	TagCategoryEquipment       TagCategory = "Equipment"
	TagCategoryRegion          TagCategory = "Region"
)

var tagToCategory = map[RecipeTag]TagCategory{
	// Beer
	RecipeTagAle:        TagCategoryBeer,
	RecipeTagLager:      TagCategoryBeer,
	RecipeTagHybrid:     TagCategoryBeer,
	RecipeTagIPA:        TagCategoryBeer,
	RecipeTagPaleAle:    TagCategoryBeer,
	RecipeTagStout:      TagCategoryBeer,
	RecipeTagPorter:     TagCategoryBeer,
	RecipeTagWheatBeer:  TagCategoryBeer,
	RecipeTagBelgian:    TagCategoryBeer,
	RecipeTagPilsner:    TagCategoryBeer,
	RecipeTagBock:       TagCategoryBeer,
	RecipeTagMarzen:     TagCategoryBeer,
	RecipeTagDunkel:     TagCategoryBeer,
	RecipeTagSour:       TagCategoryBeer,
	RecipeTagSaison:     TagCategoryBeer,
	RecipeTagBarleywine: TagCategoryBeer,
	RecipeTagAmberAle:   TagCategoryBeer,
	RecipeTagBrownAle:   TagCategoryBeer,
	RecipeTagKolsch:     TagCategoryBeer,
	RecipeTagHefeweizen: TagCategoryBeer,
	RecipeTagWitbier:    TagCategoryBeer,
	RecipeTagDubbel:     TagCategoryBeer,
	RecipeTagTripel:     TagCategoryBeer,
	RecipeTagQuad:       TagCategoryBeer,

	// Mead
	RecipeTagTraditionalMead: TagCategoryMead,
	RecipeTagMelomel:         TagCategoryMead,
	RecipeTagCyser:           TagCategoryMead,
	RecipeTagPyment:          TagCategoryMead,
	RecipeTagMetheglin:       TagCategoryMead,
	RecipeTagBraggot:         TagCategoryMead,
	RecipeTagAcerglyn:        TagCategoryMead,
	RecipeTagBochet:          TagCategoryMead,
	RecipeTagHydromel:        TagCategoryMead,
	RecipeTagSweetMead:       TagCategoryMead,
	RecipeTagDryMead:         TagCategoryMead,
	RecipeTagSessionMead:     TagCategoryMead,

	// Wine
	RecipeTagRedWine:       TagCategoryWine,
	RecipeTagWhiteWine:     TagCategoryWine,
	RecipeTagRose:          TagCategoryWine,
	RecipeTagFruitWine:     TagCategoryWine,
	RecipeTagDessertWine:   TagCategoryWine,
	RecipeTagSparklingWine: TagCategoryWine,

	// Cider
	RecipeTagHardCider:   TagCategoryCider,
	RecipeTagPerry:       TagCategoryCider,
	RecipeTagDryCider:    TagCategoryCider,
	RecipeTagSweetCider:  TagCategoryCider,
	RecipeTagSpicedCider: TagCategoryCider,

	// Spirits
	RecipeTagWhiskey:   TagCategorySpirits,
	RecipeTagVodka:     TagCategorySpirits,
	RecipeTagGin:       TagCategorySpirits,
	RecipeTagRum:       TagCategorySpirits,
	RecipeTagBrandy:    TagCategorySpirits,
	RecipeTagMoonshine: TagCategorySpirits,

	// Other Fermented
	RecipeTagKombucha:     TagCategoryOtherFermented,
	RecipeTagHardKombucha: TagCategoryOtherFermented,
	RecipeTagKvass:        TagCategoryOtherFermented,
	RecipeTagTepache:      TagCategoryOtherFermented,
	RecipeTagJun:          TagCategoryOtherFermented,

	// Country Wines
	RecipeTagDandelionWine:   TagCategoryCountryWines,
	RecipeTagElderflowerWine: TagCategoryCountryWines,
	RecipeTagRhubarbWine:     TagCategoryCountryWines,

	// Ingredients
	RecipeTagHops:   TagCategoryIngredients,
	RecipeTagMalts:  TagCategoryIngredients,
	RecipeTagYeast:  TagCategoryIngredients,
	RecipeTagFruit:  TagCategoryIngredients,
	RecipeTagSpices: TagCategoryIngredients,
	RecipeTagHoney:  TagCategoryIngredients,
	RecipeTagGrains: TagCategoryIngredients,
	RecipeTagHerbs:  TagCategoryIngredients,

	// Brewing Methods
	RecipeTagAllGrain:         TagCategoryBrewingMethods,
	RecipeTagExtract:          TagCategoryBrewingMethods,
	RecipeTagPartialMash:      TagCategoryBrewingMethods,
	RecipeTagBIAB:             TagCategoryBrewingMethods,
	RecipeTagWildFermentation: TagCategoryBrewingMethods,

	// Characteristics
	RecipeTagHighABV:    TagCategoryCharacteristics,
	RecipeTagSession:    TagCategoryCharacteristics,
	RecipeTagHoppy:      TagCategoryCharacteristics,
	RecipeTagMalty:      TagCategoryCharacteristics,
	RecipeTagSweet:      TagCategoryCharacteristics,
	RecipeTagDry:        TagCategoryCharacteristics,
	RecipeTagBitter:     TagCategoryCharacteristics,
	RecipeTagSmoky:      TagCategoryCharacteristics,
	RecipeTagBarrelAged: TagCategoryCharacteristics,

	// Seasonality
	RecipeTagSummer:  TagCategorySeasonality,
	RecipeTagWinter:  TagCategorySeasonality,
	RecipeTagSpring:  TagCategorySeasonality,
	RecipeTagFall:    TagCategorySeasonality,
	RecipeTagHoliday: TagCategorySeasonality,

	// Difficulty
	RecipeTagBeginner:     TagCategoryDifficulty,
	RecipeTagIntermediate: TagCategoryDifficulty,
	RecipeTagAdvanced:     TagCategoryDifficulty,

	// Batch Size
	RecipeTagSmallBatch: TagCategoryBatchSize,
	RecipeTagLargeBatch: TagCategoryBatchSize,

	// Equipment
	RecipeTagKeg:    TagCategoryEquipment,
	RecipeTagBottle: TagCategoryEquipment,
	RecipeTagCarboy: TagCategoryEquipment,
	RecipeTagBucket: TagCategoryEquipment,

	// Region
	RecipeTagAmerican:      TagCategoryRegion,
	RecipeTagGerman:        TagCategoryRegion,
	RecipeTagBelgianRegion: TagCategoryRegion,
	RecipeTagBritish:       TagCategoryRegion,
	RecipeTagIrish:         TagCategoryRegion,
	RecipeTagCzech:         TagCategoryRegion,
	RecipeTagFrench:        TagCategoryRegion,
	RecipeTagAsian:         TagCategoryRegion,
}

func (s RecipeTag) AsCategory() TagCategory {
	switch s {
	case RecipeTagAle,
		RecipeTagLager,
		RecipeTagHybrid,
		RecipeTagIPA,
		RecipeTagPaleAle,
		RecipeTagStout,
		RecipeTagPorter,
		RecipeTagWheatBeer,
		RecipeTagBelgian,
		RecipeTagPilsner,
		RecipeTagBock,
		RecipeTagMarzen,
		RecipeTagDunkel,
		RecipeTagSour,
		RecipeTagSaison,
		RecipeTagBarleywine,
		RecipeTagAmberAle,
		RecipeTagBrownAle,
		RecipeTagKolsch,
		RecipeTagHefeweizen,
		RecipeTagWitbier,
		RecipeTagDubbel,
		RecipeTagTripel,
		RecipeTagQuad:
		return TagCategoryBeer

	case RecipeTagTraditionalMead,
		RecipeTagMelomel,
		RecipeTagCyser,
		RecipeTagPyment,
		RecipeTagMetheglin,
		RecipeTagBraggot,
		RecipeTagAcerglyn,
		RecipeTagBochet,
		RecipeTagHydromel,
		RecipeTagSweetMead,
		RecipeTagDryMead,
		RecipeTagSessionMead:
		return TagCategoryMead

	case RecipeTagRedWine,
		RecipeTagWhiteWine,
		RecipeTagRose,
		RecipeTagFruitWine,
		RecipeTagDessertWine,
		RecipeTagSparklingWine:
		return TagCategoryWine

	case RecipeTagHardCider,
		RecipeTagPerry,
		RecipeTagDryCider,
		RecipeTagSweetCider,
		RecipeTagSpicedCider:
		return TagCategoryCider

	case RecipeTagWhiskey,
		RecipeTagVodka,
		RecipeTagGin,
		RecipeTagRum,
		RecipeTagBrandy,
		RecipeTagMoonshine:
		return TagCategorySpirits

	case RecipeTagKombucha,
		RecipeTagHardKombucha,
		RecipeTagKvass,
		RecipeTagTepache,
		RecipeTagJun:
		return TagCategoryOtherFermented

	case RecipeTagDandelionWine,
		RecipeTagElderflowerWine,
		RecipeTagRhubarbWine:
		return TagCategoryCountryWines

	case RecipeTagHops,
		RecipeTagMalts,
		RecipeTagYeast,
		RecipeTagFruit,
		RecipeTagSpices,
		RecipeTagHoney,
		RecipeTagGrains,
		RecipeTagHerbs:
		return TagCategoryIngredients

	case RecipeTagAllGrain,
		RecipeTagExtract,
		RecipeTagPartialMash,
		RecipeTagBIAB,
		RecipeTagWildFermentation:
		return TagCategoryBrewingMethods

	case RecipeTagHighABV,
		RecipeTagSession,
		RecipeTagHoppy,
		RecipeTagMalty,
		RecipeTagSweet,
		RecipeTagDry,
		RecipeTagBitter,
		RecipeTagSmoky,
		RecipeTagBarrelAged:
		return TagCategoryCharacteristics

	case RecipeTagSummer,
		RecipeTagWinter,
		RecipeTagSpring,
		RecipeTagFall,
		RecipeTagHoliday:
		return TagCategorySeasonality

	case RecipeTagBeginner,
		RecipeTagIntermediate,
		RecipeTagAdvanced:
		return TagCategoryDifficulty

	case RecipeTagSmallBatch,
		RecipeTagLargeBatch:
		return TagCategoryBatchSize

	case RecipeTagKeg,
		RecipeTagBottle,
		RecipeTagCarboy,
		RecipeTagBucket:
		return TagCategoryEquipment

	case RecipeTagAmerican,
		RecipeTagGerman,
		RecipeTagBelgianRegion,
		RecipeTagBritish,
		RecipeTagIrish,
		RecipeTagCzech,
		RecipeTagFrench,
		RecipeTagAsian:
		return TagCategoryRegion

	default:
		return ""
	}
}
