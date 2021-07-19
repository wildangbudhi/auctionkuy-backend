package utils

import (
	"strings"
)

type country struct {
	Name      string
	Code      string
	PhoneCode string
}

func NewCountryDictionary() (*Trie, map[string]string) {

	var trie *Trie = NewTrie()
	var maps map[string]string = make(map[string]string)

	for i := 0; i < len(PhoneNumberCode); i++ {
		var phoneCode string
		phoneCode = strings.ReplaceAll(PhoneNumberCode[i].PhoneCode, "+", "")
		phoneCode = strings.ReplaceAll(phoneCode, " ", "")
		var countryCode string = strings.ToLower(PhoneNumberCode[i].Code)

		trie.Insert(phoneCode)
		maps[countryCode] = phoneCode
	}

	return trie, maps

}

var PhoneNumberCode [242]country = [242]country{
	{
		Name:      "Afghanistan",
		PhoneCode: "+93",
		Code:      "AF",
	},
	{
		Name:      "Aland Islands",
		PhoneCode: "+358",
		Code:      "AX",
	},
	{
		Name:      "Albania",
		PhoneCode: "+355",
		Code:      "AL",
	},
	{
		Name:      "Algeria",
		PhoneCode: "+213",
		Code:      "DZ",
	},
	{
		Name:      "AmericanSamoa",
		PhoneCode: "+1 684",
		Code:      "AS",
	},
	{
		Name:      "Andorra",
		PhoneCode: "+376",
		Code:      "AD",
	},
	{
		Name:      "Angola",
		PhoneCode: "+244",
		Code:      "AO",
	},
	{
		Name:      "Anguilla",
		PhoneCode: "+1 264",
		Code:      "AI",
	},
	{
		Name:      "Antarctica",
		PhoneCode: "+672",
		Code:      "AQ",
	},
	{
		Name:      "Antigua and Barbuda",
		PhoneCode: "+1268",
		Code:      "AG",
	},
	{
		Name:      "Argentina",
		PhoneCode: "+54",
		Code:      "AR",
	},
	{
		Name:      "Armenia",
		PhoneCode: "+374",
		Code:      "AM",
	},
	{
		Name:      "Aruba",
		PhoneCode: "+297",
		Code:      "AW",
	},
	{
		Name:      "Australia",
		PhoneCode: "+61",
		Code:      "AU",
	},
	{
		Name:      "Austria",
		PhoneCode: "+43",
		Code:      "AT",
	},
	{
		Name:      "Azerbaijan",
		PhoneCode: "+994",
		Code:      "AZ",
	},
	{
		Name:      "Bahamas",
		PhoneCode: "+1 242",
		Code:      "BS",
	},
	{
		Name:      "Bahrain",
		PhoneCode: "+973",
		Code:      "BH",
	},
	{
		Name:      "Bangladesh",
		PhoneCode: "+880",
		Code:      "BD",
	},
	{
		Name:      "Barbados",
		PhoneCode: "+1 246",
		Code:      "BB",
	},
	{
		Name:      "Belarus",
		PhoneCode: "+375",
		Code:      "BY",
	},
	{
		Name:      "Belgium",
		PhoneCode: "+32",
		Code:      "BE",
	},
	{
		Name:      "Belize",
		PhoneCode: "+501",
		Code:      "BZ",
	},
	{
		Name:      "Benin",
		PhoneCode: "+229",
		Code:      "BJ",
	},
	{
		Name:      "Bermuda",
		PhoneCode: "+1 441",
		Code:      "BM",
	},
	{
		Name:      "Bhutan",
		PhoneCode: "+975",
		Code:      "BT",
	},
	{
		Name:      "Bolivia, Plurinational State of",
		PhoneCode: "+591",
		Code:      "BO",
	},
	{
		Name:      "Bosnia and Herzegovina",
		PhoneCode: "+387",
		Code:      "BA",
	},
	{
		Name:      "Botswana",
		PhoneCode: "+267",
		Code:      "BW",
	},
	{
		Name:      "Brazil",
		PhoneCode: "+55",
		Code:      "BR",
	},
	{
		Name:      "British Indian Ocean Territory",
		PhoneCode: "+246",
		Code:      "IO",
	},
	{
		Name:      "Brunei Darussalam",
		PhoneCode: "+673",
		Code:      "BN",
	},
	{
		Name:      "Bulgaria",
		PhoneCode: "+359",
		Code:      "BG",
	},
	{
		Name:      "Burkina Faso",
		PhoneCode: "+226",
		Code:      "BF",
	},
	{
		Name:      "Burundi",
		PhoneCode: "+257",
		Code:      "BI",
	},
	{
		Name:      "Cambodia",
		PhoneCode: "+855",
		Code:      "KH",
	},
	{
		Name:      "Cameroon",
		PhoneCode: "+237",
		Code:      "CM",
	},
	{
		Name:      "Canada",
		PhoneCode: "+1",
		Code:      "CA",
	},
	{
		Name:      "Cape Verde",
		PhoneCode: "+238",
		Code:      "CV",
	},
	{
		Name:      "Cayman Islands",
		PhoneCode: "+ 345",
		Code:      "KY",
	},
	{
		Name:      "Central African Republic",
		PhoneCode: "+236",
		Code:      "CF",
	},
	{
		Name:      "Chad",
		PhoneCode: "+235",
		Code:      "TD",
	},
	{
		Name:      "Chile",
		PhoneCode: "+56",
		Code:      "CL",
	},
	{
		Name:      "China",
		PhoneCode: "+86",
		Code:      "CN",
	},
	{
		Name:      "Christmas Island",
		PhoneCode: "+61",
		Code:      "CX",
	},
	{
		Name:      "Cocos (Keeling) Islands",
		PhoneCode: "+61",
		Code:      "CC",
	},
	{
		Name:      "Colombia",
		PhoneCode: "+57",
		Code:      "CO",
	},
	{
		Name:      "Comoros",
		PhoneCode: "+269",
		Code:      "KM",
	},
	{
		Name:      "Congo",
		PhoneCode: "+242",
		Code:      "CG",
	},
	{
		Name:      "Congo, The Democratic Republic of the Congo",
		PhoneCode: "+243",
		Code:      "CD",
	},
	{
		Name:      "Cook Islands",
		PhoneCode: "+682",
		Code:      "CK",
	},
	{
		Name:      "Costa Rica",
		PhoneCode: "+506",
		Code:      "CR",
	},
	{
		Name:      "Cote d'Ivoire",
		PhoneCode: "+225",
		Code:      "CI",
	},
	{
		Name:      "Croatia",
		PhoneCode: "+385",
		Code:      "HR",
	},
	{
		Name:      "Cuba",
		PhoneCode: "+53",
		Code:      "CU",
	},
	{
		Name:      "Cyprus",
		PhoneCode: "+357",
		Code:      "CY",
	},
	{
		Name:      "Czech Republic",
		PhoneCode: "+420",
		Code:      "CZ",
	},
	{
		Name:      "Denmark",
		PhoneCode: "+45",
		Code:      "DK",
	},
	{
		Name:      "Djibouti",
		PhoneCode: "+253",
		Code:      "DJ",
	},
	{
		Name:      "Dominica",
		PhoneCode: "+1 767",
		Code:      "DM",
	},
	{
		Name:      "Dominican Republic",
		PhoneCode: "+1 849",
		Code:      "DO",
	},
	{
		Name:      "Ecuador",
		PhoneCode: "+593",
		Code:      "EC",
	},
	{
		Name:      "Egypt",
		PhoneCode: "+20",
		Code:      "EG",
	},
	{
		Name:      "El Salvador",
		PhoneCode: "+503",
		Code:      "SV",
	},
	{
		Name:      "Equatorial Guinea",
		PhoneCode: "+240",
		Code:      "GQ",
	},
	{
		Name:      "Eritrea",
		PhoneCode: "+291",
		Code:      "ER",
	},
	{
		Name:      "Estonia",
		PhoneCode: "+372",
		Code:      "EE",
	},
	{
		Name:      "Ethiopia",
		PhoneCode: "+251",
		Code:      "ET",
	},
	{
		Name:      "Falkland Islands (Malvinas)",
		PhoneCode: "+500",
		Code:      "FK",
	},
	{
		Name:      "Faroe Islands",
		PhoneCode: "+298",
		Code:      "FO",
	},
	{
		Name:      "Fiji",
		PhoneCode: "+679",
		Code:      "FJ",
	},
	{
		Name:      "Finland",
		PhoneCode: "+358",
		Code:      "FI",
	},
	{
		Name:      "France",
		PhoneCode: "+33",
		Code:      "FR",
	},
	{
		Name:      "French Guiana",
		PhoneCode: "+594",
		Code:      "GF",
	},
	{
		Name:      "French Polynesia",
		PhoneCode: "+689",
		Code:      "PF",
	},
	{
		Name:      "Gabon",
		PhoneCode: "+241",
		Code:      "GA",
	},
	{
		Name:      "Gambia",
		PhoneCode: "+220",
		Code:      "GM",
	},
	{
		Name:      "Georgia",
		PhoneCode: "+995",
		Code:      "GE",
	},
	{
		Name:      "Germany",
		PhoneCode: "+49",
		Code:      "DE",
	},
	{
		Name:      "Ghana",
		PhoneCode: "+233",
		Code:      "GH",
	},
	{
		Name:      "Gibraltar",
		PhoneCode: "+350",
		Code:      "GI",
	},
	{
		Name:      "Greece",
		PhoneCode: "+30",
		Code:      "GR",
	},
	{
		Name:      "Greenland",
		PhoneCode: "+299",
		Code:      "GL",
	},
	{
		Name:      "Grenada",
		PhoneCode: "+1 473",
		Code:      "GD",
	},
	{
		Name:      "Guadeloupe",
		PhoneCode: "+590",
		Code:      "GP",
	},
	{
		Name:      "Guam",
		PhoneCode: "+1 671",
		Code:      "GU",
	},
	{
		Name:      "Guatemala",
		PhoneCode: "+502",
		Code:      "GT",
	},
	{
		Name:      "Guernsey",
		PhoneCode: "+44",
		Code:      "GG",
	},
	{
		Name:      "Guinea",
		PhoneCode: "+224",
		Code:      "GN",
	},
	{
		Name:      "Guinea-Bissau",
		PhoneCode: "+245",
		Code:      "GW",
	},
	{
		Name:      "Guyana",
		PhoneCode: "+595",
		Code:      "GY",
	},
	{
		Name:      "Haiti",
		PhoneCode: "+509",
		Code:      "HT",
	},
	{
		Name:      "Holy See (Vatican City State)",
		PhoneCode: "+379",
		Code:      "VA",
	},
	{
		Name:      "Honduras",
		PhoneCode: "+504",
		Code:      "HN",
	},
	{
		Name:      "Hong Kong",
		PhoneCode: "+852",
		Code:      "HK",
	},
	{
		Name:      "Hungary",
		PhoneCode: "+36",
		Code:      "HU",
	},
	{
		Name:      "Iceland",
		PhoneCode: "+354",
		Code:      "IS",
	},
	{
		Name:      "India",
		PhoneCode: "+91",
		Code:      "IN",
	},
	{
		Name:      "Indonesia",
		PhoneCode: "+62",
		Code:      "ID",
	},
	{
		Name:      "Iran, Islamic Republic of Persian Gulf",
		PhoneCode: "+98",
		Code:      "IR",
	},
	{
		Name:      "Iraq",
		PhoneCode: "+964",
		Code:      "IQ",
	},
	{
		Name:      "Ireland",
		PhoneCode: "+353",
		Code:      "IE",
	},
	{
		Name:      "Isle of Man",
		PhoneCode: "+44",
		Code:      "IM",
	},
	{
		Name:      "Israel",
		PhoneCode: "+972",
		Code:      "IL",
	},
	{
		Name:      "Italy",
		PhoneCode: "+39",
		Code:      "IT",
	},
	{
		Name:      "Jamaica",
		PhoneCode: "+1 876",
		Code:      "JM",
	},
	{
		Name:      "Japan",
		PhoneCode: "+81",
		Code:      "JP",
	},
	{
		Name:      "Jersey",
		PhoneCode: "+44",
		Code:      "JE",
	},
	{
		Name:      "Jordan",
		PhoneCode: "+962",
		Code:      "JO",
	},
	{
		Name:      "Kazakhstan",
		PhoneCode: "+7 7",
		Code:      "KZ",
	},
	{
		Name:      "Kenya",
		PhoneCode: "+254",
		Code:      "KE",
	},
	{
		Name:      "Kiribati",
		PhoneCode: "+686",
		Code:      "KI",
	},
	{
		Name:      "Korea, Democratic People's Republic of Korea",
		PhoneCode: "+850",
		Code:      "KP",
	},
	{
		Name:      "Korea, Republic of South Korea",
		PhoneCode: "+82",
		Code:      "KR",
	},
	{
		Name:      "Kosovo",
		PhoneCode: "+383",
		Code:      "XK",
	},
	{
		Name:      "Kuwait",
		PhoneCode: "+965",
		Code:      "KW",
	},
	{
		Name:      "Kyrgyzstan",
		PhoneCode: "+996",
		Code:      "KG",
	},
	{
		Name:      "Laos",
		PhoneCode: "+856",
		Code:      "LA",
	},
	{
		Name:      "Latvia",
		PhoneCode: "+371",
		Code:      "LV",
	},
	{
		Name:      "Lebanon",
		PhoneCode: "+961",
		Code:      "LB",
	},
	{
		Name:      "Lesotho",
		PhoneCode: "+266",
		Code:      "LS",
	},
	{
		Name:      "Liberia",
		PhoneCode: "+231",
		Code:      "LR",
	},
	{
		Name:      "Libyan Arab Jamahiriya",
		PhoneCode: "+218",
		Code:      "LY",
	},
	{
		Name:      "Liechtenstein",
		PhoneCode: "+423",
		Code:      "LI",
	},
	{
		Name:      "Lithuania",
		PhoneCode: "+370",
		Code:      "LT",
	},
	{
		Name:      "Luxembourg",
		PhoneCode: "+352",
		Code:      "LU",
	},
	{
		Name:      "Macao",
		PhoneCode: "+853",
		Code:      "MO",
	},
	{
		Name:      "Macedonia",
		PhoneCode: "+389",
		Code:      "MK",
	},
	{
		Name:      "Madagascar",
		PhoneCode: "+261",
		Code:      "MG",
	},
	{
		Name:      "Malawi",
		PhoneCode: "+265",
		Code:      "MW",
	},
	{
		Name:      "Malaysia",
		PhoneCode: "+60",
		Code:      "MY",
	},
	{
		Name:      "Maldives",
		PhoneCode: "+960",
		Code:      "MV",
	},
	{
		Name:      "Mali",
		PhoneCode: "+223",
		Code:      "ML",
	},
	{
		Name:      "Malta",
		PhoneCode: "+356",
		Code:      "MT",
	},
	{
		Name:      "Marshall Islands",
		PhoneCode: "+692",
		Code:      "MH",
	},
	{
		Name:      "Martinique",
		PhoneCode: "+596",
		Code:      "MQ",
	},
	{
		Name:      "Mauritania",
		PhoneCode: "+222",
		Code:      "MR",
	},
	{
		Name:      "Mauritius",
		PhoneCode: "+230",
		Code:      "MU",
	},
	{
		Name:      "Mayotte",
		PhoneCode: "+262",
		Code:      "YT",
	},
	{
		Name:      "Mexico",
		PhoneCode: "+52",
		Code:      "MX",
	},
	{
		Name:      "Micronesia, Federated States of Micronesia",
		PhoneCode: "+691",
		Code:      "FM",
	},
	{
		Name:      "Moldova",
		PhoneCode: "+373",
		Code:      "MD",
	},
	{
		Name:      "Monaco",
		PhoneCode: "+377",
		Code:      "MC",
	},
	{
		Name:      "Mongolia",
		PhoneCode: "+976",
		Code:      "MN",
	},
	{
		Name:      "Montenegro",
		PhoneCode: "+382",
		Code:      "ME",
	},
	{
		Name:      "Montserrat",
		PhoneCode: "+1664",
		Code:      "MS",
	},
	{
		Name:      "Morocco",
		PhoneCode: "+212",
		Code:      "MA",
	},
	{
		Name:      "Mozambique",
		PhoneCode: "+258",
		Code:      "MZ",
	},
	{
		Name:      "Myanmar",
		PhoneCode: "+95",
		Code:      "MM",
	},
	{
		Name:      "Namibia",
		PhoneCode: "+264",
		Code:      "NA",
	},
	{
		Name:      "Nauru",
		PhoneCode: "+674",
		Code:      "NR",
	},
	{
		Name:      "Nepal",
		PhoneCode: "+977",
		Code:      "NP",
	},
	{
		Name:      "Netherlands",
		PhoneCode: "+31",
		Code:      "NL",
	},
	{
		Name:      "Netherlands Antilles",
		PhoneCode: "+599",
		Code:      "AN",
	},
	{
		Name:      "New Caledonia",
		PhoneCode: "+687",
		Code:      "NC",
	},
	{
		Name:      "New Zealand",
		PhoneCode: "+64",
		Code:      "NZ",
	},
	{
		Name:      "Nicaragua",
		PhoneCode: "+505",
		Code:      "NI",
	},
	{
		Name:      "Niger",
		PhoneCode: "+227",
		Code:      "NE",
	},
	{
		Name:      "Nigeria",
		PhoneCode: "+234",
		Code:      "NG",
	},
	{
		Name:      "Niue",
		PhoneCode: "+683",
		Code:      "NU",
	},
	{
		Name:      "Norfolk Island",
		PhoneCode: "+672",
		Code:      "NF",
	},
	{
		Name:      "Northern Mariana Islands",
		PhoneCode: "+1 670",
		Code:      "MP",
	},
	{
		Name:      "Norway",
		PhoneCode: "+47",
		Code:      "NO",
	},
	{
		Name:      "Oman",
		PhoneCode: "+968",
		Code:      "OM",
	},
	{
		Name:      "Pakistan",
		PhoneCode: "+92",
		Code:      "PK",
	},
	{
		Name:      "Palau",
		PhoneCode: "+680",
		Code:      "PW",
	},
	{
		Name:      "Palestinian Territory, Occupied",
		PhoneCode: "+970",
		Code:      "PS",
	},
	{
		Name:      "Panama",
		PhoneCode: "+507",
		Code:      "PA",
	},
	{
		Name:      "Papua New Guinea",
		PhoneCode: "+675",
		Code:      "PG",
	},
	{
		Name:      "Paraguay",
		PhoneCode: "+595",
		Code:      "PY",
	},
	{
		Name:      "Peru",
		PhoneCode: "+51",
		Code:      "PE",
	},
	{
		Name:      "Philippines",
		PhoneCode: "+63",
		Code:      "PH",
	},
	{
		Name:      "Pitcairn",
		PhoneCode: "+872",
		Code:      "PN",
	},
	{
		Name:      "Poland",
		PhoneCode: "+48",
		Code:      "PL",
	},
	{
		Name:      "Portugal",
		PhoneCode: "+351",
		Code:      "PT",
	},
	{
		Name:      "Puerto Rico",
		PhoneCode: "+1 939",
		Code:      "PR",
	},
	{
		Name:      "Qatar",
		PhoneCode: "+974",
		Code:      "QA",
	},
	{
		Name:      "Romania",
		PhoneCode: "+40",
		Code:      "RO",
	},
	{
		Name:      "Russia",
		PhoneCode: "+7",
		Code:      "RU",
	},
	{
		Name:      "Rwanda",
		PhoneCode: "+250",
		Code:      "RW",
	},
	{
		Name:      "Reunion",
		PhoneCode: "+262",
		Code:      "RE",
	},
	{
		Name:      "Saint Barthelemy",
		PhoneCode: "+590",
		Code:      "BL",
	},
	{
		Name:      "Saint Helena, Ascension and Tristan Da Cunha",
		PhoneCode: "+290",
		Code:      "SH",
	},
	{
		Name:      "Saint Kitts and Nevis",
		PhoneCode: "+1 869",
		Code:      "KN",
	},
	{
		Name:      "Saint Lucia",
		PhoneCode: "+1 758",
		Code:      "LC",
	},
	{
		Name:      "Saint Martin",
		PhoneCode: "+590",
		Code:      "MF",
	},
	{
		Name:      "Saint Pierre and Miquelon",
		PhoneCode: "+508",
		Code:      "PM",
	},
	{
		Name:      "Saint Vincent and the Grenadines",
		PhoneCode: "+1 784",
		Code:      "VC",
	},
	{
		Name:      "Samoa",
		PhoneCode: "+685",
		Code:      "WS",
	},
	{
		Name:      "San Marino",
		PhoneCode: "+378",
		Code:      "SM",
	},
	{
		Name:      "Sao Tome and Principe",
		PhoneCode: "+239",
		Code:      "ST",
	},
	{
		Name:      "Saudi Arabia",
		PhoneCode: "+966",
		Code:      "SA",
	},
	{
		Name:      "Senegal",
		PhoneCode: "+221",
		Code:      "SN",
	},
	{
		Name:      "Serbia",
		PhoneCode: "+381",
		Code:      "RS",
	},
	{
		Name:      "Seychelles",
		PhoneCode: "+248",
		Code:      "SC",
	},
	{
		Name:      "Sierra Leone",
		PhoneCode: "+232",
		Code:      "SL",
	},
	{
		Name:      "Singapore",
		PhoneCode: "+65",
		Code:      "SG",
	},
	{
		Name:      "Slovakia",
		PhoneCode: "+421",
		Code:      "SK",
	},
	{
		Name:      "Slovenia",
		PhoneCode: "+386",
		Code:      "SI",
	},
	{
		Name:      "Solomon Islands",
		PhoneCode: "+677",
		Code:      "SB",
	},
	{
		Name:      "Somalia",
		PhoneCode: "+252",
		Code:      "SO",
	},
	{
		Name:      "South Africa",
		PhoneCode: "+27",
		Code:      "ZA",
	},
	{
		Name:      "South Georgia and the South Sandwich Islands",
		PhoneCode: "+500",
		Code:      "GS",
	},
	{
		Name:      "Spain",
		PhoneCode: "+34",
		Code:      "ES",
	},
	{
		Name:      "Sri Lanka",
		PhoneCode: "+94",
		Code:      "LK",
	},
	{
		Name:      "Sudan",
		PhoneCode: "+249",
		Code:      "SD",
	},
	{
		Name:      "Suriname",
		PhoneCode: "+597",
		Code:      "SR",
	},
	{
		Name:      "Svalbard and Jan Mayen",
		PhoneCode: "+47",
		Code:      "SJ",
	},
	{
		Name:      "Swaziland",
		PhoneCode: "+268",
		Code:      "SZ",
	},
	{
		Name:      "Sweden",
		PhoneCode: "+46",
		Code:      "SE",
	},
	{
		Name:      "Switzerland",
		PhoneCode: "+41",
		Code:      "CH",
	},
	{
		Name:      "Syrian Arab Republic",
		PhoneCode: "+963",
		Code:      "SY",
	},
	{
		Name:      "Taiwan",
		PhoneCode: "+886",
		Code:      "TW",
	},
	{
		Name:      "Tajikistan",
		PhoneCode: "+992",
		Code:      "TJ",
	},
	{
		Name:      "Tanzania, United Republic of Tanzania",
		PhoneCode: "+255",
		Code:      "TZ",
	},
	{
		Name:      "Thailand",
		PhoneCode: "+66",
		Code:      "TH",
	},
	{
		Name:      "Timor-Leste",
		PhoneCode: "+670",
		Code:      "TL",
	},
	{
		Name:      "Togo",
		PhoneCode: "+228",
		Code:      "TG",
	},
	{
		Name:      "Tokelau",
		PhoneCode: "+690",
		Code:      "TK",
	},
	{
		Name:      "Tonga",
		PhoneCode: "+676",
		Code:      "TO",
	},
	{
		Name:      "Trinidad and Tobago",
		PhoneCode: "+1 868",
		Code:      "TT",
	},
	{
		Name:      "Tunisia",
		PhoneCode: "+216",
		Code:      "TN",
	},
	{
		Name:      "Turkey",
		PhoneCode: "+90",
		Code:      "TR",
	},
	{
		Name:      "Turkmenistan",
		PhoneCode: "+993",
		Code:      "TM",
	},
	{
		Name:      "Turks and Caicos Islands",
		PhoneCode: "+1 649",
		Code:      "TC",
	},
	{
		Name:      "Tuvalu",
		PhoneCode: "+688",
		Code:      "TV",
	},
	{
		Name:      "Uganda",
		PhoneCode: "+256",
		Code:      "UG",
	},
	{
		Name:      "Ukraine",
		PhoneCode: "+380",
		Code:      "UA",
	},
	{
		Name:      "United Arab Emirates",
		PhoneCode: "+971",
		Code:      "AE",
	},
	{
		Name:      "United Kingdom",
		PhoneCode: "+44",
		Code:      "GB",
	},
	{
		Name:      "United States",
		PhoneCode: "+1",
		Code:      "US",
	},
	{
		Name:      "Uruguay",
		PhoneCode: "+598",
		Code:      "UY",
	},
	{
		Name:      "Uzbekistan",
		PhoneCode: "+998",
		Code:      "UZ",
	},
	{
		Name:      "Vanuatu",
		PhoneCode: "+678",
		Code:      "VU",
	},
	{
		Name:      "Venezuela, Bolivarian Republic of Venezuela",
		PhoneCode: "+58",
		Code:      "VE",
	},
	{
		Name:      "Vietnam",
		PhoneCode: "+84",
		Code:      "VN",
	},
	{
		Name:      "Virgin Islands, British",
		PhoneCode: "+1 284",
		Code:      "VG",
	},
	{
		Name:      "Virgin Islands, U.S.",
		PhoneCode: "+1 340",
		Code:      "VI",
	},
	{
		Name:      "Wallis and Futuna",
		PhoneCode: "+681",
		Code:      "WF",
	},
	{
		Name:      "Yemen",
		PhoneCode: "+967",
		Code:      "YE",
	},
	{
		Name:      "Zambia",
		PhoneCode: "+260",
		Code:      "ZM",
	},
	{
		Name:      "Zimbabwe",
		PhoneCode: "+263",
		Code:      "ZW",
	},
}
