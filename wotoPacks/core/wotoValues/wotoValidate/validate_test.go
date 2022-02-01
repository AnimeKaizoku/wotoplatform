package wotoValidate_test

import (
	"testing"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"
)

func TestPurifyKey(t *testing.T) {
	var (
		allKeys = []string{
			`   *   *   MANGA/\/\/\/\/\/\/\/\/~+-(*&*^&%`,
			`   * Anime - ~~~+-*&*^&%_`,
			"_   *  Light  *   Novel~    _____",
			"   *  Light  *   Novel~    ",
			"   *  Light  *   Novel~",
			"Light Novel",
			"Light_Novel",
			"Light-Novel",
			"Light_Novel",
			"Light-Novel",
			"Light-Novel~",
			"Light~Novel~",
			"Light     Novel~",
		}
	)
	p := ""
	for _, current := range allKeys {
		p = wotoValidate.PurifyKey(current)
		if p == "" {
			t.Errorf("PurifyKey(%s) = %s", current, p)
			return
		}
	}
}
