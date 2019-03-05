// Code generated from tron_parser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tron_parser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 438,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 2, 7,
	2, 99, 10, 2, 12, 2, 14, 2, 102, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 114, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 129, 10, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 143, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 152, 10,
	9, 3, 9, 3, 9, 7, 9, 156, 10, 9, 12, 9, 14, 9, 159, 11, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 7, 10, 165, 10, 10, 12, 10, 14, 10, 168, 11, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 180,
	10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 186, 10, 12, 13, 12, 14, 12,
	187, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 195, 10, 13, 3, 14, 3, 14,
	3, 14, 5, 14, 200, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 207,
	10, 14, 12, 14, 14, 14, 210, 11, 14, 3, 14, 3, 14, 5, 14, 214, 10, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7,
	16, 226, 10, 16, 12, 16, 14, 16, 229, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 5, 17, 236, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 7, 18, 246, 10, 18, 12, 18, 14, 18, 249, 11, 18, 3, 18, 3, 18,
	5, 18, 253, 10, 18, 3, 19, 3, 19, 5, 19, 257, 10, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 5, 20, 265, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 7, 21, 272, 10, 21, 12, 21, 14, 21, 275, 11, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 281, 10, 22, 3, 23, 3, 23, 3, 23, 7, 23, 286, 10, 23,
	12, 23, 14, 23, 289, 11, 23, 3, 24, 3, 24, 5, 24, 293, 10, 24, 3, 25, 3,
	25, 3, 26, 5, 26, 298, 10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26,
	305, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 313, 10,
	27, 12, 27, 14, 27, 316, 11, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 329, 10, 29, 12, 29, 14,
	29, 332, 11, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30,
	341, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 5, 31, 355, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 7, 33, 364, 10, 33, 12, 33, 14, 33, 367, 11, 33, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 5, 42, 386, 10, 42, 3, 42, 3,
	42, 7, 42, 390, 10, 42, 12, 42, 14, 42, 393, 11, 42, 3, 42, 3, 42, 3, 43,
	5, 43, 398, 10, 43, 3, 43, 3, 43, 7, 43, 402, 10, 43, 12, 43, 14, 43, 405,
	11, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 5, 46,
	415, 10, 46, 3, 46, 3, 46, 5, 46, 419, 10, 46, 3, 46, 3, 46, 3, 46, 5,
	46, 424, 10, 46, 5, 46, 426, 10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 48, 3, 48, 5, 48, 436, 10, 48, 3, 48, 2, 2, 49, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
	84, 86, 88, 90, 92, 94, 2, 7, 3, 2, 24, 25, 4, 2, 42, 49, 51, 57, 6, 2,
	42, 42, 45, 46, 48, 49, 51, 57, 3, 2, 40, 41, 3, 2, 18, 19, 2, 447, 2,
	96, 3, 2, 2, 2, 4, 113, 3, 2, 2, 2, 6, 115, 3, 2, 2, 2, 8, 120, 3, 2, 2,
	2, 10, 126, 3, 2, 2, 2, 12, 133, 3, 2, 2, 2, 14, 137, 3, 2, 2, 2, 16, 151,
	3, 2, 2, 2, 18, 160, 3, 2, 2, 2, 20, 179, 3, 2, 2, 2, 22, 181, 3, 2, 2,
	2, 24, 194, 3, 2, 2, 2, 26, 196, 3, 2, 2, 2, 28, 217, 3, 2, 2, 2, 30, 221,
	3, 2, 2, 2, 32, 235, 3, 2, 2, 2, 34, 237, 3, 2, 2, 2, 36, 254, 3, 2, 2,
	2, 38, 261, 3, 2, 2, 2, 40, 268, 3, 2, 2, 2, 42, 280, 3, 2, 2, 2, 44, 282,
	3, 2, 2, 2, 46, 292, 3, 2, 2, 2, 48, 294, 3, 2, 2, 2, 50, 297, 3, 2, 2,
	2, 52, 308, 3, 2, 2, 2, 54, 319, 3, 2, 2, 2, 56, 323, 3, 2, 2, 2, 58, 335,
	3, 2, 2, 2, 60, 344, 3, 2, 2, 2, 62, 358, 3, 2, 2, 2, 64, 360, 3, 2, 2,
	2, 66, 368, 3, 2, 2, 2, 68, 370, 3, 2, 2, 2, 70, 372, 3, 2, 2, 2, 72, 374,
	3, 2, 2, 2, 74, 376, 3, 2, 2, 2, 76, 378, 3, 2, 2, 2, 78, 380, 3, 2, 2,
	2, 80, 382, 3, 2, 2, 2, 82, 385, 3, 2, 2, 2, 84, 397, 3, 2, 2, 2, 86, 408,
	3, 2, 2, 2, 88, 410, 3, 2, 2, 2, 90, 425, 3, 2, 2, 2, 92, 427, 3, 2, 2,
	2, 94, 435, 3, 2, 2, 2, 96, 100, 5, 6, 4, 2, 97, 99, 5, 4, 3, 2, 98, 97,
	3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 7, 2, 2, 3, 104,
	3, 3, 2, 2, 2, 105, 114, 5, 10, 6, 2, 106, 114, 5, 12, 7, 2, 107, 114,
	5, 8, 5, 2, 108, 114, 5, 14, 8, 2, 109, 114, 5, 18, 10, 2, 110, 114, 5,
	22, 12, 2, 111, 114, 5, 30, 16, 2, 112, 114, 5, 88, 45, 2, 113, 105, 3,
	2, 2, 2, 113, 106, 3, 2, 2, 2, 113, 107, 3, 2, 2, 2, 113, 108, 3, 2, 2,
	2, 113, 109, 3, 2, 2, 2, 113, 110, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113,
	112, 3, 2, 2, 2, 114, 5, 3, 2, 2, 2, 115, 116, 7, 21, 2, 2, 116, 117, 7,
	5, 2, 2, 117, 118, 7, 20, 2, 2, 118, 119, 7, 8, 2, 2, 119, 7, 3, 2, 2,
	2, 120, 121, 7, 22, 2, 2, 121, 122, 5, 64, 33, 2, 122, 123, 7, 3, 2, 2,
	123, 124, 5, 50, 26, 2, 124, 125, 7, 4, 2, 2, 125, 9, 3, 2, 2, 2, 126,
	128, 7, 23, 2, 2, 127, 129, 9, 2, 2, 2, 128, 127, 3, 2, 2, 2, 128, 129,
	3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 7, 58, 2, 2, 131, 132, 7, 8,
	2, 2, 132, 11, 3, 2, 2, 2, 133, 134, 7, 27, 2, 2, 134, 135, 5, 64, 33,
	2, 135, 136, 7, 8, 2, 2, 136, 13, 3, 2, 2, 2, 137, 138, 7, 28, 2, 2, 138,
	139, 5, 16, 9, 2, 139, 142, 7, 5, 2, 2, 140, 143, 5, 90, 46, 2, 141, 143,
	5, 92, 47, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 144, 3,
	2, 2, 2, 144, 145, 7, 8, 2, 2, 145, 15, 3, 2, 2, 2, 146, 152, 7, 59, 2,
	2, 147, 148, 7, 10, 2, 2, 148, 149, 5, 64, 33, 2, 149, 150, 7, 11, 2, 2,
	150, 152, 3, 2, 2, 2, 151, 146, 3, 2, 2, 2, 151, 147, 3, 2, 2, 2, 152,
	157, 3, 2, 2, 2, 153, 154, 7, 12, 2, 2, 154, 156, 7, 59, 2, 2, 155, 153,
	3, 2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2,
	2, 2, 158, 17, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 161, 7, 50, 2, 2,
	161, 162, 5, 66, 34, 2, 162, 166, 7, 3, 2, 2, 163, 165, 5, 20, 11, 2, 164,
	163, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167,
	3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 170, 7, 4,
	2, 2, 170, 19, 3, 2, 2, 2, 171, 180, 5, 50, 26, 2, 172, 180, 5, 22, 12,
	2, 173, 180, 5, 18, 10, 2, 174, 180, 5, 14, 8, 2, 175, 180, 5, 56, 29,
	2, 176, 180, 5, 60, 31, 2, 177, 180, 5, 38, 20, 2, 178, 180, 5, 88, 45,
	2, 179, 171, 3, 2, 2, 2, 179, 172, 3, 2, 2, 2, 179, 173, 3, 2, 2, 2, 179,
	174, 3, 2, 2, 2, 179, 175, 3, 2, 2, 2, 179, 176, 3, 2, 2, 2, 179, 177,
	3, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 21, 3, 2, 2, 2, 181, 182, 7, 29,
	2, 2, 182, 183, 5, 68, 35, 2, 183, 185, 7, 3, 2, 2, 184, 186, 5, 24, 13,
	2, 185, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 7, 4, 2, 2, 190, 23, 3,
	2, 2, 2, 191, 195, 5, 14, 8, 2, 192, 195, 5, 26, 14, 2, 193, 195, 5, 88,
	45, 2, 194, 191, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 193, 3, 2, 2, 2,
	195, 25, 3, 2, 2, 2, 196, 197, 7, 59, 2, 2, 197, 199, 7, 5, 2, 2, 198,
	200, 7, 18, 2, 2, 199, 198, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201,
	3, 2, 2, 2, 201, 213, 7, 60, 2, 2, 202, 203, 7, 13, 2, 2, 203, 208, 5,
	28, 15, 2, 204, 205, 7, 15, 2, 2, 205, 207, 5, 28, 15, 2, 206, 204, 3,
	2, 2, 2, 207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2,
	2, 209, 211, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 212, 7, 14, 2, 2, 212,
	214, 3, 2, 2, 2, 213, 202, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215,
	3, 2, 2, 2, 215, 216, 7, 8, 2, 2, 216, 27, 3, 2, 2, 2, 217, 218, 5, 16,
	9, 2, 218, 219, 7, 5, 2, 2, 219, 220, 5, 90, 46, 2, 220, 29, 3, 2, 2, 2,
	221, 222, 7, 30, 2, 2, 222, 223, 5, 78, 40, 2, 223, 227, 7, 3, 2, 2, 224,
	226, 5, 32, 17, 2, 225, 224, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229, 227, 3, 2,
	2, 2, 230, 231, 7, 4, 2, 2, 231, 31, 3, 2, 2, 2, 232, 236, 5, 14, 8, 2,
	233, 236, 5, 34, 18, 2, 234, 236, 5, 88, 45, 2, 235, 232, 3, 2, 2, 2, 235,
	233, 3, 2, 2, 2, 235, 234, 3, 2, 2, 2, 236, 33, 3, 2, 2, 2, 237, 238, 7,
	31, 2, 2, 238, 239, 5, 80, 41, 2, 239, 240, 5, 36, 19, 2, 240, 241, 7,
	26, 2, 2, 241, 252, 5, 36, 19, 2, 242, 247, 7, 3, 2, 2, 243, 246, 5, 14,
	8, 2, 244, 246, 5, 88, 45, 2, 245, 243, 3, 2, 2, 2, 245, 244, 3, 2, 2,
	2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	250, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 253, 7, 4, 2, 2, 251, 253,
	7, 8, 2, 2, 252, 242, 3, 2, 2, 2, 252, 251, 3, 2, 2, 2, 253, 35, 3, 2,
	2, 2, 254, 256, 7, 10, 2, 2, 255, 257, 7, 32, 2, 2, 256, 255, 3, 2, 2,
	2, 256, 257, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259, 5, 82, 42, 2,
	259, 260, 7, 11, 2, 2, 260, 37, 3, 2, 2, 2, 261, 264, 7, 33, 2, 2, 262,
	265, 5, 40, 21, 2, 263, 265, 5, 44, 23, 2, 264, 262, 3, 2, 2, 2, 264, 263,
	3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 7, 8, 2, 2, 267, 39, 3, 2,
	2, 2, 268, 273, 5, 42, 22, 2, 269, 270, 7, 15, 2, 2, 270, 272, 5, 42, 22,
	2, 271, 269, 3, 2, 2, 2, 272, 275, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2, 273,
	274, 3, 2, 2, 2, 274, 41, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 276, 281, 7,
	60, 2, 2, 277, 278, 7, 60, 2, 2, 278, 279, 7, 34, 2, 2, 279, 281, 7, 60,
	2, 2, 280, 276, 3, 2, 2, 2, 280, 277, 3, 2, 2, 2, 281, 43, 3, 2, 2, 2,
	282, 287, 7, 58, 2, 2, 283, 284, 7, 15, 2, 2, 284, 286, 7, 58, 2, 2, 285,
	283, 3, 2, 2, 2, 286, 289, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288,
	3, 2, 2, 2, 288, 45, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 290, 293, 5, 48,
	25, 2, 291, 293, 5, 84, 43, 2, 292, 290, 3, 2, 2, 2, 292, 291, 3, 2, 2,
	2, 293, 47, 3, 2, 2, 2, 294, 295, 9, 3, 2, 2, 295, 49, 3, 2, 2, 2, 296,
	298, 7, 35, 2, 2, 297, 296, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 299,
	3, 2, 2, 2, 299, 300, 5, 46, 24, 2, 300, 301, 5, 72, 37, 2, 301, 302, 7,
	5, 2, 2, 302, 304, 7, 60, 2, 2, 303, 305, 5, 52, 27, 2, 304, 303, 3, 2,
	2, 2, 304, 305, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 307, 7, 8, 2, 2,
	307, 51, 3, 2, 2, 2, 308, 309, 7, 13, 2, 2, 309, 314, 5, 54, 28, 2, 310,
	311, 7, 15, 2, 2, 311, 313, 5, 54, 28, 2, 312, 310, 3, 2, 2, 2, 313, 316,
	3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 317, 3, 2,
	2, 2, 316, 314, 3, 2, 2, 2, 317, 318, 7, 14, 2, 2, 318, 53, 3, 2, 2, 2,
	319, 320, 5, 16, 9, 2, 320, 321, 7, 5, 2, 2, 321, 322, 5, 90, 46, 2, 322,
	55, 3, 2, 2, 2, 323, 324, 7, 36, 2, 2, 324, 325, 5, 74, 38, 2, 325, 330,
	7, 3, 2, 2, 326, 329, 5, 58, 30, 2, 327, 329, 5, 88, 45, 2, 328, 326, 3,
	2, 2, 2, 328, 327, 3, 2, 2, 2, 329, 332, 3, 2, 2, 2, 330, 328, 3, 2, 2,
	2, 330, 331, 3, 2, 2, 2, 331, 333, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 333,
	334, 7, 4, 2, 2, 334, 57, 3, 2, 2, 2, 335, 336, 5, 46, 24, 2, 336, 337,
	5, 72, 37, 2, 337, 338, 7, 5, 2, 2, 338, 340, 7, 60, 2, 2, 339, 341, 5,
	52, 27, 2, 340, 339, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 342, 3, 2,
	2, 2, 342, 343, 7, 8, 2, 2, 343, 59, 3, 2, 2, 2, 344, 345, 7, 37, 2, 2,
	345, 346, 7, 16, 2, 2, 346, 347, 5, 62, 32, 2, 347, 348, 7, 15, 2, 2, 348,
	349, 5, 46, 24, 2, 349, 350, 7, 17, 2, 2, 350, 351, 5, 76, 39, 2, 351,
	352, 7, 5, 2, 2, 352, 354, 7, 60, 2, 2, 353, 355, 5, 52, 27, 2, 354, 353,
	3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 7, 8,
	2, 2, 357, 61, 3, 2, 2, 2, 358, 359, 9, 4, 2, 2, 359, 63, 3, 2, 2, 2, 360,
	365, 7, 59, 2, 2, 361, 362, 7, 12, 2, 2, 362, 364, 7, 59, 2, 2, 363, 361,
	3, 2, 2, 2, 364, 367, 3, 2, 2, 2, 365, 363, 3, 2, 2, 2, 365, 366, 3, 2,
	2, 2, 366, 65, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 368, 369, 7, 59, 2, 2,
	369, 67, 3, 2, 2, 2, 370, 371, 7, 59, 2, 2, 371, 69, 3, 2, 2, 2, 372, 373,
	7, 59, 2, 2, 373, 71, 3, 2, 2, 2, 374, 375, 7, 59, 2, 2, 375, 73, 3, 2,
	2, 2, 376, 377, 7, 59, 2, 2, 377, 75, 3, 2, 2, 2, 378, 379, 7, 59, 2, 2,
	379, 77, 3, 2, 2, 2, 380, 381, 7, 59, 2, 2, 381, 79, 3, 2, 2, 2, 382, 383,
	7, 59, 2, 2, 383, 81, 3, 2, 2, 2, 384, 386, 7, 12, 2, 2, 385, 384, 3, 2,
	2, 2, 385, 386, 3, 2, 2, 2, 386, 391, 3, 2, 2, 2, 387, 388, 7, 59, 2, 2,
	388, 390, 7, 12, 2, 2, 389, 387, 3, 2, 2, 2, 390, 393, 3, 2, 2, 2, 391,
	389, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 394, 3, 2, 2, 2, 393, 391,
	3, 2, 2, 2, 394, 395, 5, 66, 34, 2, 395, 83, 3, 2, 2, 2, 396, 398, 7, 12,
	2, 2, 397, 396, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 403, 3, 2, 2, 2,
	399, 400, 7, 59, 2, 2, 400, 402, 7, 12, 2, 2, 401, 399, 3, 2, 2, 2, 402,
	405, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 406,
	3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 406, 407, 5, 70, 36, 2, 407, 85, 3, 2,
	2, 2, 408, 409, 9, 5, 2, 2, 409, 87, 3, 2, 2, 2, 410, 411, 7, 8, 2, 2,
	411, 89, 3, 2, 2, 2, 412, 426, 5, 64, 33, 2, 413, 415, 9, 6, 2, 2, 414,
	413, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 426,
	7, 60, 2, 2, 417, 419, 9, 6, 2, 2, 418, 417, 3, 2, 2, 2, 418, 419, 3, 2,
	2, 2, 419, 420, 3, 2, 2, 2, 420, 426, 7, 61, 2, 2, 421, 424, 7, 58, 2,
	2, 422, 424, 5, 86, 44, 2, 423, 421, 3, 2, 2, 2, 423, 422, 3, 2, 2, 2,
	424, 426, 3, 2, 2, 2, 425, 412, 3, 2, 2, 2, 425, 414, 3, 2, 2, 2, 425,
	418, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 91, 3, 2, 2, 2, 427, 428, 7,
	3, 2, 2, 428, 429, 7, 59, 2, 2, 429, 430, 7, 9, 2, 2, 430, 431, 5, 94,
	48, 2, 431, 432, 7, 4, 2, 2, 432, 93, 3, 2, 2, 2, 433, 436, 7, 58, 2, 2,
	434, 436, 5, 92, 47, 2, 435, 433, 3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436,
	95, 3, 2, 2, 2, 43, 100, 113, 128, 142, 151, 157, 166, 179, 187, 194, 199,
	208, 213, 227, 235, 245, 247, 252, 256, 264, 273, 280, 287, 292, 297, 304,
	314, 328, 330, 340, 354, 365, 385, 391, 397, 403, 414, 418, 423, 425, 435,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'='", "'\"'", "'''", "';'", "':'", "'('", "')'", "'.'",
	"'['", "']'", "','", "'<'", "'>'", "'-'", "'+'", "", "'syntax'", "'extend'",
	"'import'", "'weak'", "'public'", "'returns'", "'package'", "'option'",
	"'enum'", "'service'", "'rpc'", "'stream'", "'reserved'", "'to'", "'repeated'",
	"'oneof'", "'map'", "'inf'", "'nan'", "'true'", "'false'", "'bool'", "'bytes'",
	"'double'", "'fixed32'", "'fixed64'", "'float'", "'int32'", "'int64'",
	"'message'", "'sfixed32'", "'sfixed64'", "'sint32'", "'sint64'", "'string'",
	"'uint32'", "'uint64'",
}
var symbolicNames = []string{
	"", "LCUR", "RCUR", "EQ", "DQ", "SQ", "SEMI", "COLON", "LPAREN", "RPAREN",
	"DOT", "LSB", "RSB", "COMMA", "LCHEVR", "RCHEVR", "DASH", "PLUS", "PROTO3",
	"SYNTAX", "EXTEND", "IMPORT", "WEAK", "PUBLIC", "RETURNS", "PACKAGE", "OPTION",
	"ENUM", "SERVICE", "RPC", "STREAM", "RESERVED", "TO", "REPEATED", "ONEOF",
	"MAP", "INF", "NAN", "TRUE", "FALSE", "BOOL", "BYTES", "DOUBLE", "FIXED32",
	"FIXED64", "FLOAT", "INT32", "INT64", "MESSAGE", "SFIXED32", "SFIXED64",
	"SINT32", "SINT64", "STRING", "UINT32", "UINT64", "StrLit", "Ident", "Int_lit",
	"Float_lit", "WS", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"proto", "top_level_statement", "syntax", "extend", "importStatement",
	"packageStatement", "option", "optionName", "message", "messageBody", "enumDefinition",
	"enumBody", "enumField", "enumValueOption", "service", "serviceBody", "rpc",
	"rpcParam", "reserved", "ranges", "rangee", "fieldNames", "typer", "types",
	"field", "fieldOptions", "fieldOption", "oneof", "oneofField", "mapField",
	"keyType", "fullIdent", "messageName", "enumName", "messageOrEnumName",
	"fieldName", "oneofName", "mapName", "serviceName", "rpcName", "messageType",
	"messageOrEnumType", "bool_lit", "emptyStatement", "constant", "pron",
	"pronElem",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tron_parser struct {
	*antlr.BaseParser
}

func Newtron_parser(input antlr.TokenStream) *tron_parser {
	this := new(tron_parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tron_parser.g4"

	return this
}

// tron_parser tokens.
const (
	tron_parserEOF          = antlr.TokenEOF
	tron_parserLCUR         = 1
	tron_parserRCUR         = 2
	tron_parserEQ           = 3
	tron_parserDQ           = 4
	tron_parserSQ           = 5
	tron_parserSEMI         = 6
	tron_parserCOLON        = 7
	tron_parserLPAREN       = 8
	tron_parserRPAREN       = 9
	tron_parserDOT          = 10
	tron_parserLSB          = 11
	tron_parserRSB          = 12
	tron_parserCOMMA        = 13
	tron_parserLCHEVR       = 14
	tron_parserRCHEVR       = 15
	tron_parserDASH         = 16
	tron_parserPLUS         = 17
	tron_parserPROTO3       = 18
	tron_parserSYNTAX       = 19
	tron_parserEXTEND       = 20
	tron_parserIMPORT       = 21
	tron_parserWEAK         = 22
	tron_parserPUBLIC       = 23
	tron_parserRETURNS      = 24
	tron_parserPACKAGE      = 25
	tron_parserOPTION       = 26
	tron_parserENUM         = 27
	tron_parserSERVICE      = 28
	tron_parserRPC          = 29
	tron_parserSTREAM       = 30
	tron_parserRESERVED     = 31
	tron_parserTO           = 32
	tron_parserREPEATED     = 33
	tron_parserONEOF        = 34
	tron_parserMAP          = 35
	tron_parserINF          = 36
	tron_parserNAN          = 37
	tron_parserTRUE         = 38
	tron_parserFALSE        = 39
	tron_parserBOOL         = 40
	tron_parserBYTES        = 41
	tron_parserDOUBLE       = 42
	tron_parserFIXED32      = 43
	tron_parserFIXED64      = 44
	tron_parserFLOAT        = 45
	tron_parserINT32        = 46
	tron_parserINT64        = 47
	tron_parserMESSAGE      = 48
	tron_parserSFIXED32     = 49
	tron_parserSFIXED64     = 50
	tron_parserSINT32       = 51
	tron_parserSINT64       = 52
	tron_parserSTRING       = 53
	tron_parserUINT32       = 54
	tron_parserUINT64       = 55
	tron_parserStrLit       = 56
	tron_parserIdent        = 57
	tron_parserInt_lit      = 58
	tron_parserFloat_lit    = 59
	tron_parserWS           = 60
	tron_parserCOMMENT      = 61
	tron_parserLINE_COMMENT = 62
)

// tron_parser rules.
const (
	tron_parserRULE_proto               = 0
	tron_parserRULE_top_level_statement = 1
	tron_parserRULE_syntax              = 2
	tron_parserRULE_extend              = 3
	tron_parserRULE_importStatement     = 4
	tron_parserRULE_packageStatement    = 5
	tron_parserRULE_option              = 6
	tron_parserRULE_optionName          = 7
	tron_parserRULE_message             = 8
	tron_parserRULE_messageBody         = 9
	tron_parserRULE_enumDefinition      = 10
	tron_parserRULE_enumBody            = 11
	tron_parserRULE_enumField           = 12
	tron_parserRULE_enumValueOption     = 13
	tron_parserRULE_service             = 14
	tron_parserRULE_serviceBody         = 15
	tron_parserRULE_rpc                 = 16
	tron_parserRULE_rpcParam            = 17
	tron_parserRULE_reserved            = 18
	tron_parserRULE_ranges              = 19
	tron_parserRULE_rangee              = 20
	tron_parserRULE_fieldNames          = 21
	tron_parserRULE_typer               = 22
	tron_parserRULE_types               = 23
	tron_parserRULE_field               = 24
	tron_parserRULE_fieldOptions        = 25
	tron_parserRULE_fieldOption         = 26
	tron_parserRULE_oneof               = 27
	tron_parserRULE_oneofField          = 28
	tron_parserRULE_mapField            = 29
	tron_parserRULE_keyType             = 30
	tron_parserRULE_fullIdent           = 31
	tron_parserRULE_messageName         = 32
	tron_parserRULE_enumName            = 33
	tron_parserRULE_messageOrEnumName   = 34
	tron_parserRULE_fieldName           = 35
	tron_parserRULE_oneofName           = 36
	tron_parserRULE_mapName             = 37
	tron_parserRULE_serviceName         = 38
	tron_parserRULE_rpcName             = 39
	tron_parserRULE_messageType         = 40
	tron_parserRULE_messageOrEnumType   = 41
	tron_parserRULE_bool_lit            = 42
	tron_parserRULE_emptyStatement      = 43
	tron_parserRULE_constant            = 44
	tron_parserRULE_pron                = 45
	tron_parserRULE_pronElem            = 46
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
	p.RuleIndex = tron_parserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(tron_parserEOF, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *tron_parser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tron_parserRULE_proto)
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
		p.SetState(94)
		p.Syntax()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tron_parserSEMI)|(1<<tron_parserEXTEND)|(1<<tron_parserIMPORT)|(1<<tron_parserPACKAGE)|(1<<tron_parserOPTION)|(1<<tron_parserENUM)|(1<<tron_parserSERVICE))) != 0) || _la == tron_parserMESSAGE {
		{
			p.SetState(95)
			p.Top_level_statement()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(tron_parserEOF)
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
	p.RuleIndex = tron_parserRULE_top_level_statement
	return p
}

func (*Top_level_statementContext) IsTop_level_statementContext() {}

func NewTop_level_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_level_statementContext {
	var p = new(Top_level_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_top_level_statement

	return p
}

func (s *Top_level_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Top_level_statementContext) ImportStatement() IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *Top_level_statementContext) PackageStatement() IPackageStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *Top_level_statementContext) Extend() IExtendContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendContext)
}

func (s *Top_level_statementContext) Option() IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *Top_level_statementContext) Message() IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *Top_level_statementContext) EnumDefinition() IEnumDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *Top_level_statementContext) Service() IServiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *Top_level_statementContext) EmptyStatement() IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *Top_level_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_level_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Top_level_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterTop_level_statement(s)
	}
}

func (s *Top_level_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitTop_level_statement(s)
	}
}

func (p *tron_parser) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tron_parserRULE_top_level_statement)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.ImportStatement()
		}

	case tron_parserPACKAGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.PackageStatement()
		}

	case tron_parserEXTEND:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Extend()
		}

	case tron_parserOPTION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Option()
		}

	case tron_parserMESSAGE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Message()
		}

	case tron_parserENUM:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.EnumDefinition()
		}

	case tron_parserSERVICE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Service()
		}

	case tron_parserSEMI:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(110)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = tron_parserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(tron_parserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *SyntaxContext) PROTO3() antlr.TerminalNode {
	return s.GetToken(tron_parserPROTO3, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *tron_parser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tron_parserRULE_syntax)

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
		p.SetState(113)
		p.Match(tron_parserSYNTAX)
	}
	{
		p.SetState(114)
		p.Match(tron_parserEQ)
	}
	{
		p.SetState(115)
		p.Match(tron_parserPROTO3)
	}
	{
		p.SetState(116)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IExtendContext is an interface to support dynamic dispatch.
type IExtendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendContext differentiates from other interfaces.
	IsExtendContext()
}

type ExtendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendContext() *ExtendContext {
	var p = new(ExtendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_extend
	return p
}

func (*ExtendContext) IsExtendContext() {}

func NewExtendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendContext {
	var p = new(ExtendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_extend

	return p
}

func (s *ExtendContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(tron_parserEXTEND, 0)
}

func (s *ExtendContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ExtendContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *ExtendContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExtendContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
}

func (s *ExtendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterExtend(s)
	}
}

func (s *ExtendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitExtend(s)
	}
}

func (p *tron_parser) Extend() (localctx IExtendContext) {
	localctx = NewExtendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tron_parserRULE_extend)

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
		p.SetState(118)
		p.Match(tron_parserEXTEND)
	}
	{
		p.SetState(119)
		p.FullIdent()
	}
	{
		p.SetState(120)
		p.Match(tron_parserLCUR)
	}
	{
		p.SetState(121)
		p.Field()
	}
	{
		p.SetState(122)
		p.Match(tron_parserRCUR)
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(tron_parserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() antlr.TerminalNode {
	return s.GetToken(tron_parserStrLit, 0)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(tron_parserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(tron_parserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *tron_parser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tron_parserRULE_importStatement)
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
		p.SetState(124)
		p.Match(tron_parserIMPORT)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserWEAK || _la == tron_parserPUBLIC {
		{
			p.SetState(125)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tron_parserWEAK || _la == tron_parserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(128)
		p.Match(tron_parserStrLit)
	}
	{
		p.SetState(129)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(tron_parserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *tron_parser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tron_parserRULE_packageStatement)

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
		p.SetState(131)
		p.Match(tron_parserPACKAGE)
	}
	{
		p.SetState(132)
		p.FullIdent()
	}
	{
		p.SetState(133)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) OPTION() antlr.TerminalNode {
	return s.GetToken(tron_parserOPTION, 0)
}

func (s *OptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *OptionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *OptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *tron_parser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tron_parserRULE_option)

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
		p.SetState(135)
		p.Match(tron_parserOPTION)
	}
	{
		p.SetState(136)
		p.OptionName()
	}
	{
		p.SetState(137)
		p.Match(tron_parserEQ)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserDASH, tron_parserPLUS, tron_parserTRUE, tron_parserFALSE, tron_parserStrLit, tron_parserIdent, tron_parserInt_lit, tron_parserFloat_lit:
		{
			p.SetState(138)
			p.Constant()
		}

	case tron_parserLCUR:
		{
			p.SetState(139)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(142)
		p.Match(tron_parserSEMI)
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
	p.RuleIndex = tron_parserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tron_parserIdent)
}

func (s *OptionNameContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, i)
}

func (s *OptionNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tron_parserLPAREN, 0)
}

func (s *OptionNameContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tron_parserRPAREN, 0)
}

func (s *OptionNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tron_parserDOT)
}

func (s *OptionNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserDOT, i)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *tron_parser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tron_parserRULE_optionName)
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
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserIdent:
		{
			p.SetState(144)
			p.Match(tron_parserIdent)
		}

	case tron_parserLPAREN:
		{
			p.SetState(145)
			p.Match(tron_parserLPAREN)
		}
		{
			p.SetState(146)
			p.FullIdent()
		}
		{
			p.SetState(147)
			p.Match(tron_parserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tron_parserDOT {
		{
			p.SetState(151)
			p.Match(tron_parserDOT)
		}
		{
			p.SetState(152)
			p.Match(tron_parserIdent)
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(tron_parserMESSAGE, 0)
}

func (s *MessageContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *MessageContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
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

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *tron_parser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tron_parserRULE_message)
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
		p.SetState(158)
		p.Match(tron_parserMESSAGE)
	}
	{
		p.SetState(159)
		p.MessageName()
	}
	{
		p.SetState(160)
		p.Match(tron_parserLCUR)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tron_parserSEMI)|(1<<tron_parserDOT)|(1<<tron_parserOPTION)|(1<<tron_parserENUM)|(1<<tron_parserRESERVED))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(tron_parserREPEATED-33))|(1<<(tron_parserONEOF-33))|(1<<(tron_parserMAP-33))|(1<<(tron_parserBOOL-33))|(1<<(tron_parserBYTES-33))|(1<<(tron_parserDOUBLE-33))|(1<<(tron_parserFIXED32-33))|(1<<(tron_parserFIXED64-33))|(1<<(tron_parserFLOAT-33))|(1<<(tron_parserINT32-33))|(1<<(tron_parserINT64-33))|(1<<(tron_parserMESSAGE-33))|(1<<(tron_parserSFIXED32-33))|(1<<(tron_parserSFIXED64-33))|(1<<(tron_parserSINT32-33))|(1<<(tron_parserSINT64-33))|(1<<(tron_parserSTRING-33))|(1<<(tron_parserUINT32-33))|(1<<(tron_parserUINT64-33))|(1<<(tron_parserIdent-33)))) != 0) {
		{
			p.SetState(161)
			p.MessageBody()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(tron_parserRCUR)
	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageBodyContext) EnumDefinition() IEnumDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *MessageBodyContext) Message() IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (p *tron_parser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tron_parserRULE_messageBody)

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

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserDOT, tron_parserREPEATED, tron_parserBOOL, tron_parserBYTES, tron_parserDOUBLE, tron_parserFIXED32, tron_parserFIXED64, tron_parserFLOAT, tron_parserINT32, tron_parserINT64, tron_parserSFIXED32, tron_parserSFIXED64, tron_parserSINT32, tron_parserSINT64, tron_parserSTRING, tron_parserUINT32, tron_parserUINT64, tron_parserIdent:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Field()
		}

	case tron_parserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.EnumDefinition()
		}

	case tron_parserMESSAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Message()
		}

	case tron_parserOPTION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(172)
			p.Option()
		}

	case tron_parserONEOF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Oneof()
		}

	case tron_parserMAP:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(174)
			p.MapField()
		}

	case tron_parserRESERVED:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(175)
			p.Reserved()
		}

	case tron_parserSEMI:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(176)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumDefinitionContext is an interface to support dynamic dispatch.
type IEnumDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefinitionContext differentiates from other interfaces.
	IsEnumDefinitionContext()
}

type EnumDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefinitionContext() *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_enumDefinition
	return p
}

func (*EnumDefinitionContext) IsEnumDefinitionContext() {}

func NewEnumDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_enumDefinition

	return p
}

func (s *EnumDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefinitionContext) ENUM() antlr.TerminalNode {
	return s.GetToken(tron_parserENUM, 0)
}

func (s *EnumDefinitionContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefinitionContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *EnumDefinitionContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
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

func (s *EnumDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterEnumDefinition(s)
	}
}

func (s *EnumDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitEnumDefinition(s)
	}
}

func (p *tron_parser) EnumDefinition() (localctx IEnumDefinitionContext) {
	localctx = NewEnumDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tron_parserRULE_enumDefinition)
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
		p.SetState(179)
		p.Match(tron_parserENUM)
	}
	{
		p.SetState(180)
		p.EnumName()
	}
	{
		p.SetState(181)
		p.Match(tron_parserLCUR)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tron_parserSEMI || _la == tron_parserOPTION || _la == tron_parserIdent {
		{
			p.SetState(182)
			p.EnumBody()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
		p.Match(tron_parserRCUR)
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
	p.RuleIndex = tron_parserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_enumBody

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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *tron_parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tron_parserRULE_enumBody)

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

	switch p.GetTokenStream().LA(1) {
	case tron_parserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Option()
		}

	case tron_parserIdent:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.EnumField()
		}

	case tron_parserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = tron_parserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *EnumFieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tron_parserInt_lit, 0)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *EnumFieldContext) DASH() antlr.TerminalNode {
	return s.GetToken(tron_parserDASH, 0)
}

func (s *EnumFieldContext) LSB() antlr.TerminalNode {
	return s.GetToken(tron_parserLSB, 0)
}

func (s *EnumFieldContext) AllEnumValueOption() []IEnumValueOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem())
	var tst = make([]IEnumValueOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueOptionContext)
		}
	}

	return tst
}

func (s *EnumFieldContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
}

func (s *EnumFieldContext) RSB() antlr.TerminalNode {
	return s.GetToken(tron_parserRSB, 0)
}

func (s *EnumFieldContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tron_parserCOMMA)
}

func (s *EnumFieldContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserCOMMA, i)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *tron_parser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tron_parserRULE_enumField)
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
		p.Match(tron_parserIdent)
	}
	{
		p.SetState(195)
		p.Match(tron_parserEQ)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserDASH {
		{
			p.SetState(196)
			p.Match(tron_parserDASH)
		}

	}
	{
		p.SetState(199)
		p.Match(tron_parserInt_lit)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserLSB {
		{
			p.SetState(200)
			p.Match(tron_parserLSB)
		}
		{
			p.SetState(201)
			p.EnumValueOption()
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tron_parserCOMMA {
			{
				p.SetState(202)
				p.Match(tron_parserCOMMA)
			}
			{
				p.SetState(203)
				p.EnumValueOption()
			}

			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(209)
			p.Match(tron_parserRSB)
		}

	}
	{
		p.SetState(213)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_enumValueOption
	return p
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (p *tron_parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tron_parserRULE_enumValueOption)

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
		p.SetState(215)
		p.OptionName()
	}
	{
		p.SetState(216)
		p.Match(tron_parserEQ)
	}
	{
		p.SetState(217)
		p.Constant()
	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(tron_parserSERVICE, 0)
}

func (s *ServiceContext) ServiceName() IServiceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *ServiceContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
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

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *tron_parser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tron_parserRULE_service)
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
		p.SetState(219)
		p.Match(tron_parserSERVICE)
	}
	{
		p.SetState(220)
		p.ServiceName()
	}
	{
		p.SetState(221)
		p.Match(tron_parserLCUR)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tron_parserSEMI)|(1<<tron_parserOPTION)|(1<<tron_parserRPC))) != 0 {
		{
			p.SetState(222)
			p.ServiceBody()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(228)
		p.Match(tron_parserRCUR)
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
	p.RuleIndex = tron_parserRULE_serviceBody
	return p
}

func (*ServiceBodyContext) IsServiceBodyContext() {}

func NewServiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBodyContext {
	var p = new(ServiceBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_serviceBody

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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterServiceBody(s)
	}
}

func (s *ServiceBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitServiceBody(s)
	}
}

func (p *tron_parser) ServiceBody() (localctx IServiceBodyContext) {
	localctx = NewServiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tron_parserRULE_serviceBody)

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

	p.SetState(233)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Option()
		}

	case tron_parserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Rpc()
		}

	case tron_parserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(tron_parserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
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

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(tron_parserRETURNS, 0)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *RpcContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *RpcContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitRpc(s)
	}
}

func (p *tron_parser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tron_parserRULE_rpc)
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
		p.SetState(235)
		p.Match(tron_parserRPC)
	}
	{
		p.SetState(236)
		p.RpcName()
	}
	{
		p.SetState(237)
		p.RpcParam()
	}
	{
		p.SetState(238)
		p.Match(tron_parserRETURNS)
	}
	{
		p.SetState(239)
		p.RpcParam()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserLCUR:
		{
			p.SetState(240)
			p.Match(tron_parserLCUR)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tron_parserSEMI || _la == tron_parserOPTION {
			p.SetState(243)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case tron_parserOPTION:
				{
					p.SetState(241)
					p.Option()
				}

			case tron_parserSEMI:
				{
					p.SetState(242)
					p.EmptyStatement()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(248)
			p.Match(tron_parserRCUR)
		}

	case tron_parserSEMI:
		{
			p.SetState(249)
			p.Match(tron_parserSEMI)
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
	p.RuleIndex = tron_parserRULE_rpcParam
	return p
}

func (*RpcParamContext) IsRpcParamContext() {}

func NewRpcParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcParamContext {
	var p = new(RpcParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_rpcParam

	return p
}

func (s *RpcParamContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcParamContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tron_parserLPAREN, 0)
}

func (s *RpcParamContext) MessageType() IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RpcParamContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tron_parserRPAREN, 0)
}

func (s *RpcParamContext) STREAM() antlr.TerminalNode {
	return s.GetToken(tron_parserSTREAM, 0)
}

func (s *RpcParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterRpcParam(s)
	}
}

func (s *RpcParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitRpcParam(s)
	}
}

func (p *tron_parser) RpcParam() (localctx IRpcParamContext) {
	localctx = NewRpcParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tron_parserRULE_rpcParam)
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
		p.SetState(252)
		p.Match(tron_parserLPAREN)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserSTREAM {
		{
			p.SetState(253)
			p.Match(tron_parserSTREAM)
		}

	}
	{
		p.SetState(256)
		p.MessageType()
	}
	{
		p.SetState(257)
		p.Match(tron_parserRPAREN)
	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(tron_parserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitReserved(s)
	}
}

func (p *tron_parser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tron_parserRULE_reserved)

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
		p.SetState(259)
		p.Match(tron_parserRESERVED)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserInt_lit:
		{
			p.SetState(260)
			p.Ranges()
		}

	case tron_parserStrLit:
		{
			p.SetState(261)
			p.FieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(264)
		p.Match(tron_parserSEMI)
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
	p.RuleIndex = tron_parserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_ranges

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
	return s.GetTokens(tron_parserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *tron_parser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tron_parserRULE_ranges)
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
		p.SetState(266)
		p.Rangee()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tron_parserCOMMA {
		{
			p.SetState(267)
			p.Match(tron_parserCOMMA)
		}
		{
			p.SetState(268)
			p.Rangee()
		}

		p.SetState(273)
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

	// IsRangeeContext differentiates from other interfaces.
	IsRangeeContext()
}

type RangeeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeeContext() *RangeeContext {
	var p = new(RangeeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_rangee
	return p
}

func (*RangeeContext) IsRangeeContext() {}

func NewRangeeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeeContext {
	var p = new(RangeeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_rangee

	return p
}

func (s *RangeeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeeContext) AllInt_lit() []antlr.TerminalNode {
	return s.GetTokens(tron_parserInt_lit)
}

func (s *RangeeContext) Int_lit(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserInt_lit, i)
}

func (s *RangeeContext) TO() antlr.TerminalNode {
	return s.GetToken(tron_parserTO, 0)
}

func (s *RangeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterRangee(s)
	}
}

func (s *RangeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitRangee(s)
	}
}

func (p *tron_parser) Rangee() (localctx IRangeeContext) {
	localctx = NewRangeeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tron_parserRULE_rangee)

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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Match(tron_parserInt_lit)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Match(tron_parserInt_lit)
		}
		{
			p.SetState(276)
			p.Match(tron_parserTO)
		}
		{
			p.SetState(277)
			p.Match(tron_parserInt_lit)
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
	p.RuleIndex = tron_parserRULE_fieldNames
	return p
}

func (*FieldNamesContext) IsFieldNamesContext() {}

func NewFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNamesContext {
	var p = new(FieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_fieldNames

	return p
}

func (s *FieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNamesContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(tron_parserStrLit)
}

func (s *FieldNamesContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserStrLit, i)
}

func (s *FieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tron_parserCOMMA)
}

func (s *FieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserCOMMA, i)
}

func (s *FieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterFieldNames(s)
	}
}

func (s *FieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitFieldNames(s)
	}
}

func (p *tron_parser) FieldNames() (localctx IFieldNamesContext) {
	localctx = NewFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, tron_parserRULE_fieldNames)
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
		p.SetState(280)
		p.Match(tron_parserStrLit)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tron_parserCOMMA {
		{
			p.SetState(281)
			p.Match(tron_parserCOMMA)
		}
		{
			p.SetState(282)
			p.Match(tron_parserStrLit)
		}

		p.SetState(287)
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
	p.RuleIndex = tron_parserRULE_typer
	return p
}

func (*TyperContext) IsTyperContext() {}

func NewTyperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TyperContext {
	var p = new(TyperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_typer

	return p
}

func (s *TyperContext) GetParser() antlr.Parser { return s.parser }

func (s *TyperContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *TyperContext) MessageOrEnumType() IMessageOrEnumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageOrEnumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumTypeContext)
}

func (s *TyperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TyperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TyperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterTyper(s)
	}
}

func (s *TyperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitTyper(s)
	}
}

func (p *tron_parser) Typer() (localctx ITyperContext) {
	localctx = NewTyperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, tron_parserRULE_typer)

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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserBOOL, tron_parserBYTES, tron_parserDOUBLE, tron_parserFIXED32, tron_parserFIXED64, tron_parserFLOAT, tron_parserINT32, tron_parserINT64, tron_parserSFIXED32, tron_parserSFIXED64, tron_parserSINT32, tron_parserSINT64, tron_parserSTRING, tron_parserUINT32, tron_parserUINT64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Types()
		}

	case tron_parserDOT, tron_parserIdent:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.MessageOrEnumType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(tron_parserDOUBLE, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(tron_parserFLOAT, 0)
}

func (s *TypesContext) INT32() antlr.TerminalNode {
	return s.GetToken(tron_parserINT32, 0)
}

func (s *TypesContext) INT64() antlr.TerminalNode {
	return s.GetToken(tron_parserINT64, 0)
}

func (s *TypesContext) UINT32() antlr.TerminalNode {
	return s.GetToken(tron_parserUINT32, 0)
}

func (s *TypesContext) UINT64() antlr.TerminalNode {
	return s.GetToken(tron_parserUINT64, 0)
}

func (s *TypesContext) SINT32() antlr.TerminalNode {
	return s.GetToken(tron_parserSINT32, 0)
}

func (s *TypesContext) SINT64() antlr.TerminalNode {
	return s.GetToken(tron_parserSINT64, 0)
}

func (s *TypesContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(tron_parserFIXED32, 0)
}

func (s *TypesContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(tron_parserFIXED64, 0)
}

func (s *TypesContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(tron_parserSFIXED32, 0)
}

func (s *TypesContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(tron_parserSFIXED64, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(tron_parserBOOL, 0)
}

func (s *TypesContext) STRING() antlr.TerminalNode {
	return s.GetToken(tron_parserSTRING, 0)
}

func (s *TypesContext) BYTES() antlr.TerminalNode {
	return s.GetToken(tron_parserBYTES, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *tron_parser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, tron_parserRULE_types)
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
		p.SetState(292)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(tron_parserBOOL-40))|(1<<(tron_parserBYTES-40))|(1<<(tron_parserDOUBLE-40))|(1<<(tron_parserFIXED32-40))|(1<<(tron_parserFIXED64-40))|(1<<(tron_parserFLOAT-40))|(1<<(tron_parserINT32-40))|(1<<(tron_parserINT64-40))|(1<<(tron_parserSFIXED32-40))|(1<<(tron_parserSFIXED64-40))|(1<<(tron_parserSINT32-40))|(1<<(tron_parserSINT64-40))|(1<<(tron_parserSTRING-40))|(1<<(tron_parserUINT32-40))|(1<<(tron_parserUINT64-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = tron_parserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_field

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

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *FieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tron_parserInt_lit, 0)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(tron_parserREPEATED, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *tron_parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, tron_parserRULE_field)
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
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserREPEATED {
		{
			p.SetState(294)
			p.Match(tron_parserREPEATED)
		}

	}
	{
		p.SetState(297)
		p.Typer()
	}
	{
		p.SetState(298)
		p.FieldName()
	}
	{
		p.SetState(299)
		p.Match(tron_parserEQ)
	}
	{
		p.SetState(300)
		p.Match(tron_parserInt_lit)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserLSB {
		{
			p.SetState(301)
			p.FieldOptions()
		}

	}
	{
		p.SetState(304)
		p.Match(tron_parserSEMI)
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
	p.RuleIndex = tron_parserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) LSB() antlr.TerminalNode {
	return s.GetToken(tron_parserLSB, 0)
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
	return s.GetToken(tron_parserRSB, 0)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tron_parserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (p *tron_parser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, tron_parserRULE_fieldOptions)
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
		p.SetState(306)
		p.Match(tron_parserLSB)
	}
	{
		p.SetState(307)
		p.FieldOption()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tron_parserCOMMA {
		{
			p.SetState(308)
			p.Match(tron_parserCOMMA)
		}
		{
			p.SetState(309)
			p.FieldOption()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(315)
		p.Match(tron_parserRSB)
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
	p.RuleIndex = tron_parserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_fieldOption

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
	return s.GetToken(tron_parserEQ, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (p *tron_parser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, tron_parserRULE_fieldOption)

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
		p.SetState(317)
		p.OptionName()
	}
	{
		p.SetState(318)
		p.Match(tron_parserEQ)
	}
	{
		p.SetState(319)
		p.Constant()
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(tron_parserONEOF, 0)
}

func (s *OneofContext) OneofName() IOneofNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *OneofContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitOneof(s)
	}
}

func (p *tron_parser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, tron_parserRULE_oneof)
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
		p.SetState(321)
		p.Match(tron_parserONEOF)
	}
	{
		p.SetState(322)
		p.OneofName()
	}
	{
		p.SetState(323)
		p.Match(tron_parserLCUR)
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tron_parserSEMI || _la == tron_parserDOT || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(tron_parserBOOL-40))|(1<<(tron_parserBYTES-40))|(1<<(tron_parserDOUBLE-40))|(1<<(tron_parserFIXED32-40))|(1<<(tron_parserFIXED64-40))|(1<<(tron_parserFLOAT-40))|(1<<(tron_parserINT32-40))|(1<<(tron_parserINT64-40))|(1<<(tron_parserSFIXED32-40))|(1<<(tron_parserSFIXED64-40))|(1<<(tron_parserSINT32-40))|(1<<(tron_parserSINT64-40))|(1<<(tron_parserSTRING-40))|(1<<(tron_parserUINT32-40))|(1<<(tron_parserUINT64-40))|(1<<(tron_parserIdent-40)))) != 0) {
		p.SetState(326)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tron_parserDOT, tron_parserBOOL, tron_parserBYTES, tron_parserDOUBLE, tron_parserFIXED32, tron_parserFIXED64, tron_parserFLOAT, tron_parserINT32, tron_parserINT64, tron_parserSFIXED32, tron_parserSFIXED64, tron_parserSINT32, tron_parserSINT64, tron_parserSTRING, tron_parserUINT32, tron_parserUINT64, tron_parserIdent:
			{
				p.SetState(324)
				p.OneofField()
			}

		case tron_parserSEMI:
			{
				p.SetState(325)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(331)
		p.Match(tron_parserRCUR)
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
	p.RuleIndex = tron_parserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_oneofField

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

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *OneofFieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tron_parserInt_lit, 0)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (p *tron_parser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, tron_parserRULE_oneofField)
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
		p.SetState(333)
		p.Typer()
	}
	{
		p.SetState(334)
		p.FieldName()
	}
	{
		p.SetState(335)
		p.Match(tron_parserEQ)
	}
	{
		p.SetState(336)
		p.Match(tron_parserInt_lit)
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserLSB {
		{
			p.SetState(337)
			p.FieldOptions()
		}

	}
	{
		p.SetState(340)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(tron_parserMAP, 0)
}

func (s *MapFieldContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCHEVR, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tron_parserCOMMA, 0)
}

func (s *MapFieldContext) Typer() ITyperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyperContext)
}

func (s *MapFieldContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCHEVR, 0)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tron_parserEQ, 0)
}

func (s *MapFieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tron_parserInt_lit, 0)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
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
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *tron_parser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, tron_parserRULE_mapField)
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
		p.SetState(342)
		p.Match(tron_parserMAP)
	}
	{
		p.SetState(343)
		p.Match(tron_parserLCHEVR)
	}
	{
		p.SetState(344)
		p.KeyType()
	}
	{
		p.SetState(345)
		p.Match(tron_parserCOMMA)
	}
	{
		p.SetState(346)
		p.Typer()
	}
	{
		p.SetState(347)
		p.Match(tron_parserRCHEVR)
	}
	{
		p.SetState(348)
		p.MapName()
	}
	{
		p.SetState(349)
		p.Match(tron_parserEQ)
	}
	{
		p.SetState(350)
		p.Match(tron_parserInt_lit)
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserLSB {
		{
			p.SetState(351)
			p.FieldOptions()
		}

	}
	{
		p.SetState(354)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(tron_parserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(tron_parserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(tron_parserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(tron_parserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(tron_parserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(tron_parserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(tron_parserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(tron_parserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(tron_parserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(tron_parserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(tron_parserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(tron_parserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *tron_parser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, tron_parserRULE_keyType)
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
		p.SetState(356)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(tron_parserBOOL-40))|(1<<(tron_parserFIXED32-40))|(1<<(tron_parserFIXED64-40))|(1<<(tron_parserINT32-40))|(1<<(tron_parserINT64-40))|(1<<(tron_parserSFIXED32-40))|(1<<(tron_parserSFIXED64-40))|(1<<(tron_parserSINT32-40))|(1<<(tron_parserSINT64-40))|(1<<(tron_parserSTRING-40))|(1<<(tron_parserUINT32-40))|(1<<(tron_parserUINT64-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tron_parserIdent)
}

func (s *FullIdentContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, i)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tron_parserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *tron_parser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, tron_parserRULE_fullIdent)
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
		p.SetState(358)
		p.Match(tron_parserIdent)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tron_parserDOT {
		{
			p.SetState(359)
			p.Match(tron_parserDOT)
		}
		{
			p.SetState(360)
			p.Match(tron_parserIdent)
		}

		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_messageName
	return p
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (p *tron_parser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, tron_parserRULE_messageName)

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
		p.SetState(366)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (p *tron_parser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, tron_parserRULE_enumName)

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
		p.SetState(368)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IMessageOrEnumNameContext is an interface to support dynamic dispatch.
type IMessageOrEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageOrEnumNameContext differentiates from other interfaces.
	IsMessageOrEnumNameContext()
}

type MessageOrEnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumNameContext() *MessageOrEnumNameContext {
	var p = new(MessageOrEnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_messageOrEnumName
	return p
}

func (*MessageOrEnumNameContext) IsMessageOrEnumNameContext() {}

func NewMessageOrEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumNameContext {
	var p = new(MessageOrEnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_messageOrEnumName

	return p
}

func (s *MessageOrEnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *MessageOrEnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMessageOrEnumName(s)
	}
}

func (s *MessageOrEnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMessageOrEnumName(s)
	}
}

func (p *tron_parser) MessageOrEnumName() (localctx IMessageOrEnumNameContext) {
	localctx = NewMessageOrEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, tron_parserRULE_messageOrEnumName)

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
		p.SetState(370)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *tron_parser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, tron_parserRULE_fieldName)

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
		p.SetState(372)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_oneofName
	return p
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterOneofName(s)
	}
}

func (s *OneofNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitOneofName(s)
	}
}

func (p *tron_parser) OneofName() (localctx IOneofNameContext) {
	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, tron_parserRULE_oneofName)

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
		p.SetState(374)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_mapName
	return p
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMapName(s)
	}
}

func (p *tron_parser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, tron_parserRULE_mapName)

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
		p.SetState(376)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (p *tron_parser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, tron_parserRULE_serviceName)

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
		p.SetState(378)
		p.Match(tron_parserIdent)
	}

	return localctx
}

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_rpcName
	return p
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (p *tron_parser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, tron_parserRULE_rpcName)

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
		p.SetState(380)
		p.Match(tron_parserIdent)
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
	p.RuleIndex = tron_parserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tron_parserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tron_parserIdent)
}

func (s *MessageTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, i)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *tron_parser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, tron_parserRULE_messageType)
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
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserDOT {
		{
			p.SetState(382)
			p.Match(tron_parserDOT)
		}

	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(385)
				p.Match(tron_parserIdent)
			}
			{
				p.SetState(386)
				p.Match(tron_parserDOT)
			}

		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}
	{
		p.SetState(392)
		p.MessageName()
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
	p.RuleIndex = tron_parserRULE_messageOrEnumType
	return p
}

func (*MessageOrEnumTypeContext) IsMessageOrEnumTypeContext() {}

func NewMessageOrEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_messageOrEnumType

	return p
}

func (s *MessageOrEnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumTypeContext) MessageOrEnumName() IMessageOrEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageOrEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumNameContext)
}

func (s *MessageOrEnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tron_parserDOT)
}

func (s *MessageOrEnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserDOT, i)
}

func (s *MessageOrEnumTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tron_parserIdent)
}

func (s *MessageOrEnumTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, i)
}

func (s *MessageOrEnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterMessageOrEnumType(s)
	}
}

func (s *MessageOrEnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitMessageOrEnumType(s)
	}
}

func (p *tron_parser) MessageOrEnumType() (localctx IMessageOrEnumTypeContext) {
	localctx = NewMessageOrEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, tron_parserRULE_messageOrEnumType)
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
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tron_parserDOT {
		{
			p.SetState(394)
			p.Match(tron_parserDOT)
		}

	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(397)
				p.Match(tron_parserIdent)
			}
			{
				p.SetState(398)
				p.Match(tron_parserDOT)
			}

		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}
	{
		p.SetState(404)
		p.MessageOrEnumName()
	}

	return localctx
}

// IBool_litContext is an interface to support dynamic dispatch.
type IBool_litContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_litContext differentiates from other interfaces.
	IsBool_litContext()
}

type Bool_litContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_litContext() *Bool_litContext {
	var p = new(Bool_litContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_bool_lit
	return p
}

func (*Bool_litContext) IsBool_litContext() {}

func NewBool_litContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_litContext {
	var p = new(Bool_litContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_bool_lit

	return p
}

func (s *Bool_litContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_litContext) TRUE() antlr.TerminalNode {
	return s.GetToken(tron_parserTRUE, 0)
}

func (s *Bool_litContext) FALSE() antlr.TerminalNode {
	return s.GetToken(tron_parserFALSE, 0)
}

func (s *Bool_litContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_litContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_litContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterBool_lit(s)
	}
}

func (s *Bool_litContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitBool_lit(s)
	}
}

func (p *tron_parser) Bool_lit() (localctx IBool_litContext) {
	localctx = NewBool_litContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, tron_parserRULE_bool_lit)
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
		p.SetState(406)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tron_parserTRUE || _la == tron_parserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = tron_parserRULE_emptyStatement
	return p
}

func (*EmptyStatementContext) IsEmptyStatementContext() {}

func NewEmptyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatementContext {
	var p = new(EmptyStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_emptyStatement

	return p
}

func (s *EmptyStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tron_parserSEMI, 0)
}

func (s *EmptyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterEmptyStatement(s)
	}
}

func (s *EmptyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitEmptyStatement(s)
	}
}

func (p *tron_parser) EmptyStatement() (localctx IEmptyStatementContext) {
	localctx = NewEmptyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, tron_parserRULE_emptyStatement)

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
		p.SetState(408)
		p.Match(tron_parserSEMI)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tron_parserInt_lit, 0)
}

func (s *ConstantContext) DASH() antlr.TerminalNode {
	return s.GetToken(tron_parserDASH, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(tron_parserPLUS, 0)
}

func (s *ConstantContext) Float_lit() antlr.TerminalNode {
	return s.GetToken(tron_parserFloat_lit, 0)
}

func (s *ConstantContext) StrLit() antlr.TerminalNode {
	return s.GetToken(tron_parserStrLit, 0)
}

func (s *ConstantContext) Bool_lit() IBool_litContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_litContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_litContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *tron_parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, tron_parserRULE_constant)
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

	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tron_parserDASH || _la == tron_parserPLUS {
			{
				p.SetState(411)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tron_parserDASH || _la == tron_parserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(414)
			p.Match(tron_parserInt_lit)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tron_parserDASH || _la == tron_parserPLUS {
			{
				p.SetState(415)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tron_parserDASH || _la == tron_parserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(418)
			p.Match(tron_parserFloat_lit)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(421)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tron_parserStrLit:
			{
				p.SetState(419)
				p.Match(tron_parserStrLit)
			}

		case tron_parserTRUE, tron_parserFALSE:
			{
				p.SetState(420)
				p.Bool_lit()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IPronContext is an interface to support dynamic dispatch.
type IPronContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsPronContext differentiates from other interfaces.
	IsPronContext()
}

type PronContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyPronContext() *PronContext {
	var p = new(PronContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_pron
	return p
}

func (*PronContext) IsPronContext() {}

func NewPronContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PronContext {
	var p = new(PronContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_pron

	return p
}

func (s *PronContext) GetParser() antlr.Parser { return s.parser }

func (s *PronContext) GetId() antlr.Token { return s.id }

func (s *PronContext) SetId(v antlr.Token) { s.id = v }

func (s *PronContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserLCUR, 0)
}

func (s *PronContext) PronElem() IPronElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronElemContext)
}

func (s *PronContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tron_parserRCUR, 0)
}

func (s *PronContext) Ident() antlr.TerminalNode {
	return s.GetToken(tron_parserIdent, 0)
}

func (s *PronContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PronContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterPron(s)
	}
}

func (s *PronContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitPron(s)
	}
}

func (p *tron_parser) Pron() (localctx IPronContext) {
	localctx = NewPronContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, tron_parserRULE_pron)

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
		p.SetState(425)
		p.Match(tron_parserLCUR)
	}
	{
		p.SetState(426)

		var _m = p.Match(tron_parserIdent)

		localctx.(*PronContext).id = _m
	}
	{
		p.SetState(427)
		p.Match(tron_parserCOLON)
	}
	{
		p.SetState(428)
		p.PronElem()
	}
	{
		p.SetState(429)
		p.Match(tron_parserRCUR)
	}

	return localctx
}

// IPronElemContext is an interface to support dynamic dispatch.
type IPronElemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPronElemContext differentiates from other interfaces.
	IsPronElemContext()
}

type PronElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPronElemContext() *PronElemContext {
	var p = new(PronElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tron_parserRULE_pronElem
	return p
}

func (*PronElemContext) IsPronElemContext() {}

func NewPronElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PronElemContext {
	var p = new(PronElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tron_parserRULE_pronElem

	return p
}

func (s *PronElemContext) GetParser() antlr.Parser { return s.parser }

func (s *PronElemContext) CopyFrom(ctx *PronElemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PronElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PronSTRContext struct {
	*PronElemContext
}

func NewPronSTRContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PronSTRContext {
	var p = new(PronSTRContext)

	p.PronElemContext = NewEmptyPronElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PronElemContext))

	return p
}

func (s *PronSTRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronSTRContext) StrLit() antlr.TerminalNode {
	return s.GetToken(tron_parserStrLit, 0)
}

func (s *PronSTRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterPronSTR(s)
	}
}

func (s *PronSTRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitPronSTR(s)
	}
}

type PronOBJContext struct {
	*PronElemContext
}

func NewPronOBJContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PronOBJContext {
	var p = new(PronOBJContext)

	p.PronElemContext = NewEmptyPronElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PronElemContext))

	return p
}

func (s *PronOBJContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronOBJContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *PronOBJContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.EnterPronOBJ(s)
	}
}

func (s *PronOBJContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tron_parserListener); ok {
		listenerT.ExitPronOBJ(s)
	}
}

func (p *tron_parser) PronElem() (localctx IPronElemContext) {
	localctx = NewPronElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, tron_parserRULE_pronElem)

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

	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tron_parserStrLit:
		localctx = NewPronSTRContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(tron_parserStrLit)
		}

	case tron_parserLCUR:
		localctx = NewPronOBJContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
