//Package dog2 takes human years and returns corresponding dog years. Ensure that move this package is located under GOPATH.
package dog2

//Dogyears function takes human years and returns dog years.
func Dogyears(humanyears int) int {
	return (humanyears * 7)
}

//YearsTwo takes human years and returns dog years but instead of single instruction of multiplying human years by 7, performs a repeated addition.
func YearsTwo(humanyears int) int {
	var count int

	for i := 0; i < humanyears; i++ {
		count += 7
	}

	return count
}
