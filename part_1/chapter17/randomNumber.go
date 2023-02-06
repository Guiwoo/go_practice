package chapter17

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ReadFastAsap() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	read, _, _ := reader.ReadLine()

	writer.Write(read)
	defer writer.Flush()
}
func TimePractice() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	const LongForm = "Jan 2, 2006 at 11:00pm"
	t1, _ := time.ParseInLocation(LongForm, "Feb 5, 2023 at 11:00pm", loc)
	fmt.Println(t1, t1.Location(), t1.UTC())

	const shortForm = "2023-Feb-05"
	t2, _ := time.Parse(shortForm, "2023-Feb-05")
	fmt.Println(t2, t2.Location())

	fmt.Println(t2.Sub(t1))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomNumber() int {
	return rand.Intn(100)
}
func inputValueInt() (int, error) {
	var stdin = bufio.NewReader(os.Stdin)
	var a int
	_, err := fmt.Scanln(&a)
	if err != nil {
		stdin.ReadString('\n')
	}
	return a, err
}

func StartGame() {
	//get random number
	target := getRandomNumber()
	fmt.Println(target, " 답 입니다")
	for {
		a, err := inputValueInt()
		if err != nil {
			fmt.Println("숫자를 입력해주세요")
		}
		if a < target {
			fmt.Println("작습니다.")
		} else if a > target {
			fmt.Println("큽니다")
		} else {
			fmt.Println("정답 입니다.")
			break
		}
	}
	fmt.Println("게임이 끝났습니다.")
}
