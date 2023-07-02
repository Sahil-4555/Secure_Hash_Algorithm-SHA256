package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func rightShift(arr []int, shiftCount int) []int {
	length := len(arr)
	shifted := make([]int, length)

	// Perform right shift operation on each element
	for i := length - 1; i >= 0; i-- {
		if i-shiftCount >= 0 {
			shifted[i] = arr[i-shiftCount]
		} else {
			shifted[i] = 0
		}
	}
	return shifted
}

func rightRotate(A []int, n int) []int {
	length := len(A)
	rotations := n % length

	result := make([]int, length)

	// Copy the elements from A to the rotated positions in the result array
	for i := 0; i < length; i++ {
		result[i] = A[(i-rotations+length)%length]
	}

	return result
}

// The xorOfThree function performs the XOR operation on four input slices A, B and C element-wise.
func xorOfThree(A, B, C []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = (A[i] + B[i] + C[i]) % 2
	}
	return temp
}

// The xorOfTwo function performs the XOR operation on four input slices A and B element-wise.
func xorOfTwo(A, B []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = (A[i] + B[i]) % 2
	}
	return temp
}

// The andOp function performs the logical AND operation on two input slices A and B element-wise.
func andOp(A, B []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = A[i] & B[i]
	}
	return temp
}

// The notOp function performs the logical NOT operation on the input slice A element-wise.
func notOp(A []int) []int {
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[i] = 1 - A[i]
	}
	return temp
}

// The addTwo function adds two input slices A and B as binary numbers.
func addTwo(A, B []int) []int {
	temp := make([]int, len(A))
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i]
		temp[i] = sum % 2
		carry = sum / 2
	}
	return temp
}

// This function adds four input slices A, B, C and D as binary numbers.
func addFour(A, B, C, D []int) []int {
	temp := make([]int, len(A))
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i] + C[i] + D[i]
		temp[i] = sum % 2
		carry = sum / 2
	}
	return temp
}

// This function adds five input slices A, B, C, D, and E as binary numbers.
func addFive(A, B, C, D, E []int) []int {
	temp := make([]int, len(A))
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum := carry + A[i] + B[i] + C[i] + D[i] + E[i]
		temp[i] = sum % 2
		carry = sum / 2
	}
	return temp
}

// This function converts the binary representation stored in the input slice A into a hexadecimal string.
func getDigest(A []int) string {
	s := ""
	for i := 0; i < 256; i += 4 {
		an := A[i]*8 + A[i+1]*4 + A[i+2]*2 + A[i+3]
		if an <= 9 {
			s += strconv.Itoa(an)
		} else if an == 10 {
			s += "a"
		} else if an == 11 {
			s += "b"
		} else if an == 12 {
			s += "c"
		} else if an == 13 {
			s += "d"
		} else if an == 14 {
			s += "e"
		} else {
			s += "f"
		}
		// s += strconv.FormatInt(int64(an), 16)
	}
	return s
}

// This function reverses the order of elements in the input slice arr.
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
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
			s0 := xorOfThree(rightRotate(w[i-15], 7), rightRotate(w[i-15], 18), rightShift(w[i-15], 3))
			s1 := xorOfThree(rightRotate(w[i-2], 17), rightRotate(w[i-2], 19), rightShift(w[i-2], 10))
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
		copy(a, A)
		copy(b, B)
		copy(c, C)
		copy(d, D)
		copy(e, E)
		copy(f, F)
		copy(g, G)
		copy(h, H)

		for i := 0; i < 64; i++ {
			s1 := xorOfThree(rightRotate(e, 6), rightRotate(e, 11), rightRotate(e, 25))
			// if i == 1 {
			// 	// fmt.Println("e6: ", i, rightRotate(e, 6))
			// 	// fmt.Println("e11: ", i, rightRotate(e, 11))
			// 	// fmt.Println("e25: ", i, rightRotate(e, 25))
			// 	fmt.Println("s1: ", i, s1)
			// }
			choose := xorOfTwo(andOp(e, f), andOp(notOp(e), g))
			// if i == 1 {
			// 	fmt.Println("choose: ", i, choose)
			// }
			temp1 := andOp(addFive(h, s1, choose, hexToBinary(roundConstants[i]), w[i]), hexToBinary("0xffffffff"))
			// if i == 1 {
			// 	fmt.Println("h+s1: ", addTwo(h, s1))
			// 	fmt.Println("ans+choose: ", addTwo(addTwo(h, s1), choose))
			// 	fmt.Println("ans+k[i]: ", addTwo(addTwo(addTwo(h, s1), choose), hexToBinary(roundConstants[i])))
			// 	fmt.Println("ans+w[i]: ", addTwo(addTwo(addTwo(addTwo(h, s1), choose), hexToBinary(roundConstants[i])), w[i]))
			// }
			s0 := xorOfThree(rightRotate(a, 2), rightRotate(a, 13), rightRotate(a, 22))
			// if i == 1 {
			// 	fmt.Println("s0: ", i, s0)
			// }
			majority := xorOfThree(andOp(a, b), andOp(a, c), andOp(b, c))
			// if i == 1 {
			// 	fmt.Println("majority: ", i, majority)
			// }
			temp2 := andOp(addTwo(s0, majority), hexToBinary("0xffffffff"))
			// if i == 1 {
			// 	fmt.Println("temp2: ", i, temp2)
			// }
			copy(h, g)
			copy(g, f)
			copy(f, e)
			copy(e, andOp(addTwo(d, temp1), hexToBinary("0xffffffff")))
			// if i == 1 {
			// 	fmt.Println("e: ", i, e)
			// }
			copy(d, c)
			copy(c, b)
			copy(b, a)
			copy(a, andOp(addTwo(temp1, temp2), hexToBinary("0xffffffff")))
			// if i == 1 {
			// 	fmt.Println("a: ", i, a)
			// }
		}
		// fmt.Println("A: ", A)
		// fmt.Println("a: ", a)
		copy(A, addTwo(A, a))
		// fmt.Println("B: ", B)
		// fmt.Println("b: ", b)
		copy(B, addTwo(B, b))
		// fmt.Println("C: ", C)
		// fmt.Println("c: ", c)
		copy(C, addTwo(C, c))
		// fmt.Println("D: ", D)
		// fmt.Println("d: ", d)
		copy(D, addTwo(D, d))
		// fmt.Println("E: ", E)
		// fmt.Println("e: ", e)
		copy(E, addTwo(E, e))
		// fmt.Println("F: ", F)
		// fmt.Println("f: ", f)
		copy(F, addTwo(F, f))
		// fmt.Println("G: ", G)
		// fmt.Println("g: ", g)
		copy(G, addTwo(G, g))
		// fmt.Println("H: ", H)
		// fmt.Println("h: ", h)
		copy(H, addTwo(H, h))

		nn += 16
	}

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
