// Generated from ADLParser.g4 by ANTLR 4.7.

package parser // ADLParser

import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/wxio/goantlr"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 278,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 3, 3, 3, 7, 3,
	45, 10, 3, 12, 3, 14, 3, 48, 11, 3, 3, 3, 7, 3, 51, 10, 3, 12, 3, 14, 3,
	54, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 63, 10, 4, 12,
	4, 14, 4, 66, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 73, 10, 4, 12,
	4, 14, 4, 76, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 81, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 87, 10, 5, 3, 6, 7, 6, 90, 10, 6, 12, 6, 14, 6, 93, 11,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 98, 10, 6, 3, 6, 3, 6, 7, 6, 102, 10, 6, 12,
	6, 14, 6, 105, 11, 6, 3, 6, 3, 6, 3, 6, 7, 6, 110, 10, 6, 12, 6, 14, 6,
	113, 11, 6, 3, 6, 3, 6, 3, 6, 5, 6, 118, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6,
	123, 10, 6, 3, 6, 3, 6, 5, 6, 127, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 149, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 155,
	10, 7, 12, 7, 14, 7, 158, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 165,
	10, 7, 12, 7, 14, 7, 168, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6,
	7, 176, 10, 7, 13, 7, 14, 7, 177, 3, 7, 3, 7, 5, 7, 182, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 188, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 193, 10, 8, 7,
	8, 195, 10, 8, 12, 8, 14, 8, 198, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 7, 9, 206, 10, 9, 12, 9, 14, 9, 209, 11, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 7, 10, 218, 10, 10, 12, 10, 14, 10, 221, 11, 10,
	3, 10, 3, 10, 5, 10, 225, 10, 10, 3, 11, 7, 11, 228, 10, 11, 12, 11, 14,
	11, 231, 11, 11, 3, 11, 3, 11, 5, 11, 235, 10, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 240, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 7, 12, 252, 10, 12, 12, 12, 14, 12, 255, 11, 12, 5, 12,
	257, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 268, 10, 12, 12, 12, 14, 12, 271, 11, 12, 5, 12, 273, 10, 12,
	3, 12, 5, 12, 276, 10, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 2, 2, 2, 308, 2, 24, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 80,
	3, 2, 2, 2, 8, 86, 3, 2, 2, 2, 10, 148, 3, 2, 2, 2, 12, 181, 3, 2, 2, 2,
	14, 183, 3, 2, 2, 2, 16, 201, 3, 2, 2, 2, 18, 212, 3, 2, 2, 2, 20, 229,
	3, 2, 2, 2, 22, 275, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26, 7, 2, 2, 3,
	26, 3, 3, 2, 2, 2, 27, 29, 5, 8, 5, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2,
	2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 33, 34, 7, 20, 2, 2, 34, 39, 7, 20, 2, 2, 35, 36, 7, 13, 2,
	2, 36, 38, 7, 20, 2, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	42, 46, 7, 3, 2, 2, 43, 45, 5, 6, 4, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3,
	2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 52, 3, 2, 2, 2, 48,
	46, 3, 2, 2, 2, 49, 51, 5, 10, 6, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2,
	2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 55, 56, 7, 4, 2, 2, 56, 57, 7, 10, 2, 2, 57, 5, 3, 2, 2, 2,
	58, 59, 7, 20, 2, 2, 59, 64, 7, 20, 2, 2, 60, 61, 7, 13, 2, 2, 61, 63,
	7, 20, 2, 2, 62, 60, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 81, 7,
	10, 2, 2, 68, 69, 7, 20, 2, 2, 69, 74, 7, 20, 2, 2, 70, 71, 7, 13, 2, 2,
	71, 73, 7, 20, 2, 2, 72, 70, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3,
	2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77,
	78, 7, 13, 2, 2, 78, 79, 7, 17, 2, 2, 79, 81, 7, 10, 2, 2, 80, 58, 3, 2,
	2, 2, 80, 68, 3, 2, 2, 2, 81, 7, 3, 2, 2, 2, 82, 83, 7, 18, 2, 2, 83, 84,
	7, 20, 2, 2, 84, 87, 5, 22, 12, 2, 85, 87, 7, 24, 2, 2, 86, 82, 3, 2, 2,
	2, 86, 85, 3, 2, 2, 2, 87, 9, 3, 2, 2, 2, 88, 90, 5, 8, 5, 2, 89, 88, 3,
	2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92,
	94, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 7, 20, 2, 2, 95, 97, 7, 20,
	2, 2, 96, 98, 5, 12, 7, 2, 97, 96, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98,
	99, 3, 2, 2, 2, 99, 103, 7, 3, 2, 2, 100, 102, 5, 20, 11, 2, 101, 100,
	3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2,
	2, 2, 104, 106, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 107, 7, 4, 2, 2,
	107, 149, 7, 10, 2, 2, 108, 110, 5, 8, 5, 2, 109, 108, 3, 2, 2, 2, 110,
	113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114,
	3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 115, 7, 20, 2, 2, 115, 117, 7, 20,
	2, 2, 116, 118, 5, 12, 7, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 120, 7, 7, 2, 2, 120, 122, 7, 20, 2, 2, 121,
	123, 5, 16, 9, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 126,
	3, 2, 2, 2, 124, 125, 7, 7, 2, 2, 125, 127, 5, 22, 12, 2, 126, 124, 3,
	2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 149, 7, 10, 2,
	2, 129, 130, 7, 20, 2, 2, 130, 131, 7, 20, 2, 2, 131, 132, 5, 22, 12, 2,
	132, 133, 7, 10, 2, 2, 133, 149, 3, 2, 2, 2, 134, 135, 7, 20, 2, 2, 135,
	136, 7, 20, 2, 2, 136, 137, 7, 20, 2, 2, 137, 138, 5, 22, 12, 2, 138, 139,
	7, 10, 2, 2, 139, 149, 3, 2, 2, 2, 140, 141, 7, 20, 2, 2, 141, 142, 7,
	20, 2, 2, 142, 143, 7, 11, 2, 2, 143, 144, 7, 20, 2, 2, 144, 145, 7, 20,
	2, 2, 145, 146, 5, 22, 12, 2, 146, 147, 7, 10, 2, 2, 147, 149, 3, 2, 2,
	2, 148, 91, 3, 2, 2, 2, 148, 111, 3, 2, 2, 2, 148, 129, 3, 2, 2, 2, 148,
	134, 3, 2, 2, 2, 148, 140, 3, 2, 2, 2, 149, 11, 3, 2, 2, 2, 150, 151, 7,
	15, 2, 2, 151, 156, 7, 20, 2, 2, 152, 153, 7, 14, 2, 2, 153, 155, 7, 20,
	2, 2, 154, 152, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2,
	156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159,
	182, 7, 16, 2, 2, 160, 161, 7, 15, 2, 2, 161, 166, 5, 14, 8, 2, 162, 163,
	7, 14, 2, 2, 163, 165, 7, 20, 2, 2, 164, 162, 3, 2, 2, 2, 165, 168, 3,
	2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 169, 3, 2, 2,
	2, 168, 166, 3, 2, 2, 2, 169, 170, 7, 16, 2, 2, 170, 182, 3, 2, 2, 2, 171,
	172, 7, 15, 2, 2, 172, 175, 7, 20, 2, 2, 173, 174, 7, 14, 2, 2, 174, 176,
	5, 14, 8, 2, 175, 173, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 175, 3, 2,
	2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 7, 16, 2, 2,
	180, 182, 3, 2, 2, 2, 181, 150, 3, 2, 2, 2, 181, 160, 3, 2, 2, 2, 181,
	171, 3, 2, 2, 2, 182, 13, 3, 2, 2, 2, 183, 184, 7, 20, 2, 2, 184, 187,
	7, 15, 2, 2, 185, 188, 7, 20, 2, 2, 186, 188, 5, 14, 8, 2, 187, 185, 3,
	2, 2, 2, 187, 186, 3, 2, 2, 2, 188, 196, 3, 2, 2, 2, 189, 192, 7, 14, 2,
	2, 190, 193, 7, 20, 2, 2, 191, 193, 5, 14, 8, 2, 192, 190, 3, 2, 2, 2,
	192, 191, 3, 2, 2, 2, 193, 195, 3, 2, 2, 2, 194, 189, 3, 2, 2, 2, 195,
	198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 199,
	3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199, 200, 7, 16, 2, 2, 200, 15, 3, 2,
	2, 2, 201, 202, 7, 15, 2, 2, 202, 207, 5, 18, 10, 2, 203, 204, 7, 14, 2,
	2, 204, 206, 5, 18, 10, 2, 205, 203, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2,
	207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209,
	207, 3, 2, 2, 2, 210, 211, 7, 16, 2, 2, 211, 17, 3, 2, 2, 2, 212, 224,
	7, 20, 2, 2, 213, 214, 7, 15, 2, 2, 214, 219, 5, 18, 10, 2, 215, 216, 7,
	14, 2, 2, 216, 218, 5, 18, 10, 2, 217, 215, 3, 2, 2, 2, 218, 221, 3, 2,
	2, 2, 219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 222, 3, 2, 2, 2,
	221, 219, 3, 2, 2, 2, 222, 223, 7, 16, 2, 2, 223, 225, 3, 2, 2, 2, 224,
	213, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 19, 3, 2, 2, 2, 226, 228, 5,
	8, 5, 2, 227, 226, 3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2, 2,
	2, 229, 230, 3, 2, 2, 2, 230, 232, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232,
	234, 7, 20, 2, 2, 233, 235, 5, 16, 9, 2, 234, 233, 3, 2, 2, 2, 234, 235,
	3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 239, 7, 20, 2, 2, 237, 238, 7, 7,
	2, 2, 238, 240, 5, 22, 12, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2,
	2, 240, 241, 3, 2, 2, 2, 241, 242, 7, 10, 2, 2, 242, 21, 3, 2, 2, 2, 243,
	276, 7, 19, 2, 2, 244, 276, 7, 20, 2, 2, 245, 276, 7, 21, 2, 2, 246, 276,
	7, 22, 2, 2, 247, 256, 7, 5, 2, 2, 248, 253, 5, 22, 12, 2, 249, 250, 7,
	14, 2, 2, 250, 252, 5, 22, 12, 2, 251, 249, 3, 2, 2, 2, 252, 255, 3, 2,
	2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2,
	255, 253, 3, 2, 2, 2, 256, 248, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257,
	258, 3, 2, 2, 2, 258, 276, 7, 6, 2, 2, 259, 272, 7, 3, 2, 2, 260, 261,
	7, 19, 2, 2, 261, 262, 7, 12, 2, 2, 262, 269, 5, 22, 12, 2, 263, 264, 7,
	14, 2, 2, 264, 265, 7, 19, 2, 2, 265, 266, 7, 12, 2, 2, 266, 268, 5, 22,
	12, 2, 267, 263, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2,
	269, 270, 3, 2, 2, 2, 270, 273, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272,
	260, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 276,
	7, 4, 2, 2, 275, 243, 3, 2, 2, 2, 275, 244, 3, 2, 2, 2, 275, 245, 3, 2,
	2, 2, 275, 246, 3, 2, 2, 2, 275, 247, 3, 2, 2, 2, 275, 259, 3, 2, 2, 2,
	276, 23, 3, 2, 2, 2, 36, 30, 39, 46, 52, 64, 74, 80, 86, 91, 97, 103, 111,
	117, 122, 126, 148, 156, 166, 177, 181, 187, 192, 196, 207, 219, 224, 229,
	234, 239, 253, 256, 269, 272, 275,
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
	"Module", "ImportModule", "ImportScopedName", "Annotation", "AnnotationNotScoped",
	"AnnotationScoped", "Struct", "Union", "Newtype", "Type", "TypeParam",
	"TypeExpr", "TypeExprElem", "Field", "Json", "JsonStr", "JsonBool", "JsonNull",
	"JsonInt", "JsonFloat", "JsonArray", "JsonObj", "ModuleAnno", "DeclAnno",
	"FieldAnno",
}

