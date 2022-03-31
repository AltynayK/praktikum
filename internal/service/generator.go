package service

var (
	IdList map[int]string
	id     int = 0
)

func WriteURLByID(url string) int {
	id++
	IdList[id] = url

	return id
}

func GetURLFromID(id int) string {

	// remove row below

	return IdList[id]
}
