//go:build bench

package benchmark

import (
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

// Benchmarks are opt-in; run with `make test-bench`.

const (
	maxUint16PlusOne = 65536
	outOfRangeBase   = 70000
	outOfRangeModulo = 10
	integerSample    = 12345
	ciStringSample   = "RFID-ABC123"
	sampleTimestamp  = "2025-01-02T15:00:00Z"
)

func BenchmarkNewCiString20Type(b *testing.B) {
	b.ReportAllocs()

	const sample = "RFID-ABC1234567890"

	for i := 0; i < b.N; i++ {
		if _, err := st.NewCiString20Type(sample); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString20Type_InvalidNonASCII(b *testing.B) {
	b.ReportAllocs()

	const sample = "bad\x01value"

	for i := 0; i < b.N; i++ {
		_, _ = st.NewCiString20Type(sample)
	}
}

func BenchmarkNewCiString25Type(b *testing.B) {
	b.ReportAllocs()

	const sample = "RFID-ABC1234567890123"

	for i := 0; i < b.N; i++ {
		if _, err := st.NewCiString25Type(sample); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString50Type(b *testing.B) {
	b.ReportAllocs()

	const sample = "RFID-ABC12345678901234567890123456789012"

	for i := 0; i < b.N; i++ {
		if _, err := st.NewCiString50Type(sample); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString255Type(b *testing.B) {
	b.ReportAllocs()

	long := make([]byte, 255)
	for i := range long {
		long[i] = 'A'
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := st.NewCiString255Type(string(long)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString500Type(b *testing.B) {
	b.ReportAllocs()

	long := make([]byte, 500)
	for i := range long {
		long[i] = 'B'
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := st.NewCiString500Type(string(long)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCiString20TypeString(b *testing.B) {
	b.ReportAllocs()

	cs, _ := st.NewCiString20Type(ciStringSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cs.String()
	}
}

func BenchmarkCiString20TypeValue(b *testing.B) {
	b.ReportAllocs()

	cs, _ := st.NewCiString20Type(ciStringSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cs.Value()
	}
}

func BenchmarkNewInteger(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := st.NewInteger(i % maxUint16PlusOne); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewInteger_OutOfRange(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = st.NewInteger(outOfRangeBase + (i % outOfRangeModulo))
	}
}

func BenchmarkIntegerString(b *testing.B) {
	b.ReportAllocs()

	val, _ := st.NewInteger(integerSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = val.String()
	}
}

func BenchmarkIntegerValue(b *testing.B) {
	b.ReportAllocs()

	val, _ := st.NewInteger(integerSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = val.Value()
	}
}

func BenchmarkNewInteger_ParseStrings(b *testing.B) {
	b.ReportAllocs()

	inputs := []string{"0", "123", "65535", "99999", "-1"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := inputs[i%len(inputs)]
		v, _ := strconv.Atoi(s)
		_, _ = st.NewInteger(v)
	}
}

func BenchmarkNewDateTime(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := st.NewDateTime(sampleTimestamp); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDateTimeString(b *testing.B) {
	b.ReportAllocs()

	dt, _ := st.NewDateTime(sampleTimestamp)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dt.String()
	}
}

func BenchmarkNewIDToken(b *testing.B) {
	b.ReportAllocs()

	ci, _ := st.NewCiString20Type(ciStringSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = st.NewIDToken(ci)
	}
}

func BenchmarkNewIDTagInfo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := st.NewIDTagInfo(
			st.AuthorizationStatusAccepted,
		); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIDTagInfoBuilderChain(b *testing.B) {
	b.ReportAllocs()

	expiry, _ := st.NewDateTime(sampleTimestamp)
	ci, _ := st.NewCiString20Type("PARENT-1")
	parent := st.NewIDToken(ci)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		info, err := st.NewIDTagInfo(
			st.AuthorizationStatusAccepted,
		)
		if err != nil {
			b.Fatal(err)
		}

		info = info.
			WithExpiryDate(expiry).
			WithParentIDTag(parent)
		_ = info
	}
}
