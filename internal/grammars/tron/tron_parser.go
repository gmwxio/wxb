// Code generated from TronParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TronParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 386,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 5, 2, 40, 10, 2, 3, 2, 7, 2, 43, 10, 2, 12, 2,
	14, 2, 46, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	5, 4, 57, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 65, 10, 4, 12,
	4, 14, 4, 68, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 81, 10, 4, 12, 4, 14, 4, 84, 11, 4, 3, 4, 3, 4, 7,
	4, 88, 10, 4, 12, 4, 14, 4, 91, 11, 4, 3, 4, 3, 4, 5, 4, 95, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 106, 10, 5, 13,
	5, 14, 5, 107, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 124, 10, 5, 3, 5, 3, 5, 5, 5, 128, 10,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 139, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 7, 7, 154, 10, 7, 12, 7, 14, 7, 157, 11, 7, 3, 7, 3, 7, 3, 7,
	7, 7, 162, 10, 7, 12, 7, 14, 7, 165, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 7, 7, 173, 10, 7, 12, 7, 14, 7, 176, 11, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 7, 7, 183, 10, 7, 12, 7, 14, 7, 186, 11, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 195, 10, 7, 13, 7, 14, 7, 196, 3, 7, 5,
	7, 200, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 209, 10,
	7, 12, 7, 14, 7, 212, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 7, 7, 224, 10, 7, 12, 7, 14, 7, 227, 11, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 232, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 237, 10, 8, 5, 8, 239,
	10, 8, 3, 9, 3, 9, 3, 9, 7, 9, 244, 10, 9, 12, 9, 14, 9, 247, 11, 9, 3,
	10, 3, 10, 5, 10, 251, 10, 10, 3, 10, 5, 10, 254, 10, 10, 3, 10, 3, 10,
	7, 10, 258, 10, 10, 12, 10, 14, 10, 261, 11, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 274, 10, 11,
	12, 11, 14, 11, 277, 11, 11, 3, 11, 3, 11, 5, 11, 281, 10, 11, 3, 12, 3,
	12, 3, 12, 7, 12, 286, 10, 12, 12, 12, 14, 12, 289, 11, 12, 3, 12, 3, 12,
	3, 12, 7, 12, 294, 10, 12, 12, 12, 14, 12, 297, 11, 12, 5, 12, 299, 10,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 306, 10, 13, 5, 13, 308,
	10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 314, 10, 14, 12, 14, 14, 14,
	317, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 7, 16, 330, 10, 16, 12, 16, 14, 16, 333, 11, 16, 3, 16,
	5, 16, 336, 10, 16, 3, 16, 3, 16, 7, 16, 340, 10, 16, 12, 16, 14, 16, 343,
	11, 16, 3, 17, 3, 17, 3, 17, 6, 17, 348, 10, 17, 13, 17, 14, 17, 349, 3,
	17, 3, 17, 5, 17, 354, 10, 17, 3, 17, 3, 17, 5, 17, 358, 10, 17, 3, 17,
	3, 17, 3, 17, 5, 17, 363, 10, 17, 3, 18, 3, 18, 3, 18, 5, 18, 368, 10,
	18, 3, 18, 7, 18, 371, 10, 18, 12, 18, 14, 18, 374, 11, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 384, 10, 19, 3, 19, 2,
	2, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 2, 4, 3, 2, 25, 26, 4, 2, 15, 15, 22, 22, 2, 431, 2, 39, 3, 2, 2, 2,
	4, 49, 3, 2, 2, 2, 6, 94, 3, 2, 2, 2, 8, 138, 3, 2, 2, 2, 10, 140, 3, 2,
	2, 2, 12, 231, 3, 2, 2, 2, 14, 238, 3, 2, 2, 2, 16, 240, 3, 2, 2, 2, 18,
	248, 3, 2, 2, 2, 20, 280, 3, 2, 2, 2, 22, 298, 3, 2, 2, 2, 24, 307, 3,
	2, 2, 2, 26, 309, 3, 2, 2, 2, 28, 320, 3, 2, 2, 2, 30, 335, 3, 2, 2, 2,
	32, 362, 3, 2, 2, 2, 34, 364, 3, 2, 2, 2, 36, 383, 3, 2, 2, 2, 38, 40,
	5, 4, 3, 2, 39, 38, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 44, 3, 2, 2, 2,
	41, 43, 5, 6, 4, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3,
	2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47,
	48, 7, 2, 2, 3, 48, 3, 3, 2, 2, 2, 49, 50, 7, 4, 2, 2, 50, 51, 7, 5, 2,
	2, 51, 52, 7, 3, 2, 2, 52, 53, 7, 6, 2, 2, 53, 5, 3, 2, 2, 2, 54, 56, 7,
	28, 2, 2, 55, 57, 7, 28, 2, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2,
	57, 58, 3, 2, 2, 2, 58, 59, 7, 27, 2, 2, 59, 95, 7, 15, 2, 2, 60, 61, 7,
	28, 2, 2, 61, 66, 7, 28, 2, 2, 62, 63, 7, 19, 2, 2, 63, 65, 7, 28, 2, 2,
	64, 62, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3,
	2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 95, 7, 15, 2, 2, 70,
	71, 7, 28, 2, 2, 71, 72, 5, 30, 16, 2, 72, 73, 7, 12, 2, 2, 73, 74, 5,
	32, 17, 2, 74, 75, 7, 15, 2, 2, 75, 95, 3, 2, 2, 2, 76, 77, 7, 28, 2, 2,
	77, 82, 7, 28, 2, 2, 78, 79, 7, 19, 2, 2, 79, 81, 7, 28, 2, 2, 80, 78,
	3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 89, 7, 10, 2, 2, 86, 88, 5,
	8, 5, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89,
	90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 95, 7, 11,
	2, 2, 93, 95, 7, 15, 2, 2, 94, 54, 3, 2, 2, 2, 94, 60, 3, 2, 2, 2, 94,
	70, 3, 2, 2, 2, 94, 76, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 7, 3, 2, 2,
	2, 96, 139, 5, 10, 6, 2, 97, 98, 7, 28, 2, 2, 98, 99, 7, 28, 2, 2, 99,
	100, 7, 10, 2, 2, 100, 139, 7, 11, 2, 2, 101, 102, 7, 28, 2, 2, 102, 103,
	7, 28, 2, 2, 103, 105, 7, 10, 2, 2, 104, 106, 5, 8, 5, 2, 105, 104, 3,
	2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2,
	2, 108, 109, 3, 2, 2, 2, 109, 110, 7, 11, 2, 2, 110, 139, 3, 2, 2, 2, 111,
	112, 7, 28, 2, 2, 112, 113, 7, 28, 2, 2, 113, 114, 7, 10, 2, 2, 114, 115,
	7, 15, 2, 2, 115, 139, 7, 11, 2, 2, 116, 117, 7, 28, 2, 2, 117, 118, 5,
	22, 12, 2, 118, 119, 7, 15, 2, 2, 119, 139, 3, 2, 2, 2, 120, 121, 7, 28,
	2, 2, 121, 123, 7, 12, 2, 2, 122, 124, 7, 25, 2, 2, 123, 122, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 7, 29, 2, 2, 126,
	128, 5, 26, 14, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129,
	3, 2, 2, 2, 129, 139, 7, 15, 2, 2, 130, 131, 7, 28, 2, 2, 131, 132, 7,
	28, 2, 2, 132, 133, 5, 18, 10, 2, 133, 134, 7, 28, 2, 2, 134, 135, 5, 18,
	10, 2, 135, 136, 5, 20, 11, 2, 136, 139, 3, 2, 2, 2, 137, 139, 7, 15, 2,
	2, 138, 96, 3, 2, 2, 2, 138, 97, 3, 2, 2, 2, 138, 101, 3, 2, 2, 2, 138,
	111, 3, 2, 2, 2, 138, 116, 3, 2, 2, 2, 138, 120, 3, 2, 2, 2, 138, 130,
	3, 2, 2, 2, 138, 137, 3, 2, 2, 2, 139, 9, 3, 2, 2, 2, 140, 141, 5, 12,
	7, 2, 141, 142, 7, 12, 2, 2, 142, 143, 5, 14, 8, 2, 143, 144, 7, 15, 2,
	2, 144, 11, 3, 2, 2, 2, 145, 232, 7, 28, 2, 2, 146, 147, 7, 28, 2, 2, 147,
	232, 7, 28, 2, 2, 148, 149, 7, 28, 2, 2, 149, 150, 7, 17, 2, 2, 150, 155,
	7, 28, 2, 2, 151, 152, 7, 19, 2, 2, 152, 154, 7, 28, 2, 2, 153, 151, 3,
	2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2,
	2, 156, 158, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 163, 7, 18, 2, 2, 159,
	160, 7, 19, 2, 2, 160, 162, 7, 28, 2, 2, 161, 159, 3, 2, 2, 2, 162, 165,
	3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 232, 3, 2,
	2, 2, 165, 163, 3, 2, 2, 2, 166, 167, 7, 28, 2, 2, 167, 168, 7, 19, 2,
	2, 168, 169, 7, 28, 2, 2, 169, 174, 7, 19, 2, 2, 170, 171, 7, 28, 2, 2,
	171, 173, 7, 19, 2, 2, 172, 170, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174,
	172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2, 176, 174,
	3, 2, 2, 2, 177, 178, 7, 28, 2, 2, 178, 232, 7, 28, 2, 2, 179, 184, 7,
	19, 2, 2, 180, 181, 7, 28, 2, 2, 181, 183, 7, 19, 2, 2, 182, 180, 3, 2,
	2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2,
	185, 187, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187, 188, 7, 28, 2, 2, 188,
	232, 7, 28, 2, 2, 189, 190, 7, 28, 2, 2, 190, 199, 7, 28, 2, 2, 191, 194,
	7, 19, 2, 2, 192, 193, 7, 28, 2, 2, 193, 195, 7, 19, 2, 2, 194, 192, 3,
	2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2,
	2, 197, 198, 3, 2, 2, 2, 198, 200, 7, 28, 2, 2, 199, 191, 3, 2, 2, 2, 199,
	200, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 232, 7, 28, 2, 2, 202, 203,
	7, 28, 2, 2, 203, 204, 7, 23, 2, 2, 204, 205, 7, 28, 2, 2, 205, 210, 7,
	22, 2, 2, 206, 207, 7, 28, 2, 2, 207, 209, 7, 19, 2, 2, 208, 206, 3, 2,
	2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2,
	211, 213, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 7, 28, 2, 2, 214,
	215, 7, 24, 2, 2, 215, 232, 7, 28, 2, 2, 216, 217, 7, 28, 2, 2, 217, 218,
	7, 23, 2, 2, 218, 219, 7, 28, 2, 2, 219, 220, 7, 22, 2, 2, 220, 225, 7,
	19, 2, 2, 221, 222, 7, 28, 2, 2, 222, 224, 7, 19, 2, 2, 223, 221, 3, 2,
	2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2,
	226, 228, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 229, 7, 28, 2, 2, 229,
	230, 7, 24, 2, 2, 230, 232, 7, 28, 2, 2, 231, 145, 3, 2, 2, 2, 231, 146,
	3, 2, 2, 2, 231, 148, 3, 2, 2, 2, 231, 166, 3, 2, 2, 2, 231, 179, 3, 2,
	2, 2, 231, 189, 3, 2, 2, 2, 231, 202, 3, 2, 2, 2, 231, 216, 3, 2, 2, 2,
	232, 13, 3, 2, 2, 2, 233, 239, 5, 32, 17, 2, 234, 236, 7, 29, 2, 2, 235,
	237, 5, 26, 14, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239,
	3, 2, 2, 2, 238, 233, 3, 2, 2, 2, 238, 234, 3, 2, 2, 2, 239, 15, 3, 2,
	2, 2, 240, 245, 7, 28, 2, 2, 241, 242, 7, 19, 2, 2, 242, 244, 7, 28, 2,
	2, 243, 241, 3, 2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245,
	246, 3, 2, 2, 2, 246, 17, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 250, 7,
	17, 2, 2, 249, 251, 7, 28, 2, 2, 250, 249, 3, 2, 2, 2, 250, 251, 3, 2,
	2, 2, 251, 253, 3, 2, 2, 2, 252, 254, 7, 19, 2, 2, 253, 252, 3, 2, 2, 2,
	253, 254, 3, 2, 2, 2, 254, 259, 3, 2, 2, 2, 255, 256, 7, 28, 2, 2, 256,
	258, 7, 19, 2, 2, 257, 255, 3, 2, 2, 2, 258, 261, 3, 2, 2, 2, 259, 257,
	3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 262, 3, 2, 2, 2, 261, 259, 3, 2,
	2, 2, 262, 263, 7, 28, 2, 2, 263, 264, 7, 18, 2, 2, 264, 19, 3, 2, 2, 2,
	265, 275, 7, 10, 2, 2, 266, 267, 7, 28, 2, 2, 267, 268, 5, 30, 16, 2, 268,
	269, 7, 12, 2, 2, 269, 270, 5, 32, 17, 2, 270, 271, 7, 15, 2, 2, 271, 274,
	3, 2, 2, 2, 272, 274, 7, 15, 2, 2, 273, 266, 3, 2, 2, 2, 273, 272, 3, 2,
	2, 2, 274, 277, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2,
	276, 278, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 278, 281, 7, 11, 2, 2, 279,
	281, 7, 15, 2, 2, 280, 265, 3, 2, 2, 2, 280, 279, 3, 2, 2, 2, 281, 21,
	3, 2, 2, 2, 282, 287, 5, 24, 13, 2, 283, 284, 7, 22, 2, 2, 284, 286, 5,
	24, 13, 2, 285, 283, 3, 2, 2, 2, 286, 289, 3, 2, 2, 2, 287, 285, 3, 2,
	2, 2, 287, 288, 3, 2, 2, 2, 288, 299, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2,
	290, 295, 7, 27, 2, 2, 291, 292, 7, 22, 2, 2, 292, 294, 7, 27, 2, 2, 293,
	291, 3, 2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 295, 296,
	3, 2, 2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 298, 282, 3, 2,
	2, 2, 298, 290, 3, 2, 2, 2, 299, 23, 3, 2, 2, 2, 300, 308, 7, 29, 2, 2,
	301, 302, 7, 29, 2, 2, 302, 305, 7, 28, 2, 2, 303, 306, 7, 29, 2, 2, 304,
	306, 7, 28, 2, 2, 305, 303, 3, 2, 2, 2, 305, 304, 3, 2, 2, 2, 306, 308,
	3, 2, 2, 2, 307, 300, 3, 2, 2, 2, 307, 301, 3, 2, 2, 2, 308, 25, 3, 2,
	2, 2, 309, 310, 7, 20, 2, 2, 310, 315, 5, 28, 15, 2, 311, 312, 7, 22, 2,
	2, 312, 314, 5, 28, 15, 2, 313, 311, 3, 2, 2, 2, 314, 317, 3, 2, 2, 2,
	315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3, 2, 2, 2, 317,
	315, 3, 2, 2, 2, 318, 319, 7, 21, 2, 2, 319, 27, 3, 2, 2, 2, 320, 321,
	5, 30, 16, 2, 321, 322, 7, 12, 2, 2, 322, 323, 5, 32, 17, 2, 323, 29, 3,
	2, 2, 2, 324, 336, 7, 28, 2, 2, 325, 326, 7, 17, 2, 2, 326, 331, 7, 28,
	2, 2, 327, 328, 7, 19, 2, 2, 328, 330, 7, 28, 2, 2, 329, 327, 3, 2, 2,
	2, 330, 333, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332,
	334, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2, 334, 336, 7, 18, 2, 2, 335, 324,
	3, 2, 2, 2, 335, 325, 3, 2, 2, 2, 336, 341, 3, 2, 2, 2, 337, 338, 7, 19,
	2, 2, 338, 340, 7, 28, 2, 2, 339, 337, 3, 2, 2, 2, 340, 343, 3, 2, 2, 2,
	341, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 31, 3, 2, 2, 2, 343, 341,
	3, 2, 2, 2, 344, 347, 7, 28, 2, 2, 345, 346, 7, 19, 2, 2, 346, 348, 7,
	28, 2, 2, 347, 345, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 347, 3, 2, 2,
	2, 349, 350, 3, 2, 2, 2, 350, 363, 3, 2, 2, 2, 351, 363, 7, 28, 2, 2, 352,
	354, 9, 2, 2, 2, 353, 352, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 355,
	3, 2, 2, 2, 355, 363, 7, 29, 2, 2, 356, 358, 9, 2, 2, 2, 357, 356, 3, 2,
	2, 2, 357, 358, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 363, 7, 30, 2, 2,
	360, 363, 7, 27, 2, 2, 361, 363, 5, 34, 18, 2, 362, 344, 3, 2, 2, 2, 362,
	351, 3, 2, 2, 2, 362, 353, 3, 2, 2, 2, 362, 357, 3, 2, 2, 2, 362, 360,
	3, 2, 2, 2, 362, 361, 3, 2, 2, 2, 363, 33, 3, 2, 2, 2, 364, 365, 7, 10,
	2, 2, 365, 372, 5, 36, 19, 2, 366, 368, 9, 3, 2, 2, 367, 366, 3, 2, 2,
	2, 367, 368, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 371, 5, 36, 19, 2,
	370, 367, 3, 2, 2, 2, 371, 374, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 372,
	373, 3, 2, 2, 2, 373, 375, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 375, 376,
	7, 11, 2, 2, 376, 35, 3, 2, 2, 2, 377, 378, 7, 28, 2, 2, 378, 379, 7, 16,
	2, 2, 379, 384, 5, 32, 17, 2, 380, 381, 7, 28, 2, 2, 381, 382, 7, 16, 2,
	2, 382, 384, 5, 34, 18, 2, 383, 377, 3, 2, 2, 2, 383, 380, 3, 2, 2, 2,
	384, 37, 3, 2, 2, 2, 47, 39, 44, 56, 66, 82, 89, 94, 107, 123, 127, 138,
	155, 163, 174, 184, 196, 199, 210, 225, 231, 236, 238, 245, 250, 253, 259,
	273, 275, 280, 287, 295, 298, 305, 307, 315, 331, 335, 341, 349, 353, 357,
	362, 367, 372, 383,
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
	"ID", "INT", "FLT", "WS", "COMMENT", "LINE_COMMENT", "DOWN", "UP", "ROOT",
	"ERROR", "Import", "Package", "Option", "Extend", "Message", "Enum", "Service",
	"Oneof", "Map", "Field", "Repeated", "Reserved", "Rpc", "Keytype", "EnumValue",
	"EnumNum", "INF", "NAN", "MAX", "WEAK", "PUBLIC", "Returns", "Stream",
	"To", "TRUE", "FALSE",
}

