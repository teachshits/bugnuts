package maps

// Expected data for testdata/sym.map

var expect5 = []SymHash{64, 2, 0, 0, 0, 4, 128, 4096, 128, 1536, 64, 128, 2, 1, 0, 0, 0, 2, 64, 128, 64, 48, 2, 4, 0, 0, 0, 0, 0, 1, 2, 4, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 0, 1, 2, 4, 2, 1, 0, 1, 2, 4, 64, 2, 0, 2, 64, 128, 64, 2, 0, 2, 64, 128, 128, 4, 0, 4, 128, 4096, 128, 4, 0, 4, 128, 4096, 64, 2, 0, 2, 64, 128, 64, 2, 0, 2, 64, 128, 2, 1, 0, 1, 2, 4, 2, 1, 0, 1, 2, 4, 2, 1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 4, 64, 2, 0, 0, 0, 1, 2, 4, 2, 48, 64, 128, 128, 4, 0, 0, 0, 2, 64, 128, 64, 1536, 128, 4096}

var expectmap5 = []*[8]SymHash{
	&[8]SymHash{64, 256, 65536, 65536, 262144, 256, 64, 262144},
	&[8]SymHash{32, 512, 32768, 2097152, 524288, 8, 2, 8388608},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16384, 1024, 16384, 4, 1024, 4194304, 4194304, 4},
	&[8]SymHash{8192, 2048, 8192, 128, 2048, 131072, 131072, 128},
	&[8]SymHash{4096, 4096, 4096, 4096, 4096, 4096, 4096, 4096},
	&[8]SymHash{2048, 8192, 2048, 131072, 8192, 128, 128, 131072},
	&[8]SymHash{1536, 16416, 525312, 4194306, 49152, 8388612, 2097156, 4194312},
	&[8]SymHash{256, 64, 262144, 64, 65536, 262144, 65536, 256},
	&[8]SymHash{128, 128, 131072, 2048, 131072, 8192, 2048, 8192},
	&[8]SymHash{2, 8, 2097152, 32768, 8388608, 512, 32, 524288},
	&[8]SymHash{1, 16, 1048576, 1048576, 16777216, 16, 1, 16777216},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{512, 32, 524288, 2, 32768, 8388608, 2097152, 8},
	&[8]SymHash{256, 64, 262144, 64, 65536, 262144, 65536, 256},
	&[8]SymHash{128, 128, 131072, 2048, 131072, 8192, 2048, 8192},
	&[8]SymHash{64, 256, 65536, 65536, 262144, 256, 64, 262144},
	&[8]SymHash{48, 513, 16809984, 2097153, 1572864, 16777224, 1048578, 8388624},
	&[8]SymHash{8, 2, 8388608, 32, 2097152, 524288, 32768, 512},
	&[8]SymHash{4, 4, 4194304, 1024, 4194304, 16384, 1024, 16384},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16, 1, 16777216, 1, 1048576, 16777216, 1048576, 16},
	&[8]SymHash{8, 2, 8388608, 32, 2097152, 524288, 32768, 512},
	&[8]SymHash{4, 4, 4194304, 1024, 4194304, 16384, 1024, 16384},
	&[8]SymHash{2, 8, 2097152, 32768, 8388608, 512, 32, 524288},
	&[8]SymHash{1, 16, 1048576, 1048576, 16777216, 16, 1, 16777216},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{2097152, 8388608, 2, 524288, 8, 32, 512, 32768},
	&[8]SymHash{1048576, 16777216, 1, 16777216, 16, 1, 16, 1048576},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16777216, 1048576, 16, 16, 1, 1048576, 16777216, 1},
	&[8]SymHash{8388608, 2097152, 8, 512, 2, 32768, 524288, 32},
	&[8]SymHash{4194304, 4194304, 4, 16384, 4, 1024, 16384, 1024},
	&[8]SymHash{2097152, 8388608, 2, 524288, 8, 32, 512, 32768},
	&[8]SymHash{1048576, 16777216, 1, 16777216, 16, 1, 16, 1048576},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16777216, 1048576, 16, 16, 1, 1048576, 16777216, 1},
	&[8]SymHash{8388608, 2097152, 8, 512, 2, 32768, 524288, 32},
	&[8]SymHash{4194304, 4194304, 4, 16384, 4, 1024, 16384, 1024},
	&[8]SymHash{65536, 262144, 64, 262144, 256, 64, 256, 65536},
	&[8]SymHash{32768, 524288, 32, 8388608, 512, 2, 8, 2097152},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{524288, 32768, 512, 8, 32, 2097152, 8388608, 2},
	&[8]SymHash{262144, 65536, 256, 256, 64, 65536, 262144, 64},
	&[8]SymHash{131072, 131072, 128, 8192, 128, 2048, 8192, 2048},
	&[8]SymHash{65536, 262144, 64, 262144, 256, 64, 256, 65536},
	&[8]SymHash{32768, 524288, 32, 8388608, 512, 2, 8, 2097152},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{524288, 32768, 512, 8, 32, 2097152, 8388608, 2},
	&[8]SymHash{262144, 65536, 256, 256, 64, 65536, 262144, 64},
	&[8]SymHash{131072, 131072, 128, 8192, 128, 2048, 8192, 2048},
	&[8]SymHash{2048, 8192, 2048, 131072, 8192, 128, 128, 131072},
	&[8]SymHash{1024, 16384, 1024, 4194304, 16384, 4, 4, 4194304},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16384, 1024, 16384, 4, 1024, 4194304, 4194304, 4},
	&[8]SymHash{8192, 2048, 8192, 128, 2048, 131072, 131072, 128},
	&[8]SymHash{4096, 4096, 4096, 4096, 4096, 4096, 4096, 4096},
	&[8]SymHash{2048, 8192, 2048, 131072, 8192, 128, 128, 131072},
	&[8]SymHash{1024, 16384, 1024, 4194304, 16384, 4, 4, 4194304},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16384, 1024, 16384, 4, 1024, 4194304, 4194304, 4},
	&[8]SymHash{8192, 2048, 8192, 128, 2048, 131072, 131072, 128},
	&[8]SymHash{4096, 4096, 4096, 4096, 4096, 4096, 4096, 4096},
	&[8]SymHash{64, 256, 65536, 65536, 262144, 256, 64, 262144},
	&[8]SymHash{32, 512, 32768, 2097152, 524288, 8, 2, 8388608},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{512, 32, 524288, 2, 32768, 8388608, 2097152, 8},
	&[8]SymHash{256, 64, 262144, 64, 65536, 262144, 65536, 256},
	&[8]SymHash{128, 128, 131072, 2048, 131072, 8192, 2048, 8192},
	&[8]SymHash{64, 256, 65536, 65536, 262144, 256, 64, 262144},
	&[8]SymHash{32, 512, 32768, 2097152, 524288, 8, 2, 8388608},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{512, 32, 524288, 2, 32768, 8388608, 2097152, 8},
	&[8]SymHash{256, 64, 262144, 64, 65536, 262144, 65536, 256},
	&[8]SymHash{128, 128, 131072, 2048, 131072, 8192, 2048, 8192},
	&[8]SymHash{2, 8, 2097152, 32768, 8388608, 512, 32, 524288},
	&[8]SymHash{1, 16, 1048576, 1048576, 16777216, 16, 1, 16777216},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16, 1, 16777216, 1, 1048576, 16777216, 1048576, 16},
	&[8]SymHash{8, 2, 8388608, 32, 2097152, 524288, 32768, 512},
	&[8]SymHash{4, 4, 4194304, 1024, 4194304, 16384, 1024, 16384},
	&[8]SymHash{2, 8, 2097152, 32768, 8388608, 512, 32, 524288},
	&[8]SymHash{1, 16, 1048576, 1048576, 16777216, 16, 1, 16777216},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16, 1, 16777216, 1, 1048576, 16777216, 1048576, 16},
	&[8]SymHash{8, 2, 8388608, 32, 2097152, 524288, 32768, 512},
	&[8]SymHash{4, 4, 4194304, 1024, 4194304, 16384, 1024, 16384},
	&[8]SymHash{2097152, 8388608, 2, 524288, 8, 32, 512, 32768},
	&[8]SymHash{1048576, 16777216, 1, 16777216, 16, 1, 16, 1048576},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16777216, 1048576, 16, 16, 1, 1048576, 16777216, 1},
	&[8]SymHash{8388608, 2097152, 8, 512, 2, 32768, 524288, 32},
	&[8]SymHash{4194304, 4194304, 4, 16384, 4, 1024, 16384, 1024},
	&[8]SymHash{65536, 262144, 64, 262144, 256, 64, 256, 65536},
	&[8]SymHash{32768, 524288, 32, 8388608, 512, 2, 8, 2097152},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{16777216, 1048576, 16, 16, 1, 1048576, 16777216, 1},
	&[8]SymHash{8388608, 2097152, 8, 512, 2, 32768, 524288, 32},
	&[8]SymHash{4194304, 4194304, 4, 16384, 4, 1024, 16384, 1024},
	&[8]SymHash{2097152, 8388608, 2, 524288, 8, 32, 512, 32768},
	&[8]SymHash{1572864, 16809984, 513, 16777224, 48, 2097153, 8388624, 1048578},
	&[8]SymHash{262144, 65536, 256, 256, 64, 65536, 262144, 64},
	&[8]SymHash{131072, 131072, 128, 8192, 128, 2048, 8192, 2048},
	&[8]SymHash{2048, 8192, 2048, 131072, 8192, 128, 128, 131072},
	&[8]SymHash{1024, 16384, 1024, 4194304, 16384, 4, 4, 4194304},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{524288, 32768, 512, 8, 32, 2097152, 8388608, 2},
	&[8]SymHash{262144, 65536, 256, 256, 64, 65536, 262144, 64},
	&[8]SymHash{131072, 131072, 128, 8192, 128, 2048, 8192, 2048},
	&[8]SymHash{65536, 262144, 64, 262144, 256, 64, 256, 65536},
	&[8]SymHash{49152, 525312, 16416, 8388612, 1536, 4194306, 4194312, 2097156},
	&[8]SymHash{8192, 2048, 8192, 128, 2048, 131072, 131072, 128},
	&[8]SymHash{4096, 4096, 4096, 4096, 4096, 4096, 4096, 4096},
}

