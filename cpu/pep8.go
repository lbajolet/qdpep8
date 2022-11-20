package cpu

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Sign int

const (
	Positive Sign = iota
	Negative
)

type opcode uint8

const (
	STOP       opcode = 0
	RETTR             = 1
	MOVSPA            = 2
	MOVFLGA           = 3
	BRi               = 4
	BRx               = 5
	BRLEi             = 6
	BRLEx             = 7
	BRLTi             = 8
	BRLTx             = 9
	BREQi             = 10
	BREQx             = 11
	BRNEi             = 12
	BRNEx             = 13
	BRGEi             = 14
	BRGEx             = 15
	BRGTi             = 16
	BRGTx             = 17
	BRVi              = 18
	BRVx              = 19
	BRCi              = 20
	BRCx              = 21
	CALLi             = 22
	CALLx             = 23
	NOTA              = 24
	NOTX              = 25
	NEGA              = 26
	NEGX              = 27
	ASLA              = 28
	ASLX              = 29
	ASRA              = 30
	ASRX              = 31
	ROLA              = 32
	ROLX              = 33
	RORA              = 34
	RORX              = 35
	NOP0              = 36
	NOP1              = 37
	NOP2              = 38
	NOP3              = 39
	NOPi              = 40
	NOPd              = 41
	NOPn              = 42
	NOPs              = 43
	NOPsf             = 44
	NOPx              = 45
	NOPsx             = 46
	NOPsxf            = 47
	DECIi             = 48
	DECId             = 49
	DECIn             = 50
	DECIs             = 51
	DECIsf            = 52
	DECIx             = 53
	DECIsx            = 54
	DECIsxf           = 55
	DECOi             = 56
	DECOd             = 57
	DECOn             = 58
	DECOs             = 59
	DECOsf            = 60
	DECOx             = 61
	DECOsx            = 62
	DECOsxf           = 63
	STROi             = 64
	STROd             = 65
	STROn             = 66
	STROs             = 67
	STROsf            = 68
	STROx             = 69
	STROsx            = 70
	STROsxf           = 71
	CHARIi            = 72
	CHARId            = 73
	CHARIn            = 74
	CHARIs            = 75
	CHARIsf           = 76
	CHARIx            = 77
	CHARIsx           = 78
	CHARIsxf          = 79
	CHAROi            = 80
	CHAROd            = 81
	CHAROn            = 82
	CHAROs            = 83
	CHAROsf           = 84
	CHAROx            = 85
	CHAROsx           = 86
	CHAROsxf          = 87
	RET0              = 88
	RET1              = 89
	RET2              = 90
	RET3              = 91
	RET4              = 92
	RET5              = 93
	RET6              = 94
	RET7              = 95
	ADDSPi            = 96
	ADDSPd            = 97
	ADDSPn            = 98
	ADDSPs            = 99
	ADDSPsf           = 100
	ADDSPx            = 101
	ADDSPsx           = 102
	ADDSPsxf          = 103
	SUBSPi            = 104
	SUBSPd            = 105
	SUBSPn            = 106
	SUBSPs            = 107
	SUBSPsf           = 108
	SUBSPx            = 109
	SUBSPsx           = 110
	SUBSPsxf          = 111
	ADDAi             = 112
	ADDAd             = 113
	ADDAn             = 114
	ADDAs             = 115
	ADDAsf            = 116
	ADDAx             = 117
	ADDAsx            = 118
	ADDAsxf           = 119
	ADDXi             = 120
	ADDXd             = 121
	ADDXn             = 122
	ADDXs             = 123
	ADDXsf            = 124
	ADDXx             = 125
	ADDXsx            = 126
	ADDXsxf           = 127
	SUBAi             = 128
	SUBAd             = 129
	SUBAn             = 130
	SUBAs             = 131
	SUBAsf            = 132
	SUBAx             = 133
	SUBAsx            = 134
	SUBAsxf           = 135
	SUBXi             = 136
	SUBXd             = 137
	SUBXn             = 138
	SUBXs             = 139
	SUBXsf            = 140
	SUBXx             = 141
	SUBXsx            = 142
	SUBXsxf           = 143
	ANDAi             = 144
	ANDAd             = 145
	ANDAn             = 146
	ANDAs             = 147
	ANDAsf            = 148
	ANDAx             = 149
	ANDAsx            = 150
	ANDAsxf           = 151
	ANDXi             = 152
	ANDXd             = 153
	ANDXn             = 154
	ANDXs             = 155
	ANDXsf            = 156
	ANDXx             = 157
	ANDXsx            = 158
	ANDXsxf           = 159
	ORAi              = 160
	ORAd              = 161
	ORAn              = 162
	ORAs              = 163
	ORAsf             = 164
	ORAx              = 165
	ORAsx             = 166
	ORAsxf            = 167
	ORXi              = 168
	ORXd              = 169
	ORXn              = 170
	ORXs              = 171
	ORXsf             = 172
	ORXx              = 173
	ORXsx             = 174
	ORXsxf            = 175
	CPAi              = 176
	CPAd              = 177
	CPAn              = 178
	CPAs              = 179
	CPAsf             = 180
	CPAx              = 181
	CPAsx             = 182
	CPAsxf            = 183
	CPXi              = 184
	CPXd              = 185
	CPXn              = 186
	CPXs              = 187
	CPXsf             = 188
	CPXx              = 189
	CPXsx             = 190
	CPXsxf            = 191
	LDAi              = 192
	LDAd              = 193
	LDAn              = 194
	LDAs              = 195
	LDAsf             = 196
	LDAx              = 197
	LDAsx             = 198
	LDAsxf            = 199
	LDXi              = 200
	LDXd              = 201
	LDXn              = 202
	LDXs              = 203
	LDXsf             = 204
	LDXx              = 205
	LDXsx             = 206
	LDXsxf            = 207
	LDBYTEAi          = 208
	LDBYTEAd          = 209
	LDBYTEAn          = 210
	LDBYTEAs          = 211
	LDBYTEAsf         = 212
	LDBYTEAx          = 213
	LDBYTEAsx         = 214
	LDBYTEAsxf        = 215
	LDBYTEXi          = 216
	LDBYTEXd          = 217
	LDBYTEXn          = 218
	LDBYTEXs          = 219
	LDBYTEXsf         = 220
	LDBYTEXx          = 221
	LDBYTEXsx         = 222
	LDBYTEXsxf        = 223
	STAi              = 224
	STAd              = 225
	STAn              = 226
	STAs              = 227
	STAsf             = 228
	STAx              = 229
	STAsx             = 230
	STAsxf            = 231
	STXi              = 232
	STXd              = 233
	STXn              = 234
	STXs              = 235
	STXsf             = 236
	STXx              = 237
	STXsx             = 238
	STXsxf            = 239
	STBYTEAi          = 240
	STBYTEAd          = 241
	STBYTEAn          = 242
	STBYTEAs          = 243
	STBYTEAsf         = 244
	STBYTEAx          = 245
	STBYTEAsx         = 246
	STBYTEAsxf        = 247
	STBYTEXi          = 248
	STBYTEXd          = 249
	STBYTEXn          = 250
	STBYTEXs          = 251
	STBYTEXsf         = 252
	STBYTEXx          = 253
	STBYTEXsx         = 254
	STBYTEXsxf        = 255
)

