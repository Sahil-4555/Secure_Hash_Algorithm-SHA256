package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// roundConstants represents the round constants used in SHA-256 algorithm
var roundConstants = [64]string{
	"0x428a2f98", "0x71374491", "0xb5c0fbcf", "0xe9b5dba5",
	"0x3956c25b", "0x59f111f1", "0x923f82a4", "0xab1c5ed5",
	"0xd807aa98", "0x12835b01", "0x243185be", "0x550c7dc3",
	"0x72be5d74", "0x80deb1fe", "0x9bdc06a7", "0xc19bf174",
	"0xe49b69c1", "0xefbe4786", "0x0fc19dc6", "0x240ca1cc",
	"0x2de92c6f", "0x4a7484aa", "0x5cb0a9dc", "0x76f988da",
	"0x983e5152", "0xa831c66d", "0xb00327c8", "0xbf597fc7",
	"0xc6e00bf3", "0xd5a79147", "0x06ca6351", "0x14292967",
	"0x27b70a85", "0x2e1b2138", "0x4d2c6dfc", "0x53380d13",
	"0x650a7354", "0x766a0abb", "0x81c2c92e", "0x92722c85",
	"0xa2bfe8a1", "0xa81a664b", "0xc24b8b70", "0xc76c51a3",
	"0xd192e819", "0xd6990624", "0xf40e3585", "0x106aa070",
	"0x19a4c116", "0x1e376c08", "0x2748774c", "0x34b0bcb5",
	"0x391c0cb3", "0x4ed8aa4a", "0x5b9cca4f", "0x682e6ff3",
	"0x748f82ee", "0x78a5636f", "0x84c87814", "0x8cc70208",
	"0x90befffa", "0xa4506ceb", "0xbef9a3f7", "0xc67178f2",
}

// hexToBinary converts a hexadecimal string to binary representation
func hexToBinary(hex string) []int {
	// Remove the "0x" prefix if present
	hex = strings.TrimPrefix(hex, "0x")

	// Initialize the binary array
	binaryArray := make([]int, 32)

	// Convert each hexadecimal digit to binary
	for i, char := range hex {
		// Convert the hexadecimal digit to an integer
		var digit int
		switch {
		case char >= '0' && char <= '9':
			digit = int(char - '0')
		case char >= 'a' && char <= 'f':
			digit = int(char-'a') + 10
		case char >= 'A' && char <= 'F':
			digit = int(char-'A') + 10
		default:
			return binaryArray
		}

		for j := 0; j < 4; j++ {
			bit := (digit >> (3 - j)) & 1
			binaryArray[i*4+j] = bit
		}
	}

	return binaryArray
}

// rightShift performs right shift operation on an array of bits
func rightShift(arr []int, shiftCount int) []int {
	length := len(arr)
	// Create a new array for the shifted bits
	shifted := make([]int, length)

	// Perform the right shift operation
	for i := length - 1; i >= 0; i-- {
		if i-shiftCount >= 0 {
			shifted[i] = arr[i-shiftCount]
		} else {
			shifted[i] = 0
		}
	}
	return shifted
}

// rightRotate performs circular right rotation on an array of bits
func rightRotate(A []int, n int) []int {
	length := len(A)
	rotations := n % length

	// Create a new array for the rotated bits
	rotatedArray := make([]int, length)

	// Perform the circular right rotation
	for i := 0; i < length; i++ {
		rotatedArray[i] = A[(i-rotations+length)%length]
	}

	return rotatedArray
}

// xorOfThree performs XOR operation on three arrays of bits
func xorOfThree(A, B, C []int) []int {
	// Create a new array for the XOR result
	xorArray := make([]int, len(A))

	// Perform the XOR operation
	for i := 0; i < len(A); i++ {
		xorArray[i] = (A[i] + B[i] + C[i]) % 2
	}

	return xorArray
}

// xorOfTwo performs XOR operation on two arrays of bits
func xorOfTwo(A, B []int) []int {
	// Create a new array for the XOR result
	xorArray := make([]int, len(A))

	// Perform the XOR operation
	for i := 0; i < len(A); i++ {
		xorArray[i] = (A[i] + B[i]) % 2
	}

	return xorArray
}

// andOp performs logical AND operation on two arrays of bits
func andOp(A, B []int) []int {
	// Create a new array for the AND result
	andArray := make([]int, len(A))

	// Perform the logical AND operation
	for i := 0; i < len(A); i++ {
		andArray[i] = A[i] & B[i]
	}

	return andArray
}

// notOp performs logical NOT operation on an array of bits
func notOp(A []int) []int {
	// Create a new array for the NOT result
	notArray := make([]int, len(A))

	// Perform the logical NOT operation
	for i := 0; i < len(A); i++ {
		notArray[i] = 1 - A[i]
	}

	return notArray
}

