/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file example1.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Mar 11 00:01:09 2015
 *
 **/

package main

import (
	"fmt"
	"gonlpir"
)

func main() {
	n, err := gonlpir.NewNLPIR("/home/wu/workspace/gopath/src/go-execise/gonlpir/ictclas", gonlpir.UTF8, "0")
	if err != nil {
		panic(err)
	}
	fmt.Println(n.ParagraphProcess("我是中国人", true))
	fmt.Println(n.ParagraphProcess("我是中国人", false))
	results := n.ParagraphProcessA("我是中国人", true)
	for i := 0; i < len(results); i++ {
		fmt.Printf("%#v\n", results[i])
	}
	n.Exit()
}
