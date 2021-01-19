package openweathermap

// see https://openweathermap.org/weather-conditions
var conditionsMap = map[int]string{

	200: "Isolated Thunderstorms",
	201: "Scattered Thunderstorms",
	202: "Thunderstorms",
	210: "Thunderstorms",
	211: "Thunderstorms",
	212: "Thunderstorms",
	221: "Thunderstorms",
	230: "Scattered Thunderstorms",
	231: "Scattered Thunderstorms",
	232: "Thunderstorms",

	300: "Drizzle",
	301: "Drizzle",
	302: "Drizzle",
	310: "Drizzle",
	311: "Drizzle",
	312: "Drizzle",
	313: "Drizzle",
	314: "Drizzle",
	321: "Drizzle",

	500: "Light Rain",
	501: "Light Rain",
	502: "Heavy Rain",
	503: "Heavy Rain",
	504: "Heavy Rain",
	511: "Freezing Rain",
	520: "Rain",
	521: "Scattered Showers",
	522: "Heavy Rain",
	531: "Heavy Rain",

	600: "Light Snow",
	601: "Snow",
	602: "Heavy Snow",
	611: "Sleet",
	612: "Wintry Mix Snow / Sleet",
	613: "Scattered Snow Showers",
	615: "Scattered Snow Showers",
	616: "Scattered Snow Showers",
	620: "Wintry Mix Snow / Sleet",
	621: "Scattered Snow Showers",
	622: "Scattered Snow Showers",

	701: "Haze / Windy",
	711: "Smoke / Windy",
	721: "Haze / Windy",
	731: "Blowing Dust / Sandstorm",
	741: "Foggy",
	751: "Blowing Dust / Sandstorm",
	761: "Blowing Dust / Sandstorm",
	762: "Blowing Dust / Sandstorm",
	771: "Blowing Spray / Windy",
	781: "Tornado",

	800: "Clear",
	801: "Fair / Mostly Clear",
	802: "Partly Cloudy",
	803: "Mostly Cloudy",
	804: "Cloudy",
}
