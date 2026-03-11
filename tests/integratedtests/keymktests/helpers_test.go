package keymktests

// mockByteEnumNamer implements enuminf.ByteEnumNamer for testing.
type mockByteEnumNamer struct {
	name string
}

func (m mockByteEnumNamer) Name() string {
	return m.name
}

func (m mockByteEnumNamer) ValueByte() byte {
	return 0
}

func (m mockByteEnumNamer) String() string {
	return m.name
}