var ruleNames = []string{
	"proto", "syntax", "top_level_statement", "msg_enum_svc_ext", "associaton",
	"left_assoc", "right_assoc", "fullname", "messageType", "rpcDelim", "ranges",
	"rangee", "fieldOptions", "fieldOption", "optionName", "constant", "constantObj",
	"constantObjElem",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TronParser struct {
	*antlr.BaseParser
}

func NewTronParser(input antlr.TokenStream) *TronParser {
	this := new(TronParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TronParser.g4"

	return this
}

// TronParser tokens.
const (
	TronParserEOF              = antlr.TokenEOF
	TronParserPROTO3           = 1
	TronParserSYNTAX           = 2
	TronParserEQ_PRE           = 3
	TronParserENDPRE           = 4
	TronParserPRE_WS           = 5
	TronParserPRE_COMMENT      = 6
	TronParserPRE_LINE_COMMENT = 7
	TronParserLCUR             = 8
	TronParserRCUR             = 9
	TronParserEQ               = 10
	TronParserDQ               = 11
	TronParserSQ               = 12
	TronParserSEMI             = 13
	TronParserCOLON            = 14
	TronParserLPAREN           = 15
	TronParserRPAREN           = 16
	TronParserDOT              = 17
	TronParserLSB              = 18
	TronParserRSB              = 19
	TronParserCOMMA            = 20
	TronParserLCHEVR           = 21
	TronParserRCHEVR           = 22
	TronParserDASH             = 23
	TronParserPLUS             = 24
	TronParserSTR              = 25
	TronParserID               = 26
	TronParserINT              = 27
	TronParserFLT              = 28
	TronParserWS               = 29
	TronParserCOMMENT          = 30
	TronParserLINE_COMMENT     = 31
	TronParserDOWN             = 32
	TronParserUP               = 33
	TronParserROOT             = 34
	TronParserERROR            = 35
	TronParserImport           = 36
	TronParserPackage          = 37
	TronParserOption           = 38
	TronParserExtend           = 39
	TronParserMessage          = 40
	TronParserEnum             = 41
	TronParserService          = 42
	TronParserOneof            = 43
	TronParserMap              = 44
	TronParserField            = 45
	TronParserRepeated         = 46
	TronParserReserved         = 47
	TronParserRpc              = 48
	TronParserKeytype          = 49
	TronParserEnumValue        = 50
	TronParserEnumNum          = 51
	TronParserINF              = 52
	TronParserNAN              = 53
	TronParserMAX              = 54
	TronParserWEAK             = 55
	TronParserPUBLIC           = 56
	TronParserReturns          = 57
	TronParserStream           = 58
	TronParserTo               = 59
	TronParserTRUE             = 60
	TronParserFALSE            = 61
)

// TronParser rules.
const (
	TronParserRULE_proto               = 0
	TronParserRULE_syntax              = 1
	TronParserRULE_top_level_statement = 2
	TronParserRULE_msg_enum_svc_ext    = 3
	TronParserRULE_associaton          = 4
	TronParserRULE_left_assoc          = 5
	TronParserRULE_right_assoc         = 6
	TronParserRULE_fullname            = 7
	TronParserRULE_messageType         = 8
	TronParserRULE_rpcDelim            = 9
	TronParserRULE_ranges              = 10
	TronParserRULE_rangee              = 11
	TronParserRULE_fieldOptions        = 12
	TronParserRULE_fieldOption         = 13
	TronParserRULE_optionName          = 14
	TronParserRULE_constant            = 15
	TronParserRULE_constantObj         = 16
	TronParserRULE_constantObjElem     = 17
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
	p.RuleIndex = TronParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(TronParserEOF, 0)
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
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *TronParser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TronParserRULE_proto)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TronParserSYNTAX {
		{
			p.SetState(36)
			p.Syntax()
		}

	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TronParserSEMI || _la == TronParserID {
		{
			p.SetState(39)
			p.Top_level_statement()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.Match(TronParserEOF)
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
	p.RuleIndex = TronParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(TronParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ_PRE() antlr.TerminalNode {
	return s.GetToken(TronParserEQ_PRE, 0)
}

func (s *SyntaxContext) PROTO3() antlr.TerminalNode {
	return s.GetToken(TronParserPROTO3, 0)
}

func (s *SyntaxContext) ENDPRE() antlr.TerminalNode {
	return s.GetToken(TronParserENDPRE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *TronParser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TronParserRULE_syntax)

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
		p.SetState(47)
		p.Match(TronParserSYNTAX)
	}
	{
		p.SetState(48)
		p.Match(TronParserEQ_PRE)
	}
	{
		p.SetState(49)
		p.Match(TronParserPROTO3)
	}
	{
		p.SetState(50)
		p.Match(TronParserENDPRE)
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
	p.RuleIndex = TronParserRULE_top_level_statement
	return p
}

func (*Top_level_statementContext) IsTop_level_statementContext() {}

func NewTop_level_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_level_statementContext {
	var p = new(Top_level_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_top_level_statement

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

type OptionFileDefContext struct {
	*Top_level_statementContext
	opt antlr.Token
}

func NewOptionFileDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionFileDefContext {
	var p = new(OptionFileDefContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *OptionFileDefContext) GetOpt() antlr.Token { return s.opt }

func (s *OptionFileDefContext) SetOpt(v antlr.Token) { s.opt = v }

func (s *OptionFileDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionFileDefContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionFileDefContext) EQ() antlr.TerminalNode {
	return s.GetToken(TronParserEQ, 0)
}

func (s *OptionFileDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionFileDefContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *OptionFileDefContext) ID() antlr.TerminalNode {
	return s.GetToken(TronParserID, 0)
}

func (s *OptionFileDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterOptionFileDef(s)
	}
}

func (s *OptionFileDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitOptionFileDef(s)
	}
}

type Ext_Msg_Enum_SvcContext struct {
	*Top_level_statementContext
	mese antlr.Token
	_ID  antlr.Token
	a    []antlr.Token
}

func NewExt_Msg_Enum_SvcContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Ext_Msg_Enum_SvcContext {
	var p = new(Ext_Msg_Enum_SvcContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *Ext_Msg_Enum_SvcContext) GetMese() antlr.Token { return s.mese }

func (s *Ext_Msg_Enum_SvcContext) Get_ID() antlr.Token { return s._ID }

func (s *Ext_Msg_Enum_SvcContext) SetMese(v antlr.Token) { s.mese = v }

func (s *Ext_Msg_Enum_SvcContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Ext_Msg_Enum_SvcContext) GetA() []antlr.Token { return s.a }

func (s *Ext_Msg_Enum_SvcContext) SetA(v []antlr.Token) { s.a = v }

func (s *Ext_Msg_Enum_SvcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_Msg_Enum_SvcContext) LCUR() antlr.TerminalNode {
	return s.GetToken(TronParserLCUR, 0)
}

func (s *Ext_Msg_Enum_SvcContext) RCUR() antlr.TerminalNode {
	return s.GetToken(TronParserRCUR, 0)
}

func (s *Ext_Msg_Enum_SvcContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *Ext_Msg_Enum_SvcContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *Ext_Msg_Enum_SvcContext) AllMsg_enum_svc_ext() []IMsg_enum_svc_extContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsg_enum_svc_extContext)(nil)).Elem())
	var tst = make([]IMsg_enum_svc_extContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsg_enum_svc_extContext)
		}
	}

	return tst
}

func (s *Ext_Msg_Enum_SvcContext) Msg_enum_svc_ext(i int) IMsg_enum_svc_extContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsg_enum_svc_extContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsg_enum_svc_extContext)
}

