// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VlangParser struct {
	*antlr.BaseParser
}

var VlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'println'", "'print'", "'fn'", "'struct'", "'int'", "'float'",
		"'string'", "'bool'", "'char'", "'void'", "'len'", "'cap'", "'append'",
		"'mut'", "", "", "", "", "", "", "", "'if'", "'else'", "'while'", "'for'",
		"'in'", "'return'", "'break'", "'continue'", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'!'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='", "'>'",
		"'>='", "'='", "'++'", "'--'", "':'", "'('", "')'", "'['", "'{'", "']'",
		"'}'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "LEN", "CAP", "APPEND",
		"MUT", "BOOLEANO", "ENTERO", "FLOAT", "DECIMAL", "CADENA", "CARACTER",
		"ID", "IF_KW", "ELSE_KW", "WHILE_KW", "FOR_KW", "IN_KW", "RETURN_KW",
		"BREAK_KW", "CONTINUE_KW", "PLUS", "MINUS", "MUL", "DIV", "MOD", "NOT",
		"OR", "AND", "EQ", "NEQ", "LT", "LE", "GT", "GE", "ASSIGN", "INC", "DEC",
		"COLON", "LPAREN", "RPAREN", "LBRACK", "LCOR", "RBRACK", "RCOR", "DOT",
		"COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "stmt", "if_stmt", "if_chain", "else_stmt", "println", "print",
		"while_stmt", "for_stmt", "range", "func_call", "func_dcl", "struct_dcl",
		"transfer_stmt", "assign_stmt", "decl_stmt", "id_pattern", "expresion",
		"parametros", "func_param", "arg_list", "func_arg", "valores", "valor",
		"strct_instancia", "struct_prop", "var_type", "struct_vector", "incredecre",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 352, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 4, 0, 60, 8, 0, 11, 0, 12, 0,
		61, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 83, 8, 2, 10, 2,
		12, 2, 86, 9, 2, 1, 2, 3, 2, 89, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 95,
		8, 3, 10, 3, 12, 3, 98, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 105,
		8, 4, 10, 4, 12, 4, 108, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 122, 8, 6, 10, 6, 12, 6, 125, 9, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 133, 8, 7, 10, 7, 12, 7, 136,
		9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 145, 8, 8, 1, 8,
		1, 8, 5, 8, 149, 8, 8, 10, 8, 12, 8, 152, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 165, 8, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 173, 8, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 178, 8, 11, 10, 11, 12, 11, 181, 9, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 5, 12, 189, 8, 12, 10, 12, 12, 12, 192, 9, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 3, 13, 198, 8, 13, 1, 13, 1, 13, 3, 13, 202, 8,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 219, 8, 15, 1, 16, 1, 16, 1,
		16, 5, 16, 224, 8, 16, 10, 16, 12, 16, 227, 9, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		3, 17, 253, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		5, 17, 273, 8, 17, 10, 17, 12, 17, 276, 9, 17, 1, 18, 1, 18, 1, 18, 5,
		18, 281, 8, 18, 10, 18, 12, 18, 284, 9, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 5, 20, 291, 8, 20, 10, 20, 12, 20, 294, 9, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 307,
		8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 313, 8, 24, 10, 24, 12, 24, 316,
		9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		3, 26, 338, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 3, 28, 350, 8, 28, 1, 28, 0, 1, 34, 29, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 0, 5, 2, 0, 31, 31, 35, 35, 1, 0, 32, 34, 1, 0,
		30, 31, 1, 0, 40, 43, 1, 0, 38, 39, 380, 0, 59, 1, 0, 0, 0, 2, 77, 1, 0,
		0, 0, 4, 79, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 101, 1, 0, 0, 0, 10, 111,
		1, 0, 0, 0, 12, 116, 1, 0, 0, 0, 14, 128, 1, 0, 0, 0, 16, 139, 1, 0, 0,
		0, 18, 155, 1, 0, 0, 0, 20, 161, 1, 0, 0, 0, 22, 168, 1, 0, 0, 0, 24, 184,
		1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 203, 1, 0, 0, 0, 30, 218, 1, 0, 0,
		0, 32, 220, 1, 0, 0, 0, 34, 252, 1, 0, 0, 0, 36, 277, 1, 0, 0, 0, 38, 285,
		1, 0, 0, 0, 40, 287, 1, 0, 0, 0, 42, 295, 1, 0, 0, 0, 44, 298, 1, 0, 0,
		0, 46, 306, 1, 0, 0, 0, 48, 308, 1, 0, 0, 0, 50, 319, 1, 0, 0, 0, 52, 337,
		1, 0, 0, 0, 54, 339, 1, 0, 0, 0, 56, 349, 1, 0, 0, 0, 58, 60, 3, 2, 1,
		0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 5, 0, 0, 1, 64, 1, 1, 0, 0, 0,
		65, 78, 3, 30, 15, 0, 66, 78, 3, 28, 14, 0, 67, 78, 3, 4, 2, 0, 68, 78,
		3, 14, 7, 0, 69, 78, 3, 16, 8, 0, 70, 78, 3, 20, 10, 0, 71, 78, 3, 22,
		11, 0, 72, 78, 3, 24, 12, 0, 73, 78, 3, 48, 24, 0, 74, 78, 3, 26, 13, 0,
		75, 78, 3, 10, 5, 0, 76, 78, 3, 12, 6, 0, 77, 65, 1, 0, 0, 0, 77, 66, 1,
		0, 0, 0, 77, 67, 1, 0, 0, 0, 77, 68, 1, 0, 0, 0, 77, 69, 1, 0, 0, 0, 77,
		70, 1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77, 73, 1, 0, 0,
		0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 3, 1,
		0, 0, 0, 79, 84, 3, 6, 3, 0, 80, 81, 5, 23, 0, 0, 81, 83, 3, 6, 3, 0, 82,
		80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0,
		0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89, 3, 8, 4, 0, 88, 87,
		1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 5, 1, 0, 0, 0, 90, 91, 5, 22, 0, 0,
		91, 92, 3, 34, 17, 0, 92, 96, 5, 51, 0, 0, 93, 95, 3, 2, 1, 0, 94, 93,
		1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 99, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 53, 0, 0, 100, 7, 1,
		0, 0, 0, 101, 102, 5, 23, 0, 0, 102, 106, 5, 51, 0, 0, 103, 105, 3, 2,
		1, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0,
		106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109,
		110, 5, 53, 0, 0, 110, 9, 1, 0, 0, 0, 111, 112, 5, 1, 0, 0, 112, 113, 5,
		48, 0, 0, 113, 114, 3, 34, 17, 0, 114, 115, 5, 49, 0, 0, 115, 11, 1, 0,
		0, 0, 116, 117, 5, 2, 0, 0, 117, 118, 5, 48, 0, 0, 118, 123, 3, 34, 17,
		0, 119, 120, 5, 55, 0, 0, 120, 122, 3, 34, 17, 0, 121, 119, 1, 0, 0, 0,
		122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 49, 0, 0, 127, 13,
		1, 0, 0, 0, 128, 129, 5, 24, 0, 0, 129, 130, 3, 34, 17, 0, 130, 134, 5,
		51, 0, 0, 131, 133, 3, 2, 1, 0, 132, 131, 1, 0, 0, 0, 133, 136, 1, 0, 0,
		0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 1, 0, 0, 0, 136,
		134, 1, 0, 0, 0, 137, 138, 5, 53, 0, 0, 138, 15, 1, 0, 0, 0, 139, 140,
		5, 25, 0, 0, 140, 141, 5, 21, 0, 0, 141, 144, 5, 26, 0, 0, 142, 145, 3,
		34, 17, 0, 143, 145, 3, 18, 9, 0, 144, 142, 1, 0, 0, 0, 144, 143, 1, 0,
		0, 0, 145, 146, 1, 0, 0, 0, 146, 150, 5, 51, 0, 0, 147, 149, 3, 2, 1, 0,
		148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150,
		151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154,
		5, 53, 0, 0, 154, 17, 1, 0, 0, 0, 155, 156, 3, 34, 17, 0, 156, 157, 5,
		54, 0, 0, 157, 158, 5, 54, 0, 0, 158, 159, 5, 54, 0, 0, 159, 160, 3, 34,
		17, 0, 160, 19, 1, 0, 0, 0, 161, 162, 3, 32, 16, 0, 162, 164, 5, 48, 0,
		0, 163, 165, 3, 36, 18, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 167, 5, 49, 0, 0, 167, 21, 1, 0, 0, 0, 168,
		169, 5, 3, 0, 0, 169, 170, 5, 21, 0, 0, 170, 172, 5, 48, 0, 0, 171, 173,
		3, 40, 20, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 1,
		0, 0, 0, 174, 175, 5, 49, 0, 0, 175, 179, 5, 51, 0, 0, 176, 178, 3, 2,
		1, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0,
		179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182,
		183, 5, 53, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 5, 4, 0, 0, 185, 186,
		5, 21, 0, 0, 186, 190, 5, 51, 0, 0, 187, 189, 3, 28, 14, 0, 188, 187, 1,
		0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0,
		0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 53, 0, 0, 194,
		25, 1, 0, 0, 0, 195, 197, 5, 27, 0, 0, 196, 198, 3, 34, 17, 0, 197, 196,
		1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 202, 1, 0, 0, 0, 199, 202, 5, 28,
		0, 0, 200, 202, 5, 29, 0, 0, 201, 195, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0,
		201, 200, 1, 0, 0, 0, 202, 27, 1, 0, 0, 0, 203, 204, 3, 32, 16, 0, 204,
		205, 5, 44, 0, 0, 205, 206, 3, 34, 17, 0, 206, 29, 1, 0, 0, 0, 207, 208,
		5, 14, 0, 0, 208, 209, 5, 21, 0, 0, 209, 210, 5, 44, 0, 0, 210, 219, 3,
		34, 17, 0, 211, 212, 5, 14, 0, 0, 212, 213, 3, 32, 16, 0, 213, 214, 5,
		47, 0, 0, 214, 215, 3, 52, 26, 0, 215, 216, 5, 44, 0, 0, 216, 217, 3, 34,
		17, 0, 217, 219, 1, 0, 0, 0, 218, 207, 1, 0, 0, 0, 218, 211, 1, 0, 0, 0,
		219, 31, 1, 0, 0, 0, 220, 225, 5, 21, 0, 0, 221, 222, 5, 54, 0, 0, 222,
		224, 5, 21, 0, 0, 223, 221, 1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223,
		1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 33, 1, 0, 0, 0, 227, 225, 1, 0,
		0, 0, 228, 229, 6, 17, -1, 0, 229, 253, 3, 46, 23, 0, 230, 253, 3, 20,
		10, 0, 231, 232, 5, 48, 0, 0, 232, 233, 3, 34, 17, 0, 233, 234, 5, 49,
		0, 0, 234, 253, 1, 0, 0, 0, 235, 236, 5, 50, 0, 0, 236, 237, 3, 34, 17,
		0, 237, 238, 5, 52, 0, 0, 238, 253, 1, 0, 0, 0, 239, 240, 7, 0, 0, 0, 240,
		253, 3, 34, 17, 12, 241, 253, 5, 21, 0, 0, 242, 253, 3, 56, 28, 0, 243,
		244, 5, 21, 0, 0, 244, 245, 5, 54, 0, 0, 245, 253, 5, 21, 0, 0, 246, 247,
		5, 21, 0, 0, 247, 248, 5, 54, 0, 0, 248, 253, 3, 34, 17, 2, 249, 250, 5,
		21, 0, 0, 250, 251, 5, 44, 0, 0, 251, 253, 3, 34, 17, 1, 252, 228, 1, 0,
		0, 0, 252, 230, 1, 0, 0, 0, 252, 231, 1, 0, 0, 0, 252, 235, 1, 0, 0, 0,
		252, 239, 1, 0, 0, 0, 252, 241, 1, 0, 0, 0, 252, 242, 1, 0, 0, 0, 252,
		243, 1, 0, 0, 0, 252, 246, 1, 0, 0, 0, 252, 249, 1, 0, 0, 0, 253, 274,
		1, 0, 0, 0, 254, 255, 10, 11, 0, 0, 255, 256, 7, 1, 0, 0, 256, 273, 3,
		34, 17, 12, 257, 258, 10, 10, 0, 0, 258, 259, 7, 2, 0, 0, 259, 273, 3,
		34, 17, 11, 260, 261, 10, 9, 0, 0, 261, 262, 7, 3, 0, 0, 262, 273, 3, 34,
		17, 10, 263, 264, 10, 8, 0, 0, 264, 265, 7, 4, 0, 0, 265, 273, 3, 34, 17,
		9, 266, 267, 10, 7, 0, 0, 267, 268, 5, 37, 0, 0, 268, 273, 3, 34, 17, 8,
		269, 270, 10, 6, 0, 0, 270, 271, 5, 36, 0, 0, 271, 273, 3, 34, 17, 7, 272,
		254, 1, 0, 0, 0, 272, 257, 1, 0, 0, 0, 272, 260, 1, 0, 0, 0, 272, 263,
		1, 0, 0, 0, 272, 266, 1, 0, 0, 0, 272, 269, 1, 0, 0, 0, 273, 276, 1, 0,
		0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 35, 1, 0, 0, 0,
		276, 274, 1, 0, 0, 0, 277, 282, 3, 38, 19, 0, 278, 279, 5, 55, 0, 0, 279,
		281, 3, 38, 19, 0, 280, 278, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0, 282, 280,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 37, 1, 0, 0, 0, 284, 282, 1, 0,
		0, 0, 285, 286, 3, 32, 16, 0, 286, 39, 1, 0, 0, 0, 287, 292, 3, 42, 21,
		0, 288, 289, 5, 55, 0, 0, 289, 291, 3, 42, 21, 0, 290, 288, 1, 0, 0, 0,
		291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		41, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296, 5, 21, 0, 0, 296, 297,
		3, 52, 26, 0, 297, 43, 1, 0, 0, 0, 298, 299, 3, 46, 23, 0, 299, 45, 1,
		0, 0, 0, 300, 307, 5, 16, 0, 0, 301, 307, 5, 18, 0, 0, 302, 307, 5, 19,
		0, 0, 303, 307, 5, 15, 0, 0, 304, 307, 5, 20, 0, 0, 305, 307, 5, 17, 0,
		0, 306, 300, 1, 0, 0, 0, 306, 301, 1, 0, 0, 0, 306, 302, 1, 0, 0, 0, 306,
		303, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 305, 1, 0, 0, 0, 307, 47, 1,
		0, 0, 0, 308, 309, 5, 4, 0, 0, 309, 310, 5, 21, 0, 0, 310, 314, 5, 51,
		0, 0, 311, 313, 3, 50, 25, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0,
		0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316,
		314, 1, 0, 0, 0, 317, 318, 5, 53, 0, 0, 318, 49, 1, 0, 0, 0, 319, 320,
		3, 52, 26, 0, 320, 321, 5, 21, 0, 0, 321, 322, 5, 44, 0, 0, 322, 323, 5,
		21, 0, 0, 323, 324, 5, 51, 0, 0, 324, 325, 5, 21, 0, 0, 325, 326, 5, 47,
		0, 0, 326, 327, 3, 34, 17, 0, 327, 328, 5, 55, 0, 0, 328, 329, 1, 0, 0,
		0, 329, 330, 5, 53, 0, 0, 330, 51, 1, 0, 0, 0, 331, 338, 5, 5, 0, 0, 332,
		338, 5, 6, 0, 0, 333, 338, 5, 7, 0, 0, 334, 338, 5, 8, 0, 0, 335, 338,
		5, 9, 0, 0, 336, 338, 5, 10, 0, 0, 337, 331, 1, 0, 0, 0, 337, 332, 1, 0,
		0, 0, 337, 333, 1, 0, 0, 0, 337, 334, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0,
		337, 336, 1, 0, 0, 0, 338, 53, 1, 0, 0, 0, 339, 340, 5, 50, 0, 0, 340,
		341, 5, 21, 0, 0, 341, 342, 5, 52, 0, 0, 342, 343, 5, 48, 0, 0, 343, 344,
		5, 49, 0, 0, 344, 55, 1, 0, 0, 0, 345, 346, 5, 21, 0, 0, 346, 350, 5, 45,
		0, 0, 347, 348, 5, 21, 0, 0, 348, 350, 5, 46, 0, 0, 349, 345, 1, 0, 0,
		0, 349, 347, 1, 0, 0, 0, 350, 57, 1, 0, 0, 0, 27, 61, 77, 84, 88, 96, 106,
		123, 134, 144, 150, 164, 172, 179, 190, 197, 201, 218, 225, 252, 272, 274,
		282, 292, 306, 314, 337, 349,
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

