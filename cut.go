/*
 *Author:wangshuyao
 *Date:17:03 2019/5/15
 */

package wordninja

import (
	"math"
	"bufio"
	"regexp"
	"strings"
	"errors"
	"unicode"
	"fmt"
	"os"
	"io"
)

var maxLenWord int
var wordCost map[string]float64

func init() {
	words := loadWords("./dict/wordninja_words.txt")
	generateCutWordMap(words)
	fmt.Println("init English cut words successfully!")
}

// load the english cut words from file, return list of word.
func loadWords(file string) ([]string) {
	words, err := readFileByLine(file)
	if err != nil {
		panic("load english cut word failed," + err.Error())
	}
	return words
}

// generateCutWordMap initialize wordCost map and the length of the longest word.
// the cost of a word is calculated by `log(log(len(words))*idx)`.
func generateCutWordMap(words []string) {
	wordCost = make(map[string]float64)
	var wordLen int
	logLen := math.Log(float64(len(words)))
	for idx, word := range words {
		wordLen = len(word)
		if wordLen > maxLenWord {
			maxLenWord = wordLen
		}
		wordCost[word] = math.Log(logLen * float64(idx+1))
	}
}

type match struct {
	cost float64
	idx  int
}

type text struct {
	s string
}

// bestMatch will return the minimal cost and its appropriate character's index.
func (s *text) bestMatch(costs []float64, i int) (match, error) {
	candidates := costs[max(0, i-maxLenWord): i]
	k := 0
	var matchs []match
	for j := len(candidates) - 1; j >= 0; j-- {
		cost := getWordCost(strings.ToLower(s.s[i-k-1:i])) + float64(candidates[j])
		matchs = append(matchs, match{cost: cost, idx: k + 1})
		k++
	}
	return minCost(matchs)
}

// getWordCost return cost of word from the wordCost map.
// if the word is not exist in the map, it will return `9e99`.
func getWordCost(word string) float64 {
	if v, ok := wordCost[word]; ok {
		return v
	}
	return 9e99
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min return the minimal cost of the matchs.
func minCost(matchs []match) (match, error) {
	if len(matchs) == 0 {
		return match{}, errors.New("match.len ")
	}
	r := matchs[0]
	for _, m := range matchs {
		if m.cost < r.cost {
			r = m
		}
	}
	return r, nil
}

func Cut(s string) []string {
	eng := getEnglishText(s)
	return CutEnglish(eng)
}

// CutEnglish return the best matched words with cutting the English string `s`.
func CutEnglish(eng string) []string {
	text := text{s: eng}
	costs := []float64{0}
	for i := 1; i < len(eng)+1; i++ {
		if m, err := text.bestMatch(costs, i); err == nil {
			costs = append(costs, m.cost)
		}
	}

	var out []string
	i := len(eng)

	for i > 0 {
		m, err := text.bestMatch(costs, i)
		if err != nil {
			continue
		}

		newToken := true

		//ignore a lone apostrophe
		if !(eng[i-m.idx:i] == "'") {
			if len(out) > 0 {
				//re-attach split 's and split digits or digit followed by digit.
				if out[len(out)-1] == "'s" ||
					(unicode.IsDigit(rune(eng[i-1])) && unicode.IsDigit(rune(out[len(out)-1][0]))) {
					// combine current token with previous token.
					out[len(out)-1] = eng[i-m.idx:i] + out[len(out)-1]
					newToken = false
				}
			}
		}
		if newToken {
			word := eng[i-m.idx:i]
			out = append(out, word)
		}
		i -= m.idx
	}
	return reverse(out)
}

// reverse return reversed list of `dst`
func reverse(dst []string) (res []string) {
	for i := len(dst); i > 0; i-- {
		res = append(res, dst[i-1])
	}
	return
}

// getEnglishText return all the English characters of string `s`.
func getEnglishText(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9']+")
	return strings.Join(reg.Split(s, -1), "")
}

//ReadFileByLine return a list by reading file line by line.
func readFileByLine(file string) (lines []string, err error) {

	f, err := os.Open(file)
	if err != nil {
		return lines, errors.New("open file failed")
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if line == "\n" {
			continue
		}
		line = strings.Replace(line, "\n", "", -1)
		lines = append(lines, line)
	}
	return lines, nil
}
