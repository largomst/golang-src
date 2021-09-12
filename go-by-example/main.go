package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// 变量
	var a int
	fmt.Println(a)

	var b int = 42
	fmt.Println(b)

	var d, e, f = 0, 1, 2
	fmt.Println(d, e, f)

	c := 42
	fmt.Println(c)

	// 常量
	const g = 42
	fmt.Println(g)

	const h float32 = 42
	fmt.Println(h)

	// 循环
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	// 条件
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if j := -1; j < 0 {
		fmt.Println(j, "is negative")
	} else {
		fmt.Println(j, "is positive")
	}

	// switch
	j := 2
	switch j {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	switch {
	case j == 3:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'am bool")
		case int:
			fmt.Println("I'am int")
		case string:
			fmt.Println("I'am string")
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("12")

	// 数组
	var k [3]int
	fmt.Println(k)

	var l = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(l)

	var l2 [3]int = [3]int{1, 2, 3}
	fmt.Println(l2)

	var twoD = [...][3]int{{1, 2, 3}, {2, 3, 4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(twoD[i][j])
		}
		fmt.Println()
	}
	fmt.Println(twoD)

	// 切片
	m := make([]string, 3)
	fmt.Println(m)
	fmt.Printf("%T", m)

	m[0] = "a"
	m[1] = "b"
	m[2] = "c"
	fmt.Println(m)

	m = append(m, "d")
	fmt.Println(m)

	n := make([]string, len(m))
	copy(n, m)
	fmt.Println(n)

	o := []string{"a", "b", "c"}
	fmt.Printf("%T", o)

	// twoD2 := make([][]int, 3)
	// var twoD2 = make([][]int, 3)
	var twoD2 [][]int = make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println(twoD2)

	// hashmap
	p := make(map[int]string)
	p[1] = "tencent"
	p[2] = "文因互联"
	// p[3] = ""
	fmt.Println(p)
	fmt.Println(len(p))

	value1, prs1 := p[1]
	value2, prs2 := p[2]
	value3, prs3 := p[3]
	fmt.Println(value1, prs1)
	fmt.Println(value2, prs2)
	fmt.Println(value3, prs3)

	delete(p, 1)
	fmt.Println(p)

	// for range
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

	for i, num := range nums {
		fmt.Println(i, num)
	}

	kvs := map[string]string{"a": "apple", "m": "microsoft"}
	for k, v := range kvs {
		fmt.Printf("%s->%s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("%s\n", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// 函数
	res := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(res)
	res = plus3(1, 2, 3)
	fmt.Println(res)

	a, b = values()
	fmt.Println(a, b)

	// ...
	res = sum(1, 2, 3, 4)
	fmt.Println(res)

	numbers := []int{1, 2, 3, 4, 5}
	res = sum(numbers...)
	fmt.Println(res)

	// 闭包
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 递归
	for i := 0; i < 30; i++ {
		fmt.Println(fib(i))
	}

	// 指针
	s := 42
	fmt.Println(s)

	zeroval(s)
	fmt.Println(s)

	zeroptr(&s)
	fmt.Println(s)

	fmt.Println(&s)

	// 结构体
	fmt.Println(person{"RS", 20})
	fmt.Println(person{name: "RS", age: 20})
	fmt.Println(person{name: "RS"})
	fmt.Println(&person{name: "RS"})

	p1 := person{name: "RS", age: 20}
	fmt.Println(p1.name)

	var sp1 *person = &p1
	// sp1 := &p1
	fmt.Println(sp1.name) // . 对结构体指针会自动解引用

	sp1.age += 1
	fmt.Println(p1.age)

	// 方法
	r := rect{width: 10, height: 20}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

	me := person{name: "rui", age: 25}
	fmt.Println(me.age)
	me.grow()
	fmt.Println(me.age)

	mep := &me
	mep.grow()
	fmt.Println(me.age)

	//接口
	r1 := rect{10, 20}
	c1 := circle{10}
	measure(r1)
	measure(c1)

	// 错误处理
	for _, i := range []int{4, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{4, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	// 类型断言
	_, err := f2(42)
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	// interface{}
	body := Body{}
	body.Msg = "5"
	fmt.Printf("%#v %T\n", body.Msg, body.Msg)
	body.Msg = 5
	fmt.Printf("%#v %T\n", body.Msg, body.Msg)

	vars := [10]interface{}{1, 2, "12", true}
	fmt.Println(vars)
	for _, v := range vars {
		fmt.Printf("%#v %T\n", v, v)
	}

}

type Body struct {
	Msg interface{}
}

func plus3(a, b, c int) int {
	return a + b + c
}

func values() (int, int) {
	return 3, 7
}

func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}

}

func zeroval(val int) {
	val = 0
}

func zeroptr(iptr *int) {
	*iptr = 0

}

type person struct {
	name string
	age  int
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height

}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (p *person) grow() {
	p.age += 1
}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant't work with 42")
	} else {
		return arg + 3, nil
	}
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	} else {
		return arg + 3, nil
	}
}

type B interface {
	b()
}

type struct1 struct {
	a int
}

func (s *struct1) b() {
	fmt.Print(s.a)
}