// VlangParserInit initializes any static state used to implement VlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.once.Do(vlangParserInit)
}

// NewVlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVlangParser(input antlr.TokenStream) *VlangParser {
	VlangParserInit()
	this := new(VlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Vlang.g4"

	return this
}

// VlangParser tokens.
const (
	VlangParserEOF           = antlr.TokenEOF
	VlangParserT__0          = 1
	VlangParserT__1          = 2
	VlangParserT__2          = 3
	VlangParserT__3          = 4
	VlangParserT__4          = 5
	VlangParserT__5          = 6
	VlangParserT__6          = 7
	VlangParserT__7          = 8
	VlangParserT__8          = 9
	VlangParserT__9          = 10
	VlangParserLEN           = 11
	VlangParserCAP           = 12
	VlangParserAPPEND        = 13
	VlangParserMUT           = 14
	VlangParserBOOLEANO      = 15
	VlangParserENTERO        = 16
	VlangParserFLOAT         = 17
	VlangParserDECIMAL       = 18
	VlangParserCADENA        = 19
	VlangParserCARACTER      = 20
	VlangParserID            = 21
	VlangParserIF_KW         = 22
	VlangParserELSE_KW       = 23
	VlangParserWHILE_KW      = 24
	VlangParserFOR_KW        = 25
	VlangParserIN_KW         = 26
	VlangParserRETURN_KW     = 27
	VlangParserBREAK_KW      = 28
	VlangParserCONTINUE_KW   = 29
	VlangParserPLUS          = 30
	VlangParserMINUS         = 31
	VlangParserMUL           = 32
	VlangParserDIV           = 33
	VlangParserMOD           = 34
	VlangParserNOT           = 35
	VlangParserOR            = 36
	VlangParserAND           = 37
	VlangParserEQ            = 38
	VlangParserNEQ           = 39
	VlangParserLT            = 40
	VlangParserLE            = 41
	VlangParserGT            = 42
	VlangParserGE            = 43
	VlangParserASSIGN        = 44
	VlangParserINC           = 45
	VlangParserDEC           = 46
	VlangParserCOLON         = 47
	VlangParserLPAREN        = 48
	VlangParserRPAREN        = 49
	VlangParserLBRACK        = 50
	VlangParserLCOR          = 51
	VlangParserRBRACK        = 52
	VlangParserRCOR          = 53
	VlangParserDOT           = 54
	VlangParserCOMMA         = 55
	VlangParserWS            = 56
	VlangParserLINE_COMMENT  = 57
	VlangParserBLOCK_COMMENT = 58
)

