package gherkin

// Builtin dialects for af (Afrikaans), am (Armenian), ar (Arabic), ast (Asturian), az (Azerbaijani), bg (Bulgarian), bm (Malay), bs (Bosnian), ca (Catalan), cs (Czech), cy-GB (Welsh), da (Danish), de (German), el (Greek), em (Emoji), en (English), en-Scouse (Scouse), en-au (Australian), en-lol (LOLCAT), en-old (Old English), en-pirate (Pirate), eo (Esperanto), es (Spanish), et (Estonian), fa (Persian), fi (Finnish), fr (French), ga (Irish), gj (Gujarati), gl (Galician), he (Hebrew), hi (Hindi), hr (Croatian), ht (Creole), hu (Hungarian), id (Indonesian), is (Icelandic), it (Italian), ja (Japanese), jv (Javanese), ka (Georgian), kn (Kannada), ko (Korean), lt (Lithuanian), lu (Luxemburgish), lv (Latvian), mk-Cyrl (Macedonian), mk-Latn (Macedonian (Latin)), mn (Mongolian), nl (Dutch), no (Norwegian), pa (Panjabi), pl (Polish), pt (Portuguese), ro (Romanian), ru (Russian), sk (Slovak), sl (Slovenian), sr-Cyrl (Serbian), sr-Latn (Serbian (Latin)), sv (Swedish), ta (Tamil), th (Thai), tl (Telugu), tlh (Klingon), tr (Turkish), tt (Tatar), uk (Ukrainian), ur (Urdu), uz (Uzbek), vi (Vietnamese), zh-CN (Chinese simplified), zh-TW (Chinese traditional)
func GherkinDialectsBuildin() GherkinDialectProvider {
	return buildinDialects
}

const (
	feature         = "feature"
	background      = "background"
	scenario        = "scenario"
	scenarioOutline = "scenarioOutline"
	examples        = "examples"
	given           = "given"
	when            = "when"
	then            = "then"
	and             = "and"
	but             = "but"
)

var buildinDialects = gherkinDialectMap{
	"en": &GherkinDialect{
		"en", "English", "English", map[string][]string{
			feature: []string{
				"Flow",
			},
			background: []string{
				"Background",
			},
			scenario: []string{
				"Scenario",
				"Dialog",
			},
			scenarioOutline: []string{
				"Scenario Outline",
				"Scenario Template",
			},
			examples: []string{
				"Examples",
				"Scenarios",
			},
			given: []string{
				"Given ",
			},
			when: []string{
				"When ",
			},
			then: []string{
				"Then ",
			},
			and: []string{
				"And ",
			},
			but: []string{
				"With ",
				"Timeout ",
			},
		},
	},
}
