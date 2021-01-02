package anagram

func FoundAnagramsInWord(anagramToFound string, completeWord string) ([]string) {
	lenCompleteWord := len(completeWord)
	lenAnagramToFound := len(anagramToFound)

	if isStringNotEmpty(anagramToFound) {
		var anagramFound []string

		for currentStart := 0; isEndReachedInCompleteWord(currentStart, lenCompleteWord, lenAnagramToFound); currentStart += 1 {
			var substrWord = getSubstrWord(completeWord, currentStart, len(anagramToFound))
			anagramFound = append(anagramFound, getSubstrToAnagramOrEmpty(substrWord, anagramToFound))
		}

		return anagramFound
	} else {
		return []string{}
	}
}

func getSubstrToAnagramOrEmpty(splitWord string, anagramToFound string) (string) {
	var foundLettersMap = map[int]string{}
	var maxForLoop = len(anagramToFound)
	
	for currentSplitIdx := 0; currentSplitIdx < maxForLoop; currentSplitIdx += 1 {
		var currentLetterInSplit = string(splitWord[currentSplitIdx])

		for currentAnagramIdx := 0; currentAnagramIdx < maxForLoop; currentAnagramIdx += 1 {
			var currentLetterInAnagram = string(anagramToFound[currentAnagramIdx])

			if isSearchedLetterEqualToSplit(currentLetterInAnagram, currentLetterInSplit) &&
				isNotCurrentLetterInMap(foundLettersMap, currentAnagramIdx) {

				foundLettersMap[currentAnagramIdx] = currentLetterInAnagram
				break
			}
		}
	}
	
	return getWordIfValidOrEmpty(foundLettersMap, splitWord)
}

func getWordIfValidOrEmpty(foundLettersMap map[int]string, word string) (string){
	if len(foundLettersMap) == len(word) {
		return word
	}
	
	return ""
}

func isSearchedLetterEqualToSplit(currentLetterInSearched, currentLetterInSplit string) bool {
	return currentLetterInSearched == currentLetterInSplit
}

func isNotCurrentLetterInMap(foundLettersMap map[int]string, currentIdx int) bool {
	return foundLettersMap[currentIdx] == ""
}

func isEndReachedInCompleteWord(currentStart, lenCompleteWord, lenAnagramToFound int) (bool) {
	return currentStart <= (lenCompleteWord - lenAnagramToFound)
}

func getSubstrWord(completeWord string, currentStart int, end int) string {
	return completeWord[currentStart:(currentStart + end)]
}

func isStringNotEmpty(valueStr string) (bool) {
	return len(valueStr) > 0
}