// VlangParser rules.
const (
	VlangParserRULE_programa        = 0
	VlangParserRULE_stmt            = 1
	VlangParserRULE_if_stmt         = 2
	VlangParserRULE_if_chain        = 3
	VlangParserRULE_else_stmt       = 4
	VlangParserRULE_println         = 5
	VlangParserRULE_print           = 6
	VlangParserRULE_while_stmt      = 7
	VlangParserRULE_for_stmt        = 8
	VlangParserRULE_range           = 9
	VlangParserRULE_func_call       = 10
	VlangParserRULE_func_dcl        = 11
	VlangParserRULE_struct_dcl      = 12
	VlangParserRULE_transfer_stmt   = 13
	VlangParserRULE_assign_stmt     = 14
	VlangParserRULE_decl_stmt       = 15
	VlangParserRULE_id_pattern      = 16
	VlangParserRULE_expresion       = 17
	VlangParserRULE_parametros      = 18
	VlangParserRULE_func_param      = 19
	VlangParserRULE_arg_list        = 20
	VlangParserRULE_func_arg        = 21
	VlangParserRULE_valores         = 22
	VlangParserRULE_valor           = 23
	VlangParserRULE_strct_instancia = 24
	VlangParserRULE_struct_prop     = 25
	VlangParserRULE_var_type        = 26
	VlangParserRULE_struct_vector   = 27
	VlangParserRULE_incredecre      = 28
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VlangParserEOF, 0)
}

