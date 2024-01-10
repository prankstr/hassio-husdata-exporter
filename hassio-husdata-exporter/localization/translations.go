package localization

import "golang.org/x/text/language"

var (
	SwedishCatalog = map[string]string{
		"Radiator Return":        "Returledning",
		"Heat Carrier Return":    "Returledning Intern",
		"Heat Carrier Forward":   "Framledning",
		"Brine In / Condensor":   "Köldbärare In",
		"Brine Out / Evaporator": "Köldbärare Ut",
		"Outdoor":                "Ute",
		"Indoor":                 "Rum",
		"Hot Water Top":          "Varmvatten Topp",
		"Hot Water Mid":          "Varmvatten Laddning",
		"Hot Gas":                "Hetgas",
		"Suction Gas":            "Suggas",
		"Liquid Flow":            "Vätskeledning",
		"Air Intake":             "Luftintag",
		"Pool":                   "Pool",
		"Radiator Forward 2":     "Framledning 2",
		"Radiator Return 2":      "Returledning 2",
		"Compressor":             "Kompressor",
		"Add Heat":               "Tillskott",
		"Hot Water":              "Varmvattenladdning",
		"Add Heat Step 1":        "Tillskott Steg I",
		"Add Heat Step 2":        "Tillskott Steg II",
		"Pump Heat Circuit":      "Värmebärarpump",
		"Pump Cold Circuit":      "Köldbärarpump",
		"Pump Radiator":          "Cirkulationspump",
		"Switch Valve 1":         "Växelventil 1",
		"Switch Valve 2":         "Växelventil 2",
		"Heating Cable":          "Värmekabel",
		"Crank Case Heater":      "Vevhusvärmare",
		"Fan":                    "Fläkt",
		"High Pressostat":        "Pressostat Hög",
		"Low Pressostat":         "Pressostat Låg",
		"Heat Set 1 CurveL":      "Kurvlutning",
		"Heat Set 1 CurveL 2":    "Kurvlutning 2",
		"Heat Set 3 Parallel":    "Förskjutning Värmekurva",
		"Heat Set 3 Parallel 2":  "Förskjutning Värmekurva 2",
	}
)

func GetTranslation(lang language.Tag, key string) string {
	var catalog map[string]string
	switch lang {
	case language.Swedish:
		catalog = SwedishCatalog
	default:
		catalog = SwedishCatalog
	}

	if msg, ok := catalog[key]; ok {
		return msg
	}

	return key
}
