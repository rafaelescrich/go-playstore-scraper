package locales

// Language is the language code which the playstore supports.
type Language string

// DefLanguage doesn't represent a specific language, it's equivalent to
// setting the language to the one of the host making the requests, as
// seen by the store.
const DefLanguage Language = ""

const (
	Afrikaans                Language = "af"
	Amharic                  Language = "am"
	Arabic                   Language = "ar"
	Azeri                    Language = "az"
	Belarusian               Language = "be"
	Bulgarian                Language = "bg"
	Bengali                  Language = "bn"
	Bosnian                  Language = "bs"
	Catalan                  Language = "ca"
	Czech                    Language = "cs"
	Danish                   Language = "da"
	German                   Language = "de"
	GermanAustria            Language = "de_AT"
	GermanSwitzerland        Language = "de_CH"
	Greek                    Language = "el"
	English                  Language = "en"
	EnglishAustralia         Language = "en_AU"
	EnglishCanada            Language = "en_CA"
	EnglishUK                Language = "en_GB"
	EnglishIreland           Language = "en_IE"
	EnglishIndia             Language = "en_IN"
	EnglishSingapore         Language = "en_SG"
	EnglishUS                Language = "en_US"
	EnglishSouthAfrica       Language = "en_ZA"
	Spanish                  Language = "es"
	SpanishLatinAmerica      Language = "es_419"
	SpanishArgentina         Language = "es_AR"
	SpanishBolivia           Language = "es_BO"
	SpanishChile             Language = "es_CL"
	SpanishColombia          Language = "es_CO"
	SpanishCostaRica         Language = "es_CR"
	SpanishDominicanRepublic Language = "es_DO"
	SpanishEcuador           Language = "es_EC"
	SpanishSpain             Language = "es-ES"
	SpanishGuatemala         Language = "es_GT"
	SpanishHonduras          Language = "es_HN"
	SpanishMexico            Language = "es_MX"
	SpanishNicaragua         Language = "es_NI"
	SpanishPanama            Language = "es_PA"
	SpanishPeru              Language = "es_PE"
	SpanishPuertoRico        Language = "es_PR"
	SpanishParaguay          Language = "es_PY"
	SpanishElSalvador        Language = "es_SV"
	SpanishUS                Language = "es_US"
	SpanishUruguay           Language = "es_UY"
	SpanishVenezuela         Language = "es_VE"
	Estonian                 Language = "et"
	Basque                   Language = "eu"
	Farsi                    Language = "fa"
	Finnish                  Language = "fi"
	Filipino                 Language = "fil"
	French                   Language = "fr"
	FrenchCanada             Language = "fr_CA"
	FrenchFrance             Language = "fr_FR"
	FrenchSwitzerland        Language = "fr_CH"
	Galician                 Language = "gl"
	SwissGerman              Language = "gsw"
	Gujarati                 Language = "gu"
	Hebrew                   Language = "he"
	Hindi                    Language = "hi"
	Croatian                 Language = "hr"
	Hungarian                Language = "hu"
	Armenian                 Language = "hy"
	Indonesian               Language = "id"
	Indonesian2              Language = "in"
	Icelandic                Language = "is"
	Italian                  Language = "it"
	Japanese                 Language = "ja"
	Georgian                 Language = "ka"
	Kazakh                   Language = "kk"
	Khmer                    Language = "km"
	Kannada                  Language = "kn"
	Korean                   Language = "ko"
	Kyrgyz                   Language = "ky"
	Lingala                  Language = "ln"
	Lao                      Language = "lo"
	Lithuanian               Language = "lt"
	Latvian                  Language = "lv"
	Macedonian               Language = "mk"
	Malayalam                Language = "ml"
	Mongolian                Language = "mn"
	Marathi                  Language = "mr"
	Malay                    Language = "ms"
	Burmese                  Language = "my"
	NorwegianBokmal          Language = "nb"
	Nepali                   Language = "ne"
	Dutch                    Language = "nl"
	Norwegian                Language = "no"
	Punjabi                  Language = "pa"
	Polish                   Language = "pl"
	Portuguese               Language = "pt"
	PortugueseBrazil         Language = "pt_BR"
	PortuguesePortugal       Language = "pt_PT"
	Romanian                 Language = "ro"
	Russian                  Language = "ru"
	Sinhala                  Language = "si"
	Slovak                   Language = "sk"
	Slovenian                Language = "sl"
	Albanian                 Language = "sq"
	Serbian                  Language = "sr"
	SerbianLatin             Language = "sr_LATN"
	Swedish                  Language = "sv"
	Swahili                  Language = "sw"
	Tamil                    Language = "ta"
	Telugu                   Language = "te"
	Thai                     Language = "th"
	Tagalog                  Language = "tl"
	Turkish                  Language = "tr"
	Ukrainian                Language = "uk"
	Urdu                     Language = "ur"
	Uzbek                    Language = "uz"
	Vietnamese               Language = "vi"
	Chinese                  Language = "zh"
	ChineseSimplified        Language = "zh_CN"
	ChineseHongKong          Language = "zh_HK"
	ChineseTraditional       Language = "zh_TW"
	Zulu                     Language = "zu"
)

// AllLanguages is an array of all the language codes available in this package
var AllLanguages = [...]Language{
	DefLanguage, Afrikaans, Amharic, Arabic,
	Azeri, Belarusian, Bulgarian, Bengali,
	Bosnian, Catalan, Czech, Danish,
	German, GermanAustria, GermanSwitzerland,
	Greek, English, EnglishAustralia,
	EnglishCanada, EnglishUK, EnglishIreland,
	EnglishIndia, EnglishSingapore, EnglishUS,
	EnglishSouthAfrica, Spanish, SpanishLatinAmerica,
	SpanishArgentina, SpanishBolivia, SpanishChile,
	SpanishColombia, SpanishCostaRica, SpanishDominicanRepublic,
	SpanishEcuador, SpanishSpain, SpanishGuatemala,
	SpanishHonduras, SpanishMexico, SpanishNicaragua,
	SpanishPanama, SpanishPeru, SpanishPuertoRico,
	SpanishParaguay, SpanishElSalvador, SpanishUS,
	SpanishUruguay, SpanishVenezuela, Estonian,
	Basque, Farsi, Finnish, Filipino,
	French, FrenchCanada, FrenchFrance, FrenchSwitzerland,
	Galician, SwissGerman, Gujarati, Hebrew,
	Hindi, Croatian, Hungarian, Armenian,
	Indonesian, Indonesian2, Icelandic, Italian,
	Japanese, Georgian, Kazakh, Khmer, Kannada,
	Korean, Kyrgyz, Lingala, Lao, Lithuanian,
	Latvian, Macedonian, Malayalam, Mongolian,
	Marathi, Malay, Burmese, NorwegianBokmal,
	Nepali, Dutch, Norwegian, Punjabi, Polish,
	Portuguese, PortugueseBrazil, PortuguesePortugal,
	Romanian, Russian, Sinhala, Slovak, Slovenian,
	Albanian, Serbian, SerbianLatin, Swedish,
	Swahili, Tamil, Telugu, Thai, Tagalog,
	Turkish, Ukrainian, Urdu, Uzbek, Vietnamese,
	Chinese, ChineseSimplified, ChineseHongKong,
	ChineseTraditional, Zulu,
}
