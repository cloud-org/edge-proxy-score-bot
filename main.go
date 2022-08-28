/*
 * MIT License
 *
 * Copyright (c) 2022 cloud-org Authors
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"bytes"
	"context"
	"edge-proxy-score-bot/api"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cloud-org/msgpush"
)

var (
	dingtalk string // dingtalk 通知链接
	help     bool   // 帮助
)

func init() {
	flag.StringVar(&dingtalk, "dingtalk", "", "dingtalk webhook url")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stdout, `edge-proxy-score-bot - rankList push
Usage: edge-proxy-score-bot [-h help]
Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	// you can modify this link for other tianchi race
	url := "https://tianchi.aliyun.com/competition/proxy/api/competition/api/race/season/ranking/list?pageNum=1&pageSize=5&season=-1&raceId=531984&userId=-1"
	resp, err := api.GetRankList(context.TODO(), url)
	if err != nil {
		log.Printf("获取排行榜列表失败: %v\n", err)
		return
	}

	if resp.Code != "SUCCESS" {
		log.Printf("code 不为 SUCCESS\n")
		return
	}

	raceName := resp.Data.RaceVO.RaceName
	// rank1: xx score: xx
	var temp bytes.Buffer
	temp.WriteString(raceName)
	temp.WriteString("\n")
	for i := 0; i < len(resp.Data.List); i++ {
		item := resp.Data.List[i]
		content := fmt.Sprintf("rank%d: %s, score: %v, resourceusage: %v\n", item.Rank, item.TeamName, item.Score, item.ScoreJSONObject.ScoreResourceUsage)
		temp.WriteString(content)
	}

	msg := temp.String()
	log.Println(msg)

	if dingtalk != "" {
		w := msgpush.NewDingTalk(dingtalk)
		_ = w.SendText(msg)
	}

	return
}