var ruleNames = []string{
	"adl", "module", "imports", "annon", "top_level_statement", "typeParam",
	"typeParamError", "typeExpr", "typeExprElem", "soruBody", "jsonValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ADLParser struct {
	*antlr.BaseParser
}

func NewADLParser(input antlr.TokenStream) *ADLParser {
	this := new(ADLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ADLParser.g4"

	return this
}

// ADLParser tokens.
const (
	ADLParserEOF                 = antlr.TokenEOF
	ADLParserLCUR                = 1
	ADLParserRCUR                = 2
	ADLParserLSQ                 = 3
	ADLParserRSQ                 = 4
	ADLParserEQ                  = 5
	ADLParserDQ                  = 6
	ADLParserSQ                  = 7
	ADLParserSEMI                = 8
	ADLParserDCOLON              = 9
	ADLParserCOLON               = 10
	ADLParserDOT                 = 11
	ADLParserCOMMA               = 12
	ADLParserLCHEVR              = 13
	ADLParserRCHEVR              = 14
	ADLParserSTAR                = 15
	ADLParserAT                  = 16
	ADLParserSTR                 = 17
	ADLParserID                  = 18
	ADLParserINT                 = 19
	ADLParserFLT                 = 20
	ADLParserWS                  = 21
	ADLParserLINE_DOC            = 22
	ADLParserLINE_COMMENT        = 23
	ADLParserDOWN                = 24
	ADLParserUP                  = 25
	ADLParserROOT                = 26
	ADLParserERROR               = 27
	ADLParserADL                 = 28
	ADLParserModule              = 29
	ADLParserImportModule        = 30
	ADLParserImportScopedName    = 31
	ADLParserAnnotation          = 32
	ADLParserAnnotationNotScoped = 33
	ADLParserAnnotationScoped    = 34
	ADLParserStruct              = 35
	ADLParserUnion               = 36
	ADLParserNewtype             = 37
	ADLParserType                = 38
	ADLParserTypeParam           = 39
	ADLParserTypeExpr            = 40
	ADLParserTypeExprElem        = 41
	ADLParserField               = 42
	ADLParserJson                = 43
	ADLParserJsonStr             = 44
	ADLParserJsonBool            = 45
	ADLParserJsonNull            = 46
	ADLParserJsonInt             = 47
	ADLParserJsonFloat           = 48
	ADLParserJsonArray           = 49
	ADLParserJsonObj             = 50
	ADLParserModuleAnno          = 51
	ADLParserDeclAnno            = 52
	ADLParserFieldAnno           = 53
)

// ADLParser rules.
const (
	ADLParserRULE_adl                 = 0
	ADLParserRULE_module              = 1
	ADLParserRULE_imports             = 2
	ADLParserRULE_annon               = 3
	ADLParserRULE_top_level_statement = 4
	ADLParserRULE_typeParam           = 5
	ADLParserRULE_typeParamError      = 6
	ADLParserRULE_typeExpr            = 7
	ADLParserRULE_typeExprElem        = 8
	ADLParserRULE_soruBody            = 9
	ADLParserRULE_jsonValue           = 10
)

type IAdlContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	Module() IModuleContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	EOF() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsAdlContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AdlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdlContext() *AdlContext {
	var p = new(AdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLParserEOF, 0)
}

//provideCopyFrom
func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdlEntryListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdlExitListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (s *AdlContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Adl != nil {
		h.Adl(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *AdlContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case AdlContextVisitor:
		return t.VisitAdl(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *ADLParser) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ADLParserRULE_adl)

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
		p.Module()
	}
	{
		p.SetState(23)
		p.Match(ADLParserEOF)
	}

	return localctx
}

type IModuleContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsModuleContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *ModuleContext) CopyFrom(ctx *ModuleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type ModuleStatementContext struct {
	*ModuleContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	name []antlr.Token
}

func NewModuleStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleStatementContext {
	var p = new(ModuleStatementContext)

	p.ModuleContext = NewEmptyModuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ModuleContext))

	return p
}

type IModuleStatementContext interface {
	//Current rule
	IModuleContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	AllImports() []IImportsContext
	AllTop_level_statement() []ITop_level_statementContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext
	Imports(i int) IImportsContext
	Top_level_statement(i int) ITop_level_statementContext

	//  tokenGetterDecl
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	GetName() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ModuleStatementContext) IsModuleStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *ModuleStatementContext) GetKw() antlr.Token  { return s.kw }
func (s *ModuleStatementContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ModuleStatementContext) Get_ID() antlr.Token  { return s._ID }
func (s *ModuleStatementContext) Set_ID(v antlr.Token) { s._ID = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ModuleStatementContext) GetName() []antlr.Token  { return s.name }
func (s *ModuleStatementContext) SetName(v []antlr.Token) { s.name = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ModuleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ModuleStatementContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *ModuleStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *ModuleStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *ModuleStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ModuleStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ModuleStatementContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *ModuleStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ADLParserDOT)
}

