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
 * @file gonlpir_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Mar  8 15:40:16 2015
 *
 **/

package gonlpir

import (
	"fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

func TestInit(t *testing.T) {
	sent, err := NewSent("./sent", UTF8, "0")
	if err != nil {
		t.Error(err)
		return
	}
	score := sent.GetSentimentPoint("我非常恨美国政府，美国政府没有人性！他们挑起战争，伤害无辜！")
	fmt.Println(score)
}

//===================================================================
// Private
//===================================================================
