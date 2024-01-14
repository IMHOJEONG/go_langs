// https://go.dev/tour/moretypes/18

package moretypes_18

// import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {

			if i%3 == 0 {
				result[i][j] = uint8(i * j)
				continue
			}

			if i%3 == 1 {
				result[i][j] = uint8(i ^ j)
				continue
			}

			result[i][j] = uint8((i + j) / 2)
		}
	}

	return result
}

func main() {
	// pic.Show(Pic)
}
