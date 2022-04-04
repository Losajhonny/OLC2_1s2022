// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Compilador/gen"
import "Compilador/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 328,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 77, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 112, 10, 4, 12, 4, 14, 4, 115, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 129, 10, 5, 3, 6, 7, 6, 132, 10, 6, 12, 6, 14, 6, 135, 11, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 144, 10, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 165, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 176, 10, 9, 12, 9, 14, 9, 179, 11,
	9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 197, 10, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12,
	210, 10, 12, 12, 12, 14, 12, 213, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 227, 10, 13,
	12, 13, 14, 13, 230, 11, 13, 3, 13, 3, 13, 5, 13, 234, 10, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	247, 10, 14, 3, 14, 3, 14, 6, 14, 251, 10, 14, 13, 14, 14, 14, 252, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 259, 10, 14, 5, 14, 261, 10, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 276, 10, 15, 3, 15, 3, 15, 6, 15, 280, 10, 15, 13, 15,
	14, 15, 281, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 289, 10, 15, 3,
	15, 3, 15, 5, 15, 293, 10, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 2, 4, 6, 22, 21, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 4,
	3, 2, 6, 8, 3, 2, 3, 4, 2, 348, 2, 40, 3, 2, 2, 2, 4, 44, 3, 2, 2, 2, 6,
	76, 3, 2, 2, 2, 8, 128, 3, 2, 2, 2, 10, 133, 3, 2, 2, 2, 12, 143, 3, 2,
	2, 2, 14, 164, 3, 2, 2, 2, 16, 166, 3, 2, 2, 2, 18, 180, 3, 2, 2, 2, 20,
	196, 3, 2, 2, 2, 22, 198, 3, 2, 2, 2, 24, 214, 3, 2, 2, 2, 26, 237, 3,
	2, 2, 2, 28, 265, 3, 2, 2, 2, 30, 297, 3, 2, 2, 2, 32, 304, 3, 2, 2, 2,
	34, 312, 3, 2, 2, 2, 36, 317, 3, 2, 2, 2, 38, 323, 3, 2, 2, 2, 40, 41,
	8, 2, 1, 2, 41, 42, 5, 10, 6, 2, 42, 43, 7, 2, 2, 3, 43, 3, 3, 2, 2, 2,
	44, 45, 8, 3, 1, 2, 45, 5, 3, 2, 2, 2, 46, 47, 8, 4, 1, 2, 47, 48, 7, 3,
	2, 2, 48, 49, 5, 6, 4, 17, 49, 50, 8, 4, 1, 2, 50, 77, 3, 2, 2, 2, 51,
	52, 7, 4, 2, 2, 52, 53, 5, 6, 4, 16, 53, 54, 8, 4, 1, 2, 54, 77, 3, 2,
	2, 2, 55, 56, 7, 5, 2, 2, 56, 57, 5, 6, 4, 15, 57, 58, 8, 4, 1, 2, 58,
	77, 3, 2, 2, 2, 59, 60, 7, 12, 2, 2, 60, 61, 5, 6, 4, 2, 61, 62, 7, 13,
	2, 2, 62, 63, 8, 4, 1, 2, 63, 77, 3, 2, 2, 2, 64, 65, 7, 43, 2, 2, 65,
	77, 8, 4, 1, 2, 66, 67, 7, 46, 2, 2, 67, 77, 8, 4, 1, 2, 68, 69, 7, 14,
	2, 2, 69, 77, 8, 4, 1, 2, 70, 71, 7, 15, 2, 2, 71, 77, 8, 4, 1, 2, 72,
	73, 5, 22, 12, 2, 73, 74, 7, 16, 2, 2, 74, 75, 8, 4, 1, 2, 75, 77, 3, 2,
	2, 2, 76, 46, 3, 2, 2, 2, 76, 51, 3, 2, 2, 2, 76, 55, 3, 2, 2, 2, 76, 59,
	3, 2, 2, 2, 76, 64, 3, 2, 2, 2, 76, 66, 3, 2, 2, 2, 76, 68, 3, 2, 2, 2,
	76, 70, 3, 2, 2, 2, 76, 72, 3, 2, 2, 2, 77, 113, 3, 2, 2, 2, 78, 79, 12,
	14, 2, 2, 79, 80, 9, 2, 2, 2, 80, 81, 5, 6, 4, 15, 81, 82, 8, 4, 1, 2,
	82, 112, 3, 2, 2, 2, 83, 84, 12, 13, 2, 2, 84, 85, 9, 3, 2, 2, 85, 86,
	5, 6, 4, 14, 86, 87, 8, 4, 1, 2, 87, 112, 3, 2, 2, 2, 88, 89, 12, 12, 2,
	2, 89, 90, 5, 8, 5, 2, 90, 91, 5, 6, 4, 13, 91, 92, 8, 4, 1, 2, 92, 112,
	3, 2, 2, 2, 93, 94, 12, 11, 2, 2, 94, 95, 7, 9, 2, 2, 95, 96, 5, 4, 3,
	2, 96, 97, 5, 6, 4, 12, 97, 98, 8, 4, 1, 2, 98, 112, 3, 2, 2, 2, 99, 100,
	12, 10, 2, 2, 100, 101, 7, 10, 2, 2, 101, 102, 5, 4, 3, 2, 102, 103, 5,
	6, 4, 11, 103, 104, 8, 4, 1, 2, 104, 112, 3, 2, 2, 2, 105, 106, 12, 9,
	2, 2, 106, 107, 7, 11, 2, 2, 107, 108, 5, 4, 3, 2, 108, 109, 5, 6, 4, 10,
	109, 110, 8, 4, 1, 2, 110, 112, 3, 2, 2, 2, 111, 78, 3, 2, 2, 2, 111, 83,
	3, 2, 2, 2, 111, 88, 3, 2, 2, 2, 111, 93, 3, 2, 2, 2, 111, 99, 3, 2, 2,
	2, 111, 105, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 7, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117, 7,
	17, 2, 2, 117, 129, 8, 5, 1, 2, 118, 119, 7, 18, 2, 2, 119, 129, 8, 5,
	1, 2, 120, 121, 7, 19, 2, 2, 121, 129, 8, 5, 1, 2, 122, 123, 7, 20, 2,
	2, 123, 129, 8, 5, 1, 2, 124, 125, 7, 21, 2, 2, 125, 129, 8, 5, 1, 2, 126,
	127, 7, 22, 2, 2, 127, 129, 8, 5, 1, 2, 128, 116, 3, 2, 2, 2, 128, 118,
	3, 2, 2, 2, 128, 120, 3, 2, 2, 2, 128, 122, 3, 2, 2, 2, 128, 124, 3, 2,
	2, 2, 128, 126, 3, 2, 2, 2, 129, 9, 3, 2, 2, 2, 130, 132, 5, 12, 7, 2,
	131, 130, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 11, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 144, 5,
	14, 8, 2, 137, 144, 5, 20, 11, 2, 138, 144, 5, 24, 13, 2, 139, 144, 5,
	28, 15, 2, 140, 144, 5, 30, 16, 2, 141, 144, 5, 32, 17, 2, 142, 144, 5,
	34, 18, 2, 143, 136, 3, 2, 2, 2, 143, 137, 3, 2, 2, 2, 143, 138, 3, 2,
	2, 2, 143, 139, 3, 2, 2, 2, 143, 140, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2,
	143, 142, 3, 2, 2, 2, 144, 13, 3, 2, 2, 2, 145, 146, 7, 23, 2, 2, 146,
	147, 7, 46, 2, 2, 147, 148, 7, 24, 2, 2, 148, 149, 5, 18, 10, 2, 149, 150,
	7, 25, 2, 2, 150, 151, 8, 8, 1, 2, 151, 165, 3, 2, 2, 2, 152, 153, 7, 23,
	2, 2, 153, 154, 7, 46, 2, 2, 154, 155, 7, 24, 2, 2, 155, 156, 7, 26, 2,
	2, 156, 157, 7, 27, 2, 2, 157, 158, 5, 16, 9, 2, 158, 159, 7, 16, 2, 2,
	159, 160, 7, 28, 2, 2, 160, 161, 5, 18, 10, 2, 161, 162, 7, 25, 2, 2, 162,
	163, 8, 8, 1, 2, 163, 165, 3, 2, 2, 2, 164, 145, 3, 2, 2, 2, 164, 152,
	3, 2, 2, 2, 165, 15, 3, 2, 2, 2, 166, 167, 7, 43, 2, 2, 167, 168, 7, 29,
	2, 2, 168, 169, 7, 43, 2, 2, 169, 177, 8, 9, 1, 2, 170, 171, 7, 30, 2,
	2, 171, 172, 7, 43, 2, 2, 172, 173, 7, 29, 2, 2, 173, 174, 7, 43, 2, 2,
	174, 176, 8, 9, 1, 2, 175, 170, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177,
	175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 17, 3, 2, 2, 2, 179, 177, 3,
	2, 2, 2, 180, 181, 7, 31, 2, 2, 181, 182, 8, 10, 1, 2, 182, 19, 3, 2, 2,
	2, 183, 184, 7, 46, 2, 2, 184, 185, 7, 32, 2, 2, 185, 186, 5, 6, 4, 2,
	186, 187, 7, 25, 2, 2, 187, 188, 8, 11, 1, 2, 188, 197, 3, 2, 2, 2, 189,
	190, 5, 22, 12, 2, 190, 191, 7, 16, 2, 2, 191, 192, 7, 32, 2, 2, 192, 193,
	5, 6, 4, 2, 193, 194, 7, 25, 2, 2, 194, 195, 8, 11, 1, 2, 195, 197, 3,
	2, 2, 2, 196, 183, 3, 2, 2, 2, 196, 189, 3, 2, 2, 2, 197, 21, 3, 2, 2,
	2, 198, 199, 8, 12, 1, 2, 199, 200, 7, 46, 2, 2, 200, 201, 7, 27, 2, 2,
	201, 202, 5, 6, 4, 2, 202, 203, 8, 12, 1, 2, 203, 211, 3, 2, 2, 2, 204,
	205, 12, 4, 2, 2, 205, 206, 7, 30, 2, 2, 206, 207, 5, 6, 4, 2, 207, 208,
	8, 12, 1, 2, 208, 210, 3, 2, 2, 2, 209, 204, 3, 2, 2, 2, 210, 213, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 23, 3, 2, 2, 2,
	213, 211, 3, 2, 2, 2, 214, 215, 7, 33, 2, 2, 215, 216, 5, 6, 4, 2, 216,
	217, 8, 13, 1, 2, 217, 218, 5, 36, 19, 2, 218, 228, 8, 13, 1, 2, 219, 220,
	7, 34, 2, 2, 220, 221, 7, 33, 2, 2, 221, 222, 5, 6, 4, 2, 222, 223, 8,
	13, 1, 2, 223, 224, 5, 36, 19, 2, 224, 225, 8, 13, 1, 2, 225, 227, 3, 2,
	2, 2, 226, 219, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2,
	228, 229, 3, 2, 2, 2, 229, 233, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 231,
	232, 7, 34, 2, 2, 232, 234, 5, 36, 19, 2, 233, 231, 3, 2, 2, 2, 233, 234,
	3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 8, 13, 1, 2, 236, 25, 3, 2,
	2, 2, 237, 238, 7, 35, 2, 2, 238, 239, 5, 6, 4, 2, 239, 250, 7, 36, 2,
	2, 240, 241, 7, 37, 2, 2, 241, 242, 5, 6, 4, 2, 242, 243, 7, 24, 2, 2,
	243, 246, 8, 14, 1, 2, 244, 247, 5, 36, 19, 2, 245, 247, 5, 38, 20, 2,
	246, 244, 3, 2, 2, 2, 246, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	249, 8, 14, 1, 2, 249, 251, 3, 2, 2, 2, 250, 240, 3, 2, 2, 2, 251, 252,
	3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 260, 3, 2,
	2, 2, 254, 255, 7, 38, 2, 2, 255, 258, 7, 24, 2, 2, 256, 259, 5, 36, 19,
	2, 257, 259, 5, 38, 20, 2, 258, 256, 3, 2, 2, 2, 258, 257, 3, 2, 2, 2,
	259, 261, 3, 2, 2, 2, 260, 254, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261,
	262, 3, 2, 2, 2, 262, 263, 8, 14, 1, 2, 263, 264, 7, 39, 2, 2, 264, 27,
	3, 2, 2, 2, 265, 266, 7, 35, 2, 2, 266, 267, 5, 6, 4, 2, 267, 268, 7, 36,
	2, 2, 268, 279, 8, 15, 1, 2, 269, 270, 7, 37, 2, 2, 270, 271, 5, 6, 4,
	2, 271, 272, 7, 24, 2, 2, 272, 275, 8, 15, 1, 2, 273, 276, 5, 36, 19, 2,
	274, 276, 5, 38, 20, 2, 275, 273, 3, 2, 2, 2, 275, 274, 3, 2, 2, 2, 276,
	277, 3, 2, 2, 2, 277, 278, 8, 15, 1, 2, 278, 280, 3, 2, 2, 2, 279, 269,
	3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 281, 282, 3, 2,
	2, 2, 282, 292, 3, 2, 2, 2, 283, 284, 7, 38, 2, 2, 284, 285, 7, 24, 2,
	2, 285, 288, 8, 15, 1, 2, 286, 289, 5, 36, 19, 2, 287, 289, 5, 38, 20,
	2, 288, 286, 3, 2, 2, 2, 288, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290,
	291, 8, 15, 1, 2, 291, 293, 3, 2, 2, 2, 292, 283, 3, 2, 2, 2, 292, 293,
	3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 295, 8, 15, 1, 2, 295, 296, 7, 39,
	2, 2, 296, 29, 3, 2, 2, 2, 297, 298, 7, 40, 2, 2, 298, 299, 8, 16, 1, 2,
	299, 300, 5, 6, 4, 2, 300, 301, 8, 16, 1, 2, 301, 302, 5, 36, 19, 2, 302,
	303, 8, 16, 1, 2, 303, 31, 3, 2, 2, 2, 304, 305, 7, 41, 2, 2, 305, 306,
	8, 17, 1, 2, 306, 307, 5, 36, 19, 2, 307, 308, 7, 40, 2, 2, 308, 309, 5,
	6, 4, 2, 309, 310, 7, 25, 2, 2, 310, 311, 8, 17, 1, 2, 311, 33, 3, 2, 2,
	2, 312, 313, 7, 42, 2, 2, 313, 314, 8, 18, 1, 2, 314, 315, 5, 36, 19, 2,
	315, 316, 8, 18, 1, 2, 316, 35, 3, 2, 2, 2, 317, 318, 8, 19, 1, 2, 318,
	319, 7, 36, 2, 2, 319, 320, 5, 10, 6, 2, 320, 321, 7, 39, 2, 2, 321, 322,
	8, 19, 1, 2, 322, 37, 3, 2, 2, 2, 323, 324, 8, 20, 1, 2, 324, 325, 5, 10,
	6, 2, 325, 326, 8, 20, 1, 2, 326, 39, 3, 2, 2, 2, 22, 76, 111, 113, 128,
	133, 143, 164, 177, 196, 211, 228, 233, 246, 252, 258, 260, 275, 281, 288,
	292,
}
var literalNames = []string{
	"", "'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "'&&'", "'^'", "'||'", "'('",
	"')'", "'true'", "'false'", "']'", "'=='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'let'", "':'", "';'", "'array'", "'['", "'of'", "'..'", "','",
	"'int'", "'='", "'if'", "'else'", "'switch'", "'{'", "'case'", "'default'",
	"'}'", "'while'", "'do'", "'loop'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "NUM", "FLOAT", "STRING", "ID", "MULCOMMENT", "UNICOMMENT",
	"WHITESPACE",
}

