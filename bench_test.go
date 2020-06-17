package bench

import (
	_ "net/http/pprof"
	"testing"
	"time"

	// cevaris "github.com/cevaris/ordered_map"
	elliot "github.com/elliotchance/orderedmap"
	w8 "github.com/wk8/go-ordered-map"
)

const ITERS = 70_000_000

func BenchmarkMapInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := map[Outpoint]*UTXOEntry{}

		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			entry.scriptPubKey = make([]byte, 37)
			m[outpoint] = &entry
		}

		time.Sleep(1 * time.Hour)
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkMapInsert**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)

		// os.Stdout.Sync()

	}
}

func BenchmarkElliotMapInsert(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := elliot.NewOrderedMap()

		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m.Set(outpoint, &entry)
		}
		//var r runtime.MemStats
		//runtime.ReadMemStats(&r)
		//os.Stdout.Sync()
		//fmt.Printf("\n**BenchmarkElliotMapInsert**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		//os.Stdout.Sync()

	}

}

func BenchmarkW8MapInsert(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := w8.New()

		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m.Set(outpoint, &entry)
		}
		//var r runtime.MemStats
		//runtime.ReadMemStats(&r)
		//os.Stdout.Sync()
		//fmt.Printf("\n**BenchmarkW8MapInsert**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		//
		//os.Stdout.Sync()
	}
}

func BenchmarkElichaiMapInsert(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := NewOrderedUtxoCollection()

		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m.Add(outpoint, &entry)
		}
		//var r runtime.MemStats
		//runtime.ReadMemStats(&r)
		//os.Stdout.Sync()
		//fmt.Printf("\n**BenchmarkElichaiMapInsert**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		//os.Stdout.Sync()

	}
}

// func BenchmarkCevarisMapInsert(b *testing.B) {
//
// 	for n := 0; n < b.N; n++ {
// 		m := cevaris.NewOrderedMap()

// 		for i := 0; i < ITERS; i++ {
// 			outpoint := Outpoint{}
// 			outpoint.Index = uint32(i)
// 			entry := UTXOEntry{}
// 			entry.amount = uint64(i)
// 			entry.blockBlueScore = uint64(i) * 2
// 			m.Set(outpoint, &entry)
// 		}
// 	}
// }

func BenchmarkMapRemove(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := utxoCollection{}

		b.StopTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m[outpoint] = &entry
		}
		b.StartTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			delete(m, outpoint)
		}
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkMapRemove**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		// os.Stdout.Sync()

	}
}

func BenchmarkElliotMapRemove(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := elliot.NewOrderedMap()

		b.StopTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m.Set(outpoint, &entry)
		}
		b.StartTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			m.Delete(outpoint)
		}
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkElliotMapRemove**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		// os.Stdout.Sync()

	}
}

func BenchmarkW8MapRemove(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := w8.New()

		b.StopTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m.Set(outpoint, &entry)
		}
		b.StartTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			m.Delete(outpoint)
		}
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkW8MapRemove**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)

		// os.Stdout.Sync()
	}
}

func BenchmarkElichaiMapRemove(b *testing.B) {

	for n := 0; n < b.N; n++ {
		m := NewOrderedUtxoCollection()

		b.StopTimer()
		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			entry := UTXOEntry{}
			entry.amount = uint64(i)
			entry.blockBlueScore = uint64(i) * 2
			shit := [37]byte{}
			entry.scriptPubKey = shit[:]
			m.Add(outpoint, &entry)
		}
		b.StartTimer()

		for i := 0; i < ITERS; i++ {
			outpoint := Outpoint{}
			outpoint.Index = uint32(i)
			m.Remove(outpoint)
		}
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkElichaiMapRemove**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		// os.Stdout.Sync()

	}
}

// func BenchmarkCevarisMapRemove(b *testing.B) {
//
// 	for n := 0; n < b.N; n++ {
// 		m := cevaris.NewOrderedMap()

// 		b.StopTimer()
// 		for i := 0; i < ITERS; i++ {
// 			outpoint := Outpoint{}
// 			outpoint.Index = uint32(i)
// 			entry := UTXOEntry{}
// 			entry.amount = uint64(i)
// 			entry.blockBlueScore = uint64(i) * 2
// 			m.Set(outpoint, &entry)
// 		}
// 		b.StartTimer()
// 		for i := 0; i < ITERS; i++ {
// 			outpoint := Outpoint{}
// 			outpoint.Index = uint32(i)
// 			m.Delete(outpoint)
// 		}
// 	}
// }

