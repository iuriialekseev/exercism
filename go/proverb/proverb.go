// Package proverb implements function that generates medieval rhyme
package proverb

// Proverb returns medieval rhyme
func Proverb(rhyme []string) []string {
	var result []string

	for i, word := range rhyme {
		if i < len(rhyme)-1 {
			result = append(result, "For want of a "+word+" the "+rhyme[i+1]+" was lost.")
		} else {
			result = append(result, "And all for the want of a "+rhyme[0]+".")
		}
	}

	return result
}
