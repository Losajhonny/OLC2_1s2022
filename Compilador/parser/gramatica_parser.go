// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Compilador/gen"
import "Compilador/err"
import "Compilador/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 53, 356,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 83, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	7, 4, 118, 10, 4, 12, 4, 14, 4, 121, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 135, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 148, 10,
	6, 12, 6, 14, 6, 151, 11, 6, 3, 7, 7, 7, 154, 10, 7, 12, 7, 14, 7, 157,
	11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	169, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 190, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 201,
	10, 10, 12, 10, 14, 10, 204, 11, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 5, 12, 222, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 236, 10, 13, 12, 13, 14, 13,
	239, 11, 13, 3, 13, 3, 13, 5, 13, 243, 10, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 256, 10, 14,
	3, 14, 3, 14, 6, 14, 260, 10, 14, 13, 14, 14, 14, 261, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 268, 10, 14, 5, 14, 270, 10, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 285, 10, 15, 3, 15, 3, 15, 6, 15, 289, 10, 15, 13, 15, 14, 15, 290,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 298, 10, 15, 3, 15, 3, 15, 5,
	15, 302, 10, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 2, 4, 6, 10, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 4, 3, 2, 6, 8, 3, 2,
	3, 4, 2, 376, 2, 46, 3, 2, 2, 2, 4, 50, 3, 2, 2, 2, 6, 82, 3, 2, 2, 2,
	8, 134, 3, 2, 2, 2, 10, 136, 3, 2, 2, 2, 12, 155, 3, 2, 2, 2, 14, 168,
	3, 2, 2, 2, 16, 189, 3, 2, 2, 2, 18, 191, 3, 2, 2, 2, 20, 205, 3, 2, 2,
	2, 22, 221, 3, 2, 2, 2, 24, 223, 3, 2, 2, 2, 26, 246, 3, 2, 2, 2, 28, 274,
	3, 2, 2, 2, 30, 306, 3, 2, 2, 2, 32, 313, 3, 2, 2, 2, 34, 321, 3, 2, 2,
	2, 36, 326, 3, 2, 2, 2, 38, 337, 3, 2, 2, 2, 40, 341, 3, 2, 2, 2, 42, 345,
	3, 2, 2, 2, 44, 351, 3, 2, 2, 2, 46, 47, 8, 2, 1, 2, 47, 48, 5, 12, 7,
	2, 48, 49, 7, 2, 2, 3, 49, 3, 3, 2, 2, 2, 50, 51, 8, 3, 1, 2, 51, 5, 3,
	2, 2, 2, 52, 53, 8, 4, 1, 2, 53, 54, 7, 3, 2, 2, 54, 55, 5, 6, 4, 17, 55,
	56, 8, 4, 1, 2, 56, 83, 3, 2, 2, 2, 57, 58, 7, 4, 2, 2, 58, 59, 5, 6, 4,
	16, 59, 60, 8, 4, 1, 2, 60, 83, 3, 2, 2, 2, 61, 62, 7, 5, 2, 2, 62, 63,
	5, 6, 4, 15, 63, 64, 8, 4, 1, 2, 64, 83, 3, 2, 2, 2, 65, 66, 7, 12, 2,
	2, 66, 67, 5, 6, 4, 2, 67, 68, 7, 13, 2, 2, 68, 69, 8, 4, 1, 2, 69, 83,
	3, 2, 2, 2, 70, 71, 7, 47, 2, 2, 71, 83, 8, 4, 1, 2, 72, 73, 7, 50, 2,
	2, 73, 83, 8, 4, 1, 2, 74, 75, 7, 14, 2, 2, 75, 83, 8, 4, 1, 2, 76, 77,
	7, 15, 2, 2, 77, 83, 8, 4, 1, 2, 78, 79, 5, 10, 6, 2, 79, 80, 7, 16, 2,
	2, 80, 81, 8, 4, 1, 2, 81, 83, 3, 2, 2, 2, 82, 52, 3, 2, 2, 2, 82, 57,
	3, 2, 2, 2, 82, 61, 3, 2, 2, 2, 82, 65, 3, 2, 2, 2, 82, 70, 3, 2, 2, 2,
	82, 72, 3, 2, 2, 2, 82, 74, 3, 2, 2, 2, 82, 76, 3, 2, 2, 2, 82, 78, 3,
	2, 2, 2, 83, 119, 3, 2, 2, 2, 84, 85, 12, 14, 2, 2, 85, 86, 9, 2, 2, 2,
	86, 87, 5, 6, 4, 15, 87, 88, 8, 4, 1, 2, 88, 118, 3, 2, 2, 2, 89, 90, 12,
	13, 2, 2, 90, 91, 9, 3, 2, 2, 91, 92, 5, 6, 4, 14, 92, 93, 8, 4, 1, 2,
	93, 118, 3, 2, 2, 2, 94, 95, 12, 12, 2, 2, 95, 96, 5, 8, 5, 2, 96, 97,
	5, 6, 4, 13, 97, 98, 8, 4, 1, 2, 98, 118, 3, 2, 2, 2, 99, 100, 12, 11,
	2, 2, 100, 101, 7, 9, 2, 2, 101, 102, 5, 4, 3, 2, 102, 103, 5, 6, 4, 12,
	103, 104, 8, 4, 1, 2, 104, 118, 3, 2, 2, 2, 105, 106, 12, 10, 2, 2, 106,
	107, 7, 10, 2, 2, 107, 108, 5, 4, 3, 2, 108, 109, 5, 6, 4, 11, 109, 110,
	8, 4, 1, 2, 110, 118, 3, 2, 2, 2, 111, 112, 12, 9, 2, 2, 112, 113, 7, 11,
	2, 2, 113, 114, 5, 4, 3, 2, 114, 115, 5, 6, 4, 10, 115, 116, 8, 4, 1, 2,
	116, 118, 3, 2, 2, 2, 117, 84, 3, 2, 2, 2, 117, 89, 3, 2, 2, 2, 117, 94,
	3, 2, 2, 2, 117, 99, 3, 2, 2, 2, 117, 105, 3, 2, 2, 2, 117, 111, 3, 2,
	2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2,
	120, 7, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 17, 2, 2, 123, 135,
	8, 5, 1, 2, 124, 125, 7, 18, 2, 2, 125, 135, 8, 5, 1, 2, 126, 127, 7, 19,
	2, 2, 127, 135, 8, 5, 1, 2, 128, 129, 7, 20, 2, 2, 129, 135, 8, 5, 1, 2,
	130, 131, 7, 21, 2, 2, 131, 135, 8, 5, 1, 2, 132, 133, 7, 22, 2, 2, 133,
	135, 8, 5, 1, 2, 134, 122, 3, 2, 2, 2, 134, 124, 3, 2, 2, 2, 134, 126,
	3, 2, 2, 2, 134, 128, 3, 2, 2, 2, 134, 130, 3, 2, 2, 2, 134, 132, 3, 2,
	2, 2, 135, 9, 3, 2, 2, 2, 136, 137, 8, 6, 1, 2, 137, 138, 7, 50, 2, 2,
	138, 139, 7, 24, 2, 2, 139, 140, 5, 6, 4, 2, 140, 141, 8, 6, 1, 2, 141,
	149, 3, 2, 2, 2, 142, 143, 12, 4, 2, 2, 143, 144, 7, 23, 2, 2, 144, 145,
	5, 6, 4, 2, 145, 146, 8, 6, 1, 2, 146, 148, 3, 2, 2, 2, 147, 142, 3, 2,
	2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2,
	150, 11, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 154, 5, 14, 8, 2, 153,
	152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156,
	3, 2, 2, 2, 156, 13, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 169, 5, 16,
	9, 2, 159, 169, 5, 22, 12, 2, 160, 169, 5, 24, 13, 2, 161, 169, 5, 28,
	15, 2, 162, 169, 5, 30, 16, 2, 163, 169, 5, 32, 17, 2, 164, 169, 5, 34,
	18, 2, 165, 169, 5, 36, 19, 2, 166, 169, 5, 38, 20, 2, 167, 169, 5, 40,
	21, 2, 168, 158, 3, 2, 2, 2, 168, 159, 3, 2, 2, 2, 168, 160, 3, 2, 2, 2,
	168, 161, 3, 2, 2, 2, 168, 162, 3, 2, 2, 2, 168, 163, 3, 2, 2, 2, 168,
	164, 3, 2, 2, 2, 168, 165, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 167,
	3, 2, 2, 2, 169, 15, 3, 2, 2, 2, 170, 171, 7, 25, 2, 2, 171, 172, 7, 50,
	2, 2, 172, 173, 7, 26, 2, 2, 173, 174, 5, 20, 11, 2, 174, 175, 7, 27, 2,
	2, 175, 176, 8, 9, 1, 2, 176, 190, 3, 2, 2, 2, 177, 178, 7, 25, 2, 2, 178,
	179, 7, 50, 2, 2, 179, 180, 7, 26, 2, 2, 180, 181, 7, 28, 2, 2, 181, 182,
	7, 24, 2, 2, 182, 183, 5, 18, 10, 2, 183, 184, 7, 16, 2, 2, 184, 185, 7,
	29, 2, 2, 185, 186, 5, 20, 11, 2, 186, 187, 7, 27, 2, 2, 187, 188, 8, 9,
	1, 2, 188, 190, 3, 2, 2, 2, 189, 170, 3, 2, 2, 2, 189, 177, 3, 2, 2, 2,
	190, 17, 3, 2, 2, 2, 191, 192, 7, 47, 2, 2, 192, 193, 7, 30, 2, 2, 193,
	194, 7, 47, 2, 2, 194, 202, 8, 10, 1, 2, 195, 196, 7, 23, 2, 2, 196, 197,
	7, 47, 2, 2, 197, 198, 7, 30, 2, 2, 198, 199, 7, 47, 2, 2, 199, 201, 8,
	10, 1, 2, 200, 195, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2,
	2, 202, 203, 3, 2, 2, 2, 203, 19, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205,
	206, 7, 31, 2, 2, 206, 207, 8, 11, 1, 2, 207, 21, 3, 2, 2, 2, 208, 209,
	7, 50, 2, 2, 209, 210, 7, 32, 2, 2, 210, 211, 5, 6, 4, 2, 211, 212, 7,
	27, 2, 2, 212, 213, 8, 12, 1, 2, 213, 222, 3, 2, 2, 2, 214, 215, 5, 10,
	6, 2, 215, 216, 7, 16, 2, 2, 216, 217, 7, 32, 2, 2, 217, 218, 5, 6, 4,
	2, 218, 219, 7, 27, 2, 2, 219, 220, 8, 12, 1, 2, 220, 222, 3, 2, 2, 2,
	221, 208, 3, 2, 2, 2, 221, 214, 3, 2, 2, 2, 222, 23, 3, 2, 2, 2, 223, 224,
	7, 33, 2, 2, 224, 225, 5, 6, 4, 2, 225, 226, 8, 13, 1, 2, 226, 227, 5,
	42, 22, 2, 227, 237, 8, 13, 1, 2, 228, 229, 7, 34, 2, 2, 229, 230, 7, 33,
	2, 2, 230, 231, 5, 6, 4, 2, 231, 232, 8, 13, 1, 2, 232, 233, 5, 42, 22,
	2, 233, 234, 8, 13, 1, 2, 234, 236, 3, 2, 2, 2, 235, 228, 3, 2, 2, 2, 236,
	239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 242,
	3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 7, 34, 2, 2, 241, 243, 5, 42,
	22, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2,
	244, 245, 8, 13, 1, 2, 245, 25, 3, 2, 2, 2, 246, 247, 7, 35, 2, 2, 247,
	248, 5, 6, 4, 2, 248, 259, 7, 36, 2, 2, 249, 250, 7, 37, 2, 2, 250, 251,
	5, 6, 4, 2, 251, 252, 7, 26, 2, 2, 252, 255, 8, 14, 1, 2, 253, 256, 5,
	42, 22, 2, 254, 256, 5, 44, 23, 2, 255, 253, 3, 2, 2, 2, 255, 254, 3, 2,
	2, 2, 256, 257, 3, 2, 2, 2, 257, 258, 8, 14, 1, 2, 258, 260, 3, 2, 2, 2,
	259, 249, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261,
	262, 3, 2, 2, 2, 262, 269, 3, 2, 2, 2, 263, 264, 7, 38, 2, 2, 264, 267,
	7, 26, 2, 2, 265, 268, 5, 42, 22, 2, 266, 268, 5, 44, 23, 2, 267, 265,
	3, 2, 2, 2, 267, 266, 3, 2, 2, 2, 268, 270, 3, 2, 2, 2, 269, 263, 3, 2,
	2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 272, 8, 14, 1, 2,
	272, 273, 7, 39, 2, 2, 273, 27, 3, 2, 2, 2, 274, 275, 7, 35, 2, 2, 275,
	276, 5, 6, 4, 2, 276, 277, 7, 36, 2, 2, 277, 288, 8, 15, 1, 2, 278, 279,
	7, 37, 2, 2, 279, 280, 5, 6, 4, 2, 280, 281, 7, 26, 2, 2, 281, 284, 8,
	15, 1, 2, 282, 285, 5, 42, 22, 2, 283, 285, 5, 44, 23, 2, 284, 282, 3,
	2, 2, 2, 284, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 8, 15, 1,
	2, 287, 289, 3, 2, 2, 2, 288, 278, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290,
	288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 301, 3, 2, 2, 2, 292, 293,
	7, 38, 2, 2, 293, 294, 7, 26, 2, 2, 294, 297, 8, 15, 1, 2, 295, 298, 5,
	42, 22, 2, 296, 298, 5, 44, 23, 2, 297, 295, 3, 2, 2, 2, 297, 296, 3, 2,
	2, 2, 298, 299, 3, 2, 2, 2, 299, 300, 8, 15, 1, 2, 300, 302, 3, 2, 2, 2,
	301, 292, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303,
	304, 8, 15, 1, 2, 304, 305, 7, 39, 2, 2, 305, 29, 3, 2, 2, 2, 306, 307,
	7, 40, 2, 2, 307, 308, 8, 16, 1, 2, 308, 309, 5, 6, 4, 2, 309, 310, 8,
	16, 1, 2, 310, 311, 5, 42, 22, 2, 311, 312, 8, 16, 1, 2, 312, 31, 3, 2,
	2, 2, 313, 314, 7, 41, 2, 2, 314, 315, 8, 17, 1, 2, 315, 316, 5, 42, 22,
	2, 316, 317, 7, 40, 2, 2, 317, 318, 5, 6, 4, 2, 318, 319, 7, 27, 2, 2,
	319, 320, 8, 17, 1, 2, 320, 33, 3, 2, 2, 2, 321, 322, 7, 42, 2, 2, 322,
	323, 8, 18, 1, 2, 323, 324, 5, 42, 22, 2, 324, 325, 8, 18, 1, 2, 325, 35,
	3, 2, 2, 2, 326, 327, 7, 43, 2, 2, 327, 328, 7, 50, 2, 2, 328, 329, 7,
	32, 2, 2, 329, 330, 5, 6, 4, 2, 330, 331, 7, 44, 2, 2, 331, 332, 5, 6,
	4, 2, 332, 333, 8, 19, 1, 2, 333, 334, 7, 41, 2, 2, 334, 335, 5, 42, 22,
	2, 335, 336, 8, 19, 1, 2, 336, 37, 3, 2, 2, 2, 337, 338, 7, 45, 2, 2, 338,
	339, 7, 27, 2, 2, 339, 340, 8, 20, 1, 2, 340, 39, 3, 2, 2, 2, 341, 342,
	7, 46, 2, 2, 342, 343, 7, 27, 2, 2, 343, 344, 8, 21, 1, 2, 344, 41, 3,
	2, 2, 2, 345, 346, 8, 22, 1, 2, 346, 347, 7, 36, 2, 2, 347, 348, 5, 12,
	7, 2, 348, 349, 7, 39, 2, 2, 349, 350, 8, 22, 1, 2, 350, 43, 3, 2, 2, 2,
	351, 352, 8, 23, 1, 2, 352, 353, 5, 12, 7, 2, 353, 354, 8, 23, 1, 2, 354,
	45, 3, 2, 2, 2, 22, 82, 117, 119, 134, 149, 155, 168, 189, 202, 221, 237,
	242, 255, 261, 267, 269, 284, 290, 297, 301,
}
var literalNames = []string{
	"", "'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "'&&'", "'^'", "'||'", "'('",
	"')'", "'true'", "'false'", "']'", "'=='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "','", "'['", "'let'", "':'", "';'", "'array'", "'of'", "'..'",
	"'int'", "'='", "'if'", "'else'", "'switch'", "'{'", "'case'", "'default'",
	"'}'", "'while'", "'do'", "'loop'", "'for'", "'to'", "'break'", "'continue'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "NUM", "FLOAT", "STRING", "ID", "MULCOMMENT",
	"UNICOMMENT", "WHITESPACE",
}