func (s *Ext_Msg_Enum_SvcContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *Ext_Msg_Enum_SvcContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *Ext_Msg_Enum_SvcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterExt_Msg_Enum_Svc(s)
	}
}

func (s *Ext_Msg_Enum_SvcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitExt_Msg_Enum_Svc(s)
	}
}

type ImportStatementContext struct {
	*Top_level_statementContext
	imp          antlr.Token
	weakORpublic antlr.Token
	b            antlr.Token
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

func (s *ImportStatementContext) GetB() antlr.Token { return s.b }

func (s *ImportStatementContext) SetImp(v antlr.Token) { s.imp = v }

func (s *ImportStatementContext) SetWeakORpublic(v antlr.Token) { s.weakORpublic = v }

func (s *ImportStatementContext) SetB(v antlr.Token) { s.b = v }

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *ImportStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *ImportStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *ImportStatementContext) STR() antlr.TerminalNode {
	return s.GetToken(TronParserSTR, 0)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitImportStatement(s)
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
	return s.GetToken(TronParserSEMI, 0)
}

func (s *EmptyStmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterEmptyStm(s)
	}
}

func (s *EmptyStmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitEmptyStm(s)
	}
}

type PackageStatementContext struct {
	*Top_level_statementContext
	pkg antlr.Token
	_ID antlr.Token
	a   []antlr.Token
}

func NewPackageStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *PackageStatementContext) GetPkg() antlr.Token { return s.pkg }

func (s *PackageStatementContext) Get_ID() antlr.Token { return s._ID }

