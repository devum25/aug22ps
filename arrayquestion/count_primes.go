package arrayquestion

// 204. Count Primes
// Attempted
// Medium
// Topics
// Companies
// Hint
// Given an integer n, return the number of prime numbers that are strictly less than n.

// Example 1:

// Input: n = 10
// Output: 4
// Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

// Intuition
// The code efficiently counts prime numbers less than n using the Sieve of Eratosthenes. It initializes an array, marks multiples of primes, and counts remaining primes. Skipping evens and optimizing the sieve contribute to its efficiency.

// Approach
// Initialization:
// Initialize a boolean array isPrime of size n.
// Set all odd numbers greater than or equal to 3 to true, skipping even numbers (except 2)

// Sieve of Eratosthenes:
// Iterate through odd numbers from 3 to the square root of n.
// For each prime i, mark multiples of i as false in the array

// Count Primes:
// Count the number of true values in the array, representing prime numbers less than n.

// Optimizations:
// Skip even numbers (except 2) in the initialization and sieve steps.
// Limit sieve iterations to the square root of n.

// Complexity
// Screenshot 2023-12-08 153500.png

// Time complexity:
// The Sieve of Eratosthenes has a time complexity of approximately O(n log log n), where n is the input size.

// Space complexity:
// O(n) - The algorithm uses an array of size n to store whether each number is prime.

// Code

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	composites := make([]bool, n)

	for i := 2; i*i < n; i++ {
		if !composites[i] {
			for j := i * i; j < n; j = j + i {
				composites[j] = true
			}
		}
	}
	ans := 0
	for i := 2; i < n; i++ {
		if !composites[i] {
			ans++
		}
	}

	return ans
}
