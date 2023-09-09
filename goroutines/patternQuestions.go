package goroutines

func PrintPattern1(n int) {
	var k = 0
	for i := 1; i <= n; i++ {
		if i <= n/2+1 {
			k++
		} else {
			k--
		}
		for j := 0; j < k; j++ {
			print("*")
		}
		println()
	}
}

func PrintPattern1usingRecursion(n, x, y int) {
	if x > 2*n-1 {
		return
	}

	if x > n {
		if y > 2*n-x {
			println()
			PrintPattern1usingRecursion(n, x+1, 1)
		} else {
			print("*")
			PrintPattern1usingRecursion(n, x, y+1)
		}
	} else if y > x {
		println()
		PrintPattern1usingRecursion(n, x+1, 1)
	} else {
		print("*")
		PrintPattern1usingRecursion(n, x, y+1)
	}
}

func PrintPatternPyramidUsingRecursion(n, x, z int, isStar bool) {
	if z == n {
		return
	}

	if x == 2*n-1 {
		// you change the line using this
		println()
		PrintPatternPyramidUsingRecursion(n, 0, z+1, true)
	} else {
		// you don't only print stars
		if x >= n-z-1 && x <= n+z-1 {
			if isStar {
				print("*")
				isStar = false
			} else {
				print(" ")
				isStar = true
			}
		} else {
			print(" ")
			isStar = true
		}
		PrintPatternPyramidUsingRecursion(n, x+1, z, isStar)
	}
}

func PyramidPattern(n int) {
	l := n
	m := n
	isStar := false
	for i := 1; i <= n; i++ {
		isStar = true
		for j := 1; j <= n*2-1; j++ {
			if j >= l && j <= m {
				if isStar {
					print("*")
					isStar = false
				} else {
					print(" ")
					isStar = true
				}
			} else {
				print(" ")
			}
		}
		l--
		m++
		println()
	}
}

func DiamondPattern(n int) {
	l := n
	m := n
	isStar := false
	for i := 1; i <= n*2-1; i++ {
		isStar = true
		for j := 1; j <= n*2-1; j++ {
			if j >= l && j <= m {
				if isStar {
					print("*")
					isStar = false
				} else {
					print(" ")
					isStar = true
				}
			} else {
				print(" ")
			}
		}

		if i < n {
			l--
			m++
		} else {
			l++
			m--
		}

		println()
	}
}

func BinarySearch(arr []int, x int) int {
	// considering the array is in increasing order
	i := 0
	j := len(arr) - 1
	for i <= j {
		mid := (i + j) / 2
		if arr[mid] < x {
			i = mid + 1
		} else if arr[mid] > x {
			j = mid - 1
		} else {
			return arr[mid]
		}
	}
	return -1
}

func BinarySearchUsingRecursion(arr []int, i, j, x int) int {
	if i > j {
		return -1
	}

	mid := (i + j) / 2
	if arr[mid] == x {
		return x
	}

	if arr[mid] > x {
		return BinarySearchUsingRecursion(arr, i, mid-1, x)
	} else {
		return BinarySearchUsingRecursion(arr, mid+1, j, x)
	}
}