func (s *ModuleStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserDOT, i)
}

func (s *ModuleStatementContext) AllImports() []IImportsContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ImportsContext)(nil)).Elem())
	var tst = make([]IImportsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportsContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Imports(i int) IImportsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ImportsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *ModuleStatementContext) AllTop_level_statement() []ITop_level_statementContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*Top_level_statementContext)(nil)).Elem())
	var tst = make([]ITop_level_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITop_level_statementContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Top_level_statement(i int) ITop_level_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*Top_level_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITop_level_statementContext)
}

func (s *ModuleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleStatementEntryListener); ok {
		listenerT.EnterModuleStatement(s)
	}
}

func (s *ModuleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleStatementExitListener); ok {
		listenerT.ExitModuleStatement(s)
	}
}

func (s *ModuleStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ModuleStatement != nil {
		h.ModuleStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModuleStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModuleStatementContextVisitor:
		return t.VisitModuleStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ADLParserRULE_module)
	var //TokenTypeDecl
	_la int

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

	var _alt int

	localctx = NewModuleStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserAT || _la == ADLParserLINE_DOC {
		{
			p.SetState(25)
			p.Annon()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		var _m = p.Match(ADLParserID)
		localctx.(*ModuleStatementContext).kw = _m

	}
	{
		p.SetState(32)
		var _m = p.Match(ADLParserID)
		localctx.(*ModuleStatementContext)._ID = _m

	}
	localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserDOT {
		{
			p.SetState(33)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(34)
			var _m = p.Match(ADLParserID)
			localctx.(*ModuleStatementContext)._ID = _m

		}
		localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(ADLParserLCUR)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.Imports()
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserAT)|(1<<ADLParserID)|(1<<ADLParserLINE_DOC))) != 0 {
		{
			p.SetState(47)
			p.Top_level_statement()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(ADLParserRCUR)
	}
	{
		p.SetState(54)
		p.Match(ADLParserSEMI)
	}

	return localctx
}

type IImportsContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsImportsContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type ImportsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportsContext() *ImportsContext {
	var p = new(ImportsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_imports
	return p
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	var p = new(ImportsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *ImportsContext) CopyFrom(ctx *ImportsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type ImportModuleNameContext struct {
	*ImportsContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	a []antlr.Token
}

func NewImportModuleNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportModuleNameContext {
	var p = new(ImportModuleNameContext)

	p.ImportsContext = NewEmptyImportsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportsContext))

	return p
}

type IImportModuleNameContext interface {
	//Current rule
	IImportsContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	STAR() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllDOT() []antlr.TerminalNode
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	DOT(i int) antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	GetA() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ImportModuleNameContext) IsImportModuleNameContext() {}

//AltLabelStructDecl tokenDecls
func (s *ImportModuleNameContext) GetKw() antlr.Token  { return s.kw }
func (s *ImportModuleNameContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ImportModuleNameContext) Get_ID() antlr.Token  { return s._ID }
func (s *ImportModuleNameContext) Set_ID(v antlr.Token) { s._ID = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ImportModuleNameContext) GetA() []antlr.Token  { return s.a }
func (s *ImportModuleNameContext) SetA(v []antlr.Token) { s.a = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportModuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportModuleNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ADLParserDOT)
}

func (s *ImportModuleNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserDOT, i)
}

func (s *ImportModuleNameContext) STAR() antlr.TerminalNode {
	return s.GetToken(ADLParserSTAR, 0)
}

func (s *ImportModuleNameContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *ImportModuleNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ImportModuleNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ImportModuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportModuleNameEntryListener); ok {
		listenerT.EnterImportModuleName(s)
	}
}

func (s *ImportModuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportModuleNameExitListener); ok {
		listenerT.ExitImportModuleName(s)
	}
}

func (s *ImportModuleNameContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportModuleName != nil {
		h.ImportModuleName(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportModuleNameContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportModuleNameContextVisitor:
		return t.VisitImportModuleName(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ImportScopedNameContext struct {
	*ImportsContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	a []antlr.Token
}

func NewImportScopedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportScopedNameContext {
	var p = new(ImportScopedNameContext)

	p.ImportsContext = NewEmptyImportsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportsContext))

	return p
}

type IImportScopedNameContext interface {
	//Current rule
	IImportsContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	GetA() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ImportScopedNameContext) IsImportScopedNameContext() {}

//AltLabelStructDecl tokenDecls
func (s *ImportScopedNameContext) GetKw() antlr.Token  { return s.kw }
func (s *ImportScopedNameContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ImportScopedNameContext) Get_ID() antlr.Token  { return s._ID }
func (s *ImportScopedNameContext) Set_ID(v antlr.Token) { s._ID = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ImportScopedNameContext) GetA() []antlr.Token  { return s.a }
func (s *ImportScopedNameContext) SetA(v []antlr.Token) { s.a = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportScopedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportScopedNameContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *ImportScopedNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ImportScopedNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ImportScopedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ADLParserDOT)
}

func (s *ImportScopedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserDOT, i)
}

func (s *ImportScopedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportScopedNameEntryListener); ok {
		listenerT.EnterImportScopedName(s)
	}
}

func (s *ImportScopedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportScopedNameExitListener); ok {
		listenerT.ExitImportScopedName(s)
	}
}

func (s *ImportScopedNameContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportScopedName != nil {
		h.ImportScopedName(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportScopedNameContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportScopedNameContextVisitor:
		return t.VisitImportScopedName(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) Imports() (localctx IImportsContext) {
	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADLParserRULE_imports)
	var //TokenTypeDecl
	_la int

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

	var _alt int

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewImportScopedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			var _m = p.Match(ADLParserID)
			localctx.(*ImportScopedNameContext).kw = _m

		}
		{
			p.SetState(57)
			var _m = p.Match(ADLParserID)
			localctx.(*ImportScopedNameContext)._ID = _m

		}
		localctx.(*ImportScopedNameContext).a = append(localctx.(*ImportScopedNameContext).a, localctx.(*ImportScopedNameContext)._ID)
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserDOT {
			{
				p.SetState(58)
				p.Match(ADLParserDOT)
			}
			{
				p.SetState(59)
				var _m = p.Match(ADLParserID)
				localctx.(*ImportScopedNameContext)._ID = _m

			}
			localctx.(*ImportScopedNameContext).a = append(localctx.(*ImportScopedNameContext).a, localctx.(*ImportScopedNameContext)._ID)

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(65)
			p.Match(ADLParserSEMI)
		}

	case 2:
		localctx = NewImportModuleNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			var _m = p.Match(ADLParserID)
			localctx.(*ImportModuleNameContext).kw = _m

		}
		{
			p.SetState(67)
			var _m = p.Match(ADLParserID)
			localctx.(*ImportModuleNameContext)._ID = _m

		}
		localctx.(*ImportModuleNameContext).a = append(localctx.(*ImportModuleNameContext).a, localctx.(*ImportModuleNameContext)._ID)
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(68)
					p.Match(ADLParserDOT)
				}
				{
					p.SetState(69)
					var _m = p.Match(ADLParserID)
					localctx.(*ImportModuleNameContext)._ID = _m

				}
				localctx.(*ImportModuleNameContext).a = append(localctx.(*ImportModuleNameContext).a, localctx.(*ImportModuleNameContext)._ID)

			}
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		{
			p.SetState(75)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(76)
			p.Match(ADLParserSTAR)
		}
		{
			p.SetState(77)
			p.Match(ADLParserSEMI)
		}

	}

	return localctx
}

type IAnnonContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsAnnonContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AnnonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnonContext() *AnnonContext {
	var p = new(AnnonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_annon
	return p
}

func (*AnnonContext) IsAnnonContext() {}

func NewAnnonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnonContext {
	var p = new(AnnonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_annon

	return p
}

func (s *AnnonContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *AnnonContext) CopyFrom(ctx *AnnonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AnnonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type DocAnnoContext struct {
	*AnnonContext
	//TokenDecl
	a antlr.Token
}

func NewDocAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DocAnnoContext {
	var p = new(DocAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

type IDocAnnoContext interface {
	//Current rule
	IAnnonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	LINE_DOC() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*DocAnnoContext) IsDocAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *DocAnnoContext) GetA() antlr.Token  { return s.a }
func (s *DocAnnoContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *DocAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *DocAnnoContext) LINE_DOC() antlr.TerminalNode {
	return s.GetToken(ADLParserLINE_DOC, 0)
}

func (s *DocAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocAnnoEntryListener); ok {
		listenerT.EnterDocAnno(s)
	}
}

func (s *DocAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocAnnoExitListener); ok {
		listenerT.ExitDocAnno(s)
	}
}

func (s *DocAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DocAnno != nil {
		h.DocAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DocAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DocAnnoContextVisitor:
		return t.VisitDocAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type LocalAnnoContext struct {
	*AnnonContext
	//TokenDecl
	a antlr.Token
}

func NewLocalAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalAnnoContext {
	var p = new(LocalAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

type ILocalAnnoContext interface {
	//Current rule
	IAnnonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	AT() antlr.TerminalNode
	ID() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*LocalAnnoContext) IsLocalAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *LocalAnnoContext) GetA() antlr.Token  { return s.a }
func (s *LocalAnnoContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *LocalAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *LocalAnnoContext) AT() antlr.TerminalNode {
	return s.GetToken(ADLParserAT, 0)
}

func (s *LocalAnnoContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *LocalAnnoContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
}

func (s *LocalAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LocalAnnoEntryListener); ok {
		listenerT.EnterLocalAnno(s)
	}
}

func (s *LocalAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LocalAnnoExitListener); ok {
		listenerT.ExitLocalAnno(s)
	}
}

func (s *LocalAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.LocalAnno != nil {
		h.LocalAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *LocalAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case LocalAnnoContextVisitor:
		return t.VisitLocalAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) Annon() (localctx IAnnonContext) {
	localctx = NewAnnonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADLParserRULE_annon)

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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserAT:
		localctx = NewLocalAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(ADLParserAT)
		}
		{
			p.SetState(81)
			var _m = p.Match(ADLParserID)
			localctx.(*LocalAnnoContext).a = _m

		}
		{
			p.SetState(82)
			p.JsonValue()
		}

	case ADLParserLINE_DOC:
		localctx = NewDocAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			var _m = p.Match(ADLParserLINE_DOC)
			localctx.(*DocAnnoContext).a = _m

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITop_level_statementContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTop_level_statementContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type Top_level_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_level_statementContext() *Top_level_statementContext {
	var p = new(Top_level_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_top_level_statement
	return p
}

func (*Top_level_statementContext) IsTop_level_statementContext() {}

func NewTop_level_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_level_statementContext {
	var p = new(Top_level_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_top_level_statement

	return p
}

func (s *Top_level_statementContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *Top_level_statementContext) CopyFrom(ctx *Top_level_statementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Top_level_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_level_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type StructOrUnionContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
}

func NewStructOrUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructOrUnionContext {
	var p = new(StructOrUnionContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IStructOrUnionContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeParam() ITypeParamContext
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	AllSoruBody() []ISoruBodyContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext
	SoruBody(i int) ISoruBodyContext

	//  tokenGetterDecl
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*StructOrUnionContext) IsStructOrUnionContext() {}

//AltLabelStructDecl tokenDecls
func (s *StructOrUnionContext) GetKw() antlr.Token  { return s.kw }
func (s *StructOrUnionContext) SetKw(v antlr.Token) { s.kw = v }

func (s *StructOrUnionContext) GetA() antlr.Token  { return s.a }
func (s *StructOrUnionContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *StructOrUnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *StructOrUnionContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *StructOrUnionContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *StructOrUnionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *StructOrUnionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *StructOrUnionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *StructOrUnionContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *StructOrUnionContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *StructOrUnionContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *StructOrUnionContext) AllSoruBody() []ISoruBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*SoruBodyContext)(nil)).Elem())
	var tst = make([]ISoruBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoruBodyContext)
		}
	}

	return tst
}

func (s *StructOrUnionContext) SoruBody(i int) ISoruBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*SoruBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoruBodyContext)
}

