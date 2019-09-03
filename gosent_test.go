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
	sent := NewSent("./sent", UTF8, "0")
	score := sent.GetSentimentPoint("【今日关注】<font color='red'>中法</font><font color='red'>聚变</font>联合研究中心揭牌成立")
	fmt.Println(score)
}

//===================================================================
// Private
//===================================================================
