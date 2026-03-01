package seed

import (
	"strings"

	"github.com/dhavisiregar/masak-apa/database"
	"github.com/dhavisiregar/masak-apa/models"
)

type RecipeData struct {
	Title        string
	Instructions string
	CookingTime  int
	Difficulty   string
	Ingredients  []string
}

func SeedData() {
	var count int64
	database.DB.Model(&models.Recipe{}).Count(&count)
	if count > 0 {
		return
	}

	ingredientNames := []string{
		"telur", "mie", "bawang putih", "bawang merah", "kecap manis",
		"ayam", "nasi", "cabai", "garam", "gula",
		"minyak goreng", "tomat", "kangkung", "tahu", "tempe",
		"udang", "daging sapi", "wortel", "kentang", "buncis",
		"santan", "lengkuas", "serai", "daun salam", "kunyit",
		"jahe", "kemiri", "ketumbar", "merica", "pala",
		"daun bawang", "seledri", "bawang bombay", "saus tiram", "kecap asin",
		"tepung terigu", "mentega", "susu", "keju", "coklat",
		"pisang", "singkong", "ubi", "labu", "jagung",
	}

	ingMap := map[string]uint{}
	for _, name := range ingredientNames {
		ing := models.Ingredient{}
		database.DB.FirstOrCreate(&ing, models.Ingredient{
			Name:           name,
			NormalizedName: name,
		})
		ingMap[name] = ing.ID
	}

	recipes := []RecipeData{
		{
			Title: "Nasi Goreng Spesial",
			Instructions: "1. Panaskan minyak, tumis bawang merah dan bawang putih hingga harum. 2. Masukkan telur, orak-arik. 3. Tambahkan nasi, aduk rata. 4. Beri kecap manis, garam, merica. 5. Sajikan dengan taburan daun bawang.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"nasi", "telur", "bawang merah", "bawang putih", "kecap manis", "garam", "minyak goreng", "daun bawang"},
		},
		{
			Title: "Mie Goreng Jawa",
			Instructions: "1. Rebus mie hingga matang, tiriskan. 2. Tumis bawang putih dan bawang merah. 3. Masukkan telur orak-arik. 4. Tambahkan mie, kecap manis, saus tiram. 5. Aduk rata dan sajikan.",
			CookingTime: 15, Difficulty: "easy",
			Ingredients: []string{"mie", "telur", "bawang putih", "bawang merah", "kecap manis", "saus tiram", "minyak goreng"},
		},
		{
			Title: "Ayam Goreng Krispy",
			Instructions: "1. Marinasi ayam dengan bawang putih, kunyit, garam, ketumbar. 2. Balur dengan tepung terigu. 3. Goreng dalam minyak panas hingga keemasan dan matang.",
			CookingTime: 40, Difficulty: "medium",
			Ingredients: []string{"ayam", "bawang putih", "kunyit", "garam", "ketumbar", "tepung terigu", "minyak goreng"},
		},
		{
			Title: "Ayam Bakar Kecap",
			Instructions: "1. Lumuri ayam dengan bumbu halus. 2. Ungkep hingga setengah matang. 3. Bakar di atas bara api sambil dioles kecap manis. 4. Balik hingga matang merata.",
			CookingTime: 60, Difficulty: "medium",
			Ingredients: []string{"ayam", "kecap manis", "bawang putih", "bawang merah", "jahe", "garam", "gula"},
		},
		{
			Title: "Telur Dadar Bumbu",
			Instructions: "1. Kocok telur bersama bawang merah iris, cabai, garam. 2. Panaskan minyak dalam wajan. 3. Tuang kocokan telur, masak hingga matang dan tepi kecoklatan.",
			CookingTime: 10, Difficulty: "easy",
			Ingredients: []string{"telur", "bawang merah", "cabai", "garam", "minyak goreng"},
		},
		{
			Title: "Telur Balado",
			Instructions: "1. Goreng telur rebus hingga kekuningan. 2. Haluskan cabai, bawang merah, bawang putih, tomat. 3. Tumis bumbu halus, masukkan telur. 4. Tambahkan garam, gula, aduk rata.",
			CookingTime: 25, Difficulty: "easy",
			Ingredients: []string{"telur", "cabai", "bawang merah", "bawang putih", "tomat", "garam", "gula", "minyak goreng"},
		},
		{
			Title: "Tumis Kangkung Belacan",
			Instructions: "1. Panaskan minyak, tumis bawang putih dan cabai. 2. Masukkan kangkung, tambahkan kecap asin dan garam. 3. Aduk cepat di api besar hingga layu. Sajikan segera.",
			CookingTime: 10, Difficulty: "easy",
			Ingredients: []string{"kangkung", "bawang putih", "cabai", "kecap asin", "garam", "minyak goreng"},
		},
		{
			Title: "Tahu Goreng Crispy",
			Instructions: "1. Potong tahu, taburi garam. 2. Goreng dalam minyak panas hingga keemasan dan crispy. 3. Sajikan dengan cabai rawit atau sambal kecap.",
			CookingTime: 15, Difficulty: "easy",
			Ingredients: []string{"tahu", "garam", "minyak goreng"},
		},
		{
			Title: "Tempe Orek Manis",
			Instructions: "1. Goreng tempe hingga kering. 2. Tumis bawang putih, bawang merah, cabai. 3. Masukkan tempe, tambahkan kecap manis, gula, garam. 4. Aduk hingga bumbu meresap.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"tempe", "bawang putih", "bawang merah", "cabai", "kecap manis", "gula", "garam", "minyak goreng"},
		},
		{
			Title: "Soto Ayam",
			Instructions: "1. Rebus ayam dengan serai, daun salam, lengkuas. 2. Haluskan kunyit, kemiri, bawang putih, bawang merah. 3. Tumis bumbu halus, masukkan ke kaldu. 4. Tambahkan garam dan gula. Sajikan dengan nasi.",
			CookingTime: 60, Difficulty: "medium",
			Ingredients: []string{"ayam", "serai", "daun salam", "lengkuas", "kunyit", "kemiri", "bawang putih", "bawang merah", "garam", "gula"},
		},
		{
			Title: "Opor Ayam",
			Instructions: "1. Haluskan kemiri, ketumbar, kunyit, bawang putih, bawang merah. 2. Tumis bumbu dengan serai, daun salam, lengkuas. 3. Masukkan ayam, tambahkan santan. 4. Masak hingga ayam matang dan kuah mengental.",
			CookingTime: 60, Difficulty: "medium",
			Ingredients: []string{"ayam", "santan", "kemiri", "ketumbar", "kunyit", "bawang putih", "bawang merah", "serai", "daun salam", "lengkuas", "garam", "gula"},
		},
		{
			Title: "Rendang Daging",
			Instructions: "1. Haluskan cabai, bawang merah, bawang putih, jahe, kunyit, serai. 2. Masak daging dengan bumbu halus dan santan. 3. Masak dengan api kecil sambil diaduk hingga santan mengering dan daging berwarna coklat kehitaman.",
			CookingTime: 180, Difficulty: "hard",
			Ingredients: []string{"daging sapi", "santan", "cabai", "bawang merah", "bawang putih", "jahe", "kunyit", "serai", "lengkuas", "daun salam", "garam"},
		},
		{
			Title: "Sayur Lodeh",
			Instructions: "1. Rebus santan bersama serai, daun salam, lengkuas. 2. Masukkan sayuran: wortel, buncis, labu. 3. Tambahkan tahu tempe. 4. Bumbui dengan bawang putih, bawang merah, kemiri, garam, gula.",
			CookingTime: 30, Difficulty: "medium",
			Ingredients: []string{"santan", "wortel", "buncis", "labu", "tahu", "tempe", "bawang putih", "bawang merah", "kemiri", "serai", "daun salam", "lengkuas", "garam", "gula"},
		},
		{
			Title: "Capcay Goreng",
			Instructions: "1. Tumis bawang putih dan bawang bombay. 2. Masukkan wortel, buncis, kentang. 3. Tambahkan saus tiram, kecap asin. 4. Masak hingga sayuran matang. 5. Tambahkan daun bawang dan seledri.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"wortel", "buncis", "kentang", "bawang putih", "bawang bombay", "saus tiram", "kecap asin", "daun bawang", "seledri", "minyak goreng"},
		},
		{
			Title: "Udang Goreng Tepung",
			Instructions: "1. Bersihkan udang, sisakan ekornya. 2. Bumbui dengan bawang putih, garam, merica. 3. Balur dengan tepung terigu. 4. Goreng dalam minyak panas hingga keemasan.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"udang", "bawang putih", "garam", "merica", "tepung terigu", "minyak goreng"},
		},
		{
			Title: "Udang Saus Tiram",
			Instructions: "1. Tumis bawang putih dan bawang bombay hingga harum. 2. Masukkan udang, masak hingga berubah warna. 3. Tambahkan saus tiram, kecap manis, garam. 4. Aduk rata dan sajikan.",
			CookingTime: 15, Difficulty: "easy",
			Ingredients: []string{"udang", "bawang putih", "bawang bombay", "saus tiram", "kecap manis", "garam", "minyak goreng"},
		},
		{
			Title: "Perkedel Kentang",
			Instructions: "1. Rebus kentang hingga lunak, haluskan. 2. Campur dengan telur, bawang putih, merica, garam, seledri. 3. Bentuk bulat pipih. 4. Celup telur, goreng hingga keemasan.",
			CookingTime: 35, Difficulty: "medium",
			Ingredients: []string{"kentang", "telur", "bawang putih", "merica", "garam", "seledri", "minyak goreng"},
		},
		{
			Title: "Sup Wortel Kentang",
			Instructions: "1. Tumis bawang putih dan bawang bombay. 2. Tambahkan air, masukkan wortel dan kentang. 3. Rebus hingga lunak. 4. Bumbui dengan garam, merica, seledri.",
			CookingTime: 30, Difficulty: "easy",
			Ingredients: []string{"wortel", "kentang", "bawang putih", "bawang bombay", "garam", "merica", "seledri"},
		},
		{
			Title: "Balado Tempe",
			Instructions: "1. Goreng tempe hingga keemasan. 2. Haluskan cabai, bawang merah, bawang putih, tomat. 3. Tumis bumbu hingga matang. 4. Masukkan tempe, tambahkan garam dan gula.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"tempe", "cabai", "bawang merah", "bawang putih", "tomat", "garam", "gula", "minyak goreng"},
		},
		{
			Title: "Mie Rebus Sederhana",
			Instructions: "1. Rebus mie hingga matang. 2. Tumis bawang putih, masukkan telur orak-arik. 3. Tambahkan kecap manis, garam. 4. Masukkan mie, aduk rata. 5. Sajikan dengan taburan daun bawang.",
			CookingTime: 15, Difficulty: "easy",
			Ingredients: []string{"mie", "telur", "bawang putih", "kecap manis", "garam", "daun bawang"},
		},
		{
			Title: "Ayam Geprek",
			Instructions: "1. Goreng ayam dengan tepung berbumbu hingga crispy. 2. Haluskan cabai rawit, bawang putih, garam. 3. Geprek ayam di atas sambal. 4. Sajikan dengan nasi.",
			CookingTime: 30, Difficulty: "medium",
			Ingredients: []string{"ayam", "tepung terigu", "cabai", "bawang putih", "garam", "minyak goreng", "nasi"},
		},
		{
			Title: "Tumis Tahu Tempe",
			Instructions: "1. Goreng tahu dan tempe hingga keemasan. 2. Tumis bawang merah, bawang putih, cabai. 3. Masukkan tahu tempe, kecap manis, garam. 4. Aduk rata dan sajikan.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"tahu", "tempe", "bawang merah", "bawang putih", "cabai", "kecap manis", "garam", "minyak goreng"},
		},
		{
			Title: "Nasi Goreng Seafood",
			Instructions: "1. Tumis bawang putih dan bawang merah. 2. Masukkan udang dan orak-arik telur. 3. Tambahkan nasi, kecap manis, saus tiram. 4. Aduk rata, bumbui garam dan merica.",
			CookingTime: 25, Difficulty: "medium",
			Ingredients: []string{"nasi", "udang", "telur", "bawang putih", "bawang merah", "kecap manis", "saus tiram", "garam", "merica", "minyak goreng"},
		},
		{
			Title: "Semur Daging",
			Instructions: "1. Rebus daging hingga empuk. 2. Tumis bawang merah, bawang putih, jahe. 3. Masukkan daging, tambahkan kecap manis, pala, merica. 4. Masak hingga bumbu meresap.",
			CookingTime: 90, Difficulty: "medium",
			Ingredients: []string{"daging sapi", "bawang merah", "bawang putih", "jahe", "kecap manis", "pala", "merica", "garam", "minyak goreng"},
		},
		{
			Title: "Pisang Goreng",
			Instructions: "1. Buat adonan dari tepung terigu, gula, garam, dan air. 2. Kupas pisang, celupkan ke adonan. 3. Goreng dalam minyak panas hingga keemasan.",
			CookingTime: 15, Difficulty: "easy",
			Ingredients: []string{"pisang", "tepung terigu", "gula", "garam", "minyak goreng"},
		},
		{
			Title: "Kolak Pisang",
			Instructions: "1. Rebus santan bersama gula, garam, daun salam. 2. Masukkan pisang yang sudah dipotong. 3. Masak hingga pisang lunak dan kuah mengental.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"pisang", "santan", "gula", "garam", "daun salam"},
		},
		{
			Title: "Bubur Ayam",
			Instructions: "1. Masak nasi dengan banyak air hingga menjadi bubur. 2. Rebus ayam, suwir-suwir. 3. Tumis bawang putih, tambahkan ke bubur. 4. Sajikan dengan ayam suwir, daun bawang, seledri.",
			CookingTime: 45, Difficulty: "medium",
			Ingredients: []string{"nasi", "ayam", "bawang putih", "jahe", "garam", "daun bawang", "seledri"},
		},
		{
			Title: "Sambal Goreng Kentang",
			Instructions: "1. Goreng kentang hingga matang. 2. Haluskan cabai, bawang merah, bawang putih. 3. Tumis bumbu, masukkan kentang. 4. Tambahkan kecap manis, garam, gula.",
			CookingTime: 30, Difficulty: "easy",
			Ingredients: []string{"kentang", "cabai", "bawang merah", "bawang putih", "kecap manis", "garam", "gula", "minyak goreng"},
		},
		{
			Title: "Jagung Susu Keju",
			Instructions: "1. Rebus atau bakar jagung hingga matang. 2. Oles dengan mentega. 3. Taburi gula, kemudian siram dengan susu kental manis. 4. Taburi keju parut di atasnya.",
			CookingTime: 20, Difficulty: "easy",
			Ingredients: []string{"jagung", "mentega", "susu", "keju", "gula"},
		},
		{
			Title: "Tumis Buncis Wortel",
			Instructions: "1. Tumis bawang putih dan bawang merah. 2. Masukkan wortel, masak sebentar. 3. Tambahkan buncis, saus tiram, garam. 4. Aduk rata hingga sayuran matang namun masih renyah.",
			CookingTime: 15, Difficulty: "easy",
			Ingredients: []string{"buncis", "wortel", "bawang putih", "bawang merah", "saus tiram", "garam", "minyak goreng"},
		},
	}

	for _, rd := range recipes {
		slug := strings.ToLower(strings.ReplaceAll(rd.Title, " ", "-"))
		recipe := models.Recipe{
			Title:        rd.Title,
			Slug:         slug,
			Instructions: rd.Instructions,
			CookingTime:  rd.CookingTime,
			Difficulty:   rd.Difficulty,
		}
		if err := database.DB.Create(&recipe).Error; err != nil {
			continue
		}
		for _, ingName := range rd.Ingredients {
			var ing models.Ingredient
			if err := database.DB.Where("name = ?", ingName).First(&ing).Error; err != nil {
				continue
			}
			database.DB.Create(&models.RecipeIngredient{
				RecipeID:     recipe.ID,
				IngredientID: ing.ID,
				IsOptional:   false,
			})
		}
	}
}