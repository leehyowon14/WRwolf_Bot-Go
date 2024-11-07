package util

func GetUserNameFromMention(mention string) string {
	return mention[2 : len(mention)-1]
}
