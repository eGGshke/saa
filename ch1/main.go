package ch1

import (
    "log"
    "math/rand"
)

func Main() {
    log.Printf("Greatest common divisor is: %d\n", Gcd(10, 5))

    var a []int
    for i := 0; i < 101; i++ {
        a = append(a, rand.Intn(i+1))
    }

    log.Printf("Greatest intger in array: %d\n", FindLarge(a))

    log.Printf("Find dulbicates in array: %d\n", FindDuplicates(a))

    log.Printf("Find dividing point: %d\n", DividingPoint(a))

    Tree := &TreeNode{Val: 50}
    for i := 1; i < 101; i++ {
        Tree.Add(rand.Intn(i + 1))
    }
    log.Printf("Max value is: %d \n", Tree.GetMaxValue())
    log.Printf("In binary tree: %t \n", Tree.Exist(3))
}

// Gcd поиск наибольшего общего делителя
func Gcd(a, b int) int {
    for b != 0 {
        remainder := a % b
        a = b
        b = remainder
    }

    return a
}

// FindLarge O(1 + N + 1)
func FindLarge(a []int) int {
    var largest int
    for i := 0; i < len(a); i++ {
        if a[i] > largest {
            largest = a[i]
        }
    }

    return largest
}

// FindDuplicates O(N3)
func FindDuplicates(a []int) any {
    var d []int

    for i := 0; i < len(a); i++ {
        for j := i + 1; j < len(a); j++ {
            if a[i] == a[j] {
                have := false
                for k := 0; k < len(d); k++ {
                    if d[k] == a[i] {
                        have = true
                        break
                    }
                }

                if !have {
                    d = append(d, a[i])
                }
            }
        }
    }

    return d
}

// DividingPoint O(1)
func DividingPoint(a []int) int {
    n1 := a[0]
    n2 := a[len(a)-1]
    n3 := a[int(len(a)-1)/2]

    if n1 < n2 && n1 > n3 {
        return n1
    }

    if n2 > n1 && n2 < n3 {
        return n2
    }

    return n3
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func (n *TreeNode) Add(val int) {
    if n == nil {
        return
    }

    if val == n.Val {
        return
    }

    if val < n.Val {
        if n.Left == nil {
            n.Left = &TreeNode{Val: val}
        } else {
            n.Left.Add(val)
        }
    } else {
        if n.Right == nil {
            n.Right = &TreeNode{Val: val}
        } else {
            n.Right.Add(val)
        }
    }
}

func (n *TreeNode) Exist(need int) bool {
    for {
        if n == nil {
            return false
        }

        if need == n.Val {
            return true
        }

        if need < n.Val {
            n = n.Left
        } else {
            n = n.Right
        }
    }
}

func (n *TreeNode) GetMaxValue() int {
    for {
        if n.Right == nil {
            return n.Val
        }

        n = n.Right
    }
}