var ruleNames = []string{
	"start", "marcador", "expresion", "oprel", "lref", "instrucciones", "instruccion",
	"inst_declaracion", "ldims", "tipo", "inst_asignacion", "inst_if", "inst_switch_propuesta1",
	"inst_switch", "inst_while", "inst_doWhile", "inst_loop", "inst_for", "inst_break",
	"inst_continue", "bloque", "bloqueSinLlaves",
}

type GramaticaParser struct {
	*antlr.BaseParser
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GramaticaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// posición relativa de un símbolo
var desp int = 0

// entorno actual
var tope *entorno.Entorno

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserT__0       = 1
	GramaticaParserT__1       = 2
	GramaticaParserT__2       = 3
	GramaticaParserT__3       = 4
	GramaticaParserT__4       = 5
	GramaticaParserT__5       = 6
	GramaticaParserT__6       = 7
	GramaticaParserT__7       = 8
	GramaticaParserT__8       = 9
	GramaticaParserT__9       = 10
	GramaticaParserT__10      = 11
	GramaticaParserT__11      = 12
	GramaticaParserT__12      = 13
	GramaticaParserT__13      = 14
	GramaticaParserT__14      = 15
	GramaticaParserT__15      = 16
	GramaticaParserT__16      = 17
	GramaticaParserT__17      = 18
	GramaticaParserT__18      = 19
	GramaticaParserT__19      = 20
	GramaticaParserT__20      = 21
	GramaticaParserT__21      = 22
	GramaticaParserT__22      = 23
	GramaticaParserT__23      = 24
	GramaticaParserT__24      = 25
	GramaticaParserT__25      = 26
	GramaticaParserT__26      = 27
	GramaticaParserT__27      = 28
	GramaticaParserT__28      = 29
	GramaticaParserT__29      = 30
	GramaticaParserT__30      = 31
	GramaticaParserT__31      = 32
	GramaticaParserT__32      = 33
	GramaticaParserT__33      = 34
	GramaticaParserT__34      = 35
	GramaticaParserT__35      = 36
	GramaticaParserT__36      = 37
	GramaticaParserT__37      = 38
	GramaticaParserT__38      = 39
	GramaticaParserT__39      = 40
	GramaticaParserT__40      = 41
	GramaticaParserT__41      = 42
	GramaticaParserT__42      = 43
	GramaticaParserT__43      = 44
	GramaticaParserNUM        = 45
	GramaticaParserFLOAT      = 46
	GramaticaParserSTRING     = 47
	GramaticaParserID         = 48
	GramaticaParserMULCOMMENT = 49
	GramaticaParserUNICOMMENT = 50
	GramaticaParserWHITESPACE = 51
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start                  = 0
	GramaticaParserRULE_marcador               = 1
	GramaticaParserRULE_expresion              = 2
	GramaticaParserRULE_oprel                  = 3
	GramaticaParserRULE_lref                   = 4
	GramaticaParserRULE_instrucciones          = 5
	GramaticaParserRULE_instruccion            = 6
	GramaticaParserRULE_inst_declaracion       = 7
	GramaticaParserRULE_ldims                  = 8
	GramaticaParserRULE_tipo                   = 9
	GramaticaParserRULE_inst_asignacion        = 10
	GramaticaParserRULE_inst_if                = 11
	GramaticaParserRULE_inst_switch_propuesta1 = 12
	GramaticaParserRULE_inst_switch            = 13
	GramaticaParserRULE_inst_while             = 14
	GramaticaParserRULE_inst_doWhile           = 15
	GramaticaParserRULE_inst_loop              = 16
	GramaticaParserRULE_inst_for               = 17
	GramaticaParserRULE_inst_break             = 18
	GramaticaParserRULE_inst_continue          = 19
	GramaticaParserRULE_bloque                 = 20
	GramaticaParserRULE_bloqueSinLlaves        = 21
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

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
	tope = entorno.NewEntorno(nil)
	{
		p.SetState(45)
		p.Instrucciones()
	}
	{
		p.SetState(46)
		p.Match(GramaticaParserEOF)
	}

	return localctx
}