func (s *PackageStatementContext) SetPkg(v antlr.Token) { s.pkg = v }

func (s *PackageStatementContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PackageStatementContext) GetA() []antlr.Token { return s.a }

func (s *PackageStatementContext) SetA(v []antlr.Token) { s.a = v }

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *PackageStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *PackageStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *PackageStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *PackageStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *TronParser) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TronParserRULE_top_level_statement)
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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewImportStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)

			var _m = p.Match(TronParserID)

			localctx.(*ImportStatementContext).imp = _m
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserID {
			{
				p.SetState(53)

				var _m = p.Match(TronParserID)

				localctx.(*ImportStatementContext).weakORpublic = _m
			}

		}
		{
			p.SetState(56)

			var _m = p.Match(TronParserSTR)

			localctx.(*ImportStatementContext).b = _m
		}
		{
			p.SetState(57)
			p.Match(TronParserSEMI)
		}

	case 2:
		localctx = NewPackageStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)

			var _m = p.Match(TronParserID)

			localctx.(*PackageStatementContext).pkg = _m
		}
		{
			p.SetState(59)

			var _m = p.Match(TronParserID)

			localctx.(*PackageStatementContext)._ID = _m
		}
		localctx.(*PackageStatementContext).a = append(localctx.(*PackageStatementContext).a, localctx.(*PackageStatementContext)._ID)
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserDOT {
			{
				p.SetState(60)
				p.Match(TronParserDOT)
			}
			{
				p.SetState(61)

				var _m = p.Match(TronParserID)

				localctx.(*PackageStatementContext)._ID = _m
			}
			localctx.(*PackageStatementContext).a = append(localctx.(*PackageStatementContext).a, localctx.(*PackageStatementContext)._ID)

			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(67)
			p.Match(TronParserSEMI)
		}

	case 3:
		localctx = NewOptionFileDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)

			var _m = p.Match(TronParserID)

			localctx.(*OptionFileDefContext).opt = _m
		}
		{
			p.SetState(69)
			p.OptionName()
		}
		{
			p.SetState(70)
			p.Match(TronParserEQ)
		}
		{
			p.SetState(71)
			p.Constant()
		}
		{
			p.SetState(72)
			p.Match(TronParserSEMI)
		}

	case 4:
		localctx = NewExt_Msg_Enum_SvcContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)

			var _m = p.Match(TronParserID)

			localctx.(*Ext_Msg_Enum_SvcContext).mese = _m
		}

		{
			p.SetState(75)

			var _m = p.Match(TronParserID)

			localctx.(*Ext_Msg_Enum_SvcContext)._ID = _m
		}
		localctx.(*Ext_Msg_Enum_SvcContext).a = append(localctx.(*Ext_Msg_Enum_SvcContext).a, localctx.(*Ext_Msg_Enum_SvcContext)._ID)
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserDOT {
			{
				p.SetState(76)
				p.Match(TronParserDOT)
			}
			{
				p.SetState(77)

				var _m = p.Match(TronParserID)

				localctx.(*Ext_Msg_Enum_SvcContext)._ID = _m
			}
			localctx.(*Ext_Msg_Enum_SvcContext).a = append(localctx.(*Ext_Msg_Enum_SvcContext).a, localctx.(*Ext_Msg_Enum_SvcContext)._ID)

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		{
			p.SetState(83)
			p.Match(TronParserLCUR)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TronParserSEMI)|(1<<TronParserDOT)|(1<<TronParserID))) != 0 {
			{
				p.SetState(84)
				p.Msg_enum_svc_ext()
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(90)
			p.Match(TronParserRCUR)
		}

	case 5:
		localctx = NewEmptyStmContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.Match(TronParserSEMI)
		}

	}

	return localctx
}

// IMsg_enum_svc_extContext is an interface to support dynamic dispatch.
type IMsg_enum_svc_extContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsg_enum_svc_extContext differentiates from other interfaces.
	IsMsg_enum_svc_extContext()
}

type Msg_enum_svc_extContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsg_enum_svc_extContext() *Msg_enum_svc_extContext {
	var p = new(Msg_enum_svc_extContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_msg_enum_svc_ext
	return p
}

func (*Msg_enum_svc_extContext) IsMsg_enum_svc_extContext() {}

func NewMsg_enum_svc_extContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Msg_enum_svc_extContext {
	var p = new(Msg_enum_svc_extContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_msg_enum_svc_ext

	return p
}

func (s *Msg_enum_svc_extContext) GetParser() antlr.Parser { return s.parser }

func (s *Msg_enum_svc_extContext) CopyFrom(ctx *Msg_enum_svc_extContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Msg_enum_svc_extContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Msg_enum_svc_extContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EmptyTopLvlContext struct {
	*Msg_enum_svc_extContext
	msg antlr.Token
}

func NewEmptyTopLvlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyTopLvlContext {
	var p = new(EmptyTopLvlContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *EmptyTopLvlContext) GetMsg() antlr.Token { return s.msg }

func (s *EmptyTopLvlContext) SetMsg(v antlr.Token) { s.msg = v }

func (s *EmptyTopLvlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyTopLvlContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *EmptyTopLvlContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *EmptyTopLvlContext) LCUR() antlr.TerminalNode {
	return s.GetToken(TronParserLCUR, 0)
}

func (s *EmptyTopLvlContext) RCUR() antlr.TerminalNode {
	return s.GetToken(TronParserRCUR, 0)
}

func (s *EmptyTopLvlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterEmptyTopLvl(s)
	}
}

func (s *EmptyTopLvlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitEmptyTopLvl(s)
	}
}

type EmptyTopLvlStmContext struct {
	*Msg_enum_svc_extContext
	msg antlr.Token
}

func NewEmptyTopLvlStmContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyTopLvlStmContext {
	var p = new(EmptyTopLvlStmContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *EmptyTopLvlStmContext) GetMsg() antlr.Token { return s.msg }

func (s *EmptyTopLvlStmContext) SetMsg(v antlr.Token) { s.msg = v }

func (s *EmptyTopLvlStmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyTopLvlStmContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *EmptyTopLvlStmContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *EmptyTopLvlStmContext) LCUR() antlr.TerminalNode {
	return s.GetToken(TronParserLCUR, 0)
}

func (s *EmptyTopLvlStmContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *EmptyTopLvlStmContext) RCUR() antlr.TerminalNode {
	return s.GetToken(TronParserRCUR, 0)
}

func (s *EmptyTopLvlStmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterEmptyTopLvlStm(s)
	}
}

func (s *EmptyTopLvlStmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitEmptyTopLvlStm(s)
	}
}

type TLIOptionContext struct {
	*Msg_enum_svc_extContext
}

func NewTLIOptionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TLIOptionContext {
	var p = new(TLIOptionContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *TLIOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TLIOptionContext) ID() antlr.TerminalNode {
	return s.GetToken(TronParserID, 0)
}

func (s *TLIOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(TronParserEQ, 0)
}

func (s *TLIOptionContext) INT() antlr.TerminalNode {
	return s.GetToken(TronParserINT, 0)
}

func (s *TLIOptionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *TLIOptionContext) DASH() antlr.TerminalNode {
	return s.GetToken(TronParserDASH, 0)
}

func (s *TLIOptionContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *TLIOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterTLIOption(s)
	}
}

func (s *TLIOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitTLIOption(s)
	}
}

type EmptyStmStmContext struct {
	*Msg_enum_svc_extContext
}

func NewEmptyStmStmContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyStmStmContext {
	var p = new(EmptyStmStmContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *EmptyStmStmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStmStmContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *EmptyStmStmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterEmptyStmStm(s)
	}
}

func (s *EmptyStmStmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitEmptyStmStm(s)
	}
}

type RPCSigContext struct {
	*Msg_enum_svc_extContext
	rpcID antlr.Token
	rets  antlr.Token
}

func NewRPCSigContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RPCSigContext {
	var p = new(RPCSigContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *RPCSigContext) GetRpcID() antlr.Token { return s.rpcID }

func (s *RPCSigContext) GetRets() antlr.Token { return s.rets }

func (s *RPCSigContext) SetRpcID(v antlr.Token) { s.rpcID = v }

func (s *RPCSigContext) SetRets(v antlr.Token) { s.rets = v }

func (s *RPCSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RPCSigContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *RPCSigContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *RPCSigContext) AllMessageType() []IMessageTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem())
	var tst = make([]IMessageTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageTypeContext)
		}
	}

	return tst
}

func (s *RPCSigContext) MessageType(i int) IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RPCSigContext) RpcDelim() IRpcDelimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcDelimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcDelimContext)
}

func (s *RPCSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRPCSig(s)
	}
}

func (s *RPCSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRPCSig(s)
	}
}

type AssocContext struct {
	*Msg_enum_svc_extContext
}

func NewAssocContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssocContext {
	var p = new(AssocContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *AssocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssocContext) Associaton() IAssociatonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociatonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociatonContext)
}

func (s *AssocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterAssoc(s)
	}
}

func (s *AssocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitAssoc(s)
	}
}

type RangeContext struct {
	*Msg_enum_svc_extContext
	res antlr.Token
}

func NewRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeContext {
	var p = new(RangeContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *RangeContext) GetRes() antlr.Token { return s.res }

func (s *RangeContext) SetRes(v antlr.Token) { s.res = v }

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *RangeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *RangeContext) ID() antlr.TerminalNode {
	return s.GetToken(TronParserID, 0)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRange(s)
	}
}

