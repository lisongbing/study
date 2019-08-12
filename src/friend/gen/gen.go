package gen

import (
	"diabloserver/public"
)

//生成若干个不重复的随机数
func RandomTestBase() {
	//测试5次
	for i := 0; i < 5; i++ {
		//nums := GenerateRandomNumber(10, 30, 10,nil)
		//fmt.Println(nums)
	}
}

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(Seed *int64,start int, end int, count int,current []int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样

	for len(nums) < count {
		//生成随机数
		num := int(public.Random(Seed,int64(start),int64(end)))
		//查重
		exist := false
			for _, v := range current {
				if v == num {
					exist = true
					break
				}
			}

		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
