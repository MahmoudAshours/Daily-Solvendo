type ProductOfNumbers struct {
	prefixProducts []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefixProducts: []int{1}, // Initialize with 1 to handle the first product calculation
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		// Reset the prefix products if we encounter a 0
		this.prefixProducts = []int{1}
	} else {
		// Append the product of the last prefix product and the new number
		this.prefixProducts = append(
			this.prefixProducts,
			this.prefixProducts[len(this.prefixProducts)-1]*num,
		)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefixProducts)
	if k >= n {
		// If k is greater than or equal to the length of the prefix products, it means we have a 0 in the last k numbers
		return 0
	}
	// Calculate the product of the last k numbers by dividing the last prefix product by the prefix product at n-k-1
	return this.prefixProducts[n-1] / this.prefixProducts[n-k-1]
}