// IMarcadorContext is an interface to support dynamic dispatch.
type IMarcadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista returns the lista attribute.
	GetLista() []string

	// SetLista sets the lista attribute.
	SetLista([]string)

	// IsMarcadorContext differentiates from other interfaces.
	IsMarcadorContext()
}

type MarcadorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  []string
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_marcador
	return p
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, lista []string) *MarcadorContext {
	var p = new(MarcadorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_marcador

	p.lista = lista

	return p
}

func (s *MarcadorContext) GetParser() antlr.Parser { return s.parser }

func (s *MarcadorContext) GetLista() []string { return s.lista }

func (s *MarcadorContext) SetLista(v []string) { s.lista = v }

func (s *MarcadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarcadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterMarcador(s)
	}
}

func (s *MarcadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitMarcador(s)
	}
}

func (p *GramaticaParser) Marcador(lista []string) (localctx IMarcadorContext) {
	this := p
	_ = this

	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState(), lista)
	p.EnterRule(localctx, 2, GramaticaParserRULE_marcador)

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
	gen.Soltar(lista)

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// GetRef returns the ref rule contexts.
	GetRef() ILrefContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// SetRef sets the ref rule contexts.
	SetRef(ILrefContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// GetDir returns the dir attribute.
	GetDir() string

	// GetLv returns the lv attribute.
	GetLv() []string

	// GetLf returns the lf attribute.
	GetLf() []string

	// GetCad returns the cad attribute.
	GetCad() string

	// SetDir sets the dir attribute.
	SetDir(string)

	// SetLv sets the lv attribute.
	SetLv([]string)

	// SetLf sets the lf attribute.
	SetLf([]string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	lv     []string
	lf     []string
	cad    string
	e1     IExpresionContext
	op     antlr.Token
	e      IExpresionContext
	n      antlr.Token
	id     antlr.Token
	ref    ILrefContext
	e2     IExpresionContext
	opr    IOprelContext
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) GetN() antlr.Token { return s.n }

func (s *ExpresionContext) GetId() antlr.Token { return s.id }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) SetN(v antlr.Token) { s.n = v }

func (s *ExpresionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpresionContext) GetE1() IExpresionContext { return s.e1 }

func (s *ExpresionContext) GetE() IExpresionContext { return s.e }

func (s *ExpresionContext) GetRef() ILrefContext { return s.ref }

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) GetOpr() IOprelContext { return s.opr }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetE(v IExpresionContext) { s.e = v }

func (s *ExpresionContext) SetRef(v ILrefContext) { s.ref = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *ExpresionContext) GetDir() string { return s.dir }

func (s *ExpresionContext) GetLv() []string { return s.lv }

func (s *ExpresionContext) GetLf() []string { return s.lf }

func (s *ExpresionContext) GetCad() string { return s.cad }

func (s *ExpresionContext) SetDir(v string) { s.dir = v }

func (s *ExpresionContext) SetLv(v []string) { s.lv = v }

func (s *ExpresionContext) SetLf(v []string) { s.lf = v }

func (s *ExpresionContext) SetCad(v string) { s.cad = v }

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *ExpresionContext) Lref() ILrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILrefContext)
}

func (s *ExpresionContext) Oprel() IOprelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOprelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOprelContext)
}

func (s *ExpresionContext) Marcador() IMarcadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarcadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *GramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *GramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GramaticaParserRULE_expresion, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(51)

			var _m = p.Match(GramaticaParserT__0)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(52)

			var _x = p.expresion(15)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.AddExpresionUnaria(localctx.(*ExpresionContext).dir, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetE().GetDir())

	case 2:
		{
			p.SetState(55)

			var _m = p.Match(GramaticaParserT__1)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(56)

			var _x = p.expresion(14)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.AddExpresionUnaria(localctx.(*ExpresionContext).dir, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetE().GetDir())

	case 3:
		{
			p.SetState(59)

			var _m = p.Match(GramaticaParserT__2)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(60)

			var _x = p.expresion(13)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLv()

	case 4:
		{
			p.SetState(63)
			p.Match(GramaticaParserT__9)
		}
		{
			p.SetState(64)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e = _x
		}
		{
			p.SetState(65)
			p.Match(GramaticaParserT__10)
		}

		localctx.(*ExpresionContext).dir = localctx.(*ExpresionContext).GetE().GetDir()
		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLv()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).cad = localctx.(*ExpresionContext).GetE().GetCad()

	case 5:
		{
			p.SetState(68)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*ExpresionContext).n = _m
		}

		localctx.(*ExpresionContext).dir = (func() string {
			if localctx.(*ExpresionContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetN().GetText()
			}
		}())
		localctx.(*ExpresionContext).cad = ""

	case 6:
		{
			p.SetState(70)

			var _m = p.Match(GramaticaParserID)

			localctx.(*ExpresionContext).id = _m
		}

		localctx.(*ExpresionContext).dir = (func() string {
			if localctx.(*ExpresionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetId().GetText()
			}
		}())
		localctx.(*ExpresionContext).cad = ""

	case 7:
		{
			p.SetState(72)
			p.Match(GramaticaParserT__11)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewLabel())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewLabel())
		gen.AddGoto(localctx.(*ExpresionContext).lv[0])

	case 8:
		{
			p.SetState(74)
			p.Match(GramaticaParserT__12)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewLabel())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewLabel())
		gen.AddGoto(localctx.(*ExpresionContext).lf[0])

	case 9:
		{
			p.SetState(76)

			var _x = p.lref(0)

			localctx.(*ExpresionContext).ref = _x
		}
		{
			p.SetState(77)
			p.Match(GramaticaParserT__13)
		}

		// orden de columnas
		// localctx.(*ExpresionContext).dir = gen.NewTemp()
		// gen.AddGetArray(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetRef().GetId(), localctx.(*ExpresionContext).GetRef().GetAux())

		// orden de filas
		tmp := gen.NewTemp()
		gen.AddExpresion(tmp, localctx.(*ExpresionContext).GetRef().GetAux(), "+", localctx.(*ExpresionContext).GetRef().GetDir())
		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.AddGetArray(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetRef().GetId(), tmp)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(83)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserT__3)|(1<<GramaticaParserT__4)|(1<<GramaticaParserT__5))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(84)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.AddExpresion(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetE1().GetDir(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetE2().GetDir())

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__0 || _la == GramaticaParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.AddExpresion(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetE1().GetDir(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetE2().GetDir())

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(93)

					var _x = p.Oprel()

					localctx.(*ExpresionContext).opr = _x
				}
				{
					p.SetState(94)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewLabel())
				localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewLabel())
				localctx.(*ExpresionContext).cad = gen.AddIf(localctx.(*ExpresionContext).GetE1().GetDir(), localctx.(*ExpresionContext).GetOpr().GetOp(), localctx.(*ExpresionContext).GetE2().GetDir(), localctx.(*ExpresionContext).lv[0])
				gen.AddGoto(localctx.(*ExpresionContext).lf[0])

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(98)

					var _m = p.Match(GramaticaParserT__6)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(99)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLv())
				}
				{
					p.SetState(100)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLf(), localctx.(*ExpresionContext).GetE2().GetLf())

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(104)

					var _m = p.Match(GramaticaParserT__7)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(105)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(106)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).e2 = _x
				}

				gen.Soltar(localctx.(*ExpresionContext).GetE1().GetLv())
				gen.AddIfCad(localctx.(*ExpresionContext).GetE2().GetCad(), localctx.(*ExpresionContext).GetE2().GetLf()[0])
				gen.AddGoto(localctx.(*ExpresionContext).GetE2().GetLv()[0])
				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(110)

					var _m = p.Match(GramaticaParserT__8)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(111)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(112)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()
				localctx.(*ExpresionContext).lv = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLv(), localctx.(*ExpresionContext).GetE2().GetLv())

			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IOprelContext is an interface to support dynamic dispatch.
type IOprelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// SetOpe sets the ope token.
	SetOpe(antlr.Token)

	// GetOp returns the op attribute.
	GetOp() string

	// SetOp sets the op attribute.
	SetOp(string)

	// IsOprelContext differentiates from other interfaces.
	IsOprelContext()
}

type OprelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	ope    antlr.Token
}

func NewEmptyOprelContext() *OprelContext {
	var p = new(OprelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_oprel
	return p
}

func (*OprelContext) IsOprelContext() {}

func NewOprelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OprelContext {
	var p = new(OprelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_oprel

	return p
}

func (s *OprelContext) GetParser() antlr.Parser { return s.parser }

func (s *OprelContext) GetOpe() antlr.Token { return s.ope }

func (s *OprelContext) SetOpe(v antlr.Token) { s.ope = v }

func (s *OprelContext) GetOp() string { return s.op }

func (s *OprelContext) SetOp(v string) { s.op = v }

func (s *OprelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OprelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OprelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterOprel(s)
	}
}

func (s *OprelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitOprel(s)
	}
}

func (p *GramaticaParser) Oprel() (localctx IOprelContext) {
	this := p
	_ = this

	localctx = NewOprelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_oprel)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)

			var _m = p.Match(GramaticaParserT__14)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)

			var _m = p.Match(GramaticaParserT__15)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)

			var _m = p.Match(GramaticaParserT__16)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__17:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)

			var _m = p.Match(GramaticaParserT__17)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__18:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)

			var _m = p.Match(GramaticaParserT__18)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__19:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(130)

			var _m = p.Match(GramaticaParserT__19)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILrefContext is an interface to support dynamic dispatch.
type ILrefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId_ returns the id_ token.
	GetId_() antlr.Token

	// SetId_ sets the id_ token.
	SetId_(antlr.Token)

	// GetRef returns the ref rule contexts.
	GetRef() ILrefContext

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetRef sets the ref rule contexts.
	SetRef(ILrefContext)

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetId returns the id attribute.
	GetId() string

	// GetAux returns the aux attribute.
	GetAux() string

	// GetDir returns the dir attribute.
	GetDir() string

	// GetDim_ returns the dim_ attribute.
	GetDim_() int

	// SetId sets the id attribute.
	SetId(string)

	// SetAux sets the aux attribute.
	SetAux(string)

	// SetDir sets the dir attribute.
	SetDir(string)

	// SetDim_ sets the dim_ attribute.
	SetDim_(int)

	// IsLrefContext differentiates from other interfaces.
	IsLrefContext()
}

type LrefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     string
	aux    string
	dir    string
	dim_   int
	ref    ILrefContext
	id_    antlr.Token
	e      IExpresionContext
}

func NewEmptyLrefContext() *LrefContext {
	var p = new(LrefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_lref
	return p
}

func (*LrefContext) IsLrefContext() {}

func NewLrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LrefContext {
	var p = new(LrefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_lref

	return p
}

func (s *LrefContext) GetParser() antlr.Parser { return s.parser }

func (s *LrefContext) GetId_() antlr.Token { return s.id_ }

func (s *LrefContext) SetId_(v antlr.Token) { s.id_ = v }

func (s *LrefContext) GetRef() ILrefContext { return s.ref }

func (s *LrefContext) GetE() IExpresionContext { return s.e }

func (s *LrefContext) SetRef(v ILrefContext) { s.ref = v }

func (s *LrefContext) SetE(v IExpresionContext) { s.e = v }

func (s *LrefContext) GetId() string { return s.id }

func (s *LrefContext) GetAux() string { return s.aux }

func (s *LrefContext) GetDir() string { return s.dir }

func (s *LrefContext) GetDim_() int { return s.dim_ }

func (s *LrefContext) SetId(v string) { s.id = v }

func (s *LrefContext) SetAux(v string) { s.aux = v }

func (s *LrefContext) SetDir(v string) { s.dir = v }

func (s *LrefContext) SetDim_(v int) { s.dim_ = v }

func (s *LrefContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *LrefContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LrefContext) Lref() ILrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILrefContext)
}

func (s *LrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterLref(s)
	}
}

func (s *LrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitLref(s)
	}
}

func (p *GramaticaParser) Lref() (localctx ILrefContext) {
	return p.lref(0)
}

func (p *GramaticaParser) lref(_p int) (localctx ILrefContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLrefContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILrefContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, GramaticaParserRULE_lref, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(135)

		var _m = p.Match(GramaticaParserID)

		localctx.(*LrefContext).id_ = _m
	}
	{
		p.SetState(136)
		p.Match(GramaticaParserT__21)
	}
	{
		p.SetState(137)

		var _x = p.expresion(0)

		localctx.(*LrefContext).e = _x
	}

	// orden de columnas
	// localctx.(*LrefContext).id = (func() string { if localctx.(*LrefContext).GetId_() == nil { return "" } else { return localctx.(*LrefContext).GetId_().GetText() }}())
	// localctx.(*LrefContext).aux = gen.NewTemp()
	// localctx.(*LrefContext).dim_ = 1
	// gen.AddAsign(localctx.(*LrefContext).aux, localctx.(*LrefContext).GetE().GetDir())

	// orden de filas
	localctx.(*LrefContext).id = (func() string {
		if localctx.(*LrefContext).GetId_() == nil {
			return ""
		} else {
			return localctx.(*LrefContext).GetId_().GetText()
		}
	}())
	localctx.(*LrefContext).aux = gen.NewTemp()
	localctx.(*LrefContext).dir = localctx.(*LrefContext).GetE().GetDir()
	localctx.(*LrefContext).dim_ = 1
	gen.AddAsign(localctx.(*LrefContext).aux, "0")

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLrefContext(p, _parentctx, _parentState)
			localctx.(*LrefContext).ref = _prevctx
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_lref)
			p.SetState(140)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(141)
				p.Match(GramaticaParserT__20)
			}
			{
				p.SetState(142)

				var _x = p.expresion(0)

				localctx.(*LrefContext).e = _x
			}

			// orden en columnas
			// tmp1 := gen.NewTemp()
			// gen.AddExpresion(tmp1, localctx.(*LrefContext).GetE().GetDir(), "-", "1")
			// tmp2 := gen.NewTemp()
			// gen.AddExpresion(tmp2, tmp1, "*", entorno.GetTamOrdenColumnas(localctx.(*LrefContext).GetRef().GetId(), localctx.(*LrefContext).GetRef().GetDim_(), tope))
			// gen.AddExpresion(localctx.(*LrefContext).GetRef().GetAux(), localctx.(*LrefContext).GetRef().GetAux(), "+", tmp2)

			// orden en filas
			tmp1 := gen.NewTemp()
			gen.AddExpresion(tmp1, localctx.(*LrefContext).GetRef().GetDir(), "-", "1")
			tmp2 := gen.NewTemp()
			gen.AddExpresion(tmp2, tmp1, "*", entorno.GetTamOrdenFilas(localctx.(*LrefContext).GetRef().GetId(), localctx.(*LrefContext).GetRef().GetDim_(), tope))
			gen.AddExpresion(localctx.(*LrefContext).GetRef().GetAux(), localctx.(*LrefContext).GetRef().GetAux(), "+", tmp2)

			localctx.(*LrefContext).id = localctx.(*LrefContext).GetRef().GetId()
			localctx.(*LrefContext).aux = localctx.(*LrefContext).GetRef().GetAux()
			localctx.(*LrefContext).dir = localctx.(*LrefContext).GetE().GetDir()
			localctx.(*LrefContext).dim_ = localctx.(*LrefContext).GetRef().GetDim_() + 1

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *GramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramaticaParserRULE_instrucciones)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-23)&-(0x1f+1)) == 0 && ((1<<uint((_la-23)))&((1<<(GramaticaParserT__22-23))|(1<<(GramaticaParserT__30-23))|(1<<(GramaticaParserT__32-23))|(1<<(GramaticaParserT__37-23))|(1<<(GramaticaParserT__38-23))|(1<<(GramaticaParserT__39-23))|(1<<(GramaticaParserT__40-23))|(1<<(GramaticaParserT__42-23))|(1<<(GramaticaParserT__43-23))|(1<<(GramaticaParserID-23)))) != 0 {
		{
			p.SetState(150)
			p.Instruccion()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Inst_declaracion() IInst_declaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_declaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_declaracionContext)
}

func (s *InstruccionContext) Inst_asignacion() IInst_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_asignacionContext)
}

func (s *InstruccionContext) Inst_if() IInst_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_ifContext)
}

func (s *InstruccionContext) Inst_switch() IInst_switchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_switchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_switchContext)
}

func (s *InstruccionContext) Inst_while() IInst_whileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_whileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_whileContext)
}

func (s *InstruccionContext) Inst_doWhile() IInst_doWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_doWhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_doWhileContext)
}

func (s *InstruccionContext) Inst_loop() IInst_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_loopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_loopContext)
}

func (s *InstruccionContext) Inst_for() IInst_forContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_forContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_forContext)
}

func (s *InstruccionContext) Inst_break() IInst_breakContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_breakContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_breakContext)
}

func (s *InstruccionContext) Inst_continue() IInst_continueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_continueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_continueContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *GramaticaParser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramaticaParserRULE_instruccion)

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

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Inst_declaracion()
		}

	case GramaticaParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Inst_asignacion()
		}

	case GramaticaParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Inst_if()
		}

	case GramaticaParserT__32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
			p.Inst_switch()
		}

	case GramaticaParserT__37:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)
			p.Inst_while()
		}

	case GramaticaParserT__38:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(161)
			p.Inst_doWhile()
		}

	case GramaticaParserT__39:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(162)
			p.Inst_loop()
		}

	case GramaticaParserT__40:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(163)
			p.Inst_for()
		}

	case GramaticaParserT__42:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(164)
			p.Inst_break()
		}

	case GramaticaParserT__43:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(165)
			p.Inst_continue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInst_declaracionContext is an interface to support dynamic dispatch.
