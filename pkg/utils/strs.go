package utils

//import "log"

// '3' -> 0x03, 'A' -> 0x0A
func char2Bcd(c uint8) int {
	var b int
	if c > 0x39 {
		b = int(c - 0x37)
	} else {
		b = int(c - 0x30)
	}
	return b
}

// [ 0x32, 0x33 ] -> 0x23
// func symHex2Hex(s []uint8, offset int) int {
// 	var r int
// 	r = char2Bcd(s[offset]) << 4
// 	r += char2Bcd(s[offset + 1])
// 	return r
// }
// "23" -> 0x23
func sym2Hex(s string, offset int) int {
	// b := []uint8(s)
	// r := symHex2Hex(b, offset)
	var r int
	r = char2Bcd(s[offset]) << 4
	r += char2Bcd(s[offset+1])
	return r
}

func Str2Hex(s string) []int {
	len_dst := len(s) / 2
	dst := make([]int, len_dst)
	for i := 0; i < len_dst; i++ {
		dst[i] = sym2Hex(s, i*2)
	}
	//log.Printf("src= %s, dst= % X ", s, dst)
	return dst
}

//------------------------------

// 0xC1 -> [1,0,0,0, 0,0,1,1]
func hex2Bits(c int) []int {
	b := make([]int, 8)
	for i := 0; i < 8; i++ {
		if c&(1<<i) != 0 {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
	return b
}
func hexBuf2Bits(src []int) []int {
	len_src := len(src)
	dst := make([]int, len_src*8)
	for i := 0; i < len_src; i++ {
		a := hex2Bits(src[i])
		for j := 0; j < 8; j++ {
			dst[i*8+j] = a[j]
		}
	}
	return dst
}

// "0482" -> [0,0,1,0,0,0,0,0, 0,1,0,0,0,0,0,1]
func Str2Bits(s string) []int {
	b := Str2Hex(s)
	dst := hexBuf2Bits(b)
	//log.Printf("src= %s, dst= % X ", s, dst)
	return dst
}

//====================================

func bcd2CharH(b int) uint8 {
	var c uint8
	c = uint8(b >> 4 & 0x0f)
	if c > 9 {
		c = c + 0x37
	} else {
		c = c + 0x30
	}
	return c
}
func bcd2CharL(b int) uint8 {
	var c uint8
	c = uint8(b & 0x0f)
	if c > 9 {
		c = c + 0x37
	} else {
		c = c + 0x30
	}
	return c
}

// 0x25 0x31 -> "2531"
func Hex2Str(buf []int, len_src int) string {
	len_dst := len_src * 2
	strBuf := make([]uint8, len_dst)
	for i := 0; i < len_src; i++ {
		strBuf[i*2] = bcd2CharH(buf[i])
		strBuf[i*2+1] = bcd2CharL(buf[i])
	}
	s := string(strBuf)
	return s
}

func StrNCmp(b1 []byte, b2 []byte) int {
	len1 := len(b1)
	len2 := len(b2)
	len := 0
	if len1 >= len2 {
		len = len2
	} else {
		len = len1
	}
	for i := 0; i < len; i++ {
		if b1[i] != b2[i] {
			return 1
		}
	}
	return 0
}
