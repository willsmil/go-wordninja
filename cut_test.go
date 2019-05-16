/*
 *Author:wangshuyao
 *Date:10:24 2019/5/16
 */

package wordninja

import (
	"testing"
	"fmt"
)

func Test_simple(t *testing.T) {
	text := "derekanderson"
	fmt.Println(Cut(text)) // [derek anderson]
}

func Test_with_underscores(t *testing.T) {
	text := "derek_anderson"
	fmt.Println(Cut(text)) // [derek anderson]
}

func Test_caps(t *testing.T) {
	text := "DEREKANDERSON"
	fmt.Println(Cut(text)) // [DEREK ANDERSON]
}

func Test_digits(t *testing.T) {
	text := "win32intel"
	fmt.Println(Cut(text)) // [win 32 intel]
}

func Test_apostrophes(t *testing.T) {
	text := `"that'sthesheriff'sbadge"`
	fmt.Println(Cut(text)) // [that's the sheriff's badge]
}

func Test_with_chinese(t *testing.T) {
	text := "you你aresuch真awonderful好man"
	fmt.Println(Cut(text)) // [you are such a wonderful man]
}
