// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/wxio/goantlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 265,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 3, 7, 3, 34, 10, 3,
	12, 3, 14, 3, 37, 11, 3, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14, 3, 43, 11, 3,
	3, 3, 5, 3, 46, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4,
	54, 11, 4, 3, 4, 5, 4, 57, 10, 4, 3, 4, 7, 4, 60, 10, 4, 12, 4, 14, 4,
	63, 11, 4, 3, 4, 5, 4, 66, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 71, 10, 4, 12,
	4, 14, 4, 74, 11, 4, 3, 4, 5, 4, 77, 10, 4, 3, 4, 7, 4, 80, 10, 4, 12,
	4, 14, 4, 83, 11, 4, 3, 4, 5, 4, 86, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 91,
	10, 4, 12, 4, 14, 4, 94, 11, 4, 3, 4, 5, 4, 97, 10, 4, 3, 4, 5, 4, 100,
	10, 4, 3, 4, 7, 4, 103, 10, 4, 12, 4, 14, 4, 106, 11, 4, 3, 4, 5, 4, 109,
	10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 114, 10, 4, 12, 4, 14, 4, 117, 11, 4, 3,
	4, 5, 4, 120, 10, 4, 3, 4, 5, 4, 123, 10, 4, 3, 4, 7, 4, 126, 10, 4, 12,
	4, 14, 4, 129, 11, 4, 3, 4, 5, 4, 132, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 140, 10, 4, 12, 4, 14, 4, 143, 11, 4, 3, 4, 3, 4, 7, 4,
	147, 10, 4, 12, 4, 14, 4, 150, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 156,
	10, 4, 12, 4, 14, 4, 159, 11, 4, 3, 4, 3, 4, 5, 4, 163, 10, 4, 3, 4, 7,
	4, 166, 10, 4, 12, 4, 14, 4, 169, 11, 4, 3, 4, 5, 4, 172, 10, 4, 5, 4,
	174, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 179, 10, 5, 12, 5, 14, 5, 182, 11,
	5, 3, 5, 5, 5, 185, 10, 5, 3, 5, 5, 5, 188, 10, 5, 3, 5, 5, 5, 191, 10,
	5, 3, 6, 3, 6, 3, 6, 6, 6, 196, 10, 6, 13, 6, 14, 6, 197, 3, 6, 3, 6, 5,
	6, 202, 10, 6, 3, 7, 3, 7, 3, 7, 6, 7, 207, 10, 7, 13, 7, 14, 7, 208, 3,
	7, 3, 7, 5, 7, 213, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 240, 10, 8, 13, 8, 14, 8, 241, 3,
	8, 3, 8, 5, 8, 246, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 254,
	10, 8, 13, 8, 14, 8, 255, 3, 8, 3, 8, 5, 8, 260, 10, 8, 3, 8, 5, 8, 263,
	10, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 3, 2, 34, 35, 3, 2,
	36, 37, 2, 311, 2, 16, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 173, 3, 2, 2,
	2, 8, 175, 3, 2, 2, 2, 10, 192, 3, 2, 2, 2, 12, 203, 3, 2, 2, 2, 14, 262,
	3, 2, 2, 2, 16, 17, 7, 26, 2, 2, 17, 18, 7, 30, 2, 2, 18, 19, 7, 26, 2,
	2, 19, 20, 5, 4, 3, 2, 20, 21, 7, 27, 2, 2, 21, 22, 7, 27, 2, 2, 22, 23,
	7, 2, 2, 3, 23, 3, 3, 2, 2, 2, 24, 45, 7, 31, 2, 2, 25, 29, 7, 26, 2, 2,
	26, 28, 7, 33, 2, 2, 27, 26, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3,
	2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 35, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32,
	34, 7, 32, 2, 2, 33, 32, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2,
	2, 2, 35, 36, 3, 2, 2, 2, 36, 41, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 40,
	5, 6, 4, 2, 39, 38, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	41, 42, 3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 46, 7,
	27, 2, 2, 45, 25, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 5, 3, 2, 2, 2, 47,
	65, 7, 34, 2, 2, 48, 52, 7, 26, 2, 2, 49, 51, 7, 33, 2, 2, 50, 49, 3, 2,
	2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 56,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 57, 7, 38, 2, 2, 56, 55, 3, 2, 2, 2,
	56, 57, 3, 2, 2, 2, 57, 61, 3, 2, 2, 2, 58, 60, 5, 8, 5, 2, 59, 58, 3,
	2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62,
	64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 66, 7, 27, 2, 2, 65, 48, 3, 2,
	2, 2, 65, 66, 3, 2, 2, 2, 66, 174, 3, 2, 2, 2, 67, 85, 7, 35, 2, 2, 68,
	72, 7, 26, 2, 2, 69, 71, 7, 33, 2, 2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2,
	2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72,
	3, 2, 2, 2, 75, 77, 7, 38, 2, 2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2,
	77, 81, 3, 2, 2, 2, 78, 80, 5, 8, 5, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3,
	2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83,
	81, 3, 2, 2, 2, 84, 86, 7, 27, 2, 2, 85, 68, 3, 2, 2, 2, 85, 86, 3, 2,
	2, 2, 86, 174, 3, 2, 2, 2, 87, 108, 7, 37, 2, 2, 88, 92, 7, 26, 2, 2, 89,
	91, 7, 33, 2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2,
	2, 2, 92, 93, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 97,
	7, 38, 2, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2,
	98, 100, 5, 10, 6, 2, 99, 98, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 104,
	3, 2, 2, 2, 101, 103, 5, 14, 8, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2,
	2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2,
	106, 104, 3, 2, 2, 2, 107, 109, 7, 27, 2, 2, 108, 88, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 174, 3, 2, 2, 2, 110, 131, 7, 36, 2, 2, 111, 115,
	7, 26, 2, 2, 112, 114, 7, 33, 2, 2, 113, 112, 3, 2, 2, 2, 114, 117, 3,
	2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 119, 3, 2, 2,
	2, 117, 115, 3, 2, 2, 2, 118, 120, 7, 38, 2, 2, 119, 118, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121, 123, 5, 10, 6, 2, 122, 121,
	3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 127, 3, 2, 2, 2, 124, 126, 5, 14,
	8, 2, 125, 124, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2,
	127, 128, 3, 2, 2, 2, 128, 130, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130,
	132, 7, 27, 2, 2, 131, 111, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 174,
	3, 2, 2, 2, 133, 174, 7, 50, 2, 2, 134, 174, 7, 51, 2, 2, 135, 174, 7,
	52, 2, 2, 136, 137, 9, 2, 2, 2, 137, 141, 7, 26, 2, 2, 138, 140, 7, 33,
	2, 2, 139, 138, 3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2,
	141, 142, 3, 2, 2, 2, 142, 144, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 144,
	148, 7, 29, 2, 2, 145, 147, 5, 8, 5, 2, 146, 145, 3, 2, 2, 2, 147, 150,
	3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 151, 3, 2,
	2, 2, 150, 148, 3, 2, 2, 2, 151, 174, 7, 27, 2, 2, 152, 153, 9, 3, 2, 2,
	153, 157, 7, 26, 2, 2, 154, 156, 7, 33, 2, 2, 155, 154, 3, 2, 2, 2, 156,
	159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 160,
	3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 162, 7, 29, 2, 2, 161, 163, 5, 10,
	6, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 167, 3, 2, 2, 2,
	164, 166, 5, 14, 8, 2, 165, 164, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167,
	165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167,
	3, 2, 2, 2, 170, 172, 7, 27, 2, 2, 171, 170, 3, 2, 2, 2, 171, 172, 3, 2,
	2, 2, 172, 174, 3, 2, 2, 2, 173, 47, 3, 2, 2, 2, 173, 67, 3, 2, 2, 2, 173,
	87, 3, 2, 2, 2, 173, 110, 3, 2, 2, 2, 173, 133, 3, 2, 2, 2, 173, 134, 3,
	2, 2, 2, 173, 135, 3, 2, 2, 2, 173, 136, 3, 2, 2, 2, 173, 152, 3, 2, 2,
	2, 174, 7, 3, 2, 2, 2, 175, 190, 7, 41, 2, 2, 176, 180, 7, 26, 2, 2, 177,
	179, 7, 33, 2, 2, 178, 177, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 178,
	3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2,
	2, 2, 183, 185, 5, 10, 6, 2, 184, 183, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2,
	185, 187, 3, 2, 2, 2, 186, 188, 5, 14, 8, 2, 187, 186, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 191, 7, 27, 2, 2, 190, 176,
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 9, 3, 2, 2, 2, 192, 201, 7, 39,
	2, 2, 193, 195, 7, 26, 2, 2, 194, 196, 5, 12, 7, 2, 195, 194, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198,
	199, 3, 2, 2, 2, 199, 200, 7, 27, 2, 2, 200, 202, 3, 2, 2, 2, 201, 193,
	3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 11, 3, 2, 2, 2, 203, 212, 7, 40,
	2, 2, 204, 206, 7, 26, 2, 2, 205, 207, 5, 12, 7, 2, 206, 205, 3, 2, 2,
	2, 207, 208, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209,
	210, 3, 2, 2, 2, 210, 211, 7, 27, 2, 2, 211, 213, 3, 2, 2, 2, 212, 204,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 13, 3, 2, 2, 2, 214, 215, 7, 42,
	2, 2, 215, 216, 7, 26, 2, 2, 216, 217, 7, 43, 2, 2, 217, 263, 7, 27, 2,
	2, 218, 219, 7, 42, 2, 2, 219, 220, 7, 26, 2, 2, 220, 221, 7, 44, 2, 2,
	221, 263, 7, 27, 2, 2, 222, 223, 7, 42, 2, 2, 223, 224, 7, 26, 2, 2, 224,
	225, 7, 45, 2, 2, 225, 263, 7, 27, 2, 2, 226, 227, 7, 42, 2, 2, 227, 228,
	7, 26, 2, 2, 228, 229, 7, 46, 2, 2, 229, 263, 7, 27, 2, 2, 230, 231, 7,
	42, 2, 2, 231, 232, 7, 26, 2, 2, 232, 233, 7, 47, 2, 2, 233, 263, 7, 27,
	2, 2, 234, 235, 7, 42, 2, 2, 235, 236, 7, 26, 2, 2, 236, 245, 7, 48, 2,
	2, 237, 239, 7, 26, 2, 2, 238, 240, 5, 14, 8, 2, 239, 238, 3, 2, 2, 2,
	240, 241, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242,
	243, 3, 2, 2, 2, 243, 244, 7, 27, 2, 2, 244, 246, 3, 2, 2, 2, 245, 237,
	3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 263, 7, 27,
	2, 2, 248, 249, 7, 42, 2, 2, 249, 250, 7, 26, 2, 2, 250, 259, 7, 49, 2,
	2, 251, 253, 7, 26, 2, 2, 252, 254, 5, 14, 8, 2, 253, 252, 3, 2, 2, 2,
	254, 255, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256,
	257, 3, 2, 2, 2, 257, 258, 7, 27, 2, 2, 258, 260, 3, 2, 2, 2, 259, 251,
	3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 263, 7, 27,
	2, 2, 262, 214, 3, 2, 2, 2, 262, 218, 3, 2, 2, 2, 262, 222, 3, 2, 2, 2,
	262, 226, 3, 2, 2, 2, 262, 230, 3, 2, 2, 2, 262, 234, 3, 2, 2, 2, 262,
	248, 3, 2, 2, 2, 263, 15, 3, 2, 2, 2, 44, 29, 35, 41, 45, 52, 56, 61, 65,
	72, 76, 81, 85, 92, 96, 99, 104, 108, 115, 119, 122, 127, 131, 141, 148,
	157, 162, 167, 171, 173, 180, 184, 187, 190, 197, 201, 208, 212, 241, 245,
	255, 259, 262,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'['", "']'", "'='", "'\"'", "'''", "';'", "'::'", "':'",
	"'.'", "','", "'<'", "'>'", "'*'", "'@'",
}
var symbolicNames = []string{
	"", "LCUR", "RCUR", "LSQ", "RSQ", "EQ", "DQ", "SQ", "SEMI", "DCOLON", "COLON",
	"DOT", "COMMA", "LCHEVR", "RCHEVR", "STAR", "AT", "STR", "ID", "INT", "FLT",
	"WS", "LINE_DOC", "LINE_COMMENT", "DOWN", "UP", "ROOT", "ERROR", "ADL",
	"Module", "Import", "Annotation", "Struct", "Union", "Newtype", "Type",
	"TypeParam", "TypeExpr", "TypeExprElem", "Field", "Json", "JsonStr", "JsonBool",
	"JsonNull", "JsonInt", "JsonFloat", "JsonArray", "JsonObj", "ModuleAnno",
	"DeclAnno", "FieldAnno",
}

var ruleNames = []string{
	"adl", "module", "tld", "nameBody", "typeExpr_", "typeExprElem_", "jsonVal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ADLWalker struct {
	*antlr.BaseParser
}

func NewADLWalker(input antlr.TokenStream) *ADLWalker {
	this := new(ADLWalker)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ADLWalker.g4"

	return this
}

// ADLWalker tokens.
const (
	ADLWalkerEOF          = antlr.TokenEOF
	ADLWalkerLCUR         = 1
	ADLWalkerRCUR         = 2
	ADLWalkerLSQ          = 3
	ADLWalkerRSQ          = 4
	ADLWalkerEQ           = 5
	ADLWalkerDQ           = 6
	ADLWalkerSQ           = 7
	ADLWalkerSEMI         = 8
	ADLWalkerDCOLON       = 9
	ADLWalkerCOLON        = 10
	ADLWalkerDOT          = 11
	ADLWalkerCOMMA        = 12
	ADLWalkerLCHEVR       = 13
	ADLWalkerRCHEVR       = 14
	ADLWalkerSTAR         = 15
	ADLWalkerAT           = 16
	ADLWalkerSTR          = 17
	ADLWalkerID           = 18
	ADLWalkerINT          = 19
	ADLWalkerFLT          = 20
	ADLWalkerWS           = 21
	ADLWalkerLINE_DOC     = 22
	ADLWalkerLINE_COMMENT = 23
	ADLWalkerDOWN         = 24
	ADLWalkerUP           = 25
	ADLWalkerROOT         = 26
	ADLWalkerERROR        = 27
	ADLWalkerADL          = 28
	ADLWalkerModule       = 29
	ADLWalkerImport       = 30
	ADLWalkerAnnotation   = 31
	ADLWalkerStruct       = 32
	ADLWalkerUnion        = 33
	ADLWalkerNewtype      = 34
	ADLWalkerType         = 35
	ADLWalkerTypeParam    = 36
	ADLWalkerTypeExpr     = 37
	ADLWalkerTypeExprElem = 38
	ADLWalkerField        = 39
	ADLWalkerJson         = 40
	ADLWalkerJsonStr      = 41
	ADLWalkerJsonBool     = 42
	ADLWalkerJsonNull     = 43
	ADLWalkerJsonInt      = 44
	ADLWalkerJsonFloat    = 45
	ADLWalkerJsonArray    = 46
	ADLWalkerJsonObj      = 47
	ADLWalkerModuleAnno   = 48
	ADLWalkerDeclAnno     = 49
	ADLWalkerFieldAnno    = 50
)

// ADLWalker rules.
const (
	ADLWalkerRULE_adl           = 0
	ADLWalkerRULE_module        = 1
	ADLWalkerRULE_tld           = 2
	ADLWalkerRULE_nameBody      = 3
	ADLWalkerRULE_typeExpr_     = 4
	ADLWalkerRULE_typeExprElem_ = 5
	ADLWalkerRULE_jsonVal       = 6
)

// IAdlContext is an interface to support dynamic dispatch.
type IAdlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdlContext differentiates from other interfaces.
	IsAdlContext()
}

type AdlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdlContext() *AdlContext {
	var p = new(AdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

func (s *AdlContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *AdlContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *AdlContext) ADL() antlr.TerminalNode {
	return s.GetToken(ADLWalkerADL, 0)
}

func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *AdlContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLWalkerEOF, 0)
}

func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (p *ADLWalker) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ADLWalkerRULE_adl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Match(ADLWalkerDOWN)
	}
	{
		p.SetState(15)
		p.Match(ADLWalkerADL)
	}
	{
		p.SetState(16)
		p.Match(ADLWalkerDOWN)
	}
	{
		p.SetState(17)
		p.Module()
	}
	{
		p.SetState(18)
		p.Match(ADLWalkerUP)
	}
	{
		p.SetState(19)
		p.Match(ADLWalkerUP)
	}
	{
		p.SetState(20)
		p.Match(ADLWalkerEOF)
	}

	return localctx
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) Module() antlr.TerminalNode {
	return s.GetToken(ADLWalkerModule, 0)
}

func (s *ModuleContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *ModuleContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *ModuleContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *ModuleContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *ModuleContext) AllImport() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerImport)
}

func (s *ModuleContext) Import(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerImport, i)
}

func (s *ModuleContext) AllTld() []ITldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *ModuleContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *ADLWalker) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ADLWalkerRULE_module)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(ADLWalkerModule)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(23)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(24)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerImport {
			{
				p.SetState(30)
				p.Match(ADLWalkerImport)
			}

			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ADLWalkerStruct-32))|(1<<(ADLWalkerUnion-32))|(1<<(ADLWalkerNewtype-32))|(1<<(ADLWalkerType-32))|(1<<(ADLWalkerModuleAnno-32))|(1<<(ADLWalkerDeclAnno-32))|(1<<(ADLWalkerFieldAnno-32)))) != 0 {
			{
				p.SetState(36)
				p.Tld()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// ITldContext is an interface to support dynamic dispatch.
type ITldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTldContext differentiates from other interfaces.
	IsTldContext()
}

type TldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTldContext() *TldContext {
	var p = new(TldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_tld

	return p
}

func (s *TldContext) GetParser() antlr.Parser { return s.parser }

func (s *TldContext) CopyFrom(ctx *TldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeContext struct {
	*TldContext
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) Type() antlr.TerminalNode {
	return s.GetToken(ADLWalkerType, 0)
}

func (s *TypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *TypeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *TypeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *TypeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *TypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *TypeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *TypeContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitType(s)
	}
}

type TypeParamErrorContext struct {
	*TldContext
}

func NewTypeParamErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamErrorContext {
	var p = new(TypeParamErrorContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *TypeParamErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamErrorContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *TypeParamErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(ADLWalkerERROR, 0)
}

func (s *TypeParamErrorContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *TypeParamErrorContext) Struct() antlr.TerminalNode {
	return s.GetToken(ADLWalkerStruct, 0)
}

func (s *TypeParamErrorContext) Union() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUnion, 0)
}

func (s *TypeParamErrorContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *TypeParamErrorContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *TypeParamErrorContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *TypeParamErrorContext) Type() antlr.TerminalNode {
	return s.GetToken(ADLWalkerType, 0)
}

func (s *TypeParamErrorContext) Newtype() antlr.TerminalNode {
	return s.GetToken(ADLWalkerNewtype, 0)
}

func (s *TypeParamErrorContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *TypeParamErrorContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeParamErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitTypeParamError(s)
	}
}

type NewtypeContext struct {
	*TldContext
}

func NewNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewtypeContext {
	var p = new(NewtypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *NewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewtypeContext) Newtype() antlr.TerminalNode {
	return s.GetToken(ADLWalkerNewtype, 0)
}

func (s *NewtypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *NewtypeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *NewtypeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *NewtypeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *NewtypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *NewtypeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *NewtypeContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *NewtypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *NewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterNewtype(s)
	}
}

func (s *NewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitNewtype(s)
	}
}

type ModAnnoContext struct {
	*TldContext
}

func NewModAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModAnnoContext {
	var p = new(ModAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *ModAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModAnnoContext) ModuleAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerModuleAnno, 0)
}

func (s *ModAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterModAnno(s)
	}
}

func (s *ModAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitModAnno(s)
	}
}

type FieldAnnoContext struct {
	*TldContext
}

func NewFieldAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnoContext {
	var p = new(FieldAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *FieldAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAnnoContext) FieldAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerFieldAnno, 0)
}

func (s *FieldAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterFieldAnno(s)
	}
}

func (s *FieldAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitFieldAnno(s)
	}
}

type UnionContext struct {
	*TldContext
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) Union() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUnion, 0)
}

func (s *UnionContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *UnionContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *UnionContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *UnionContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *UnionContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *UnionContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *UnionContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitUnion(s)
	}
}

type DeclAnnoContext struct {
	*TldContext
}

func NewDeclAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnoContext {
	var p = new(DeclAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *DeclAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAnnoContext) DeclAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDeclAnno, 0)
}

func (s *DeclAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterDeclAnno(s)
	}
}

func (s *DeclAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitDeclAnno(s)
	}
}

type StructContext struct {
	*TldContext
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) Struct() antlr.TerminalNode {
	return s.GetToken(ADLWalkerStruct, 0)
}

func (s *StructContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *StructContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *StructContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *StructContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *StructContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *StructContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *StructContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (p *ADLWalker) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADLWalkerRULE_tld)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(ADLWalkerStruct)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(46)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(47)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(52)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(53)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerField {
				{
					p.SetState(56)
					p.NameBody()
				}

				p.SetState(61)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(62)
				p.Match(ADLWalkerUP)
			}

		}

	case 2:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(ADLWalkerUnion)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(66)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(67)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(72)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(73)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerField {
				{
					p.SetState(76)
					p.NameBody()
				}

				p.SetState(81)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(82)
				p.Match(ADLWalkerUP)
			}

		}

	case 3:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Match(ADLWalkerType)
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(86)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(87)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(93)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(96)
					p.TypeExpr_()
				}

			}
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerJson {
				{
					p.SetState(99)
					p.JsonVal()
				}

				p.SetState(104)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(105)
				p.Match(ADLWalkerUP)
			}

		}

	case 4:
		localctx = NewNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.Match(ADLWalkerNewtype)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(109)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(110)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(115)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(116)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(119)
					p.TypeExpr_()
				}

			}
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerJson {
				{
					p.SetState(122)
					p.JsonVal()
				}

				p.SetState(127)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(128)
				p.Match(ADLWalkerUP)
			}

		}

	case 5:
		localctx = NewModAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Match(ADLWalkerModuleAnno)
		}

	case 6:
		localctx = NewDeclAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.Match(ADLWalkerDeclAnno)
		}

	case 7:
		localctx = NewFieldAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.Match(ADLWalkerFieldAnno)
		}

	case 8:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ADLWalkerStruct || _la == ADLWalkerUnion) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(135)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(136)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(142)
			p.Match(ADLWalkerERROR)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerField {
			{
				p.SetState(143)
				p.NameBody()
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(ADLWalkerUP)
		}

	case 9:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(150)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ADLWalkerNewtype || _la == ADLWalkerType) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(151)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(152)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(158)
			p.Match(ADLWalkerERROR)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerTypeExpr {
			{
				p.SetState(159)
				p.TypeExpr_()
			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerJson {
			{
				p.SetState(162)
				p.JsonVal()
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(168)
				p.Match(ADLWalkerUP)
			}

		}

	}

	return localctx
}

// INameBodyContext is an interface to support dynamic dispatch.
type INameBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameBodyContext differentiates from other interfaces.
	IsNameBodyContext()
}

type NameBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameBodyContext() *NameBodyContext {
	var p = new(NameBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_nameBody
	return p
}

func (*NameBodyContext) IsNameBodyContext() {}

func NewNameBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameBodyContext {
	var p = new(NameBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_nameBody

	return p
}

func (s *NameBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *NameBodyContext) CopyFrom(ctx *NameBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NameBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldContext struct {
	*NameBodyContext
}

func NewFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldContext {
	var p = new(FieldContext)

	p.NameBodyContext = NewEmptyNameBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameBodyContext))

	return p
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) Field() antlr.TerminalNode {
	return s.GetToken(ADLWalkerField, 0)
}

func (s *FieldContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *FieldContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *FieldContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *FieldContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *FieldContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *FieldContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ADLWalker) NameBody() (localctx INameBodyContext) {
	localctx = NewNameBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADLWalkerRULE_nameBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFieldContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(ADLWalkerField)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(174)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(175)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerTypeExpr {
			{
				p.SetState(181)
				p.TypeExpr_()
			}

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerJson {
			{
				p.SetState(184)
				p.JsonVal()
			}

		}
		{
			p.SetState(187)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// ITypeExpr_Context is an interface to support dynamic dispatch.
type ITypeExpr_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExpr_Context differentiates from other interfaces.
	IsTypeExpr_Context()
}

type TypeExpr_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExpr_Context() *TypeExpr_Context {
	var p = new(TypeExpr_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_typeExpr_
	return p
}

func (*TypeExpr_Context) IsTypeExpr_Context() {}

func NewTypeExpr_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExpr_Context {
	var p = new(TypeExpr_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_typeExpr_

	return p
}

func (s *TypeExpr_Context) GetParser() antlr.Parser { return s.parser }

func (s *TypeExpr_Context) TypeExpr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExpr, 0)
}

func (s *TypeExpr_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *TypeExpr_Context) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *TypeExpr_Context) AllTypeExprElem_() []ITypeExprElem_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeExpr_Context) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

func (s *TypeExpr_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpr_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExpr_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitTypeExpr_(s)
	}
}

func (p *ADLWalker) TypeExpr_() (localctx ITypeExpr_Context) {
	localctx = NewTypeExpr_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADLWalkerRULE_typeExpr_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(ADLWalkerTypeExpr)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(191)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLWalkerTypeExprElem {
			{
				p.SetState(192)
				p.TypeExprElem_()
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(197)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// ITypeExprElem_Context is an interface to support dynamic dispatch.
type ITypeExprElem_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExprElem_Context differentiates from other interfaces.
	IsTypeExprElem_Context()
}

type TypeExprElem_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElem_Context() *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_typeExprElem_
	return p
}

func (*TypeExprElem_Context) IsTypeExprElem_Context() {}

func NewTypeExprElem_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_typeExprElem_

	return p
}

func (s *TypeExprElem_Context) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprElem_Context) CopyFrom(ctx *TypeExprElem_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElem_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElem_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeParamsContext struct {
	*TypeExprElem_Context
}

func NewTypeParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamsContext {
	var p = new(TypeParamsContext)

	p.TypeExprElem_Context = NewEmptyTypeExprElem_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElem_Context))

	return p
}

func (s *TypeParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamsContext) TypeExprElem() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExprElem, 0)
}

func (s *TypeParamsContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *TypeParamsContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *TypeParamsContext) AllTypeExprElem_() []ITypeExprElem_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeParamsContext) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

func (s *TypeParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterTypeParams(s)
	}
}

func (s *TypeParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitTypeParams(s)
	}
}

func (p *ADLWalker) TypeExprElem_() (localctx ITypeExprElem_Context) {
	localctx = NewTypeExprElem_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ADLWalkerRULE_typeExprElem_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewTypeParamsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(ADLWalkerTypeExprElem)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(202)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLWalkerTypeExprElem {
			{
				p.SetState(203)
				p.TypeExprElem_()
			}

			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(208)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// IJsonValContext is an interface to support dynamic dispatch.
type IJsonValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonValContext differentiates from other interfaces.
	IsJsonValContext()
}

type JsonValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValContext() *JsonValContext {
	var p = new(JsonValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_jsonVal
	return p
}

func (*JsonValContext) IsJsonValContext() {}

func NewJsonValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValContext {
	var p = new(JsonValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_jsonVal

	return p
}

func (s *JsonValContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonValContext) CopyFrom(ctx *JsonValContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JsonStrContext struct {
	*JsonValContext
}

func NewJsonStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonStrContext {
	var p = new(JsonStrContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonStrContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonStrContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonStrContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonStr, 0)
}

func (s *JsonStrContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonStr(s)
	}
}

func (s *JsonStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonStr(s)
	}
}

type JsonArrayContext struct {
	*JsonValContext
}

func NewJsonArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayContext {
	var p = new(JsonArrayContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonArrayContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonArrayContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *JsonArrayContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *JsonArrayContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonArray, 0)
}

func (s *JsonArrayContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *JsonArrayContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *JsonArrayContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonArrayContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonArray(s)
	}
}

func (s *JsonArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonArray(s)
	}
}

type JsonFloatContext struct {
	*JsonValContext
}

func NewJsonFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonFloatContext {
	var p = new(JsonFloatContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonFloatContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonFloatContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonFloatContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonFloat, 0)
}

func (s *JsonFloatContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonFloat(s)
	}
}

func (s *JsonFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonFloat(s)
	}
}

type JsonObjContext struct {
	*JsonValContext
}

func NewJsonObjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjContext {
	var p = new(JsonObjContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonObjContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonObjContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *JsonObjContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *JsonObjContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonObj, 0)
}

func (s *JsonObjContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *JsonObjContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *JsonObjContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonObjContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonObj(s)
	}
}

func (s *JsonObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonObj(s)
	}
}

type JsonBoolContext struct {
	*JsonValContext
}

func NewJsonBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonBoolContext {
	var p = new(JsonBoolContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonBoolContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonBoolContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonBoolContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonBool, 0)
}

func (s *JsonBoolContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonBool(s)
	}
}

func (s *JsonBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonBool(s)
	}
}

type JsonIntContext struct {
	*JsonValContext
}

func NewJsonIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonIntContext {
	var p = new(JsonIntContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonIntContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonIntContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonIntContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonInt, 0)
}

func (s *JsonIntContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonInt(s)
	}
}

func (s *JsonIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonInt(s)
	}
}

type JsonNullContext struct {
	*JsonValContext
}

func NewJsonNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonNullContext {
	var p = new(JsonNullContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonNullContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonNullContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonNullContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonNull, 0)
}

func (s *JsonNullContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonNull(s)
	}
}

func (s *JsonNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonNull(s)
	}
}

func (p *ADLWalker) JsonVal() (localctx IJsonValContext) {
	localctx = NewJsonValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ADLWalkerRULE_jsonVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(213)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(214)
			p.Match(ADLWalkerJsonStr)
		}
		{
			p.SetState(215)
			p.Match(ADLWalkerUP)
		}

	case 2:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(217)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(218)
			p.Match(ADLWalkerJsonBool)
		}
		{
			p.SetState(219)
			p.Match(ADLWalkerUP)
		}

	case 3:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(221)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(222)
			p.Match(ADLWalkerJsonNull)
		}
		{
			p.SetState(223)
			p.Match(ADLWalkerUP)
		}

	case 4:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(224)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(225)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(226)
			p.Match(ADLWalkerJsonInt)
		}
		{
			p.SetState(227)
			p.Match(ADLWalkerUP)
		}

	case 5:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(228)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(229)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(230)
			p.Match(ADLWalkerJsonFloat)
		}
		{
			p.SetState(231)
			p.Match(ADLWalkerUP)
		}

	case 6:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(232)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(233)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(234)
			p.Match(ADLWalkerJsonArray)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(235)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == ADLWalkerJson {
				{
					p.SetState(236)
					p.JsonVal()
				}

				p.SetState(239)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(241)
				p.Match(ADLWalkerUP)
			}

		}
		{
			p.SetState(245)
			p.Match(ADLWalkerUP)
		}

	case 7:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(246)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(247)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(248)
			p.Match(ADLWalkerJsonObj)
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(249)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == ADLWalkerJson {
				{
					p.SetState(250)
					p.JsonVal()
				}

				p.SetState(253)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(255)
				p.Match(ADLWalkerUP)
			}

		}
		{
			p.SetState(259)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}