func (oc opcode) BaseOp() string {
	switch oc {
	case STOP:
		return "STOP"
	case RETTR:
		return "RETTR"
	case MOVSPA:
		return "MOVSPA"
	case MOVFLGA:
		return "MOVFLGA"
	case BRi, BRx:
		return "BR"
	case BRLEi, BRLEx:
		return "BRLE"
	case BRLTi, BRLTx:
		return "BRLT"
	case BREQi, BREQx:
		return "BREQ"
	case BRNEi, BRNEx:
		return "BRNE"
	case BRGEi, BRGEx:
		return "BRGE"
	case BRGTi, BRGTx:
		return "BRGT"
	case BRVi, BRVx:
		return "BRV"
	case BRCi, BRCx:
		return "BRC"
	case CALLi, CALLx:
		return "CALL"
	case NOTA, NOTX:
		return "NOT"
	case NEGA, NEGX:
		return "NEG"
	case ASLA, ASLX:
		return "ASL"
	case ASRA, ASRX:
		return "ASR"
	case ROLA, ROLX:
		return "ROL"
	case RORA, RORX:
		return "ROR"
	case NOP0, NOP1, NOP2, NOP3, NOPi, NOPd, NOPn, NOPs, NOPsf, NOPx, NOPsx, NOPsxf:
		return "NOP"
	case DECIi, DECId, DECIn, DECIs, DECIsf, DECIx, DECIsx, DECIsxf:
		return "DECI"
	case DECOi, DECOd, DECOn, DECOs, DECOsf, DECOx, DECOsx, DECOsxf:
		return "DECO"
	case STROi, STROd, STROn, STROs, STROsf, STROx, STROsx, STROsxf:
		return "STRO"
	case CHARIi, CHARId, CHARIn, CHARIs, CHARIsf, CHARIx, CHARIsx, CHARIsxf:
		return "CHARI"
	case CHAROi, CHAROd, CHAROn, CHAROs, CHAROsf, CHAROx, CHAROsx, CHAROsxf:
		return "CHARO"
	case RET0, RET1, RET2, RET3, RET4, RET5, RET6, RET7:
		return "RET"
	case ADDSPi, ADDSPd, ADDSPn, ADDSPs, ADDSPsf, ADDSPx, ADDSPsx, ADDSPsxf:
		return "ADDSP"
	case SUBSPi, SUBSPd, SUBSPn, SUBSPs, SUBSPsf, SUBSPx, SUBSPsx, SUBSPsxf:
		return "SUBSP"
	case ADDAi, ADDAd, ADDAn, ADDAs, ADDAsf, ADDAx, ADDAsx, ADDAsxf,
		ADDXi, ADDXd, ADDXn, ADDXs, ADDXsf, ADDXx, ADDXsx, ADDXsxf:
		return "ADD"
	case SUBAi, SUBAd, SUBAn, SUBAs, SUBAsf, SUBAx, SUBAsx, SUBAsxf,
		SUBXi, SUBXd, SUBXn, SUBXs, SUBXsf, SUBXx, SUBXsx, SUBXsxf:
		return "SUB"
	case ANDAi, ANDAd, ANDAn, ANDAs, ANDAsf, ANDAx, ANDAsx, ANDAsxf,
		ANDXi, ANDXd, ANDXn, ANDXs, ANDXsf, ANDXx, ANDXsx, ANDXsxf:
		return "AND"
	case ORAi, ORAd, ORAn, ORAs, ORAsf, ORAx, ORAsx, ORAsxf,
		ORXi, ORXd, ORXn, ORXs, ORXsf, ORXx, ORXsx, ORXsxf:
		return "OR"
	case CPAi, CPAd, CPAn, CPAs, CPAsf, CPAx, CPAsx, CPAsxf,
		CPXi, CPXd, CPXn, CPXs, CPXsf, CPXx, CPXsx, CPXsxf:
		return "CP"
	case LDAi, LDAd, LDAn, LDAs, LDAsf, LDAx, LDAsx, LDAsxf,
		LDXi, LDXd, LDXn, LDXs, LDXsf, LDXx, LDXsx, LDXsxf:
		return "LD"
	case LDBYTEAi, LDBYTEAd, LDBYTEAn, LDBYTEAs, LDBYTEAsf, LDBYTEAx, LDBYTEAsx, LDBYTEAsxf,
		LDBYTEXi, LDBYTEXd, LDBYTEXn, LDBYTEXs, LDBYTEXsf, LDBYTEXx, LDBYTEXsx, LDBYTEXsxf:
		return "LDBYTE"
	case STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf:
		return "ST"
	case STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf:
		return "STBYTE"
	}
	panic("Unknown opcode")
}

