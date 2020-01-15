package main

import (
	"fmt"
	"math"
)

func Paginator(CurrentPage, PageSize int, nums int64) map[string]interface{} {

	var LastPage int   //前一页地址
	var NextPage int   //后一页地址
	var StartIndex int //当前页起始索引
	var EndIndex int   //当前页结束索引
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(PageSize))) //page总数
	if CurrentPage > totalpages {
		CurrentPage = totalpages
	}
	if CurrentPage <= 0 {
		CurrentPage = 1
	}
	var pages []int
	switch {
	case CurrentPage >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		LastPage = CurrentPage - 1
		NextPage = int(math.Min(float64(totalpages), float64(CurrentPage+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
		// StartIndex = LastPage * PageSize
		// EndIndex = CurrentPage*PageSize - 1
		// if EndIndex > int(nums)-1 {
		// 	EndIndex = int(nums) - 1
		// }

	case CurrentPage >= 3 && totalpages > 5:
		start := CurrentPage - 3 + 1
		pages = make([]int, 5)
		LastPage = CurrentPage - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		LastPage = CurrentPage - 1
		NextPage = CurrentPage + 1
		// StartIndex = LastPage * PageSize
		// EndIndex = CurrentPage*PageSize - 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		LastPage = int(math.Max(float64(1), float64(CurrentPage-1)))
		NextPage = int(math.Min(float64(CurrentPage+1), float64(totalpages)))
		EndIndex = CurrentPage*PageSize - 1
		StartIndex = CurrentPage*PageSize - PageSize
		if StartIndex < 0 {
			StartIndex = 0
		}
		if EndIndex > int(nums)-1 {
			EndIndex = int(nums) - 1
		}
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["LastPage"] = LastPage
	paginatorMap["NextPage"] = NextPage
	paginatorMap["currpage"] = CurrentPage
	paginatorMap["StartIndex"] = StartIndex
	paginatorMap["EndIndex"] = EndIndex
	return paginatorMap
}

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println("a= ", a)
	PageSize := 3
	var rsCount int64 = 10
	CurrentPage := 4
	result := Paginator(CurrentPage, PageSize, rsCount)
	StartIndex := result["StartIndex"].(int)
	EndIndex := result["EndIndex"].(int)
	fmt.Println(a[StartIndex:EndIndex])
	fmt.Println(result)
}
