package sakamotocommands

// Constant used to determine the validity of a youtube link
const (
	// VIDEO const
	VIDEO = 1

	// PLAYLIST const
	PLAYLIST = 2

	// NONVALIDLINK const
	NONVALIDLINK = 0
)

var (
	// SearchSortOptions list Sorting options for imgur api
	SearchSortOptions = []string{
		"top",
		"rising",
		"viral",
		"time",
	}

	// SearchRangeOptions list time range options for imgur api
	SearchRangeOptions = []string{
		"day",
		"week",
		"month",
		"year",
		"all",
	}

	// SEARCHSORT : current imgur sort option
	SEARCHSORT = "top"

	// SEARCHRANGE : current imgur range option
	SEARCHRANGE = "all"
)

// Soundbox : Path map for soundbox -> map[command]path
var Soundbox = map[string]string{
	"JEAGER":      "SoundBox/JEAGER.mp3",
	"JEANNE":      "SoundBox/JEANNE.mp3",
	"cklair":      "SoundBox/mouicklair.mp3",
	"whee":        "SoundBox/whee.mp3",
	"bruh":        "SoundBox/bruh.mp3",
	"oof":         "SoundBox/roblox.mp3",
	"marionon":    "SoundBox/mario_non.mp3",
	"thomas":      "SoundBox/thomas.mp3",
	"sanic":       "SoundBox/sanic.mp3",
	"running":     "SoundBox/why_are_u_running.mp3",
	"SPITONHIM":   "SoundBox/SPITONHIM.mp3",
	"dewae":       "SoundBox/dewae.mp3",
	"johncena":    "SoundBox/johncena.mp3",
	"sensibilite": "SoundBox/sensibilite.mp3",
	"qualifie":    "SoundBox/qualifie.mp3",
	"vega":        "SoundBox/vega.mp3",
	"rengar":      "SoundBox/rengar.mp3",
	"quenouille":  "SoundBox/quenouille.mp3",
	"mince":       "SoundBox/oh_mince.mp3",
	"troposphere": "SoundBox/reussite.mp3",
	"ratz":        "SoundBox/ratz.mp3",
	"doremi":      "SoundBox/doremi.mp3",
	"guile":       "SoundBox/guile.mp3",
	"zombie":      "SoundBox/zombie.mp3",
	"ally":        "SoundBox/ally.mp3",
	"boule":       "SoundBox/la-boule-magique.mp3",
}