func (s *StructOrUnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructOrUnionEntryListener); ok {
		listenerT.EnterStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructOrUnionExitListener); ok {
		listenerT.ExitStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.StructOrUnion != nil {
		h.StructOrUnion(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *StructOrUnionContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StructOrUnionContextVisitor:
		return t.VisitStructOrUnion(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type DeclAnnotationContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
}

func NewDeclAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnotationContext {
	var p = new(DeclAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IDeclAnnotationContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token
	GetB() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*DeclAnnotationContext) IsDeclAnnotationContext() {}

//AltLabelStructDecl tokenDecls
func (s *DeclAnnotationContext) GetKw() antlr.Token  { return s.kw }
func (s *DeclAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

func (s *DeclAnnotationContext) GetA() antlr.Token  { return s.a }
func (s *DeclAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *DeclAnnotationContext) GetB() antlr.Token  { return s.b }
func (s *DeclAnnotationContext) SetB(v antlr.Token) { s.b = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *DeclAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *DeclAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *DeclAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *DeclAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *DeclAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *DeclAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnotationEntryListener); ok {
		listenerT.EnterDeclAnnotation(s)
	}
}

func (s *DeclAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnotationExitListener); ok {
		listenerT.ExitDeclAnnotation(s)
	}
}

func (s *DeclAnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DeclAnnotation != nil {
		h.DeclAnnotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DeclAnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DeclAnnotationContextVisitor:
		return t.VisitDeclAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type FieldAnnotationContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
	//TokenDecl
	c antlr.Token
}

func NewFieldAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnotationContext {
	var p = new(FieldAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IFieldAnnotationContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DCOLON() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token
	GetB() antlr.Token
	GetC() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FieldAnnotationContext) IsFieldAnnotationContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldAnnotationContext) GetKw() antlr.Token  { return s.kw }
func (s *FieldAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

func (s *FieldAnnotationContext) GetA() antlr.Token  { return s.a }
func (s *FieldAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldAnnotationContext) GetB() antlr.Token  { return s.b }
func (s *FieldAnnotationContext) SetB(v antlr.Token) { s.b = v }

func (s *FieldAnnotationContext) GetC() antlr.Token  { return s.c }
func (s *FieldAnnotationContext) SetC(v antlr.Token) { s.c = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldAnnotationContext) DCOLON() antlr.TerminalNode {
	return s.GetToken(ADLParserDCOLON, 0)
}

func (s *FieldAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *FieldAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *FieldAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *FieldAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *FieldAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnotationEntryListener); ok {
		listenerT.EnterFieldAnnotation(s)
	}
}

func (s *FieldAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnotationExitListener); ok {
		listenerT.ExitFieldAnnotation(s)
	}
}

func (s *FieldAnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FieldAnnotation != nil {
		h.FieldAnnotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldAnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldAnnotationContextVisitor:
		return t.VisitFieldAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type TypeOrNewtypeContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
}

func NewTypeOrNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOrNewtypeContext {
	var p = new(TypeOrNewtypeContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type ITypeOrNewtypeContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeParam() ITypeParamContext
	TypeExpr() ITypeExprContext
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllEQ() []antlr.TerminalNode
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	EQ(i int) antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token
	GetB() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeOrNewtypeContext) IsTypeOrNewtypeContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeOrNewtypeContext) GetKw() antlr.Token  { return s.kw }
func (s *TypeOrNewtypeContext) SetKw(v antlr.Token) { s.kw = v }

func (s *TypeOrNewtypeContext) GetA() antlr.Token  { return s.a }
func (s *TypeOrNewtypeContext) SetA(v antlr.Token) { s.a = v }

func (s *TypeOrNewtypeContext) GetB() antlr.Token  { return s.b }
func (s *TypeOrNewtypeContext) SetB(v antlr.Token) { s.b = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeOrNewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeOrNewtypeContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(ADLParserEQ)
}

func (s *TypeOrNewtypeContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserEQ, i)
}

func (s *TypeOrNewtypeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *TypeOrNewtypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *TypeOrNewtypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *TypeOrNewtypeContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *TypeOrNewtypeContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *TypeOrNewtypeContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *TypeOrNewtypeContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeOrNewtypeContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *TypeOrNewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeOrNewtypeEntryListener); ok {
		listenerT.EnterTypeOrNewtype(s)
	}
}

func (s *TypeOrNewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeOrNewtypeExitListener); ok {
		listenerT.ExitTypeOrNewtype(s)
	}
}

func (s *TypeOrNewtypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeOrNewtype != nil {
		h.TypeOrNewtype(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeOrNewtypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeOrNewtypeContextVisitor:
		return t.VisitTypeOrNewtype(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ModuleAnnotationContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
}

func NewModuleAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleAnnotationContext {
	var p = new(ModuleAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IModuleAnnotationContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ModuleAnnotationContext) IsModuleAnnotationContext() {}

//AltLabelStructDecl tokenDecls
func (s *ModuleAnnotationContext) GetKw() antlr.Token  { return s.kw }
func (s *ModuleAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ModuleAnnotationContext) GetA() antlr.Token  { return s.a }
func (s *ModuleAnnotationContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ModuleAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ModuleAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ModuleAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *ModuleAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ModuleAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ModuleAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleAnnotationEntryListener); ok {
		listenerT.EnterModuleAnnotation(s)
	}
}

func (s *ModuleAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleAnnotationExitListener); ok {
		listenerT.ExitModuleAnnotation(s)
	}
}

func (s *ModuleAnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ModuleAnnotation != nil {
		h.ModuleAnnotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModuleAnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModuleAnnotationContextVisitor:
		return t.VisitModuleAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADLParserRULE_top_level_statement)
	var //TokenTypeDecl
	_la int

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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructOrUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserAT || _la == ADLParserLINE_DOC {
			{
				p.SetState(86)
				p.Annon()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(92)
			var _m = p.Match(ADLParserID)
			localctx.(*StructOrUnionContext).kw = _m

		}
		{
			p.SetState(93)
			var _m = p.Match(ADLParserID)
			localctx.(*StructOrUnionContext).a = _m

		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(94)
				p.TypeParam()
			}

		}
		{
			p.SetState(97)
			p.Match(ADLParserLCUR)
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserAT)|(1<<ADLParserID)|(1<<ADLParserLINE_DOC))) != 0 {
			{
				p.SetState(98)
				p.SoruBody()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(104)
			p.Match(ADLParserRCUR)
		}
		{
			p.SetState(105)
			p.Match(ADLParserSEMI)
		}

	case 2:
		localctx = NewTypeOrNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserAT || _la == ADLParserLINE_DOC {
			{
				p.SetState(106)
				p.Annon()
			}

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			var _m = p.Match(ADLParserID)
			localctx.(*TypeOrNewtypeContext).kw = _m

		}
		{
			p.SetState(113)
			var _m = p.Match(ADLParserID)
			localctx.(*TypeOrNewtypeContext).a = _m

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(114)
				p.TypeParam()
			}

		}
		{
			p.SetState(117)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(118)
			var _m = p.Match(ADLParserID)
			localctx.(*TypeOrNewtypeContext).b = _m

		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(119)
				p.TypeExpr()
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserEQ {
			{
				p.SetState(122)
				p.Match(ADLParserEQ)
			}
			{
				p.SetState(123)
				p.JsonValue()
			}

		}
		{
			p.SetState(126)
			p.Match(ADLParserSEMI)
		}

	case 3:
		localctx = NewModuleAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			var _m = p.Match(ADLParserID)
			localctx.(*ModuleAnnotationContext).kw = _m

		}
		{
			p.SetState(128)
			var _m = p.Match(ADLParserID)
			localctx.(*ModuleAnnotationContext).a = _m

		}
		{
			p.SetState(129)
			p.JsonValue()
		}
		{
			p.SetState(130)
			p.Match(ADLParserSEMI)
		}

	case 4:
		localctx = NewDeclAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(132)
			var _m = p.Match(ADLParserID)
			localctx.(*DeclAnnotationContext).kw = _m

		}
		{
			p.SetState(133)
			var _m = p.Match(ADLParserID)
			localctx.(*DeclAnnotationContext).a = _m

		}
		{
			p.SetState(134)
			var _m = p.Match(ADLParserID)
			localctx.(*DeclAnnotationContext).b = _m

		}
		{
			p.SetState(135)
			p.JsonValue()
		}
		{
			p.SetState(136)
			p.Match(ADLParserSEMI)
		}

	case 5:
		localctx = NewFieldAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			var _m = p.Match(ADLParserID)
			localctx.(*FieldAnnotationContext).kw = _m

		}
		{
			p.SetState(139)
			var _m = p.Match(ADLParserID)
			localctx.(*FieldAnnotationContext).a = _m

		}
		{
			p.SetState(140)
			p.Match(ADLParserDCOLON)
		}
		{
			p.SetState(141)
			var _m = p.Match(ADLParserID)
			localctx.(*FieldAnnotationContext).b = _m

		}
		{
			p.SetState(142)
			var _m = p.Match(ADLParserID)
			localctx.(*FieldAnnotationContext).c = _m

		}
		{
			p.SetState(143)
			p.JsonValue()
		}
		{
			p.SetState(144)
			p.Match(ADLParserSEMI)
		}

	}

	return localctx
}

type ITypeParamContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeParamContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamContext() *TypeParamContext {
	var p = new(TypeParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeParam
	return p
}

func (*TypeParamContext) IsTypeParamContext() {}

func NewTypeParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamContext {
	var p = new(TypeParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeParam

	return p
}

func (s *TypeParamContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeParamContext) CopyFrom(ctx *TypeParamContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeParameterContext struct {
	*TypeParamContext
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	typep []antlr.Token
}

func NewTypeParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParameterContext {
	var p = new(TypeParameterContext)

	p.TypeParamContext = NewEmptyTypeParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeParamContext))

	return p
}

type ITypeParameterContext interface {
	//Current rule
	ITypeParamContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls

	//  tokenTypeDecls
	//  tokenListDecls
	GetTypep() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeParameterContext) IsTypeParameterContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeParameterContext) Get_ID() antlr.Token  { return s._ID }
func (s *TypeParameterContext) Set_ID(v antlr.Token) { s._ID = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *TypeParameterContext) GetTypep() []antlr.Token  { return s.typep }
func (s *TypeParameterContext) SetTypep(v []antlr.Token) { s.typep = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeParameterContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeParameterContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *TypeParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *TypeParameterContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeParameterContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *TypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParameterEntryListener); ok {
		listenerT.EnterTypeParameter(s)
	}
}

func (s *TypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParameterExitListener); ok {
		listenerT.ExitTypeParameter(s)
	}
}

func (s *TypeParameterContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParameter != nil {
		h.TypeParameter(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParameterContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParameterContextVisitor:
		return t.VisitTypeParameter(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ErrorTypeParamContext struct {
	*TypeParamContext
}

func NewErrorTypeParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ErrorTypeParamContext {
	var p = new(ErrorTypeParamContext)

	p.TypeParamContext = NewEmptyTypeParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeParamContext))

	return p
}

type IErrorTypeParamContext interface {
	//Current rule
	ITypeParamContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeParamError() []ITypeParamErrorContext
	//  ruleListIndexedGetterDecl
	TypeParamError(i int) ITypeParamErrorContext

	//  tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ErrorTypeParamContext) IsErrorTypeParamContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ErrorTypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ErrorTypeParamContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *ErrorTypeParamContext) AllTypeParamError() []ITypeParamErrorContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem())
	var tst = make([]ITypeParamErrorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParamErrorContext)
		}
	}

	return tst
}

func (s *ErrorTypeParamContext) TypeParamError(i int) ITypeParamErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParamErrorContext)
}

func (s *ErrorTypeParamContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *ErrorTypeParamContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *ErrorTypeParamContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *ErrorTypeParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ErrorTypeParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ErrorTypeParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ErrorTypeParamEntryListener); ok {
		listenerT.EnterErrorTypeParam(s)
	}
}

func (s *ErrorTypeParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ErrorTypeParamExitListener); ok {
		listenerT.ExitErrorTypeParam(s)
	}
}

