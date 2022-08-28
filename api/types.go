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

package api

type RankListResult struct {
	Data    RankListData `json:"data"`
	Code    string       `json:"code"`
	TraceID string       `json:"traceId"`
}

type RankListData struct {
	List   []RankList `json:"list"`   // 排行榜列表
	RaceVO RaceVO     `json:"raceVO"` // 赛题相关信息
	Total  int        `json:"total"`  // 目前总共的提交队伍数
}

type RaceVO struct {
	RaceID   int    `json:"raceId"`
	RaceName string `json:"raceName"`
	Brief    string `json:"brief"`
	Bonus    int    `json:"bonus"`
}

type TeamMemberList struct {
	ID                    int     `json:"id"`
	GmtCreate             string  `json:"gmtCreate"`
	GmtModified           string  `json:"gmtModified"`
	UserID                int64   `json:"userId"`
	NickName              string  `json:"nickName"`
	GraduateSchool        string  `json:"graduateSchool"`
	Organization          string  `json:"organization"`
	IsStudent             bool    `json:"isStudent"`
	TeamID                int     `json:"teamId"`
	IsLeader              bool    `json:"isLeader"`
	RaceID                int     `json:"raceId"`
	Status                int     `json:"status"`
	JoinType              int     `json:"joinType"`
	JoinSuccess           bool    `json:"joinSuccess"`
	LatestSore            float64 `json:"latestSore"`
	ImperialKitchenStatus int     `json:"imperialKitchenStatus"`
}

type ScoreJSONObject struct {
	Score              string `json:"score"`
	ScoreResourceUsage string `json:"score_resourceUsage"`
	ScoreConsistency   string `json:"score_consistency"`
	ScoreFilter        string `json:"score_filter"`
	ScoreFunctional    string `json:"score_functional"`
}

type RankList struct {
	TeamLeaderOrganization string           `json:"teamLeaderOrganization"`
	TeamType               int              `json:"teamType"`
	TeamMemberList         []TeamMemberList `json:"teamMemberList"`
	MobileScore            string           `json:"mobileScore"`
	Trends                 int              `json:"trends"`
	ScoreID                int              `json:"scoreId"`
	RaceID                 int              `json:"raceId"`
	ScoreJSONObject        ScoreJSONObject  `json:"scoreJsonObject,omitempty"`
	GmtSubmit              string           `json:"gmtSubmit"`
	SeasonEnd              bool             `json:"seasonEnd"`
	ScoreJSON              string           `json:"scoreJson"`
	RateTaskID             int              `json:"rateTaskId"`
	Season                 int              `json:"season"`
	TeamID                 int              `json:"teamId"`
	TeamName               string           `json:"teamName"`
	Rank                   int              `json:"rank"`
	GmtCreate              string           `json:"gmtCreate"`
	GmtModified            string           `json:"gmtModified"`
	Score                  float64          `json:"score"`
	ID                     int              `json:"id"`
}