type MsgEnumSvcExtContext struct {
	*Msg_enum_svc_extContext
	mese antlr.Token
}

func NewMsgEnumSvcExtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgEnumSvcExtContext {
	var p = new(MsgEnumSvcExtContext)

	p.Msg_enum_svc_extContext = NewEmptyMsg_enum_svc_extContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Msg_enum_svc_extContext))

	return p
}

func (s *MsgEnumSvcExtContext) GetMese() antlr.Token { return s.mese }

func (s *MsgEnumSvcExtContext) SetMese(v antlr.Token) { s.mese = v }

func (s *MsgEnumSvcExtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgEnumSvcExtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *MsgEnumSvcExtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *MsgEnumSvcExtContext) LCUR() antlr.TerminalNode {
	return s.GetToken(TronParserLCUR, 0)
}

func (s *MsgEnumSvcExtContext) RCUR() antlr.TerminalNode {
	return s.GetToken(TronParserRCUR, 0)
}

func (s *MsgEnumSvcExtContext) AllMsg_enum_svc_ext() []IMsg_enum_svc_extContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsg_enum_svc_extContext)(nil)).Elem())
	var tst = make([]IMsg_enum_svc_extContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsg_enum_svc_extContext)
		}
	}

	return tst
}

func (s *MsgEnumSvcExtContext) Msg_enum_svc_ext(i int) IMsg_enum_svc_extContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsg_enum_svc_extContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsg_enum_svc_extContext)
}

func (s *MsgEnumSvcExtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterMsgEnumSvcExt(s)
	}
}

func (s *MsgEnumSvcExtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitMsgEnumSvcExt(s)
	}
}

func (p *TronParser) Msg_enum_svc_ext() (localctx IMsg_enum_svc_extContext) {
	localctx = NewMsg_enum_svc_extContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TronParserRULE_msg_enum_svc_ext)
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

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssocContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Associaton()
		}

	case 2:
		localctx = NewEmptyTopLvlContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)

			var _m = p.Match(TronParserID)

			localctx.(*EmptyTopLvlContext).msg = _m
		}
		{
			p.SetState(96)
			p.Match(TronParserID)
		}
		{
			p.SetState(97)
			p.Match(TronParserLCUR)
		}
		{
			p.SetState(98)
			p.Match(TronParserRCUR)
		}

	case 3:
		localctx = NewMsgEnumSvcExtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)

			var _m = p.Match(TronParserID)

			localctx.(*MsgEnumSvcExtContext).mese = _m
		}
		{
			p.SetState(100)
			p.Match(TronParserID)
		}
		{
			p.SetState(101)
			p.Match(TronParserLCUR)
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TronParserSEMI)|(1<<TronParserDOT)|(1<<TronParserID))) != 0) {
			{
				p.SetState(102)
				p.Msg_enum_svc_ext()
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Match(TronParserRCUR)
		}

	case 4:
		localctx = NewEmptyTopLvlStmContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)

			var _m = p.Match(TronParserID)

			localctx.(*EmptyTopLvlStmContext).msg = _m
		}
		{
			p.SetState(110)
			p.Match(TronParserID)
		}
		{
			p.SetState(111)
			p.Match(TronParserLCUR)
		}
		{
			p.SetState(112)
			p.Match(TronParserSEMI)
		}
		{
			p.SetState(113)
			p.Match(TronParserRCUR)
		}

	case 5:
		localctx = NewRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)

			var _m = p.Match(TronParserID)

			localctx.(*RangeContext).res = _m
		}
		{
			p.SetState(115)
			p.Ranges()
		}
		{
			p.SetState(116)
			p.Match(TronParserSEMI)
		}

	case 6:
		localctx = NewTLIOptionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(118)
			p.Match(TronParserID)
		}
		{
			p.SetState(119)
			p.Match(TronParserEQ)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserDASH {
			{
				p.SetState(120)
				p.Match(TronParserDASH)
			}

		}
		{
			p.SetState(123)
			p.Match(TronParserINT)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserLSB {
			{
				p.SetState(124)
				p.FieldOptions()
			}

		}
		{
			p.SetState(127)
			p.Match(TronParserSEMI)
		}

	case 7:
		localctx = NewRPCSigContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(128)

			var _m = p.Match(TronParserID)

			localctx.(*RPCSigContext).rpcID = _m
		}
		{
			p.SetState(129)
			p.Match(TronParserID)
		}
		{
			p.SetState(130)
			p.MessageType()
		}
		{
			p.SetState(131)

			var _m = p.Match(TronParserID)

			localctx.(*RPCSigContext).rets = _m
		}
		{
			p.SetState(132)
			p.MessageType()
		}
		{
			p.SetState(133)
			p.RpcDelim()
		}

	case 8:
		localctx = NewEmptyStmStmContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)
			p.Match(TronParserSEMI)
		}

	}

	return localctx
}

// IAssociatonContext is an interface to support dynamic dispatch.
type IAssociatonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociatonContext differentiates from other interfaces.
	IsAssociatonContext()
}

type AssociatonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociatonContext() *AssociatonContext {
	var p = new(AssociatonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_associaton
	return p
}

func (*AssociatonContext) IsAssociatonContext() {}

func NewAssociatonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociatonContext {
	var p = new(AssociatonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_associaton

	return p
}

func (s *AssociatonContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociatonContext) Left_assoc() ILeft_assocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeft_assocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeft_assocContext)
}

func (s *AssociatonContext) EQ() antlr.TerminalNode {
	return s.GetToken(TronParserEQ, 0)
}

func (s *AssociatonContext) Right_assoc() IRight_assocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRight_assocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRight_assocContext)
}

func (s *AssociatonContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, 0)
}

func (s *AssociatonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociatonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociatonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterAssociaton(s)
	}
}

func (s *AssociatonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitAssociaton(s)
	}
}

func (p *TronParser) Associaton() (localctx IAssociatonContext) {
	localctx = NewAssociatonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TronParserRULE_associaton)

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
		p.SetState(138)
		p.Left_assoc()
	}
	{
		p.SetState(139)
		p.Match(TronParserEQ)
	}
	{
		p.SetState(140)
		p.Right_assoc()
	}
	{
		p.SetState(141)
		p.Match(TronParserSEMI)
	}

	return localctx
}

// ILeft_assocContext is an interface to support dynamic dispatch.
type ILeft_assocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeft_assocContext differentiates from other interfaces.
	IsLeft_assocContext()
}

type Left_assocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeft_assocContext() *Left_assocContext {
	var p = new(Left_assocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_left_assoc
	return p
}

func (*Left_assocContext) IsLeft_assocContext() {}

func NewLeft_assocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Left_assocContext {
	var p = new(Left_assocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_left_assoc

	return p
}

func (s *Left_assocContext) GetParser() antlr.Parser { return s.parser }

func (s *Left_assocContext) CopyFrom(ctx *Left_assocContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Left_assocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Left_assocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleLocalContext struct {
	*Left_assocContext
	_ID antlr.Token
	a   []antlr.Token
}

func NewSingleLocalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleLocalContext {
	var p = new(SingleLocalContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *SingleLocalContext) Get_ID() antlr.Token { return s._ID }

func (s *SingleLocalContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *SingleLocalContext) GetA() []antlr.Token { return s.a }

func (s *SingleLocalContext) SetA(v []antlr.Token) { s.a = v }

func (s *SingleLocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleLocalContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *SingleLocalContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *SingleLocalContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *SingleLocalContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *SingleLocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterSingleLocal(s)
	}
}

func (s *SingleLocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitSingleLocal(s)
	}
}

type EnumLeftContext struct {
	*Left_assocContext
	a antlr.Token
}

func NewEnumLeftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumLeftContext {
	var p = new(EnumLeftContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *EnumLeftContext) GetA() antlr.Token { return s.a }

func (s *EnumLeftContext) SetA(v antlr.Token) { s.a = v }

func (s *EnumLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumLeftContext) ID() antlr.TerminalNode {
	return s.GetToken(TronParserID, 0)
}

func (s *EnumLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterEnumLeft(s)
	}
}

func (s *EnumLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitEnumLeft(s)
	}
}

type OptContext struct {
	*Left_assocContext
	a antlr.Token
}

func NewOptContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptContext {
	var p = new(OptContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *OptContext) GetA() antlr.Token { return s.a }

func (s *OptContext) SetA(v antlr.Token) { s.a = v }

func (s *OptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TronParserLPAREN, 0)
}

func (s *OptContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *OptContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *OptContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TronParserRPAREN, 0)
}

func (s *OptContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *OptContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *OptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterOpt(s)
	}
}

func (s *OptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitOpt(s)
	}
}

type Opt_SingleContext struct {
	*Left_assocContext
	a antlr.Token
	b antlr.Token
}

func NewOpt_SingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Opt_SingleContext {
	var p = new(Opt_SingleContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *Opt_SingleContext) GetA() antlr.Token { return s.a }

func (s *Opt_SingleContext) GetB() antlr.Token { return s.b }

func (s *Opt_SingleContext) SetA(v antlr.Token) { s.a = v }

func (s *Opt_SingleContext) SetB(v antlr.Token) { s.b = v }

func (s *Opt_SingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Opt_SingleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *Opt_SingleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *Opt_SingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterOpt_Single(s)
	}
}

