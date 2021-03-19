package module

import (
	"github.com/liangyaopei/checker"
	"testing"
)

func BenchmarkIp(b *testing.B) {
	type Test struct {
		IP string
	}
	test := Test{IP: "127.0.0.1"}

	ipChecker := NewChecker()
	ipChecker.Add(main.Ip("IP"), "wrong ip")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = ipChecker.Check(test)
	}
}

func BenchmarkNot(b *testing.B) {
	type Test struct {
		NotIP string
	}
	test := Test{NotIP: "127.0.0.1.1"}

	notIPChecker := NewChecker()
	notIPChecker.Add(Not(main.Ip("IP")), "wrong ip")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = notIPChecker.Check(test)
	}
}

func BenchmarkMap(b *testing.B) {

	kvMap := make(map[main.keyStruct]main.valueStruct)
	keys := []main.keyStruct{{1}, {2}, {3}}
	for _, key := range keys {
		kvMap[key] = main.valueStruct{Value: 9}
	}
	m := main.mapStruct{
		kvMap,
	}

	mapChecker := NewChecker()
	mapRule := main.Map("Map",
		RangeInt("Key", 1, 10),
		InInt("Value", 8, 9, 10))
	mapChecker.Add(mapRule, "invalid map")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = mapChecker.Check(m)
	}
}