var expect = []SymHash{65536, 512, 4, 0, 8, 1024, 131072, 16777216, 9437184, 4718592, 2359296, 131072, 512, 256, 2, 0, 4, 512, 65536, 131072, 73728, 36864, 18432, 1024, 4, 2, 1, 0, 2, 256, 512, 1024, 576, 288, 144, 8, 4, 2, 65, 2, 16448, 2105344, 269484032, 2105344, 4398047559681, 34359746560, 268435520, 8, 512, 256, 8320, 256, 512, 1024, 512, 256, 8320, 256, 512, 1024, 65536, 512, 1064960, 512, 65536, 131072, 65536, 512, 1064960, 512, 65536, 131072, 131072, 1024, 136314880, 1024, 131072, 16777216, 131072, 1024, 136314880, 1024, 131072, 16777216, 65536, 512, 1064960, 512, 65536, 131072, 65536, 512, 1064960, 512, 65536, 131072, 540672, 4224, 2233382993921, 256, 512, 1024, 512, 256, 2233382993921, 4224, 540672, 69206016, 540672, 4224, 34359738433, 2, 16448, 2105344, 269484032, 2105344, 4415226380321, 34359746816, 8858370112, 69206016, 65536, 512, 4, 0, 2, 256, 512, 1024, 18432, 36864, 73728, 131072, 131072, 1024, 8, 0, 4, 512, 65536, 131072, 2359296, 4718592, 9437184, 16777216}