func (oc opcode) hasAddr() bool {
	switch oc {
	case BRLEi, BRLEx,
		BRLTi, BRLTx,
		BREQi, BREQx,
		BRNEi, BRNEx,
		BRGEi, BRGEx,
		BRGTi, BRGTx,
		BRVi, BRVx,
		BRCi, BRCx,
		CALLi, CALLx,
		NOPi, NOPd, NOPn, NOPs, NOPsf, NOPx, NOPsx, NOPsxf,
		DECIi, DECId, DECIn, DECIs, DECIsf, DECIx, DECIsx, DECIsxf,
		CHARIi, CHARId, CHARIn, CHARIs, CHARIsf, CHARIx, CHARIsx, CHARIsxf,
		STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf,
		STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf,
		STROi, STROd, STROn, STROs, STROsf, STROx, STROsx, STROsxf,
		DECOi, DECOd, DECOn, DECOs, DECOsf, DECOx, DECOsx, DECOsxf,
		CHAROi, CHAROd, CHAROn, CHAROs, CHAROsf, CHAROx, CHAROsx, CHAROsxf,
		ADDSPi, ADDSPd, ADDSPn, ADDSPs, ADDSPsf, ADDSPx, ADDSPsx, ADDSPsxf,
		SUBSPi, SUBSPd, SUBSPn, SUBSPs, SUBSPsf, SUBSPx, SUBSPsx, SUBSPsxf,
		ADDAi, ADDAd, ADDAn, ADDAs, ADDAsf, ADDAx, ADDAsx, ADDAsxf,
		ADDXi, ADDXd, ADDXn, ADDXs, ADDXsf, ADDXx, ADDXsx, ADDXsxf,
		SUBAi, SUBAd, SUBAn, SUBAs, SUBAsf, SUBAx, SUBAsx, SUBAsxf,
		SUBXi, SUBXd, SUBXn, SUBXs, SUBXsf, SUBXx, SUBXsx, SUBXsxf,
		ANDAi, ANDAd, ANDAn, ANDAs, ANDAsf, ANDAx, ANDAsx, ANDAsxf,
		ANDXi, ANDXd, ANDXn, ANDXs, ANDXsf, ANDXx, ANDXsx, ANDXsxf,
		ORAi, ORAd, ORAn, ORAs, ORAsf, ORAx, ORAsx, ORAsxf,
		ORXi, ORXd, ORXn, ORXs, ORXsf, ORXx, ORXsx, ORXsxf,
		CPAi, CPAd, CPAn, CPAs, CPAsf, CPAx, CPAsx, CPAsxf,
		CPXi, CPXd, CPXn, CPXs, CPXsf, CPXx, CPXsx, CPXsxf,
		LDAi, LDAd, LDAn, LDAs, LDAsf, LDAx, LDAsx, LDAsxf,
		LDXi, LDXd, LDXn, LDXs, LDXsf, LDXx, LDXsx, LDXsxf,
		LDBYTEAi, LDBYTEAd, LDBYTEAn, LDBYTEAs, LDBYTEAsf, LDBYTEAx, LDBYTEAsx, LDBYTEAsxf,
		LDBYTEXi, LDBYTEXd, LDBYTEXn, LDBYTEXs, LDBYTEXsf, LDBYTEXx, LDBYTEXsx, LDBYTEXsxf:
		return true
	}
	return false
}

func (oc opcode) getMode() (AddrMode, error) {
	switch oc {
	case BRi, BRx,
		BRLEi, BRLEx,
		BRLTi, BRLTx,
		BREQi, BREQx,
		BRNEi, BRNEx,
		BRGEi, BRGEx,
		BRGTi, BRGTx,
		BRVi, BRVx,
		BRCi, BRCx,
		CALLi, CALLx:
		am := oc & 0x1
		if am == 0 {
			return i, nil
		}
		return x, nil

	case NOPi, NOPd, NOPn, NOPs, NOPsf, NOPx, NOPsx, NOPsxf:
		am := oc & 0x7
		if am != 0 {
			return i, fmt.Errorf("invalid addressing mode %s for NOP", AddrMode(am))
		}
		return AddrMode(am), nil

	case DECIi, DECId, DECIn, DECIs, DECIsf, DECIx, DECIsx, DECIsxf,
		CHARIi, CHARId, CHARIn, CHARIs, CHARIsf, CHARIx, CHARIsx, CHARIsxf,
		STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf,
		STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf:
		am := oc & 0x7
		if am == 0 {
			return i, fmt.Errorf("invalid addressing mode i for %s", oc.BaseOp())
		}
		return AddrMode(am), nil

	case STROi, STROd, STROn, STROs, STROsf, STROx, STROsx, STROsxf:
		am := oc & 0x7
		switch AddrMode(am) {
		case d, n, sf:
			return AddrMode(am), nil
		}
		return i, fmt.Errorf("invalid addressing mode %s for STRO", AddrMode(am))

	case DECOi, DECOd, DECOn, DECOs, DECOsf, DECOx, DECOsx, DECOsxf,
		CHAROi, CHAROd, CHAROn, CHAROs, CHAROsf, CHAROx, CHAROsx, CHAROsxf,
		ADDSPi, ADDSPd, ADDSPn, ADDSPs, ADDSPsf, ADDSPx, ADDSPsx, ADDSPsxf,
		SUBSPi, SUBSPd, SUBSPn, SUBSPs, SUBSPsf, SUBSPx, SUBSPsx, SUBSPsxf,
		ADDAi, ADDAd, ADDAn, ADDAs, ADDAsf, ADDAx, ADDAsx, ADDAsxf,
		ADDXi, ADDXd, ADDXn, ADDXs, ADDXsf, ADDXx, ADDXsx, ADDXsxf,
		SUBAi, SUBAd, SUBAn, SUBAs, SUBAsf, SUBAx, SUBAsx, SUBAsxf,
		SUBXi, SUBXd, SUBXn, SUBXs, SUBXsf, SUBXx, SUBXsx, SUBXsxf,
		ANDAi, ANDAd, ANDAn, ANDAs, ANDAsf, ANDAx, ANDAsx, ANDAsxf,
		ANDXi, ANDXd, ANDXn, ANDXs, ANDXsf, ANDXx, ANDXsx, ANDXsxf,
		ORAi, ORAd, ORAn, ORAs, ORAsf, ORAx, ORAsx, ORAsxf,
		ORXi, ORXd, ORXn, ORXs, ORXsf, ORXx, ORXsx, ORXsxf,
		CPAi, CPAd, CPAn, CPAs, CPAsf, CPAx, CPAsx, CPAsxf,
		CPXi, CPXd, CPXn, CPXs, CPXsf, CPXx, CPXsx, CPXsxf,
		LDAi, LDAd, LDAn, LDAs, LDAsf, LDAx, LDAsx, LDAsxf,
		LDXi, LDXd, LDXn, LDXs, LDXsf, LDXx, LDXsx, LDXsxf,
		LDBYTEAi, LDBYTEAd, LDBYTEAn, LDBYTEAs, LDBYTEAsf, LDBYTEAx, LDBYTEAsx, LDBYTEAsxf,
		LDBYTEXi, LDBYTEXd, LDBYTEXn, LDBYTEXs, LDBYTEXsf, LDBYTEXx, LDBYTEXsx, LDBYTEXsxf:
		am := oc & 0x7
		return AddrMode(am), nil
	}

	return i, fmt.Errorf("getMode: no known instruction for addressing mode %s (opcode %d)", oc.BaseOp(), oc)
}