var ruleNames = []string{
	"start", "marcador", "expresion", "oprel", "instrucciones", "instruccion",
	"inst_declaracion", "ldims", "tipo", "inst_asignacion", "lref", "inst_if",
	"inst_switch_propuesta1", "inst_switch_propuesta2", "inst_while", "inst_doWhile",
	"inst_loop", "bloque", "bloqueSinLLaves",
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
	GramaticaParserNUM        = 41
	GramaticaParserFLOAT      = 42
	GramaticaParserSTRING     = 43
	GramaticaParserID         = 44
	GramaticaParserMULCOMMENT = 45
	GramaticaParserUNICOMMENT = 46
	GramaticaParserWHITESPACE = 47
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start                  = 0
	GramaticaParserRULE_marcador               = 1
	GramaticaParserRULE_expresion              = 2
	GramaticaParserRULE_oprel                  = 3
	GramaticaParserRULE_instrucciones          = 4
	GramaticaParserRULE_instruccion            = 5
	GramaticaParserRULE_inst_declaracion       = 6
	GramaticaParserRULE_ldims                  = 7
	GramaticaParserRULE_tipo                   = 8
	GramaticaParserRULE_inst_asignacion        = 9
	GramaticaParserRULE_lref                   = 10
	GramaticaParserRULE_inst_if                = 11
	GramaticaParserRULE_inst_switch_propuesta1 = 12
	GramaticaParserRULE_inst_switch_propuesta2 = 13
	GramaticaParserRULE_inst_while             = 14
	GramaticaParserRULE_inst_doWhile           = 15
	GramaticaParserRULE_inst_loop              = 16
	GramaticaParserRULE_bloque                 = 17
	GramaticaParserRULE_bloqueSinLLaves        = 18
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
		p.SetState(39)
		p.Instrucciones()
	}
	{
		p.SetState(40)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(45)

			var _m = p.Match(GramaticaParserT__0)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(46)

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
			p.SetState(49)

			var _m = p.Match(GramaticaParserT__1)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(50)

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
			p.SetState(53)

			var _m = p.Match(GramaticaParserT__2)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(54)

			var _x = p.expresion(13)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLv()

	case 4:
		{
			p.SetState(57)
			p.Match(GramaticaParserT__9)
		}
		{
			p.SetState(58)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e = _x
		}
		{
			p.SetState(59)
			p.Match(GramaticaParserT__10)
		}

		localctx.(*ExpresionContext).dir = localctx.(*ExpresionContext).GetE().GetDir()
		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLv()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).cad = localctx.(*ExpresionContext).GetE().GetCad()

	case 5:
		{
			p.SetState(62)

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
			p.SetState(64)

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
			p.SetState(66)
			p.Match(GramaticaParserT__11)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewLabel())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewLabel())
		gen.AddGoto(localctx.(*ExpresionContext).lv[0])

	case 8:
		{
			p.SetState(68)
			p.Match(GramaticaParserT__12)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewLabel())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewLabel())
		gen.AddGoto(localctx.(*ExpresionContext).lf[0])

	case 9:
		{
			p.SetState(70)

			var _x = p.lref(0)

			localctx.(*ExpresionContext).ref = _x
		}
		{
			p.SetState(71)
			p.Match(GramaticaParserT__13)
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.AddGetArray(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetRef().GetId(), localctx.(*ExpresionContext).GetRef().GetAux())

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(77)

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
					p.SetState(78)

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
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(82)

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
					p.SetState(83)

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
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(87)

					var _x = p.Oprel()

					localctx.(*ExpresionContext).opr = _x
				}
				{
					p.SetState(88)

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
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(92)

					var _m = p.Match(GramaticaParserT__6)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(93)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLv())
				}
				{
					p.SetState(94)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLf(), localctx.(*ExpresionContext).GetE2().GetLf())

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(98)

					var _m = p.Match(GramaticaParserT__7)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(99)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(100)

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
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(104)

					var _m = p.Match(GramaticaParserT__8)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(105)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(106)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()
				localctx.(*ExpresionContext).lv = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLv(), localctx.(*ExpresionContext).GetE2().GetLv())

			}

		}
		p.SetState(113)
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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)

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
			p.SetState(116)

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
			p.SetState(118)

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
			p.SetState(120)

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
			p.SetState(122)

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
			p.SetState(124)

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
	p.EnterRule(localctx, 8, GramaticaParserRULE_instrucciones)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(GramaticaParserT__20-21))|(1<<(GramaticaParserT__30-21))|(1<<(GramaticaParserT__32-21))|(1<<(GramaticaParserT__37-21))|(1<<(GramaticaParserT__38-21))|(1<<(GramaticaParserT__39-21))|(1<<(GramaticaParserID-21)))) != 0 {
		{
			p.SetState(128)
			p.Instruccion()
		}

		p.SetState(133)
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

func (s *InstruccionContext) Inst_switch_propuesta2() IInst_switch_propuesta2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_switch_propuesta2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_switch_propuesta2Context)
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
	p.EnterRule(localctx, 10, GramaticaParserRULE_instruccion)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Inst_declaracion()
		}

	case GramaticaParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Inst_asignacion()
		}

	case GramaticaParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.Inst_if()
		}

	case GramaticaParserT__32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.Inst_switch_propuesta2()
		}

	case GramaticaParserT__37:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.Inst_while()
		}

	case GramaticaParserT__38:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(139)
			p.Inst_doWhile()
		}

	case GramaticaParserT__39:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(140)
			p.Inst_loop()
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
	p.EnterRule(localctx, 12, GramaticaParserRULE_inst_declaracion)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Match(GramaticaParserT__20)
		}
		{
			p.SetState(144)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Inst_declaracionContext).id = _m
		}
		{
			p.SetState(145)
			p.Match(GramaticaParserT__21)
		}
		{
			p.SetState(146)

			var _x = p.Tipo()

			localctx.(*Inst_declaracionContext).t = _x
		}
		{
			p.SetState(147)
			p.Match(GramaticaParserT__22)
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
			p.SetState(150)
			p.Match(GramaticaParserT__20)
		}
		{
			p.SetState(151)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Inst_declaracionContext).id = _m
		}
		{
			p.SetState(152)
			p.Match(GramaticaParserT__21)
		}
		{
			p.SetState(153)
			p.Match(GramaticaParserT__23)
		}
		{
			p.SetState(154)
			p.Match(GramaticaParserT__24)
		}
		{
			p.SetState(155)

			var _x = p.Ldims()

			localctx.(*Inst_declaracionContext).dims = _x
		}
		{
			p.SetState(156)
			p.Match(GramaticaParserT__13)
		}
		{
			p.SetState(157)
			p.Match(GramaticaParserT__25)
		}
		{
			p.SetState(158)

			var _x = p.Tipo()

			localctx.(*Inst_declaracionContext).t = _x
		}
		{
			p.SetState(159)
			p.Match(GramaticaParserT__22)
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
	p.EnterRule(localctx, 14, GramaticaParserRULE_ldims)
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
		p.SetState(164)

		var _m = p.Match(GramaticaParserNUM)

		localctx.(*LdimsContext).n1 = _m
	}
	{
		p.SetState(165)
		p.Match(GramaticaParserT__26)
	}
	{
		p.SetState(166)

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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__27 {
		{
			p.SetState(168)
			p.Match(GramaticaParserT__27)
		}
		{
			p.SetState(169)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*LdimsContext).n3 = _m
		}
		{
			p.SetState(170)
			p.Match(GramaticaParserT__26)
		}
		{
			p.SetState(171)

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

		p.SetState(177)
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
	p.EnterRule(localctx, 16, GramaticaParserRULE_tipo)

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
	p.EnterRule(localctx, 18, GramaticaParserRULE_inst_asignacion)

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

	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Inst_asignacionContext).id = _m
		}
		{
			p.SetState(182)
			p.Match(GramaticaParserT__29)
		}
		{
			p.SetState(183)

			var _x = p.expresion(0)

			localctx.(*Inst_asignacionContext).e = _x
		}
		{
			p.SetState(184)
			p.Match(GramaticaParserT__22)
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
			p.SetState(187)

			var _x = p.lref(0)

			localctx.(*Inst_asignacionContext).ref = _x
		}
		{
			p.SetState(188)
			p.Match(GramaticaParserT__13)
		}
		{
			p.SetState(189)
			p.Match(GramaticaParserT__29)
		}
		{
			p.SetState(190)

			var _x = p.expresion(0)

			localctx.(*Inst_asignacionContext).e = _x
		}
		{
			p.SetState(191)
			p.Match(GramaticaParserT__22)
		}
		gen.AddSetArray(localctx.(*Inst_asignacionContext).GetRef().GetId(), localctx.(*Inst_asignacionContext).GetRef().GetAux(), localctx.(*Inst_asignacionContext).GetE().GetDir())

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

	// GetDim_ returns the dim_ attribute.
	GetDim_() int

	// SetId sets the id attribute.
	SetId(string)

	// SetAux sets the aux attribute.
	SetAux(string)

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

