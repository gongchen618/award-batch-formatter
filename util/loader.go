package util

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/xuri/excelize/v2"
	"strconv"
	"unicode"
)

type StudentAwardCollection struct {
	Name  string
	Award []string
}

func LoadData(file string) []StudentAwardCollection {
	f, err := excelize.OpenFile(file)
	if err != nil {
		log.Error(err.Error())
	}
	contestInfoData := loadContest(f, "contest")
	studentAwardData := loadStudent(f, "data")
	studentAwardAdditionData := loadStudent(f, "contest+")

	var studentAwardCollection []StudentAwardCollection
	for sid, student := range studentAwardData {
		studentAwardP := StudentAwardCollection{
			Name: student.Name,
		}
		for i, award := range student.Award {
			if len(award) == 0 {
				continue
			}
			awardStr := ""
			awardRune := []rune(award)
			for j := 0; j < len(awardRune); j++ {
				err := unicode.IsNumber(awardRune[j])
				if err == false {
					awardStr = fmt.Sprintf("%s%s", awardStr, string(awardRune[j]))
				} else {
					num, _ := strconv.Atoi(string(awardRune[j]))
					awardStr = fmt.Sprintf("%s%s", awardStr, contestInfoData[i].AwardLevel[num])
				}
			}
			studentAwardP.Award = append(studentAwardP.Award, fmt.Sprintf("%s%s", contestInfoData[i].FullName, awardStr))
		}
		for _, award := range studentAwardAdditionData[sid].Award {
			if len(award) == 0 {
				continue
			}
			studentAwardP.Award = append(studentAwardP.Award, award)
		}
		studentAwardCollection = append(studentAwardCollection, studentAwardP)
	}

	return studentAwardCollection
}

type contestInfo struct {
	FullName   string
	AwardLevel []string
}

func loadContest(f *excelize.File, sheet string) []contestInfo {
	columns, err := f.GetCols(sheet)
	if err != nil {
		log.Error(err.Error())
	}
	var contestInfoRes []contestInfo
	for _, col := range columns[1:] {
		if len(col[1]) == 0 {
			continue
		}
		info := contestInfo{
			FullName:   col[1],
			AwardLevel: col[2:],
		}
		contestInfoRes = append(contestInfoRes, info)
	}
	return contestInfoRes
}

type studentAward struct {
	Name  string
	Award []string
}

func loadStudent(f *excelize.File, sheet string) []studentAward {
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Error(err.Error())
	}
	var studentAwardRes []studentAward
	for _, row := range rows[1:] {
		if len(row) == 0 || len(row[0]) == 0 {
			continue
		}
		info := studentAward{
			Name:  row[0],
			Award: row[1:],
		}
		studentAwardRes = append(studentAwardRes, info)
	}
	return studentAwardRes
}