//func BenchmarkMapGet(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := utxoCollection{}
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m[outpoint] = &entry
//		}
//		b.StartTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_ = m[outpoint]
//		}
//	}
//}
//
//func BenchmarkElliotGet(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := elliot.NewOrderedMap()
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m.Set(outpoint, &entry)
//		}
//		b.StartTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_, _ = m.Get(outpoint)
//		}
//	}
//}
//
//func BenchmarkW8MapGet(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := w8.New()
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m.Set(outpoint, &entry)
//		}
//		b.StartTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_, _ = m.Get(outpoint)
//		}
//	}
//}
//
//func BenchmarkElichaiMapGet(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := NewOrderedUtxoCollection()
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m.Add(outpoint, &entry)
//		}
//		b.StartTimer()
//
//
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_, _ = m.Get(outpoint)
//		}
//	}
//}

// func BenchmarkCevarisMapGet(b *testing.B) {
//
// 	for n := 0; n < b.N; n++ {
// 		m := cevaris.NewOrderedMap()

// 		b.StopTimer()
// 		for i := 0; i < ITERS; i++ {
// 			outpoint := Outpoint{}
// 			outpoint.Index = uint32(i)
// 			entry := UTXOEntry{}
// 			entry.amount = uint64(i)
// 			entry.blockBlueScore = uint64(i) * 2
// 			m.Set(outpoint, &entry)
// 		}
// 		b.StartTimer()
// 		for i := 0; i < ITERS; i++ {
// 			outpoint := Outpoint{}
// 			outpoint.Index = uint32(i)
// 			_, _ = m.Get(outpoint)
// 		}
// 	}
// }

//func BenchmarkMapGetContains(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := utxoCollection{}
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m[outpoint] = &entry
//		}
//		b.StartTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_ = m[outpoint]
//		}
//	}
//}
//
//func BenchmarkElliotContains(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := elliot.NewOrderedMap()
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m.Set(outpoint, &entry)
//		}
//		b.StartTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_, _ = m.Get(outpoint)
//		}
//	}
//}
//
//func BenchmarkW8MapContains(b *testing.B) {
//
//	for n := 0; n < b.N; n++ {
//		m := w8.New()
//
//		b.StopTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			entry := UTXOEntry{}
//			entry.amount = uint64(i)
//			entry.blockBlueScore = uint64(i) * 2
//			m.Set(outpoint, &entry)
//		}
//		b.StartTimer()
//		for i := 0; i < ITERS; i++ {
//			outpoint := Outpoint{}
//			outpoint.Index = uint32(i)
//			_, _ = m.Get(outpoint)
//		}
//	}
//}

func getOutpointEntry(i uint32) (Outpoint, UTXOEntry) {
	outpoint := Outpoint{Index: i}
	entry := UTXOEntry{}
	entry.amount = uint64(i)
	entry.blockBlueScore = uint64(i) * 2
	entry.scriptPubKey = make([]byte, 37)
	return outpoint, entry
}

func BenchmarkMapRemoveFirst(b *testing.B) {

	m := utxoCollection{}
	var i uint32
	for i = 0; i < ITERS; i++ {
		outpoint, entry := getOutpointEntry(i)
		m[outpoint] = &entry
	}
	firstOutpoint := Outpoint{Index: 0}

	for n := 0; n < b.N; n++ {
		delete(m, firstOutpoint)
		b.StopTimer()
		firstOutpoint.Index += 1
		newOutpoint, newEntry := getOutpointEntry(i)
		m[newOutpoint] = &newEntry
		i += 1
		b.StartTimer()
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkMapRemove**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		// os.Stdout.Sync()

	}
}

func BenchmarkElichaiMapRemoveFirst(b *testing.B) {
	m := NewOrderedUtxoCollection()

	var i uint32
	for i = 0; i < ITERS; i++ {
		outpoint, entry := getOutpointEntry(i)
		m.Add(outpoint, &entry)
	}

	firstOutpoint := Outpoint{Index: 0}

	for n := 0; n < b.N; n++ {
		m.Remove(firstOutpoint)
		b.StopTimer()
		firstOutpoint.Index += 1
		newOutpoint, newEntry := getOutpointEntry(i)
		m.Add(newOutpoint, &newEntry)
		i += 1
		b.StartTimer()
		// var r runtime.MemStats
		// runtime.ReadMemStats(&r)
		// os.Stdout.Sync()
		// fmt.Printf("\n**BenchmarkElichaiMapRemove**: HeapAlloc: %d, TotalAlloc: %d, Mallocs: %d, Frees: %d StackSys: %d, StackInuse: %d, Sys: %d\n", r.HeapAlloc, r.TotalAlloc, r.Mallocs, r.Frees, r.StackSys, r.StackInuse, r.Sys)
		// os.Stdout.Sync()

	}
}