func (s *Opt_SingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitOpt_Single(s)
	}
}

type MapLocalLeftContext struct {
	*Left_assocContext
	mpt antlr.Token
	kt  antlr.Token
	_ID antlr.Token
	v   []antlr.Token
	c   antlr.Token
}

func NewMapLocalLeftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapLocalLeftContext {
	var p = new(MapLocalLeftContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *MapLocalLeftContext) GetMpt() antlr.Token { return s.mpt }

func (s *MapLocalLeftContext) GetKt() antlr.Token { return s.kt }

func (s *MapLocalLeftContext) Get_ID() antlr.Token { return s._ID }

func (s *MapLocalLeftContext) GetC() antlr.Token { return s.c }

func (s *MapLocalLeftContext) SetMpt(v antlr.Token) { s.mpt = v }

func (s *MapLocalLeftContext) SetKt(v antlr.Token) { s.kt = v }

func (s *MapLocalLeftContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *MapLocalLeftContext) SetC(v antlr.Token) { s.c = v }

func (s *MapLocalLeftContext) GetV() []antlr.Token { return s.v }

func (s *MapLocalLeftContext) SetV(v []antlr.Token) { s.v = v }

func (s *MapLocalLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLocalLeftContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(TronParserLCHEVR, 0)
}

func (s *MapLocalLeftContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TronParserCOMMA, 0)
}

func (s *MapLocalLeftContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *MapLocalLeftContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *MapLocalLeftContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(TronParserRCHEVR, 0)
}

func (s *MapLocalLeftContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *MapLocalLeftContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *MapLocalLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterMapLocalLeft(s)
	}
}

func (s *MapLocalLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitMapLocalLeft(s)
	}
}

type MapLeftContext struct {
	*Left_assocContext
	mpt antlr.Token
	kt  antlr.Token
	_ID antlr.Token
	v   []antlr.Token
	c   antlr.Token
}

func NewMapLeftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapLeftContext {
	var p = new(MapLeftContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *MapLeftContext) GetMpt() antlr.Token { return s.mpt }

func (s *MapLeftContext) GetKt() antlr.Token { return s.kt }

func (s *MapLeftContext) Get_ID() antlr.Token { return s._ID }

func (s *MapLeftContext) GetC() antlr.Token { return s.c }

func (s *MapLeftContext) SetMpt(v antlr.Token) { s.mpt = v }

func (s *MapLeftContext) SetKt(v antlr.Token) { s.kt = v }

func (s *MapLeftContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *MapLeftContext) SetC(v antlr.Token) { s.c = v }

func (s *MapLeftContext) GetV() []antlr.Token { return s.v }

func (s *MapLeftContext) SetV(v []antlr.Token) { s.v = v }

func (s *MapLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLeftContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(TronParserLCHEVR, 0)
}

func (s *MapLeftContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TronParserCOMMA, 0)
}

func (s *MapLeftContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(TronParserRCHEVR, 0)
}

func (s *MapLeftContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *MapLeftContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *MapLeftContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *MapLeftContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *MapLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterMapLeft(s)
	}
}

func (s *MapLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitMapLeft(s)
	}
}

type RepeatedContext struct {
	*Left_assocContext
	a   antlr.Token
	_ID antlr.Token
	b   []antlr.Token
	c   antlr.Token
}

func NewRepeatedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatedContext {
	var p = new(RepeatedContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *RepeatedContext) GetA() antlr.Token { return s.a }

func (s *RepeatedContext) Get_ID() antlr.Token { return s._ID }

func (s *RepeatedContext) GetC() antlr.Token { return s.c }

func (s *RepeatedContext) SetA(v antlr.Token) { s.a = v }

func (s *RepeatedContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RepeatedContext) SetC(v antlr.Token) { s.c = v }

func (s *RepeatedContext) GetB() []antlr.Token { return s.b }

func (s *RepeatedContext) SetB(v []antlr.Token) { s.b = v }

func (s *RepeatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatedContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *RepeatedContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *RepeatedContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *RepeatedContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *RepeatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRepeated(s)
	}
}

func (s *RepeatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRepeated(s)
	}
}

type SingleFull_RepLocalContext struct {
	*Left_assocContext
	a   antlr.Token
	_ID antlr.Token
	b   []antlr.Token
	c   antlr.Token
}

func NewSingleFull_RepLocalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleFull_RepLocalContext {
	var p = new(SingleFull_RepLocalContext)

	p.Left_assocContext = NewEmptyLeft_assocContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Left_assocContext))

	return p
}

func (s *SingleFull_RepLocalContext) GetA() antlr.Token { return s.a }

func (s *SingleFull_RepLocalContext) Get_ID() antlr.Token { return s._ID }

func (s *SingleFull_RepLocalContext) GetC() antlr.Token { return s.c }

func (s *SingleFull_RepLocalContext) SetA(v antlr.Token) { s.a = v }

func (s *SingleFull_RepLocalContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *SingleFull_RepLocalContext) SetC(v antlr.Token) { s.c = v }

func (s *SingleFull_RepLocalContext) GetB() []antlr.Token { return s.b }

func (s *SingleFull_RepLocalContext) SetB(v []antlr.Token) { s.b = v }

func (s *SingleFull_RepLocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleFull_RepLocalContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *SingleFull_RepLocalContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *SingleFull_RepLocalContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *SingleFull_RepLocalContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *SingleFull_RepLocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterSingleFull_RepLocal(s)
	}
}

func (s *SingleFull_RepLocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitSingleFull_RepLocal(s)
	}
}