func (s *ErrorTypeParamContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ErrorTypeParam != nil {
		h.ErrorTypeParam(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ErrorTypeParamContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ErrorTypeParamContextVisitor:
		return t.VisitErrorTypeParam(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) TypeParam() (localctx ITypeParamContext) {
	localctx = NewTypeParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ADLParserRULE_typeParam)
	var //TokenTypeDecl
	_la int

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

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(ADLParserLCHEVR)
		}
		{
			p.SetState(149)
			var _m = p.Match(ADLParserID)
			localctx.(*TypeParameterContext)._ID = _m

		}
		localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserCOMMA {
			{
				p.SetState(150)
				p.Match(ADLParserCOMMA)
			}
			{
				p.SetState(151)
				var _m = p.Match(ADLParserID)
				localctx.(*TypeParameterContext)._ID = _m

			}
			localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(157)
			p.Match(ADLParserRCHEVR)
		}

	case 2:
		localctx = NewErrorTypeParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(ADLParserLCHEVR)
		}
		{
			p.SetState(159)
			p.TypeParamError()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserCOMMA {
			{
				p.SetState(160)
				p.Match(ADLParserCOMMA)
			}
			{
				p.SetState(161)
				p.Match(ADLParserID)
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(167)
			p.Match(ADLParserRCHEVR)
		}

	case 3:
		localctx = NewErrorTypeParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Match(ADLParserLCHEVR)
		}
		{
			p.SetState(170)
			p.Match(ADLParserID)
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLParserCOMMA {
			{
				p.SetState(171)
				p.Match(ADLParserCOMMA)
			}
			{
				p.SetState(172)
				p.TypeParamError()
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(177)
			p.Match(ADLParserRCHEVR)
		}

	}

	return localctx
}

type ITypeParamErrorContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeParamError() []ITypeParamErrorContext
	//  ruleListIndexedGetterDecl
	TypeParamError(i int) ITypeParamErrorContext
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTypeParamErrorContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeParamErrorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamErrorContext() *TypeParamErrorContext {
	var p = new(TypeParamErrorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeParamError
	return p
}

func (*TypeParamErrorContext) IsTypeParamErrorContext() {}

func NewTypeParamErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamErrorContext {
	var p = new(TypeParamErrorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeParamError

	return p
}

func (s *TypeParamErrorContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *TypeParamErrorContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *TypeParamErrorContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *TypeParamErrorContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeParamErrorContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeParamErrorContext) AllTypeParamError() []ITypeParamErrorContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem())
	var tst = make([]ITypeParamErrorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParamErrorContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) TypeParamError(i int) ITypeParamErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParamErrorContext)
}

func (s *TypeParamErrorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeParamErrorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

//provideCopyFrom
func (s *TypeParamErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeParamErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamErrorEntryListener); ok {
		listenerT.EnterTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamErrorExitListener); ok {
		listenerT.ExitTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParamError != nil {
		h.TypeParamError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParamErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParamErrorContextVisitor:
		return t.VisitTypeParamError(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *ADLParser) TypeParamError() (localctx ITypeParamErrorContext) {
	localctx = NewTypeParamErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ADLParserRULE_typeParamError)
	var //TokenTypeDecl
	_la int

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
		p.SetState(181)
		p.Match(ADLParserID)
	}
	{
		p.SetState(182)
		p.Match(ADLParserLCHEVR)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(183)
			p.Match(ADLParserID)
		}

	case 2:
		{
			p.SetState(184)
			p.TypeParamError()
		}

	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(187)
			p.Match(ADLParserCOMMA)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(188)
				p.Match(ADLParserID)
			}

		case 2:
			{
				p.SetState(189)
				p.TypeParamError()
			}

		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(197)
		p.Match(ADLParserRCHEVR)
	}

	return localctx
}

type ITypeExprContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeExprContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeExpr
	return p
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeExprContext) CopyFrom(ctx *TypeExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeExpressionContext struct {
	*TypeExprContext
	//RuleContextDecl
	_typeExprElem ITypeExprElemContext
	//RuleContextListDecl
	typep []ITypeExprElemContext
}

func NewTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionContext {
	var p = new(TypeExpressionContext)

	p.TypeExprContext = NewEmptyTypeExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprContext))

	return p
}

type ITypeExpressionContext interface {
	//Current rule
	ITypeExprContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem() []ITypeExprElemContext
	//  ruleListIndexedGetterDecl
	TypeExprElem(i int) ITypeExprElemContext

	//  tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls

	//  ruleContextListDecls
	GetTypep() []ITypeExprElemContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeExpressionContext) IsTypeExpressionContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *TypeExpressionContext) Get_typeExprElem() ITypeExprElemContext  { return s._typeExprElem }
func (s *TypeExpressionContext) Set_typeExprElem(v ITypeExprElemContext) { s._typeExprElem = v }

//AltLabelStructDecl ruleContextListDecls
func (s *TypeExpressionContext) GetTypep() []ITypeExprElemContext  { return s.typep }
func (s *TypeExpressionContext) SetTypep(v []ITypeExprElemContext) { s.typep = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeExpressionContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeExpressionContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeExpressionContext) AllTypeExprElem() []ITypeExprElemContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem())
	var tst = make([]ITypeExprElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElemContext)
		}
	}

	return tst
}

func (s *TypeExpressionContext) TypeExprElem(i int) ITypeExprElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElemContext)
}

func (s *TypeExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *TypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionEntryListener); ok {
		listenerT.EnterTypeExpression(s)
	}
}

func (s *TypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionExitListener); ok {
		listenerT.ExitTypeExpression(s)
	}
}

func (s *TypeExpressionContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpression != nil {
		h.TypeExpression(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExpressionContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExpressionContextVisitor:
		return t.VisitTypeExpression(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ADLParserRULE_typeExpr)
	var //TokenTypeDecl
	_la int

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

	localctx = NewTypeExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(ADLParserLCHEVR)
	}
	{
		p.SetState(200)

		var _x = p.TypeExprElem()

		localctx.(*TypeExpressionContext)._typeExprElem = _x

	}
	localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(201)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(202)

			var _x = p.TypeExprElem()

			localctx.(*TypeExpressionContext)._typeExprElem = _x

		}
		localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.Match(ADLParserRCHEVR)
	}

	return localctx
}

type ITypeExprElemContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeExprElemContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExprElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElemContext() *TypeExprElemContext {
	var p = new(TypeExprElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeExprElem
	return p
}

func (*TypeExprElemContext) IsTypeExprElemContext() {}

func NewTypeExprElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElemContext {
	var p = new(TypeExprElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeExprElem

	return p
}

func (s *TypeExprElemContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeExprElemContext) CopyFrom(ctx *TypeExprElemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeExpressionElemContext struct {
	*TypeExprElemContext
	//TokenDecl
	a antlr.Token
	//RuleContextDecl
	_typeExprElem ITypeExprElemContext
	//RuleContextListDecl
	typep []ITypeExprElemContext
}

func NewTypeExpressionElemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionElemContext {
	var p = new(TypeExpressionElemContext)

	p.TypeExprElemContext = NewEmptyTypeExprElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElemContext))

	return p
}

type ITypeExpressionElemContext interface {
	//Current rule
	ITypeExprElemContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem() []ITypeExprElemContext
	//  ruleListIndexedGetterDecl
	TypeExprElem(i int) ITypeExprElemContext

	//  tokenGetterDecl
	ID() antlr.TerminalNode
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls

	//  ruleContextListDecls
	GetTypep() []ITypeExprElemContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeExpressionElemContext) IsTypeExpressionElemContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeExpressionElemContext) GetA() antlr.Token  { return s.a }
func (s *TypeExpressionElemContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *TypeExpressionElemContext) Get_typeExprElem() ITypeExprElemContext  { return s._typeExprElem }
func (s *TypeExpressionElemContext) Set_typeExprElem(v ITypeExprElemContext) { s._typeExprElem = v }

//AltLabelStructDecl ruleContextListDecls
func (s *TypeExpressionElemContext) GetTypep() []ITypeExprElemContext  { return s.typep }
func (s *TypeExpressionElemContext) SetTypep(v []ITypeExprElemContext) { s.typep = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeExpressionElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeExpressionElemContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
}

func (s *TypeExpressionElemContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeExpressionElemContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeExpressionElemContext) AllTypeExprElem() []ITypeExprElemContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem())
	var tst = make([]ITypeExprElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElemContext)
		}
	}

	return tst
}