func (oc opcode) hasReg() bool {
	switch oc {
	case NOTA, NOTX,
		NEGA, NEGX,
		ASLA, ASLX,
		ASRA, ASRX,
		ROLA, ROLX,
		RORA, RORX,
		ADDAi, ADDAd, ADDAn, ADDAs, ADDAsf, ADDAx, ADDAsx, ADDAsxf,
		ADDXi, ADDXd, ADDXn, ADDXs, ADDXsf, ADDXx, ADDXsx, ADDXsxf,
		SUBAi, SUBAd, SUBAn, SUBAs, SUBAsf, SUBAx, SUBAsx, SUBAsxf,
		SUBXi, SUBXd, SUBXn, SUBXs, SUBXsf, SUBXx, SUBXsx, SUBXsxf,
		ANDAi, ANDAd, ANDAn, ANDAs, ANDAsf, ANDAx, ANDAsx, ANDAsxf,
		ANDXi, ANDXd, ANDXn, ANDXs, ANDXsf, ANDXx, ANDXsx, ANDXsxf,
		ORAi, ORAd, ORAn, ORAs, ORAsf, ORAx, ORAsx, ORAsxf,
		ORXi, ORXd, ORXn, ORXs, ORXsf, ORXx, ORXsx, ORXsxf,
		CPAi, CPAd, CPAn, CPAs, CPAsf, CPAx, CPAsx, CPAsxf,
		CPXi, CPXd, CPXn, CPXs, CPXsf, CPXx, CPXsx, CPXsxf,
		LDAi, LDAd, LDAn, LDAs, LDAsf, LDAx, LDAsx, LDAsxf,
		LDXi, LDXd, LDXn, LDXs, LDXsf, LDXx, LDXsx, LDXsxf,
		LDBYTEAi, LDBYTEAd, LDBYTEAn, LDBYTEAs, LDBYTEAsf, LDBYTEAx, LDBYTEAsx, LDBYTEAsxf,
		LDBYTEXi, LDBYTEXd, LDBYTEXn, LDBYTEXs, LDBYTEXsf, LDBYTEXx, LDBYTEXsx, LDBYTEXsxf,
		STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf,
		STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf:
		return true
	}

	return false
}

func (oc opcode) register() register {
	switch oc {
	case NOTA, NOTX,
		NEGA, NEGX,
		ASLA, ASLX,
		ASRA, ASRX,
		ROLA, ROLX,
		RORA, RORX:
		regbit := oc & 0x1
		if regbit == 0 {
			return A
		}
		return X
	case ADDAi, ADDAd, ADDAn, ADDAs, ADDAsf, ADDAx, ADDAsx, ADDAsxf,
		ADDXi, ADDXd, ADDXn, ADDXs, ADDXsf, ADDXx, ADDXsx, ADDXsxf,
		SUBAi, SUBAd, SUBAn, SUBAs, SUBAsf, SUBAx, SUBAsx, SUBAsxf,
		SUBXi, SUBXd, SUBXn, SUBXs, SUBXsf, SUBXx, SUBXsx, SUBXsxf,
		ANDAi, ANDAd, ANDAn, ANDAs, ANDAsf, ANDAx, ANDAsx, ANDAsxf,
		ANDXi, ANDXd, ANDXn, ANDXs, ANDXsf, ANDXx, ANDXsx, ANDXsxf,
		ORAi, ORAd, ORAn, ORAs, ORAsf, ORAx, ORAsx, ORAsxf,
		ORXi, ORXd, ORXn, ORXs, ORXsf, ORXx, ORXsx, ORXsxf,
		CPAi, CPAd, CPAn, CPAs, CPAsf, CPAx, CPAsx, CPAsxf,
		CPXi, CPXd, CPXn, CPXs, CPXsf, CPXx, CPXsx, CPXsxf,
		LDAi, LDAd, LDAn, LDAs, LDAsf, LDAx, LDAsx, LDAsxf,
		LDXi, LDXd, LDXn, LDXs, LDXsf, LDXx, LDXsx, LDXsxf,
		LDBYTEAi, LDBYTEAd, LDBYTEAn, LDBYTEAs, LDBYTEAsf, LDBYTEAx, LDBYTEAsx, LDBYTEAsxf,
		LDBYTEXi, LDBYTEXd, LDBYTEXn, LDBYTEXs, LDBYTEXsf, LDBYTEXx, LDBYTEXsx, LDBYTEXsxf,
		STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf,
		STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf:
		regbit := (oc & 0x8) >> 3
		if regbit == 0 {
			return A
		}
		return X
	}
	panic(fmt.Sprintf("register: failed to get register from opcode %d (%s)", oc, oc.BaseOp()))
}

