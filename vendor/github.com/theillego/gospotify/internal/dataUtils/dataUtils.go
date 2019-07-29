package dataUtils

func GetChunksFromStringArray(arr []string, chunkLength int) (chunks [][]string)  {

	chunks = [][]string{} // slice
	arrLength := len(arr)

	if chunkLength != 0 {
		for i := 0; i < arrLength; i += chunkLength {
			if i + chunkLength > arrLength && chunkLength < arrLength {
				chunkLength = arrLength - i
			} else if chunkLength > arrLength {
				chunkLength = arrLength
			}

			chunks = append(chunks, arr[i:i + chunkLength])
		}
	}

	return

}

