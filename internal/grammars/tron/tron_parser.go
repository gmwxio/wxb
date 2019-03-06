// Code generated from tronParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tronParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 424,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 5, 2, 64, 10, 2, 3, 2, 7,
	2, 67, 10, 2, 12, 2, 14, 2, 70, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 5, 4, 81, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 89, 10, 4, 12, 4, 14, 4, 92, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 99, 10, 4, 12, 4, 14, 4, 102, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 118, 10,
	4, 12, 4, 14, 4, 121, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 128, 10,
	4, 13, 4, 14, 4, 129, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 138, 10,
	4, 12, 4, 14, 4, 141, 11, 4, 3, 4, 3, 4, 5, 4, 145, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 6, 5, 152, 10, 5, 13, 5, 14, 5, 153, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 7, 5, 162, 10, 5, 12, 5, 14, 5, 165, 11, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 173, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 179, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 184, 10, 7, 3, 7, 3, 7, 5, 7, 188,
	10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 195, 10, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 205, 10, 9, 12, 9, 14, 9, 208, 11,
	9, 3, 9, 3, 9, 5, 9, 212, 10, 9, 3, 10, 3, 10, 5, 10, 216, 10, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 226, 10, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 233, 10, 13, 12, 13, 14, 13, 236,
	11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 243, 10, 14, 5, 14, 245,
	10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 250, 10, 15, 12, 15, 14, 15, 253, 11,
	15, 3, 16, 5, 16, 256, 10, 16, 3, 16, 3, 16, 7, 16, 260, 10, 16, 12, 16,
	14, 16, 263, 11, 16, 3, 16, 3, 16, 3, 17, 5, 17, 268, 10, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 275, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 286, 10, 19, 12, 19, 14, 19,
	289, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 298,
	10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 5, 21, 312, 10, 21, 3, 21, 3, 21, 3, 22, 5, 22, 317,
	10, 22, 3, 22, 3, 22, 7, 22, 321, 10, 22, 12, 22, 14, 22, 324, 11, 22,
	3, 22, 3, 22, 3, 23, 5, 23, 329, 10, 23, 3, 23, 3, 23, 7, 23, 333, 10,
	23, 12, 23, 14, 23, 336, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 352,
	10, 26, 12, 26, 14, 26, 355, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 368, 10, 28, 12, 28, 14,
	28, 371, 11, 28, 3, 28, 5, 28, 374, 10, 28, 3, 28, 3, 28, 7, 28, 378, 10,
	28, 12, 28, 14, 28, 381, 11, 28, 3, 29, 3, 29, 3, 29, 6, 29, 386, 10, 29,
	13, 29, 14, 29, 387, 3, 29, 3, 29, 5, 29, 392, 10, 29, 3, 29, 3, 29, 5,
	29, 396, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 401, 10, 29, 3, 30, 3, 30,
	3, 30, 5, 30, 406, 10, 30, 3, 30, 7, 30, 409, 10, 30, 12, 30, 14, 30, 412,
	11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31,
	422, 10, 31, 3, 31, 2, 2, 32, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	2, 4, 3, 2, 25, 26, 4, 2, 15, 15, 22, 22, 2, 460, 2, 63, 3, 2, 2, 2, 4,
	73, 3, 2, 2, 2, 6, 144, 3, 2, 2, 2, 8, 172, 3, 2, 2, 2, 10, 178, 3, 2,
	2, 2, 12, 180, 3, 2, 2, 2, 14, 194, 3, 2, 2, 2, 16, 196, 3, 2, 2, 2, 18,
	213, 3, 2, 2, 2, 20, 220, 3, 2, 2, 2, 22, 222, 3, 2, 2, 2, 24, 229, 3,
	2, 2, 2, 26, 244, 3, 2, 2, 2, 28, 246, 3, 2, 2, 2, 30, 255, 3, 2, 2, 2,
	32, 267, 3, 2, 2, 2, 34, 278, 3, 2, 2, 2, 36, 280, 3, 2, 2, 2, 38, 292,
	3, 2, 2, 2, 40, 301, 3, 2, 2, 2, 42, 316, 3, 2, 2, 2, 44, 328, 3, 2, 2,
	2, 46, 339, 3, 2, 2, 2, 48, 341, 3, 2, 2, 2, 50, 347, 3, 2, 2, 2, 52, 358,
	3, 2, 2, 2, 54, 373, 3, 2, 2, 2, 56, 400, 3, 2, 2, 2, 58, 402, 3, 2, 2,
	2, 60, 421, 3, 2, 2, 2, 62, 64, 5, 4, 3, 2, 63, 62, 3, 2, 2, 2, 63, 64,
	3, 2, 2, 2, 64, 68, 3, 2, 2, 2, 65, 67, 5, 6, 4, 2, 66, 65, 3, 2, 2, 2,
	67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71, 3,
	2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 2, 2, 3, 72, 3, 3, 2, 2, 2, 73,
	74, 7, 4, 2, 2, 74, 75, 7, 5, 2, 2, 75, 76, 7, 3, 2, 2, 76, 77, 7, 6, 2,
	2, 77, 5, 3, 2, 2, 2, 78, 80, 7, 28, 2, 2, 79, 81, 7, 28, 2, 2, 80, 79,
	3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 7, 27, 2, 2,
	83, 145, 7, 15, 2, 2, 84, 85, 7, 28, 2, 2, 85, 90, 7, 28, 2, 2, 86, 87,
	7, 19, 2, 2, 87, 89, 7, 28, 2, 2, 88, 86, 3, 2, 2, 2, 89, 92, 3, 2, 2,
	2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 90,
	3, 2, 2, 2, 93, 145, 7, 15, 2, 2, 94, 95, 7, 28, 2, 2, 95, 100, 7, 28,
	2, 2, 96, 97, 7, 19, 2, 2, 97, 99, 7, 28, 2, 2, 98, 96, 3, 2, 2, 2, 99,
	102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 103, 3,
	2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 7, 10, 2, 2, 104, 105, 5, 32,
	17, 2, 105, 106, 7, 11, 2, 2, 106, 145, 3, 2, 2, 2, 107, 108, 7, 28, 2,
	2, 108, 109, 5, 54, 28, 2, 109, 110, 7, 12, 2, 2, 110, 111, 5, 56, 29,
	2, 111, 112, 7, 15, 2, 2, 112, 145, 3, 2, 2, 2, 113, 114, 7, 28, 2, 2,
	114, 115, 7, 28, 2, 2, 115, 119, 7, 10, 2, 2, 116, 118, 5, 8, 5, 2, 117,
	116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120,
	3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 145, 7, 11,
	2, 2, 123, 124, 7, 28, 2, 2, 124, 125, 7, 28, 2, 2, 125, 127, 7, 10, 2,
	2, 126, 128, 5, 10, 6, 2, 127, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132,
	7, 11, 2, 2, 132, 145, 3, 2, 2, 2, 133, 134, 7, 28, 2, 2, 134, 135, 7,
	28, 2, 2, 135, 139, 7, 10, 2, 2, 136, 138, 5, 14, 8, 2, 137, 136, 3, 2,
	2, 2, 138, 141, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2,
	140, 142, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 142, 145, 7, 11, 2, 2, 143,
	145, 7, 15, 2, 2, 144, 78, 3, 2, 2, 2, 144, 84, 3, 2, 2, 2, 144, 94, 3,
	2, 2, 2, 144, 107, 3, 2, 2, 2, 144, 113, 3, 2, 2, 2, 144, 123, 3, 2, 2,
	2, 144, 133, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 7, 3, 2, 2, 2, 146,
	173, 5, 32, 17, 2, 147, 148, 7, 28, 2, 2, 148, 149, 7, 28, 2, 2, 149, 151,
	7, 10, 2, 2, 150, 152, 5, 10, 6, 2, 151, 150, 3, 2, 2, 2, 152, 153, 3,
	2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 3, 2, 2,
	2, 155, 156, 7, 11, 2, 2, 156, 173, 3, 2, 2, 2, 157, 158, 7, 28, 2, 2,
	158, 159, 7, 28, 2, 2, 159, 163, 7, 10, 2, 2, 160, 162, 5, 8, 5, 2, 161,
	160, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164,
	3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 173, 7, 11,
	2, 2, 167, 173, 5, 48, 25, 2, 168, 173, 5, 36, 19, 2, 169, 173, 5, 40,
	21, 2, 170, 173, 5, 22, 12, 2, 171, 173, 5, 46, 24, 2, 172, 146, 3, 2,
	2, 2, 172, 147, 3, 2, 2, 2, 172, 157, 3, 2, 2, 2, 172, 167, 3, 2, 2, 2,
	172, 168, 3, 2, 2, 2, 172, 169, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172,
	171, 3, 2, 2, 2, 173, 9, 3, 2, 2, 2, 174, 179, 5, 48, 25, 2, 175, 179,
	5, 12, 7, 2, 176, 179, 5, 22, 12, 2, 177, 179, 5, 46, 24, 2, 178, 174,
	3, 2, 2, 2, 178, 175, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 177, 3, 2,
	2, 2, 179, 11, 3, 2, 2, 2, 180, 181, 7, 28, 2, 2, 181, 183, 7, 12, 2, 2,
	182, 184, 7, 25, 2, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184,
	185, 3, 2, 2, 2, 185, 187, 7, 29, 2, 2, 186, 188, 5, 50, 26, 2, 187, 186,
	3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 7, 15,
	2, 2, 190, 13, 3, 2, 2, 2, 191, 195, 5, 48, 25, 2, 192, 195, 5, 16, 9,
	2, 193, 195, 5, 46, 24, 2, 194, 191, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2,
	194, 193, 3, 2, 2, 2, 195, 15, 3, 2, 2, 2, 196, 197, 7, 28, 2, 2, 197,
	198, 7, 28, 2, 2, 198, 199, 5, 18, 10, 2, 199, 200, 7, 28, 2, 2, 200, 211,
	5, 18, 10, 2, 201, 206, 7, 10, 2, 2, 202, 205, 5, 48, 25, 2, 203, 205,
	5, 46, 24, 2, 204, 202, 3, 2, 2, 2, 204, 203, 3, 2, 2, 2, 205, 208, 3,
	2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 3, 2, 2,
	2, 208, 206, 3, 2, 2, 2, 209, 212, 7, 11, 2, 2, 210, 212, 7, 15, 2, 2,
	211, 201, 3, 2, 2, 2, 211, 210, 3, 2, 2, 2, 212, 17, 3, 2, 2, 2, 213, 215,
	7, 17, 2, 2, 214, 216, 5, 20, 11, 2, 215, 214, 3, 2, 2, 2, 215, 216, 3,
	2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 5, 42, 22, 2, 218, 219, 7, 18,
	2, 2, 219, 19, 3, 2, 2, 2, 220, 221, 7, 28, 2, 2, 221, 21, 3, 2, 2, 2,
	222, 225, 7, 28, 2, 2, 223, 226, 5, 24, 13, 2, 224, 226, 5, 28, 15, 2,
	225, 223, 3, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227,
	228, 7, 15, 2, 2, 228, 23, 3, 2, 2, 2, 229, 234, 5, 26, 14, 2, 230, 231,
	7, 22, 2, 2, 231, 233, 5, 26, 14, 2, 232, 230, 3, 2, 2, 2, 233, 236, 3,
	2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 25, 3, 2, 2,
	2, 236, 234, 3, 2, 2, 2, 237, 245, 7, 29, 2, 2, 238, 239, 7, 29, 2, 2,
	239, 242, 7, 28, 2, 2, 240, 243, 7, 29, 2, 2, 241, 243, 7, 28, 2, 2, 242,
	240, 3, 2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 245, 3, 2, 2, 2, 244, 237,
	3, 2, 2, 2, 244, 238, 3, 2, 2, 2, 245, 27, 3, 2, 2, 2, 246, 251, 7, 27,
	2, 2, 247, 248, 7, 22, 2, 2, 248, 250, 7, 27, 2, 2, 249, 247, 3, 2, 2,
	2, 250, 253, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252,
	29, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 256, 7, 19, 2, 2, 255, 254,
	3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 261, 3, 2, 2, 2, 257, 258, 7, 28,
	2, 2, 258, 260, 7, 19, 2, 2, 259, 257, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2,
	261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 264, 3, 2, 2, 2, 263,
	261, 3, 2, 2, 2, 264, 265, 7, 28, 2, 2, 265, 31, 3, 2, 2, 2, 266, 268,
	5, 34, 18, 2, 267, 266, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 269, 3,
	2, 2, 2, 269, 270, 5, 30, 16, 2, 270, 271, 7, 28, 2, 2, 271, 272, 7, 12,
	2, 2, 272, 274, 7, 29, 2, 2, 273, 275, 5, 50, 26, 2, 274, 273, 3, 2, 2,
	2, 274, 275, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 7, 15, 2, 2, 277,
	33, 3, 2, 2, 2, 278, 279, 7, 28, 2, 2, 279, 35, 3, 2, 2, 2, 280, 281, 7,
	28, 2, 2, 281, 282, 7, 28, 2, 2, 282, 287, 7, 10, 2, 2, 283, 286, 5, 38,
	20, 2, 284, 286, 5, 46, 24, 2, 285, 283, 3, 2, 2, 2, 285, 284, 3, 2, 2,
	2, 286, 289, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288,
	290, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 290, 291, 7, 11, 2, 2, 291, 37,
	3, 2, 2, 2, 292, 293, 5, 30, 16, 2, 293, 294, 7, 28, 2, 2, 294, 295, 7,
	12, 2, 2, 295, 297, 7, 29, 2, 2, 296, 298, 5, 50, 26, 2, 297, 296, 3, 2,
	2, 2, 297, 298, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 300, 7, 15, 2, 2,
	300, 39, 3, 2, 2, 2, 301, 302, 7, 28, 2, 2, 302, 303, 7, 23, 2, 2, 303,
	304, 7, 28, 2, 2, 304, 305, 7, 22, 2, 2, 305, 306, 5, 30, 16, 2, 306, 307,
	7, 24, 2, 2, 307, 308, 7, 28, 2, 2, 308, 309, 7, 12, 2, 2, 309, 311, 7,
	29, 2, 2, 310, 312, 5, 50, 26, 2, 311, 310, 3, 2, 2, 2, 311, 312, 3, 2,
	2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 7, 15, 2, 2, 314, 41, 3, 2, 2, 2,
	315, 317, 7, 19, 2, 2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317,
	322, 3, 2, 2, 2, 318, 319, 7, 28, 2, 2, 319, 321, 7, 19, 2, 2, 320, 318,
	3, 2, 2, 2, 321, 324, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2,
	2, 2, 323, 325, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 325, 326, 7, 28, 2, 2,
	326, 43, 3, 2, 2, 2, 327, 329, 7, 19, 2, 2, 328, 327, 3, 2, 2, 2, 328,
	329, 3, 2, 2, 2, 329, 334, 3, 2, 2, 2, 330, 331, 7, 28, 2, 2, 331, 333,
	7, 19, 2, 2, 332, 330, 3, 2, 2, 2, 333, 336, 3, 2, 2, 2, 334, 332, 3, 2,
	2, 2, 334, 335, 3, 2, 2, 2, 335, 337, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2,
	337, 338, 7, 28, 2, 2, 338, 45, 3, 2, 2, 2, 339, 340, 7, 15, 2, 2, 340,
	47, 3, 2, 2, 2, 341, 342, 7, 28, 2, 2, 342, 343, 5, 54, 28, 2, 343, 344,
	7, 12, 2, 2, 344, 345, 5, 56, 29, 2, 345, 346, 7, 15, 2, 2, 346, 49, 3,
	2, 2, 2, 347, 348, 7, 20, 2, 2, 348, 353, 5, 52, 27, 2, 349, 350, 7, 22,
	2, 2, 350, 352, 5, 52, 27, 2, 351, 349, 3, 2, 2, 2, 352, 355, 3, 2, 2,
	2, 353, 351, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 356, 3, 2, 2, 2, 355,
	353, 3, 2, 2, 2, 356, 357, 7, 21, 2, 2, 357, 51, 3, 2, 2, 2, 358, 359,
	5, 54, 28, 2, 359, 360, 7, 12, 2, 2, 360, 361, 5, 56, 29, 2, 361, 53, 3,
	2, 2, 2, 362, 374, 7, 28, 2, 2, 363, 364, 7, 17, 2, 2, 364, 369, 7, 28,
	2, 2, 365, 366, 7, 19, 2, 2, 366, 368, 7, 28, 2, 2, 367, 365, 3, 2, 2,
	2, 368, 371, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370,
	372, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 372, 374, 7, 18, 2, 2, 373, 362,
	3, 2, 2, 2, 373, 363, 3, 2, 2, 2, 374, 379, 3, 2, 2, 2, 375, 376, 7, 19,
	2, 2, 376, 378, 7, 28, 2, 2, 377, 375, 3, 2, 2, 2, 378, 381, 3, 2, 2, 2,
	379, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 55, 3, 2, 2, 2, 381, 379,
	3, 2, 2, 2, 382, 385, 7, 28, 2, 2, 383, 384, 7, 19, 2, 2, 384, 386, 7,
	28, 2, 2, 385, 383, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 385, 3, 2, 2,
	2, 387, 388, 3, 2, 2, 2, 388, 401, 3, 2, 2, 2, 389, 401, 7, 28, 2, 2, 390,
	392, 9, 2, 2, 2, 391, 390, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 393,
	3, 2, 2, 2, 393, 401, 7, 29, 2, 2, 394, 396, 9, 2, 2, 2, 395, 394, 3, 2,
	2, 2, 395, 396, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 401, 7, 30, 2, 2,
	398, 401, 7, 27, 2, 2, 399, 401, 5, 58, 30, 2, 400, 382, 3, 2, 2, 2, 400,
	389, 3, 2, 2, 2, 400, 391, 3, 2, 2, 2, 400, 395, 3, 2, 2, 2, 400, 398,
	3, 2, 2, 2, 400, 399, 3, 2, 2, 2, 401, 57, 3, 2, 2, 2, 402, 403, 7, 10,
	2, 2, 403, 410, 5, 60, 31, 2, 404, 406, 9, 3, 2, 2, 405, 404, 3, 2, 2,
	2, 405, 406, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 409, 5, 60, 31, 2,
	408, 405, 3, 2, 2, 2, 409, 412, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 410,
	411, 3, 2, 2, 2, 411, 413, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 413, 414,
	7, 11, 2, 2, 414, 59, 3, 2, 2, 2, 415, 416, 7, 28, 2, 2, 416, 417, 7, 16,
	2, 2, 417, 422, 5, 56, 29, 2, 418, 419, 7, 28, 2, 2, 419, 420, 7, 16, 2,
	2, 420, 422, 5, 58, 30, 2, 421, 415, 3, 2, 2, 2, 421, 418, 3, 2, 2, 2,
	422, 61, 3, 2, 2, 2, 50, 63, 68, 80, 90, 100, 119, 129, 139, 144, 153,
	163, 172, 178, 183, 187, 194, 204, 206, 211, 215, 225, 234, 242, 244, 251,
	255, 261, 267, 274, 285, 287, 297, 311, 316, 322, 328, 334, 353, 369, 373,
	379, 387, 391, 395, 400, 405, 410, 421,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'syntax'", "", "", "", "", "", "'{'", "'}'", "", "'\"'", "'''",
	"", "':'", "'('", "')'", "'.'", "'['", "']'", "','", "'<'", "'>'", "'-'",
	"'+'",
}
var symbolicNames = []string{
	"", "PROTO3", "SYNTAX", "EQ_PRE", "ENDPRE", "PRE_WS", "PRE_COMMENT", "PRE_LINE_COMMENT",
	"LCUR", "RCUR", "EQ", "DQ", "SQ", "SEMI", "COLON", "LPAREN", "RPAREN",
	"DOT", "LSB", "RSB", "COMMA", "LCHEVR", "RCHEVR", "DASH", "PLUS", "STR",
	"ID", "INT", "FLT", "WS", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"proto", "syntax", "top_level_statement", "messageBody", "enumBody", "enumField",
	"serviceBody", "rpc", "rpcParam", "rpcStream", "reserved", "ranges", "rangee",
	"fieldNames", "typer", "field", "fieldRepeat", "oneof", "oneofField", "mapField",
	"messageType", "messageOrEnumType", "emptyStatement", "option", "fieldOptions",
	"fieldOption", "optionName", "constant", "constantObj", "constantObjElem",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tronParser struct {
	*antlr.BaseParser
}

func NewtronParser(input antlr.TokenStream) *tronParser {
	this := new(tronParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tronParser.g4"

	return this
}

// tronParser tokens.
const (
	tronParserEOF              = antlr.TokenEOF
	tronParserPROTO3           = 1
	tronParserSYNTAX           = 2
	tronParserEQ_PRE           = 3
	tronParserENDPRE           = 4
	tronParserPRE_WS           = 5
	tronParserPRE_COMMENT      = 6
	tronParserPRE_LINE_COMMENT = 7
	tronParserLCUR             = 8
	tronParserRCUR             = 9
	tronParserEQ               = 10
	tronParserDQ               = 11
	tronParserSQ               = 12
	tronParserSEMI             = 13
	tronParserCOLON            = 14
	tronParserLPAREN           = 15
	tronParserRPAREN           = 16
	tronParserDOT              = 17
	tronParserLSB              = 18
	tronParserRSB              = 19
	tronParserCOMMA            = 20
	tronParserLCHEVR           = 21
	tronParserRCHEVR           = 22
	tronParserDASH             = 23
	tronParserPLUS             = 24
	tronParserSTR              = 25
	tronParserID               = 26
	tronParserINT              = 27
	tronParserFLT              = 28
	tronParserWS               = 29
	tronParserCOMMENT          = 30
	tronParserLINE_COMMENT     = 31
)

// tronParser rules.
const (
	tronParserRULE_proto               = 0
	tronParserRULE_syntax              = 1
	tronParserRULE_top_level_statement = 2
	tronParserRULE_messageBody         = 3
	tronParserRULE_enumBody            = 4
	tronParserRULE_enumField           = 5
	tronParserRULE_serviceBody         = 6
	tronParserRULE_rpc                 = 7
	tronParserRULE_rpcParam            = 8
	tronParserRULE_rpcStream           = 9
	tronParserRULE_reserved            = 10
	tronParserRULE_ranges              = 11
	tronParserRULE_rangee              = 12
	tronParserRULE_fieldNames          = 13
	tronParserRULE_typer               = 14
	tronParserRULE_field               = 15
	tronParserRULE_fieldRepeat         = 16
	tronParserRULE_oneof               = 17
	tronParserRULE_oneofField          = 18
	tronParserRULE_mapField            = 19
	tronParserRULE_messageType         = 20
	tronParserRULE_messageOrEnumType   = 21
	tronParserRULE_emptyStatement      = 22
	tronParserRULE_option              = 23
	tronParserRULE_fieldOptions        = 24
	tronParserRULE_fieldOption         = 25
	tronParserRULE_optionName          = 26
	tronParserRULE_constant            = 27
	tronParserRULE_constantObj         = 28
	tronParserRULE_constantObjElem     = 29
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(tronParserEOF, 0)
}

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) AllTop_level_statement() []ITop_level_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITop_level_statementContext)(nil)).Elem())
	var tst = make([]ITop_level_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITop_level_statementContext)
		}
	}

	return tst
}

func (s *ProtoContext) Top_level_statement(i int) ITop_level_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITop_level_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITop_level_statementContext)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *tronParser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tronParserRULE_proto)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserSYNTAX {
		{
			p.SetState(60)
			p.Syntax()
		}

	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserSEMI || _la == tronParserID {
		{
			p.SetState(63)
			p.Top_level_statement()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
		p.Match(tronParserEOF)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(tronParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ_PRE() antlr.TerminalNode {
	return s.GetToken(tronParserEQ_PRE, 0)
}

func (s *SyntaxContext) PROTO3() antlr.TerminalNode {
	return s.GetToken(tronParserPROTO3, 0)
}

func (s *SyntaxContext) ENDPRE() antlr.TerminalNode {
	return s.GetToken(tronParserENDPRE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *tronParser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tronParserRULE_syntax)

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
		p.SetState(71)
		p.Match(tronParserSYNTAX)
	}
	{
		p.SetState(72)
		p.Match(tronParserEQ_PRE)
	}
	{
		p.SetState(73)
		p.Match(tronParserPROTO3)
	}
	{
		p.SetState(74)
		p.Match(tronParserENDPRE)
	}

	return localctx
}

// ITop_level_statementContext is an interface to support dynamic dispatch.
type ITop_level_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTop_level_statementContext differentiates from other interfaces.
	IsTop_level_statementContext()
}

type Top_level_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_level_statementContext() *Top_level_statementContext {
	var p = new(Top_level_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_top_level_statement
	return p
}

func (*Top_level_statementContext) IsTop_level_statementContext() {}

func NewTop_level_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_level_statementContext {
	var p = new(Top_level_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_top_level_statement

	return p
}

func (s *Top_level_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Top_level_statementContext) CopyFrom(ctx *Top_level_statementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Top_level_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_level_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OptionDefContext struct {
	*Top_level_statementContext
	opt antlr.Token
}

func NewOptionDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionDefContext {
	var p = new(OptionDefContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *OptionDefContext) GetOpt() antlr.Token { return s.opt }

func (s *OptionDefContext) SetOpt(v antlr.Token) { s.opt = v }

func (s *OptionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionDefContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionDefContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *OptionDefContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *OptionDefContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *OptionDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOptionDef(s)
	}
}

func (s *OptionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOptionDef(s)
	}
}

type MessageContext struct {
	*Top_level_statementContext
	msg antlr.Token
}

func NewMessageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MessageContext {
	var p = new(MessageContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *MessageContext) GetMsg() antlr.Token { return s.msg }

func (s *MessageContext) SetMsg(v antlr.Token) { s.msg = v }

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *MessageContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *MessageContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *MessageContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *MessageContext) AllMessageBody() []IMessageBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem())
	var tst = make([]IMessageBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageBodyContext)
		}
	}

	return tst
}

func (s *MessageContext) MessageBody(i int) IMessageBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitMessage(s)
	}
}

type ExtendContext struct {
	*Top_level_statementContext
	ext antlr.Token
}

func NewExtendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExtendContext {
	var p = new(ExtendContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *ExtendContext) GetExt() antlr.Token { return s.ext }

func (s *ExtendContext) SetExt(v antlr.Token) { s.ext = v }

func (s *ExtendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *ExtendContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *ExtendContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *ExtendContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExtendContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *ExtendContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *ExtendContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *ExtendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterExtend(s)
	}
}

func (s *ExtendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitExtend(s)
	}
}

type ServiceContext struct {
	*Top_level_statementContext
	svc antlr.Token
}

func NewServiceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ServiceContext {
	var p = new(ServiceContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *ServiceContext) GetSvc() antlr.Token { return s.svc }

func (s *ServiceContext) SetSvc(v antlr.Token) { s.svc = v }

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *ServiceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *ServiceContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *ServiceContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *ServiceContext) AllServiceBody() []IServiceBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceBodyContext)(nil)).Elem())
	var tst = make([]IServiceBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceBodyContext)
		}
	}

	return tst
}

func (s *ServiceContext) ServiceBody(i int) IServiceBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceBodyContext)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitService(s)
	}
}

type ImportStatementContext struct {
	*Top_level_statementContext
	imp          antlr.Token
	weakORpublic antlr.Token
}

func NewImportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *ImportStatementContext) GetImp() antlr.Token { return s.imp }

func (s *ImportStatementContext) GetWeakORpublic() antlr.Token { return s.weakORpublic }

func (s *ImportStatementContext) SetImp(v antlr.Token) { s.imp = v }

func (s *ImportStatementContext) SetWeakORpublic(v antlr.Token) { s.weakORpublic = v }

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) STR() antlr.TerminalNode {
	return s.GetToken(tronParserSTR, 0)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *ImportStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *ImportStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

type EnumDefinitionContext struct {
	*Top_level_statementContext
	en antlr.Token
}

func NewEnumDefinitionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *EnumDefinitionContext) GetEn() antlr.Token { return s.en }

func (s *EnumDefinitionContext) SetEn(v antlr.Token) { s.en = v }

func (s *EnumDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefinitionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *EnumDefinitionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *EnumDefinitionContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *EnumDefinitionContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *EnumDefinitionContext) AllEnumBody() []IEnumBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem())
	var tst = make([]IEnumBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumBodyContext)
		}
	}

	return tst
}

func (s *EnumDefinitionContext) EnumBody(i int) IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterEnumDefinition(s)
	}
}

func (s *EnumDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitEnumDefinition(s)
	}
}

type EmptyStmContext struct {
	*Top_level_statementContext
}

func NewEmptyStmContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyStmContext {
	var p = new(EmptyStmContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *EmptyStmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStmContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *EmptyStmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterEmptyStm(s)
	}
}

func (s *EmptyStmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitEmptyStm(s)
	}
}

type PackageStatementContext struct {
	*Top_level_statementContext
	pkg antlr.Token
}

func NewPackageStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *PackageStatementContext) GetPkg() antlr.Token { return s.pkg }

func (s *PackageStatementContext) SetPkg(v antlr.Token) { s.pkg = v }

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *PackageStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *PackageStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *PackageStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *tronParser) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tronParserRULE_top_level_statement)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewImportStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)

			var _m = p.Match(tronParserID)

			localctx.(*ImportStatementContext).imp = _m
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tronParserID {
			{
				p.SetState(77)

				var _m = p.Match(tronParserID)

				localctx.(*ImportStatementContext).weakORpublic = _m
			}

		}
		{
			p.SetState(80)
			p.Match(tronParserSTR)
		}
		{
			p.SetState(81)
			p.Match(tronParserSEMI)
		}

	case 2:
		localctx = NewPackageStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)

			var _m = p.Match(tronParserID)

			localctx.(*PackageStatementContext).pkg = _m
		}
		{
			p.SetState(83)
			p.Match(tronParserID)
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserDOT {
			{
				p.SetState(84)
				p.Match(tronParserDOT)
			}
			{
				p.SetState(85)
				p.Match(tronParserID)
			}

			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(91)
			p.Match(tronParserSEMI)
		}

	case 3:
		localctx = NewExtendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)

			var _m = p.Match(tronParserID)

			localctx.(*ExtendContext).ext = _m
		}
		{
			p.SetState(93)
			p.Match(tronParserID)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserDOT {
			{
				p.SetState(94)
				p.Match(tronParserDOT)
			}
			{
				p.SetState(95)
				p.Match(tronParserID)
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(tronParserLCUR)
		}
		{
			p.SetState(102)
			p.Field()
		}
		{
			p.SetState(103)
			p.Match(tronParserRCUR)
		}

	case 4:
		localctx = NewOptionDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)

			var _m = p.Match(tronParserID)

			localctx.(*OptionDefContext).opt = _m
		}
		{
			p.SetState(106)
			p.OptionName()
		}
		{
			p.SetState(107)
			p.Match(tronParserEQ)
		}

		{
			p.SetState(108)
			p.Constant()
		}

		{
			p.SetState(109)
			p.Match(tronParserSEMI)
		}

	case 5:
		localctx = NewMessageContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)

			var _m = p.Match(tronParserID)

			localctx.(*MessageContext).msg = _m
		}
		{
			p.SetState(112)
			p.Match(tronParserID)
		}
		{
			p.SetState(113)
			p.Match(tronParserLCUR)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tronParserSEMI)|(1<<tronParserDOT)|(1<<tronParserID))) != 0 {
			{
				p.SetState(114)
				p.MessageBody()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
			p.Match(tronParserRCUR)
		}

	case 6:
		localctx = NewEnumDefinitionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)

			var _m = p.Match(tronParserID)

			localctx.(*EnumDefinitionContext).en = _m
		}
		{
			p.SetState(122)
			p.Match(tronParserID)
		}
		{
			p.SetState(123)
			p.Match(tronParserLCUR)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tronParserSEMI || _la == tronParserID {
			{
				p.SetState(124)
				p.EnumBody()
			}

			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(tronParserRCUR)
		}

	case 7:
		localctx = NewServiceContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(131)

			var _m = p.Match(tronParserID)

			localctx.(*ServiceContext).svc = _m
		}
		{
			p.SetState(132)
			p.Match(tronParserID)
		}
		{
			p.SetState(133)
			p.Match(tronParserLCUR)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserSEMI || _la == tronParserID {
			{
				p.SetState(134)
				p.ServiceBody()
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(140)
			p.Match(tronParserRCUR)
		}

	case 8:
		localctx = NewEmptyStmContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(141)
			p.Match(tronParserSEMI)
		}

	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEn returns the en token.
	GetEn() antlr.Token

	// GetMsg returns the msg token.
	GetMsg() antlr.Token

	// SetEn sets the en token.
	SetEn(antlr.Token)

	// SetMsg sets the msg token.
	SetMsg(antlr.Token)

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	en     antlr.Token
	msg    antlr.Token
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) GetEn() antlr.Token { return s.en }

func (s *MessageBodyContext) GetMsg() antlr.Token { return s.msg }

func (s *MessageBodyContext) SetEn(v antlr.Token) { s.en = v }

func (s *MessageBodyContext) SetMsg(v antlr.Token) { s.msg = v }

func (s *MessageBodyContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageBodyContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *MessageBodyContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *MessageBodyContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *MessageBodyContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *MessageBodyContext) AllEnumBody() []IEnumBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem())
	var tst = make([]IEnumBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumBodyContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) EnumBody(i int) IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *MessageBodyContext) AllMessageBody() []IMessageBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem())
	var tst = make([]IMessageBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageBodyContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageBody(i int) IMessageBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageBodyContext) Option() IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *MessageBodyContext) Oneof() IOneofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageBodyContext) MapField() IMapFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageBodyContext) Reserved() IReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageBodyContext) EmptyStatement() IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (p *tronParser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tronParserRULE_messageBody)
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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)

			var _m = p.Match(tronParserID)

			localctx.(*MessageBodyContext).en = _m
		}
		{
			p.SetState(146)
			p.Match(tronParserID)
		}
		{
			p.SetState(147)
			p.Match(tronParserLCUR)
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tronParserSEMI || _la == tronParserID {
			{
				p.SetState(148)
				p.EnumBody()
			}

			p.SetState(151)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(153)
			p.Match(tronParserRCUR)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)

			var _m = p.Match(tronParserID)

			localctx.(*MessageBodyContext).msg = _m
		}
		{
			p.SetState(156)
			p.Match(tronParserID)
		}
		{
			p.SetState(157)
			p.Match(tronParserLCUR)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tronParserSEMI)|(1<<tronParserDOT)|(1<<tronParserID))) != 0 {
			{
				p.SetState(158)
				p.MessageBody()
			}

			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(164)
			p.Match(tronParserRCUR)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(165)
			p.Option()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(166)
			p.Oneof()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(167)
			p.MapField()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(168)
			p.Reserved()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(169)
			p.EmptyStatement()
		}

	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) Option() IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *EnumBodyContext) EnumField() IEnumFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumBodyContext) Reserved() IReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *EnumBodyContext) EmptyStatement() IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *tronParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tronParserRULE_enumBody)

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

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Option()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.EnumField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(174)
			p.Reserved()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(175)
			p.EmptyStatement()
		}

	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *EnumFieldContext) INT() antlr.TerminalNode {
	return s.GetToken(tronParserINT, 0)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *EnumFieldContext) DASH() antlr.TerminalNode {
	return s.GetToken(tronParserDASH, 0)
}

func (s *EnumFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *tronParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tronParserRULE_enumField)
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
		p.SetState(178)
		p.Match(tronParserID)
	}
	{
		p.SetState(179)
		p.Match(tronParserEQ)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDASH {
		{
			p.SetState(180)
			p.Match(tronParserDASH)
		}

	}
	{
		p.SetState(183)
		p.Match(tronParserINT)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(184)
			p.FieldOptions()
		}

	}
	{
		p.SetState(187)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IServiceBodyContext is an interface to support dynamic dispatch.
type IServiceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceBodyContext differentiates from other interfaces.
	IsServiceBodyContext()
}

type ServiceBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceBodyContext() *ServiceBodyContext {
	var p = new(ServiceBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_serviceBody
	return p
}

func (*ServiceBodyContext) IsServiceBodyContext() {}

func NewServiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBodyContext {
	var p = new(ServiceBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_serviceBody

	return p
}

func (s *ServiceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceBodyContext) Option() IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ServiceBodyContext) Rpc() IRpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceBodyContext) EmptyStatement() IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *ServiceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterServiceBody(s)
	}
}

func (s *ServiceBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitServiceBody(s)
	}
}

func (p *tronParser) ServiceBody() (localctx IServiceBodyContext) {
	localctx = NewServiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tronParserRULE_serviceBody)

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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Option()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Rpc()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.EmptyStatement()
		}

	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRpcID returns the rpcID token.
	GetRpcID() antlr.Token

	// GetRets returns the rets token.
	GetRets() antlr.Token

	// SetRpcID sets the rpcID token.
	SetRpcID(antlr.Token)

	// SetRets sets the rets token.
	SetRets(antlr.Token)

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	rpcID  antlr.Token
	rets   antlr.Token
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) GetRpcID() antlr.Token { return s.rpcID }

func (s *RpcContext) GetRets() antlr.Token { return s.rets }

func (s *RpcContext) SetRpcID(v antlr.Token) { s.rpcID = v }

func (s *RpcContext) SetRets(v antlr.Token) { s.rets = v }

func (s *RpcContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *RpcContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *RpcContext) AllRpcParam() []IRpcParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRpcParamContext)(nil)).Elem())
	var tst = make([]IRpcParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRpcParamContext)
		}
	}

	return tst
}

func (s *RpcContext) RpcParam(i int) IRpcParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRpcParamContext)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *RpcContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *RpcContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *RpcContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *RpcContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *RpcContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitRpc(s)
	}
}

func (p *tronParser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tronParserRULE_rpc)
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
		p.SetState(194)

		var _m = p.Match(tronParserID)

		localctx.(*RpcContext).rpcID = _m
	}
	{
		p.SetState(195)
		p.Match(tronParserID)
	}
	{
		p.SetState(196)
		p.RpcParam()
	}
	{
		p.SetState(197)

		var _m = p.Match(tronParserID)

		localctx.(*RpcContext).rets = _m
	}
	{
		p.SetState(198)
		p.RpcParam()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserLCUR:
		{
			p.SetState(199)
			p.Match(tronParserLCUR)
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserSEMI || _la == tronParserID {
			p.SetState(202)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case tronParserID:
				{
					p.SetState(200)
					p.Option()
				}

			case tronParserSEMI:
				{
					p.SetState(201)
					p.EmptyStatement()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(207)
			p.Match(tronParserRCUR)
		}

	case tronParserSEMI:
		{
			p.SetState(208)
			p.Match(tronParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRpcParamContext is an interface to support dynamic dispatch.
type IRpcParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcParamContext differentiates from other interfaces.
	IsRpcParamContext()
}

type RpcParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcParamContext() *RpcParamContext {
	var p = new(RpcParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_rpcParam
	return p
}

func (*RpcParamContext) IsRpcParamContext() {}

func NewRpcParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcParamContext {
	var p = new(RpcParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_rpcParam

	return p
}

func (s *RpcParamContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcParamContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tronParserLPAREN, 0)
}

func (s *RpcParamContext) MessageType() IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RpcParamContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tronParserRPAREN, 0)
}

func (s *RpcParamContext) RpcStream() IRpcStreamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcStreamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcStreamContext)
}

func (s *RpcParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterRpcParam(s)
	}
}

func (s *RpcParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitRpcParam(s)
	}
}

func (p *tronParser) RpcParam() (localctx IRpcParamContext) {
	localctx = NewRpcParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tronParserRULE_rpcParam)

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
		p.Match(tronParserLPAREN)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(212)
			p.RpcStream()
		}

	}
	{
		p.SetState(215)
		p.MessageType()
	}
	{
		p.SetState(216)
		p.Match(tronParserRPAREN)
	}

	return localctx
}

// IRpcStreamContext is an interface to support dynamic dispatch.
type IRpcStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStream returns the stream token.
	GetStream() antlr.Token

	// SetStream sets the stream token.
	SetStream(antlr.Token)

	// IsRpcStreamContext differentiates from other interfaces.
	IsRpcStreamContext()
}

type RpcStreamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	stream antlr.Token
}

func NewEmptyRpcStreamContext() *RpcStreamContext {
	var p = new(RpcStreamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_rpcStream
	return p
}

func (*RpcStreamContext) IsRpcStreamContext() {}

func NewRpcStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcStreamContext {
	var p = new(RpcStreamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_rpcStream

	return p
}

func (s *RpcStreamContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcStreamContext) GetStream() antlr.Token { return s.stream }

func (s *RpcStreamContext) SetStream(v antlr.Token) { s.stream = v }

func (s *RpcStreamContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *RpcStreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcStreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcStreamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterRpcStream(s)
	}
}

func (s *RpcStreamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitRpcStream(s)
	}
}

func (p *tronParser) RpcStream() (localctx IRpcStreamContext) {
	localctx = NewRpcStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tronParserRULE_rpcStream)

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
		p.SetState(218)

		var _m = p.Match(tronParserID)

		localctx.(*RpcStreamContext).stream = _m
	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRes returns the res token.
	GetRes() antlr.Token

	// SetRes sets the res token.
	SetRes(antlr.Token)

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	res    antlr.Token
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) GetRes() antlr.Token { return s.res }

func (s *ReservedContext) SetRes(v antlr.Token) { s.res = v }

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *ReservedContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) FieldNames() IFieldNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitReserved(s)
	}
}

func (p *tronParser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tronParserRULE_reserved)

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
		p.SetState(220)

		var _m = p.Match(tronParserID)

		localctx.(*ReservedContext).res = _m
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserINT:
		{
			p.SetState(221)
			p.Ranges()
		}

	case tronParserSTR:
		{
			p.SetState(222)
			p.FieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(225)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRangee() []IRangeeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRangeeContext)(nil)).Elem())
	var tst = make([]IRangeeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRangeeContext)
		}
	}

	return tst
}

func (s *RangesContext) Rangee(i int) IRangeeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRangeeContext)
}

func (s *RangesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tronParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tronParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *tronParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tronParserRULE_ranges)
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
		p.SetState(227)
		p.Rangee()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(228)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(229)
			p.Rangee()
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRangeeContext is an interface to support dynamic dispatch.
type IRangeeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTo returns the to token.
	GetTo() antlr.Token

	// GetMax returns the max token.
	GetMax() antlr.Token

	// SetTo sets the to token.
	SetTo(antlr.Token)

	// SetMax sets the max token.
	SetMax(antlr.Token)

	// IsRangeeContext differentiates from other interfaces.
	IsRangeeContext()
}

type RangeeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	to     antlr.Token
	max    antlr.Token
}

func NewEmptyRangeeContext() *RangeeContext {
	var p = new(RangeeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_rangee
	return p
}

func (*RangeeContext) IsRangeeContext() {}

func NewRangeeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeeContext {
	var p = new(RangeeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_rangee

	return p
}

func (s *RangeeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeeContext) GetTo() antlr.Token { return s.to }

func (s *RangeeContext) GetMax() antlr.Token { return s.max }

func (s *RangeeContext) SetTo(v antlr.Token) { s.to = v }

func (s *RangeeContext) SetMax(v antlr.Token) { s.max = v }

func (s *RangeeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(tronParserINT)
}

func (s *RangeeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserINT, i)
}

func (s *RangeeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *RangeeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *RangeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterRangee(s)
	}
}

func (s *RangeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitRangee(s)
	}
}

func (p *tronParser) Rangee() (localctx IRangeeContext) {
	localctx = NewRangeeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tronParserRULE_rangee)

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

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Match(tronParserINT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(tronParserINT)
		}
		{
			p.SetState(237)

			var _m = p.Match(tronParserID)

			localctx.(*RangeeContext).to = _m
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tronParserINT:
			{
				p.SetState(238)
				p.Match(tronParserINT)
			}

		case tronParserID:
			{
				p.SetState(239)

				var _m = p.Match(tronParserID)

				localctx.(*RangeeContext).max = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IFieldNamesContext is an interface to support dynamic dispatch.
type IFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNamesContext differentiates from other interfaces.
	IsFieldNamesContext()
}

type FieldNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNamesContext() *FieldNamesContext {
	var p = new(FieldNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_fieldNames
	return p
}

func (*FieldNamesContext) IsFieldNamesContext() {}

func NewFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNamesContext {
	var p = new(FieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_fieldNames

	return p
}

func (s *FieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNamesContext) AllSTR() []antlr.TerminalNode {
	return s.GetTokens(tronParserSTR)
}

func (s *FieldNamesContext) STR(i int) antlr.TerminalNode {
	return s.GetToken(tronParserSTR, i)
}

func (s *FieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tronParserCOMMA)
}

func (s *FieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tronParserCOMMA, i)
}

func (s *FieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterFieldNames(s)
	}
}

func (s *FieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitFieldNames(s)
	}
}

func (p *tronParser) FieldNames() (localctx IFieldNamesContext) {
	localctx = NewFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tronParserRULE_fieldNames)
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
		p.SetState(244)
		p.Match(tronParserSTR)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(245)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(246)
			p.Match(tronParserSTR)
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITyperContext is an interface to support dynamic dispatch.
type ITyperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyperContext differentiates from other interfaces.
	IsTyperContext()
}

type TyperContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyperContext() *TyperContext {
	var p = new(TyperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_typer
	return p
}

func (*TyperContext) IsTyperContext() {}

func NewTyperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TyperContext {
	var p = new(TyperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_typer

	return p
}

func (s *TyperContext) GetParser() antlr.Parser { return s.parser }

func (s *TyperContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *TyperContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *TyperContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *TyperContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *TyperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TyperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TyperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterTyper(s)
	}
}

func (s *TyperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitTyper(s)
	}
}

func (p *tronParser) Typer() (localctx ITyperContext) {
	localctx = NewTyperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tronParserRULE_typer)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDOT {
		{
			p.SetState(252)
			p.Match(tronParserDOT)
		}

	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(255)
				p.Match(tronParserID)
			}
			{
				p.SetState(256)
				p.Match(tronParserDOT)
			}

		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	{
		p.SetState(262)
		p.Match(tronParserID)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Typer() ITyperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyperContext)
}

func (s *FieldContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *FieldContext) INT() antlr.TerminalNode {
	return s.GetToken(tronParserINT, 0)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *FieldContext) FieldRepeat() IFieldRepeatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldRepeatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldRepeatContext)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *tronParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tronParserRULE_field)
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
	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(264)
			p.FieldRepeat()
		}

	}
	{
		p.SetState(267)
		p.Typer()
	}
	{
		p.SetState(268)
		p.Match(tronParserID)
	}
	{
		p.SetState(269)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(270)
		p.Match(tronParserINT)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(271)
			p.FieldOptions()
		}

	}
	{
		p.SetState(274)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IFieldRepeatContext is an interface to support dynamic dispatch.
type IFieldRepeatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRepeated returns the repeated token.
	GetRepeated() antlr.Token

	// SetRepeated sets the repeated token.
	SetRepeated(antlr.Token)

	// IsFieldRepeatContext differentiates from other interfaces.
	IsFieldRepeatContext()
}

type FieldRepeatContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	repeated antlr.Token
}

func NewEmptyFieldRepeatContext() *FieldRepeatContext {
	var p = new(FieldRepeatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_fieldRepeat
	return p
}

func (*FieldRepeatContext) IsFieldRepeatContext() {}

func NewFieldRepeatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldRepeatContext {
	var p = new(FieldRepeatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_fieldRepeat

	return p
}

func (s *FieldRepeatContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldRepeatContext) GetRepeated() antlr.Token { return s.repeated }

func (s *FieldRepeatContext) SetRepeated(v antlr.Token) { s.repeated = v }

func (s *FieldRepeatContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *FieldRepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldRepeatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldRepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterFieldRepeat(s)
	}
}

func (s *FieldRepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitFieldRepeat(s)
	}
}

func (p *tronParser) FieldRepeat() (localctx IFieldRepeatContext) {
	localctx = NewFieldRepeatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tronParserRULE_fieldRepeat)

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
		p.SetState(276)

		var _m = p.Match(tronParserID)

		localctx.(*FieldRepeatContext).repeated = _m
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOf returns the of token.
	GetOf() antlr.Token

	// SetOf sets the of token.
	SetOf(antlr.Token)

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	of     antlr.Token
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) GetOf() antlr.Token { return s.of }

func (s *OneofContext) SetOf(v antlr.Token) { s.of = v }

func (s *OneofContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *OneofContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *OneofContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *OneofContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem())
	var tst = make([]IOneofFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofFieldContext)
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOneof(s)
	}
}

func (p *tronParser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tronParserRULE_oneof)
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
		p.SetState(278)

		var _m = p.Match(tronParserID)

		localctx.(*OneofContext).of = _m
	}
	{
		p.SetState(279)
		p.Match(tronParserID)
	}
	{
		p.SetState(280)
		p.Match(tronParserLCUR)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tronParserSEMI)|(1<<tronParserDOT)|(1<<tronParserID))) != 0 {
		p.SetState(283)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tronParserDOT, tronParserID:
			{
				p.SetState(281)
				p.OneofField()
			}

		case tronParserSEMI:
			{
				p.SetState(282)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(288)
		p.Match(tronParserRCUR)
	}

	return localctx
}

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) Typer() ITyperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyperContext)
}

func (s *OneofFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *OneofFieldContext) INT() antlr.TerminalNode {
	return s.GetToken(tronParserINT, 0)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (p *tronParser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tronParserRULE_oneofField)
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
		p.SetState(290)
		p.Typer()
	}
	{
		p.SetState(291)
		p.Match(tronParserID)
	}
	{
		p.SetState(292)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(293)
		p.Match(tronParserINT)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(294)
			p.FieldOptions()
		}

	}
	{
		p.SetState(297)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMp returns the mp token.
	GetMp() antlr.Token

	// GetKeyType returns the keyType token.
	GetKeyType() antlr.Token

	// SetMp sets the mp token.
	SetMp(antlr.Token)

	// SetKeyType sets the keyType token.
	SetKeyType(antlr.Token)

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	mp      antlr.Token
	keyType antlr.Token
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) GetMp() antlr.Token { return s.mp }

func (s *MapFieldContext) GetKeyType() antlr.Token { return s.keyType }

func (s *MapFieldContext) SetMp(v antlr.Token) { s.mp = v }

func (s *MapFieldContext) SetKeyType(v antlr.Token) { s.keyType = v }

func (s *MapFieldContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(tronParserLCHEVR, 0)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tronParserCOMMA, 0)
}

func (s *MapFieldContext) Typer() ITyperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyperContext)
}

func (s *MapFieldContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(tronParserRCHEVR, 0)
}

func (s *MapFieldContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *MapFieldContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *MapFieldContext) INT() antlr.TerminalNode {
	return s.GetToken(tronParserINT, 0)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *tronParser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tronParserRULE_mapField)
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
		p.SetState(299)

		var _m = p.Match(tronParserID)

		localctx.(*MapFieldContext).mp = _m
	}
	{
		p.SetState(300)
		p.Match(tronParserLCHEVR)
	}
	{
		p.SetState(301)

		var _m = p.Match(tronParserID)

		localctx.(*MapFieldContext).keyType = _m
	}
	{
		p.SetState(302)
		p.Match(tronParserCOMMA)
	}
	{
		p.SetState(303)
		p.Typer()
	}
	{
		p.SetState(304)
		p.Match(tronParserRCHEVR)
	}
	{
		p.SetState(305)
		p.Match(tronParserID)
	}
	{
		p.SetState(306)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(307)
		p.Match(tronParserINT)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(308)
			p.FieldOptions()
		}

	}
	{
		p.SetState(311)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *MessageTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *tronParser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tronParserRULE_messageType)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDOT {
		{
			p.SetState(313)
			p.Match(tronParserDOT)
		}

	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(316)
				p.Match(tronParserID)
			}
			{
				p.SetState(317)
				p.Match(tronParserDOT)
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}
	{
		p.SetState(323)
		p.Match(tronParserID)
	}

	return localctx
}

// IMessageOrEnumTypeContext is an interface to support dynamic dispatch.
type IMessageOrEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageOrEnumTypeContext differentiates from other interfaces.
	IsMessageOrEnumTypeContext()
}

type MessageOrEnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumTypeContext() *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_messageOrEnumType
	return p
}

func (*MessageOrEnumTypeContext) IsMessageOrEnumTypeContext() {}

func NewMessageOrEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_messageOrEnumType

	return p
}

func (s *MessageOrEnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *MessageOrEnumTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *MessageOrEnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *MessageOrEnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *MessageOrEnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterMessageOrEnumType(s)
	}
}

func (s *MessageOrEnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitMessageOrEnumType(s)
	}
}

func (p *tronParser) MessageOrEnumType() (localctx IMessageOrEnumTypeContext) {
	localctx = NewMessageOrEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, tronParserRULE_messageOrEnumType)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDOT {
		{
			p.SetState(325)
			p.Match(tronParserDOT)
		}

	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(328)
				p.Match(tronParserID)
			}
			{
				p.SetState(329)
				p.Match(tronParserDOT)
			}

		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}
	{
		p.SetState(335)
		p.Match(tronParserID)
	}

	return localctx
}

// IEmptyStatementContext is an interface to support dynamic dispatch.
type IEmptyStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatementContext differentiates from other interfaces.
	IsEmptyStatementContext()
}

type EmptyStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatementContext() *EmptyStatementContext {
	var p = new(EmptyStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_emptyStatement
	return p
}

func (*EmptyStatementContext) IsEmptyStatementContext() {}

func NewEmptyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatementContext {
	var p = new(EmptyStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_emptyStatement

	return p
}

func (s *EmptyStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *EmptyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterEmptyStatement(s)
	}
}

func (s *EmptyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitEmptyStatement(s)
	}
}

func (p *tronParser) EmptyStatement() (localctx IEmptyStatementContext) {
	localctx = NewEmptyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, tronParserRULE_emptyStatement)

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
		p.SetState(337)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpt returns the opt token.
	GetOpt() antlr.Token

	// SetOpt sets the opt token.
	SetOpt(antlr.Token)

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	opt    antlr.Token
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) GetOpt() antlr.Token { return s.opt }

func (s *OptionContext) SetOpt(v antlr.Token) { s.opt = v }

func (s *OptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *OptionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *OptionContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *OptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *tronParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, tronParserRULE_option)

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
		p.SetState(339)

		var _m = p.Match(tronParserID)

		localctx.(*OptionContext).opt = _m
	}
	{
		p.SetState(340)
		p.OptionName()
	}
	{
		p.SetState(341)
		p.Match(tronParserEQ)
	}

	{
		p.SetState(342)
		p.Constant()
	}

	{
		p.SetState(343)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) LSB() antlr.TerminalNode {
	return s.GetToken(tronParserLSB, 0)
}

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem())
	var tst = make([]IFieldOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldOptionContext)
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) RSB() antlr.TerminalNode {
	return s.GetToken(tronParserRSB, 0)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tronParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tronParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (p *tronParser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, tronParserRULE_fieldOptions)
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
		p.SetState(345)
		p.Match(tronParserLSB)
	}
	{
		p.SetState(346)
		p.FieldOption()
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(347)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(348)
			p.FieldOption()
		}

		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(354)
		p.Match(tronParserRSB)
	}

	return localctx
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (p *tronParser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, tronParserRULE_fieldOption)

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
		p.SetState(356)
		p.OptionName()
	}
	{
		p.SetState(357)
		p.Match(tronParserEQ)
	}

	{
		p.SetState(358)
		p.Constant()
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *OptionNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *OptionNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tronParserLPAREN, 0)
}

func (s *OptionNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tronParserRPAREN, 0)
}

func (s *OptionNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *OptionNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *tronParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, tronParserRULE_optionName)
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
	p.SetState(371)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserID:
		{
			p.SetState(360)
			p.Match(tronParserID)
		}

	case tronParserLPAREN:
		{
			p.SetState(361)
			p.Match(tronParserLPAREN)
		}
		{
			p.SetState(362)
			p.Match(tronParserID)
		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserDOT {
			{
				p.SetState(363)
				p.Match(tronParserDOT)
			}
			{
				p.SetState(364)
				p.Match(tronParserID)
			}

			p.SetState(369)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(370)
			p.Match(tronParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserDOT {
		{
			p.SetState(373)
			p.Match(tronParserDOT)
		}
		{
			p.SetState(374)
			p.Match(tronParserID)
		}

		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetId() antlr.Token { return s.id }

func (s *ConstantContext) SetId(v antlr.Token) { s.id = v }

func (s *ConstantContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tronParserID)
}

func (s *ConstantContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tronParserID, i)
}

func (s *ConstantContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *ConstantContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(tronParserINT, 0)
}

func (s *ConstantContext) DASH() antlr.TerminalNode {
	return s.GetToken(tronParserDASH, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(tronParserPLUS, 0)
}

func (s *ConstantContext) FLT() antlr.TerminalNode {
	return s.GetToken(tronParserFLT, 0)
}

func (s *ConstantContext) STR() antlr.TerminalNode {
	return s.GetToken(tronParserSTR, 0)
}

func (s *ConstantContext) ConstantObj() IConstantObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantObjContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *tronParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, tronParserRULE_constant)
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

	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Match(tronParserID)
		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tronParserDOT {
			{
				p.SetState(381)
				p.Match(tronParserDOT)
			}
			{
				p.SetState(382)
				p.Match(tronParserID)
			}

			p.SetState(385)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)

			var _m = p.Match(tronParserID)

			localctx.(*ConstantContext).id = _m
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tronParserDASH || _la == tronParserPLUS {
			{
				p.SetState(388)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tronParserDASH || _la == tronParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(391)
			p.Match(tronParserINT)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tronParserDASH || _la == tronParserPLUS {
			{
				p.SetState(392)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tronParserDASH || _la == tronParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(395)
			p.Match(tronParserFLT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(396)
			p.Match(tronParserSTR)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(397)
			p.ConstantObj()
		}

	}

	return localctx
}

// IConstantObjContext is an interface to support dynamic dispatch.
type IConstantObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantObjContext differentiates from other interfaces.
	IsConstantObjContext()
}

type ConstantObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantObjContext() *ConstantObjContext {
	var p = new(ConstantObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_constantObj
	return p
}

func (*ConstantObjContext) IsConstantObjContext() {}

func NewConstantObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantObjContext {
	var p = new(ConstantObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_constantObj

	return p
}

func (s *ConstantObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantObjContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *ConstantObjContext) AllConstantObjElem() []IConstantObjElemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantObjElemContext)(nil)).Elem())
	var tst = make([]IConstantObjElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantObjElemContext)
		}
	}

	return tst
}

func (s *ConstantObjContext) ConstantObjElem(i int) IConstantObjElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantObjElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantObjElemContext)
}

func (s *ConstantObjContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *ConstantObjContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tronParserCOMMA)
}

func (s *ConstantObjContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tronParserCOMMA, i)
}

func (s *ConstantObjContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(tronParserSEMI)
}

func (s *ConstantObjContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, i)
}

func (s *ConstantObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterConstantObj(s)
	}
}

func (s *ConstantObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitConstantObj(s)
	}
}

func (p *tronParser) ConstantObj() (localctx IConstantObjContext) {
	localctx = NewConstantObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, tronParserRULE_constantObj)
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
		p.SetState(400)
		p.Match(tronParserLCUR)
	}
	{
		p.SetState(401)
		p.ConstantObjElem()
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tronParserSEMI)|(1<<tronParserCOMMA)|(1<<tronParserID))) != 0 {
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tronParserSEMI || _la == tronParserCOMMA {
			{
				p.SetState(402)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tronParserSEMI || _la == tronParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(405)
			p.ConstantObjElem()
		}

		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(411)
		p.Match(tronParserRCUR)
	}

	return localctx
}

// IConstantObjElemContext is an interface to support dynamic dispatch.
type IConstantObjElemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantObjElemContext differentiates from other interfaces.
	IsConstantObjElemContext()
}

type ConstantObjElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantObjElemContext() *ConstantObjElemContext {
	var p = new(ConstantObjElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_constantObjElem
	return p
}

func (*ConstantObjElemContext) IsConstantObjElemContext() {}

func NewConstantObjElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantObjElemContext {
	var p = new(ConstantObjElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_constantObjElem

	return p
}

func (s *ConstantObjElemContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantObjElemContext) CopyFrom(ctx *ConstantObjElemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantObjElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantObjElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PronSTRContext struct {
	*ConstantObjElemContext
	id antlr.Token
}

func NewPronSTRContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PronSTRContext {
	var p = new(PronSTRContext)

	p.ConstantObjElemContext = NewEmptyConstantObjElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantObjElemContext))

	return p
}

func (s *PronSTRContext) GetId() antlr.Token { return s.id }

func (s *PronSTRContext) SetId(v antlr.Token) { s.id = v }

func (s *PronSTRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronSTRContext) COLON() antlr.TerminalNode {
	return s.GetToken(tronParserCOLON, 0)
}

func (s *PronSTRContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PronSTRContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *PronSTRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterPronSTR(s)
	}
}

func (s *PronSTRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitPronSTR(s)
	}
}

type PronOBJContext struct {
	*ConstantObjElemContext
	id antlr.Token
}

func NewPronOBJContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PronOBJContext {
	var p = new(PronOBJContext)

	p.ConstantObjElemContext = NewEmptyConstantObjElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantObjElemContext))

	return p
}

func (s *PronOBJContext) GetId() antlr.Token { return s.id }

func (s *PronOBJContext) SetId(v antlr.Token) { s.id = v }

func (s *PronOBJContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronOBJContext) COLON() antlr.TerminalNode {
	return s.GetToken(tronParserCOLON, 0)
}

func (s *PronOBJContext) ConstantObj() IConstantObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantObjContext)
}

func (s *PronOBJContext) ID() antlr.TerminalNode {
	return s.GetToken(tronParserID, 0)
}

func (s *PronOBJContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterPronOBJ(s)
	}
}

func (s *PronOBJContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitPronOBJ(s)
	}
}

func (p *tronParser) ConstantObjElem() (localctx IConstantObjElemContext) {
	localctx = NewConstantObjElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, tronParserRULE_constantObjElem)

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

	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPronSTRContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(413)

			var _m = p.Match(tronParserID)

			localctx.(*PronSTRContext).id = _m
		}
		{
			p.SetState(414)
			p.Match(tronParserCOLON)
		}
		{
			p.SetState(415)
			p.Constant()
		}

	case 2:
		localctx = NewPronOBJContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)

			var _m = p.Match(tronParserID)

			localctx.(*PronOBJContext).id = _m
		}
		{
			p.SetState(417)
			p.Match(tronParserCOLON)
		}
		{
			p.SetState(418)
			p.ConstantObj()
		}

	}

	return localctx
}
