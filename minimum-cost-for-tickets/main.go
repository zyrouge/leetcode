package main

// func mincostTickets(days []int, costs []int) int {
// 	n := len(days)
// 	m := make(map[int]int)
// 	var dfs func(start int) int
// 	dfs = func(start int) int {
// 		if start == n {
// 			return 0
// 		}
// 		if found, ok := m[start]; ok {
// 			return found
// 		}
// 		i := start + 1
// 		total := costs[0] + dfs(i)
// 		for i < n && days[i]-days[start] < 7 {
// 			i++
// 		}
// 		if i > start {
// 			total = min(total, costs[1]+dfs(i))
// 		}
// 		for i < n && days[i]-days[start] < 30 {
// 			i++
// 		}
// 		if i > start {
// 			total = min(total, costs[2]+dfs(i))
// 		}
// 		m[start] = total
// 		fmt.Println(start, total)
// 		return total
// 	}
// 	return dfs(0)
// }

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	days = append(days, days[n-1]+30)
	totals := make([]int, n+1)
	day7, day30 := n, n
	for i := n - 1; i >= 0; i-- {
		for day7 > i+1 && days[day7-1]-days[i] >= 7 {
			day7--
		}
		for day30 > i+1 && days[day30-1]-days[i] >= 30 {
			day30--
		}
		totals[i] = min(costs[0]+totals[i+1], costs[1]+totals[day7], costs[2]+totals[day30])
	}
	return totals[0]
}