type IInst_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITipoContext

	// GetDims returns the dims rule contexts.
	GetDims() ILdimsContext

	// SetT sets the t rule contexts.
	SetT(ITipoContext)

	// SetDims sets the dims rule contexts.
	SetDims(ILdimsContext)

	// IsInst_declaracionContext differentiates from other interfaces.
	IsInst_declaracionContext()
}

type Inst_declaracionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      ITipoContext
	dims   ILdimsContext
}

func NewEmptyInst_declaracionContext() *Inst_declaracionContext {
	var p = new(Inst_declaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_declaracion
	return p
}

func (*Inst_declaracionContext) IsInst_declaracionContext() {}

func NewInst_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_declaracionContext {
	var p = new(Inst_declaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_declaracion

	return p
}

func (s *Inst_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_declaracionContext) GetId() antlr.Token { return s.id }

func (s *Inst_declaracionContext) SetId(v antlr.Token) { s.id = v }

func (s *Inst_declaracionContext) GetT() ITipoContext { return s.t }

func (s *Inst_declaracionContext) GetDims() ILdimsContext { return s.dims }

func (s *Inst_declaracionContext) SetT(v ITipoContext) { s.t = v }

func (s *Inst_declaracionContext) SetDims(v ILdimsContext) { s.dims = v }

func (s *Inst_declaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Inst_declaracionContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Inst_declaracionContext) Ldims() ILdimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILdimsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILdimsContext)
}

func (s *Inst_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_declaracion(s)
	}
}

func (s *Inst_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_declaracion(s)
	}
}

func (p *GramaticaParser) Inst_declaracion() (localctx IInst_declaracionContext) {
	this := p
	_ = this

	localctx = NewInst_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramaticaParserRULE_inst_declaracion)

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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(GramaticaParserT__22)
		}
		{
			p.SetState(169)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Inst_declaracionContext).id = _m
		}
		{
			p.SetState(170)
			p.Match(GramaticaParserT__23)
		}
		{
			p.SetState(171)

			var _x = p.Tipo()

			localctx.(*Inst_declaracionContext).t = _x
		}
		{
			p.SetState(172)
			p.Match(GramaticaParserT__24)
		}

		s := entorno.NewVar((func() string {
			if localctx.(*Inst_declaracionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*Inst_declaracionContext).GetId().GetText()
			}
		}()), localctx.(*Inst_declaracionContext).GetT().GetCad(), desp)
		desp = desp + 1
		tope.Put((func() string {
			if localctx.(*Inst_declaracionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*Inst_declaracionContext).GetId().GetText()
			}
		}()), s)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(GramaticaParserT__22)
		}
		{
			p.SetState(176)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Inst_declaracionContext).id = _m
		}
		{
			p.SetState(177)
			p.Match(GramaticaParserT__23)
		}
		{
			p.SetState(178)
			p.Match(GramaticaParserT__25)
		}
		{
			p.SetState(179)
			p.Match(GramaticaParserT__21)
		}
		{
			p.SetState(180)

			var _x = p.Ldims()

			localctx.(*Inst_declaracionContext).dims = _x
		}
		{
			p.SetState(181)
			p.Match(GramaticaParserT__13)
		}
		{
			p.SetState(182)
			p.Match(GramaticaParserT__26)
		}
		{
			p.SetState(183)

			var _x = p.Tipo()

			localctx.(*Inst_declaracionContext).t = _x
		}
		{
			p.SetState(184)
			p.Match(GramaticaParserT__24)
		}

		s := entorno.NewVarArray((func() string {
			if localctx.(*Inst_declaracionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*Inst_declaracionContext).GetId().GetText()
			}
		}()), localctx.(*Inst_declaracionContext).GetT().GetCad(), desp, 0, localctx.(*Inst_declaracionContext).GetDims().GetDim_())
		desp = desp + 1
		tope.Put((func() string {
			if localctx.(*Inst_declaracionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*Inst_declaracionContext).GetId().GetText()
			}
		}()), s)

	}

	return localctx
}

// ILdimsContext is an interface to support dynamic dispatch.
type ILdimsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN1 returns the n1 token.
	GetN1() antlr.Token

	// GetN2 returns the n2 token.
	GetN2() antlr.Token

	// GetN3 returns the n3 token.
	GetN3() antlr.Token

	// GetN4 returns the n4 token.
	GetN4() antlr.Token

	// SetN1 sets the n1 token.
	SetN1(antlr.Token)

	// SetN2 sets the n2 token.
	SetN2(antlr.Token)

	// SetN3 sets the n3 token.
	SetN3(antlr.Token)

	// SetN4 sets the n4 token.
	SetN4(antlr.Token)

	// GetDim_ returns the dim_ attribute.
	GetDim_() []*entorno.Dimension

	// SetDim_ sets the dim_ attribute.
	SetDim_([]*entorno.Dimension)

	// IsLdimsContext differentiates from other interfaces.
	IsLdimsContext()
}

type LdimsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dim_   []*entorno.Dimension
	n1     antlr.Token
	n2     antlr.Token
	n3     antlr.Token
	n4     antlr.Token
}

func NewEmptyLdimsContext() *LdimsContext {
	var p = new(LdimsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_ldims
	return p
}

func (*LdimsContext) IsLdimsContext() {}

func NewLdimsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LdimsContext {
	var p = new(LdimsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_ldims

	return p
}

func (s *LdimsContext) GetParser() antlr.Parser { return s.parser }

func (s *LdimsContext) GetN1() antlr.Token { return s.n1 }

func (s *LdimsContext) GetN2() antlr.Token { return s.n2 }

func (s *LdimsContext) GetN3() antlr.Token { return s.n3 }

func (s *LdimsContext) GetN4() antlr.Token { return s.n4 }

func (s *LdimsContext) SetN1(v antlr.Token) { s.n1 = v }

func (s *LdimsContext) SetN2(v antlr.Token) { s.n2 = v }

func (s *LdimsContext) SetN3(v antlr.Token) { s.n3 = v }

func (s *LdimsContext) SetN4(v antlr.Token) { s.n4 = v }

func (s *LdimsContext) GetDim_() []*entorno.Dimension { return s.dim_ }

func (s *LdimsContext) SetDim_(v []*entorno.Dimension) { s.dim_ = v }

func (s *LdimsContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserNUM)
}

func (s *LdimsContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, i)
}

func (s *LdimsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LdimsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LdimsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterLdims(s)
	}
}

func (s *LdimsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitLdims(s)
	}
}

func (p *GramaticaParser) Ldims() (localctx ILdimsContext) {
	this := p
	_ = this

	localctx = NewLdimsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramaticaParserRULE_ldims)
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
		p.SetState(189)

		var _m = p.Match(GramaticaParserNUM)

		localctx.(*LdimsContext).n1 = _m
	}
	{
		p.SetState(190)
		p.Match(GramaticaParserT__27)
	}
	{
		p.SetState(191)

		var _m = p.Match(GramaticaParserNUM)

		localctx.(*LdimsContext).n2 = _m
	}

	d := entorno.NewDimension((func() int {
		if localctx.(*LdimsContext).GetN1() == nil {
			return 0
		} else {
			i, _ := strconv.Atoi(localctx.(*LdimsContext).GetN1().GetText())
			return i
		}
	}()), (func() int {
		if localctx.(*LdimsContext).GetN2() == nil {
			return 0
		} else {
			i, _ := strconv.Atoi(localctx.(*LdimsContext).GetN2().GetText())
			return i
		}
	}()))
	localctx.(*LdimsContext).dim_ = append(localctx.(*LdimsContext).dim_, d)

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__20 {
		{
			p.SetState(193)
			p.Match(GramaticaParserT__20)
		}
		{
			p.SetState(194)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*LdimsContext).n3 = _m
		}
		{
			p.SetState(195)
			p.Match(GramaticaParserT__27)
		}
		{
			p.SetState(196)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*LdimsContext).n4 = _m
		}

		d := entorno.NewDimension((func() int {
			if localctx.(*LdimsContext).GetN3() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*LdimsContext).GetN3().GetText())
				return i
			}
		}()), (func() int {
			if localctx.(*LdimsContext).GetN4() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*LdimsContext).GetN4().GetText())
				return i
			}
		}()))
		localctx.(*LdimsContext).dim_ = append(localctx.(*LdimsContext).dim_, d)

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCad returns the cad attribute.
	GetCad() string

	// SetCad sets the cad attribute.
	SetCad(string)

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cad    string
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_tipo
	return p
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) GetCad() string { return s.cad }

func (s *TipoContext) SetCad(v string) { s.cad = v }

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *GramaticaParser) Tipo() (localctx ITipoContext) {
	this := p
	_ = this

	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_tipo)

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
		p.SetState(203)
		p.Match(GramaticaParserT__28)
	}
	localctx.(*TipoContext).cad = "int"

	return localctx
}

// IInst_asignacionContext is an interface to support dynamic dispatch.
type IInst_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// GetRef returns the ref rule contexts.
	GetRef() ILrefContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// SetRef sets the ref rule contexts.
	SetRef(ILrefContext)

	// IsInst_asignacionContext differentiates from other interfaces.
	IsInst_asignacionContext()
}

type Inst_asignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	e      IExpresionContext
	ref    ILrefContext
}

