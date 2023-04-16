// Code generated from ../C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var clexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clexerLexerInit() {
	staticData := &clexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "'#include'", "'('", "')'", "'{'", "'}'", "'struct'", "','",
		"'asm'", "'if'", "'return'", "'goto'", "':'", "'for'", "'while'", "'do'",
		"'switch'", "'break'", "'continue'", "'-'", "'!'", "'&'", "'^'", "'|'",
		"'&&'", "'||'", "'?'", "'.'", "'='", "'const'", "'+='", "'-='", "'case'",
		"'default'", "'else if'", "'else'", "'*'", "'/'", "'%'", "'+'", "'=='",
		"'!='", "'<'", "'>'", "'<='", "'>='", "'++'", "'--'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "STRING",
		"ID", "WS", "SINGLE_COMMENT", "MULTILINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "INT",
		"STRING", "ID", "WS", "SINGLE_COMMENT", "MULTILINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 54, 336, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 4, 48, 285, 8, 48, 11, 48, 12,
		48, 286, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 293, 8, 49, 10, 49, 12, 49,
		296, 9, 49, 1, 49, 1, 49, 1, 50, 1, 50, 5, 50, 302, 8, 50, 10, 50, 12,
		50, 305, 9, 50, 1, 51, 4, 51, 308, 8, 51, 11, 51, 12, 51, 309, 1, 51, 1,
		51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 318, 8, 52, 10, 52, 12, 52, 321,
		9, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 5, 53, 329, 8, 53, 10,
		53, 12, 53, 332, 9, 53, 1, 53, 1, 53, 1, 53, 2, 319, 330, 0, 54, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
		21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59,
		30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77,
		39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95,
		48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 1, 0, 5, 1, 0,
		48, 57, 2, 0, 34, 34, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 342, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0,
		71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0,
		0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0,
		0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0,
		0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0,
		1, 109, 1, 0, 0, 0, 3, 111, 1, 0, 0, 0, 5, 120, 1, 0, 0, 0, 7, 122, 1,
		0, 0, 0, 9, 124, 1, 0, 0, 0, 11, 126, 1, 0, 0, 0, 13, 128, 1, 0, 0, 0,
		15, 135, 1, 0, 0, 0, 17, 137, 1, 0, 0, 0, 19, 141, 1, 0, 0, 0, 21, 144,
		1, 0, 0, 0, 23, 151, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0, 27, 158, 1, 0, 0,
		0, 29, 162, 1, 0, 0, 0, 31, 168, 1, 0, 0, 0, 33, 171, 1, 0, 0, 0, 35, 178,
		1, 0, 0, 0, 37, 184, 1, 0, 0, 0, 39, 193, 1, 0, 0, 0, 41, 195, 1, 0, 0,
		0, 43, 197, 1, 0, 0, 0, 45, 199, 1, 0, 0, 0, 47, 201, 1, 0, 0, 0, 49, 203,
		1, 0, 0, 0, 51, 206, 1, 0, 0, 0, 53, 209, 1, 0, 0, 0, 55, 211, 1, 0, 0,
		0, 57, 213, 1, 0, 0, 0, 59, 215, 1, 0, 0, 0, 61, 221, 1, 0, 0, 0, 63, 224,
		1, 0, 0, 0, 65, 227, 1, 0, 0, 0, 67, 232, 1, 0, 0, 0, 69, 240, 1, 0, 0,
		0, 71, 248, 1, 0, 0, 0, 73, 253, 1, 0, 0, 0, 75, 255, 1, 0, 0, 0, 77, 257,
		1, 0, 0, 0, 79, 259, 1, 0, 0, 0, 81, 261, 1, 0, 0, 0, 83, 264, 1, 0, 0,
		0, 85, 267, 1, 0, 0, 0, 87, 269, 1, 0, 0, 0, 89, 271, 1, 0, 0, 0, 91, 274,
		1, 0, 0, 0, 93, 277, 1, 0, 0, 0, 95, 280, 1, 0, 0, 0, 97, 284, 1, 0, 0,
		0, 99, 288, 1, 0, 0, 0, 101, 299, 1, 0, 0, 0, 103, 307, 1, 0, 0, 0, 105,
		313, 1, 0, 0, 0, 107, 324, 1, 0, 0, 0, 109, 110, 5, 59, 0, 0, 110, 2, 1,
		0, 0, 0, 111, 112, 5, 35, 0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 110,
		0, 0, 114, 115, 5, 99, 0, 0, 115, 116, 5, 108, 0, 0, 116, 117, 5, 117,
		0, 0, 117, 118, 5, 100, 0, 0, 118, 119, 5, 101, 0, 0, 119, 4, 1, 0, 0,
		0, 120, 121, 5, 40, 0, 0, 121, 6, 1, 0, 0, 0, 122, 123, 5, 41, 0, 0, 123,
		8, 1, 0, 0, 0, 124, 125, 5, 123, 0, 0, 125, 10, 1, 0, 0, 0, 126, 127, 5,
		125, 0, 0, 127, 12, 1, 0, 0, 0, 128, 129, 5, 115, 0, 0, 129, 130, 5, 116,
		0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 117, 0, 0, 132, 133, 5, 99,
		0, 0, 133, 134, 5, 116, 0, 0, 134, 14, 1, 0, 0, 0, 135, 136, 5, 44, 0,
		0, 136, 16, 1, 0, 0, 0, 137, 138, 5, 97, 0, 0, 138, 139, 5, 115, 0, 0,
		139, 140, 5, 109, 0, 0, 140, 18, 1, 0, 0, 0, 141, 142, 5, 105, 0, 0, 142,
		143, 5, 102, 0, 0, 143, 20, 1, 0, 0, 0, 144, 145, 5, 114, 0, 0, 145, 146,
		5, 101, 0, 0, 146, 147, 5, 116, 0, 0, 147, 148, 5, 117, 0, 0, 148, 149,
		5, 114, 0, 0, 149, 150, 5, 110, 0, 0, 150, 22, 1, 0, 0, 0, 151, 152, 5,
		103, 0, 0, 152, 153, 5, 111, 0, 0, 153, 154, 5, 116, 0, 0, 154, 155, 5,
		111, 0, 0, 155, 24, 1, 0, 0, 0, 156, 157, 5, 58, 0, 0, 157, 26, 1, 0, 0,
		0, 158, 159, 5, 102, 0, 0, 159, 160, 5, 111, 0, 0, 160, 161, 5, 114, 0,
		0, 161, 28, 1, 0, 0, 0, 162, 163, 5, 119, 0, 0, 163, 164, 5, 104, 0, 0,
		164, 165, 5, 105, 0, 0, 165, 166, 5, 108, 0, 0, 166, 167, 5, 101, 0, 0,
		167, 30, 1, 0, 0, 0, 168, 169, 5, 100, 0, 0, 169, 170, 5, 111, 0, 0, 170,
		32, 1, 0, 0, 0, 171, 172, 5, 115, 0, 0, 172, 173, 5, 119, 0, 0, 173, 174,
		5, 105, 0, 0, 174, 175, 5, 116, 0, 0, 175, 176, 5, 99, 0, 0, 176, 177,
		5, 104, 0, 0, 177, 34, 1, 0, 0, 0, 178, 179, 5, 98, 0, 0, 179, 180, 5,
		114, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182, 5, 97, 0, 0, 182, 183, 5,
		107, 0, 0, 183, 36, 1, 0, 0, 0, 184, 185, 5, 99, 0, 0, 185, 186, 5, 111,
		0, 0, 186, 187, 5, 110, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 105,
		0, 0, 189, 190, 5, 110, 0, 0, 190, 191, 5, 117, 0, 0, 191, 192, 5, 101,
		0, 0, 192, 38, 1, 0, 0, 0, 193, 194, 5, 45, 0, 0, 194, 40, 1, 0, 0, 0,
		195, 196, 5, 33, 0, 0, 196, 42, 1, 0, 0, 0, 197, 198, 5, 38, 0, 0, 198,
		44, 1, 0, 0, 0, 199, 200, 5, 94, 0, 0, 200, 46, 1, 0, 0, 0, 201, 202, 5,
		124, 0, 0, 202, 48, 1, 0, 0, 0, 203, 204, 5, 38, 0, 0, 204, 205, 5, 38,
		0, 0, 205, 50, 1, 0, 0, 0, 206, 207, 5, 124, 0, 0, 207, 208, 5, 124, 0,
		0, 208, 52, 1, 0, 0, 0, 209, 210, 5, 63, 0, 0, 210, 54, 1, 0, 0, 0, 211,
		212, 5, 46, 0, 0, 212, 56, 1, 0, 0, 0, 213, 214, 5, 61, 0, 0, 214, 58,
		1, 0, 0, 0, 215, 216, 5, 99, 0, 0, 216, 217, 5, 111, 0, 0, 217, 218, 5,
		110, 0, 0, 218, 219, 5, 115, 0, 0, 219, 220, 5, 116, 0, 0, 220, 60, 1,
		0, 0, 0, 221, 222, 5, 43, 0, 0, 222, 223, 5, 61, 0, 0, 223, 62, 1, 0, 0,
		0, 224, 225, 5, 45, 0, 0, 225, 226, 5, 61, 0, 0, 226, 64, 1, 0, 0, 0, 227,
		228, 5, 99, 0, 0, 228, 229, 5, 97, 0, 0, 229, 230, 5, 115, 0, 0, 230, 231,
		5, 101, 0, 0, 231, 66, 1, 0, 0, 0, 232, 233, 5, 100, 0, 0, 233, 234, 5,
		101, 0, 0, 234, 235, 5, 102, 0, 0, 235, 236, 5, 97, 0, 0, 236, 237, 5,
		117, 0, 0, 237, 238, 5, 108, 0, 0, 238, 239, 5, 116, 0, 0, 239, 68, 1,
		0, 0, 0, 240, 241, 5, 101, 0, 0, 241, 242, 5, 108, 0, 0, 242, 243, 5, 115,
		0, 0, 243, 244, 5, 101, 0, 0, 244, 245, 5, 32, 0, 0, 245, 246, 5, 105,
		0, 0, 246, 247, 5, 102, 0, 0, 247, 70, 1, 0, 0, 0, 248, 249, 5, 101, 0,
		0, 249, 250, 5, 108, 0, 0, 250, 251, 5, 115, 0, 0, 251, 252, 5, 101, 0,
		0, 252, 72, 1, 0, 0, 0, 253, 254, 5, 42, 0, 0, 254, 74, 1, 0, 0, 0, 255,
		256, 5, 47, 0, 0, 256, 76, 1, 0, 0, 0, 257, 258, 5, 37, 0, 0, 258, 78,
		1, 0, 0, 0, 259, 260, 5, 43, 0, 0, 260, 80, 1, 0, 0, 0, 261, 262, 5, 61,
		0, 0, 262, 263, 5, 61, 0, 0, 263, 82, 1, 0, 0, 0, 264, 265, 5, 33, 0, 0,
		265, 266, 5, 61, 0, 0, 266, 84, 1, 0, 0, 0, 267, 268, 5, 60, 0, 0, 268,
		86, 1, 0, 0, 0, 269, 270, 5, 62, 0, 0, 270, 88, 1, 0, 0, 0, 271, 272, 5,
		60, 0, 0, 272, 273, 5, 61, 0, 0, 273, 90, 1, 0, 0, 0, 274, 275, 5, 62,
		0, 0, 275, 276, 5, 61, 0, 0, 276, 92, 1, 0, 0, 0, 277, 278, 5, 43, 0, 0,
		278, 279, 5, 43, 0, 0, 279, 94, 1, 0, 0, 0, 280, 281, 5, 45, 0, 0, 281,
		282, 5, 45, 0, 0, 282, 96, 1, 0, 0, 0, 283, 285, 7, 0, 0, 0, 284, 283,
		1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0,
		0, 0, 287, 98, 1, 0, 0, 0, 288, 294, 5, 34, 0, 0, 289, 293, 8, 1, 0, 0,
		290, 291, 5, 92, 0, 0, 291, 293, 9, 0, 0, 0, 292, 289, 1, 0, 0, 0, 292,
		290, 1, 0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295,
		1, 0, 0, 0, 295, 297, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 298, 5, 34,
		0, 0, 298, 100, 1, 0, 0, 0, 299, 303, 7, 2, 0, 0, 300, 302, 7, 3, 0, 0,
		301, 300, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303,
		304, 1, 0, 0, 0, 304, 102, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 308,
		7, 4, 0, 0, 307, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 307, 1, 0,
		0, 0, 309, 310, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 312, 6, 51, 0, 0,
		312, 104, 1, 0, 0, 0, 313, 314, 5, 47, 0, 0, 314, 315, 5, 47, 0, 0, 315,
		319, 1, 0, 0, 0, 316, 318, 9, 0, 0, 0, 317, 316, 1, 0, 0, 0, 318, 321,
		1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 322, 1, 0,
		0, 0, 321, 319, 1, 0, 0, 0, 322, 323, 5, 10, 0, 0, 323, 106, 1, 0, 0, 0,
		324, 325, 5, 47, 0, 0, 325, 326, 5, 42, 0, 0, 326, 330, 1, 0, 0, 0, 327,
		329, 9, 0, 0, 0, 328, 327, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 331,
		1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0, 332, 330, 1, 0,
		0, 0, 333, 334, 5, 42, 0, 0, 334, 335, 5, 47, 0, 0, 335, 108, 1, 0, 0,
		0, 8, 0, 286, 292, 294, 303, 309, 319, 330, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CLexerInit initializes any static state used to implement CLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CLexerInit() {
	staticData := &clexerLexerStaticData
	staticData.once.Do(clexerLexerInit)
}

// NewCLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCLexer(input antlr.CharStream) *CLexer {
	CLexerInit()
	l := new(CLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &clexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "C.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CLexer tokens.
const (
	CLexerT__0              = 1
	CLexerT__1              = 2
	CLexerT__2              = 3
	CLexerT__3              = 4
	CLexerT__4              = 5
	CLexerT__5              = 6
	CLexerT__6              = 7
	CLexerT__7              = 8
	CLexerT__8              = 9
	CLexerT__9              = 10
	CLexerT__10             = 11
	CLexerT__11             = 12
	CLexerT__12             = 13
	CLexerT__13             = 14
	CLexerT__14             = 15
	CLexerT__15             = 16
	CLexerT__16             = 17
	CLexerT__17             = 18
	CLexerT__18             = 19
	CLexerT__19             = 20
	CLexerT__20             = 21
	CLexerT__21             = 22
	CLexerT__22             = 23
	CLexerT__23             = 24
	CLexerT__24             = 25
	CLexerT__25             = 26
	CLexerT__26             = 27
	CLexerT__27             = 28
	CLexerT__28             = 29
	CLexerT__29             = 30
	CLexerT__30             = 31
	CLexerT__31             = 32
	CLexerT__32             = 33
	CLexerT__33             = 34
	CLexerT__34             = 35
	CLexerT__35             = 36
	CLexerT__36             = 37
	CLexerT__37             = 38
	CLexerT__38             = 39
	CLexerT__39             = 40
	CLexerT__40             = 41
	CLexerT__41             = 42
	CLexerT__42             = 43
	CLexerT__43             = 44
	CLexerT__44             = 45
	CLexerT__45             = 46
	CLexerT__46             = 47
	CLexerT__47             = 48
	CLexerINT               = 49
	CLexerSTRING            = 50
	CLexerID                = 51
	CLexerWS                = 52
	CLexerSINGLE_COMMENT    = 53
	CLexerMULTILINE_COMMENT = 54
)