var expectmap = []*[8]SymHash{
	&[8]SymHash{65536, 262144, 1073741824, 1073741824, 4294967296, 262144, 65536, 4294967296},
	&[8]SymHash{32768, 524288, 536870912, 137438953472, 8589934592, 2048, 512, 549755813888},
	&[8]SymHash{16384, 1048576, 268435456, 17592186044416, 17179869184, 16, 4, 70368744177664},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{134217728, 2097152, 134217728, 8, 2097152, 35184372088832, 35184372088832, 8},
	&[8]SymHash{67108864, 4194304, 67108864, 1024, 4194304, 274877906944, 274877906944, 1024},
	&[8]SymHash{33554432, 8388608, 33554432, 131072, 8388608, 2147483648, 2147483648, 131072},
	&[8]SymHash{16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216},
	&[8]SymHash{9437184, 33570816, 17188257792, 2147483652, 301989888, 70368744308736, 17592186175488, 2147483664},
	&[8]SymHash{4718592, 67141632, 8594128896, 274877907456, 603979776, 549755814912, 137438954496, 274877908992},
	&[8]SymHash{2359296, 134283264, 4297064448, 35184372154368, 1207959552, 4294967304, 1073741832, 35184372350976},
	&[8]SymHash{131072, 131072, 2147483648, 8388608, 2147483648, 33554432, 8388608, 33554432},
	&[8]SymHash{512, 2048, 137438953472, 536870912, 549755813888, 524288, 32768, 8589934592},
	&[8]SymHash{256, 4096, 68719476736, 68719476736, 1099511627776, 4096, 256, 1099511627776},
	&[8]SymHash{128, 8192, 34359738368, 8796093022208, 2199023255552, 32, 2, 140737488355328},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{1048576, 16384, 17179869184, 4, 268435456, 70368744177664, 17592186044416, 16},
	&[8]SymHash{524288, 32768, 8589934592, 512, 536870912, 549755813888, 137438953472, 2048},
	&[8]SymHash{262144, 65536, 4294967296, 65536, 1073741824, 4294967296, 1073741824, 262144},
	&[8]SymHash{131072, 131072, 2147483648, 8388608, 2147483648, 33554432, 8388608, 33554432},
	&[8]SymHash{73728, 262272, 2200096997376, 1073741826, 38654705664, 140737488617472, 8796093087744, 4294967328},
	&[8]SymHash{36864, 524544, 1100048498688, 137438953728, 77309411328, 1099511629824, 68719477248, 549755817984},
	&[8]SymHash{18432, 1049088, 550024249344, 17592186077184, 154618822656, 8589934608, 536870916, 70368744701952},
	&[8]SymHash{1024, 1024, 274877906944, 4194304, 274877906944, 67108864, 4194304, 67108864},
	&[8]SymHash{4, 16, 17592186044416, 268435456, 70368744177664, 1048576, 16384, 17179869184},
	&[8]SymHash{2, 32, 8796093022208, 34359738368, 140737488355328, 8192, 128, 2199023255552},
	&[8]SymHash{1, 64, 4398046511104, 4398046511104, 281474976710656, 64, 1, 281474976710656},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{8192, 128, 2199023255552, 2, 34359738368, 140737488355328, 8796093022208, 32},
	&[8]SymHash{4096, 256, 1099511627776, 256, 68719476736, 1099511627776, 68719476736, 4096},
	&[8]SymHash{2048, 512, 549755813888, 32768, 137438953472, 8589934592, 536870912, 524288},
	&[8]SymHash{1024, 1024, 274877906944, 4194304, 274877906944, 67108864, 4194304, 67108864},
	&[8]SymHash{576, 2049, 281612415664128, 536870913, 4947802324992, 281474977234944, 4398046543872, 8589934656},
	&[8]SymHash{288, 4098, 140806207832064, 68719476864, 9895604649984, 2199023259648, 34359738624, 1099511635968},
	&[8]SymHash{144, 8196, 70403103916032, 8796093038592, 19791209299968, 17179869216, 268435458, 140737489403904},
	&[8]SymHash{8, 8, 35184372088832, 2097152, 35184372088832, 134217728, 2097152, 134217728},
	&[8]SymHash{17592186044416, 70368744177664, 4, 17179869184, 16, 16384, 1048576, 268435456},
	&[8]SymHash{8796093022208, 140737488355328, 2, 2199023255552, 32, 128, 8192, 34359738368},
	&[8]SymHash{285873023221760, 285873023221760, 65, 281474976710720, 65, 4398046511105, 281474976710720, 4398046511105},
	&[8]SymHash{140737488355328, 8796093022208, 32, 8192, 2, 34359738368, 2199023255552, 128},
	&[8]SymHash{70368744177728, 17592186044417, 281474976710672, 1048577, 4398046511108, 281475245146112, 4415226380288, 16448},
	&[8]SymHash{35184372088864, 35184372088834, 140737488355336, 134217856, 8796093022216, 2199025352704, 34493956096, 2105344},
	&[8]SymHash{17592186044432, 70368744177668, 70368744177668, 17179885568, 17592186044432, 17179885568, 269484032, 269484032},
	&[8]SymHash{8796093022216, 140737488355336, 35184372088834, 2199025352704, 35184372088864, 134217856, 2105344, 34493956096},
	&[8]SymHash{285873023221764, 285873023221776, 17592186044481, 281475245146176, 70368744177729, 4398047559681, 281474976727104, 4415226380289},
	&[8]SymHash{140737488355330, 8796093022240, 8796093022240, 34359746560, 140737488355330, 34359746560, 2199023255680, 2199023255680},
	&[8]SymHash{70368744177665, 17592186044480, 4398046511120, 4398047559680, 281474976710660, 268435520, 17179869185, 281474976727040},
	&[8]SymHash{35184372088832, 35184372088832, 8, 134217728, 8, 2097152, 134217728, 2097152},
	&[8]SymHash{137438953472, 549755813888, 512, 8589934592, 2048, 32768, 524288, 536870912},
	&[8]SymHash{68719476736, 1099511627776, 256, 1099511627776, 4096, 256, 4096, 68719476736},
	&[8]SymHash{2233382993920, 2233382993920, 8320, 140737488355360, 8320, 8796093022210, 140737488355360, 8796093022210},
	&[8]SymHash{1099511627776, 68719476736, 4096, 4096, 256, 68719476736, 1099511627776, 256},
	&[8]SymHash{549755813888, 137438953472, 2048, 524288, 512, 536870912, 8589934592, 32768},
	&[8]SymHash{274877906944, 274877906944, 1024, 67108864, 1024, 4194304, 67108864, 4194304},
	&[8]SymHash{137438953472, 549755813888, 512, 8589934592, 2048, 32768, 524288, 536870912},
	&[8]SymHash{68719476736, 1099511627776, 256, 1099511627776, 4096, 256, 4096, 68719476736},
	&[8]SymHash{2233382993920, 2233382993920, 8320, 140737488355360, 8320, 8796093022210, 140737488355360, 8796093022210},
	&[8]SymHash{1099511627776, 68719476736, 4096, 4096, 256, 68719476736, 1099511627776, 256},
	&[8]SymHash{549755813888, 137438953472, 2048, 524288, 512, 536870912, 8589934592, 32768},
	&[8]SymHash{274877906944, 274877906944, 1024, 67108864, 1024, 4194304, 67108864, 4194304},
	&[8]SymHash{1073741824, 4294967296, 65536, 4294967296, 262144, 65536, 262144, 1073741824},
	&[8]SymHash{536870912, 8589934592, 32768, 549755813888, 524288, 512, 2048, 137438953472},
	&[8]SymHash{17448304640, 17448304640, 1064960, 70368744177680, 1064960, 17592186044420, 70368744177680, 17592186044420},
	&[8]SymHash{8589934592, 536870912, 524288, 2048, 32768, 137438953472, 549755813888, 512},
	&[8]SymHash{4294967296, 1073741824, 262144, 262144, 65536, 1073741824, 4294967296, 65536},
	&[8]SymHash{2147483648, 2147483648, 131072, 33554432, 131072, 8388608, 33554432, 8388608},
	&[8]SymHash{1073741824, 4294967296, 65536, 4294967296, 262144, 65536, 262144, 1073741824},
	&[8]SymHash{536870912, 8589934592, 32768, 549755813888, 524288, 512, 2048, 137438953472},
	&[8]SymHash{17448304640, 17448304640, 1064960, 70368744177680, 1064960, 17592186044420, 70368744177680, 17592186044420},
	&[8]SymHash{8589934592, 536870912, 524288, 2048, 32768, 137438953472, 549755813888, 512},
	&[8]SymHash{4294967296, 1073741824, 262144, 262144, 65536, 1073741824, 4294967296, 65536},
	&[8]SymHash{2147483648, 2147483648, 131072, 33554432, 131072, 8388608, 33554432, 8388608},
	&[8]SymHash{8388608, 33554432, 8388608, 2147483648, 33554432, 131072, 131072, 2147483648},
	&[8]SymHash{4194304, 67108864, 4194304, 274877906944, 67108864, 1024, 1024, 274877906944},
	&[8]SymHash{136314880, 136314880, 136314880, 35184372088840, 136314880, 35184372088840, 35184372088840, 35184372088840},
	&[8]SymHash{67108864, 4194304, 67108864, 1024, 4194304, 274877906944, 274877906944, 1024},
	&[8]SymHash{33554432, 8388608, 33554432, 131072, 8388608, 2147483648, 2147483648, 131072},
	&[8]SymHash{16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216},
	&[8]SymHash{8388608, 33554432, 8388608, 2147483648, 33554432, 131072, 131072, 2147483648},
	&[8]SymHash{4194304, 67108864, 4194304, 274877906944, 67108864, 1024, 1024, 274877906944},
	&[8]SymHash{136314880, 136314880, 136314880, 35184372088840, 136314880, 35184372088840, 35184372088840, 35184372088840},
	&[8]SymHash{67108864, 4194304, 67108864, 1024, 4194304, 274877906944, 274877906944, 1024},
	&[8]SymHash{33554432, 8388608, 33554432, 131072, 8388608, 2147483648, 2147483648, 131072},
	&[8]SymHash{16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216},
	&[8]SymHash{65536, 262144, 1073741824, 1073741824, 4294967296, 262144, 65536, 4294967296},
	&[8]SymHash{32768, 524288, 536870912, 137438953472, 8589934592, 2048, 512, 549755813888},
	&[8]SymHash{1064960, 1064960, 17448304640, 17592186044420, 17448304640, 70368744177680, 17592186044420, 70368744177680},
	&[8]SymHash{524288, 32768, 8589934592, 512, 536870912, 549755813888, 137438953472, 2048},
	&[8]SymHash{262144, 65536, 4294967296, 65536, 1073741824, 4294967296, 1073741824, 262144},
	&[8]SymHash{131072, 131072, 2147483648, 8388608, 2147483648, 33554432, 8388608, 33554432},
	&[8]SymHash{65536, 262144, 1073741824, 1073741824, 4294967296, 262144, 65536, 4294967296},
	&[8]SymHash{32768, 524288, 536870912, 137438953472, 8589934592, 2048, 512, 549755813888},
	&[8]SymHash{1064960, 1064960, 17448304640, 17592186044420, 17448304640, 70368744177680, 17592186044420, 70368744177680},
	&[8]SymHash{524288, 32768, 8589934592, 512, 536870912, 549755813888, 137438953472, 2048},
	&[8]SymHash{262144, 65536, 4294967296, 65536, 1073741824, 4294967296, 1073741824, 262144},
	&[8]SymHash{131072, 131072, 2147483648, 8388608, 2147483648, 33554432, 8388608, 33554432},
	&[8]SymHash{17592186044928, 70368744179712, 137438953476, 17716740096, 549755813904, 540672, 1081344, 8858370048},
	&[8]SymHash{8796093022464, 140737488359424, 68719476738, 2267742732288, 1099511627808, 4224, 8448, 1133871366144},
	&[8]SymHash{4398046519424, 281474976718976, 2233382993921, 290271069732866, 2233382993984, 140737488355361, 8796093022274, 145135534866464},
	&[8]SymHash{4096, 256, 1099511627776, 256, 68719476736, 1099511627776, 68719476736, 4096},
	&[8]SymHash{2048, 512, 549755813888, 32768, 137438953472, 8589934592, 536870912, 524288},
	&[8]SymHash{1024, 1024, 274877906944, 4194304, 274877906944, 67108864, 4194304, 67108864},
	&[8]SymHash{512, 2048, 137438953472, 536870912, 549755813888, 524288, 32768, 8589934592},
	&[8]SymHash{256, 4096, 68719476736, 68719476736, 1099511627776, 4096, 256, 1099511627776},
	&[8]SymHash{281474976718976, 4398046519424, 2233382993984, 8796093022274, 2233382993921, 145135534866464, 290271069732866, 140737488355361},
	&[8]SymHash{140737488359424, 8796093022464, 1099511627808, 8448, 68719476738, 1133871366144, 2267742732288, 4224},
	&[8]SymHash{70368744179712, 17592186044928, 549755813904, 1081344, 137438953476, 8858370048, 17716740096, 540672},
	&[8]SymHash{35184372089856, 35184372089856, 274877906952, 138412032, 274877906952, 69206016, 138412032, 69206016},
	&[8]SymHash{137438953476, 549755813904, 17592186044928, 8858370048, 70368744179712, 1081344, 540672, 17716740096},
	&[8]SymHash{68719476738, 1099511627808, 8796093022464, 1133871366144, 140737488359424, 8448, 4224, 2267742732288},
	&[8]SymHash{34359738433, 2199023255617, 285873023221888, 145135534866433, 285873023229952, 281474976710722, 4398046511137, 290271069732928},
	&[8]SymHash{32, 2, 140737488355328, 128, 8796093022208, 2199023255552, 34359738368, 8192},
	&[8]SymHash{281474976710672, 4398046511108, 70368744177728, 16448, 17592186044417, 4415226380288, 281475245146112, 1048577},
	&[8]SymHash{140737488355336, 8796093022216, 35184372088864, 2105344, 35184372088834, 34493956096, 2199025352704, 134217856},
	&[8]SymHash{70368744177668, 17592186044432, 17592186044432, 269484032, 70368744177668, 269484032, 17179885568, 17179885568},
	&[8]SymHash{35184372088834, 35184372088864, 8796093022216, 34493956096, 140737488355336, 2105344, 134217856, 2199025352704},
	&[8]SymHash{19791209300033, 70403103916097, 285873023229956, 4415226380321, 285873023221904, 290271069749312, 145135535915009, 281475245146178},
	&[8]SymHash{9895604650016, 140806207832066, 140737488359426, 2199023259776, 8796093022496, 2267742732416, 1133871374336, 34359746816},
	&[8]SymHash{4947802325008, 281612415664132, 70368744179713, 281474977251328, 17592186044992, 17716740097, 8858370112, 4398047592448},
	&[8]SymHash{274877906952, 274877906952, 35184372089856, 69206016, 35184372089856, 138412032, 69206016, 138412032},
	&[8]SymHash{1073741824, 4294967296, 65536, 4294967296, 262144, 65536, 262144, 1073741824},
	&[8]SymHash{536870912, 8589934592, 32768, 549755813888, 524288, 512, 2048, 137438953472},
	&[8]SymHash{268435456, 17179869184, 16384, 70368744177664, 1048576, 4, 16, 17592186044416},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{2199023255552, 34359738368, 8192, 32, 128, 8796093022208, 140737488355328, 2},
	&[8]SymHash{1099511627776, 68719476736, 4096, 4096, 256, 68719476736, 1099511627776, 256},
	&[8]SymHash{549755813888, 137438953472, 2048, 524288, 512, 536870912, 8589934592, 32768},
	&[8]SymHash{274877906944, 274877906944, 1024, 67108864, 1024, 4194304, 67108864, 4194304},
	&[8]SymHash{154618822656, 550024249344, 1049088, 8589934608, 18432, 17592186077184, 70368744701952, 536870916},
	&[8]SymHash{77309411328, 1100048498688, 524544, 1099511629824, 36864, 137438953728, 549755817984, 68719477248},
	&[8]SymHash{38654705664, 2200096997376, 262272, 140737488617472, 73728, 1073741826, 4294967328, 8796093087744},
	&[8]SymHash{2147483648, 2147483648, 131072, 33554432, 131072, 8388608, 33554432, 8388608},
	&[8]SymHash{8388608, 33554432, 8388608, 2147483648, 33554432, 131072, 131072, 2147483648},
	&[8]SymHash{4194304, 67108864, 4194304, 274877906944, 67108864, 1024, 1024, 274877906944},
	&[8]SymHash{2097152, 134217728, 2097152, 35184372088832, 134217728, 8, 8, 35184372088832},
	&[8]SymHash{0, 0, 0, 0, 0, 0, 0, 0},
	&[8]SymHash{17179869184, 268435456, 1048576, 16, 16384, 17592186044416, 70368744177664, 4},
	&[8]SymHash{8589934592, 536870912, 524288, 2048, 32768, 137438953472, 549755813888, 512},
	&[8]SymHash{4294967296, 1073741824, 262144, 262144, 65536, 1073741824, 4294967296, 65536},
	&[8]SymHash{2147483648, 2147483648, 131072, 33554432, 131072, 8388608, 33554432, 8388608},
	&[8]SymHash{1207959552, 4297064448, 134283264, 4294967304, 2359296, 35184372154368, 35184372350976, 1073741832},
	&[8]SymHash{603979776, 8594128896, 67141632, 549755814912, 4718592, 274877907456, 274877908992, 137438954496},
	&[8]SymHash{301989888, 17188257792, 33570816, 70368744308736, 9437184, 2147483652, 2147483664, 17592186175488},
	&[8]SymHash{16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216, 16777216},
}