func (s *ProgramaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VlangParserRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&996163614) != 0) {
		{
			p.SetState(58)
			p.Stmt()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(VlangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Assign_stmt() IAssign_stmtContext
	If_stmt() IIf_stmtContext
	While_stmt() IWhile_stmtContext
	For_stmt() IFor_stmtContext
	Func_call() IFunc_callContext
	Func_dcl() IFunc_dclContext
	Struct_dcl() IStruct_dclContext
	Strct_instancia() IStrct_instanciaContext
	Transfer_stmt() ITransfer_stmtContext
	Println_() IPrintlnContext
	Print_() IPrintContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *StmtContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StmtContext) Struct_dcl() IStruct_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_dclContext)
}

func (s *StmtContext) Strct_instancia() IStrct_instanciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrct_instanciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrct_instanciaContext)
}

func (s *StmtContext) Transfer_stmt() ITransfer_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_stmtContext)
}

func (s *StmtContext) Println_() IPrintlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnContext)
}

func (s *StmtContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VlangParserRULE_stmt)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Decl_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Assign_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.If_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.While_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.For_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.Func_call()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(71)
			p.Func_dcl()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(72)
			p.Struct_dcl()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(73)
			p.Strct_instancia()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(74)
			p.Transfer_stmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(75)
			p.Println_()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(76)
			p.Print_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStmtContext struct {
	If_stmtContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_chainContext)
}

func (s *IfStmtContext) AllELSE_KW() []antlr.TerminalNode {
	return s.GetTokens(VlangParserELSE_KW)
}

func (s *IfStmtContext) ELSE_KW(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserELSE_KW, i)
}

