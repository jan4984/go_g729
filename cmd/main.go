package main

import (
	"flag"
	"os"
	"io"
	"encoding/binary"
	"github.com/jan4984/go_g729"
	"fmt"
)

func main(){
	f := "-"
	fo := "-"
	flag.StringVar(&f, "input", f, "input g729b file with format 21 6b xx xx <data> frame, - for stdin")
	flag.StringVar(&fo, "output", fo, "output pcm file, - for stdout")
	flag.Parse()
	var inF,outF *os.File
	var err error

	if f == "-" {
		inF = os.Stdin
	}else{
		inF, err = os.Open(f)
		if err != nil{
			panic(err)
		}
		defer inF.Close()
	}

	if fo == "-" {
		inF = os.Stdout
	}else{
		outF, err = os.Create(fo)
		if err != nil{
			panic(err)
		}
		defer outF.Close()
	}

	decoder := go_g729.NewDecoder()
	defer decoder.Destroy()
	h:= make([]byte, 4)
	for {
		_, err =io.ReadFull(inF, h)
		if err != nil {
			fmt.Println("quit for reading input:", err)
			break
		}
		l := ((binary.LittleEndian.Uint32(h) & 0xFFFF0000) >> 16) / 8
		d := make([]byte, l)
		_, err = io.ReadFull(inF, d)
		if err != nil {
			panic(err)
		}
		pcm := decoder.Decode(d)
		outF.Write(pcm)
	}
}