func (s *LrefContext) GetDim_() int { return s.dim_ }

func (s *LrefContext) SetId(v string) { s.id = v }

func (s *LrefContext) SetAux(v string) { s.aux = v }

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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, GramaticaParserRULE_lref, _p)

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
		p.SetState(197)

		var _m = p.Match(GramaticaParserID)

		localctx.(*LrefContext).id_ = _m
	}
	{
		p.SetState(198)
		p.Match(GramaticaParserT__24)
	}
	{
		p.SetState(199)

		var _x = p.expresion(0)

		localctx.(*LrefContext).e = _x
	}

	localctx.(*LrefContext).id = (func() string {
		if localctx.(*LrefContext).GetId_() == nil {
			return ""
		} else {
			return localctx.(*LrefContext).GetId_().GetText()
		}
	}())
	localctx.(*LrefContext).aux = gen.NewTemp()
	localctx.(*LrefContext).dim_ = 1
	gen.AddAsign(localctx.(*LrefContext).aux, localctx.(*LrefContext).GetE().GetDir())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLrefContext(p, _parentctx, _parentState)
			localctx.(*LrefContext).ref = _prevctx
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_lref)
			p.SetState(202)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(203)
				p.Match(GramaticaParserT__27)
			}
			{
				p.SetState(204)

				var _x = p.expresion(0)

				localctx.(*LrefContext).e = _x
			}

			tmp1 := gen.NewTemp()
			gen.AddExpresion(tmp1, localctx.(*LrefContext).GetE().GetDir(), "-", "1")

			tmp2 := gen.NewTemp()
			gen.AddExpresion(tmp2, tmp1, "*", entorno.GetTamOrdenColumn(localctx.(*LrefContext).GetRef().GetId(), localctx.(*LrefContext).GetRef().GetDim_(), tope))
			gen.AddExpresion(localctx.(*LrefContext).GetRef().GetAux(), localctx.(*LrefContext).GetRef().GetAux(), "+", tmp2)

			localctx.(*LrefContext).id = localctx.(*LrefContext).GetRef().GetId()
			localctx.(*LrefContext).aux = localctx.(*LrefContext).GetRef().GetAux()
			localctx.(*LrefContext).dim_ = localctx.(*LrefContext).GetRef().GetDim_() + 1

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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
		p.SetState(212)
		p.Match(GramaticaParserT__30)
	}
	{
		p.SetState(213)

		var _x = p.expresion(0)

		localctx.(*Inst_ifContext).e1 = _x
	}

	localctx.(*Inst_ifContext).lsalida = gen.NewLabel()
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLv())

	{
		p.SetState(215)
		p.Bloque()
	}

	gen.AddGoto(localctx.(*Inst_ifContext).lsalida)
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLf())

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(217)
				p.Match(GramaticaParserT__31)
			}
			{
				p.SetState(218)
				p.Match(GramaticaParserT__30)
			}
			{
				p.SetState(219)

				var _x = p.expresion(0)

				localctx.(*Inst_ifContext).e2 = _x
			}

			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLv())

			{
				p.SetState(221)
				p.Bloque()
			}

			gen.AddGoto(localctx.(*Inst_ifContext).lsalida)
			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLf())

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__31 {
		{
			p.SetState(229)
			p.Match(GramaticaParserT__31)
		}
		{
			p.SetState(230)
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

func (s *Inst_switch_propuesta1Context) AllBloqueSinLLaves() []IBloqueSinLLavesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem())
	var tst = make([]IBloqueSinLLavesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueSinLLavesContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) BloqueSinLLaves(i int) IBloqueSinLLavesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueSinLLavesContext)
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
		p.Match(GramaticaParserT__32)
	}
	{
		p.SetState(236)

		var _x = p.expresion(0)

		localctx.(*Inst_switch_propuesta1Context).e1 = _x
	}
	{
		p.SetState(237)
		p.Match(GramaticaParserT__33)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserT__34 {
		{
			p.SetState(238)
			p.Match(GramaticaParserT__34)
		}
		{
			p.SetState(239)

			var _x = p.expresion(0)

			localctx.(*Inst_switch_propuesta1Context).e2 = _x
		}
		{
			p.SetState(240)
			p.Match(GramaticaParserT__21)
		}

		localctx.(*Inst_switch_propuesta1Context).lv = gen.NewLabel()
		gen.AddIf(localctx.(*Inst_switch_propuesta1Context).GetE1().GetDir(), "!=", localctx.(*Inst_switch_propuesta1Context).GetE2().GetDir(), localctx.(*Inst_switch_propuesta1Context).lv)

		p.SetState(244)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(242)
				p.Bloque()
			}

		case GramaticaParserT__20, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__34, GramaticaParserT__35, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserID:
			{
				p.SetState(243)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.AddGoto(localctx.(*Inst_switch_propuesta1Context).lsalida)
		gen.AddLabel(localctx.(*Inst_switch_propuesta1Context).lv)

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__35 {
		{
			p.SetState(252)
			p.Match(GramaticaParserT__35)
		}
		{
			p.SetState(253)
			p.Match(GramaticaParserT__21)
		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(254)
				p.Bloque()
			}

		case GramaticaParserT__20, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserID:
			{
				p.SetState(255)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	gen.AddLabel(localctx.(*Inst_switch_propuesta1Context).lsalida)

	{
		p.SetState(261)
		p.Match(GramaticaParserT__36)
	}

	return localctx
}

// IInst_switch_propuesta2Context is an interface to support dynamic dispatch.
type IInst_switch_propuesta2Context interface {
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

	// GetLprueba returns the lprueba attribute.
	GetLprueba() string

	// GetCad returns the cad attribute.
	GetCad() string

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// GetLv returns the lv attribute.
	GetLv() string

	// GetDefecto returns the defecto attribute.
	GetDefecto() bool

	// SetLprueba sets the lprueba attribute.
	SetLprueba(string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// SetLv sets the lv attribute.
	SetLv(string)

	// SetDefecto sets the defecto attribute.
	SetDefecto(bool)

	// IsInst_switch_propuesta2Context differentiates from other interfaces.
	IsInst_switch_propuesta2Context()
}

type Inst_switch_propuesta2Context struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lprueba string
	cad     string
	lsalida string
	lv      string
	defecto bool
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_switch_propuesta2Context() *Inst_switch_propuesta2Context {
	var p = new(Inst_switch_propuesta2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta2
	return p
}

func (*Inst_switch_propuesta2Context) IsInst_switch_propuesta2Context() {}

func NewInst_switch_propuesta2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_switch_propuesta2Context {
	var p = new(Inst_switch_propuesta2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta2

	return p
}

func (s *Inst_switch_propuesta2Context) GetParser() antlr.Parser { return s.parser }

func (s *Inst_switch_propuesta2Context) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_switch_propuesta2Context) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_switch_propuesta2Context) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_switch_propuesta2Context) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_switch_propuesta2Context) GetLprueba() string { return s.lprueba }

func (s *Inst_switch_propuesta2Context) GetCad() string { return s.cad }

func (s *Inst_switch_propuesta2Context) GetLsalida() string { return s.lsalida }

func (s *Inst_switch_propuesta2Context) GetLv() string { return s.lv }

func (s *Inst_switch_propuesta2Context) GetDefecto() bool { return s.defecto }

func (s *Inst_switch_propuesta2Context) SetLprueba(v string) { s.lprueba = v }

func (s *Inst_switch_propuesta2Context) SetCad(v string) { s.cad = v }

func (s *Inst_switch_propuesta2Context) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_switch_propuesta2Context) SetLv(v string) { s.lv = v }

func (s *Inst_switch_propuesta2Context) SetDefecto(v bool) { s.defecto = v }

func (s *Inst_switch_propuesta2Context) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta2Context) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_switch_propuesta2Context) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta2Context) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_switch_propuesta2Context) AllBloqueSinLLaves() []IBloqueSinLLavesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem())
	var tst = make([]IBloqueSinLLavesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueSinLLavesContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta2Context) BloqueSinLLaves(i int) IBloqueSinLLavesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueSinLLavesContext)
}

