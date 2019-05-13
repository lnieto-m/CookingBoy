package sakamotocommands

import "testing"

func TestLinkRegex(t *testing.T) {
	testTables := []struct {
		regx        string
		returnValue int
	}{
		{"", NONVALIDLINK},
		{"fsfdsfsdfsdf4345346%&%$@@$435@#%#$sd", NONVALIDLINK},
		{"youtube.com/watch?v=5454rew", NONVALIDLINK},
		{"http://youtube.com/", NONVALIDLINK},
		{"https://youtube.com/", NONVALIDLINK},
		{"https://youtube.en/watch?v=gtr_d4", VIDEO},
		{"http://www.youtube.de/watch?v=gtr_d4", VIDEO},
		{"https://youtube.fr/watch?v=tr6-ytr&list=gtr_d4", PLAYLIST},
		{"https://www.youtube.jp/watch?list=gtr_d4", NONVALIDLINK},
		{"https://youtube.en/watch?v=#@$@#%^^$_d4", VIDEO},
	}
	for _, table := range testTables {
		validity := checkLinkValidity(table.regx)
		if validity != table.returnValue {
			t.Errorf("checkLinkValidity failed for %s, got: %d, want: %d.\n", table.regx, validity, table.returnValue)
		}
	}
}
