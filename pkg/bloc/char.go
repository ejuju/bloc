package bloc

import (
	"image"
	"image/color"
	"math"
)

const Width = 7
const Height = 15

type Char [Height]uint8

func CharFromRune(v rune) Char {
	if int(v) >= len(ASCIIChars) {
		return CharUnknown
	}
	return ASCIIChars[v]
}

var ASCIIChars = [128]Char{
	0:    CharUnknown,         // Blank (non-graphic character).
	1:    CharUnknown,         // Blank (non-graphic character).
	2:    CharUnknown,         // Blank (non-graphic character).
	3:    CharUnknown,         // Blank (non-graphic character).
	4:    CharUnknown,         // Blank (non-graphic character).
	5:    CharUnknown,         // Blank (non-graphic character).
	6:    CharUnknown,         // Blank (non-graphic character).
	7:    CharUnknown,         // Blank (non-graphic character).
	8:    CharUnknown,         // Blank (non-graphic character).
	9:    CharUnknown,         // Blank (non-graphic character).
	10:   CharUnknown,         // Blank (non-graphic character).
	11:   CharUnknown,         // Blank (non-graphic character).
	12:   CharUnknown,         // Blank (non-graphic character).
	13:   CharUnknown,         // Blank (non-graphic character).
	14:   CharUnknown,         // Blank (non-graphic character).
	15:   CharUnknown,         // Blank (non-graphic character).
	16:   CharUnknown,         // Blank (non-graphic character).
	17:   CharUnknown,         // Blank (non-graphic character).
	18:   CharUnknown,         // Blank (non-graphic character).
	19:   CharUnknown,         // Blank (non-graphic character).
	20:   CharUnknown,         // Blank (non-graphic character).
	21:   CharUnknown,         // Blank (non-graphic character).
	22:   CharUnknown,         // Blank (non-graphic character).
	23:   CharUnknown,         // Blank (non-graphic character).
	24:   CharUnknown,         // Blank (non-graphic character).
	25:   CharUnknown,         // Blank (non-graphic character).
	26:   CharUnknown,         // Blank (non-graphic character).
	27:   CharUnknown,         // Blank (non-graphic character).
	28:   CharUnknown,         // Blank (non-graphic character).
	29:   CharUnknown,         // Blank (non-graphic character).
	30:   CharUnknown,         // Blank (non-graphic character).
	31:   CharUnknown,         // Blank (non-graphic character).
	' ':  CharSpace,           // 32
	'!':  CharExclamationMark, // 33
	'"':  CharDoubleQuotes,    // 34
	'#':  CharHashtag,         // 35
	'$':  CharDollar,          // 36
	'%':  CharPercent,         // 37
	'&':  CharAmpersand,       // 38
	'\'': CharSingleQuote,     // 39
	'(':  CharParenLeft,       // 40
	')':  CharParenRight,      // 41
	'*':  CharAsterisk,        // 42
	'+':  CharPlus,            // 43
	',':  CharComma,           // 44
	'-':  CharHyphen,          // 45
	'.':  CharDot,             // 46
	'/':  CharSlash,           // 47
	'0':  Char0,
	'1':  Char1,
	'2':  Char2,
	'3':  Char3,
	'4':  Char4,
	'5':  Char5,
	'6':  Char6,
	'7':  Char7,
	'8':  Char8,
	'9':  Char9,
	':':  CharColon,        // 58
	';':  CharSemiColon,    // 59
	'<':  CharLessThan,     // 60
	'=':  CharEqual,        // 61
	'>':  CharGreaterThan,  // 62
	'?':  CharQuestionMark, // 63
	'@':  CharAt,           // 64
	'A':  CharA,
	'B':  CharB,
	'C':  CharC,
	'D':  CharD,
	'E':  CharE,
	'F':  CharF,
	'G':  CharG,
	'H':  CharH,
	'I':  CharI,
	'J':  CharJ,
	'K':  CharK,
	'L':  CharL,
	'M':  CharM,
	'N':  CharN,
	'O':  CharO,
	'P':  CharP,
	'Q':  CharQ,
	'R':  CharR,
	'S':  CharS,
	'T':  CharT,
	'U':  CharU,
	'V':  CharV,
	'W':  CharW,
	'X':  CharX,
	'Y':  CharY,
	'Z':  CharZ,
	'[':  CharBracketLeft,  // 91
	'\\': CharBackSlash,    // 92
	']':  CharBracketRight, // 93
	'^':  CharCircumflex,   // 94
	'_':  CharUnderscore,   // 95
	'`':  CharBacktick,     // 96
	'a':  CharALower,
	'b':  CharBLower,
	'c':  CharCLower,
	'd':  CharDLower,
	'e':  CharELower,
	'f':  CharFLower,
	'g':  CharGLower,
	'h':  CharHLower,
	'i':  CharILower,
	'j':  CharJLower,
	'k':  CharKLower,
	'l':  CharLLower,
	'm':  CharMLower,
	'n':  CharNLower,
	'o':  CharOLower,
	'p':  CharPLower,
	'q':  CharQLower,
	'r':  CharRLower,
	's':  CharSLower,
	't':  CharTLower,
	'u':  CharULower,
	'v':  CharVLower,
	'w':  CharWLower,
	'x':  CharXLower,
	'y':  CharYLower,
	'z':  CharZLower,
	'{':  CharBraceLeft,   // 123
	'|':  CharVerticalBar, // 124
	'}':  CharBraceRight,  // 125
	'~':  CharTilde,       // 126
	127:  CharUnknown,     // Blank (non-graphic character).
}

func (c Char) Image() *image.Gray {
	img := image.NewGray(image.Rect(0, 0, Width, Height))
	for y, row := range c {
		for x, cell := range uint7Cells(row) {
			clr := uint8(0)
			if cell {
				clr = math.MaxUint8
			}
			img.SetGray(x, y, color.Gray{clr})
		}
	}
	return img
}

func uint7Cells(v uint8) [Width]bool {
	return [...]bool{
		v>>6&0b1 > 0,
		v>>5&0b1 > 0,
		v>>4&0b1 > 0,
		v>>3&0b1 > 0,
		v>>2&0b1 > 0,
		v>>1&0b1 > 0,
		v>>0&0b1 > 0,
	}
}
