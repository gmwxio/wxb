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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 466,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 3, 2,
	7, 2, 89, 10, 2, 12, 2, 14, 2, 92, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 109,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 116, 10, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 125, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 134, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 143, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 152,
	10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 161, 10, 10,
	12, 10, 14, 10, 164, 11, 10, 3, 10, 5, 10, 167, 10, 10, 3, 10, 3, 10, 7,
	10, 171, 10, 10, 12, 10, 14, 10, 174, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	7, 11, 180, 10, 11, 12, 11, 14, 11, 183, 11, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 5, 12, 191, 10, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 7, 13, 200, 10, 13, 12, 13, 14, 13, 203, 11, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 211, 10, 14, 12, 14, 14, 14, 214,
	11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 5, 15, 226, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 232, 10,
	16, 13, 16, 14, 16, 233, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 5, 17, 241,
	10, 17, 3, 18, 3, 18, 3, 18, 5, 18, 246, 10, 18, 3, 18, 3, 18, 5, 18, 250,
	10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19,
	260, 10, 19, 12, 19, 14, 19, 263, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 7, 20, 271, 10, 20, 12, 20, 14, 20, 274, 11, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 5, 21, 281, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 7, 22, 291, 10, 22, 12, 22, 14, 22, 294, 11, 22,
	3, 22, 3, 22, 5, 22, 298, 10, 22, 3, 23, 3, 23, 5, 23, 302, 10, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 310, 10, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 7, 25, 317, 10, 25, 12, 25, 14, 25, 320, 11, 25, 3,
	26, 3, 26, 3, 26, 3, 26, 5, 26, 326, 10, 26, 3, 27, 3, 27, 3, 27, 7, 27,
	331, 10, 27, 12, 27, 14, 27, 334, 11, 27, 3, 28, 3, 28, 5, 28, 338, 10,
	28, 3, 29, 3, 29, 3, 30, 5, 30, 343, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 5, 30, 350, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 7,
	31, 358, 10, 31, 12, 31, 14, 31, 361, 11, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 374, 10, 33, 12,
	33, 14, 33, 377, 11, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	5, 34, 386, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 400, 10, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 37, 5, 37, 407, 10, 37, 3, 37, 3, 37, 7, 37, 411, 10, 37, 12,
	37, 14, 37, 414, 11, 37, 3, 37, 3, 37, 3, 38, 5, 38, 419, 10, 38, 3, 38,
	3, 38, 7, 38, 423, 10, 38, 12, 38, 14, 38, 426, 11, 38, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 7, 41, 437, 10, 41, 12, 41,
	14, 41, 440, 11, 41, 3, 41, 5, 41, 443, 10, 41, 3, 41, 3, 41, 5, 41, 447,
	10, 41, 3, 41, 3, 41, 3, 41, 5, 41, 452, 10, 41, 5, 41, 454, 10, 41, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 5, 43, 464, 10, 43,
	3, 43, 2, 2, 44, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 2, 7, 3, 2, 29, 30, 4, 2, 47, 54, 56,
	62, 6, 2, 47, 47, 50, 51, 53, 54, 56, 62, 3, 2, 45, 46, 3, 2, 25, 26, 2,
	487, 2, 86, 3, 2, 2, 2, 4, 95, 3, 2, 2, 2, 6, 108, 3, 2, 2, 2, 8, 110,
	3, 2, 2, 2, 10, 119, 3, 2, 2, 2, 12, 128, 3, 2, 2, 2, 14, 137, 3, 2, 2,
	2, 16, 146, 3, 2, 2, 2, 18, 166, 3, 2, 2, 2, 20, 175, 3, 2, 2, 2, 22, 188,
	3, 2, 2, 2, 24, 195, 3, 2, 2, 2, 26, 206, 3, 2, 2, 2, 28, 225, 3, 2, 2,
	2, 30, 227, 3, 2, 2, 2, 32, 240, 3, 2, 2, 2, 34, 242, 3, 2, 2, 2, 36, 253,
	3, 2, 2, 2, 38, 266, 3, 2, 2, 2, 40, 280, 3, 2, 2, 2, 42, 282, 3, 2, 2,
	2, 44, 299, 3, 2, 2, 2, 46, 306, 3, 2, 2, 2, 48, 313, 3, 2, 2, 2, 50, 325,
	3, 2, 2, 2, 52, 327, 3, 2, 2, 2, 54, 337, 3, 2, 2, 2, 56, 339, 3, 2, 2,
	2, 58, 342, 3, 2, 2, 2, 60, 353, 3, 2, 2, 2, 62, 364, 3, 2, 2, 2, 64, 368,
	3, 2, 2, 2, 66, 380, 3, 2, 2, 2, 68, 389, 3, 2, 2, 2, 70, 403, 3, 2, 2,
	2, 72, 406, 3, 2, 2, 2, 74, 418, 3, 2, 2, 2, 76, 429, 3, 2, 2, 2, 78, 431,
	3, 2, 2, 2, 80, 453, 3, 2, 2, 2, 82, 455, 3, 2, 2, 2, 84, 463, 3, 2, 2,
	2, 86, 90, 5, 4, 3, 2, 87, 89, 5, 6, 4, 2, 88, 87, 3, 2, 2, 2, 89, 92,
	3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 93, 3, 2, 2, 2,
	92, 90, 3, 2, 2, 2, 93, 94, 7, 2, 2, 3, 94, 3, 3, 2, 2, 2, 95, 96, 7, 4,
	2, 2, 96, 97, 7, 5, 2, 2, 97, 98, 7, 3, 2, 2, 98, 99, 7, 6, 2, 2, 99, 5,
	3, 2, 2, 2, 100, 109, 5, 22, 12, 2, 101, 109, 5, 24, 13, 2, 102, 109, 5,
	20, 11, 2, 103, 109, 5, 8, 5, 2, 104, 109, 5, 26, 14, 2, 105, 109, 5, 30,
	16, 2, 106, 109, 5, 38, 20, 2, 107, 109, 5, 78, 40, 2, 108, 100, 3, 2,
	2, 2, 108, 101, 3, 2, 2, 2, 108, 102, 3, 2, 2, 2, 108, 103, 3, 2, 2, 2,
	108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108,
	107, 3, 2, 2, 2, 109, 7, 3, 2, 2, 2, 110, 111, 7, 33, 2, 2, 111, 112, 5,
	18, 10, 2, 112, 115, 7, 12, 2, 2, 113, 116, 5, 80, 41, 2, 114, 116, 5,
	82, 42, 2, 115, 113, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 117, 3, 2,
	2, 2, 117, 118, 7, 15, 2, 2, 118, 9, 3, 2, 2, 2, 119, 120, 7, 33, 2, 2,
	120, 121, 5, 18, 10, 2, 121, 124, 7, 12, 2, 2, 122, 125, 5, 80, 41, 2,
	123, 125, 5, 82, 42, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125,
	126, 3, 2, 2, 2, 126, 127, 7, 15, 2, 2, 127, 11, 3, 2, 2, 2, 128, 129,
	7, 33, 2, 2, 129, 130, 5, 18, 10, 2, 130, 133, 7, 12, 2, 2, 131, 134, 5,
	80, 41, 2, 132, 134, 5, 82, 42, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2,
	2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 7, 15, 2, 2, 136, 13, 3, 2, 2, 2,
	137, 138, 7, 33, 2, 2, 138, 139, 5, 18, 10, 2, 139, 142, 7, 12, 2, 2, 140,
	143, 5, 80, 41, 2, 141, 143, 5, 82, 42, 2, 142, 140, 3, 2, 2, 2, 142, 141,
	3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 15, 2, 2, 145, 15, 3, 2,
	2, 2, 146, 147, 7, 33, 2, 2, 147, 148, 5, 18, 10, 2, 148, 151, 7, 12, 2,
	2, 149, 152, 5, 80, 41, 2, 150, 152, 5, 82, 42, 2, 151, 149, 3, 2, 2, 2,
	151, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 15, 2, 2, 154,
	17, 3, 2, 2, 2, 155, 167, 7, 64, 2, 2, 156, 157, 7, 17, 2, 2, 157, 162,
	7, 64, 2, 2, 158, 159, 7, 19, 2, 2, 159, 161, 7, 64, 2, 2, 160, 158, 3,
	2, 2, 2, 161, 164, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2,
	2, 163, 165, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 165, 167, 7, 18, 2, 2, 166,
	155, 3, 2, 2, 2, 166, 156, 3, 2, 2, 2, 167, 172, 3, 2, 2, 2, 168, 169,
	7, 19, 2, 2, 169, 171, 7, 64, 2, 2, 170, 168, 3, 2, 2, 2, 171, 174, 3,
	2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 19, 3, 2, 2,
	2, 174, 172, 3, 2, 2, 2, 175, 176, 7, 27, 2, 2, 176, 181, 7, 64, 2, 2,
	177, 178, 7, 19, 2, 2, 178, 180, 7, 64, 2, 2, 179, 177, 3, 2, 2, 2, 180,
	183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 184,
	3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 185, 7, 10, 2, 2, 185, 186, 5, 58,
	30, 2, 186, 187, 7, 11, 2, 2, 187, 21, 3, 2, 2, 2, 188, 190, 7, 28, 2,
	2, 189, 191, 9, 2, 2, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	192, 3, 2, 2, 2, 192, 193, 7, 63, 2, 2, 193, 194, 7, 15, 2, 2, 194, 23,
	3, 2, 2, 2, 195, 196, 7, 32, 2, 2, 196, 201, 7, 64, 2, 2, 197, 198, 7,
	19, 2, 2, 198, 200, 7, 64, 2, 2, 199, 197, 3, 2, 2, 2, 200, 203, 3, 2,
	2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204, 3, 2, 2, 2,
	203, 201, 3, 2, 2, 2, 204, 205, 7, 15, 2, 2, 205, 25, 3, 2, 2, 2, 206,
	207, 7, 55, 2, 2, 207, 208, 7, 64, 2, 2, 208, 212, 7, 10, 2, 2, 209, 211,
	5, 28, 15, 2, 210, 209, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 210, 3,
	2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215, 3, 2, 2, 2, 214, 212, 3, 2, 2,
	2, 215, 216, 7, 11, 2, 2, 216, 27, 3, 2, 2, 2, 217, 226, 5, 58, 30, 2,
	218, 226, 5, 30, 16, 2, 219, 226, 5, 26, 14, 2, 220, 226, 5, 10, 6, 2,
	221, 226, 5, 64, 33, 2, 222, 226, 5, 68, 35, 2, 223, 226, 5, 46, 24, 2,
	224, 226, 5, 78, 40, 2, 225, 217, 3, 2, 2, 2, 225, 218, 3, 2, 2, 2, 225,
	219, 3, 2, 2, 2, 225, 220, 3, 2, 2, 2, 225, 221, 3, 2, 2, 2, 225, 222,
	3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 29, 3, 2,
	2, 2, 227, 228, 7, 34, 2, 2, 228, 229, 7, 64, 2, 2, 229, 231, 7, 10, 2,
	2, 230, 232, 5, 32, 17, 2, 231, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2,
	233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235,
	236, 7, 11, 2, 2, 236, 31, 3, 2, 2, 2, 237, 241, 5, 12, 7, 2, 238, 241,
	5, 34, 18, 2, 239, 241, 5, 78, 40, 2, 240, 237, 3, 2, 2, 2, 240, 238, 3,
	2, 2, 2, 240, 239, 3, 2, 2, 2, 241, 33, 3, 2, 2, 2, 242, 243, 7, 64, 2,
	2, 243, 245, 7, 12, 2, 2, 244, 246, 7, 25, 2, 2, 245, 244, 3, 2, 2, 2,
	245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 249, 7, 65, 2, 2, 248,
	250, 5, 36, 19, 2, 249, 248, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 251,
	3, 2, 2, 2, 251, 252, 7, 15, 2, 2, 252, 35, 3, 2, 2, 2, 253, 254, 7, 20,
	2, 2, 254, 255, 5, 18, 10, 2, 255, 256, 7, 12, 2, 2, 256, 261, 5, 80, 41,
	2, 257, 258, 7, 22, 2, 2, 258, 260, 5, 36, 19, 2, 259, 257, 3, 2, 2, 2,
	260, 263, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262,
	264, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 264, 265, 7, 21, 2, 2, 265, 37,
	3, 2, 2, 2, 266, 267, 7, 35, 2, 2, 267, 268, 7, 64, 2, 2, 268, 272, 7,
	10, 2, 2, 269, 271, 5, 40, 21, 2, 270, 269, 3, 2, 2, 2, 271, 274, 3, 2,
	2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 275, 3, 2, 2, 2,
	274, 272, 3, 2, 2, 2, 275, 276, 7, 11, 2, 2, 276, 39, 3, 2, 2, 2, 277,
	281, 5, 14, 8, 2, 278, 281, 5, 42, 22, 2, 279, 281, 5, 78, 40, 2, 280,
	277, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 279, 3, 2, 2, 2, 281, 41, 3,
	2, 2, 2, 282, 283, 7, 36, 2, 2, 283, 284, 7, 64, 2, 2, 284, 285, 5, 44,
	23, 2, 285, 286, 7, 31, 2, 2, 286, 297, 5, 44, 23, 2, 287, 292, 7, 10,
	2, 2, 288, 291, 5, 16, 9, 2, 289, 291, 5, 78, 40, 2, 290, 288, 3, 2, 2,
	2, 290, 289, 3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292,
	293, 3, 2, 2, 2, 293, 295, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 298,
	7, 11, 2, 2, 296, 298, 7, 15, 2, 2, 297, 287, 3, 2, 2, 2, 297, 296, 3,
	2, 2, 2, 298, 43, 3, 2, 2, 2, 299, 301, 7, 17, 2, 2, 300, 302, 7, 37, 2,
	2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303,
	304, 5, 72, 37, 2, 304, 305, 7, 18, 2, 2, 305, 45, 3, 2, 2, 2, 306, 309,
	7, 38, 2, 2, 307, 310, 5, 48, 25, 2, 308, 310, 5, 52, 27, 2, 309, 307,
	3, 2, 2, 2, 309, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312, 7, 15,
	2, 2, 312, 47, 3, 2, 2, 2, 313, 318, 5, 50, 26, 2, 314, 315, 7, 22, 2,
	2, 315, 317, 5, 50, 26, 2, 316, 314, 3, 2, 2, 2, 317, 320, 3, 2, 2, 2,
	318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 49, 3, 2, 2, 2, 320, 318,
	3, 2, 2, 2, 321, 326, 7, 65, 2, 2, 322, 323, 7, 65, 2, 2, 323, 324, 7,
	39, 2, 2, 324, 326, 7, 65, 2, 2, 325, 321, 3, 2, 2, 2, 325, 322, 3, 2,
	2, 2, 326, 51, 3, 2, 2, 2, 327, 332, 7, 63, 2, 2, 328, 329, 7, 22, 2, 2,
	329, 331, 7, 63, 2, 2, 330, 328, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332,
	330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 53, 3, 2, 2, 2, 334, 332, 3,
	2, 2, 2, 335, 338, 5, 56, 29, 2, 336, 338, 5, 74, 38, 2, 337, 335, 3, 2,
	2, 2, 337, 336, 3, 2, 2, 2, 338, 55, 3, 2, 2, 2, 339, 340, 9, 3, 2, 2,
	340, 57, 3, 2, 2, 2, 341, 343, 7, 40, 2, 2, 342, 341, 3, 2, 2, 2, 342,
	343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 5, 54, 28, 2, 345, 346,
	7, 64, 2, 2, 346, 347, 7, 12, 2, 2, 347, 349, 7, 65, 2, 2, 348, 350, 5,
	60, 31, 2, 349, 348, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 3, 2,
	2, 2, 351, 352, 7, 15, 2, 2, 352, 59, 3, 2, 2, 2, 353, 354, 7, 20, 2, 2,
	354, 359, 5, 62, 32, 2, 355, 356, 7, 22, 2, 2, 356, 358, 5, 62, 32, 2,
	357, 355, 3, 2, 2, 2, 358, 361, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 359,
	360, 3, 2, 2, 2, 360, 362, 3, 2, 2, 2, 361, 359, 3, 2, 2, 2, 362, 363,
	7, 21, 2, 2, 363, 61, 3, 2, 2, 2, 364, 365, 5, 18, 10, 2, 365, 366, 7,
	12, 2, 2, 366, 367, 5, 80, 41, 2, 367, 63, 3, 2, 2, 2, 368, 369, 7, 41,
	2, 2, 369, 370, 7, 64, 2, 2, 370, 375, 7, 10, 2, 2, 371, 374, 5, 66, 34,
	2, 372, 374, 5, 78, 40, 2, 373, 371, 3, 2, 2, 2, 373, 372, 3, 2, 2, 2,
	374, 377, 3, 2, 2, 2, 375, 373, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376,
	378, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 378, 379, 7, 11, 2, 2, 379, 65,
	3, 2, 2, 2, 380, 381, 5, 54, 28, 2, 381, 382, 7, 64, 2, 2, 382, 383, 7,
	12, 2, 2, 383, 385, 7, 65, 2, 2, 384, 386, 5, 60, 31, 2, 385, 384, 3, 2,
	2, 2, 385, 386, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 7, 15, 2, 2,
	388, 67, 3, 2, 2, 2, 389, 390, 7, 42, 2, 2, 390, 391, 7, 23, 2, 2, 391,
	392, 5, 70, 36, 2, 392, 393, 7, 22, 2, 2, 393, 394, 5, 54, 28, 2, 394,
	395, 7, 24, 2, 2, 395, 396, 7, 64, 2, 2, 396, 397, 7, 12, 2, 2, 397, 399,
	7, 65, 2, 2, 398, 400, 5, 60, 31, 2, 399, 398, 3, 2, 2, 2, 399, 400, 3,
	2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 402, 7, 15, 2, 2, 402, 69, 3, 2, 2,
	2, 403, 404, 9, 4, 2, 2, 404, 71, 3, 2, 2, 2, 405, 407, 7, 19, 2, 2, 406,
	405, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 412, 3, 2, 2, 2, 408, 409,
	7, 64, 2, 2, 409, 411, 7, 19, 2, 2, 410, 408, 3, 2, 2, 2, 411, 414, 3,
	2, 2, 2, 412, 410, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 415, 3, 2, 2,
	2, 414, 412, 3, 2, 2, 2, 415, 416, 7, 64, 2, 2, 416, 73, 3, 2, 2, 2, 417,
	419, 7, 19, 2, 2, 418, 417, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 424,
	3, 2, 2, 2, 420, 421, 7, 64, 2, 2, 421, 423, 7, 19, 2, 2, 422, 420, 3,
	2, 2, 2, 423, 426, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2,
	2, 425, 427, 3, 2, 2, 2, 426, 424, 3, 2, 2, 2, 427, 428, 7, 64, 2, 2, 428,
	75, 3, 2, 2, 2, 429, 430, 9, 5, 2, 2, 430, 77, 3, 2, 2, 2, 431, 432, 7,
	15, 2, 2, 432, 79, 3, 2, 2, 2, 433, 438, 7, 64, 2, 2, 434, 435, 7, 19,
	2, 2, 435, 437, 7, 64, 2, 2, 436, 434, 3, 2, 2, 2, 437, 440, 3, 2, 2, 2,
	438, 436, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 454, 3, 2, 2, 2, 440,
	438, 3, 2, 2, 2, 441, 443, 9, 6, 2, 2, 442, 441, 3, 2, 2, 2, 442, 443,
	3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 454, 7, 65, 2, 2, 445, 447, 9, 6,
	2, 2, 446, 445, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2,
	448, 454, 7, 66, 2, 2, 449, 452, 7, 63, 2, 2, 450, 452, 5, 76, 39, 2, 451,
	449, 3, 2, 2, 2, 451, 450, 3, 2, 2, 2, 452, 454, 3, 2, 2, 2, 453, 433,
	3, 2, 2, 2, 453, 442, 3, 2, 2, 2, 453, 446, 3, 2, 2, 2, 453, 451, 3, 2,
	2, 2, 454, 81, 3, 2, 2, 2, 455, 456, 7, 10, 2, 2, 456, 457, 7, 64, 2, 2,
	457, 458, 7, 16, 2, 2, 458, 459, 5, 84, 43, 2, 459, 460, 7, 11, 2, 2, 460,
	83, 3, 2, 2, 2, 461, 464, 7, 63, 2, 2, 462, 464, 5, 82, 42, 2, 463, 461,
	3, 2, 2, 2, 463, 462, 3, 2, 2, 2, 464, 85, 3, 2, 2, 2, 50, 90, 108, 115,
	124, 133, 142, 151, 162, 166, 172, 181, 190, 201, 212, 225, 233, 240, 245,
	249, 261, 272, 280, 290, 292, 297, 301, 309, 318, 325, 332, 337, 342, 349,
	359, 373, 375, 385, 399, 406, 412, 418, 424, 438, 442, 446, 451, 453, 463,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'syntax'", "", "", "", "", "", "'{'", "'}'", "", "'\"'", "'''",
	"", "':'", "'('", "')'", "'.'", "'['", "']'", "','", "'<'", "'>'", "'-'",
	"'+'", "'extend'", "'import'", "'weak'", "'public'", "'returns'", "'package'",
	"'option'", "'enum'", "'service'", "'rpc'", "'stream'", "'reserved'", "'to'",
	"'repeated'", "'oneof'", "'map'", "'inf'", "'nan'", "'true'", "'false'",
	"'bool'", "'bytes'", "'double'", "'fixed32'", "'fixed64'", "'float'", "'int32'",
	"'int64'", "'message'", "'sfixed32'", "'sfixed64'", "'sint32'", "'sint64'",
	"'string'", "'uint32'", "'uint64'",
}
var symbolicNames = []string{
	"", "PROTO3", "SYNTAX", "EQ_PRE", "ENDPRE", "PRE_WS", "PRE_COMMENT", "PRE_LINE_COMMENT",
	"LCUR", "RCUR", "EQ", "DQ", "SQ", "SEMI", "COLON", "LPAREN", "RPAREN",
	"DOT", "LSB", "RSB", "COMMA", "LCHEVR", "RCHEVR", "DASH", "PLUS", "EXTEND",
	"IMPORT", "WEAK", "PUBLIC", "RETURNS", "PACKAGE", "OPTION", "ENUM", "SERVICE",
	"RPC", "STREAM", "RESERVED", "TO", "REPEATED", "ONEOF", "MAP", "INF", "NAN",
	"TRUE", "FALSE", "BOOL", "BYTES", "DOUBLE", "FIXED32", "FIXED64", "FLOAT",
	"INT32", "INT64", "MESSAGE", "SFIXED32", "SFIXED64", "SINT32", "SINT64",
	"STRING", "UINT32", "UINT64", "StrLit", "Ident", "Int_lit", "Float_lit",
	"WS", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"proto", "syntax", "top_level_statement", "optionFile", "optionMessage",
	"optionEnum", "optionService", "optionRpc", "optionName", "extend", "importStatement",
	"packageStatement", "message", "messageBody", "enumDefinition", "enumBody",
	"enumField", "enumValueOption", "service", "serviceBody", "rpc", "rpcParam",
	"reserved", "ranges", "rangee", "fieldNames", "typer", "types", "field",
	"fieldOptions", "fieldOption", "oneof", "oneofField", "mapField", "keyType",
	"messageType", "messageOrEnumType", "bool_lit", "emptyStatement", "constant",
	"pron", "pronElem",
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
	tronParserEXTEND           = 25
	tronParserIMPORT           = 26
	tronParserWEAK             = 27
	tronParserPUBLIC           = 28
	tronParserRETURNS          = 29
	tronParserPACKAGE          = 30
	tronParserOPTION           = 31
	tronParserENUM             = 32
	tronParserSERVICE          = 33
	tronParserRPC              = 34
	tronParserSTREAM           = 35
	tronParserRESERVED         = 36
	tronParserTO               = 37
	tronParserREPEATED         = 38
	tronParserONEOF            = 39
	tronParserMAP              = 40
	tronParserINF              = 41
	tronParserNAN              = 42
	tronParserTRUE             = 43
	tronParserFALSE            = 44
	tronParserBOOL             = 45
	tronParserBYTES            = 46
	tronParserDOUBLE           = 47
	tronParserFIXED32          = 48
	tronParserFIXED64          = 49
	tronParserFLOAT            = 50
	tronParserINT32            = 51
	tronParserINT64            = 52
	tronParserMESSAGE          = 53
	tronParserSFIXED32         = 54
	tronParserSFIXED64         = 55
	tronParserSINT32           = 56
	tronParserSINT64           = 57
	tronParserSTRING           = 58
	tronParserUINT32           = 59
	tronParserUINT64           = 60
	tronParserStrLit           = 61
	tronParserIdent            = 62
	tronParserInt_lit          = 63
	tronParserFloat_lit        = 64
	tronParserWS               = 65
	tronParserCOMMENT          = 66
	tronParserLINE_COMMENT     = 67
)

// tronParser rules.
const (
	tronParserRULE_proto               = 0
	tronParserRULE_syntax              = 1
	tronParserRULE_top_level_statement = 2
	tronParserRULE_optionFile          = 3
	tronParserRULE_optionMessage       = 4
	tronParserRULE_optionEnum          = 5
	tronParserRULE_optionService       = 6
	tronParserRULE_optionRpc           = 7
	tronParserRULE_optionName          = 8
	tronParserRULE_extend              = 9
	tronParserRULE_importStatement     = 10
	tronParserRULE_packageStatement    = 11
	tronParserRULE_message             = 12
	tronParserRULE_messageBody         = 13
	tronParserRULE_enumDefinition      = 14
	tronParserRULE_enumBody            = 15
	tronParserRULE_enumField           = 16
	tronParserRULE_enumValueOption     = 17
	tronParserRULE_service             = 18
	tronParserRULE_serviceBody         = 19
	tronParserRULE_rpc                 = 20
	tronParserRULE_rpcParam            = 21
	tronParserRULE_reserved            = 22
	tronParserRULE_ranges              = 23
	tronParserRULE_rangee              = 24
	tronParserRULE_fieldNames          = 25
	tronParserRULE_typer               = 26
	tronParserRULE_types               = 27
	tronParserRULE_field               = 28
	tronParserRULE_fieldOptions        = 29
	tronParserRULE_fieldOption         = 30
	tronParserRULE_oneof               = 31
	tronParserRULE_oneofField          = 32
	tronParserRULE_mapField            = 33
	tronParserRULE_keyType             = 34
	tronParserRULE_messageType         = 35
	tronParserRULE_messageOrEnumType   = 36
	tronParserRULE_bool_lit            = 37
	tronParserRULE_emptyStatement      = 38
	tronParserRULE_constant            = 39
	tronParserRULE_pron                = 40
	tronParserRULE_pronElem            = 41
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

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(tronParserEOF, 0)
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
	{
		p.SetState(84)
		p.Syntax()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tronParserSEMI)|(1<<tronParserEXTEND)|(1<<tronParserIMPORT)|(1<<tronParserPACKAGE)|(1<<tronParserOPTION))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(tronParserENUM-32))|(1<<(tronParserSERVICE-32))|(1<<(tronParserMESSAGE-32)))) != 0) {
		{
			p.SetState(85)
			p.Top_level_statement()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
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
		p.SetState(93)
		p.Match(tronParserSYNTAX)
	}
	{
		p.SetState(94)
		p.Match(tronParserEQ_PRE)
	}
	{
		p.SetState(95)
		p.Match(tronParserPROTO3)
	}
	{
		p.SetState(96)
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

func (s *Top_level_statementContext) OptionFile() IOptionFileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionFileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionFileContext)
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
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterTop_level_statement(s)
	}
}

func (s *Top_level_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitTop_level_statement(s)
	}
}

func (p *tronParser) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tronParserRULE_top_level_statement)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.ImportStatement()
		}

	case tronParserPACKAGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.PackageStatement()
		}

	case tronParserEXTEND:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Extend()
		}

	case tronParserOPTION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.OptionFile()
		}

	case tronParserMESSAGE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Message()
		}

	case tronParserENUM:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(103)
			p.EnumDefinition()
		}

	case tronParserSERVICE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(104)
			p.Service()
		}

	case tronParserSEMI:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(105)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOptionFileContext is an interface to support dynamic dispatch.
type IOptionFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionFileContext differentiates from other interfaces.
	IsOptionFileContext()
}

type OptionFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionFileContext() *OptionFileContext {
	var p = new(OptionFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_optionFile
	return p
}

func (*OptionFileContext) IsOptionFileContext() {}

func NewOptionFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionFileContext {
	var p = new(OptionFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_optionFile

	return p
}

func (s *OptionFileContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionFileContext) CopyFrom(ctx *OptionFileContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OptionFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Option_FileContext struct {
	*OptionFileContext
}

func NewOption_FileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Option_FileContext {
	var p = new(Option_FileContext)

	p.OptionFileContext = NewEmptyOptionFileContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OptionFileContext))

	return p
}

func (s *Option_FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_FileContext) OPTION() antlr.TerminalNode {
	return s.GetToken(tronParserOPTION, 0)
}

func (s *Option_FileContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *Option_FileContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *Option_FileContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *Option_FileContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Option_FileContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *Option_FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOption_File(s)
	}
}

func (s *Option_FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOption_File(s)
	}
}

func (p *tronParser) OptionFile() (localctx IOptionFileContext) {
	localctx = NewOptionFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tronParserRULE_optionFile)

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

	localctx = NewOption_FileContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(tronParserOPTION)
	}
	{
		p.SetState(109)
		p.OptionName()
	}
	{
		p.SetState(110)
		p.Match(tronParserEQ)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserDASH, tronParserPLUS, tronParserTRUE, tronParserFALSE, tronParserStrLit, tronParserIdent, tronParserInt_lit, tronParserFloat_lit:
		{
			p.SetState(111)
			p.Constant()
		}

	case tronParserLCUR:
		{
			p.SetState(112)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(115)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IOptionMessageContext is an interface to support dynamic dispatch.
type IOptionMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionMessageContext differentiates from other interfaces.
	IsOptionMessageContext()
}

type OptionMessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionMessageContext() *OptionMessageContext {
	var p = new(OptionMessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_optionMessage
	return p
}

func (*OptionMessageContext) IsOptionMessageContext() {}

func NewOptionMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionMessageContext {
	var p = new(OptionMessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_optionMessage

	return p
}

func (s *OptionMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionMessageContext) CopyFrom(ctx *OptionMessageContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OptionMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Option_MsgContext struct {
	*OptionMessageContext
}

func NewOption_MsgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Option_MsgContext {
	var p = new(Option_MsgContext)

	p.OptionMessageContext = NewEmptyOptionMessageContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OptionMessageContext))

	return p
}

func (s *Option_MsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_MsgContext) OPTION() antlr.TerminalNode {
	return s.GetToken(tronParserOPTION, 0)
}

func (s *Option_MsgContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *Option_MsgContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *Option_MsgContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *Option_MsgContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Option_MsgContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *Option_MsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOption_Msg(s)
	}
}

func (s *Option_MsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOption_Msg(s)
	}
}

func (p *tronParser) OptionMessage() (localctx IOptionMessageContext) {
	localctx = NewOptionMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tronParserRULE_optionMessage)

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

	localctx = NewOption_MsgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(tronParserOPTION)
	}
	{
		p.SetState(118)
		p.OptionName()
	}
	{
		p.SetState(119)
		p.Match(tronParserEQ)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserDASH, tronParserPLUS, tronParserTRUE, tronParserFALSE, tronParserStrLit, tronParserIdent, tronParserInt_lit, tronParserFloat_lit:
		{
			p.SetState(120)
			p.Constant()
		}

	case tronParserLCUR:
		{
			p.SetState(121)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(124)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IOptionEnumContext is an interface to support dynamic dispatch.
type IOptionEnumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionEnumContext differentiates from other interfaces.
	IsOptionEnumContext()
}

type OptionEnumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionEnumContext() *OptionEnumContext {
	var p = new(OptionEnumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_optionEnum
	return p
}

func (*OptionEnumContext) IsOptionEnumContext() {}

func NewOptionEnumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionEnumContext {
	var p = new(OptionEnumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_optionEnum

	return p
}

func (s *OptionEnumContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionEnumContext) CopyFrom(ctx *OptionEnumContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OptionEnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionEnumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Option_EnumContext struct {
	*OptionEnumContext
}

func NewOption_EnumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Option_EnumContext {
	var p = new(Option_EnumContext)

	p.OptionEnumContext = NewEmptyOptionEnumContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OptionEnumContext))

	return p
}

func (s *Option_EnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_EnumContext) OPTION() antlr.TerminalNode {
	return s.GetToken(tronParserOPTION, 0)
}

func (s *Option_EnumContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *Option_EnumContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *Option_EnumContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *Option_EnumContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Option_EnumContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *Option_EnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOption_Enum(s)
	}
}

func (s *Option_EnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOption_Enum(s)
	}
}

func (p *tronParser) OptionEnum() (localctx IOptionEnumContext) {
	localctx = NewOptionEnumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tronParserRULE_optionEnum)

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

	localctx = NewOption_EnumContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(tronParserOPTION)
	}
	{
		p.SetState(127)
		p.OptionName()
	}
	{
		p.SetState(128)
		p.Match(tronParserEQ)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserDASH, tronParserPLUS, tronParserTRUE, tronParserFALSE, tronParserStrLit, tronParserIdent, tronParserInt_lit, tronParserFloat_lit:
		{
			p.SetState(129)
			p.Constant()
		}

	case tronParserLCUR:
		{
			p.SetState(130)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(133)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IOptionServiceContext is an interface to support dynamic dispatch.
type IOptionServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionServiceContext differentiates from other interfaces.
	IsOptionServiceContext()
}

type OptionServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionServiceContext() *OptionServiceContext {
	var p = new(OptionServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_optionService
	return p
}

func (*OptionServiceContext) IsOptionServiceContext() {}

func NewOptionServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionServiceContext {
	var p = new(OptionServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_optionService

	return p
}

func (s *OptionServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionServiceContext) CopyFrom(ctx *OptionServiceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OptionServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Option_ServiceContext struct {
	*OptionServiceContext
}

func NewOption_ServiceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Option_ServiceContext {
	var p = new(Option_ServiceContext)

	p.OptionServiceContext = NewEmptyOptionServiceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OptionServiceContext))

	return p
}

func (s *Option_ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_ServiceContext) OPTION() antlr.TerminalNode {
	return s.GetToken(tronParserOPTION, 0)
}

func (s *Option_ServiceContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *Option_ServiceContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *Option_ServiceContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *Option_ServiceContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Option_ServiceContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *Option_ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOption_Service(s)
	}
}

func (s *Option_ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOption_Service(s)
	}
}

func (p *tronParser) OptionService() (localctx IOptionServiceContext) {
	localctx = NewOptionServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tronParserRULE_optionService)

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

	localctx = NewOption_ServiceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(tronParserOPTION)
	}
	{
		p.SetState(136)
		p.OptionName()
	}
	{
		p.SetState(137)
		p.Match(tronParserEQ)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserDASH, tronParserPLUS, tronParserTRUE, tronParserFALSE, tronParserStrLit, tronParserIdent, tronParserInt_lit, tronParserFloat_lit:
		{
			p.SetState(138)
			p.Constant()
		}

	case tronParserLCUR:
		{
			p.SetState(139)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(142)
		p.Match(tronParserSEMI)
	}

	return localctx
}

// IOptionRpcContext is an interface to support dynamic dispatch.
type IOptionRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionRpcContext differentiates from other interfaces.
	IsOptionRpcContext()
}

type OptionRpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionRpcContext() *OptionRpcContext {
	var p = new(OptionRpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tronParserRULE_optionRpc
	return p
}

func (*OptionRpcContext) IsOptionRpcContext() {}

func NewOptionRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionRpcContext {
	var p = new(OptionRpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_optionRpc

	return p
}

func (s *OptionRpcContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionRpcContext) CopyFrom(ctx *OptionRpcContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OptionRpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionRpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Option_RpcContext struct {
	*OptionRpcContext
}

func NewOption_RpcContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Option_RpcContext {
	var p = new(Option_RpcContext)

	p.OptionRpcContext = NewEmptyOptionRpcContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OptionRpcContext))

	return p
}

func (s *Option_RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_RpcContext) OPTION() antlr.TerminalNode {
	return s.GetToken(tronParserOPTION, 0)
}

func (s *Option_RpcContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *Option_RpcContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *Option_RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *Option_RpcContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Option_RpcContext) Pron() IPronContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronContext)
}

func (s *Option_RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterOption_Rpc(s)
	}
}

func (s *Option_RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitOption_Rpc(s)
	}
}

func (p *tronParser) OptionRpc() (localctx IOptionRpcContext) {
	localctx = NewOptionRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tronParserRULE_optionRpc)

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

	localctx = NewOption_RpcContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(tronParserOPTION)
	}
	{
		p.SetState(145)
		p.OptionName()
	}
	{
		p.SetState(146)
		p.Match(tronParserEQ)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserDASH, tronParserPLUS, tronParserTRUE, tronParserFALSE, tronParserStrLit, tronParserIdent, tronParserInt_lit, tronParserFloat_lit:
		{
			p.SetState(147)
			p.Constant()
		}

	case tronParserLCUR:
		{
			p.SetState(148)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(151)
		p.Match(tronParserSEMI)
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

func (s *OptionNameContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tronParserIdent)
}

func (s *OptionNameContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tronParserIdent, i)
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
	p.EnterRule(localctx, 16, tronParserRULE_optionName)
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
	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserIdent:
		{
			p.SetState(153)
			p.Match(tronParserIdent)
		}

	case tronParserLPAREN:
		{
			p.SetState(154)
			p.Match(tronParserLPAREN)
		}
		{
			p.SetState(155)
			p.Match(tronParserIdent)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserDOT {
			{
				p.SetState(156)
				p.Match(tronParserDOT)
			}
			{
				p.SetState(157)
				p.Match(tronParserIdent)
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(163)
			p.Match(tronParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserDOT {
		{
			p.SetState(166)
			p.Match(tronParserDOT)
		}
		{
			p.SetState(167)
			p.Match(tronParserIdent)
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = tronParserRULE_extend
	return p
}

func (*ExtendContext) IsExtendContext() {}

func NewExtendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendContext {
	var p = new(ExtendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_extend

	return p
}

func (s *ExtendContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(tronParserEXTEND, 0)
}

func (s *ExtendContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tronParserIdent)
}

func (s *ExtendContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tronParserIdent, i)
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

func (s *ExtendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *tronParser) Extend() (localctx IExtendContext) {
	localctx = NewExtendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tronParserRULE_extend)
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
		p.SetState(173)
		p.Match(tronParserEXTEND)
	}
	{
		p.SetState(174)
		p.Match(tronParserIdent)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserDOT {
		{
			p.SetState(175)
			p.Match(tronParserDOT)
		}
		{
			p.SetState(176)
			p.Match(tronParserIdent)
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(tronParserLCUR)
	}
	{
		p.SetState(183)
		p.Field()
	}
	{
		p.SetState(184)
		p.Match(tronParserRCUR)
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
	p.RuleIndex = tronParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(tronParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() antlr.TerminalNode {
	return s.GetToken(tronParserStrLit, 0)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(tronParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(tronParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *tronParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tronParserRULE_importStatement)
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
		p.SetState(186)
		p.Match(tronParserIMPORT)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserWEAK || _la == tronParserPUBLIC {
		{
			p.SetState(187)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tronParserWEAK || _la == tronParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(190)
		p.Match(tronParserStrLit)
	}
	{
		p.SetState(191)
		p.Match(tronParserSEMI)
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
	p.RuleIndex = tronParserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(tronParserPACKAGE, 0)
}

func (s *PackageStatementContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tronParserIdent)
}

func (s *PackageStatementContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tronParserIdent, i)
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

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *tronParser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tronParserRULE_packageStatement)
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
		p.SetState(193)
		p.Match(tronParserPACKAGE)
	}
	{
		p.SetState(194)
		p.Match(tronParserIdent)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserDOT {
		{
			p.SetState(195)
			p.Match(tronParserDOT)
		}
		{
			p.SetState(196)
			p.Match(tronParserIdent)
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(tronParserSEMI)
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
	p.RuleIndex = tronParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(tronParserMESSAGE, 0)
}

func (s *MessageContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
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

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *tronParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tronParserRULE_message)
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
		p.Match(tronParserMESSAGE)
	}
	{
		p.SetState(205)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(206)
		p.Match(tronParserLCUR)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tronParserSEMI)|(1<<tronParserDOT)|(1<<tronParserOPTION))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(tronParserENUM-32))|(1<<(tronParserRESERVED-32))|(1<<(tronParserREPEATED-32))|(1<<(tronParserONEOF-32))|(1<<(tronParserMAP-32))|(1<<(tronParserBOOL-32))|(1<<(tronParserBYTES-32))|(1<<(tronParserDOUBLE-32))|(1<<(tronParserFIXED32-32))|(1<<(tronParserFIXED64-32))|(1<<(tronParserFLOAT-32))|(1<<(tronParserINT32-32))|(1<<(tronParserINT64-32))|(1<<(tronParserMESSAGE-32))|(1<<(tronParserSFIXED32-32))|(1<<(tronParserSFIXED64-32))|(1<<(tronParserSINT32-32))|(1<<(tronParserSINT64-32))|(1<<(tronParserSTRING-32))|(1<<(tronParserUINT32-32))|(1<<(tronParserUINT64-32))|(1<<(tronParserIdent-32)))) != 0) {
		{
			p.SetState(207)
			p.MessageBody()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
		p.Match(tronParserRCUR)
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

func (s *MessageBodyContext) OptionMessage() IOptionMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionMessageContext)
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
	p.EnterRule(localctx, 26, tronParserRULE_messageBody)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserDOT, tronParserREPEATED, tronParserBOOL, tronParserBYTES, tronParserDOUBLE, tronParserFIXED32, tronParserFIXED64, tronParserFLOAT, tronParserINT32, tronParserINT64, tronParserSFIXED32, tronParserSFIXED64, tronParserSINT32, tronParserSINT64, tronParserSTRING, tronParserUINT32, tronParserUINT64, tronParserIdent:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Field()
		}

	case tronParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.EnumDefinition()
		}

	case tronParserMESSAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(217)
			p.Message()
		}

	case tronParserOPTION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(218)
			p.OptionMessage()
		}

	case tronParserONEOF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(219)
			p.Oneof()
		}

	case tronParserMAP:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(220)
			p.MapField()
		}

	case tronParserRESERVED:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(221)
			p.Reserved()
		}

	case tronParserSEMI:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(222)
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
	p.RuleIndex = tronParserRULE_enumDefinition
	return p
}

func (*EnumDefinitionContext) IsEnumDefinitionContext() {}

func NewEnumDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_enumDefinition

	return p
}

func (s *EnumDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefinitionContext) ENUM() antlr.TerminalNode {
	return s.GetToken(tronParserENUM, 0)
}

func (s *EnumDefinitionContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
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

func (s *EnumDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *tronParser) EnumDefinition() (localctx IEnumDefinitionContext) {
	localctx = NewEnumDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tronParserRULE_enumDefinition)
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
		p.SetState(225)
		p.Match(tronParserENUM)
	}
	{
		p.SetState(226)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(227)
		p.Match(tronParserLCUR)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tronParserSEMI || _la == tronParserOPTION || _la == tronParserIdent {
		{
			p.SetState(228)
			p.EnumBody()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(tronParserRCUR)
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

func (s *EnumBodyContext) OptionEnum() IOptionEnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionEnumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionEnumContext)
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
	p.EnterRule(localctx, 30, tronParserRULE_enumBody)

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

	p.SetState(238)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.OptionEnum()
		}

	case tronParserIdent:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.EnumField()
		}

	case tronParserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
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

func (s *EnumFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *EnumFieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tronParserInt_lit, 0)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *EnumFieldContext) DASH() antlr.TerminalNode {
	return s.GetToken(tronParserDASH, 0)
}

func (s *EnumFieldContext) EnumValueOption() IEnumValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
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
	p.EnterRule(localctx, 32, tronParserRULE_enumField)
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
		p.SetState(240)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(241)
		p.Match(tronParserEQ)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDASH {
		{
			p.SetState(242)
			p.Match(tronParserDASH)
		}

	}
	{
		p.SetState(245)
		p.Match(tronParserInt_lit)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(246)
			p.EnumValueOption()
		}

	}
	{
		p.SetState(249)
		p.Match(tronParserSEMI)
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
	p.RuleIndex = tronParserRULE_enumValueOption
	return p
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) LSB() antlr.TerminalNode {
	return s.GetToken(tronParserLSB, 0)
}

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) RSB() antlr.TerminalNode {
	return s.GetToken(tronParserRSB, 0)
}

func (s *EnumValueOptionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tronParserCOMMA)
}

func (s *EnumValueOptionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tronParserCOMMA, i)
}

func (s *EnumValueOptionContext) AllEnumValueOption() []IEnumValueOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem())
	var tst = make([]IEnumValueOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueOptionContext)
		}
	}

	return tst
}

func (s *EnumValueOptionContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (p *tronParser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tronParserRULE_enumValueOption)
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
		p.SetState(251)
		p.Match(tronParserLSB)
	}
	{
		p.SetState(252)
		p.OptionName()
	}
	{
		p.SetState(253)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(254)
		p.Constant()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(255)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(256)
			p.EnumValueOption()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(262)
		p.Match(tronParserRSB)
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
	p.RuleIndex = tronParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(tronParserSERVICE, 0)
}

func (s *ServiceContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
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

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *tronParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tronParserRULE_service)
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
		p.SetState(264)
		p.Match(tronParserSERVICE)
	}
	{
		p.SetState(265)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(266)
		p.Match(tronParserLCUR)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(tronParserSEMI-13))|(1<<(tronParserOPTION-13))|(1<<(tronParserRPC-13)))) != 0 {
		{
			p.SetState(267)
			p.ServiceBody()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(tronParserRCUR)
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

func (s *ServiceBodyContext) OptionService() IOptionServiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionServiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionServiceContext)
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
	p.EnterRule(localctx, 38, tronParserRULE_serviceBody)

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
	case tronParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.OptionService()
		}

	case tronParserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.Rpc()
		}

	case tronParserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(277)
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

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(tronParserRPC, 0)
}

func (s *RpcContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
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
	return s.GetToken(tronParserRETURNS, 0)
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

func (s *RpcContext) AllOptionRpc() []IOptionRpcContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionRpcContext)(nil)).Elem())
	var tst = make([]IOptionRpcContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionRpcContext)
		}
	}

	return tst
}

func (s *RpcContext) OptionRpc(i int) IOptionRpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionRpcContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionRpcContext)
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
	p.EnterRule(localctx, 40, tronParserRULE_rpc)
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
		p.Match(tronParserRPC)
	}
	{
		p.SetState(281)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(282)
		p.RpcParam()
	}
	{
		p.SetState(283)
		p.Match(tronParserRETURNS)
	}
	{
		p.SetState(284)
		p.RpcParam()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserLCUR:
		{
			p.SetState(285)
			p.Match(tronParserLCUR)
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserSEMI || _la == tronParserOPTION {
			p.SetState(288)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case tronParserOPTION:
				{
					p.SetState(286)
					p.OptionRpc()
				}

			case tronParserSEMI:
				{
					p.SetState(287)
					p.EmptyStatement()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(293)
			p.Match(tronParserRCUR)
		}

	case tronParserSEMI:
		{
			p.SetState(294)
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

func (s *RpcParamContext) STREAM() antlr.TerminalNode {
	return s.GetToken(tronParserSTREAM, 0)
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
	p.EnterRule(localctx, 42, tronParserRULE_rpcParam)
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
		p.SetState(297)
		p.Match(tronParserLPAREN)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserSTREAM {
		{
			p.SetState(298)
			p.Match(tronParserSTREAM)
		}

	}
	{
		p.SetState(301)
		p.MessageType()
	}
	{
		p.SetState(302)
		p.Match(tronParserRPAREN)
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

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(tronParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
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
	p.EnterRule(localctx, 44, tronParserRULE_reserved)

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
		p.SetState(304)
		p.Match(tronParserRESERVED)
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserInt_lit:
		{
			p.SetState(305)
			p.Ranges()
		}

	case tronParserStrLit:
		{
			p.SetState(306)
			p.FieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(309)
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
	p.EnterRule(localctx, 46, tronParserRULE_ranges)
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
		p.SetState(311)
		p.Rangee()
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(312)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(313)
			p.Rangee()
		}

		p.SetState(318)
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

func (s *RangeeContext) AllInt_lit() []antlr.TerminalNode {
	return s.GetTokens(tronParserInt_lit)
}

func (s *RangeeContext) Int_lit(i int) antlr.TerminalNode {
	return s.GetToken(tronParserInt_lit, i)
}

func (s *RangeeContext) TO() antlr.TerminalNode {
	return s.GetToken(tronParserTO, 0)
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
	p.EnterRule(localctx, 48, tronParserRULE_rangee)

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

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(tronParserInt_lit)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(tronParserInt_lit)
		}
		{
			p.SetState(321)
			p.Match(tronParserTO)
		}
		{
			p.SetState(322)
			p.Match(tronParserInt_lit)
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

func (s *FieldNamesContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(tronParserStrLit)
}

func (s *FieldNamesContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(tronParserStrLit, i)
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
	p.EnterRule(localctx, 50, tronParserRULE_fieldNames)
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
		p.SetState(325)
		p.Match(tronParserStrLit)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(326)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(327)
			p.Match(tronParserStrLit)
		}

		p.SetState(332)
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
	p.EnterRule(localctx, 52, tronParserRULE_typer)

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

	p.SetState(335)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserBOOL, tronParserBYTES, tronParserDOUBLE, tronParserFIXED32, tronParserFIXED64, tronParserFLOAT, tronParserINT32, tronParserINT64, tronParserSFIXED32, tronParserSFIXED64, tronParserSINT32, tronParserSINT64, tronParserSTRING, tronParserUINT32, tronParserUINT64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)
			p.Types()
		}

	case tronParserDOT, tronParserIdent:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
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
	p.RuleIndex = tronParserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(tronParserDOUBLE, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(tronParserFLOAT, 0)
}

func (s *TypesContext) INT32() antlr.TerminalNode {
	return s.GetToken(tronParserINT32, 0)
}

func (s *TypesContext) INT64() antlr.TerminalNode {
	return s.GetToken(tronParserINT64, 0)
}

func (s *TypesContext) UINT32() antlr.TerminalNode {
	return s.GetToken(tronParserUINT32, 0)
}

func (s *TypesContext) UINT64() antlr.TerminalNode {
	return s.GetToken(tronParserUINT64, 0)
}

func (s *TypesContext) SINT32() antlr.TerminalNode {
	return s.GetToken(tronParserSINT32, 0)
}

func (s *TypesContext) SINT64() antlr.TerminalNode {
	return s.GetToken(tronParserSINT64, 0)
}

func (s *TypesContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(tronParserFIXED32, 0)
}

func (s *TypesContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(tronParserFIXED64, 0)
}

func (s *TypesContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(tronParserSFIXED32, 0)
}

func (s *TypesContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(tronParserSFIXED64, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(tronParserBOOL, 0)
}

func (s *TypesContext) STRING() antlr.TerminalNode {
	return s.GetToken(tronParserSTRING, 0)
}

func (s *TypesContext) BYTES() antlr.TerminalNode {
	return s.GetToken(tronParserBYTES, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *tronParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, tronParserRULE_types)
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
		p.SetState(337)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(tronParserBOOL-45))|(1<<(tronParserBYTES-45))|(1<<(tronParserDOUBLE-45))|(1<<(tronParserFIXED32-45))|(1<<(tronParserFIXED64-45))|(1<<(tronParserFLOAT-45))|(1<<(tronParserINT32-45))|(1<<(tronParserINT64-45))|(1<<(tronParserSFIXED32-45))|(1<<(tronParserSFIXED64-45))|(1<<(tronParserSINT32-45))|(1<<(tronParserSINT64-45))|(1<<(tronParserSTRING-45))|(1<<(tronParserUINT32-45))|(1<<(tronParserUINT64-45)))) != 0) {
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

func (s *FieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *FieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tronParserInt_lit, 0)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tronParserSEMI, 0)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(tronParserREPEATED, 0)
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
	p.EnterRule(localctx, 56, tronParserRULE_field)
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
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserREPEATED {
		{
			p.SetState(339)
			p.Match(tronParserREPEATED)
		}

	}
	{
		p.SetState(342)
		p.Typer()
	}
	{
		p.SetState(343)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(344)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(345)
		p.Match(tronParserInt_lit)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(346)
			p.FieldOptions()
		}

	}
	{
		p.SetState(349)
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
	p.EnterRule(localctx, 58, tronParserRULE_fieldOptions)
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
		p.SetState(351)
		p.Match(tronParserLSB)
	}
	{
		p.SetState(352)
		p.FieldOption()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserCOMMA {
		{
			p.SetState(353)
			p.Match(tronParserCOMMA)
		}
		{
			p.SetState(354)
			p.FieldOption()
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(360)
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
	p.EnterRule(localctx, 60, tronParserRULE_fieldOption)

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
		p.OptionName()
	}
	{
		p.SetState(363)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(364)
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

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(tronParserONEOF, 0)
}

func (s *OneofContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
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
	p.EnterRule(localctx, 62, tronParserRULE_oneof)
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
		p.SetState(366)
		p.Match(tronParserONEOF)
	}
	{
		p.SetState(367)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(368)
		p.Match(tronParserLCUR)
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tronParserSEMI || _la == tronParserDOT || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(tronParserBOOL-45))|(1<<(tronParserBYTES-45))|(1<<(tronParserDOUBLE-45))|(1<<(tronParserFIXED32-45))|(1<<(tronParserFIXED64-45))|(1<<(tronParserFLOAT-45))|(1<<(tronParserINT32-45))|(1<<(tronParserINT64-45))|(1<<(tronParserSFIXED32-45))|(1<<(tronParserSFIXED64-45))|(1<<(tronParserSINT32-45))|(1<<(tronParserSINT64-45))|(1<<(tronParserSTRING-45))|(1<<(tronParserUINT32-45))|(1<<(tronParserUINT64-45))|(1<<(tronParserIdent-45)))) != 0) {
		p.SetState(371)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tronParserDOT, tronParserBOOL, tronParserBYTES, tronParserDOUBLE, tronParserFIXED32, tronParserFIXED64, tronParserFLOAT, tronParserINT32, tronParserINT64, tronParserSFIXED32, tronParserSFIXED64, tronParserSINT32, tronParserSINT64, tronParserSTRING, tronParserUINT32, tronParserUINT64, tronParserIdent:
			{
				p.SetState(369)
				p.OneofField()
			}

		case tronParserSEMI:
			{
				p.SetState(370)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(376)
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

func (s *OneofFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *OneofFieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tronParserInt_lit, 0)
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
	p.EnterRule(localctx, 64, tronParserRULE_oneofField)
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
		p.SetState(378)
		p.Typer()
	}
	{
		p.SetState(379)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(380)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(381)
		p.Match(tronParserInt_lit)
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(382)
			p.FieldOptions()
		}

	}
	{
		p.SetState(385)
		p.Match(tronParserSEMI)
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

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(tronParserMAP, 0)
}

func (s *MapFieldContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(tronParserLCHEVR, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
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

func (s *MapFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(tronParserEQ, 0)
}

func (s *MapFieldContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tronParserInt_lit, 0)
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
	p.EnterRule(localctx, 66, tronParserRULE_mapField)
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
		p.SetState(387)
		p.Match(tronParserMAP)
	}
	{
		p.SetState(388)
		p.Match(tronParserLCHEVR)
	}
	{
		p.SetState(389)
		p.KeyType()
	}
	{
		p.SetState(390)
		p.Match(tronParserCOMMA)
	}
	{
		p.SetState(391)
		p.Typer()
	}
	{
		p.SetState(392)
		p.Match(tronParserRCHEVR)
	}
	{
		p.SetState(393)
		p.Match(tronParserIdent)
	}
	{
		p.SetState(394)
		p.Match(tronParserEQ)
	}
	{
		p.SetState(395)
		p.Match(tronParserInt_lit)
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserLSB {
		{
			p.SetState(396)
			p.FieldOptions()
		}

	}
	{
		p.SetState(399)
		p.Match(tronParserSEMI)
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
	p.RuleIndex = tronParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(tronParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(tronParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(tronParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(tronParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(tronParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(tronParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(tronParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(tronParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(tronParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(tronParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(tronParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(tronParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *tronParser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, tronParserRULE_keyType)
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
		p.SetState(401)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(tronParserBOOL-45))|(1<<(tronParserFIXED32-45))|(1<<(tronParserFIXED64-45))|(1<<(tronParserINT32-45))|(1<<(tronParserINT64-45))|(1<<(tronParserSFIXED32-45))|(1<<(tronParserSFIXED64-45))|(1<<(tronParserSINT32-45))|(1<<(tronParserSINT64-45))|(1<<(tronParserSTRING-45))|(1<<(tronParserUINT32-45))|(1<<(tronParserUINT64-45)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *MessageTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tronParserIdent)
}

func (s *MessageTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tronParserIdent, i)
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
	p.EnterRule(localctx, 70, tronParserRULE_messageType)
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
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDOT {
		{
			p.SetState(403)
			p.Match(tronParserDOT)
		}

	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(406)
				p.Match(tronParserIdent)
			}
			{
				p.SetState(407)
				p.Match(tronParserDOT)
			}

		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}
	{
		p.SetState(413)
		p.Match(tronParserIdent)
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

func (s *MessageOrEnumTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tronParserIdent)
}

func (s *MessageOrEnumTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tronParserIdent, i)
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
	p.EnterRule(localctx, 72, tronParserRULE_messageOrEnumType)
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
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tronParserDOT {
		{
			p.SetState(415)
			p.Match(tronParserDOT)
		}

	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(418)
				p.Match(tronParserIdent)
			}
			{
				p.SetState(419)
				p.Match(tronParserDOT)
			}

		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}
	{
		p.SetState(425)
		p.Match(tronParserIdent)
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
	p.RuleIndex = tronParserRULE_bool_lit
	return p
}

func (*Bool_litContext) IsBool_litContext() {}

func NewBool_litContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_litContext {
	var p = new(Bool_litContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_bool_lit

	return p
}

func (s *Bool_litContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_litContext) TRUE() antlr.TerminalNode {
	return s.GetToken(tronParserTRUE, 0)
}

func (s *Bool_litContext) FALSE() antlr.TerminalNode {
	return s.GetToken(tronParserFALSE, 0)
}

func (s *Bool_litContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_litContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_litContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterBool_lit(s)
	}
}

func (s *Bool_litContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitBool_lit(s)
	}
}

func (p *tronParser) Bool_lit() (localctx IBool_litContext) {
	localctx = NewBool_litContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, tronParserRULE_bool_lit)
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
		p.SetState(427)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tronParserTRUE || _la == tronParserFALSE) {
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
	p.EnterRule(localctx, 76, tronParserRULE_emptyStatement)

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
		p.SetState(429)
		p.Match(tronParserSEMI)
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

func (s *ConstantContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(tronParserIdent)
}

func (s *ConstantContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(tronParserIdent, i)
}

func (s *ConstantContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tronParserDOT)
}

func (s *ConstantContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tronParserDOT, i)
}

func (s *ConstantContext) Int_lit() antlr.TerminalNode {
	return s.GetToken(tronParserInt_lit, 0)
}

func (s *ConstantContext) DASH() antlr.TerminalNode {
	return s.GetToken(tronParserDASH, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(tronParserPLUS, 0)
}

func (s *ConstantContext) Float_lit() antlr.TerminalNode {
	return s.GetToken(tronParserFloat_lit, 0)
}

func (s *ConstantContext) StrLit() antlr.TerminalNode {
	return s.GetToken(tronParserStrLit, 0)
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
	p.EnterRule(localctx, 78, tronParserRULE_constant)
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

	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(tronParserIdent)
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == tronParserDOT {
			{
				p.SetState(432)
				p.Match(tronParserDOT)
			}
			{
				p.SetState(433)
				p.Match(tronParserIdent)
			}

			p.SetState(438)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tronParserDASH || _la == tronParserPLUS {
			{
				p.SetState(439)
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
			p.SetState(442)
			p.Match(tronParserInt_lit)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tronParserDASH || _la == tronParserPLUS {
			{
				p.SetState(443)
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
			p.SetState(446)
			p.Match(tronParserFloat_lit)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(449)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case tronParserStrLit:
			{
				p.SetState(447)
				p.Match(tronParserStrLit)
			}

		case tronParserTRUE, tronParserFALSE:
			{
				p.SetState(448)
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
	p.RuleIndex = tronParserRULE_pron
	return p
}

func (*PronContext) IsPronContext() {}

func NewPronContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PronContext {
	var p = new(PronContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_pron

	return p
}

func (s *PronContext) GetParser() antlr.Parser { return s.parser }

func (s *PronContext) GetId() antlr.Token { return s.id }

func (s *PronContext) SetId(v antlr.Token) { s.id = v }

func (s *PronContext) LCUR() antlr.TerminalNode {
	return s.GetToken(tronParserLCUR, 0)
}

func (s *PronContext) PronElem() IPronElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPronElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPronElemContext)
}

func (s *PronContext) RCUR() antlr.TerminalNode {
	return s.GetToken(tronParserRCUR, 0)
}

func (s *PronContext) Ident() antlr.TerminalNode {
	return s.GetToken(tronParserIdent, 0)
}

func (s *PronContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PronContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PronContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterPron(s)
	}
}

func (s *PronContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitPron(s)
	}
}

func (p *tronParser) Pron() (localctx IPronContext) {
	localctx = NewPronContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, tronParserRULE_pron)

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
		p.SetState(453)
		p.Match(tronParserLCUR)
	}
	{
		p.SetState(454)

		var _m = p.Match(tronParserIdent)

		localctx.(*PronContext).id = _m
	}
	{
		p.SetState(455)
		p.Match(tronParserCOLON)
	}
	{
		p.SetState(456)
		p.PronElem()
	}
	{
		p.SetState(457)
		p.Match(tronParserRCUR)
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
	p.RuleIndex = tronParserRULE_pronElem
	return p
}

func (*PronElemContext) IsPronElemContext() {}

func NewPronElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PronElemContext {
	var p = new(PronElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tronParserRULE_pronElem

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
	return s.GetToken(tronParserStrLit, 0)
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
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.EnterPronOBJ(s)
	}
}

func (s *PronOBJContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tronParserListener); ok {
		listenerT.ExitPronOBJ(s)
	}
}

func (p *tronParser) PronElem() (localctx IPronElemContext) {
	localctx = NewPronElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, tronParserRULE_pronElem)

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

	p.SetState(461)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tronParserStrLit:
		localctx = NewPronSTRContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.Match(tronParserStrLit)
		}

	case tronParserLCUR:
		localctx = NewPronOBJContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.Pron()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
