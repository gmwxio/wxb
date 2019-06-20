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
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 7, 3, 30, 10, 3, 12, 3, 14, 3, 33, 11, 3, 3, 3, 7, 3,
	36, 10, 3, 12, 3, 14, 3, 39, 11, 3, 3, 3, 7, 3, 42, 10, 3, 12, 3, 14, 3,
	45, 11, 3, 3, 3, 5, 3, 48, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 53, 10, 4, 12,
	4, 14, 4, 56, 11, 4, 3, 4, 5, 4, 59, 10, 4, 3, 4, 7, 4, 62, 10, 4, 12,
	4, 14, 4, 65, 11, 4, 3, 4, 5, 4, 68, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 73,
	10, 4, 12, 4, 14, 4, 76, 11, 4, 3, 4, 5, 4, 79, 10, 4, 3, 4, 7, 4, 82,
	10, 4, 12, 4, 14, 4, 85, 11, 4, 3, 4, 5, 4, 88, 10, 4, 3, 4, 3, 4, 3, 4,
	7, 4, 93, 10, 4, 12, 4, 14, 4, 96, 11, 4, 3, 4, 5, 4, 99, 10, 4, 3, 4,
	5, 4, 102, 10, 4, 3, 4, 7, 4, 105, 10, 4, 12, 4, 14, 4, 108, 11, 4, 3,
	4, 5, 4, 111, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 116, 10, 4, 12, 4, 14, 4,
	119, 11, 4, 3, 4, 5, 4, 122, 10, 4, 3, 4, 5, 4, 125, 10, 4, 3, 4, 7, 4,
	128, 10, 4, 12, 4, 14, 4, 131, 11, 4, 3, 4, 5, 4, 134, 10, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 154, 10, 4, 12, 4, 14, 4, 157, 11, 4,
	3, 4, 3, 4, 7, 4, 161, 10, 4, 12, 4, 14, 4, 164, 11, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 170, 10, 4, 12, 4, 14, 4, 173, 11, 4, 3, 4, 3, 4, 5, 4,
	177, 10, 4, 3, 4, 7, 4, 180, 10, 4, 12, 4, 14, 4, 183, 11, 4, 3, 4, 5,
	4, 186, 10, 4, 5, 4, 188, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 193, 10, 5, 12,
	5, 14, 5, 196, 11, 5, 3, 5, 5, 5, 199, 10, 5, 3, 5, 5, 5, 202, 10, 5, 3,
	5, 5, 5, 205, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 212, 10, 6, 3,
	7, 3, 7, 3, 7, 6, 7, 217, 10, 7, 13, 7, 14, 7, 218, 3, 7, 3, 7, 5, 7, 223,
	10, 7, 3, 8, 3, 8, 3, 8, 6, 8, 228, 10, 8, 13, 8, 14, 8, 229, 3, 8, 3,
	8, 5, 8, 234, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6,
	9, 244, 10, 9, 13, 9, 14, 9, 245, 3, 9, 3, 9, 5, 9, 250, 10, 9, 3, 9, 3,
	9, 3, 9, 6, 9, 255, 10, 9, 13, 9, 14, 9, 256, 3, 9, 3, 9, 5, 9, 261, 10,
	9, 5, 9, 263, 10, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 4,
	3, 2, 34, 35, 3, 2, 36, 37, 2, 311, 2, 18, 3, 2, 2, 2, 4, 26, 3, 2, 2,
	2, 6, 187, 3, 2, 2, 2, 8, 189, 3, 2, 2, 2, 10, 206, 3, 2, 2, 2, 12, 213,
	3, 2, 2, 2, 14, 224, 3, 2, 2, 2, 16, 262, 3, 2, 2, 2, 18, 19, 7, 26, 2,
	2, 19, 20, 7, 30, 2, 2, 20, 21, 7, 26, 2, 2, 21, 22, 5, 4, 3, 2, 22, 23,
	7, 27, 2, 2, 23, 24, 7, 27, 2, 2, 24, 25, 7, 2, 2, 3, 25, 3, 3, 2, 2, 2,
	26, 47, 7, 31, 2, 2, 27, 31, 7, 26, 2, 2, 28, 30, 5, 10, 6, 2, 29, 28,
	3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2,
	32, 37, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 36, 7, 32, 2, 2, 35, 34, 3,
	2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38,
	43, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 42, 5, 6, 4, 2, 41, 40, 3, 2, 2,
	2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46,
	3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 48, 7, 27, 2, 2, 47, 27, 3, 2, 2, 2,
	47, 48, 3, 2, 2, 2, 48, 5, 3, 2, 2, 2, 49, 67, 7, 34, 2, 2, 50, 54, 7,
	26, 2, 2, 51, 53, 5, 10, 6, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2,
	54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3,
	2, 2, 2, 57, 59, 7, 38, 2, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59,
	63, 3, 2, 2, 2, 60, 62, 5, 8, 5, 2, 61, 60, 3, 2, 2, 2, 62, 65, 3, 2, 2,
	2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 66, 68, 7, 27, 2, 2, 67, 50, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 188, 3, 2, 2, 2, 69, 87, 7, 35, 2, 2, 70, 74, 7, 26, 2, 2, 71, 73,
	5, 10, 6, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	74, 75, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79, 7,
	38, 2, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 83, 3, 2, 2, 2, 80,
	82, 5, 8, 5, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2,
	2, 83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 88,
	7, 27, 2, 2, 87, 70, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 188, 3, 2, 2,
	2, 89, 110, 7, 37, 2, 2, 90, 94, 7, 26, 2, 2, 91, 93, 5, 10, 6, 2, 92,
	91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2,
	2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 99, 7, 38, 2, 2, 98, 97,
	3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 102, 5, 12, 7,
	2, 101, 100, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 106, 3, 2, 2, 2, 103,
	105, 5, 16, 9, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2,
	2, 2, 109, 111, 7, 27, 2, 2, 110, 90, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2,
	111, 188, 3, 2, 2, 2, 112, 133, 7, 36, 2, 2, 113, 117, 7, 26, 2, 2, 114,
	116, 5, 10, 6, 2, 115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115,
	3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2,
	2, 2, 120, 122, 7, 38, 2, 2, 121, 120, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2,
	122, 124, 3, 2, 2, 2, 123, 125, 5, 12, 7, 2, 124, 123, 3, 2, 2, 2, 124,
	125, 3, 2, 2, 2, 125, 129, 3, 2, 2, 2, 126, 128, 5, 16, 9, 2, 127, 126,
	3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2,
	2, 2, 130, 132, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 134, 7, 27, 2, 2,
	133, 113, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 188, 3, 2, 2, 2, 135,
	136, 7, 50, 2, 2, 136, 137, 7, 26, 2, 2, 137, 138, 5, 16, 9, 2, 138, 139,
	7, 27, 2, 2, 139, 188, 3, 2, 2, 2, 140, 141, 7, 51, 2, 2, 141, 142, 7,
	26, 2, 2, 142, 143, 5, 16, 9, 2, 143, 144, 7, 27, 2, 2, 144, 188, 3, 2,
	2, 2, 145, 146, 7, 52, 2, 2, 146, 147, 7, 26, 2, 2, 147, 148, 5, 16, 9,
	2, 148, 149, 7, 27, 2, 2, 149, 188, 3, 2, 2, 2, 150, 151, 9, 2, 2, 2, 151,
	155, 7, 26, 2, 2, 152, 154, 5, 10, 6, 2, 153, 152, 3, 2, 2, 2, 154, 157,
	3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2,
	2, 2, 157, 155, 3, 2, 2, 2, 158, 162, 7, 29, 2, 2, 159, 161, 5, 8, 5, 2,
	160, 159, 3, 2, 2, 2, 161, 164, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162,
	163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 165, 188,
	7, 27, 2, 2, 166, 167, 9, 3, 2, 2, 167, 171, 7, 26, 2, 2, 168, 170, 5,
	10, 6, 2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171, 169, 3, 2, 2,
	2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 174,
	176, 7, 29, 2, 2, 175, 177, 5, 12, 7, 2, 176, 175, 3, 2, 2, 2, 176, 177,
	3, 2, 2, 2, 177, 181, 3, 2, 2, 2, 178, 180, 5, 16, 9, 2, 179, 178, 3, 2,
	2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2,
	182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 186, 7, 27, 2, 2, 185,
	184, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187, 49, 3,
	2, 2, 2, 187, 69, 3, 2, 2, 2, 187, 89, 3, 2, 2, 2, 187, 112, 3, 2, 2, 2,
	187, 135, 3, 2, 2, 2, 187, 140, 3, 2, 2, 2, 187, 145, 3, 2, 2, 2, 187,
	150, 3, 2, 2, 2, 187, 166, 3, 2, 2, 2, 188, 7, 3, 2, 2, 2, 189, 204, 7,
	41, 2, 2, 190, 194, 7, 26, 2, 2, 191, 193, 5, 10, 6, 2, 192, 191, 3, 2,
	2, 2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2,
	195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 199, 5, 12, 7, 2, 198,
	197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201, 3, 2, 2, 2, 200, 202,
	5, 16, 9, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 205, 7, 27, 2, 2, 204, 190, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2,
	205, 9, 3, 2, 2, 2, 206, 211, 7, 33, 2, 2, 207, 208, 7, 26, 2, 2, 208,
	209, 5, 16, 9, 2, 209, 210, 7, 27, 2, 2, 210, 212, 3, 2, 2, 2, 211, 207,
	3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 11, 3, 2, 2, 2, 213, 222, 7, 39,
	2, 2, 214, 216, 7, 26, 2, 2, 215, 217, 5, 14, 8, 2, 216, 215, 3, 2, 2,
	2, 217, 218, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219,
	220, 3, 2, 2, 2, 220, 221, 7, 27, 2, 2, 221, 223, 3, 2, 2, 2, 222, 214,
	3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 13, 3, 2, 2, 2, 224, 233, 7, 40,
	2, 2, 225, 227, 7, 26, 2, 2, 226, 228, 5, 14, 8, 2, 227, 226, 3, 2, 2,
	2, 228, 229, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 232, 7, 27, 2, 2, 232, 234, 3, 2, 2, 2, 233, 225,
	3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 15, 3, 2, 2, 2, 235, 263, 7, 43,
	2, 2, 236, 263, 7, 44, 2, 2, 237, 263, 7, 45, 2, 2, 238, 263, 7, 46, 2,
	2, 239, 263, 7, 47, 2, 2, 240, 249, 7, 48, 2, 2, 241, 243, 7, 26, 2, 2,
	242, 244, 5, 16, 9, 2, 243, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245,
	243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248,
	7, 27, 2, 2, 248, 250, 3, 2, 2, 2, 249, 241, 3, 2, 2, 2, 249, 250, 3, 2,
	2, 2, 250, 263, 3, 2, 2, 2, 251, 260, 7, 49, 2, 2, 252, 254, 7, 26, 2,
	2, 253, 255, 5, 16, 9, 2, 254, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256,
	254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259,
	7, 27, 2, 2, 259, 261, 3, 2, 2, 2, 260, 252, 3, 2, 2, 2, 260, 261, 3, 2,
	2, 2, 261, 263, 3, 2, 2, 2, 262, 235, 3, 2, 2, 2, 262, 236, 3, 2, 2, 2,
	262, 237, 3, 2, 2, 2, 262, 238, 3, 2, 2, 2, 262, 239, 3, 2, 2, 2, 262,
	240, 3, 2, 2, 2, 262, 251, 3, 2, 2, 2, 263, 17, 3, 2, 2, 2, 45, 31, 37,
	43, 47, 54, 58, 63, 67, 74, 78, 83, 87, 94, 98, 101, 106, 110, 117, 121,
	124, 129, 133, 155, 162, 171, 176, 181, 185, 187, 194, 198, 201, 204, 211,
	218, 222, 229, 233, 245, 249, 256, 260, 262,
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
	"adl", "module", "tld", "nameBody", "annotation", "typeExpr_", "typeExprElem_",
	"jsonVal",
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
	ADLWalkerRULE_annotation    = 4
	ADLWalkerRULE_typeExpr_     = 5
	ADLWalkerRULE_typeExprElem_ = 6
	ADLWalkerRULE_jsonVal       = 7
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
		p.SetState(16)
		p.Match(ADLWalkerDOWN)
	}
	{
		p.SetState(17)
		p.Match(ADLWalkerADL)
	}
	{
		p.SetState(18)
		p.Match(ADLWalkerDOWN)
	}
	{
		p.SetState(19)
		p.Module()
	}
	{
		p.SetState(20)
		p.Match(ADLWalkerUP)
	}
	{
		p.SetState(21)
		p.Match(ADLWalkerUP)
	}
	{
		p.SetState(22)
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

func (s *ModuleContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ModuleContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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
		p.SetState(24)
		p.Match(ADLWalkerModule)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(25)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(26)
				p.Annotation()
			}

			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerImport {
			{
				p.SetState(32)
				p.Match(ADLWalkerImport)
			}

			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ADLWalkerStruct-32))|(1<<(ADLWalkerUnion-32))|(1<<(ADLWalkerNewtype-32))|(1<<(ADLWalkerType-32))|(1<<(ADLWalkerModuleAnno-32))|(1<<(ADLWalkerDeclAnno-32))|(1<<(ADLWalkerFieldAnno-32)))) != 0 {
			{
				p.SetState(38)
				p.Tld()
			}

			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(44)
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

func (s *TypeContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TypeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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

func (s *TypeParamErrorContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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

func (s *NewtypeContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NewtypeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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

func (s *ModAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *ModAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *ModAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
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

func (s *FieldAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *FieldAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *FieldAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
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

func (s *UnionContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *UnionContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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

func (s *DeclAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *DeclAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *DeclAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
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

func (s *StructContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StructContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(ADLWalkerStruct)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(48)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(49)
					p.Annotation()
				}

				p.SetState(54)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(55)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerField {
				{
					p.SetState(58)
					p.NameBody()
				}

				p.SetState(63)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(64)
				p.Match(ADLWalkerUP)
			}

		}

	case 2:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(ADLWalkerUnion)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(68)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(69)
					p.Annotation()
				}

				p.SetState(74)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(75)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerField {
				{
					p.SetState(78)
					p.NameBody()
				}

				p.SetState(83)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(84)
				p.Match(ADLWalkerUP)
			}

		}

	case 3:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Match(ADLWalkerType)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(88)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(89)
					p.Annotation()
				}

				p.SetState(94)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(95)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(98)
					p.TypeExpr_()
				}

			}
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ADLWalkerJsonStr-41))|(1<<(ADLWalkerJsonBool-41))|(1<<(ADLWalkerJsonNull-41))|(1<<(ADLWalkerJsonInt-41))|(1<<(ADLWalkerJsonFloat-41))|(1<<(ADLWalkerJsonArray-41))|(1<<(ADLWalkerJsonObj-41)))) != 0 {
				{
					p.SetState(101)
					p.JsonVal()
				}

				p.SetState(106)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(107)
				p.Match(ADLWalkerUP)
			}

		}

	case 4:
		localctx = NewNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Match(ADLWalkerNewtype)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(111)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(112)
					p.Annotation()
				}

				p.SetState(117)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(118)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(121)
					p.TypeExpr_()
				}

			}
			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ADLWalkerJsonStr-41))|(1<<(ADLWalkerJsonBool-41))|(1<<(ADLWalkerJsonNull-41))|(1<<(ADLWalkerJsonInt-41))|(1<<(ADLWalkerJsonFloat-41))|(1<<(ADLWalkerJsonArray-41))|(1<<(ADLWalkerJsonObj-41)))) != 0 {
				{
					p.SetState(124)
					p.JsonVal()
				}

				p.SetState(129)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(130)
				p.Match(ADLWalkerUP)
			}

		}

	case 5:
		localctx = NewModAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(133)
			p.Match(ADLWalkerModuleAnno)
		}
		{
			p.SetState(134)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(135)
			p.JsonVal()
		}
		{
			p.SetState(136)
			p.Match(ADLWalkerUP)
		}

	case 6:
		localctx = NewDeclAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(138)
			p.Match(ADLWalkerDeclAnno)
		}
		{
			p.SetState(139)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(140)
			p.JsonVal()
		}
		{
			p.SetState(141)
			p.Match(ADLWalkerUP)
		}

	case 7:
		localctx = NewFieldAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(143)
			p.Match(ADLWalkerFieldAnno)
		}
		{
			p.SetState(144)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(145)
			p.JsonVal()
		}
		{
			p.SetState(146)
			p.Match(ADLWalkerUP)
		}

	case 8:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(148)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ADLWalkerStruct || _la == ADLWalkerUnion) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(149)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(150)
				p.Annotation()
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(156)
			p.Match(ADLWalkerERROR)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerField {
			{
				p.SetState(157)
				p.NameBody()
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(163)
			p.Match(ADLWalkerUP)
		}

	case 9:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(164)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ADLWalkerNewtype || _la == ADLWalkerType) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(165)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(166)
				p.Annotation()
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(172)
			p.Match(ADLWalkerERROR)
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerTypeExpr {
			{
				p.SetState(173)
				p.TypeExpr_()
			}

		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ADLWalkerJsonStr-41))|(1<<(ADLWalkerJsonBool-41))|(1<<(ADLWalkerJsonNull-41))|(1<<(ADLWalkerJsonInt-41))|(1<<(ADLWalkerJsonFloat-41))|(1<<(ADLWalkerJsonArray-41))|(1<<(ADLWalkerJsonObj-41)))) != 0 {
			{
				p.SetState(176)
				p.JsonVal()
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(182)
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

func (s *FieldContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FieldContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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
		p.SetState(187)
		p.Match(ADLWalkerField)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(188)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(189)
				p.Annotation()
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerTypeExpr {
			{
				p.SetState(195)
				p.TypeExpr_()
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ADLWalkerJsonStr-41))|(1<<(ADLWalkerJsonBool-41))|(1<<(ADLWalkerJsonNull-41))|(1<<(ADLWalkerJsonInt-41))|(1<<(ADLWalkerJsonFloat-41))|(1<<(ADLWalkerJsonArray-41))|(1<<(ADLWalkerJsonObj-41)))) != 0 {
			{
				p.SetState(198)
				p.JsonVal()
			}

		}
		{
			p.SetState(201)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) Annotation() antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, 0)
}

