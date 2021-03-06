/*
Package cp857 transforms between utf8 runes and the Code page 857 bytes.
*/
package cp857

type utf8Enc struct {
	len  uint8
	data [3]byte
}

var dec = [256]utf8Enc{
	{1, [3]byte{0x00, 0x00, 0x00}}, {1, [3]byte{0x01, 0x00, 0x00}},
	{1, [3]byte{0x02, 0x00, 0x00}}, {1, [3]byte{0x03, 0x00, 0x00}},
	{1, [3]byte{0x04, 0x00, 0x00}}, {1, [3]byte{0x05, 0x00, 0x00}},
	{1, [3]byte{0x06, 0x00, 0x00}}, {1, [3]byte{0x07, 0x00, 0x00}},
	{1, [3]byte{0x08, 0x00, 0x00}}, {1, [3]byte{0x09, 0x00, 0x00}},
	{1, [3]byte{0x0a, 0x00, 0x00}}, {1, [3]byte{0x0b, 0x00, 0x00}},
	{1, [3]byte{0x0c, 0x00, 0x00}}, {1, [3]byte{0x0d, 0x00, 0x00}},
	{1, [3]byte{0x0e, 0x00, 0x00}}, {1, [3]byte{0x0f, 0x00, 0x00}},
	{1, [3]byte{0x10, 0x00, 0x00}}, {1, [3]byte{0x11, 0x00, 0x00}},
	{1, [3]byte{0x12, 0x00, 0x00}}, {1, [3]byte{0x13, 0x00, 0x00}},
	{1, [3]byte{0x14, 0x00, 0x00}}, {1, [3]byte{0x15, 0x00, 0x00}},
	{1, [3]byte{0x16, 0x00, 0x00}}, {1, [3]byte{0x17, 0x00, 0x00}},
	{1, [3]byte{0x18, 0x00, 0x00}}, {1, [3]byte{0x19, 0x00, 0x00}},
	{1, [3]byte{0x1a, 0x00, 0x00}}, {1, [3]byte{0x1b, 0x00, 0x00}},
	{1, [3]byte{0x1c, 0x00, 0x00}}, {1, [3]byte{0x1d, 0x00, 0x00}},
	{1, [3]byte{0x1e, 0x00, 0x00}}, {1, [3]byte{0x1f, 0x00, 0x00}},
	{1, [3]byte{0x20, 0x00, 0x00}}, {1, [3]byte{0x21, 0x00, 0x00}},
	{1, [3]byte{0x22, 0x00, 0x00}}, {1, [3]byte{0x23, 0x00, 0x00}},
	{1, [3]byte{0x24, 0x00, 0x00}}, {1, [3]byte{0x25, 0x00, 0x00}},
	{1, [3]byte{0x26, 0x00, 0x00}}, {1, [3]byte{0x27, 0x00, 0x00}},
	{1, [3]byte{0x28, 0x00, 0x00}}, {1, [3]byte{0x29, 0x00, 0x00}},
	{1, [3]byte{0x2a, 0x00, 0x00}}, {1, [3]byte{0x2b, 0x00, 0x00}},
	{1, [3]byte{0x2c, 0x00, 0x00}}, {1, [3]byte{0x2d, 0x00, 0x00}},
	{1, [3]byte{0x2e, 0x00, 0x00}}, {1, [3]byte{0x2f, 0x00, 0x00}},
	{1, [3]byte{0x30, 0x00, 0x00}}, {1, [3]byte{0x31, 0x00, 0x00}},
	{1, [3]byte{0x32, 0x00, 0x00}}, {1, [3]byte{0x33, 0x00, 0x00}},
	{1, [3]byte{0x34, 0x00, 0x00}}, {1, [3]byte{0x35, 0x00, 0x00}},
	{1, [3]byte{0x36, 0x00, 0x00}}, {1, [3]byte{0x37, 0x00, 0x00}},
	{1, [3]byte{0x38, 0x00, 0x00}}, {1, [3]byte{0x39, 0x00, 0x00}},
	{1, [3]byte{0x3a, 0x00, 0x00}}, {1, [3]byte{0x3b, 0x00, 0x00}},
	{1, [3]byte{0x3c, 0x00, 0x00}}, {1, [3]byte{0x3d, 0x00, 0x00}},
	{1, [3]byte{0x3e, 0x00, 0x00}}, {1, [3]byte{0x3f, 0x00, 0x00}},
	{1, [3]byte{0x40, 0x00, 0x00}}, {1, [3]byte{0x41, 0x00, 0x00}},
	{1, [3]byte{0x42, 0x00, 0x00}}, {1, [3]byte{0x43, 0x00, 0x00}},
	{1, [3]byte{0x44, 0x00, 0x00}}, {1, [3]byte{0x45, 0x00, 0x00}},
	{1, [3]byte{0x46, 0x00, 0x00}}, {1, [3]byte{0x47, 0x00, 0x00}},
	{1, [3]byte{0x48, 0x00, 0x00}}, {1, [3]byte{0x49, 0x00, 0x00}},
	{1, [3]byte{0x4a, 0x00, 0x00}}, {1, [3]byte{0x4b, 0x00, 0x00}},
	{1, [3]byte{0x4c, 0x00, 0x00}}, {1, [3]byte{0x4d, 0x00, 0x00}},
	{1, [3]byte{0x4e, 0x00, 0x00}}, {1, [3]byte{0x4f, 0x00, 0x00}},
	{1, [3]byte{0x50, 0x00, 0x00}}, {1, [3]byte{0x51, 0x00, 0x00}},
	{1, [3]byte{0x52, 0x00, 0x00}}, {1, [3]byte{0x53, 0x00, 0x00}},
	{1, [3]byte{0x54, 0x00, 0x00}}, {1, [3]byte{0x55, 0x00, 0x00}},
	{1, [3]byte{0x56, 0x00, 0x00}}, {1, [3]byte{0x57, 0x00, 0x00}},
	{1, [3]byte{0x58, 0x00, 0x00}}, {1, [3]byte{0x59, 0x00, 0x00}},
	{1, [3]byte{0x5a, 0x00, 0x00}}, {1, [3]byte{0x5b, 0x00, 0x00}},
	{1, [3]byte{0x5c, 0x00, 0x00}}, {1, [3]byte{0x5d, 0x00, 0x00}},
	{1, [3]byte{0x5e, 0x00, 0x00}}, {1, [3]byte{0x5f, 0x00, 0x00}},
	{1, [3]byte{0x60, 0x00, 0x00}}, {1, [3]byte{0x61, 0x00, 0x00}},
	{1, [3]byte{0x62, 0x00, 0x00}}, {1, [3]byte{0x63, 0x00, 0x00}},
	{1, [3]byte{0x64, 0x00, 0x00}}, {1, [3]byte{0x65, 0x00, 0x00}},
	{1, [3]byte{0x66, 0x00, 0x00}}, {1, [3]byte{0x67, 0x00, 0x00}},
	{1, [3]byte{0x68, 0x00, 0x00}}, {1, [3]byte{0x69, 0x00, 0x00}},
	{1, [3]byte{0x6a, 0x00, 0x00}}, {1, [3]byte{0x6b, 0x00, 0x00}},
	{1, [3]byte{0x6c, 0x00, 0x00}}, {1, [3]byte{0x6d, 0x00, 0x00}},
	{1, [3]byte{0x6e, 0x00, 0x00}}, {1, [3]byte{0x6f, 0x00, 0x00}},
	{1, [3]byte{0x70, 0x00, 0x00}}, {1, [3]byte{0x71, 0x00, 0x00}},
	{1, [3]byte{0x72, 0x00, 0x00}}, {1, [3]byte{0x73, 0x00, 0x00}},
	{1, [3]byte{0x74, 0x00, 0x00}}, {1, [3]byte{0x75, 0x00, 0x00}},
	{1, [3]byte{0x76, 0x00, 0x00}}, {1, [3]byte{0x77, 0x00, 0x00}},
	{1, [3]byte{0x78, 0x00, 0x00}}, {1, [3]byte{0x79, 0x00, 0x00}},
	{1, [3]byte{0x7a, 0x00, 0x00}}, {1, [3]byte{0x7b, 0x00, 0x00}},
	{1, [3]byte{0x7c, 0x00, 0x00}}, {1, [3]byte{0x7d, 0x00, 0x00}},
	{1, [3]byte{0x7e, 0x00, 0x00}}, {1, [3]byte{0x7f, 0x00, 0x00}},
	{2, [3]byte{0xc3, 0x87, 0x00}}, {2, [3]byte{0xc3, 0xbc, 0x00}},
	{2, [3]byte{0xc3, 0xa9, 0x00}}, {2, [3]byte{0xc3, 0xa2, 0x00}},
	{2, [3]byte{0xc3, 0xa4, 0x00}}, {2, [3]byte{0xc3, 0xa0, 0x00}},
	{2, [3]byte{0xc3, 0xa5, 0x00}}, {2, [3]byte{0xc3, 0xa7, 0x00}},
	{2, [3]byte{0xc3, 0xaa, 0x00}}, {2, [3]byte{0xc3, 0xab, 0x00}},
	{2, [3]byte{0xc3, 0xa8, 0x00}}, {2, [3]byte{0xc3, 0xaf, 0x00}},
	{2, [3]byte{0xc3, 0xae, 0x00}}, {2, [3]byte{0xc4, 0xb1, 0x00}},
	{2, [3]byte{0xc3, 0x84, 0x00}}, {2, [3]byte{0xc3, 0x85, 0x00}},
	{2, [3]byte{0xc3, 0x89, 0x00}}, {2, [3]byte{0xc3, 0xa6, 0x00}},
	{2, [3]byte{0xc3, 0x86, 0x00}}, {2, [3]byte{0xc3, 0xb4, 0x00}},
	{2, [3]byte{0xc3, 0xb6, 0x00}}, {2, [3]byte{0xc3, 0xb2, 0x00}},
	{2, [3]byte{0xc3, 0xbb, 0x00}}, {2, [3]byte{0xc3, 0xb9, 0x00}},
	{2, [3]byte{0xc4, 0xb0, 0x00}}, {2, [3]byte{0xc3, 0x96, 0x00}},
	{2, [3]byte{0xc3, 0x9c, 0x00}}, {2, [3]byte{0xc3, 0xb8, 0x00}},
	{2, [3]byte{0xc2, 0xa3, 0x00}}, {2, [3]byte{0xc3, 0x98, 0x00}},
	{2, [3]byte{0xc5, 0x9e, 0x00}}, {2, [3]byte{0xc5, 0x9f, 0x00}},
	{2, [3]byte{0xc3, 0xa1, 0x00}}, {2, [3]byte{0xc3, 0xad, 0x00}},
	{2, [3]byte{0xc3, 0xb3, 0x00}}, {2, [3]byte{0xc3, 0xba, 0x00}},
	{2, [3]byte{0xc3, 0xb1, 0x00}}, {2, [3]byte{0xc3, 0x91, 0x00}},
	{2, [3]byte{0xc4, 0x9e, 0x00}}, {2, [3]byte{0xc4, 0x9f, 0x00}},
	{2, [3]byte{0xc2, 0xbf, 0x00}}, {2, [3]byte{0xc2, 0xae, 0x00}},
	{2, [3]byte{0xc2, 0xac, 0x00}}, {2, [3]byte{0xc2, 0xbd, 0x00}},
	{2, [3]byte{0xc2, 0xbc, 0x00}}, {2, [3]byte{0xc2, 0xa1, 0x00}},
	{2, [3]byte{0xc2, 0xab, 0x00}}, {2, [3]byte{0xc2, 0xbb, 0x00}},
	{3, [3]byte{0xe2, 0x96, 0x91}}, {3, [3]byte{0xe2, 0x96, 0x92}},
	{3, [3]byte{0xe2, 0x96, 0x93}}, {3, [3]byte{0xe2, 0x94, 0x82}},
	{3, [3]byte{0xe2, 0x94, 0xa4}}, {2, [3]byte{0xc3, 0x81, 0x00}},
	{2, [3]byte{0xc3, 0x82, 0x00}}, {2, [3]byte{0xc3, 0x80, 0x00}},
	{2, [3]byte{0xc2, 0xa9, 0x00}}, {3, [3]byte{0xe2, 0x95, 0xa3}},
	{3, [3]byte{0xe2, 0x95, 0x91}}, {3, [3]byte{0xe2, 0x95, 0x97}},
	{3, [3]byte{0xe2, 0x95, 0x9d}}, {2, [3]byte{0xc2, 0xa2, 0x00}},
	{2, [3]byte{0xc2, 0xa5, 0x00}}, {3, [3]byte{0xe2, 0x94, 0x90}},
	{3, [3]byte{0xe2, 0x94, 0x94}}, {3, [3]byte{0xe2, 0x94, 0xb4}},
	{3, [3]byte{0xe2, 0x94, 0xac}}, {3, [3]byte{0xe2, 0x94, 0x9c}},
	{3, [3]byte{0xe2, 0x94, 0x80}}, {3, [3]byte{0xe2, 0x94, 0xbc}},
	{2, [3]byte{0xc3, 0xa3, 0x00}}, {2, [3]byte{0xc3, 0x83, 0x00}},
	{3, [3]byte{0xe2, 0x95, 0x9a}}, {3, [3]byte{0xe2, 0x95, 0x94}},
	{3, [3]byte{0xe2, 0x95, 0xa9}}, {3, [3]byte{0xe2, 0x95, 0xa6}},
	{3, [3]byte{0xe2, 0x95, 0xa0}}, {3, [3]byte{0xe2, 0x95, 0x90}},
	{3, [3]byte{0xe2, 0x95, 0xac}}, {2, [3]byte{0xc2, 0xa4, 0x00}},
	{2, [3]byte{0xc2, 0xba, 0x00}}, {2, [3]byte{0xc2, 0xaa, 0x00}},
	{2, [3]byte{0xc3, 0x8a, 0x00}}, {2, [3]byte{0xc3, 0x8b, 0x00}},
	{2, [3]byte{0xc3, 0x88, 0x00}}, {3, [3]byte{0xef, 0xbf, 0xbd}},
	{2, [3]byte{0xc3, 0x8d, 0x00}}, {2, [3]byte{0xc3, 0x8e, 0x00}},
	{2, [3]byte{0xc3, 0x8f, 0x00}}, {3, [3]byte{0xe2, 0x94, 0x98}},
	{3, [3]byte{0xe2, 0x94, 0x8c}}, {3, [3]byte{0xe2, 0x96, 0x88}},
	{3, [3]byte{0xe2, 0x96, 0x84}}, {2, [3]byte{0xc2, 0xa6, 0x00}},
	{2, [3]byte{0xc3, 0x8c, 0x00}}, {3, [3]byte{0xe2, 0x96, 0x80}},
	{2, [3]byte{0xc3, 0x93, 0x00}}, {2, [3]byte{0xc3, 0x9f, 0x00}},
	{2, [3]byte{0xc3, 0x94, 0x00}}, {2, [3]byte{0xc3, 0x92, 0x00}},
	{2, [3]byte{0xc3, 0xb5, 0x00}}, {2, [3]byte{0xc3, 0x95, 0x00}},
	{2, [3]byte{0xc2, 0xb5, 0x00}}, {3, [3]byte{0xef, 0xbf, 0xbd}},
	{2, [3]byte{0xc3, 0x97, 0x00}}, {2, [3]byte{0xc3, 0x9a, 0x00}},
	{2, [3]byte{0xc3, 0x9b, 0x00}}, {2, [3]byte{0xc3, 0x99, 0x00}},
	{2, [3]byte{0xc3, 0xac, 0x00}}, {2, [3]byte{0xc3, 0xbf, 0x00}},
	{2, [3]byte{0xc2, 0xaf, 0x00}}, {2, [3]byte{0xc2, 0xb4, 0x00}},
	{2, [3]byte{0xc2, 0xad, 0x00}}, {2, [3]byte{0xc2, 0xb1, 0x00}},
	{3, [3]byte{0xef, 0xbf, 0xbd}}, {2, [3]byte{0xc2, 0xbe, 0x00}},
	{2, [3]byte{0xc2, 0xb6, 0x00}}, {2, [3]byte{0xc2, 0xa7, 0x00}},
	{2, [3]byte{0xc3, 0xb7, 0x00}}, {2, [3]byte{0xc2, 0xb8, 0x00}},
	{2, [3]byte{0xc2, 0xb0, 0x00}}, {2, [3]byte{0xc2, 0xa8, 0x00}},
	{2, [3]byte{0xc2, 0xb7, 0x00}}, {2, [3]byte{0xc2, 0xb9, 0x00}},
	{2, [3]byte{0xc2, 0xb3, 0x00}}, {2, [3]byte{0xc2, 0xb2, 0x00}},
	{3, [3]byte{0xe2, 0x96, 0xa0}}, {2, [3]byte{0xc2, 0xa0, 0x00}},
}
var enc = [256]uint32{
	0x00000000, 0x01000001, 0x02000002, 0x03000003, 0x04000004, 0x05000005, 0x06000006, 0x07000007,
	0x08000008, 0x09000009, 0x0a00000a, 0x0b00000b, 0x0c00000c, 0x0d00000d, 0x0e00000e, 0x0f00000f,
	0x10000010, 0x11000011, 0x12000012, 0x13000013, 0x14000014, 0x15000015, 0x16000016, 0x17000017,
	0x18000018, 0x19000019, 0x1a00001a, 0x1b00001b, 0x1c00001c, 0x1d00001d, 0x1e00001e, 0x1f00001f,
	0x20000020, 0x21000021, 0x22000022, 0x23000023, 0x24000024, 0x25000025, 0x26000026, 0x27000027,
	0x28000028, 0x29000029, 0x2a00002a, 0x2b00002b, 0x2c00002c, 0x2d00002d, 0x2e00002e, 0x2f00002f,
	0x30000030, 0x31000031, 0x32000032, 0x33000033, 0x34000034, 0x35000035, 0x36000036, 0x37000037,
	0x38000038, 0x39000039, 0x3a00003a, 0x3b00003b, 0x3c00003c, 0x3d00003d, 0x3e00003e, 0x3f00003f,
	0x40000040, 0x41000041, 0x42000042, 0x43000043, 0x44000044, 0x45000045, 0x46000046, 0x47000047,
	0x48000048, 0x49000049, 0x4a00004a, 0x4b00004b, 0x4c00004c, 0x4d00004d, 0x4e00004e, 0x4f00004f,
	0x50000050, 0x51000051, 0x52000052, 0x53000053, 0x54000054, 0x55000055, 0x56000056, 0x57000057,
	0x58000058, 0x59000059, 0x5a00005a, 0x5b00005b, 0x5c00005c, 0x5d00005d, 0x5e00005e, 0x5f00005f,
	0x60000060, 0x61000061, 0x62000062, 0x63000063, 0x64000064, 0x65000065, 0x66000066, 0x67000067,
	0x68000068, 0x69000069, 0x6a00006a, 0x6b00006b, 0x6c00006c, 0x6d00006d, 0x6e00006e, 0x6f00006f,
	0x70000070, 0x71000071, 0x72000072, 0x73000073, 0x74000074, 0x75000075, 0x76000076, 0x77000077,
	0x78000078, 0x79000079, 0x7a00007a, 0x7b00007b, 0x7c00007c, 0x7d00007d, 0x7e00007e, 0x7f00007f,
	0xff0000a0, 0xad0000a1, 0xbd0000a2, 0x9c0000a3, 0xcf0000a4, 0xbe0000a5, 0xdd0000a6, 0xf50000a7,
	0xf90000a8, 0xb80000a9, 0xd10000aa, 0xae0000ab, 0xaa0000ac, 0xf00000ad, 0xa90000ae, 0xee0000af,
	0xf80000b0, 0xf10000b1, 0xfd0000b2, 0xfc0000b3, 0xef0000b4, 0xe60000b5, 0xf40000b6, 0xfa0000b7,
	0xf70000b8, 0xfb0000b9, 0xd00000ba, 0xaf0000bb, 0xac0000bc, 0xab0000bd, 0xf30000be, 0xa80000bf,
	0xb70000c0, 0xb50000c1, 0xb60000c2, 0xc70000c3, 0x8e0000c4, 0x8f0000c5, 0x920000c6, 0x800000c7,
	0xd40000c8, 0x900000c9, 0xd20000ca, 0xd30000cb, 0xde0000cc, 0xd60000cd, 0xd70000ce, 0xd80000cf,
	0xa50000d1, 0xe30000d2, 0xe00000d3, 0xe20000d4, 0xe50000d5, 0x990000d6, 0xe80000d7, 0x9d0000d8,
	0xeb0000d9, 0xe90000da, 0xea0000db, 0x9a0000dc, 0xe10000df, 0x850000e0, 0xa00000e1, 0x830000e2,
	0xc60000e3, 0x840000e4, 0x860000e5, 0x910000e6, 0x870000e7, 0x8a0000e8, 0x820000e9, 0x880000ea,
	0x890000eb, 0xec0000ec, 0xa10000ed, 0x8c0000ee, 0x8b0000ef, 0xa40000f1, 0x950000f2, 0xa20000f3,
	0x930000f4, 0xe40000f5, 0x940000f6, 0xf60000f7, 0x9b0000f8, 0x970000f9, 0xa30000fa, 0x960000fb,
	0x810000fc, 0xed0000ff, 0xa600011e, 0xa700011f, 0x98000130, 0x8d000131, 0x9e00015e, 0x9f00015f,
	0xc4002500, 0xb3002502, 0xda00250c, 0xbf002510, 0xc0002514, 0xd9002518, 0xc300251c, 0xb4002524,
	0xc200252c, 0xc1002534, 0xc500253c, 0xcd002550, 0xba002551, 0xc9002554, 0xbb002557, 0xc800255a,
	0xbc00255d, 0xcc002560, 0xb9002563, 0xcb002566, 0xca002569, 0xce00256c, 0xdf002580, 0xdc002584,
	0xdb002588, 0xb0002591, 0xb1002592, 0xb2002593, 0xfe0025a0, 0xfe0025a0, 0xfe0025a0, 0xfe0025a0,
}

// DecodeByte returns the Code page 857's rune decoding of the byte b.
func DecodeByte(b byte) rune {
	switch y := dec[b]; y.len {
	case 1:
		return rune(y.data[0])
	case 2:
		return rune(y.data[0]&0x1f)<<6 | rune(y.data[1]&0x3f)
	default:
		return rune(y.data[0]&0x0f)<<12 | rune(y.data[1]&0x3f)<<6 |
			rune(y.data[2]&0x3f)
	}
}

// EncodeRune returns the Code page 857's byte encoding of the rune r. ok is
// whether r is in the repertoire. If not, b is set to the  the ASCII
// substitute character '\x1a'.
func EncodeRune(r rune) (b byte, ok bool) {
	if r < 0x80 {
		return byte(r), true
	}
	for lo, hi := 0x80, 0x100; ; {
		if lo >= hi {
			return 0x1a, false
		}
		mi := (lo + hi) / 2
		r0 := rune(enc[mi] & (1<<24 - 1))
		if r0 < r {
			lo = mi + 1
		} else if r0 > r {
			hi = mi
		} else {
			return byte(enc[mi] >> 24), true
		}
	}
}
