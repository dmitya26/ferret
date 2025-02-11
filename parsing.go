package main

import (
	"fmt"
	"regexp"
)

/*
{ 0       // leftBracket is incremeneted  // idx stored on queue.
	{ 1   // leftBracket is incrememented  // idx stored on queue.


	} 1  // queue is popped and paired with rightBracket idx.
	{ 2  // leftBracket incremented // idx is stored on queue.
	{ 3  // leftBracket incremeented // idx iss stored on queue.

	} 3  // queue is popped and paired with rightBracket idx.
	} 2  // queue is popped and paired with rightBracket idx.
} 0 // rightBracket is incremented; queue is popped and paired with rightBracket idx.

If the length of the queue at the end is greater than zero, then call a syntax error at the indexes on the remaining values on the queue.
*/

/*
type JsonPreprocessQueue struct {
	leftBracket []int
	pairs       [][]int
	data        map[string][]int
}
*/

type Json struct {
	data  string
	child *Json
}

/*
func (j *JsonPreprocessQueue) queueGetPop() int {
	result := j.leftBracket[len(j.leftBracket)-1]
	j.leftBracket = j.leftBracket[:len(j.leftBracket)-1]
	return result
}

func (j *JsonPreprocessQueue) queuePush(idx int) {
	j.leftBracket = append(j.leftBracket, idx)
}
*/

/*
func resolveBrackets(arg string, j *JsonPreprocessQueue) error {
	for i := range arg {
		if arg[i] == byte('{') {
			j.queuePush(i)
		} else if arg[i] == byte('}') {
			idx := j.queueGetPop()
			j.pairs = append(j.pairs, []int{idx, i})
		}

	}

	if len(j.leftBracket) > 0 {
		leftBracket := make([]string, len(j.leftBracket))
		for i, c := range j.leftBracket {
			leftBracket[i] = strconv.Itoa(c)
		}

		return errors.New(fmt.Sprintf("Unresolved bracket(s): ", strings.Join(leftBracket, ", ")))
	}
	return nil
}

func resolveData(arg string, j *JsonPreprocessQueue) error {
	for i := range arg {
		if arg[i]
	}
	return nil
}
*/

func UnmarshalJson(raw string) (*Json, error) {
	json := &Json{}

	var preprocessed string
	for i := range raw {
		if raw[i] != byte(' ') || raw[i] != byte('\n') {
			preprocessed += string(raw[i])
		}
	}

	re := regexp.MustCompile(`:\{|,\s*`)
	preprocessed_split := re.Split(preprocessed, -1)
	for i := range preprocessed_split {
		fmt.Println(preprocessed_split[i])
	}

	return json, nil
}

func main() {
	UnmarshalJson("{'heyo':{'hello':'hello'},'epic!':{},{}}")
}