func (s *TypeExpressionElemContext) TypeExprElem(i int) ITypeExprElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElemContext)
}

func (s *TypeExpressionElemContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeExpressionElemContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *TypeExpressionElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionElemEntryListener); ok {
		listenerT.EnterTypeExpressionElem(s)
	}
}

func (s *TypeExpressionElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionElemExitListener); ok {
		listenerT.ExitTypeExpressionElem(s)
	}
}

func (s *TypeExpressionElemContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpressionElem != nil {
		h.TypeExpressionElem(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExpressionElemContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExpressionElemContextVisitor:
		return t.VisitTypeExpressionElem(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) TypeExprElem() (localctx ITypeExprElemContext) {
	localctx = NewTypeExprElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ADLParserRULE_typeExprElem)
	var //TokenTypeDecl
	_la int

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

	localctx = NewTypeExpressionElemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		var _m = p.Match(ADLParserID)
		localctx.(*TypeExpressionElemContext).a = _m

	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(211)
			p.Match(ADLParserLCHEVR)
		}
		{
			p.SetState(212)

			var _x = p.TypeExprElem()

			localctx.(*TypeExpressionElemContext)._typeExprElem = _x

		}
		localctx.(*TypeExpressionElemContext).typep = append(localctx.(*TypeExpressionElemContext).typep, localctx.(*TypeExpressionElemContext)._typeExprElem)
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserCOMMA {
			{
				p.SetState(213)
				p.Match(ADLParserCOMMA)
			}
			{
				p.SetState(214)

				var _x = p.TypeExprElem()

				localctx.(*TypeExpressionElemContext)._typeExprElem = _x

			}
			localctx.(*TypeExpressionElemContext).typep = append(localctx.(*TypeExpressionElemContext).typep, localctx.(*TypeExpressionElemContext)._typeExprElem)

			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(220)
			p.Match(ADLParserRCHEVR)
		}

	}

	return localctx
}

type ISoruBodyContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsSoruBodyContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type SoruBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoruBodyContext() *SoruBodyContext {
	var p = new(SoruBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_soruBody
	return p
}

func (*SoruBodyContext) IsSoruBodyContext() {}

func NewSoruBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoruBodyContext {
	var p = new(SoruBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_soruBody

	return p
}

func (s *SoruBodyContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *SoruBodyContext) CopyFrom(ctx *SoruBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SoruBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoruBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type FieldStatementContext struct {
	*SoruBodyContext
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
}

func NewFieldStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldStatementContext {
	var p = new(FieldStatementContext)

	p.SoruBodyContext = NewEmptySoruBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SoruBodyContext))

	return p
}

type IFieldStatementContext interface {
	//Current rule
	ISoruBodyContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr() ITypeExprContext
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	EQ() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token
	GetB() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FieldStatementContext) IsFieldStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldStatementContext) GetA() antlr.Token  { return s.a }
func (s *FieldStatementContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldStatementContext) GetB() antlr.Token  { return s.b }
func (s *FieldStatementContext) SetB(v antlr.Token) { s.b = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *FieldStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *FieldStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *FieldStatementContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *FieldStatementContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *FieldStatementContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *FieldStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(ADLParserEQ, 0)
}

func (s *FieldStatementContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *FieldStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldStatementEntryListener); ok {
		listenerT.EnterFieldStatement(s)
	}
}

func (s *FieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldStatementExitListener); ok {
		listenerT.ExitFieldStatement(s)
	}
}

func (s *FieldStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FieldStatement != nil {
		h.FieldStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldStatementContextVisitor:
		return t.VisitFieldStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) SoruBody() (localctx ISoruBodyContext) {
	localctx = NewSoruBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ADLParserRULE_soruBody)
	var //TokenTypeDecl
	_la int

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

	localctx = NewFieldStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserAT || _la == ADLParserLINE_DOC {
		{
			p.SetState(224)
			p.Annon()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(230)
		var _m = p.Match(ADLParserID)
		localctx.(*FieldStatementContext).a = _m

	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(231)
			p.TypeExpr()
		}

	}
	{
		p.SetState(234)
		var _m = p.Match(ADLParserID)
		localctx.(*FieldStatementContext).b = _m

	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserEQ {
		{
			p.SetState(235)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(236)
			p.JsonValue()
		}

	}
	{
		p.SetState(239)
		p.Match(ADLParserSEMI)
	}

	return localctx
}

type IJsonValueContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsJsonValueContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type JsonValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValueContext() *JsonValueContext {
	var p = new(JsonValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_jsonValue
	return p
}

func (*JsonValueContext) IsJsonValueContext() {}

func NewJsonValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValueContext {
	var p = new(JsonValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_jsonValue

	return p
}

func (s *JsonValueContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *JsonValueContext) CopyFrom(ctx *JsonValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TrueFalseNullContext struct {
	*JsonValueContext
	//TokenDecl
	kw antlr.Token
}

func NewTrueFalseNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueFalseNullContext {
	var p = new(TrueFalseNullContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type ITrueFalseNullContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ID() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TrueFalseNullContext) IsTrueFalseNullContext() {}

//AltLabelStructDecl tokenDecls
func (s *TrueFalseNullContext) GetKw() antlr.Token  { return s.kw }
func (s *TrueFalseNullContext) SetKw(v antlr.Token) { s.kw = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TrueFalseNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TrueFalseNullContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
}

func (s *TrueFalseNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrueFalseNullEntryListener); ok {
		listenerT.EnterTrueFalseNull(s)
	}
}

func (s *TrueFalseNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrueFalseNullExitListener); ok {
		listenerT.ExitTrueFalseNull(s)
	}
}

func (s *TrueFalseNullContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TrueFalseNull != nil {
		h.TrueFalseNull(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TrueFalseNullContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TrueFalseNullContextVisitor:
		return t.VisitTrueFalseNull(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ObjStatementContext struct {
	*JsonValueContext
	//TokenDecl
	_STR antlr.Token
	//TokenListDecl
	k []antlr.Token
	//RuleContextDecl
	_jsonValue IJsonValueContext
	//RuleContextListDecl
	v []IJsonValueContext
}

func NewObjStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjStatementContext {
	var p = new(ObjStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IObjStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonValue() []IJsonValueContext
	//  ruleListIndexedGetterDecl
	JsonValue(i int) IJsonValueContext

	//  tokenGetterDecl
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOLON() []antlr.TerminalNode
	AllSTR() []antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COLON(i int) antlr.TerminalNode
	STR(i int) antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls

	//  tokenTypeDecls
	//  tokenListDecls
	GetK() []antlr.Token
	//  ruleContextDecls

	//  ruleContextListDecls
	GetV() []IJsonValueContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ObjStatementContext) IsObjStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *ObjStatementContext) Get_STR() antlr.Token  { return s._STR }
func (s *ObjStatementContext) Set_STR(v antlr.Token) { s._STR = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ObjStatementContext) GetK() []antlr.Token  { return s.k }
func (s *ObjStatementContext) SetK(v []antlr.Token) { s.k = v }

//AltLabelStructDecl ruleContextDecls
func (s *ObjStatementContext) Get_jsonValue() IJsonValueContext  { return s._jsonValue }
func (s *ObjStatementContext) Set_jsonValue(v IJsonValueContext) { s._jsonValue = v }

//AltLabelStructDecl ruleContextListDecls
func (s *ObjStatementContext) GetV() []IJsonValueContext  { return s.v }
func (s *ObjStatementContext) SetV(v []IJsonValueContext) { s.v = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ObjStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ObjStatementContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *ObjStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *ObjStatementContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOLON)
}

func (s *ObjStatementContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOLON, i)
}

func (s *ObjStatementContext) AllSTR() []antlr.TerminalNode {
	return s.GetTokens(ADLParserSTR)
}

func (s *ObjStatementContext) STR(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserSTR, i)
}

func (s *ObjStatementContext) AllJsonValue() []IJsonValueContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValueContext)(nil)).Elem())
	var tst = make([]IJsonValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValueContext)
		}
	}

	return tst
}

func (s *ObjStatementContext) JsonValue(i int) IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ObjStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *ObjStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *ObjStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjStatementEntryListener); ok {
		listenerT.EnterObjStatement(s)
	}
}