type AddrMode int

const (
	i AddrMode = iota
	d
	n
	s
	sf
	x
	sx
	sxf
)

func (am AddrMode) String() string {
	switch am {
	case i:
		return "i"
	case d:
		return "d"
	case n:
		return "n"
	case s:
		return "s"
	case sf:
		return "sf"
	case x:
		return "x"
	case sx:
		return "sx"
	case sxf:
		return "sxf"
	}
	panic("unknown addressing mode")
}

type register int

const (
	A register = 0
	X          = 1
)

func (reg register) String() string {
	switch reg {
	case A:
		return "A"
	case X:
		return "X"
	}
	panic("unknown register")
}

var rdchar = make([]byte, 1)

func chari(in io.Reader) (byte, error) {
	b, err := in.Read(rdchar)
	if b == 0 || err != nil {
		return 0, fmt.Errorf("no more chars to consume")
	}
	return rdchar[0], nil
}

func deci(in io.Reader) int {
	c, err := chari(in)
	if err != nil {
		fmt.Print("Invalid DECI input\n")
		os.Exit(1)
	}

	if c != '-' && (c < '0' || c > '9') {
		fmt.Print("Invalid DECI input\n")
		os.Exit(1)
	}

	neg := c == '-'

	val := 0
	for c >= '0' && c <= '9' {
		val *= 10
		val += int(c - '0')
		c, err = chari(in)
		if err != nil {
			break
		}
	}

	if neg {
		val = -val
	}

	return val
}

func doadd(lop, rop uint16) (res uint16, n, z, v, c bool) {
	res32 := uint32(lop) + uint32(rop)
	cf := res32 & 0x10000

	if cf != 0 {
		c = true
	}
	res = uint16(res32 & 0xFFFF)

	if res >= 0x8000 {
		n = true
	}

	if res == 0 {
		z = true
	}

	v = addOverflowed(lop, rop, res)

	return res, n, z, v, c
}

func dosub(lop, rop uint16) (res uint16, n, z, v, c bool) {
	return doadd(lop, twocompl(rop))
}

func twocompl(val uint16) uint16 {
	return (val ^ 0xFFFF) + 1
}

func addOverflowed(lop, rop, res uint16) bool {
	lopSign := sign(lop)
	ropSign := sign(rop)
	resSign := sign(res)

	if lopSign == Positive && ropSign == Positive && resSign == Negative {
		return true
	}

	if lopSign == Negative && ropSign == Negative && resSign == Positive {
		return true
	}

	return false
}

func sign(op uint16) Sign {
	if op >= 0x8000 {
		return Negative
	}

	return Positive
}

// Pep8CPU is the emulated cpu for a PEP/8 machine
type Pep8CPU struct {
	// A register, accumulator -> 0 default
	A uint16
	// X register, index register -> 0 default
	X uint16
	// PC register, Program Counter -> 0 default
	PC uint16
	// IR register, Instruction Register -> 0 default
	IR uint16
	// SP register, Stack Pointer -> 0xFFFF default
	SP uint16
	// opcode register, keeps the opcode to load
	opcode opcode
	// Spec register, keeps the operand specifier for the instruction
	Spec uint16
	// Operand register, keeps the resolved operand for the instruction
	Operand uint16
	// AddrMode is the addressing mode for the operand
	AddrMode AddrMode
	// RAM is the memory allocated for a PEP/8 machine, i.e. 64kiB
	RAM []byte
	// N is the negative flag
	N bool
	// Z is the zero flag
	Z bool
	// N is the overflow flag
	V bool
	// N is the carry flag
	C bool

	// Input stream
	In io.Reader
	// Output stream
	Out io.Writer
	// NoEOFChariStop will not immediately stop when a CHARI yields no more input, and return 0s instead
	NoEOFChariStop bool
	// Trace will output the state of the CPU after each execution cycle
	Trace bool
}

func NewPep8Cpu() *Pep8CPU {
	return &Pep8CPU{
		PC:  0,
		SP:  0xFFFF,
		RAM: make([]byte, 65536),
		In:  os.Stdin,
		Out: os.Stdout,
	}
}

var regexbyte = regexp.MustCompile("[a-fA-F0-9]{2}")

func convert(input []byte) byte {
	if len(input) != 2 {
		panic(fmt.Sprintf("failed to convert byte %s: must be two chars long", input))
	}

	hi := convertByte(input[0])
	lo := convertByte(input[1])

	return (hi << 4) | lo
}

func convertByte(in byte) byte {
	if in >= '0' && in <= '9' {
		return in - '0'
	}

	if in >= 'A' && in <= 'F' {
		return 10 + in - 'A'
	}

	if in >= 'a' && in <= 'f' {
		return 10 + in - 'a'
	}

	panic(fmt.Sprintf("invalid byte: %c", in))
}

// LoadFromFile loads a pep8 program from an object code file
func (cpu *Pep8CPU) LoadFromFile(pepo string) error {
	cnts, err := os.ReadFile(pepo)
	if err != nil {
		return err
	}

	bytes := regexbyte.FindAll(cnts, -1)

	prgm := make([]byte, len(bytes))
	for i, byte := range bytes {
		val := convert(byte)
		prgm[i] = val
	}

	return cpu.Load(prgm)
}

// Load will load a program from a byte array, copy it into RAM, and init all registers to their default values
func (cpu *Pep8CPU) Load(pepo []byte) error {
	copy(cpu.RAM, pepo)
	return nil
}