// addTwo adds two arrays of bits as binary numbers
func addTwo(A, B []int) []int {
	// Create a new array for the sum
	sumArray := make([]int, len(A))

	// Initialize the carry
	carry := 0

	// Perform the addition
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i]
		sumArray[i] = sum % 2
		carry = sum / 2
	}

	return sumArray
}

// addFour adds four arrays of bits as binary numbers
func addFour(A, B, C, D []int) []int {
	// Create a new array for the sum
	sumArray := make([]int, len(A))

	// Initialize the carry
	carry := 0

	// Perform the addition
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i] + C[i] + D[i]
		sumArray[i] = sum % 2
		carry = sum / 2
	}

	return sumArray
}

// addFive adds five arrays of bits as binary numbers
func addFive(A, B, C, D, E []int) []int {
	// Create a new array for the sum
	sumArray := make([]int, len(A))

	// Initialize the carry
	carry := 0

	// Perform the addition
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i] + C[i] + D[i] + E[i]
		sumArray[i] = sum % 2
		carry = sum / 2
	}

	return sumArray
}

// This function converts the binary representation stored in the input slice A into a hexadecimal string.
func getDigest(A []int) string {
	// Initialize an empty string
	s := ""

	for i := 0; i < 256; i += 4 {
		// Convert four binary bits to a decimal value
		an := A[i]*8 + A[i+1]*4 + A[i+2]*2 + A[i+3]
		if an <= 9 {
			// Append the decimal value to the string
			s += strconv.Itoa(an)
		} else if an == 10 {
			// Append 'a' if the decimal value is 10
			s += "a"
		} else if an == 11 {
			// Append 'b' if the decimal value is 11
			s += "b"
		} else if an == 12 {
			// Append 'c' if the decimal value is 12
			s += "c"
		} else if an == 13 {
			// Append 'd' if the decimal value is 13
			s += "d"
		} else if an == 14 {
			// Append 'e' if the decimal value is 14
			s += "e"
		} else {
			// Append 'f' if the decimal value is 15
			s += "f"
		}
		// s += strconv.FormatInt(int64(an), 16)
	}
	// Return the hexadecimal string
	return s
}

// This function reverses the order of elements in the input slice arr.
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		// Swap elements at positions i and j
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// This function takes an input slice of integers (inputArray), which represents the input message, and performs various operations to prepare the message for hashing.
func Encrypt(inputArray []int) []int {

	// Convert the all string into binary
	length := len(inputArray)
	messageBit := []int{}

	// convert the char to binary
	for i := 0; i < length; i++ {
		a := inputArray[i]
		anss := []int{}
		// Convert the character to binary representation
		for a != 0 {
			h := a % 2
			anss = append(anss, h)
			a = a / 2
		}

		// Append leading zeros until the binary representation becomes 8 bits
		tmp := len(anss)
		for j := 0; j < 8-tmp; j++ {
			anss = append(anss, 0)
		}
		reverse(anss)

		// Append the binary representation of the character to the message bits
		for j := 0; j < 8; j++ {
			messageBit = append(messageBit, anss[j])
		}
	}

	// Append a '1' to the message bits
	messageBit = append(messageBit, 1)

	// Append zeros until the length of message bits is a multiple of 512 and has a remainder of 448
	for len(messageBit)%512 != 448 {
		messageBit = append(messageBit, 0)
	}

	// Calculate the length of the original message in bits
	length = length * 8

	// Create a slice of size 64 to store the length of the message in binary representation
	ml := make([]int, 64)
	p := 63

	// Convert the length to binary representation
	for length != 0 {
		h := length % 2
		ml[p] = h
		length = length / 2
		p = p - 1
	}

	// Append the binary representation of the length to the message bits
	for j := 0; j < 64; j++ {
		messageBit = append(messageBit, ml[j])
	}

	// Return the prepared message bits
	return messageBit
}

