package archive

import "fmt"



//def divide_chunks(l, n):
//
//# looping till length l
//for i in range(0, len(l), n):
//yield l[i:i + n]
//
//# How many elements each
//# list should have
//n = 5

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

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

			fmt.Println(i, chunkLength, arrLength)
			chunks = append(chunks, arr[i:i + chunkLength])
		}
	}

	return

}

type Blah struct {
	hello string
}

func (b *Blah) DoSomeStuff() {
	var someGlobal string
	blue := b.hello
	fmt.Println(someGlobal, blue)

}

func main () {

	//caw := []string{"hello", "bello", "nello", "hello", "bello", "nello", "bello", "nello", "hello", "bello", "nello", "bello", "nello", "hello", "bello", "bello", "nello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "bill cosby"}
	////fmt.Println(len(caw))
	//fmt.Println(GetChunksFromStringArray(caw, 4))


}
