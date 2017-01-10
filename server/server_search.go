package server

//see github.com/jpillora/scraper for config specification
//cloud-torrent uses "<id>-item" handlers
var defaultSearchConfig = []byte(`{
	"kat": {
		"name": "Kickass Torrents",
		"url": "https://kat.how/search.php?q={{query}}/{{page:1}}/?field=seeders&sorder=desc",
		"list": "#mainSearchTable table tr[id]",
		"result": {
			"name":".cellMainLink",
			"path":[".cellMainLink", "@href"],
			"magnet": ["a[title=Torrent\\ magnet\\ link]", "@href"],
			"size": "td.nobr.center",
			"seeds": ".green.center",
			"peers": ".red.center"
		}
	},
	"tpb": {
		"name": "The Pirate Bay",
		"url": "https://thepiratebay.org/search/{{query}}/{{page:0}}/7//",
		"list": "#searchResult > tbody > tr",
		"result": {
			"name":"a.detLink",
			"path":["a.detLink","@href"],
			"magnet": ["a[title=Download\\ this\\ torrent\\ using\\ magnet]","@href"],
			"size": "/Size (\\d+(\\.\\d+).[KMG]iB)/",
			"seeds": "td:nth-child(3)",
			"peers": "td:nth-child(4)"
		}
	},
        "et": {
		"name": "ExtraTorrent",
		"url": "https://extratorrent.cc/search/?search={{query}}&s_cat=&pp=&srt=seeds&order=desc&page={{page:1}}",
		"list": "table.tl tr.tlr",
		"result": {
			"name":["td.tli > a"],
			"torrent": ["td:nth-child(1) a","@href","s/torrent_download/download/"],
			"url": ["td.tli > a","@href"],
			"size": "td:nth-child(5)",
			"seeds": "td.sy",
			"peers": "td.ly"
		}
	}
}`)