func (s *IfStmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VlangParserRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIfStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.If_chain()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(80)
				p.Match(VlangParserELSE_KW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(81)
				p.If_chain()
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserELSE_KW {
		{
			p.SetState(87)
			p.Else_stmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfChainContext struct {
	If_chainContext
}

func NewIfChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfChainContext {
	var p = new(IfChainContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IfChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfChainContext) IF_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserIF_KW, 0)
}

func (s *IfChainContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IfChainContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *IfChainContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *IfChainContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfChainContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IfChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIfChain(s)
	}
}

func (s *IfChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIfChain(s)
	}
}

func (s *IfChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIfChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VlangParserRULE_if_chain)
	var _la int

	localctx = NewIfChainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(VlangParserIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.expresion(0)
	}
	{
		p.SetState(92)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&996163614) != 0 {
		{
			p.SetState(93)
			p.Stmt()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ELSE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserELSE_KW, 0)
}

func (s *ElseStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ElseStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VlangParserRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(VlangParserELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&996163614) != 0 {
		{
			p.SetState(103)
			p.Stmt()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintlnContext is an interface to support dynamic dispatch.
type IPrintlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrintlnContext differentiates from other interfaces.
	IsPrintlnContext()
}

type PrintlnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnContext() *PrintlnContext {
	var p = new(PrintlnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_println
	return p
}

func InitEmptyPrintlnContext(p *PrintlnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_println
}

func (*PrintlnContext) IsPrintlnContext() {}

func NewPrintlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnContext {
	var p = new(PrintlnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_println

	return p
}

func (s *PrintlnContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnContext) CopyAll(ctx *PrintlnContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintlnStmtContext struct {
	PrintlnContext
}

func NewPrintlnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnStmtContext {
	var p = new(PrintlnStmtContext)

	InitEmptyPrintlnContext(&p.PrintlnContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrintlnContext))

	return p
}

func (s *PrintlnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *PrintlnStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PrintlnStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *PrintlnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrintlnStmt(s)
	}
}

func (s *PrintlnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrintlnStmt(s)
	}
}

func (s *PrintlnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrintlnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Println_() (localctx IPrintlnContext) {
	localctx = NewPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VlangParserRULE_println)
	localctx = NewPrintlnStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(VlangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.expresion(0)
	}
	{
		p.SetState(114)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) CopyAll(ctx *PrintContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintStmtContext struct {
	PrintContext
}

func NewPrintStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStmtContext {
	var p = new(PrintStmtContext)

	InitEmptyPrintContext(&p.PrintContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrintContext))

	return p
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *PrintStmtContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *PrintStmtContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *PrintStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *PrintStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VlangParserRULE_print)
	var _la int

	localctx = NewPrintStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(VlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.expresion(0)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(119)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.expresion(0)
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) WHILE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserWHILE_KW, 0)
}

func (s *WhileStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *WhileStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *WhileStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VlangParserRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(VlangParserWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.expresion(0)
	}
	{
		p.SetState(130)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&996163614) != 0 {
		{
			p.SetState(131)
			p.Stmt()
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(137)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR_KW, 0)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ForStmtContext) IN_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserIN_KW, 0)
}

func (s *ForStmtContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *ForStmtContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *ForStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VlangParserRULE_for_stmt)
	var _la int

	localctx = NewForStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(VlangParserFOR_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(VlangParserIN_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(142)
			p.expresion(0)
		}

	case 2:
		{
			p.SetState(143)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(146)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&996163614) != 0 {
		{
			p.SetState(147)
			p.Stmt()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumericRangeContext struct {
	RangeContext
}

func NewNumericRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericRangeContext {
	var p = new(NumericRangeContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *NumericRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericRangeContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *NumericRangeContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *NumericRangeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VlangParserDOT)
}

func (s *NumericRangeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, i)
}

func (s *NumericRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterNumericRange(s)
	}
}

func (s *NumericRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitNumericRange(s)
	}
}

func (s *NumericRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitNumericRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VlangParserRULE_range)
	localctx = NewNumericRangeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.expresion(0)
	}
	{
		p.SetState(156)
		p.Match(VlangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(VlangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(VlangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) CopyAll(ctx *Func_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	Func_callContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncCallContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VlangParserRULE_func_call)
	var _la int

	localctx = NewFuncCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Id_pattern()
	}
	{
		p.SetState(162)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserID {
		{
			p.SetState(163)
			p.Parametros()
		}

	}
	{
		p.SetState(166)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_dclContext is an interface to support dynamic dispatch.
type IFunc_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_dclContext differentiates from other interfaces.
	IsFunc_dclContext()
}

type Func_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_dclContext() *Func_dclContext {
	var p = new(Func_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_dcl
	return p
}

func InitEmptyFunc_dclContext(p *Func_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_dcl
}

func (*Func_dclContext) IsFunc_dclContext() {}

func NewFunc_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_dclContext {
	var p = new(Func_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_dcl

	return p
}

func (s *Func_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_dclContext) CopyAll(ctx *Func_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclContext struct {
	Func_dclContext
}

func NewFuncDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclContext {
	var p = new(FuncDeclContext)

	InitEmptyFunc_dclContext(&p.Func_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_dclContext))

	return p
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncDeclContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *FuncDeclContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *FuncDeclContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *FuncDeclContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncDeclContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_dcl() (localctx IFunc_dclContext) {
	localctx = NewFunc_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VlangParserRULE_func_dcl)
	var _la int

	localctx = NewFuncDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(VlangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserID {
		{
			p.SetState(171)
			p.Arg_list()
		}

	}
	{
		p.SetState(174)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&996163614) != 0 {
		{
			p.SetState(176)
			p.Stmt()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_dclContext is an interface to support dynamic dispatch.
type IStruct_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_dclContext differentiates from other interfaces.
	IsStruct_dclContext()
}

type Struct_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_dclContext() *Struct_dclContext {
	var p = new(Struct_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_dcl
	return p
}

func InitEmptyStruct_dclContext(p *Struct_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_dcl
}

func (*Struct_dclContext) IsStruct_dclContext() {}

func NewStruct_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_dclContext {
	var p = new(Struct_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_dcl

	return p
}

func (s *Struct_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_dclContext) CopyAll(ctx *Struct_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclContext struct {
	Struct_dclContext
}

func NewStructDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclContext {
	var p = new(StructDeclContext)

	InitEmptyStruct_dclContext(&p.Struct_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_dclContext))

	return p
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructDeclContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *StructDeclContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *StructDeclContext) AllAssign_stmt() []IAssign_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			len++
		}
	}

	tst := make([]IAssign_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssign_stmtContext); ok {
			tst[i] = t.(IAssign_stmtContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) Assign_stmt(i int) IAssign_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_dcl() (localctx IStruct_dclContext) {
	localctx = NewStruct_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VlangParserRULE_struct_dcl)
	var _la int

	localctx = NewStructDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(VlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserID {
		{
			p.SetState(187)
			p.Assign_stmt()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransfer_stmtContext is an interface to support dynamic dispatch.
type ITransfer_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransfer_stmtContext differentiates from other interfaces.
	IsTransfer_stmtContext()
}

type Transfer_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_stmtContext() *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_transfer_stmt
	return p
}

func InitEmptyTransfer_stmtContext(p *Transfer_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_transfer_stmt
}

func (*Transfer_stmtContext) IsTransfer_stmtContext() {}

func NewTransfer_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_transfer_stmt

	return p
}

func (s *Transfer_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_stmtContext) CopyAll(ctx *Transfer_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Transfer_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Transfer_stmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserCONTINUE_KW, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Transfer_stmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserBREAK_KW, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Transfer_stmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN_KW() antlr.TerminalNode {
	return s.GetToken(VlangParserRETURN_KW, 0)
}

func (s *ReturnStmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Transfer_stmt() (localctx ITransfer_stmtContext) {
	localctx = NewTransfer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VlangParserRULE_transfer_stmt)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserRETURN_KW:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Match(VlangParserRETURN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(196)
				p.expresion(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VlangParserBREAK_KW:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Match(VlangParserBREAK_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCONTINUE_KW:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Match(VlangParserCONTINUE_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) CopyAll(ctx *Assign_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DirectAssignContext struct {
	Assign_stmtContext
}

func NewDirectAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DirectAssignContext {
	var p = new(DirectAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *DirectAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *DirectAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *DirectAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DirectAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDirectAssign(s)
	}
}

func (s *DirectAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDirectAssign(s)
	}
}

func (s *DirectAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDirectAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VlangParserRULE_assign_stmt)
	localctx = NewDirectAssignContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Id_pattern()
	}
	{
		p.SetState(204)
		p.Match(VlangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) CopyAll(ctx *Decl_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclAssignContext struct {
	Decl_stmtContext
}

func NewDeclAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAssignContext {
	var p = new(DeclAssignContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAssignContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *DeclAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DeclAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *DeclAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclAssign(s)
	}
}

func (s *DeclAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclAssign(s)
	}
}

func (s *DeclAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclAssignTypeContext struct {
	Decl_stmtContext
}

func NewDeclAssignTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAssignTypeContext {
	var p = new(DeclAssignTypeContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclAssignTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAssignTypeContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *DeclAssignTypeContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *DeclAssignTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *DeclAssignTypeContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *DeclAssignTypeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *DeclAssignTypeContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclAssignTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclAssignType(s)
	}
}

func (s *DeclAssignTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclAssignType(s)
	}
}

func (s *DeclAssignTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclAssignType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VlangParserRULE_decl_stmt)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(VlangParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.expresion(0)
		}

	case 2:
		localctx = NewDeclAssignTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Match(VlangParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Id_pattern()
		}
		{
			p.SetState(213)
			p.Match(VlangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Var_type()
		}
		{
			p.SetState(215)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IId_patternContext is an interface to support dynamic dispatch.
type IId_patternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsId_patternContext differentiates from other interfaces.
	IsId_patternContext()
}

type Id_patternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_patternContext() *Id_patternContext {
	var p = new(Id_patternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_id_pattern
	return p
}

func InitEmptyId_patternContext(p *Id_patternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_id_pattern
}

func (*Id_patternContext) IsId_patternContext() {}

func NewId_patternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_patternContext {
	var p = new(Id_patternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_id_pattern

	return p
}

func (s *Id_patternContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_patternContext) CopyAll(ctx *Id_patternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Id_patternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_patternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdPatternContext struct {
	Id_patternContext
}

func NewIdPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdPatternContext {
	var p = new(IdPatternContext)

	InitEmptyId_patternContext(&p.Id_patternContext)
	p.parser = parser
	p.CopyAll(ctx.(*Id_patternContext))

	return p
}

func (s *IdPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdPatternContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *IdPatternContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *IdPatternContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VlangParserDOT)
}

func (s *IdPatternContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, i)
}

func (s *IdPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIdPattern(s)
	}
}

func (s *IdPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIdPattern(s)
	}
}

func (s *IdPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIdPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Id_pattern() (localctx IId_patternContext) {
	localctx = NewId_patternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VlangParserRULE_id_pattern)
	var _la int

	localctx = NewIdPatternContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserDOT {
		{
			p.SetState(221)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CorchetesexpreContext struct {
	ExpresionContext
}

func NewCorchetesexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CorchetesexpreContext {
	var p = new(CorchetesexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *CorchetesexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CorchetesexpreContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *CorchetesexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CorchetesexpreContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *CorchetesexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCorchetesexpre(s)
	}
}

func (s *CorchetesexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCorchetesexpre(s)
	}
}

func (s *CorchetesexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCorchetesexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncionexpreContext struct {
	ExpresionContext
}

func NewFuncionexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionexpreContext {
	var p = new(FuncionexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *FuncionexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionexpreContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *FuncionexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncionexpre(s)
	}
}

func (s *FuncionexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncionexpre(s)
	}
}

func (s *FuncionexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncionexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioContext struct {
	ExpresionContext
	op antlr.Token
}

func NewUnarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioContext {
	var p = new(UnarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *UnarioContext) GetOp() antlr.Token { return s.op }

func (s *UnarioContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *UnarioContext) NOT() antlr.TerminalNode {
	return s.GetToken(VlangParserNOT, 0)
}

func (s *UnarioContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *UnarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterUnario(s)
	}
}

func (s *UnarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitUnario(s)
	}
}

func (s *UnarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitUnario(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentesisexpreContext struct {
	ExpresionContext
}

func NewParentesisexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentesisexpreContext {
	var p = new(ParentesisexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ParentesisexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentesisexpreContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ParentesisexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParentesisexpreContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ParentesisexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParentesisexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncredecrContext struct {
	ExpresionContext
}

func NewIncredecrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncredecrContext {
	var p = new(IncredecrContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IncredecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecrContext) Incredecre() IIncredecreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncredecreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncredecreContext)
}

func (s *IncredecrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncredecr(s)
	}
}

func (s *IncredecrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncredecr(s)
	}
}

func (s *IncredecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncredecr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorexpresionContext struct {
	ExpresionContext
}

func NewValorexpresionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorexpresionContext {
	var p = new(ValorexpresionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ValorexpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorexpresionContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValorexpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorexpresion(s)
	}
}

func (s *ValorexpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorexpresion(s)
	}
}

func (s *ValorexpresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorexpresion(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionforContext struct {
	ExpresionContext
}

func NewAsignacionforContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionforContext {
	var p = new(AsignacionforContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionforContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionforContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionforContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionforContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionfor(s)
	}
}

func (s *AsignacionforContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionfor(s)
	}
}

func (s *AsignacionforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionfor(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExpresionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expdotexp1Context struct {
	ExpresionContext
}

func NewExpdotexp1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expdotexp1Context {
	var p = new(Expdotexp1Context)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expdotexp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expdotexp1Context) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *Expdotexp1Context) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *Expdotexp1Context) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *Expdotexp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp1(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpdotexpContext struct {
	ExpresionContext
}

func NewExpdotexpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpdotexpContext {
	var p = new(ExpdotexpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExpdotexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpdotexpContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ExpdotexpContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *ExpdotexpContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpdotexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp(s)
	}
}

func (s *ExpdotexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp(s)
	}
}

func (s *ExpdotexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewBinaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpContext {
	var p = new(BinaryExpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *BinaryExpContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpContext) GetLeft() IExpresionContext { return s.left }

func (s *BinaryExpContext) GetRight() IExpresionContext { return s.right }

func (s *BinaryExpContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *BinaryExpContext) SetRight(v IExpresionContext) { s.right = v }

func (s *BinaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *BinaryExpContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *BinaryExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(VlangParserDIV, 0)
}

func (s *BinaryExpContext) MOD() antlr.TerminalNode {
	return s.GetToken(VlangParserMOD, 0)
}

func (s *BinaryExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUS, 0)
}

func (s *BinaryExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *BinaryExpContext) LT() antlr.TerminalNode {
	return s.GetToken(VlangParserLT, 0)
}

func (s *BinaryExpContext) LE() antlr.TerminalNode {
	return s.GetToken(VlangParserLE, 0)
}

func (s *BinaryExpContext) GT() antlr.TerminalNode {
	return s.GetToken(VlangParserGT, 0)
}

func (s *BinaryExpContext) GE() antlr.TerminalNode {
	return s.GetToken(VlangParserGE, 0)
}

func (s *BinaryExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *BinaryExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserNEQ, 0)
}

func (s *BinaryExpContext) AND() antlr.TerminalNode {
	return s.GetToken(VlangParserAND, 0)
}

func (s *BinaryExpContext) OR() antlr.TerminalNode {
	return s.GetToken(VlangParserOR, 0)
}

func (s *BinaryExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBinaryExp(s)
	}
}

func (s *BinaryExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBinaryExp(s)
	}
}

func (s *BinaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBinaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *VlangParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, VlangParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValorexpresionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(229)
			p.Valor()
		}

	case 2:
		localctx = NewFuncionexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.Func_call()
		}

	case 3:
		localctx = NewParentesisexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.expresion(0)
		}
		{
			p.SetState(233)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewCorchetesexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.expresion(0)
		}
		{
			p.SetState(237)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserMINUS || _la == VlangParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(240)
			p.expresion(12)
		}

	case 6:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(241)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIncredecrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Incredecre()
		}

	case 8:
		localctx = NewExpdotexp1Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewExpdotexpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(246)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.expresion(2)
		}

	case 10:
		localctx = NewAsignacionforContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(249)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.expresion(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(255)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30064771072) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(256)

					var _x = p.expresion(12)

					localctx.(*BinaryExpContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(258)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserPLUS || _la == VlangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(259)

					var _x = p.expresion(11)

					localctx.(*BinaryExpContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(261)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16492674416640) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(262)

					var _x = p.expresion(10)

					localctx.(*BinaryExpContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(264)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserEQ || _la == VlangParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(265)

					var _x = p.expresion(9)

					localctx.(*BinaryExpContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(267)

					var _m = p.Match(VlangParserAND)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(268)

					var _x = p.expresion(8)

					localctx.(*BinaryExpContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(270)

					var _m = p.Match(VlangParserOR)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(271)

					var _x = p.expresion(7)

					localctx.(*BinaryExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) CopyAll(ctx *ParametrosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamListContext struct {
	ParametrosContext
}

func NewParamListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListContext {
	var p = new(ParamListContext)

	InitEmptyParametrosContext(&p.ParametrosContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParametrosContext))

	return p
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) AllFunc_param() []IFunc_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunc_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_paramContext); ok {
			tst[i] = t.(IFunc_paramContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Func_param(i int) IFunc_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_paramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_paramContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VlangParserRULE_parametros)
	var _la int

	localctx = NewParamListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Func_param()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(278)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Func_param()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_paramContext is an interface to support dynamic dispatch.
type IFunc_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_paramContext differentiates from other interfaces.
	IsFunc_paramContext()
}

type Func_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_paramContext() *Func_paramContext {
	var p = new(Func_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_param
	return p
}

func InitEmptyFunc_paramContext(p *Func_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_param
}

func (*Func_paramContext) IsFunc_paramContext() {}

func NewFunc_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_paramContext {
	var p = new(Func_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_param

	return p
}

func (s *Func_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_paramContext) CopyAll(ctx *Func_paramContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncParamContext struct {
	Func_paramContext
}

func NewFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamContext {
	var p = new(FuncParamContext)

	InitEmptyFunc_paramContext(&p.Func_paramContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_paramContext))

	return p
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_param() (localctx IFunc_paramContext) {
	localctx = NewFunc_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VlangParserRULE_func_param)
	localctx = NewFuncParamContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Id_pattern()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) CopyAll(ctx *Arg_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgListContext struct {
	Arg_listContext
}

func NewArgListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgListContext {
	var p = new(ArgListContext)

	InitEmptyArg_listContext(&p.Arg_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arg_listContext))

	return p
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) AllFunc_arg() []IFunc_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_argContext); ok {
			len++
		}
	}

	tst := make([]IFunc_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_argContext); ok {
			tst[i] = t.(IFunc_argContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Func_arg(i int) IFunc_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_argContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_argContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VlangParserRULE_arg_list)
	var _la int

	localctx = NewArgListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Func_arg()
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(288)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Func_arg()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_argContext is an interface to support dynamic dispatch.
type IFunc_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_argContext differentiates from other interfaces.
	IsFunc_argContext()
}

type Func_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_argContext() *Func_argContext {
	var p = new(Func_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_arg
	return p
}

func InitEmptyFunc_argContext(p *Func_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_func_arg
}

func (*Func_argContext) IsFunc_argContext() {}

func NewFunc_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_argContext {
	var p = new(Func_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_func_arg

	return p
}

func (s *Func_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_argContext) CopyAll(ctx *Func_argContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgContext struct {
	Func_argContext
}

func NewFuncArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgContext {
	var p = new(FuncArgContext)

	InitEmptyFunc_argContext(&p.Func_argContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_argContext))

	return p
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncArgContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Func_arg() (localctx IFunc_argContext) {
	localctx = NewFunc_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VlangParserRULE_func_arg)
	localctx = NewFuncArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Var_type()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Valor() IValorContext

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValoresContext() *ValoresContext {
	var p = new(ValoresContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
	return p
}

func InitEmptyValoresContext(p *ValoresContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
}

func (*ValoresContext) IsValoresContext() {}

func NewValoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValoresContext {
	var p = new(ValoresContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valores

	return p
}

func (s *ValoresContext) GetParser() antlr.Parser { return s.parser }

func (s *ValoresContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValores(s)
	}
}

func (s *ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValores(s)
	}
}

func (s *ValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValores(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valores() (localctx IValoresContext) {
	localctx = NewValoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VlangParserRULE_valores)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Valor()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValorContext() *ValorContext {
	var p = new(ValorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
	return p
}

func InitEmptyValorContext(p *ValorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
}

func (*ValorContext) IsValorContext() {}

func NewValorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValorContext {
	var p = new(ValorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valor

	return p
}

func (s *ValorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValorContext) CopyAll(ctx *ValorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValorDecimalContext struct {
	ValorContext
}

func NewValorDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorDecimalContext {
	var p = new(ValorDecimalContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorDecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(VlangParserDECIMAL, 0)
}

func (s *ValorDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorDecimal(s)
	}
}

func (s *ValorDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorDecimal(s)
	}
}

func (s *ValorDecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorEnteroContext struct {
	ValorContext
}

func NewValorEnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorEnteroContext {
	var p = new(ValorEnteroContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorEnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorEnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(VlangParserENTERO, 0)
}

func (s *ValorEnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorEntero(s)
	}
}

func (s *ValorEnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorEntero(s)
	}
}

func (s *ValorEnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorEntero(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorBooleanoContext struct {
	ValorContext
}

func NewValorBooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorBooleanoContext {
	var p = new(ValorBooleanoContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorBooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorBooleanoContext) BOOLEANO() antlr.TerminalNode {
	return s.GetToken(VlangParserBOOLEANO, 0)
}

func (s *ValorBooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorBooleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCaracterContext struct {
	ValorContext
}

func NewValorCaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCaracterContext {
	var p = new(ValorCaracterContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(VlangParserCARACTER, 0)
}

func (s *ValorCaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCaracter(s)
	}
}

func (s *ValorCaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCaracter(s)
	}
}

func (s *ValorCaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCaracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCadenaContext struct {
	ValorContext
}

func NewValorCadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCadenaContext {
	var p = new(ValorCadenaContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(VlangParserCADENA, 0)
}

func (s *ValorCadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCadena(s)
	}
}

func (s *ValorCadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCadena(s)
	}
}

func (s *ValorCadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCadena(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorFloatContext struct {
	ValorContext
}

func NewValorFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorFloatContext {
	var p = new(ValorFloatContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(VlangParserFLOAT, 0)
}

func (s *ValorFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorFloat(s)
	}
}

func (s *ValorFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorFloat(s)
	}
}

func (s *ValorFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valor() (localctx IValorContext) {
	localctx = NewValorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VlangParserRULE_valor)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserENTERO:
		localctx = NewValorEnteroContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.Match(VlangParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserDECIMAL:
		localctx = NewValorDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(301)
			p.Match(VlangParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCADENA:
		localctx = NewValorCadenaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(302)
			p.Match(VlangParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserBOOLEANO:
		localctx = NewValorBooleanoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(303)
			p.Match(VlangParserBOOLEANO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCARACTER:
		localctx = NewValorCaracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(304)
			p.Match(VlangParserCARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserFLOAT:
		localctx = NewValorFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(305)
			p.Match(VlangParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrct_instanciaContext is an interface to support dynamic dispatch.
type IStrct_instanciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStrct_instanciaContext differentiates from other interfaces.
	IsStrct_instanciaContext()
}

type Strct_instanciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrct_instanciaContext() *Strct_instanciaContext {
	var p = new(Strct_instanciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_strct_instancia
	return p
}

func InitEmptyStrct_instanciaContext(p *Strct_instanciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_strct_instancia
}

func (*Strct_instanciaContext) IsStrct_instanciaContext() {}

func NewStrct_instanciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Strct_instanciaContext {
	var p = new(Strct_instanciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_strct_instancia

	return p
}

func (s *Strct_instanciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Strct_instanciaContext) CopyAll(ctx *Strct_instanciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Strct_instanciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Strct_instanciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructInstanciaContext struct {
	Strct_instanciaContext
}

func NewStructInstanciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstanciaContext {
	var p = new(StructInstanciaContext)

	InitEmptyStrct_instanciaContext(&p.Strct_instanciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Strct_instanciaContext))

	return p
}

func (s *StructInstanciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanciaContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructInstanciaContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *StructInstanciaContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *StructInstanciaContext) AllStruct_prop() []IStruct_propContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_propContext); ok {
			len++
		}
	}

	tst := make([]IStruct_propContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_propContext); ok {
			tst[i] = t.(IStruct_propContext)
			i++
		}
	}

	return tst
}

func (s *StructInstanciaContext) Struct_prop(i int) IStruct_propContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_propContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_propContext)
}

func (s *StructInstanciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructInstancia(s)
	}
}

func (s *StructInstanciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructInstancia(s)
	}
}

func (s *StructInstanciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructInstancia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Strct_instancia() (localctx IStrct_instanciaContext) {
	localctx = NewStrct_instanciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VlangParserRULE_strct_instancia)
	var _la int

	localctx = NewStructInstanciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(VlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2016) != 0 {
		{
			p.SetState(311)
			p.Struct_prop()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(317)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_propContext is an interface to support dynamic dispatch.
type IStruct_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_propContext differentiates from other interfaces.
	IsStruct_propContext()
}

type Struct_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_propContext() *Struct_propContext {
	var p = new(Struct_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_prop
	return p
}

func InitEmptyStruct_propContext(p *Struct_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_prop
}

func (*Struct_propContext) IsStruct_propContext() {}

func NewStruct_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_propContext {
	var p = new(Struct_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_prop

	return p
}

func (s *Struct_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_propContext) CopyAll(ctx *Struct_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructAttrContext struct {
	Struct_propContext
}

func NewStructAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttrContext {
	var p = new(StructAttrContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *StructAttrContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *StructAttrContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *StructAttrContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *StructAttrContext) LCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserLCOR, 0)
}

func (s *StructAttrContext) RCOR() antlr.TerminalNode {
	return s.GetToken(VlangParserRCOR, 0)
}

func (s *StructAttrContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *StructAttrContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *StructAttrContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *StructAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructAttr(s)
	}
}

func (s *StructAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructAttr(s)
	}
}

func (s *StructAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_prop() (localctx IStruct_propContext) {
	localctx = NewStruct_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VlangParserRULE_struct_prop)
	localctx = NewStructAttrContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Var_type()
	}
	{
		p.SetState(320)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(VlangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(VlangParserLCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(324)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.expresion(0)
	}
	{
		p.SetState(327)
		p.Match(VlangParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(329)
		p.Match(VlangParserRCOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) CopyAll(ctx *Var_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VoidTypeContext struct {
	Var_typeContext
}

func NewVoidTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidTypeContext {
	var p = new(VoidTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *VoidTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterVoidType(s)
	}
}

func (s *VoidTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitVoidType(s)
	}
}

func (s *VoidTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitVoidType(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolTypeContext struct {
	Var_typeContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringTypeContext struct {
	Var_typeContext
}

func NewStringTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringTypeContext {
	var p = new(StringTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharTypeContext struct {
	Var_typeContext
}

func NewCharTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharTypeContext {
	var p = new(CharTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *CharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCharType(s)
	}
}

func (s *CharTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCharType(s)
	}
}

func (s *CharTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCharType(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntTypeContext struct {
	Var_typeContext
}

func NewIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTypeContext {
	var p = new(IntTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatTypeContext struct {
	Var_typeContext
}

func NewFloatTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatTypeContext {
	var p = new(FloatTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *FloatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFloatType(s)
	}
}

func (s *FloatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFloatType(s)
	}
}

func (s *FloatTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFloatType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VlangParserRULE_var_type)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserT__4:
		localctx = NewIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(VlangParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__5:
		localctx = NewFloatTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(VlangParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__6:
		localctx = NewStringTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(333)
			p.Match(VlangParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__7:
		localctx = NewBoolTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(334)
			p.Match(VlangParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__8:
		localctx = NewCharTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(335)
			p.Match(VlangParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__9:
		localctx = NewVoidTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(336)
			p.Match(VlangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_vectorContext is an interface to support dynamic dispatch.
type IStruct_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_vectorContext differentiates from other interfaces.
	IsStruct_vectorContext()
}

type Struct_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_vectorContext() *Struct_vectorContext {
	var p = new(Struct_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_vector
	return p
}

func InitEmptyStruct_vectorContext(p *Struct_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_struct_vector
}

func (*Struct_vectorContext) IsStruct_vectorContext() {}

func NewStruct_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_vectorContext {
	var p = new(Struct_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_struct_vector

	return p
}

func (s *Struct_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_vectorContext) CopyAll(ctx *Struct_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructVectorContext struct {
	Struct_vectorContext
}

func NewStructVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVectorContext {
	var p = new(StructVectorContext)

	InitEmptyStruct_vectorContext(&p.Struct_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_vectorContext))

	return p
}

func (s *StructVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVectorContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *StructVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructVectorContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *StructVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *StructVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *StructVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructVector(s)
	}
}

func (s *StructVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructVector(s)
	}
}

func (s *StructVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Struct_vector() (localctx IStruct_vectorContext) {
	localctx = NewStruct_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VlangParserRULE_struct_vector)
	localctx = NewStructVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Match(VlangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncredecreContext is an interface to support dynamic dispatch.
type IIncredecreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncredecreContext differentiates from other interfaces.
	IsIncredecreContext()
}

type IncredecreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncredecreContext() *IncredecreContext {
	var p = new(IncredecreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
	return p
}

func InitEmptyIncredecreContext(p *IncredecreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
}

func (*IncredecreContext) IsIncredecreContext() {}

func NewIncredecreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncredecreContext {
	var p = new(IncredecreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_incredecre

	return p
}

func (s *IncredecreContext) GetParser() antlr.Parser { return s.parser }

func (s *IncredecreContext) CopyAll(ctx *IncredecreContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IncredecreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IncrementoContext struct {
	IncredecreContext
}

func NewIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementoContext {
	var p = new(IncrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *IncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IncrementoContext) INC() antlr.TerminalNode {
	return s.GetToken(VlangParserINC, 0)
}

func (s *IncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncremento(s)
	}
}

func (s *IncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncremento(s)
	}
}

func (s *IncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecrementoContext struct {
	IncredecreContext
}

func NewDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementoContext {
	var p = new(DecrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *DecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DecrementoContext) DEC() antlr.TerminalNode {
	return s.GetToken(VlangParserDEC, 0)
}

func (s *DecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDecremento(s)
	}
}

func (s *DecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDecremento(s)
	}
}

func (s *DecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Incredecre() (localctx IIncredecreContext) {
	localctx = NewIncredecreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VlangParserRULE_incredecre)
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(VlangParserINC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Match(VlangParserDEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VlangParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