func (s *Inst_switch_propuesta2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_switch_propuesta2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_switch_propuesta2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_switch_propuesta2(s)
	}
}

func (s *Inst_switch_propuesta2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_switch_propuesta2(s)
	}
}

func (p *GramaticaParser) Inst_switch_propuesta2() (localctx IInst_switch_propuesta2Context) {
	this := p
	_ = this

	localctx = NewInst_switch_propuesta2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GramaticaParserRULE_inst_switch_propuesta2)

	localctx.(*Inst_switch_propuesta2Context).lprueba = gen.NewLabel()
	localctx.(*Inst_switch_propuesta2Context).lsalida = gen.NewLabel()
	localctx.(*Inst_switch_propuesta2Context).cad = ""
	localctx.(*Inst_switch_propuesta2Context).defecto = false

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
		p.SetState(263)
		p.Match(GramaticaParserT__32)
	}
	{
		p.SetState(264)

		var _x = p.expresion(0)

		localctx.(*Inst_switch_propuesta2Context).e1 = _x
	}
	{
		p.SetState(265)
		p.Match(GramaticaParserT__33)
	}

	gen.AddGoto(localctx.(*Inst_switch_propuesta2Context).lprueba)

	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserT__34 {
		{
			p.SetState(267)
			p.Match(GramaticaParserT__34)
		}
		{
			p.SetState(268)

			var _x = p.expresion(0)

			localctx.(*Inst_switch_propuesta2Context).e2 = _x
		}
		{
			p.SetState(269)
			p.Match(GramaticaParserT__21)
		}

		localctx.(*Inst_switch_propuesta2Context).lv = gen.NewLabel()
		gen.AddLabel(localctx.(*Inst_switch_propuesta2Context).lv)
		localctx.(*Inst_switch_propuesta2Context).cad += "if " + localctx.(*Inst_switch_propuesta2Context).GetE1().GetDir() + " = " + localctx.(*Inst_switch_propuesta2Context).GetE2().GetDir() + " then goto " + localctx.(*Inst_switch_propuesta2Context).lv + "\n"

		p.SetState(273)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(271)
				p.Bloque()
			}

		case GramaticaParserT__20, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__34, GramaticaParserT__35, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserID:
			{
				p.SetState(272)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.AddGoto(localctx.(*Inst_switch_propuesta2Context).lsalida)

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__35 {
		{
			p.SetState(281)
			p.Match(GramaticaParserT__35)
		}
		{
			p.SetState(282)
			p.Match(GramaticaParserT__21)
		}

		localctx.(*Inst_switch_propuesta2Context).lv = gen.NewLabel()
		localctx.(*Inst_switch_propuesta2Context).defecto = true
		gen.AddLabel(localctx.(*Inst_switch_propuesta2Context).lv)

		p.SetState(286)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__33:
			{
				p.SetState(284)
				p.Bloque()
			}

		case GramaticaParserT__20, GramaticaParserT__30, GramaticaParserT__32, GramaticaParserT__36, GramaticaParserT__37, GramaticaParserT__38, GramaticaParserT__39, GramaticaParserID:
			{
				p.SetState(285)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.AddGoto(localctx.(*Inst_switch_propuesta2Context).lsalida)

	}

	gen.AddLabel(localctx.(*Inst_switch_propuesta2Context).lprueba)
	gen.Gen(localctx.(*Inst_switch_propuesta2Context).cad)
	if localctx.(*Inst_switch_propuesta2Context).defecto {
		gen.AddGoto(localctx.(*Inst_switch_propuesta2Context).lv)
	}
	gen.AddLabel(localctx.(*Inst_switch_propuesta2Context).lsalida)

	{
		p.SetState(293)
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
		p.SetState(295)
		p.Match(GramaticaParserT__37)
	}
	gen.AddLabel(localctx.(*Inst_whileContext).linicio)
	{
		p.SetState(297)

		var _x = p.expresion(0)

		localctx.(*Inst_whileContext).e = _x
	}
	gen.Soltar(localctx.(*Inst_whileContext).GetE().GetLv())
	{
		p.SetState(299)
		p.Bloque()
	}

	gen.AddGoto(localctx.(*Inst_whileContext).linicio)
	gen.Soltar(localctx.(*Inst_whileContext).GetE().GetLf())

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
		p.SetState(302)
		p.Match(GramaticaParserT__38)
	}

	localctx.(*Inst_doWhileContext).linicio = gen.NewLabel()
	gen.AddLabel(localctx.(*Inst_doWhileContext).linicio)

	{
		p.SetState(304)
		p.Bloque()
	}
	{
		p.SetState(305)
		p.Match(GramaticaParserT__37)
	}
	{
		p.SetState(306)

		var _x = p.expresion(0)

		localctx.(*Inst_doWhileContext).e = _x
	}
	{
		p.SetState(307)
		p.Match(GramaticaParserT__22)
	}

	gen.Soltar(localctx.(*Inst_doWhileContext).GetE().GetLv())
	gen.AddGoto(localctx.(*Inst_doWhileContext).linicio)
	gen.Soltar(localctx.(*Inst_doWhileContext).GetE().GetLf())

	return localctx
}