func NewEmptyInst_asignacionContext() *Inst_asignacionContext {
	var p = new(Inst_asignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_asignacion
	return p
}

func (*Inst_asignacionContext) IsInst_asignacionContext() {}

func NewInst_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_asignacionContext {
	var p = new(Inst_asignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_asignacion

	return p
}

func (s *Inst_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_asignacionContext) GetId() antlr.Token { return s.id }

func (s *Inst_asignacionContext) SetId(v antlr.Token) { s.id = v }

func (s *Inst_asignacionContext) GetE() IExpresionContext { return s.e }

func (s *Inst_asignacionContext) GetRef() ILrefContext { return s.ref }

func (s *Inst_asignacionContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_asignacionContext) SetRef(v ILrefContext) { s.ref = v }

func (s *Inst_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Inst_asignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_asignacionContext) Lref() ILrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILrefContext)
}

func (s *Inst_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_asignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_asignacion(s)
	}
}

func (s *Inst_asignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_asignacion(s)
	}
}

func (p *GramaticaParser) Inst_asignacion() (localctx IInst_asignacionContext) {
	this := p
	_ = this

	localctx = NewInst_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramaticaParserRULE_inst_asignacion)

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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Inst_asignacionContext).id = _m
		}
		{
			p.SetState(207)
			p.Match(GramaticaParserT__29)
		}
		{
			p.SetState(208)

			var _x = p.expresion(0)

			localctx.(*Inst_asignacionContext).e = _x
		}
		{
			p.SetState(209)
			p.Match(GramaticaParserT__24)
		}
		gen.AddAsign((func() string {
			if localctx.(*Inst_asignacionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*Inst_asignacionContext).GetId().GetText()
			}
		}()), localctx.(*Inst_asignacionContext).GetE().GetDir())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)

			var _x = p.lref(0)

			localctx.(*Inst_asignacionContext).ref = _x
		}
		{
			p.SetState(213)
			p.Match(GramaticaParserT__13)
		}
		{
			p.SetState(214)
			p.Match(GramaticaParserT__29)
		}
		{
			p.SetState(215)

			var _x = p.expresion(0)

			localctx.(*Inst_asignacionContext).e = _x
		}
		{
			p.SetState(216)
			p.Match(GramaticaParserT__24)
		}

		// orden de columnas
		// gen.AddSetArray(localctx.(*Inst_asignacionContext).GetRef().GetId(), localctx.(*Inst_asignacionContext).GetRef().GetAux(), localctx.(*Inst_asignacionContext).GetE().GetDir())

		// orden de filas
		tmp := gen.NewTemp()
		gen.AddExpresion(tmp, localctx.(*Inst_asignacionContext).GetRef().GetAux(), "+", localctx.(*Inst_asignacionContext).GetRef().GetDir())
		gen.AddSetArray(localctx.(*Inst_asignacionContext).GetRef().GetId(), tmp, localctx.(*Inst_asignacionContext).GetE().GetDir())

	}

	return localctx
}

// IInst_ifContext is an interface to support dynamic dispatch.
type IInst_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// IsInst_ifContext differentiates from other interfaces.
	IsInst_ifContext()
}

type Inst_ifContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lsalida string
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_ifContext() *Inst_ifContext {
	var p = new(Inst_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_if
	return p
}

func (*Inst_ifContext) IsInst_ifContext() {}

func NewInst_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_ifContext {
	var p = new(Inst_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_if

	return p
}

func (s *Inst_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_ifContext) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_ifContext) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_ifContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_ifContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_ifContext) GetLsalida() string { return s.lsalida }

func (s *Inst_ifContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_ifContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_ifContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_ifContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_ifContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_if(s)
	}
}

func (s *Inst_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_if(s)
	}
}

func (p *GramaticaParser) Inst_if() (localctx IInst_ifContext) {
	this := p
	_ = this

	localctx = NewInst_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramaticaParserRULE_inst_if)
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
		p.SetState(221)
		p.Match(GramaticaParserT__30)
	}
	{
		p.SetState(222)

		var _x = p.expresion(0)

		localctx.(*Inst_ifContext).e1 = _x
	}

	localctx.(*Inst_ifContext).lsalida = gen.NewLabel()
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLv())

	{
		p.SetState(224)
		p.Bloque()
	}

	gen.AddGoto(localctx.(*Inst_ifContext).lsalida)
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLf())

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(226)
				p.Match(GramaticaParserT__31)
			}
			{
				p.SetState(227)
				p.Match(GramaticaParserT__30)
			}
			{
				p.SetState(228)

				var _x = p.expresion(0)

				localctx.(*Inst_ifContext).e2 = _x
			}

			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLv())

			{
				p.SetState(230)
				p.Bloque()
			}

			gen.AddGoto(localctx.(*Inst_ifContext).lsalida)
			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLf())

		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__31 {
		{
			p.SetState(238)
			p.Match(GramaticaParserT__31)
		}
		{
			p.SetState(239)
			p.Bloque()
		}

	}
	gen.AddLabel(localctx.(*Inst_ifContext).lsalida)

	return localctx
}

// IInst_switch_propuesta1Context is an interface to support dynamic dispatch.
type IInst_switch_propuesta1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// GetLv returns the lv attribute.
	GetLv() string

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// SetLv sets the lv attribute.
	SetLv(string)

	// IsInst_switch_propuesta1Context differentiates from other interfaces.
	IsInst_switch_propuesta1Context()
}

type Inst_switch_propuesta1Context struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lsalida string
	lv      string
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_switch_propuesta1Context() *Inst_switch_propuesta1Context {
	var p = new(Inst_switch_propuesta1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta1
	return p
}

func (*Inst_switch_propuesta1Context) IsInst_switch_propuesta1Context() {}

func NewInst_switch_propuesta1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_switch_propuesta1Context {
	var p = new(Inst_switch_propuesta1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta1

	return p
}

func (s *Inst_switch_propuesta1Context) GetParser() antlr.Parser { return s.parser }

func (s *Inst_switch_propuesta1Context) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_switch_propuesta1Context) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_switch_propuesta1Context) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_switch_propuesta1Context) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_switch_propuesta1Context) GetLsalida() string { return s.lsalida }

func (s *Inst_switch_propuesta1Context) GetLv() string { return s.lv }

func (s *Inst_switch_propuesta1Context) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_switch_propuesta1Context) SetLv(v string) { s.lv = v }

func (s *Inst_switch_propuesta1Context) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_switch_propuesta1Context) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_switch_propuesta1Context) AllBloqueSinLlaves() []IBloqueSinLlavesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueSinLlavesContext)(nil)).Elem())
	var tst = make([]IBloqueSinLlavesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueSinLlavesContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) BloqueSinLlaves(i int) IBloqueSinLlavesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueSinLlavesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueSinLlavesContext)
}

func (s *Inst_switch_propuesta1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_switch_propuesta1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_switch_propuesta1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_switch_propuesta1(s)
	}
}

func (s *Inst_switch_propuesta1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_switch_propuesta1(s)
	}
}

func (p *GramaticaParser) Inst_switch_propuesta1() (localctx IInst_switch_propuesta1Context) {
	this := p
	_ = this

	localctx = NewInst_switch_propuesta1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GramaticaParserRULE_inst_switch_propuesta1)

	localctx.(*Inst_switch_propuesta1Context).lsalida = gen.NewLabel()
	gen.NextPtr()
	gen.Display[gen.Ptr].Salida = localctx.(*Inst_switch_propuesta1Context).lsalida

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
		p.Match(GramaticaParserT__32)
	}
	{
		p.SetState(245)

		var _x = p.expresion(0)

		localctx.(*Inst_switch_propuesta1Context).e1 = _x
	}
	{
		p.SetState(246)
		p.Match(GramaticaParserT__33)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserT__34 {
		{
			p.SetState(247)
			p.Match(GramaticaParserT__34)
		}
		{
			p.SetState(248)

			var _x = p.expresion(0)

			localctx.(*Inst_switch_propuesta1Context).e2 = _x
		}
		{
			p.SetState(249)
			p.Match(GramaticaParserT__23)
		}

		localctx.(*Inst_switch_propuesta1Context).lv = gen.NewLabel()
		gen.AddIf(localctx.(*Inst_switch_propuesta1Context).GetE1().GetDir(), "!=", localctx.(*Inst_switch_propuesta1Context).GetE2().GetDir(), localctx.(*Inst_switch_propuesta1Context).lv)

		p.SetState(253)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(251)
				p.Bloque()
			}

		case GramaticaParserT__22, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__34, GramaticaParserT__35, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserT__40, GramaticaParserT__42, GramaticaParserT__43, GramaticaParserID:
			{
				p.SetState(252)
				p.BloqueSinLlaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.AddGoto(localctx.(*Inst_switch_propuesta1Context).lsalida)
		gen.AddLabel(localctx.(*Inst_switch_propuesta1Context).lv)

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__35 {
		{
			p.SetState(261)
			p.Match(GramaticaParserT__35)
		}
		{
			p.SetState(262)
			p.Match(GramaticaParserT__23)
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(263)
				p.Bloque()
			}

		case GramaticaParserT__22, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserT__40, GramaticaParserT__42, GramaticaParserT__43, GramaticaParserID:
			{
				p.SetState(264)
				p.BloqueSinLlaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	gen.AddLabel(localctx.(*Inst_switch_propuesta1Context).lsalida)
	gen.PrevPtr()

	{
		p.SetState(270)
		p.Match(GramaticaParserT__36)
	}

	return localctx
}

// IInst_switchContext is an interface to support dynamic dispatch.
type IInst_switchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// GetLprueba returns the lprueba attribute.
	GetLprueba() string

	// GetLv returns the lv attribute.
	GetLv() string

	// GetCad returns the cad attribute.
	GetCad() string

	// GetDefecto returns the defecto attribute.
	GetDefecto() bool

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// SetLprueba sets the lprueba attribute.
	SetLprueba(string)

	// SetLv sets the lv attribute.
	SetLv(string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// SetDefecto sets the defecto attribute.
	SetDefecto(bool)

	// IsInst_switchContext differentiates from other interfaces.
	IsInst_switchContext()
}

type Inst_switchContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lsalida string
	lprueba string
	lv      string
	cad     string
	defecto bool
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_switchContext() *Inst_switchContext {
	var p = new(Inst_switchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_switch
	return p
}

func (*Inst_switchContext) IsInst_switchContext() {}

func NewInst_switchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_switchContext {
	var p = new(Inst_switchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_switch

	return p
}

func (s *Inst_switchContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_switchContext) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_switchContext) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_switchContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_switchContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_switchContext) GetLsalida() string { return s.lsalida }

func (s *Inst_switchContext) GetLprueba() string { return s.lprueba }

func (s *Inst_switchContext) GetLv() string { return s.lv }

func (s *Inst_switchContext) GetCad() string { return s.cad }

func (s *Inst_switchContext) GetDefecto() bool { return s.defecto }

func (s *Inst_switchContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_switchContext) SetLprueba(v string) { s.lprueba = v }

func (s *Inst_switchContext) SetLv(v string) { s.lv = v }

func (s *Inst_switchContext) SetCad(v string) { s.cad = v }

func (s *Inst_switchContext) SetDefecto(v bool) { s.defecto = v }

func (s *Inst_switchContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_switchContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_switchContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_switchContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_switchContext) AllBloqueSinLlaves() []IBloqueSinLlavesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueSinLlavesContext)(nil)).Elem())
	var tst = make([]IBloqueSinLlavesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueSinLlavesContext)
		}
	}

	return tst
}