func (cpu *Pep8CPU) Run() error {
	cpu.PC = 0
	cpu.SP = 0xFFFF
	for {
		cont := cpu.DoNextCycle()
		if !cont {
			break
		}
	}
	return nil
}

// DoNextCycle executes one cycle, i.e.:
//
// 1. fetch the instruction at PC
// 2. decode/validate instruction
// 3. increment PC
// 4. execute instruction
func (cpu *Pep8CPU) DoNextCycle() bool {
	cpu.opcode = opcode(cpu.RAM[cpu.PC])
	cpu.Spec = 0
	incr := 1
	if cpu.needSpec() {
		cpu.Spec = cpu.read16(cpu.PC + 1)
		cpu.getAddrMode()
		cpu.getOp()
		incr = 3
	}
	cpu.PC += uint16(incr)
	cont := cpu.Exec()
	if cpu.Trace {
		cpu.dumpState()
	}
	return cont
}

func (cpu *Pep8CPU) dumpRAM() {
	cur := 0
	for off := 0; off <= 0xFFFF; off++ {
		if cur == 0 {
			fmt.Printf("0x%04x |", off)
		}
		cur++
		fmt.Printf(" %02x", cpu.RAM[off])
		if cur >= 8 {
			fmt.Printf(" |\n")
			cur = 0
		}
	}
}

func (cpu *Pep8CPU) dumpState() {
	fmt.Printf("PC = %04x; SP = %04x; A %04x; X = %04x; Spec = %04x; N = %d, Z = %d, V = %d, C = %d; opcode = %02x; %s \n",
		cpu.PC, cpu.SP, cpu.A, cpu.X, cpu.Spec,
		booltoInt(cpu.N), booltoInt(cpu.Z), booltoInt(cpu.V), booltoInt(cpu.C),
		cpu.opcode,
		cpu.instruction())
}

func (cpu *Pep8CPU) instruction() string {
	instr := strings.Builder{}
	instr.WriteString(cpu.opcode.BaseOp())
	if cpu.opcode.hasReg() {
		instr.WriteString(cpu.opcode.register().String())
	}
	if cpu.opcode.hasAddr() {
		instr.WriteRune(',')
		addr, _ := cpu.opcode.getMode()
		instr.WriteString(addr.String())
	}

	return instr.String()
}

func booltoInt(flg bool) int {
	if flg {
		return 1
	}
	return 0
}

func (cpu *Pep8CPU) getAddrMode() {
	am, err := cpu.opcode.getMode()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	cpu.AddrMode = am
}

func (cpu *Pep8CPU) getOp() {
	switch cpu.opcode {
	case DECIi, DECId, DECIn, DECIs, DECIsf, DECIx, DECIsx, DECIsxf,
		CHARIi, CHARId, CHARIn, CHARIs, CHARIsf, CHARIx, CHARIsx, CHARIsxf,
		STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf,
		STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf,
		STROi, STROd, STROn, STROs, STROsf, STROx, STROsx, STROsxf:
		cpu.getOpAddr()
	default:
		cpu.getOpRd()
	}
}

func (cpu *Pep8CPU) getOpRd() {
	switch cpu.AddrMode {
	case i:
		cpu.Operand = cpu.Spec
	case d:
		cpu.Operand = cpu.read16(cpu.Spec)
	case x:
		cpu.Operand = cpu.read16(cpu.Spec + cpu.X)
	case n:
		cpu.Operand = cpu.read16(cpu.read16(cpu.Spec))
	case s:
		cpu.Operand = cpu.read16(cpu.Spec + cpu.SP)
	case sx:
		cpu.Operand = cpu.read16(cpu.Spec + cpu.SP + cpu.X)
	case sf:
		cpu.Operand = cpu.read16(cpu.read16(cpu.SP + cpu.Spec))
	case sxf:
		cpu.Operand = cpu.read16(cpu.read16(cpu.SP+cpu.Spec) + cpu.X)
	}
}

func (cpu *Pep8CPU) getOpAddr() {
	switch cpu.AddrMode {
	case i:
		fmt.Printf("invalid addressing mode for in-memory operation: i\n")
		os.Exit(1)
	case d:
		cpu.Operand = cpu.Spec
	case x:
		cpu.Operand = cpu.Spec + cpu.X
	case n:
		cpu.Operand = cpu.read16(cpu.Spec)
	case s:
		cpu.Operand = cpu.Spec + cpu.SP
	case sx:
		cpu.Operand = cpu.Spec + cpu.SP + cpu.X
	case sf:
		cpu.Operand = cpu.read16(cpu.SP + cpu.Spec)
	case sxf:
		cpu.Operand = cpu.read16(cpu.SP+cpu.Spec) + cpu.X
	}
}

func (cpu *Pep8CPU) read16(addr uint16) uint16 {
	b1 := uint16(cpu.RAM[addr])
	b2 := uint16(cpu.RAM[addr+1])
	return b1<<8 | b2
}

func (cpu *Pep8CPU) write16(val uint16, addr uint16) {
	cpu.RAM[addr] = uint8(val >> 8)
	cpu.RAM[addr+1] = uint8(val & 0xFF)
}

func (cpu *Pep8CPU) write8(val uint8, addr uint16) {
	cpu.RAM[addr] = val
}

func (cpu *Pep8CPU) needSpec() bool {
	oc := cpu.opcode
	switch oc {
	case STOP, RETTR, MOVSPA, MOVFLGA,
		NOTA, NOTX,
		NEGA, NEGX,
		ASLA, ASLX,
		ASRA, ASRX,
		ROLA, ROLX,
		RORA, RORX,
		NOP0, NOP1, NOP2, NOP3,
		RET0, RET1, RET2, RET3, RET4, RET5, RET6, RET7:
		return false
	}
	return true
}

