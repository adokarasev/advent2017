package passphrase

import (
	"bufio"
	"sort"
	"strings"
)

//Passphrase if passphrase does not contain duplicate words
func Passphrase(passphrase string) bool {
	tokens := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(passphrase))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		t := scanner.Text()
		if _, ok := tokens[t]; ok {
			return false
		}
		tokens[t] = true
	}
	return true
}

// StrongPasshrase validates passphrase does not contain two words
// that are anagrams of each other - that is, a passphrase is invalid
// if any word's letters can be rearranged to form any other word in the passphrase.
func StrongPasshrase(passphrase string) bool {
	tokens := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(passphrase))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		r := []rune(scanner.Text())
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		t := string(r)
		if _, ok := tokens[t]; ok {
			return false
		}
		tokens[t] = true
	}
	return true
}