func (s *AnnotationContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *AnnotationContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *AnnotationContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *ADLWalker) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADLWalkerRULE_annotation)
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
		p.SetState(204)
		p.Match(ADLWalkerAnnotation)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(205)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(206)
			p.JsonVal()
		}
		{
			p.SetState(207)
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
	p.EnterRule(localctx, 10, ADLWalkerRULE_typeExpr_)
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
		p.SetState(211)
		p.Match(ADLWalkerTypeExpr)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(212)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLWalkerTypeExprElem {
			{
				p.SetState(213)
				p.TypeExprElem_()
			}

			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(218)
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
	p.EnterRule(localctx, 12, ADLWalkerRULE_typeExprElem_)
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
		p.SetState(222)
		p.Match(ADLWalkerTypeExprElem)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(223)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLWalkerTypeExprElem {
			{
				p.SetState(224)
				p.TypeExprElem_()
			}

			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(229)
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

func (s *JsonStrContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonStr, 0)
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

func (s *JsonArrayContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonArray, 0)
}

func (s *JsonArrayContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonArrayContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
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

func (s *JsonFloatContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonFloat, 0)
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

func (s *JsonObjContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonObj, 0)
}

func (s *JsonObjContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonObjContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
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

func (s *JsonBoolContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonBool, 0)
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

func (s *JsonIntContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonInt, 0)
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

func (s *JsonNullContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonNull, 0)
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
	p.EnterRule(localctx, 14, ADLWalkerRULE_jsonVal)
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

	switch p.GetTokenStream().LA(1) {
	case ADLWalkerJsonStr:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Match(ADLWalkerJsonStr)
		}

	case ADLWalkerJsonBool:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(ADLWalkerJsonBool)
		}

	case ADLWalkerJsonNull:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)
			p.Match(ADLWalkerJsonNull)
		}

	case ADLWalkerJsonInt:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(236)
			p.Match(ADLWalkerJsonInt)
		}

	case ADLWalkerJsonFloat:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(237)
			p.Match(ADLWalkerJsonFloat)
		}

	case ADLWalkerJsonArray:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(238)
			p.Match(ADLWalkerJsonArray)
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(239)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ADLWalkerJsonStr-41))|(1<<(ADLWalkerJsonBool-41))|(1<<(ADLWalkerJsonNull-41))|(1<<(ADLWalkerJsonInt-41))|(1<<(ADLWalkerJsonFloat-41))|(1<<(ADLWalkerJsonArray-41))|(1<<(ADLWalkerJsonObj-41)))) != 0) {
				{
					p.SetState(240)
					p.JsonVal()
				}

				p.SetState(243)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(245)
				p.Match(ADLWalkerUP)
			}

		}

	case ADLWalkerJsonObj:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(249)
			p.Match(ADLWalkerJsonObj)
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(250)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ADLWalkerJsonStr-41))|(1<<(ADLWalkerJsonBool-41))|(1<<(ADLWalkerJsonNull-41))|(1<<(ADLWalkerJsonInt-41))|(1<<(ADLWalkerJsonFloat-41))|(1<<(ADLWalkerJsonArray-41))|(1<<(ADLWalkerJsonObj-41)))) != 0) {
				{
					p.SetState(251)
					p.JsonVal()
				}

				p.SetState(254)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(256)
				p.Match(ADLWalkerUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