func (s *Inst_switchContext) BloqueSinLlaves(i int) IBloqueSinLlavesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueSinLlavesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueSinLlavesContext)
}

func (s *Inst_switchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_switchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_switchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_switch(s)
	}
}

func (s *Inst_switchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_switch(s)
	}
}

func (p *GramaticaParser) Inst_switch() (localctx IInst_switchContext) {
	this := p
	_ = this

	localctx = NewInst_switchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GramaticaParserRULE_inst_switch)

	localctx.(*Inst_switchContext).lprueba = gen.NewLabel()
	localctx.(*Inst_switchContext).lsalida = gen.NewLabel()
	localctx.(*Inst_switchContext).cad = ""
	localctx.(*Inst_switchContext).defecto = false
	gen.NextPtr()
	gen.Display[gen.Ptr].Salida = localctx.(*Inst_switchContext).lsalida

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
		p.SetState(272)
		p.Match(GramaticaParserT__32)
	}
	{
		p.SetState(273)

		var _x = p.expresion(0)

		localctx.(*Inst_switchContext).e1 = _x
	}
	{
		p.SetState(274)
		p.Match(GramaticaParserT__33)
	}

	gen.AddGoto(localctx.(*Inst_switchContext).lprueba)

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserT__34 {
		{
			p.SetState(276)
			p.Match(GramaticaParserT__34)
		}
		{
			p.SetState(277)

			var _x = p.expresion(0)

			localctx.(*Inst_switchContext).e2 = _x
		}
		{
			p.SetState(278)
			p.Match(GramaticaParserT__23)
		}

		localctx.(*Inst_switchContext).lv = gen.NewLabel()
		gen.AddLabel(localctx.(*Inst_switchContext).lv)
		localctx.(*Inst_switchContext).cad += "if " + localctx.(*Inst_switchContext).GetE1().GetDir() + " = " + localctx.(*Inst_switchContext).GetE2().GetDir() + " then goto " + localctx.(*Inst_switchContext).lv + "\n"

		p.SetState(282)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(280)
				p.Bloque()
			}

		case GramaticaParserT__22, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__34, GramaticaParserT__35, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserT__40, GramaticaParserT__42, GramaticaParserT__43, GramaticaParserID:
			{
				p.SetState(281)
				p.BloqueSinLlaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.AddGoto(localctx.(*Inst_switchContext).lsalida)

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__35 {
		{
			p.SetState(290)
			p.Match(GramaticaParserT__35)
		}
		{
			p.SetState(291)
			p.Match(GramaticaParserT__23)
		}

		localctx.(*Inst_switchContext).lv = gen.NewLabel()
		localctx.(*Inst_switchContext).defecto = true
		gen.AddLabel(localctx.(*Inst_switchContext).lv)

		p.SetState(295)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(293)
				p.Bloque()
			}

		case GramaticaParserT__22, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserT__40, GramaticaParserT__42, GramaticaParserT__43, GramaticaParserID:
			{
				p.SetState(294)
				p.BloqueSinLlaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.AddGoto(localctx.(*Inst_switchContext).lsalida)

	}

	gen.AddGoto(localctx.(*Inst_switchContext).lsalida)
	gen.AddLabel(localctx.(*Inst_switchContext).lprueba)
	gen.Gen(localctx.(*Inst_switchContext).cad)
	if localctx.(*Inst_switchContext).defecto {
		gen.AddGoto(localctx.(*Inst_switchContext).lv)
	}
	gen.AddLabel(localctx.(*Inst_switchContext).lsalida)
	gen.PrevPtr()

	{
		p.SetState(302)
		p.Match(GramaticaParserT__36)
	}

	return localctx
}

// IInst_whileContext is an interface to support dynamic dispatch.
type IInst_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// IsInst_whileContext differentiates from other interfaces.
	IsInst_whileContext()
}

type Inst_whileContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	e       IExpresionContext
}

func NewEmptyInst_whileContext() *Inst_whileContext {
	var p = new(Inst_whileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_while
	return p
}

func (*Inst_whileContext) IsInst_whileContext() {}

func NewInst_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_whileContext {
	var p = new(Inst_whileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_while

	return p
}

func (s *Inst_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_whileContext) GetE() IExpresionContext { return s.e }

func (s *Inst_whileContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_whileContext) GetLinicio() string { return s.linicio }

func (s *Inst_whileContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_whileContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_whileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_while(s)
	}
}

func (s *Inst_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_while(s)
	}
}

func (p *GramaticaParser) Inst_while() (localctx IInst_whileContext) {
	this := p
	_ = this

	localctx = NewInst_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GramaticaParserRULE_inst_while)

	localctx.(*Inst_whileContext).linicio = gen.NewLabel()

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
		p.Match(GramaticaParserT__37)
	}
	gen.AddLabel(localctx.(*Inst_whileContext).linicio)
	{
		p.SetState(306)

		var _x = p.expresion(0)

		localctx.(*Inst_whileContext).e = _x
	}

	gen.NextPtr()
	gen.Display[gen.Ptr].Salida = localctx.(*Inst_whileContext).GetE().GetLf()[0]
	gen.Display[gen.Ptr].Inicio = localctx.(*Inst_whileContext).linicio
	gen.Soltar(localctx.(*Inst_whileContext).GetE().GetLv())

	{
		p.SetState(308)
		p.Bloque()
	}

	gen.AddGoto(localctx.(*Inst_whileContext).linicio)
	gen.Soltar(localctx.(*Inst_whileContext).GetE().GetLf())
	gen.PrevPtr()

	return localctx
}

// IInst_doWhileContext is an interface to support dynamic dispatch.
type IInst_doWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// IsInst_doWhileContext differentiates from other interfaces.
	IsInst_doWhileContext()
}

type Inst_doWhileContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	e       IExpresionContext
}

func NewEmptyInst_doWhileContext() *Inst_doWhileContext {
	var p = new(Inst_doWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_doWhile
	return p
}

func (*Inst_doWhileContext) IsInst_doWhileContext() {}

func NewInst_doWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_doWhileContext {
	var p = new(Inst_doWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_doWhile

	return p
}

func (s *Inst_doWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_doWhileContext) GetE() IExpresionContext { return s.e }

func (s *Inst_doWhileContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_doWhileContext) GetLinicio() string { return s.linicio }

func (s *Inst_doWhileContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_doWhileContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_doWhileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_doWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_doWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_doWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_doWhile(s)
	}
}

func (s *Inst_doWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_doWhile(s)
	}
}

func (p *GramaticaParser) Inst_doWhile() (localctx IInst_doWhileContext) {
	this := p
	_ = this

	localctx = NewInst_doWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GramaticaParserRULE_inst_doWhile)

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
		p.Match(GramaticaParserT__38)
	}

	localctx.(*Inst_doWhileContext).linicio = gen.NewLabel()
	gen.AddLabel(localctx.(*Inst_doWhileContext).linicio)
	gen.NextPtr()
	gen.Display[gen.Ptr].Salida = gen.NewLabel()
	gen.Display[gen.Ptr].Inicio = localctx.(*Inst_doWhileContext).linicio

	{
		p.SetState(313)
		p.Bloque()
	}
	{
		p.SetState(314)
		p.Match(GramaticaParserT__37)
	}
	{
		p.SetState(315)

		var _x = p.expresion(0)

		localctx.(*Inst_doWhileContext).e = _x
	}
	{
		p.SetState(316)
		p.Match(GramaticaParserT__24)
	}

	gen.Soltar(localctx.(*Inst_doWhileContext).GetE().GetLv())
	gen.AddGoto(localctx.(*Inst_doWhileContext).linicio)
	gen.Soltar(localctx.(*Inst_doWhileContext).GetE().GetLf())
	gen.AddLabel(gen.Display[gen.Ptr].Salida)
	gen.PrevPtr()

	return localctx
}

// IInst_loopContext is an interface to support dynamic dispatch.
type IInst_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// IsInst_loopContext differentiates from other interfaces.
	IsInst_loopContext()
}

type Inst_loopContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	lsalida string
}