func (p *TronParser) Left_assoc() (localctx ILeft_assocContext) {
	localctx = NewLeft_assocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TronParserRULE_left_assoc)
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

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEnumLeftContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)

			var _m = p.Match(TronParserID)

			localctx.(*EnumLeftContext).a = _m
		}

	case 2:
		localctx = NewOpt_SingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)

			var _m = p.Match(TronParserID)

			localctx.(*Opt_SingleContext).a = _m
		}
		{
			p.SetState(145)

			var _m = p.Match(TronParserID)

			localctx.(*Opt_SingleContext).b = _m
		}

	case 3:
		localctx = NewOptContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)

			var _m = p.Match(TronParserID)

			localctx.(*OptContext).a = _m
		}
		{
			p.SetState(147)
			p.Match(TronParserLPAREN)
		}
		{
			p.SetState(148)
			p.Match(TronParserID)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserDOT {
			{
				p.SetState(149)
				p.Match(TronParserDOT)
			}
			{
				p.SetState(150)
				p.Match(TronParserID)
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(156)
			p.Match(TronParserRPAREN)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserDOT {
			{
				p.SetState(157)
				p.Match(TronParserDOT)
			}
			{
				p.SetState(158)
				p.Match(TronParserID)
			}

			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		localctx = NewSingleFull_RepLocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)

			var _m = p.Match(TronParserID)

			localctx.(*SingleFull_RepLocalContext).a = _m
		}
		{
			p.SetState(165)
			p.Match(TronParserDOT)
		}
		{
			p.SetState(166)

			var _m = p.Match(TronParserID)

			localctx.(*SingleFull_RepLocalContext)._ID = _m
		}
		localctx.(*SingleFull_RepLocalContext).b = append(localctx.(*SingleFull_RepLocalContext).b, localctx.(*SingleFull_RepLocalContext)._ID)
		{
			p.SetState(167)
			p.Match(TronParserDOT)
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(168)

					var _m = p.Match(TronParserID)

					localctx.(*SingleFull_RepLocalContext)._ID = _m
				}
				localctx.(*SingleFull_RepLocalContext).b = append(localctx.(*SingleFull_RepLocalContext).b, localctx.(*SingleFull_RepLocalContext)._ID)
				{
					p.SetState(169)
					p.Match(TronParserDOT)
				}

			}
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}
		{
			p.SetState(175)

			var _m = p.Match(TronParserID)

			localctx.(*SingleFull_RepLocalContext)._ID = _m
		}
		localctx.(*SingleFull_RepLocalContext).b = append(localctx.(*SingleFull_RepLocalContext).b, localctx.(*SingleFull_RepLocalContext)._ID)
		{
			p.SetState(176)

			var _m = p.Match(TronParserID)

			localctx.(*SingleFull_RepLocalContext).c = _m
		}

	case 5:
		localctx = NewSingleLocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Match(TronParserDOT)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(178)

					var _m = p.Match(TronParserID)

					localctx.(*SingleLocalContext)._ID = _m
				}
				localctx.(*SingleLocalContext).a = append(localctx.(*SingleLocalContext).a, localctx.(*SingleLocalContext)._ID)
				{
					p.SetState(179)
					p.Match(TronParserDOT)
				}

			}
			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}
		{
			p.SetState(185)

			var _m = p.Match(TronParserID)

			localctx.(*SingleLocalContext)._ID = _m
		}
		localctx.(*SingleLocalContext).a = append(localctx.(*SingleLocalContext).a, localctx.(*SingleLocalContext)._ID)
		{
			p.SetState(186)
			p.Match(TronParserID)
		}

	case 6:
		localctx = NewRepeatedContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(187)

			var _m = p.Match(TronParserID)

			localctx.(*RepeatedContext).a = _m
		}
		{
			p.SetState(188)

			var _m = p.Match(TronParserID)

			localctx.(*RepeatedContext)._ID = _m
		}
		localctx.(*RepeatedContext).b = append(localctx.(*RepeatedContext).b, localctx.(*RepeatedContext)._ID)
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserDOT {
			{
				p.SetState(189)
				p.Match(TronParserDOT)
			}
			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(190)

						var _m = p.Match(TronParserID)

						localctx.(*RepeatedContext)._ID = _m
					}
					localctx.(*RepeatedContext).b = append(localctx.(*RepeatedContext).b, localctx.(*RepeatedContext)._ID)
					{
						p.SetState(191)
						p.Match(TronParserDOT)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(194)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
			}
			{
				p.SetState(196)

				var _m = p.Match(TronParserID)

				localctx.(*RepeatedContext)._ID = _m
			}
			localctx.(*RepeatedContext).b = append(localctx.(*RepeatedContext).b, localctx.(*RepeatedContext)._ID)

		}
		{
			p.SetState(199)

			var _m = p.Match(TronParserID)

			localctx.(*RepeatedContext).c = _m
		}

	case 7:
		localctx = NewMapLeftContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(200)

			var _m = p.Match(TronParserID)

			localctx.(*MapLeftContext).mpt = _m
		}
		{
			p.SetState(201)
			p.Match(TronParserLCHEVR)
		}
		{
			p.SetState(202)

			var _m = p.Match(TronParserID)

			localctx.(*MapLeftContext).kt = _m
		}
		{
			p.SetState(203)
			p.Match(TronParserCOMMA)
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(204)

					var _m = p.Match(TronParserID)

					localctx.(*MapLeftContext)._ID = _m
				}
				localctx.(*MapLeftContext).v = append(localctx.(*MapLeftContext).v, localctx.(*MapLeftContext)._ID)
				{
					p.SetState(205)
					p.Match(TronParserDOT)
				}

			}
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}
		{
			p.SetState(211)

			var _m = p.Match(TronParserID)

			localctx.(*MapLeftContext)._ID = _m
		}
		localctx.(*MapLeftContext).v = append(localctx.(*MapLeftContext).v, localctx.(*MapLeftContext)._ID)
		{
			p.SetState(212)
			p.Match(TronParserRCHEVR)
		}
		{
			p.SetState(213)

			var _m = p.Match(TronParserID)

			localctx.(*MapLeftContext).c = _m
		}

	case 8:
		localctx = NewMapLocalLeftContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(214)

			var _m = p.Match(TronParserID)

			localctx.(*MapLocalLeftContext).mpt = _m
		}
		{
			p.SetState(215)
			p.Match(TronParserLCHEVR)
		}
		{
			p.SetState(216)

			var _m = p.Match(TronParserID)

			localctx.(*MapLocalLeftContext).kt = _m
		}
		{
			p.SetState(217)
			p.Match(TronParserCOMMA)
		}
		{
			p.SetState(218)
			p.Match(TronParserDOT)
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(219)

					var _m = p.Match(TronParserID)

					localctx.(*MapLocalLeftContext)._ID = _m
				}
				localctx.(*MapLocalLeftContext).v = append(localctx.(*MapLocalLeftContext).v, localctx.(*MapLocalLeftContext)._ID)
				{
					p.SetState(220)
					p.Match(TronParserDOT)
				}

			}
			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}
		{
			p.SetState(226)

			var _m = p.Match(TronParserID)

			localctx.(*MapLocalLeftContext)._ID = _m
		}
		localctx.(*MapLocalLeftContext).v = append(localctx.(*MapLocalLeftContext).v, localctx.(*MapLocalLeftContext)._ID)
		{
			p.SetState(227)
			p.Match(TronParserRCHEVR)
		}
		{
			p.SetState(228)

			var _m = p.Match(TronParserID)

			localctx.(*MapLocalLeftContext).c = _m
		}

	}

	return localctx
}

// IRight_assocContext is an interface to support dynamic dispatch.
type IRight_assocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRight_assocContext differentiates from other interfaces.
	IsRight_assocContext()
}

type Right_assocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRight_assocContext() *Right_assocContext {
	var p = new(Right_assocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_right_assoc
	return p
}

func (*Right_assocContext) IsRight_assocContext() {}

func NewRight_assocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Right_assocContext {
	var p = new(Right_assocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_right_assoc

	return p
}

func (s *Right_assocContext) GetParser() antlr.Parser { return s.parser }

func (s *Right_assocContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Right_assocContext) INT() antlr.TerminalNode {
	return s.GetToken(TronParserINT, 0)
}

func (s *Right_assocContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *Right_assocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Right_assocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Right_assocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRight_assoc(s)
	}
}

func (s *Right_assocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRight_assoc(s)
	}
}

func (p *TronParser) Right_assoc() (localctx IRight_assocContext) {
	localctx = NewRight_assocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TronParserRULE_right_assoc)
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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Constant()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Match(TronParserINT)
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserLSB {
			{
				p.SetState(233)
				p.FieldOptions()
			}

		}

	}

	return localctx
}

// IFullnameContext is an interface to support dynamic dispatch.
type IFullnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullnameContext differentiates from other interfaces.
	IsFullnameContext()
}

type FullnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullnameContext() *FullnameContext {
	var p = new(FullnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_fullname
	return p
}

func (*FullnameContext) IsFullnameContext() {}

func NewFullnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullnameContext {
	var p = new(FullnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_fullname

	return p
}

func (s *FullnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FullnameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *FullnameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *FullnameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *FullnameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *FullnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterFullname(s)
	}
}

func (s *FullnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitFullname(s)
	}
}

func (p *TronParser) Fullname() (localctx IFullnameContext) {
	localctx = NewFullnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TronParserRULE_fullname)
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
		p.SetState(238)
		p.Match(TronParserID)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TronParserDOT {
		{
			p.SetState(239)
			p.Match(TronParserDOT)
		}
		{
			p.SetState(240)
			p.Match(TronParserID)
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStream returns the stream token.
	GetStream() antlr.Token

	// SetStream sets the stream token.
	SetStream(antlr.Token)

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	stream antlr.Token
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) GetStream() antlr.Token { return s.stream }

func (s *MessageTypeContext) SetStream(v antlr.Token) { s.stream = v }

func (s *MessageTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TronParserLPAREN, 0)
}

func (s *MessageTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *MessageTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *MessageTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TronParserRPAREN, 0)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *TronParser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TronParserRULE_messageType)
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
	{
		p.SetState(246)
		p.Match(TronParserLPAREN)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(247)

			var _m = p.Match(TronParserID)

			localctx.(*MessageTypeContext).stream = _m
		}

	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TronParserDOT {
		{
			p.SetState(250)
			p.Match(TronParserDOT)
		}

	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(253)
				p.Match(TronParserID)
			}
			{
				p.SetState(254)
				p.Match(TronParserDOT)
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	{
		p.SetState(260)
		p.Match(TronParserID)
	}
	{
		p.SetState(261)
		p.Match(TronParserRPAREN)
	}

	return localctx
}

// IRpcDelimContext is an interface to support dynamic dispatch.
type IRpcDelimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpt returns the opt token.
	GetOpt() antlr.Token

	// SetOpt sets the opt token.
	SetOpt(antlr.Token)

	// IsRpcDelimContext differentiates from other interfaces.
	IsRpcDelimContext()
}

type RpcDelimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	opt    antlr.Token
}

func NewEmptyRpcDelimContext() *RpcDelimContext {
	var p = new(RpcDelimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronParserRULE_rpcDelim
	return p
}

func (*RpcDelimContext) IsRpcDelimContext() {}

func NewRpcDelimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcDelimContext {
	var p = new(RpcDelimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_rpcDelim

	return p
}

func (s *RpcDelimContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcDelimContext) GetOpt() antlr.Token { return s.opt }

func (s *RpcDelimContext) SetOpt(v antlr.Token) { s.opt = v }

func (s *RpcDelimContext) LCUR() antlr.TerminalNode {
	return s.GetToken(TronParserLCUR, 0)
}

func (s *RpcDelimContext) RCUR() antlr.TerminalNode {
	return s.GetToken(TronParserRCUR, 0)
}

func (s *RpcDelimContext) AllOptionName() []IOptionNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionNameContext)(nil)).Elem())
	var tst = make([]IOptionNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionNameContext)
		}
	}

	return tst
}

func (s *RpcDelimContext) OptionName(i int) IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *RpcDelimContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(TronParserEQ)
}

func (s *RpcDelimContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(TronParserEQ, i)
}

func (s *RpcDelimContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *RpcDelimContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *RpcDelimContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(TronParserSEMI)
}

func (s *RpcDelimContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, i)
}

func (s *RpcDelimContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *RpcDelimContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *RpcDelimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcDelimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcDelimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRpcDelim(s)
	}
}

func (s *RpcDelimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRpcDelim(s)
	}
}

