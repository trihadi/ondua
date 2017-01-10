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
	"abb": {
		"name": "1337X",
		"url": "http://1337x.to/search/{{query}}/seeders/desc/{{page:1}}/",
		"list": ".box-info-detail table.table tr",
		"result": {
			"name":[".coll-1 a:nth-child(2)"],
			"url":[".coll-1 a:nth-child(2)", "@href"],
			"seeds": ".coll-2",
			"peers": ".coll-3",
			"size": [".coll-4", "/([\\d\\.]+ [KMGT]?B)/"]
		}
	},
	"abb-item": {
		"name": "1337X (Item)",
		"url": "http://1337x.to{{item}}",
		"result": {
			"magnet": ["a.btn-magnet","@href"]
		}
	}
}`)
