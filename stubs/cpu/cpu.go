package cpu

type x86features struct {
	HasAES, HasPCLMULQDQ, HasSSE41, HasSSSE3 bool
	HasADX, HasAVX, HasAVX2, HasBMI2, HasSHA bool
}
type arm64features struct {
	HasAES, HasPMULL, HasSHA2, HasSHA512, HasSHA3 bool
}
type s390xfeatures struct {
	HasAES, HasAESCBC, HasAESCTR, HasAESGCM bool
	HasGHASH, HasSHA256, HasSHA512, HasSHA3 bool
	HasECDSA                                bool
}
type loong64features struct {
	HasLSX, HasLASX bool
}

var X86 = x86features{
	HasAES: true, HasPCLMULQDQ: true, HasSSE41: true, HasSSSE3: true,
	HasADX: true, HasAVX: true, HasAVX2: true, HasBMI2: true, HasSHA: true,
}
var ARM64 arm64features
var S390X s390xfeatures
var Loong64 loong64features
