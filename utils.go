package steamapi

import "strconv"

func steamIDs2SplitArray(steamIDs []uint64, maxLen int) [][]string {
	strIds := make([][]string, 0)
	var curArr []string
	for i, id := range steamIDs {
		if i%maxLen == 0 {
			strIds = append(strIds, curArr)
			curArr = []string{}
		}

		curArr = append(curArr, strconv.FormatUint(id, 10))
	}

	if len(curArr) > 0 {
		strIds = append(strIds, curArr)
	}

	return strIds
}
