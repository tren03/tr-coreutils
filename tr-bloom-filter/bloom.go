package main

import "fmt"

type Obj struct {
	Name string
}

const FNV_PRIME = 16777619
const OFFSET_BASIS = 216636261
const NOS = 64

var bloomArray uint64 = 0

func SetBit(n uint64, pos uint32) uint64 {
	n |= (1 << pos)
	return n
}

func ClearBit(n uint64, pos uint32) uint64 {
	mask := ^(uint64(1) << pos)
	n &= mask
	return n
}

func ViewBit(n uint64, pos uint32) int8 {
	if (n & (1 << pos)) != 0 {
		return 1
	}
	return 0
}

func FNVhash(o *Obj) uint32 {
	fmt.Println(o.Name)
	var hash uint32 = OFFSET_BASIS
	for _, v := range o.Name {
		hash = uint32(hash) ^ uint32(v)
		hash = hash * FNV_PRIME
	}
	return hash
}

func Bloom() {
	fmt.Println(bloomArray)
}

func AddItem(b uint64, o Obj) {
	hash := FNVhash(&o)
	hashMod := hash % NOS
	if ViewBit(b, hashMod) == 1 {
		fmt.Println("already marked", hashMod)
		return
	}
	bloomArray = uint64(SetBit(b, hashMod))
	fmt.Println("Marking new bit", hashMod)
}

func CheckItem(b uint64, o Obj) {
	hash := FNVhash(&o)
	hashMod := hash % NOS
	if ViewBit(b, hashMod) == 1 {
		fmt.Println("item may already be set")
		return
	}
	fmt.Println("item available")
}

func main() {
	o := Obj{Name: "jack"}
	Bloom()
	fmt.Println(FNVhash(&o))
	AddItem(bloomArray, o)
	o.Name = "tra"
	AddItem(bloomArray, o)

	fmt.Println("-----")

	o.Name = "jack"
	CheckItem(bloomArray, o)
	o.Name = "tack"
	CheckItem(bloomArray, o)
	o.Name = "alsj"
	CheckItem(bloomArray, o)
	o.Name = "akld"
	CheckItem(bloomArray, o)
	o.Name = "tra"
	CheckItem(bloomArray, o)
	
}
