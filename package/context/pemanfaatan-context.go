package main

import (
	"context"
	"encoding/json"
	"fmt"
)

// membuat data yang akan masukkan ke context
type Data struct {
	Method string
	Value  string
}

func main() {
	ctx := context.Background()

	dataA := Data{
		Method: "GET",
		Value:  "Mengambil data user pada context A",
	}
	ctxLevel1_A := context.WithValue(ctx, "data", dataA)

	dataB := Data{
		Method: "POST",
		Value:  "Menambah data user pada context B",
	}
	ctxLevel1_B := context.WithValue(ctx, "data", dataB)

	processA(ctxLevel1_A)
	processB(ctxLevel1_B)
}

// tipe data dari context adalah context.Context
func processA(ctx context.Context) {
	var data Data
	fmt.Println("Proses context level 1 A")

	// dataCtx memiliki tipe data interface{}
	dataCtx := ctx.Value("data")

	// jadi dataCtx harus diubah kedalam []byte
	dataCtxByte, err := json.Marshal(dataCtx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// lalu di parsing ke struct
	err = json.Unmarshal(dataCtxByte, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Try process with method :", data.Method)
	fmt.Println(data.Value)
}

func processB(ctx context.Context) {
	var data Data
	fmt.Println("Proses context level 1 B")

	dataCtx := ctx.Value("data")

	dataCtxByte, err := json.Marshal((dataCtx))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(dataCtxByte, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Try process with method :", data.Method)
	fmt.Println(data.Value)
}