// IInst_loopContext is an interface to support dynamic dispatch.
type IInst_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// IsInst_loopContext differentiates from other interfaces.
	IsInst_loopContext()
}

type Inst_loopContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
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

func (s *Inst_loopContext) SetLinicio(v string) { s.linicio = v }

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
		p.SetState(310)
		p.Match(GramaticaParserT__39)
	}

	localctx.(*Inst_loopContext).linicio = gen.NewLabel()
	gen.AddLabel(localctx.(*Inst_loopContext).linicio)

	{
		p.SetState(312)
		p.Bloque()
	}

	gen.AddGoto(localctx.(*Inst_loopContext).linicio)

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
	p.EnterRule(localctx, 34, GramaticaParserRULE_bloque)

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
		p.SetState(316)
		p.Match(GramaticaParserT__33)
	}
	{
		p.SetState(317)
		p.Instrucciones()
	}
	{
		p.SetState(318)
		p.Match(GramaticaParserT__36)
	}

	tope = entorno.Pop()

	return localctx
}

// IBloqueSinLLavesContext is an interface to support dynamic dispatch.
type IBloqueSinLLavesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueSinLLavesContext differentiates from other interfaces.
	IsBloqueSinLLavesContext()
}

type BloqueSinLLavesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueSinLLavesContext() *BloqueSinLLavesContext {
	var p = new(BloqueSinLLavesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_bloqueSinLLaves
	return p
}

func (*BloqueSinLLavesContext) IsBloqueSinLLavesContext() {}

func NewBloqueSinLLavesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueSinLLavesContext {
	var p = new(BloqueSinLLavesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_bloqueSinLLaves

	return p
}

func (s *BloqueSinLLavesContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueSinLLavesContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueSinLLavesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueSinLLavesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueSinLLavesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBloqueSinLLaves(s)
	}
}

func (s *BloqueSinLLavesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBloqueSinLLaves(s)
	}
}

func (p *GramaticaParser) BloqueSinLLaves() (localctx IBloqueSinLLavesContext) {
	this := p
	_ = this

	localctx = NewBloqueSinLLavesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GramaticaParserRULE_bloqueSinLLaves)

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
		p.SetState(322)
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

	case 10:
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
