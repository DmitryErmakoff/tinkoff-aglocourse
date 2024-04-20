package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	q1, q2, qm1, qm2 *list.List
}

func (q *Queue) push(v, i int) {
	q.q1.PushBack([]int{v, i})
	if q.qm1.Len() == 0 {
		q.qm1.PushBack([]int{v, i})
	} else {
		if v > q.qm1.Back().Value.([]int)[0] {
			q.qm1.PushBack([]int{v, i})
		} else {
			q.qm1.PushBack([]int{q.qm1.Back().Value.([]int)[0], q.qm1.Back().Value.([]int)[1]})
		}
	}
}

func (q *Queue) pop() {
	if q.q2.Len() == 0 {
		for q.q1.Len() > 0 {
			value := q.q1.Back().Value.([]int)
			q.q2.PushBack([]int{value[0], value[1]})
			if q.qm2.Len() == 0 {
				q.qm2.PushBack([]int{value[0], value[1]})
			} else {
				if q.qm2.Back().Value.([]int)[0] > value[0] {
					q.qm2.PushBack([]int{q.qm2.Back().Value.([]int)[0], q.qm2.Back().Value.([]int)[1]})
				} else {
					q.qm2.PushBack([]int{value[0], value[1]})
				}
			}
			q.q1.Remove(q.q1.Back())
			q.qm1.Remove(q.qm1.Back())
		}
	}
	//value := q.q2.Back().Value.([]int)
	q.q2.Remove(q.q2.Back())
	q.qm2.Remove(q.qm2.Back())
}

func (q *Queue) min() []int {
	if q.q1.Len() == 0 {
		return q.qm2.Back().Value.([]int)
	}
	if q.q2.Len() == 0 {
		return q.qm1.Back().Value.([]int)
	}
	if q.qm1.Back().Value.([]int)[0] < q.qm2.Back().Value.([]int)[0] {
		return q.qm2.Back().Value.([]int)
	}
	return q.qm1.Back().Value.([]int)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 1; i < n-1; i++ {
		fmt.Scan(&a[i])
	}

	dp_otv := make([][]int, n)
	dp := Queue{list.New(), list.New(), list.New(), list.New()}
	dp.push(0, 0)
	dp_otv[0] = []int{0, -1}

	for i := 1; i < min(n, k); i++ {
		dp_otv[i] = []int{dp.min()[0] + a[i], dp.min()[1]}
		dp.push(dp_otv[i][0], i)
	}

	for i := k; i < n; i++ {
		dp_otv[i] = []int{dp.min()[0] + a[i], dp.min()[1]}
		dp.pop()
		dp.push(dp_otv[i][0], i)
	}

	fmt.Println(dp_otv[n-1][0])
	i := n - 1
	var otv []int
	for i != -1 {
		otv = append(otv, i+1)
		i = dp_otv[i][1]
	}

	for i := 1; i < len(otv); i++ {
		fmt.Print(otv[i], " ")
	}
}
