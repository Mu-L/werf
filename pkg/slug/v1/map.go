package v1

var mapping = map[string]string{
	"0": "0",
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"²": "2",
	"³": "3",
	"&": "-",
	"(": "-",
	"{": "-",
	"}": "-",
	")": "-",
	"-": "-",
	"_": "-",
	"[": "-",
	"]": "-",
	",": "-",
	".": "-",
	";": "-",
	":": "-",
	"=": "-",
	"+": "-",
	"<": "-",
	">": "-",
	"~": "-",
	"µ": "u",
	"A": "a",
	"B": "b",
	"C": "c",
	"D": "d",
	"E": "e",
	"F": "f",
	"G": "g",
	"H": "h",
	"I": "i",
	"J": "j",
	"K": "k",
	"L": "l",
	"M": "m",
	"N": "n",
	"O": "o",
	"P": "p",
	"Q": "q",
	"R": "r",
	"S": "s",
	"T": "t",
	"U": "u",
	"V": "v",
	"W": "w",
	"X": "x",
	"Y": "y",
	"Z": "z",
	"a": "a",
	"b": "b",
	"c": "c",
	"d": "d",
	"e": "e",
	"f": "f",
	"g": "g",
	"h": "h",
	"i": "i",
	"j": "j",
	"k": "k",
	"l": "l",
	"m": "m",
	"n": "n",
	"o": "o",
	"p": "p",
	"q": "q",
	"r": "r",
	"s": "s",
	"t": "t",
	"u": "u",
	"v": "v",
	"w": "w",
	"x": "x",
	"y": "y",
	"z": "z",
	"À": "a",
	"Á": "a",
	"Â": "a",
	"Ã": "a",
	"Ä": "a",
	"Å": "a",
	"Æ": "ae",
	"Ç": "c",
	"È": "e",
	"É": "e",
	"Ê": "e",
	"Ë": "e",
	"Ì": "i",
	"Í": "i",
	"Î": "i",
	"Ï": "i",
	"Ð": "d",
	"Ñ": "n",
	"Ò": "o",
	"Ó": "o",
	"Ô": "o",
	"Õ": "o",
	"Ö": "o",
	"×": "x",
	"Ø": "o",
	"Ù": "u",
	"Ú": "u",
	"Û": "u",
	"Ü": "u",
	"Ý": "y",
	"Þ": "th",
	"ß": "b",
	"à": "a",
	"á": "a",
	"â": "a",
	"ã": "a",
	"ä": "a",
	"å": "a",
	"æ": "ae",
	"ç": "c",
	"è": "e",
	"é": "e",
	"ê": "e",
	"ë": "e",
	"ì": "i",
	"í": "i",
	"î": "i",
	"ï": "i",
	"ð": "o",
	"ñ": "n",
	"ò": "o",
	"ó": "o",
	"ô": "o",
	"õ": "o",
	"ö": "o",
	"ø": "o",
	"ù": "u",
	"ú": "u",
	"û": "u",
	"ü": "u",
	"ý": "y",
	"þ": "th",
	"ÿ": "y",
	"Ā": "a",
	"ā": "a",
	"Ă": "a",
	"ă": "a",
	"Ą": "a",
	"ą": "a",
	"Ć": "c",
	"ć": "c",
	"Ĉ": "c",
	"ĉ": "c",
	"Ċ": "c",
	"ċ": "c",
	"Č": "c",
	"č": "c",
	"Ď": "d",
	"ď": "d",
	"Đ": "d",
	"đ": "d",
	"Ē": "e",
	"ē": "e",
	"Ĕ": "e",
	"ĕ": "e",
	"Ė": "e",
	"ė": "e",
	"Ę": "e",
	"ę": "e",
	"Ě": "e",
	"ě": "e",
	"Ĝ": "g",
	"ĝ": "g",
	"Ğ": "g",
	"ğ": "g",
	"Ġ": "g",
	"ġ": "g",
	"Ģ": "g",
	"ģ": "g",
	"Ĥ": "h",
	"ĥ": "h",
	"Ħ": "h",
	"ħ": "h",
	"Ĩ": "i",
	"ĩ": "i",
	"Ī": "i",
	"ī": "i",
	"Ĭ": "i",
	"ĭ": "i",
	"Į": "i",
	"į": "i",
	"İ": "l",
	"ı": "l",
	"Ĳ": "ij",
	"ĳ": "ij",
	"Ĵ": "j",
	"ĵ": "j",
	"Ķ": "k",
	"ķ": "k",
	"ĸ": "k",
	"Ĺ": "l",
	"ĺ": "l",
	"Ļ": "l",
	"ļ": "l",
	"Ľ": "l",
	"ľ": "l",
	"Ŀ": "l",
	"ŀ": "l",
	"Ł": "l",
	"ł": "l",
	"Ń": "n",
	"ń": "n",
	"Ņ": "n",
	"ņ": "n",
	"Ň": "n",
	"ň": "n",
	"ŉ": "n",
	"Ŋ": "n",
	"ŋ": "n",
	"Ō": "o",
	"ō": "o",
	"Ŏ": "o",
	"ŏ": "o",
	"Ő": "o",
	"ő": "o",
	"Œ": "oe",
	"œ": "oe",
	"Ŕ": "r",
	"ŕ": "r",
	"Ŗ": "r",
	"ŗ": "r",
	"Ř": "r",
	"ř": "r",
	"Ś": "s",
	"ś": "s",
	"Ŝ": "s",
	"ŝ": "s",
	"Ş": "s",
	"ş": "s",
	"Š": "s",
	"š": "s",
	"Ţ": "t",
	"ţ": "t",
	"Ť": "t",
	"ť": "t",
	"Ŧ": "t",
	"ŧ": "t",
	"Ũ": "u",
	"ũ": "u",
	"Ū": "u",
	"ū": "u",
	"Ŭ": "u",
	"ŭ": "u",
	"Ů": "u",
	"ů": "u",
	"Ű": "u",
	"ű": "u",
	"Ų": "u",
	"ų": "u",
	"Ŵ": "w",
	"ŵ": "w",
	"Ŷ": "y",
	"ŷ": "y",
	"Ÿ": "y",
	"Ź": "z",
	"ź": "z",
	"Ż": "z",
	"ż": "z",
	"Ž": "z",
	"ž": "z",
	"ſ": "s",
	"ƀ": "b",
	"Ɓ": "b",
	"Ƅ": "b",
	"ƅ": "b",
	"Ƈ": "c",
	"ƈ": "c",
	"Ɖ": "d",
	"Ɗ": "d",
	"Ƌ": "nd",
	"ƌ": "nd",
	"Ɛ": "e",
	"Ƒ": "f",
	"ƒ": "f",
	"Ɠ": "g",
	"Ɩ": "i",
	"Ɨ": "i",
	"Ƙ": "k",
	"ƙ": "k",
	"Ɯ": "w",
	"Ɲ": "n",
	"ƞ": "n",
	"Ɵ": "o",
	"Ơ": "o",
	"ơ": "o",
	"Ƥ": "p",
	"ƥ": "p",
	"Ʀ": "r",
	"Ƨ": "s",
	"ƨ": "s",
	"Ʃ": "s",
	"ƪ": "t",
	"ƫ": "t",
	"Ƭ": "t",
	"ƭ": "t",
	"Ʈ": "t",
	"Ư": "u",
	"ư": "u",
	"Ʊ": "u",
	"Ʋ": "u",
	"Ƶ": "z",
	"ƶ": "z",
	"Ʒ": "z",
	"Ƽ": "q",
	"ƽ": "q",
	"Ǆ": "dz",
	"ǅ": "dz",
	"ǆ": "dz",
	"Ǉ": "lj",
	"ǈ": "lj",
	"ǉ": "lj",
	"Ǌ": "nj",
	"ǋ": "nj",
	"ǌ": "nj",
	"Ǎ": "a",
	"ǎ": "a",
	"Ǐ": "i",
	"ǐ": "i",
	"Ǒ": "o",
	"ǒ": "o",
	"Ǔ": "u",
	"ǔ": "u",
	"Ǖ": "u",
	"ǖ": "u",
	"Ǘ": "u",
	"ǘ": "u",
	"Ǚ": "u",
	"ǚ": "u",
	"Ǜ": "u",
	"ǜ": "u",
	"ǝ": "e",
	"Ǟ": "a",
	"ǟ": "a",
	"Ǡ": "a",
	"ǡ": "a",
	"Ǣ": "ae",
	"ǣ": "ae",
	"Ǥ": "g",
	"ǥ": "g",
	"Ǧ": "g",
	"ǧ": "g",
	"Ǩ": "k",
	"ǩ": "k",
	"Ǫ": "o",
	"ǫ": "o",
	"Ǭ": "o",
	"ǭ": "o",
	"ǰ": "j",
	"Ǳ": "dz",
	"ǲ": "dz",
	"ǳ": "dz",
	"Ǵ": "g",
	"ǵ": "g",
	"Ǹ": "n",
	"ǹ": "n",
	"Ǻ": "a",
	"ǻ": "a",
	"Ǽ": "ae",
	"ǽ": "ae",
	"Ǿ": "o",
	"ǿ": "o",
	"Ȁ": "a",
	"ȁ": "a",
	"Ȃ": "a",
	"ȃ": "a",
	"Ȅ": "e",
	"ȅ": "e",
	"Ȇ": "e",
	"ȇ": "e",
	"Ȉ": "i",
	"ȉ": "i",
	"Ȋ": "i",
	"ȋ": "i",
	"Ȍ": "o",
	"ȍ": "o",
	"Ȏ": "o",
	"ȏ": "o",
	"Ȑ": "r",
	"ȑ": "r",
	"Ȓ": "r",
	"ȓ": "r",
	"Ȕ": "u",
	"ȕ": "u",
	"Ȗ": "u",
	"ȗ": "u",
	"Ș": "s",
	"ș": "s",
	"Ț": "t",
	"ț": "t",
	"Ȟ": "h",
	"ȟ": "h",
	"Ƞ": "n",
	"Ȥ": "z",
	"ȥ": "z",
	"Ȧ": "a",
	"ȧ": "a",
	"Ȩ": "e",
	"ȩ": "e",
	"Ȫ": "o",
	"ȫ": "o",
	"Ȭ": "o",
	"ȭ": "o",
	"Ȯ": "o",
	"ȯ": "o",
	"Ȱ": "o",
	"ȱ": "o",
	"Ȳ": "y",
	"ȳ": "y",
	"ȷ": "j",
	"ȸ": "db",
	"ȹ": "qp",
	"Ⱥ": "a",
	"Ȼ": "c",
	"ȼ": "c",
	"Ƚ": "l",
	"Ⱦ": "l",
	"ȿ": "s",
	"ɀ": "z",
	"Ƀ": "b",
	"Ʉ": "u",
	"Ʌ": "v",
	"Ɇ": "e",
	"ɇ": "e",
	"Ɉ": "j",
	"ɉ": "j",
	"Ɋ": "q",
	"ɋ": "q",
	"Ɍ": "r",
	"ɍ": "r",
	"Ɏ": "y",
	"ɏ": "y",
	"ɑ": "a",
	"ɗ": "d",
	"ɠ": "g",
	"ɡ": "g",
	"ɢ": "g",
	"ɯ": "w",
	"ɱ": "m",
	"ɵ": "o",
	"ɶ": "oe",
	"ᴀ": "a",
	"ᴁ": "ae",
	"ᴂ": "ae",
	"ᴃ": "b",
	"ᴄ": "c",
	"ᴅ": "d",
	"ᴆ": "d",
	"ᴇ": "e",
	"ᴊ": "j",
	"ᴋ": "k",
	"ᴌ": "l",
	"ᴍ": "m",
	"ᴎ": "n",
	"ᴏ": "o",
	"Ḁ": "a",
	"ḁ": "a",
	"Ḃ": "b",
	"ḃ": "b",
	"Ḅ": "b",
	"ḅ": "b",
	"Ḇ": "b",
	"ḇ": "b",
	"Ḉ": "c",
	"ḉ": "c",
	"Ḋ": "d",
	"ḋ": "d",
	"Ḍ": "d",
	"ḍ": "d",
	"Ḏ": "d",
	"ḏ": "d",
	"Ḑ": "d",
	"ḑ": "d",
	"Ḓ": "d",
	"ḓ": "d",
	"Ḕ": "e",
	"ḕ": "e",
	"Ḗ": "e",
	"ḗ": "e",
	"Ḙ": "e",
	"ḙ": "e",
	"Ḛ": "e",
	"ḛ": "e",
	"Ḝ": "e",
	"ḝ": "e",
	"Ḟ": "f",
	"ḟ": "f",
	"Ḡ": "g",
	"ḡ": "g",
	"Ḣ": "h",
	"ḣ": "h",
	"Ḥ": "h",
	"ḥ": "h",
	"Ḧ": "h",
	"ḧ": "h",
	"Ḩ": "h",
	"ḩ": "h",
	"Ḫ": "h",
	"ḫ": "h",
	"Ḭ": "i",
	"ḭ": "i",
	"Ḯ": "i",
	"ḯ": "i",
	"Ḱ": "k",
	"ḱ": "k",
	"Ḳ": "k",
	"ḳ": "k",
	"Ḵ": "k",
	"ḵ": "k",
	"Ḷ": "l",
	"ḷ": "l",
	"Ḹ": "l",
	"ḹ": "l",
	"Ḻ": "l",
	"ḻ": "l",
	"Ḽ": "l",
	"ḽ": "l",
	"Ḿ": "m",
	"ḿ": "m",
	"Ṁ": "m",
	"ṁ": "m",
	"Ṃ": "m",
	"ṃ": "m",
	"Ṅ": "n",
	"ṅ": "n",
	"Ṇ": "n",
	"ṇ": "n",
	"Ṉ": "n",
	"ṉ": "n",
	"Ṋ": "n",
	"ṋ": "n",
	"Ṍ": "o",
	"ṍ": "o",
	"Ṏ": "o",
	"ṏ": "o",
	"Ṑ": "o",
	"ṑ": "o",
	"Ṓ": "o",
	"ṓ": "o",
	"Ṕ": "p",
	"ṕ": "p",
	"Ṗ": "p",
	"ṗ": "p",
	"Ṙ": "r",
	"ṙ": "r",
	"Ṛ": "r",
	"ṛ": "r",
	"Ṝ": "r",
	"ṝ": "r",
	"Ṟ": "r",
	"ṟ": "r",
	"Ṡ": "s",
	"ṡ": "s",
	"Ṣ": "s",
	"ṣ": "s",
	"Ṥ": "s",
	"ṥ": "s",
	"Ṧ": "s",
	"ṧ": "s",
	"Ṩ": "s",
	"ṩ": "s",
	"Ṫ": "t",
	"ṫ": "t",
	"Ṭ": "t",
	"ṭ": "t",
	"Ṯ": "t",
	"ṯ": "t",
	"Ṱ": "t",
	"ṱ": "t",
	"Ṳ": "u",
	"ṳ": "u",
	"Ṵ": "u",
	"ṵ": "u",
	"Ṷ": "u",
	"ṷ": "u",
	"Ṹ": "u",
	"ṹ": "u",
	"Ṻ": "u",
	"ṻ": "u",
	"Ṽ": "v",
	"ṽ": "v",
	"Ṿ": "v",
	"ṿ": "v",
	"Ẁ": "w",
	"ẁ": "w",
	"Ẃ": "w",
	"ẃ": "w",
	"Ẅ": "w",
	"ẅ": "w",
	"Ẇ": "w",
	"ẇ": "w",
	"Ẉ": "w",
	"ẉ": "w",
	"Ẋ": "x",
	"ẋ": "x",
	"Ẍ": "x",
	"ẍ": "x",
	"Ẏ": "y",
	"ẏ": "y",
	"Ẑ": "z",
	"ẑ": "z",
	"Ẓ": "z",
	"ẓ": "z",
	"Ẕ": "z",
	"ẕ": "z",
	"ẖ": "h",
	"ẗ": "t",
	"ẘ": "w",
	"ẙ": "y",
	"ẚ": "a",
	"Ạ": "a",
	"ạ": "a",
	"Ả": "a",
	"ả": "a",
	"Ấ": "a",
	"ấ": "a",
	"Ầ": "a",
	"ầ": "a",
	"Ẩ": "a",
	"ẩ": "a",
	"Ẫ": "a",
	"ẫ": "a",
	"Ậ": "a",
	"ậ": "a",
	"Ắ": "a",
	"ắ": "a",
	"Ằ": "a",
	"ằ": "a",
	"Ẳ": "a",
	"ẳ": "a",
	"Ẵ": "a",
	"ẵ": "a",
	"Ặ": "a",
	"ặ": "a",
	"Ẹ": "e",
	"ẹ": "e",
	"Ẻ": "e",
	"ẻ": "e",
	"Ẽ": "e",
	"ẽ": "e",
	"Ế": "e",
	"ế": "e",
	"Ề": "e",
	"ề": "e",
	"Ể": "e",
	"ể": "e",
	"Ễ": "e",
	"ễ": "e",
	"Ệ": "e",
	"ệ": "e",
	"Ọ": "o",
	"ọ": "o",
	"Ỏ": "o",
	"ỏ": "o",
	"Ố": "o",
	"ố": "o",
	"Ồ": "o",
	"ồ": "o",
	"Ổ": "o",
	"ổ": "o",
	"Ỗ": "o",
	"ỗ": "o",
	"Ộ": "o",
	"ộ": "o",
	"Ớ": "o",
	"ớ": "o",
	"Ờ": "o",
	"ờ": "o",
	"Ở": "o",
	"ở": "o",
	"Ỡ": "o",
	"ỡ": "o",
	"Ợ": "o",
	"ợ": "o",
	"Ụ": "u",
	"ụ": "u",
	"Ủ": "u",
	"ủ": "u",
	"Ứ": "u",
	"ứ": "u",
	"Ừ": "u",
	"ừ": "u",
	"Ử": "u",
	"ử": "u",
	"Ữ": "u",
	"ữ": "u",
	"Ự": "u",
	"ự": "u",
	"Ỳ": "y",
	"ỳ": "y",
	"Ỵ": "y",
	"ỵ": "y",
	"Ỷ": "y",
	"ỷ": "y",
	"Ỹ": "y",
	"ỹ": "y",
	"ₐ": "a",
	"ₑ": "e",
	"ₒ": "o",
	"ₓ": "x",
	"℀": "ac",
	"℁": "as",
	"ℂ": "c",
	"℃": "c",
	"℄": "c",
	"℅": "co",
	"℆": "cu",
	"℉": "f",
	"ℊ": "g",
	"ℋ": "h",
	"ℌ": "h",
	"ℍ": "h",
	"ℎ": "h",
	"ℏ": "h",
	"ℕ": "n",
	"℗": "p",
	"ℙ": "p",
	"ℚ": "q",
	"ℛ": "r",
	"ℜ": "r",
	"ℝ": "r",
	"℞": "px",
	"℟": "r",
	"℠": "sm",
	"℡": "tel",
	"™": "tm",
	"ℤ": "z",
	"K": "k",
	"Å": "a",
	"ℬ": "b",
	"℮": "e",
	"ℯ": "e",
	"⒐": "9",
	"⒑": "10",
	"⒒": "11",
	"⒓": "12",
	"⒔": "13",
	"⒕": "14",
	"⒖": "15",
	"⒗": "16",
	"⒘": "17",
	"⒙": "18",
	"⒚": "19",
	"⒛": "20",
	"⒜": "a",
	"⒝": "b",
	"⒞": "c",
	"⒟": "d",
	"⒠": "e",
	"⒡": "f",
	"⒢": "g",
	"⒣": "h",
	"⒤": "i",
	"⒥": "j",
	"⒦": "k",
	"⒧": "l",
	"⒨": "m",
	"⒩": "n",
	"⒪": "o",
	"⒫": "p",
	"⒬": "q",
	"⒭": "r",
	"⒮": "s",
	"⒯": "t",
	"⒰": "u",
	"⒱": "v",
	"⒲": "w",
	"⒳": "x",
	"⒴": "y",
	"⒵": "z",
	"Ⓐ": "a",
	"Ⓑ": "b",
	"Ⓒ": "c",
	"Ⓓ": "d",
	"Ⓔ": "e",
	"Ⓕ": "f",
	"Ⓖ": "g",
	"Ⓗ": "h",
	"Ⓘ": "i",
	"Ⓙ": "j",
	"Ⓚ": "k",
	"Ⓛ": "l",
	"Ⓜ": "m",
	"Ⓝ": "n",
	"Ⓞ": "o",
	"Ⓟ": "p",
	"Ⓠ": "q",
	"Ⓡ": "r",
	"Ⓢ": "s",
	"Ⓣ": "t",
	"Ⓤ": "u",
	"Ⓥ": "v",
	"Ⓦ": "w",
	"Ⓧ": "x",
	"Ⓨ": "y",
	"Ⓩ": "z",
	"ⓐ": "a",
	"ⓑ": "b",
	"ⓒ": "c",
	"ⓓ": "d",
	"ⓔ": "e",
	"ⓕ": "f",
	"ⓖ": "g",
	"ⓗ": "h",
	"ⓘ": "i",
	"ⓙ": "j",
	"ⓚ": "k",
	"ⓛ": "l",
	"ⓜ": "m",
	"ⓝ": "n",
	"ⓞ": "o",
	"ⓟ": "p",
	"ⓠ": "q",
	"ⓡ": "r",
	"ⓢ": "s",
	"ⓣ": "t",
	"ⓤ": "u",
	"ⓥ": "v",
	"ⓦ": "w",
	"ⓧ": "x",
	"ⓨ": "y",
	"ⓩ": "z",
	"⓪": "0",
	"ꜰ": "f",
	"ꜱ": "s",
	"Ꜳ": "aa",
	"ꜳ": "aa",
	"Ꜵ": "ao",
	"ꜵ": "ao",
	"Ꜷ": "aj",
	"ꜷ": "aj",
	"Ꜹ": "av",
	"ꜹ": "av",
	"Ꜻ": "af",
	"ꜻ": "af",
	"Ꜽ": "af",
	"ꜽ": "af",
	"А": "a",
	"Б": "b",
	"В": "v",
	"Г": "g",
	"Д": "d",
	"Е": "ye",
	"Ё": "yo",
	"Ж": "zh",
	"З": "z",
	"И": "i",
	"Й": "j",
	"К": "k",
	"Л": "l",
	"М": "m",
	"Н": "n",
	"О": "o",
	"П": "p",
	"Р": "r",
	"С": "s",
	"Т": "t",
	"У": "u",
	"Ф": "f",
	"Х": "h",
	"Ц": "c",
	"Ч": "ch",
	"Ш": "sh",
	"Щ": "shch",
	"Ъ": "",
	"Ы": "y",
	"Ь": "",
	"Э": "e",
	"Ю": "yu",
	"Я": "ya",
	"а": "a",
	"б": "b",
	"в": "v",
	"г": "g",
	"д": "d",
	"е": "ye",
	"ё": "yo",
	"ж": "zh",
	"з": "z",
	"и": "i",
	"й": "j",
	"к": "k",
	"л": "l",
	"м": "m",
	"н": "n",
	"о": "o",
	"п": "p",
	"р": "r",
	"с": "s",
	"т": "t",
	"у": "u",
	"ф": "f",
	"х": "h",
	"ц": "c",
	"ч": "ch",
	"ш": "sh",
	"щ": "shch",
	"ъ": "",
	"ы": "y",
	"ь": "",
	"э": "e",
	"ю": "yu",
	"я": "ya",
	"Ａ": "a",
	"Ｂ": "b",
	"Ｃ": "c",
	"Ｄ": "d",
	"Ｅ": "e",
	"Ｆ": "f",
	"Ｇ": "g",
	"Ｈ": "h",
	"Ｉ": "i",
	"Ｊ": "j",
	"Ｋ": "k",
	"Ｌ": "l",
	"Ｍ": "m",
	"Ｎ": "n",
	"Ｏ": "o",
	"Ｐ": "p",
	"Ｑ": "q",
	"Ｒ": "r",
	"Ｓ": "s",
	"Ｔ": "t",
	"Ｕ": "u",
	"Ｖ": "v",
	"Ｗ": "w",
	"Ｘ": "x",
	"Ｙ": "y",
	"Ｚ": "z",
	"＿": "-",
	"ａ": "a",
	"ｂ": "b",
	"ｃ": "c",
	"ｄ": "d",
	"ｅ": "e",
	"ｆ": "f",
	"ｇ": "g",
	"ｈ": "h",
	"ｉ": "i",
	"ｊ": "j",
	"ｋ": "k",
	"ｌ": "l",
	"ｍ": "m",
	"ｎ": "n",
	"ｏ": "o",
	"ｐ": "p",
	"ｑ": "q",
	"ｒ": "r",
	"ｓ": "s",
	"ｔ": "t",
	"ｕ": "u",
	"ｖ": "v",
	"ｗ": "w",
	"ｘ": "x",
	"ｙ": "y",
	"ｚ": "z",
	"～": "-",
	"/": "-",
	" ": "-",
}