// Execute the next instruction
//
// Returns whether or not to continue execution after that
func (cpu *Pep8CPU) Exec() bool {
	switch cpu.opcode {
	case STOP:
		return false
	case RETTR:
		fmt.Printf("Unsupported instruction: RETTR\n")
		return false
	case MOVSPA:
		cpu.movspa()
	case MOVFLGA:
		cpu.movflga()
	case BRi, BRx:
		cpu.br()
	case BRLEi, BRLEx:
		cpu.brle()
	case BRLTi, BRLTx:
		cpu.brlt()
	case BREQi, BREQx:
		cpu.breq()
	case BRNEi, BRNEx:
		cpu.brne()
	case BRGEi, BRGEx:
		cpu.brge()
	case BRGTi, BRGTx:
		cpu.brgt()
	case BRVi, BRVx:
		cpu.brv()
	case BRCi, BRCx:
		cpu.brc()
	case CALLi, CALLx:
		cpu.call()
	case NOTA, NOTX:
		cpu.not()
	case NEGA, NEGX:
		cpu.neg()
	case ASLA, ASLX:
		cpu.asl()
	case ASRA, ASRX:
		cpu.asr()
	case ROLA, ROLX:
		cpu.rol()
	case RORA, RORX:
		cpu.ror()
	case NOP0, NOP1, NOP2, NOP3, NOPi, NOPd, NOPn, NOPs, NOPsf, NOPx, NOPsx, NOPsxf:
		cpu.nop()
	case DECIi, DECId, DECIn, DECIs, DECIsf, DECIx, DECIsx, DECIsxf:
		cpu.deci()
	case DECOi, DECOd, DECOn, DECOs, DECOsf, DECOx, DECOsx, DECOsxf:
		cpu.deco()
	case STROi, STROd, STROn, STROs, STROsf, STROx, STROsx, STROsxf:
		cpu.stro()
	case CHARIi, CHARId, CHARIn, CHARIs, CHARIsf, CHARIx, CHARIsx, CHARIsxf:
		cpu.chari()
	case CHAROi, CHAROd, CHAROn, CHAROs, CHAROsf, CHAROx, CHAROsx, CHAROsxf:
		cpu.charo()
	case RET0, RET1, RET2, RET3, RET4, RET5, RET6, RET7:
		cpu.ret()
	case ADDSPi, ADDSPd, ADDSPn, ADDSPs, ADDSPsf, ADDSPx, ADDSPsx, ADDSPsxf:
		cpu.addsp()
	case SUBSPi, SUBSPd, SUBSPn, SUBSPs, SUBSPsf, SUBSPx, SUBSPsx, SUBSPsxf:
		cpu.subsp()
	case ADDAi, ADDAd, ADDAn, ADDAs, ADDAsf, ADDAx, ADDAsx, ADDAsxf,
		ADDXi, ADDXd, ADDXn, ADDXs, ADDXsf, ADDXx, ADDXsx, ADDXsxf:
		cpu.add()
	case SUBAi, SUBAd, SUBAn, SUBAs, SUBAsf, SUBAx, SUBAsx, SUBAsxf,
		SUBXi, SUBXd, SUBXn, SUBXs, SUBXsf, SUBXx, SUBXsx, SUBXsxf:
		cpu.sub()
	case ANDAi, ANDAd, ANDAn, ANDAs, ANDAsf, ANDAx, ANDAsx, ANDAsxf,
		ANDXi, ANDXd, ANDXn, ANDXs, ANDXsf, ANDXx, ANDXsx, ANDXsxf:
		cpu.and()
	case ORAi, ORAd, ORAn, ORAs, ORAsf, ORAx, ORAsx, ORAsxf,
		ORXi, ORXd, ORXn, ORXs, ORXsf, ORXx, ORXsx, ORXsxf:
		cpu.or()
	case CPAi, CPAd, CPAn, CPAs, CPAsf, CPAx, CPAsx, CPAsxf,
		CPXi, CPXd, CPXn, CPXs, CPXsf, CPXx, CPXsx, CPXsxf:
		cpu.cp()
	case LDAi, LDAd, LDAn, LDAs, LDAsf, LDAx, LDAsx, LDAsxf,
		LDXi, LDXd, LDXn, LDXs, LDXsf, LDXx, LDXsx, LDXsxf:
		cpu.ld()
	case LDBYTEAi, LDBYTEAd, LDBYTEAn, LDBYTEAs, LDBYTEAsf, LDBYTEAx, LDBYTEAsx, LDBYTEAsxf,
		LDBYTEXi, LDBYTEXd, LDBYTEXn, LDBYTEXs, LDBYTEXsf, LDBYTEXx, LDBYTEXsx, LDBYTEXsxf:
		cpu.ldbyte()
	case STAi, STAd, STAn, STAs, STAsf, STAx, STAsx, STAsxf,
		STXi, STXd, STXn, STXs, STXsf, STXx, STXsx, STXsxf:
		cpu.st()
	case STBYTEAi, STBYTEAd, STBYTEAn, STBYTEAs, STBYTEAsf, STBYTEAx, STBYTEAsx, STBYTEAsxf,
		STBYTEXi, STBYTEXd, STBYTEXn, STBYTEXs, STBYTEXsf, STBYTEXx, STBYTEXsx, STBYTEXsxf:
		cpu.stbyte()
	}

	return true
}

func (cpu *Pep8CPU) movspa() {
	cpu.A = cpu.SP
}

func (cpu *Pep8CPU) movflga() {
	aval := 0

	if cpu.C {
		aval |= 1
	}
	if cpu.V {
		aval |= 2
	}
	if cpu.Z {
		aval |= 4
	}
	if cpu.N {
		aval |= 8
	}

	cpu.A = uint16(aval)
}

func (cpu *Pep8CPU) br() {
	cpu.PC = cpu.Operand
}