func (p *TronParser) RpcDelim() (localctx IRpcDelimContext) {
	localctx = NewRpcDelimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TronParserRULE_rpcDelim)
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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronParserLCUR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(TronParserLCUR)
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserSEMI || _la == TronParserID {
			p.SetState(271)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case TronParserID:
				{
					p.SetState(264)

					var _m = p.Match(TronParserID)

					localctx.(*RpcDelimContext).opt = _m
				}
				{
					p.SetState(265)
					p.OptionName()
				}
				{
					p.SetState(266)
					p.Match(TronParserEQ)
				}
				{
					p.SetState(267)
					p.Constant()
				}
				{
					p.SetState(268)
					p.Match(TronParserSEMI)
				}

			case TronParserSEMI:
				{
					p.SetState(270)
					p.Match(TronParserSEMI)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
			p.Match(TronParserRCUR)
		}

	case TronParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(TronParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = TronParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_ranges

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
	return s.GetTokens(TronParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TronParserCOMMA, i)
}

func (s *RangesContext) AllSTR() []antlr.TerminalNode {
	return s.GetTokens(TronParserSTR)
}

func (s *RangesContext) STR(i int) antlr.TerminalNode {
	return s.GetToken(TronParserSTR, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *TronParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TronParserRULE_ranges)
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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Rangee()
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserCOMMA {
			{
				p.SetState(281)
				p.Match(TronParserCOMMA)
			}
			{
				p.SetState(282)
				p.Rangee()
			}

			p.SetState(287)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TronParserSTR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Match(TronParserSTR)
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserCOMMA {
			{
				p.SetState(289)
				p.Match(TronParserCOMMA)
			}
			{
				p.SetState(290)
				p.Match(TronParserSTR)
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = TronParserRULE_rangee
	return p
}

func (*RangeeContext) IsRangeeContext() {}

func NewRangeeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeeContext {
	var p = new(RangeeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_rangee

	return p
}

func (s *RangeeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeeContext) GetTo() antlr.Token { return s.to }

func (s *RangeeContext) GetMax() antlr.Token { return s.max }

func (s *RangeeContext) SetTo(v antlr.Token) { s.to = v }

func (s *RangeeContext) SetMax(v antlr.Token) { s.max = v }

func (s *RangeeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(TronParserINT)
}

func (s *RangeeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserINT, i)
}

func (s *RangeeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *RangeeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *RangeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterRangee(s)
	}
}

func (s *RangeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitRangee(s)
	}
}

func (p *TronParser) Rangee() (localctx IRangeeContext) {
	localctx = NewRangeeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TronParserRULE_rangee)

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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(TronParserINT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Match(TronParserINT)
		}
		{
			p.SetState(300)

			var _m = p.Match(TronParserID)

			localctx.(*RangeeContext).to = _m
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TronParserINT:
			{
				p.SetState(301)
				p.Match(TronParserINT)
			}

		case TronParserID:
			{
				p.SetState(302)

				var _m = p.Match(TronParserID)

				localctx.(*RangeeContext).max = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

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
	p.RuleIndex = TronParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) LSB() antlr.TerminalNode {
	return s.GetToken(TronParserLSB, 0)
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
	return s.GetToken(TronParserRSB, 0)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TronParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TronParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (p *TronParser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TronParserRULE_fieldOptions)
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
		p.SetState(307)
		p.Match(TronParserLSB)
	}
	{
		p.SetState(308)
		p.FieldOption()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TronParserCOMMA {
		{
			p.SetState(309)
			p.Match(TronParserCOMMA)
		}
		{
			p.SetState(310)
			p.FieldOption()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(316)
		p.Match(TronParserRSB)
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
	p.RuleIndex = TronParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_fieldOption

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
	return s.GetToken(TronParserEQ, 0)
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
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (p *TronParser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TronParserRULE_fieldOption)

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
		p.SetState(318)
		p.OptionName()
	}
	{
		p.SetState(319)
		p.Match(TronParserEQ)
	}
	{
		p.SetState(320)
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
	p.RuleIndex = TronParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *OptionNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *OptionNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TronParserLPAREN, 0)
}

func (s *OptionNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TronParserRPAREN, 0)
}

func (s *OptionNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *OptionNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *TronParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TronParserRULE_optionName)
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
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronParserID:
		{
			p.SetState(322)
			p.Match(TronParserID)
		}

	case TronParserLPAREN:
		{
			p.SetState(323)
			p.Match(TronParserLPAREN)
		}
		{
			p.SetState(324)
			p.Match(TronParserID)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TronParserDOT {
			{
				p.SetState(325)
				p.Match(TronParserDOT)
			}
			{
				p.SetState(326)
				p.Match(TronParserID)
			}

			p.SetState(331)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(332)
			p.Match(TronParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TronParserDOT {
		{
			p.SetState(335)
			p.Match(TronParserDOT)
		}
		{
			p.SetState(336)
			p.Match(TronParserID)
		}

		p.SetState(341)
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
	p.RuleIndex = TronParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetId() antlr.Token { return s.id }

func (s *ConstantContext) SetId(v antlr.Token) { s.id = v }

func (s *ConstantContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TronParserID)
}

func (s *ConstantContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TronParserID, i)
}

func (s *ConstantContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TronParserDOT)
}

func (s *ConstantContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TronParserDOT, i)
}

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(TronParserINT, 0)
}

func (s *ConstantContext) DASH() antlr.TerminalNode {
	return s.GetToken(TronParserDASH, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TronParserPLUS, 0)
}

func (s *ConstantContext) FLT() antlr.TerminalNode {
	return s.GetToken(TronParserFLT, 0)
}

func (s *ConstantContext) STR() antlr.TerminalNode {
	return s.GetToken(TronParserSTR, 0)
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
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *TronParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TronParserRULE_constant)
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

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)
			p.Match(TronParserID)
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == TronParserDOT {
			{
				p.SetState(343)
				p.Match(TronParserDOT)
			}
			{
				p.SetState(344)
				p.Match(TronParserID)
			}

			p.SetState(347)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)

			var _m = p.Match(TronParserID)

			localctx.(*ConstantContext).id = _m
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserDASH || _la == TronParserPLUS {
			{
				p.SetState(350)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TronParserDASH || _la == TronParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(353)
			p.Match(TronParserINT)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserDASH || _la == TronParserPLUS {
			{
				p.SetState(354)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TronParserDASH || _la == TronParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(357)
			p.Match(TronParserFLT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(358)
			p.Match(TronParserSTR)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(359)
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
	p.RuleIndex = TronParserRULE_constantObj
	return p
}

func (*ConstantObjContext) IsConstantObjContext() {}

func NewConstantObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantObjContext {
	var p = new(ConstantObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_constantObj

	return p
}

func (s *ConstantObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantObjContext) LCUR() antlr.TerminalNode {
	return s.GetToken(TronParserLCUR, 0)
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
	return s.GetToken(TronParserRCUR, 0)
}

func (s *ConstantObjContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TronParserCOMMA)
}

func (s *ConstantObjContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TronParserCOMMA, i)
}

func (s *ConstantObjContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(TronParserSEMI)
}

func (s *ConstantObjContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(TronParserSEMI, i)
}

func (s *ConstantObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterConstantObj(s)
	}
}

func (s *ConstantObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitConstantObj(s)
	}
}

func (p *TronParser) ConstantObj() (localctx IConstantObjContext) {
	localctx = NewConstantObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TronParserRULE_constantObj)
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
		p.SetState(362)
		p.Match(TronParserLCUR)
	}
	{
		p.SetState(363)
		p.ConstantObjElem()
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TronParserSEMI)|(1<<TronParserCOMMA)|(1<<TronParserID))) != 0 {
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronParserSEMI || _la == TronParserCOMMA {
			{
				p.SetState(364)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TronParserSEMI || _la == TronParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(367)
			p.ConstantObjElem()
		}

		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(373)
		p.Match(TronParserRCUR)
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
	p.RuleIndex = TronParserRULE_constantObjElem
	return p
}

func (*ConstantObjElemContext) IsConstantObjElemContext() {}

func NewConstantObjElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantObjElemContext {
	var p = new(ConstantObjElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronParserRULE_constantObjElem

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
	return s.GetToken(TronParserCOLON, 0)
}

func (s *PronSTRContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PronSTRContext) ID() antlr.TerminalNode {
	return s.GetToken(TronParserID, 0)
}

func (s *PronSTRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterPronSTR(s)
	}
}

func (s *PronSTRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
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
	return s.GetToken(TronParserCOLON, 0)
}

func (s *PronOBJContext) ConstantObj() IConstantObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantObjContext)
}

func (s *PronOBJContext) ID() antlr.TerminalNode {
	return s.GetToken(TronParserID, 0)
}

func (s *PronOBJContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.EnterPronOBJ(s)
	}
}

func (s *PronOBJContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronParserListener); ok {
		listenerT.ExitPronOBJ(s)
	}
}

func (p *TronParser) ConstantObjElem() (localctx IConstantObjElemContext) {
	localctx = NewConstantObjElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TronParserRULE_constantObjElem)

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

	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPronSTRContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)

			var _m = p.Match(TronParserID)

			localctx.(*PronSTRContext).id = _m
		}
		{
			p.SetState(376)
			p.Match(TronParserCOLON)
		}
		{
			p.SetState(377)
			p.Constant()
		}

	case 2:
		localctx = NewPronOBJContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)

			var _m = p.Match(TronParserID)

			localctx.(*PronOBJContext).id = _m
		}
		{
			p.SetState(379)
			p.Match(TronParserCOLON)
		}
		{
			p.SetState(380)
			p.ConstantObj()
		}

	}

	return localctx
}
