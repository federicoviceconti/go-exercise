package test

import (
	"testing"
	"go-exercise-book/anagrams/anagram"
)

func TestAnagramsWorksWithTheSameWord(t *testing.T) {
	var completeWord = "aaba"
	var anagramToFound = "aaba"

	if len(anagram.FoundAnagramsInWord(completeWord, anagramToFound)) == 0 {
		t.Fatalf("Function doesn't work with same word")
	}
}

func TestAnagramsWorksWithAnotherWord(t *testing.T) {
	var completeWord = "abbabbabba"
	var anagramToFound = "abba"

	if len(anagram.FoundAnagramsInWord(completeWord, anagramToFound)) > 0 {
		t.Fatalf("Function doesn't work with another word")
	}
}

func TestAnagramsIfContainsAnagram(t *testing.T) {
	var completeWord = "abdba"
	var anagramToFound = "adbba"

	for _, value := range anagram.FoundAnagramsInWord(anagramToFound, completeWord) {
		if value == completeWord {
			return
		}
	}

	t.Fatal("The word is not contained!")
}