func (cpu *Pep8CPU) brle() {
	if cpu.Z || cpu.N {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) brlt() {
	if cpu.N {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) breq() {
	if cpu.Z {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) brne() {
	if !cpu.Z {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) brge() {
	if !cpu.N {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) brgt() {
	if !cpu.N && !cpu.Z {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) brv() {
	if cpu.V {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) brc() {
	if cpu.C {
		cpu.PC = cpu.Operand
	}
}

func (cpu *Pep8CPU) call() {
	cpu.SP -= 2
	cpu.write16(cpu.PC, cpu.SP)
	cpu.PC = cpu.Operand
}

func (cpu *Pep8CPU) not() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	val = val ^ uint16(0xFFFF)

	cpu.Z = false
	cpu.N = false

	if val < 0 {
		cpu.N = true
	} else if val == 0 {
		cpu.Z = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) neg() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	val = (val ^ uint16(0xFFFF)) + 1

	cpu.Z = false
	cpu.N = false

	if val < 0 {
		cpu.N = true
	} else if val == 0 {
		cpu.Z = true
	}

	if val == 0x8000 {
		cpu.V = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) asl() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	cf := val & 0x8000
	val = val << 1

	if cf != 0 {
		cpu.C = true
	}

	if val&0x8000 != cf {
		cpu.V = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) asr() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	sf := val & 0x8000
	cf := val & 1
	val = val >> 1

	if cf != 0 {
		cpu.C = true
	}

	val = sf | val

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) rol() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	oc := uint16(0)
	if cpu.C {
		oc = 1
	}

	nc := val & 0x8000
	val = val << 1
	val |= oc

	cpu.C = false
	if nc != 0 {
		cpu.C = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) ror() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	oc := uint16(0)
	if cpu.C {
		oc = 1
	}

	nc := val & 0x1
	val = val >> 1
	val |= oc << 15

	cpu.C = false
	if nc != 0 {
		cpu.C = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) nop() {}

func (cpu *Pep8CPU) deci() {
	val := deci(cpu.In)
	if val > 32767 {
		cpu.V = true
	}
	if val < -32768 {
		cpu.V = true
	}

	if val == 0 {
		cpu.Z = true
	}

	if val < 0 {
		cpu.N = true
	}

	cpu.write16(uint16(val&0xFFFF), cpu.Operand)
}

func (cpu *Pep8CPU) deco() {
	fmt.Fprintf(cpu.Out, "%d", int16(cpu.Operand))
}

func (cpu *Pep8CPU) stro() {
	addr := cpu.Operand
	for cpu.RAM[addr] != 0 {
		fmt.Fprintf(cpu.Out, "%c", cpu.RAM[addr])
		addr++
	}
}

func (cpu *Pep8CPU) chari() {
	b, err := chari(cpu.In)
	if err != nil && !cpu.NoEOFChariStop {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	cpu.RAM[cpu.Operand] = b
}

func (cpu *Pep8CPU) charo() {
	chr := cpu.Operand >> 8
	fmt.Fprintf(cpu.Out, "%c", chr)
}

func (cpu *Pep8CPU) ret() {
	loc := cpu.opcode & 0x7
	cpu.SP -= uint16(loc)
	retaddr := cpu.read16(cpu.SP)
	cpu.SP -= 2
	cpu.PC = retaddr
}

func (cpu *Pep8CPU) addsp() {
	res, n, z, v, c := doadd(cpu.SP, cpu.Operand)
	cpu.SP = res
	cpu.N = n
	cpu.Z = z
	cpu.V = v
	cpu.C = c
}

func (cpu *Pep8CPU) subsp() {
	res, n, z, v, c := dosub(cpu.SP, cpu.Operand)
	cpu.SP = res
	cpu.N = n
	cpu.Z = z
	cpu.V = v
	cpu.C = c
}

func (cpu *Pep8CPU) add() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	res, n, z, v, c := doadd(val, cpu.Operand)
	cpu.N = n
	cpu.Z = z
	cpu.V = v
	cpu.C = c

	switch reg {
	case A:
		cpu.A = res
	case X:
		cpu.X = res
	}
}

func (cpu *Pep8CPU) sub() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	res, n, z, v, c := dosub(val, cpu.Operand)
	cpu.N = n
	cpu.Z = z
	cpu.V = v
	cpu.C = c

	switch reg {
	case A:
		cpu.A = res
	case X:
		cpu.X = res
	}
}

func (cpu *Pep8CPU) and() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	val = val & cpu.Operand

	cpu.Z = false
	cpu.N = false

	if val == 0 {
		cpu.Z = true
	}
	if val >= 0x8000 {
		cpu.N = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) or() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	val = val | cpu.Operand

	cpu.Z = false
	cpu.N = false

	if val == 0 {
		cpu.Z = true
	}
	if val >= 0x8000 {
		cpu.N = true
	}

	switch reg {
	case A:
		cpu.A = val
	case X:
		cpu.X = val
	}
}

func (cpu *Pep8CPU) cp() {
	reg := cpu.opcode.register()
	val := cpu.A
	if reg == X {
		val = cpu.X
	}

	_, n, z, v, c := dosub(val, cpu.Operand)
	cpu.N = n
	cpu.Z = z
	cpu.V = v
	cpu.C = c
}

func (cpu *Pep8CPU) ld() {
	switch cpu.opcode.register() {
	case A:
		cpu.A = cpu.Operand
	case X:
		cpu.X = cpu.Operand
	}
}

func (cpu *Pep8CPU) ldbyte() {
	switch cpu.opcode.register() {
	case A:
		cpu.A = (cpu.Operand & 0xFF00) >> 8
	case X:
		cpu.X = (cpu.Operand & 0xFF00) >> 8
	}
}

func (cpu *Pep8CPU) st() {
	val := cpu.A
	if cpu.opcode.register() == X {
		val = cpu.X
	}
	cpu.write16(val, cpu.Operand)
}

func (cpu *Pep8CPU) stbyte() {
	val := cpu.A
	if cpu.opcode.register() == X {
		val = cpu.X
	}
	cpu.write8(uint8(val>>8), cpu.Operand)
}
