package utils

import (
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	unitv4 "tesseract/pkg/apis/unit/v1alpha1"
)

func GetUnitMetricValue(metrics []unitv4.UnitMetricStatus, key string, comontfilters map[string]string) (string, bool) {
	for i := range metrics {
		if metrics[i].Key == key {
			if len(comontfilters) == 0 {
				return metrics[i].Value, true
			}

			support := true
			for k, v := range comontfilters {
				if metrics[i].Comment[k] != v {
					support = false
					break
				}
			}
			if support {
				return metrics[i].Value, true
			}
		}
	}

	return "", false
}

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func ContainsStringV2(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

// MatchUsername username
func MatchUsername(s string) bool {
	//pattern := "^[a-zA-Z0-9][a-zA-Z0-9_-]{1,32}$"
	pattern := "[a-zA-Z0-9\\\\_\\\\-\\\\.]{1,32}$"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

// MatchDBname dbname
func MatchDBname(s string) bool {
	pattern := "[a-zA-Z0-9\\\\_\\\\-\\\\.]{1,64}$"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func MatchServiceName(s string) bool {
	pattern := "^[a-zA-Z][a-z0-9A-Z-]{4,29}$"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func MatchNameV2(s string) bool {
	pattern := "[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func MatchTopicName(s string) bool {
	pattern := "[a-zA-Z0-9\\\\_\\\\-\\\\.]{1,255}$"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

// MatchString namespace secret正则
func MatchString(s string) bool {
	pattern := "[a-z]([-a-z0-9]*[a-z0-9])?"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

// return which in p1 and not in p2
// eg:
// ps1 := []string{"select, insert, update"}
// ps2 := []string{"insert, update, delete"}
// diffPrivileges(ps1, ps2) = []string{"select"}
func StringsDiffFunc(ps1, ps2 []string) []string {
	sort.Strings(ps1)
	sort.Strings(ps2)
	diff := make([]string, 0, 50)
	for _, p1 := range ps1 {
		found := false
		for _, p2 := range ps2 {
			if strings.Trim(p1, " ") == strings.Trim(p2, " ") {
				found = true
			}
		}
		if !found {
			diff = append(diff, strings.Trim(p1, " "))
		}
	}
	return diff
}

func RemoveRepeatElement(s []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

// RandSeq 生成小写字母和数字随机组合
func RandSeq(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(36)]
	}
	return string(b)
}
