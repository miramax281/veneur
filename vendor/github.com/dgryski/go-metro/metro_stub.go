// +build amd64
// +build !noasm

package metro

//go:generate python -m peachpy.x86_64 metro.py -S -o metro_amd64.s -mabi=goasm
//go:noescape

func Hash64(buffer []byte, seed uint64) uint64
