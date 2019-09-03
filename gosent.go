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
 * @file gonlpir.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Mar  7 13:36:13 2015
 *
 **/

package gonlpir

/*
#cgo linux CFLAGS: -DOS_LINUX -I./sent/include
#cgo linux LDFLAGS: -L./sent/lib/linux64 -lSentimentNew

#include <stdio.h>
#include <stdlib.h>
#include <Sentiment.h>
*/
import "C"
import (
	"unsafe"
)

var global_sent_init bool = false

func NewSent(dataPath string, encoding int, licence string) *NLPIR {
	p := &NLPIR{}

	p.dataPath = dataPath
	p.encoding = encoding
	p.licence = licence

	d := C.CString(dataPath)
	defer C.free(unsafe.Pointer(d))

	l := C.CString(licence)
	defer C.free(unsafe.Pointer(l))

	C.ST_Init(d, C.int(encoding), l)

	global_sent_init = true

	return p
}
func (this *NLPIR) ExitSent() {
	if global_sent_init {
		C.ST_Exit()
	}
	global_sent_init = false
}

func (this *NLPIR) GetSentimentPoint(text string) int {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	return int(C.ST_GetSentimentPoint(t))
}