// This is the main hash function. It takes a string input inp, converts it to an array of ASCII values (inpBytes), and performs the SHA-1 hashing algorithm on the input message.
func Hash(inp string) string {

	// converts the individual char of string to its ascii value
	message := []rune(inp)

	inpBytes := make([]int, len(message))

	for i, c := range message {
		inpBytes[i] = int(c)
	}

	messageBit := Encrypt(inpBytes)

	// INITIALIZE HASH VALUES
	A := []int{0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1} // A=[0x6a09e667] <- this is in hexadecimal form
	B := []int{1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1} // B=[0xbb67ae85]
	C := []int{0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0} // C=[0x3c6ef372]
	D := []int{1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0} // D=[0xa54ff53a]
	E := []int{0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1} // E=[0x510e527f]
	F := []int{1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0} // F=[0x9b05688c]
	G := []int{0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1} // G=[0x1f83d9ab]
	H := []int{0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1} // H=[0x5be0cd19]

	size := len(messageBit)
	nn := 0
	for n := 0; n < size; n += 512 {
		var w [][]int

		for i := nn; i < nn+16; i++ {
			var temp []int
			for j := i * 32; j < (i+1)*32; j++ {
				temp = append(temp, messageBit[j])
			}
			w = append(w, temp)
		}

		for i := 16; i < 64; i++ {
			// s0 = (w[i-15] rightrotate 7) xor (w[i-15] rightrotate 18) xor (w[i-15] rightshift 3)
			s0 := xorOfThree(rightRotate(w[i-15], 7), rightRotate(w[i-15], 18), rightShift(w[i-15], 3))
			// s1 = (w[i- 2] rightrotate 17) xor (w[i- 2] rightrotate 19) xor (w[i- 2] rightshift 10)
			s1 := xorOfThree(rightRotate(w[i-2], 17), rightRotate(w[i-2], 19), rightShift(w[i-2], 10))
			// w[i] = w[i-16] + s0 + w[i-7] + s1
			w = append(w, addFour(w[i-16], s0, w[i-7], s1))
		}

		a := make([]int, len(A))
		b := make([]int, len(B))
		c := make([]int, len(C))
		d := make([]int, len(D))
		e := make([]int, len(E))
		f := make([]int, len(F))
		g := make([]int, len(G))
		h := make([]int, len(H))
		// a = A
		copy(a, A)
		// b = B
		copy(b, B)
		// c = C
		copy(c, C)
		// d = D
		copy(d, D)
		// e = E
		copy(e, E)
		// f = F
		copy(f, F)
		// g = G
		copy(g, G)
		// h = H
		copy(h, H)

		for i := 0; i < 64; i++ {
			// s1 = (e rightrotate 6) xor (e rightrotate 11) xor (e rightrotate 25)
			s1 := xorOfThree(rightRotate(e, 6), rightRotate(e, 11), rightRotate(e, 25))
			//choose = (e and f) xor ((not e) and g)
			choose := xorOfTwo(andOp(e, f), andOp(notOp(e), g))
			// temp1 = h + s1 + choose + roundconstants[i] + w[i]
			temp1 := andOp(addFive(h, s1, choose, hexToBinary(roundConstants[i]), w[i]), hexToBinary("0xffffffff"))
			// s0 = (a rightrotate 2) xor (a rightrotate 13) xor (a rightrotate 22)
			s0 := xorOfThree(rightRotate(a, 2), rightRotate(a, 13), rightRotate(a, 22))
			// majority = (a and b) xor (a and c) xor (b and c)
			majority := xorOfThree(andOp(a, b), andOp(a, c), andOp(b, c))
			// temp2 = s0 + majority
			temp2 := andOp(addTwo(s0, majority), hexToBinary("0xffffffff"))

			// h = g
			copy(h, g)
			// g = f
			copy(g, f)
			// f = e
			copy(f, e)
			// e = d + temp1
			copy(e, andOp(addTwo(d, temp1), hexToBinary("0xffffffff")))
			// d = c
			copy(d, c)
			// c = b
			copy(c, b)
			// b = a
			copy(b, a)
			// a = temp1 + temp2
			copy(a, andOp(addTwo(temp1, temp2), hexToBinary("0xffffffff")))
		}
		// A = A + a
		copy(A, addTwo(A, a))
		// B = B + b
		copy(B, addTwo(B, b))
		// C = C + c
		copy(C, addTwo(C, c))
		// D = D + d
		copy(D, addTwo(D, d))
		// E = E + e
		copy(E, addTwo(E, e))
		// F = F + f
		copy(F, addTwo(F, f))
		// G = G + g
		copy(G, addTwo(G, g))
		// H = H + h
		copy(H, addTwo(H, h))

		nn += 16
	}

	// digest = A append B append C append D append E append F append G append H
	dig := getDigest(append(append(append(append(append(append(append(A, B...), C...), D...), E...), F...), G...), H...))
	return dig
}

// The main function handles user input and calls the Hash function to compute the SHA-256 hash of the input message.
func main() {

	// taking input string.
	var input string
	fmt.Println("Enter Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	// if input is empty.
	if strings.TrimSpace(input) == "" {
		fmt.Println("Please Enter Input.")
		os.Exit(1)
	}

	// Output the digest
	fmt.Println("Digest (SHA-256):", Hash(input))
}