func NewEmptyInst_loopContext() *Inst_loopContext {
	var p = new(Inst_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_loop
	return p
}

func (*Inst_loopContext) IsInst_loopContext() {}

func NewInst_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_loopContext {
	var p = new(Inst_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_loop

	return p
}

func (s *Inst_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_loopContext) GetLinicio() string { return s.linicio }

func (s *Inst_loopContext) GetLsalida() string { return s.lsalida }

func (s *Inst_loopContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_loopContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_loopContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_loop(s)
	}
}

func (s *Inst_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_loop(s)
	}
}

func (p *GramaticaParser) Inst_loop() (localctx IInst_loopContext) {
	this := p
	_ = this

	localctx = NewInst_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GramaticaParserRULE_inst_loop)

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
		p.SetState(319)
		p.Match(GramaticaParserT__39)
	}

	localctx.(*Inst_loopContext).linicio = gen.NewLabel()
	gen.AddLabel(localctx.(*Inst_loopContext).linicio)
	gen.NextPtr()
	gen.Display[gen.Ptr].Salida = gen.NewLabel()
	gen.Display[gen.Ptr].Inicio = localctx.(*Inst_loopContext).linicio
	gen.Display[gen.Ptr].ContB = 0

	{
		p.SetState(321)
		p.Bloque()
	}

	gen.AddGoto(localctx.(*Inst_loopContext).linicio)
	gen.AddLabel(gen.Display[gen.Ptr].Salida)
	if gen.Display[gen.Ptr].ContB == 0 {
		err.NewError("Error, break no utilizado")
	}
	gen.PrevPtr()

	return localctx
}

// IInst_forContext is an interface to support dynamic dispatch.
type IInst_forContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// GetTmp returns the tmp attribute.
	GetTmp() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// SetTmp sets the tmp attribute.
	SetTmp(string)

	// IsInst_forContext differentiates from other interfaces.
	IsInst_forContext()
}

type Inst_forContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	lsalida string
	tmp     string
	id      antlr.Token
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_forContext() *Inst_forContext {
	var p = new(Inst_forContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_for
	return p
}

func (*Inst_forContext) IsInst_forContext() {}

func NewInst_forContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_forContext {
	var p = new(Inst_forContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_for

	return p
}

func (s *Inst_forContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_forContext) GetId() antlr.Token { return s.id }

func (s *Inst_forContext) SetId(v antlr.Token) { s.id = v }

func (s *Inst_forContext) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_forContext) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_forContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_forContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_forContext) GetLinicio() string { return s.linicio }

func (s *Inst_forContext) GetLsalida() string { return s.lsalida }

func (s *Inst_forContext) GetTmp() string { return s.tmp }

func (s *Inst_forContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_forContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_forContext) SetTmp(v string) { s.tmp = v }

func (s *Inst_forContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_forContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Inst_forContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_forContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_forContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_for(s)
	}
}

func (s *Inst_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_for(s)
	}
}

func (p *GramaticaParser) Inst_for() (localctx IInst_forContext) {
	this := p
	_ = this

	localctx = NewInst_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GramaticaParserRULE_inst_for)

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
		p.SetState(324)
		p.Match(GramaticaParserT__40)
	}
	{
		p.SetState(325)

		var _m = p.Match(GramaticaParserID)

		localctx.(*Inst_forContext).id = _m
	}
	{
		p.SetState(326)
		p.Match(GramaticaParserT__29)
	}
	{
		p.SetState(327)

		var _x = p.expresion(0)

		localctx.(*Inst_forContext).e1 = _x
	}
	{
		p.SetState(328)
		p.Match(GramaticaParserT__41)
	}
	{
		p.SetState(329)

		var _x = p.expresion(0)

		localctx.(*Inst_forContext).e2 = _x
	}

	localctx.(*Inst_forContext).linicio = gen.NewLabel()
	localctx.(*Inst_forContext).lsalida = gen.NewLabel()

	gen.NextPtr()
	gen.Display[gen.Ptr].Salida = localctx.(*Inst_forContext).lsalida
	gen.Display[gen.Ptr].Inicio = localctx.(*Inst_forContext).linicio

	localctx.(*Inst_forContext).tmp = gen.NewTemp()

	sim := entorno.NewVar((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), "int", desp)
	desp = desp + 1
	tope.Put((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), sim)

	gen.AddAsign((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), localctx.(*Inst_forContext).GetE1().GetDir())
	gen.AddExpresionUnaria(localctx.(*Inst_forContext).tmp, "+", "1")
	gen.AddIf((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), ">", localctx.(*Inst_forContext).GetE2().GetDir(), localctx.(*Inst_forContext).linicio)
	gen.AddExpresionUnaria(localctx.(*Inst_forContext).tmp, "-", "1")

	gen.AddLabel(localctx.(*Inst_forContext).linicio)
	gen.AddIf((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), "=", localctx.(*Inst_forContext).GetE2().GetDir(), localctx.(*Inst_forContext).lsalida)

	{
		p.SetState(331)
		p.Match(GramaticaParserT__38)
	}
	{
		p.SetState(332)
		p.Bloque()
	}

	gen.AddExpresion((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), (func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), "+", localctx.(*Inst_forContext).tmp)
	gen.AddGoto(localctx.(*Inst_forContext).linicio)
	gen.AddLabel(localctx.(*Inst_forContext).lsalida)
	gen.PrevPtr()

	return localctx
}

// IInst_breakContext is an interface to support dynamic dispatch.
type IInst_breakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInst_breakContext differentiates from other interfaces.
	IsInst_breakContext()
}

type Inst_breakContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_breakContext() *Inst_breakContext {
	var p = new(Inst_breakContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_break
	return p
}

func (*Inst_breakContext) IsInst_breakContext() {}

func NewInst_breakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_breakContext {
	var p = new(Inst_breakContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_break

	return p
}

func (s *Inst_breakContext) GetParser() antlr.Parser { return s.parser }
func (s *Inst_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_breakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_break(s)
	}
}

func (s *Inst_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_break(s)
	}
}

func (p *GramaticaParser) Inst_break() (localctx IInst_breakContext) {
	this := p
	_ = this

	localctx = NewInst_breakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GramaticaParserRULE_inst_break)

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
		p.SetState(335)
		p.Match(GramaticaParserT__42)
	}
	{
		p.SetState(336)
		p.Match(GramaticaParserT__24)
	}

	if gen.Ptr == 0 {
		err.NewError("Error, break fuera de instruccion")
	} else {
		gen.AddGoto(gen.Display[gen.Ptr].Salida)
		gen.Display[gen.Ptr].ContB = gen.Display[gen.Ptr].ContB + 1
	}

	return localctx
}

// IInst_continueContext is an interface to support dynamic dispatch.
type IInst_continueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInst_continueContext differentiates from other interfaces.
	IsInst_continueContext()
}

type Inst_continueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInst_continueContext() *Inst_continueContext {
	var p = new(Inst_continueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_continue
	return p
}

func (*Inst_continueContext) IsInst_continueContext() {}

func NewInst_continueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_continueContext {
	var p = new(Inst_continueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_continue

	return p
}

func (s *Inst_continueContext) GetParser() antlr.Parser { return s.parser }
func (s *Inst_continueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_continueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_continueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_continue(s)
	}
}

func (s *Inst_continueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_continue(s)
	}
}

func (p *GramaticaParser) Inst_continue() (localctx IInst_continueContext) {
	this := p
	_ = this

	localctx = NewInst_continueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GramaticaParserRULE_inst_continue)

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
		p.Match(GramaticaParserT__43)
	}
	{
		p.SetState(340)
		p.Match(GramaticaParserT__24)
	}

	if gen.Ptr == 0 {
		err.NewError("Error, continue fuera de instruccion")
	} else {
		gen.AddGoto(gen.Display[gen.Ptr].Inicio)
	}

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *GramaticaParser) Bloque() (localctx IBloqueContext) {
	this := p
	_ = this

	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GramaticaParserRULE_bloque)

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

	entorno.Push(tope)
	tope = entorno.NewEntorno(tope)

	{
		p.SetState(344)
		p.Match(GramaticaParserT__33)
	}
	{
		p.SetState(345)
		p.Instrucciones()
	}
	{
		p.SetState(346)
		p.Match(GramaticaParserT__36)
	}

	tope = entorno.Pop()

	return localctx
}

// IBloqueSinLlavesContext is an interface to support dynamic dispatch.
type IBloqueSinLlavesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueSinLlavesContext differentiates from other interfaces.
	IsBloqueSinLlavesContext()
}

type BloqueSinLlavesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueSinLlavesContext() *BloqueSinLlavesContext {
	var p = new(BloqueSinLlavesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_bloqueSinLlaves
	return p
}

func (*BloqueSinLlavesContext) IsBloqueSinLlavesContext() {}

func NewBloqueSinLlavesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueSinLlavesContext {
	var p = new(BloqueSinLlavesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_bloqueSinLlaves

	return p
}

func (s *BloqueSinLlavesContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueSinLlavesContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueSinLlavesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueSinLlavesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueSinLlavesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBloqueSinLlaves(s)
	}
}

func (s *BloqueSinLlavesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBloqueSinLlaves(s)
	}
}

func (p *GramaticaParser) BloqueSinLlaves() (localctx IBloqueSinLlavesContext) {
	this := p
	_ = this

	localctx = NewBloqueSinLlavesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GramaticaParserRULE_bloqueSinLlaves)

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

	entorno.Push(tope)
	tope = entorno.NewEntorno(tope)

	{
		p.SetState(350)
		p.Instrucciones()
	}

	tope = entorno.Pop()

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	case 4:
		var t *LrefContext = nil
		if localctx != nil {
			t = localctx.(*LrefContext)
		}
		return p.Lref_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Lref_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
