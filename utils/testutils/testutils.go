package testutils

import (
	"github.com/adetunjii/netflakes/model"
)

var TestMovieSample = model.Movie{
	Title:        "A New Hope",
	EpisodeID:    4,
	OpeningCrawl: "It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....",
	Characters: []string{
		"https://swapi.dev/api/people/1/",
		"https://swapi.dev/api/people/2/",
		"https://swapi.dev/api/people/3/",
		"https://swapi.dev/api/people/4/",
		"https://swapi.dev/api/people/5/",
		"https://swapi.dev/api/people/6/",
		"https://swapi.dev/api/people/7/",
		"https://swapi.dev/api/people/8/",
		"https://swapi.dev/api/people/9/",
		"https://swapi.dev/api/people/10/",
		"https://swapi.dev/api/people/12/",
		"https://swapi.dev/api/people/13/",
		"https://swapi.dev/api/people/14/",
		"https://swapi.dev/api/people/15/",
		"https://swapi.dev/api/people/16/",
		"https://swapi.dev/api/people/18/",
		"https://swapi.dev/api/people/19/",
		"https://swapi.dev/api/people/81/",
	},
	Created: "2014-12-10T14:23:31.880000Z",
}

var TestMovieArray = []model.Movie{
	TestMovieSample,
}