func (s *ObjStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjStatementExitListener); ok {
		listenerT.ExitObjStatement(s)
	}
}

func (s *ObjStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ObjStatement != nil {
		h.ObjStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ObjStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ObjStatementContextVisitor:
		return t.VisitObjStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type FloatStatementContext struct {
	*JsonValueContext
	//TokenDecl
	f antlr.Token
}

func NewFloatStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatStatementContext {
	var p = new(FloatStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IFloatStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	FLT() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetF() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FloatStatementContext) IsFloatStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *FloatStatementContext) GetF() antlr.Token  { return s.f }
func (s *FloatStatementContext) SetF(v antlr.Token) { s.f = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FloatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FloatStatementContext) FLT() antlr.TerminalNode {
	return s.GetToken(ADLParserFLT, 0)
}

func (s *FloatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloatStatementEntryListener); ok {
		listenerT.EnterFloatStatement(s)
	}
}

func (s *FloatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloatStatementExitListener); ok {
		listenerT.ExitFloatStatement(s)
	}
}

func (s *FloatStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FloatStatement != nil {
		h.FloatStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FloatStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FloatStatementContextVisitor:
		return t.VisitFloatStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ArrayStatementContext struct {
	*JsonValueContext
	//RuleContextDecl
	_jsonValue IJsonValueContext
	//RuleContextListDecl
	jv []IJsonValueContext
}

func NewArrayStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayStatementContext {
	var p = new(ArrayStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IArrayStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonValue() []IJsonValueContext
	//  ruleListIndexedGetterDecl
	JsonValue(i int) IJsonValueContext

	//  tokenGetterDecl
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls

	//  ruleContextListDecls
	GetJv() []IJsonValueContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ArrayStatementContext) IsArrayStatementContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *ArrayStatementContext) Get_jsonValue() IJsonValueContext  { return s._jsonValue }
func (s *ArrayStatementContext) Set_jsonValue(v IJsonValueContext) { s._jsonValue = v }

//AltLabelStructDecl ruleContextListDecls
func (s *ArrayStatementContext) GetJv() []IJsonValueContext  { return s.jv }
func (s *ArrayStatementContext) SetJv(v []IJsonValueContext) { s.jv = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ArrayStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ArrayStatementContext) LSQ() antlr.TerminalNode {
	return s.GetToken(ADLParserLSQ, 0)
}

func (s *ArrayStatementContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ADLParserRSQ, 0)
}

func (s *ArrayStatementContext) AllJsonValue() []IJsonValueContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValueContext)(nil)).Elem())
	var tst = make([]IJsonValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValueContext)
		}
	}

	return tst
}

func (s *ArrayStatementContext) JsonValue(i int) IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ArrayStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *ArrayStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *ArrayStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayStatementEntryListener); ok {
		listenerT.EnterArrayStatement(s)
	}
}

func (s *ArrayStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayStatementExitListener); ok {
		listenerT.ExitArrayStatement(s)
	}
}

func (s *ArrayStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ArrayStatement != nil {
		h.ArrayStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ArrayStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ArrayStatementContextVisitor:
		return t.VisitArrayStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type NumberStatementContext struct {
	*JsonValueContext
	//TokenDecl
	n antlr.Token
}

func NewNumberStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberStatementContext {
	var p = new(NumberStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type INumberStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	INT() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetN() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*NumberStatementContext) IsNumberStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *NumberStatementContext) GetN() antlr.Token  { return s.n }
func (s *NumberStatementContext) SetN(v antlr.Token) { s.n = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NumberStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *NumberStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(ADLParserINT, 0)
}

func (s *NumberStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumberStatementEntryListener); ok {
		listenerT.EnterNumberStatement(s)
	}
}

func (s *NumberStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumberStatementExitListener); ok {
		listenerT.ExitNumberStatement(s)
	}
}

func (s *NumberStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.NumberStatement != nil {
		h.NumberStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NumberStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NumberStatementContextVisitor:
		return t.VisitNumberStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type StringStatementContext struct {
	*JsonValueContext
	//TokenDecl
	s antlr.Token
}

func NewStringStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringStatementContext {
	var p = new(StringStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IStringStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	STR() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetS() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*StringStatementContext) IsStringStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *StringStatementContext) GetS() antlr.Token  { return s.s }
func (s *StringStatementContext) SetS(v antlr.Token) { s.s = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *StringStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *StringStatementContext) STR() antlr.TerminalNode {
	return s.GetToken(ADLParserSTR, 0)
}

func (s *StringStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StringStatementEntryListener); ok {
		listenerT.EnterStringStatement(s)
	}
}

func (s *StringStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StringStatementExitListener); ok {
		listenerT.ExitStringStatement(s)
	}
}

func (s *StringStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*ADLParserHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.StringStatement != nil {
		h.StringStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *StringStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StringStatementContextVisitor:
		return t.VisitStringStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *ADLParser) JsonValue() (localctx IJsonValueContext) {
	localctx = NewJsonValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ADLParserRULE_jsonValue)
	var //TokenTypeDecl
	_la int

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserSTR:
		localctx = NewStringStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			var _m = p.Match(ADLParserSTR)
			localctx.(*StringStatementContext).s = _m

		}

	case ADLParserID:
		localctx = NewTrueFalseNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			var _m = p.Match(ADLParserID)
			localctx.(*TrueFalseNullContext).kw = _m

		}

	case ADLParserINT:
		localctx = NewNumberStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			var _m = p.Match(ADLParserINT)
			localctx.(*NumberStatementContext).n = _m

		}

	case ADLParserFLT:
		localctx = NewFloatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)
			var _m = p.Match(ADLParserFLT)
			localctx.(*FloatStatementContext).f = _m

		}

	case ADLParserLSQ:
		localctx = NewArrayStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(245)
			p.Match(ADLParserLSQ)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserLCUR)|(1<<ADLParserLSQ)|(1<<ADLParserSTR)|(1<<ADLParserID)|(1<<ADLParserINT)|(1<<ADLParserFLT))) != 0 {
			{
				p.SetState(246)

				var _x = p.JsonValue()

				localctx.(*ArrayStatementContext)._jsonValue = _x

			}
			localctx.(*ArrayStatementContext).jv = append(localctx.(*ArrayStatementContext).jv, localctx.(*ArrayStatementContext)._jsonValue)
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLParserCOMMA {
				{
					p.SetState(247)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(248)

					var _x = p.JsonValue()

					localctx.(*ArrayStatementContext)._jsonValue = _x

				}
				localctx.(*ArrayStatementContext).jv = append(localctx.(*ArrayStatementContext).jv, localctx.(*ArrayStatementContext)._jsonValue)

				p.SetState(253)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(256)
			p.Match(ADLParserRSQ)
		}

	case ADLParserLCUR:
		localctx = NewObjStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(257)
			p.Match(ADLParserLCUR)
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserSTR {
			{
				p.SetState(258)
				var _m = p.Match(ADLParserSTR)
				localctx.(*ObjStatementContext)._STR = _m

			}
			localctx.(*ObjStatementContext).k = append(localctx.(*ObjStatementContext).k, localctx.(*ObjStatementContext)._STR)
			{
				p.SetState(259)
				p.Match(ADLParserCOLON)
			}
			{
				p.SetState(260)

				var _x = p.JsonValue()

				localctx.(*ObjStatementContext)._jsonValue = _x

			}
			localctx.(*ObjStatementContext).v = append(localctx.(*ObjStatementContext).v, localctx.(*ObjStatementContext)._jsonValue)
			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLParserCOMMA {
				{
					p.SetState(261)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(262)
					var _m = p.Match(ADLParserSTR)
					localctx.(*ObjStatementContext)._STR = _m

				}
				localctx.(*ObjStatementContext).k = append(localctx.(*ObjStatementContext).k, localctx.(*ObjStatementContext)._STR)
				{
					p.SetState(263)
					p.Match(ADLParserCOLON)
				}
				{
					p.SetState(264)

					var _x = p.JsonValue()

					localctx.(*ObjStatementContext)._jsonValue = _x

				}
				localctx.(*ObjStatementContext).v = append(localctx.(*ObjStatementContext).v, localctx.(*ObjStatementContext)._jsonValue)

				p.SetState(269)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(272)
			p.Match(ADLParserRCUR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
