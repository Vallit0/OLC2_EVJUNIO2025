package arm

// General purpose registers
const (
	X0  = "x0"
	X1  = "x1"
	X2  = "x2"
	X3  = "x3"
	X4  = "x4"
	X5  = "x5"
	X6  = "x6"
	X7  = "x7"
	X8  = "x8"
	X9  = "x9"
	X10 = "x10"
	X11 = "x11"
	X12 = "x12"
	X13 = "x13"
	X14 = "x14"
	X15 = "x15"
	X16 = "x16"
	X17 = "x17"
	X18 = "x18"
	X19 = "x19"
	X20 = "x20"
	X21 = "x21"
	X22 = "x22"
	X23 = "x23"
	X24 = "x24"
	X25 = "x25"
	X26 = "x26"
	X27 = "x27"
	X28 = "x28"
	X29 = "x29"
	X30 = "x30"

	// Special registers
	SP  = "sp"  // Stack pointer
	PC  = "pc"  // Program Counter
	XZR = "xzr" // Zero register (always zero)

	// Aliases
	FP = X29 // Frame pointer
	LR = X30 // Link register

	// SIMD and floating-point registers
	// SIMD -> Single Instruction, Multiple Data
	V0  = "v0"
	V1  = "v1"
	V2  = "v2"
	V3  = "v3"
	V4  = "v4"
	V5  = "v5"
	V6  = "v6"
	V7  = "v7"
	V8  = "v8"
	V9  = "v9"
	V10 = "v10"
	V11 = "v11"
	V12 = "v12"
	V13 = "v13"
	V14 = "v14"
	V15 = "v15"
	V16 = "v16"
	V17 = "v17"
	V18 = "v18"
	V19 = "v19"
	V20 = "v20"
	V21 = "v21"
	V22 = "v22"
	V23 = "v23"
	V24 = "v24"
	V25 = "v25"
	V26 = "v26"
	V27 = "v27"
	V28 = "v28"
	V29 = "v29"
	V30 = "v30"
	V31 = "v31"
)
