/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
type MinStack struct {
    top *Node
}

type Node struct {
    value int
    prev *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{nil}
}

func (this *MinStack) Push(x int) {
    node := &Node{x, this.top}
    this.top = node
}

func (this *MinStack) Pop() {
    if this.top == nil {
        return
    }
    this.top = this.top.prev
}

func (this *MinStack) Top() int {
    if this.top == nil {
        return math.MinInt32
    }
    return this.top.value
}

func (this *MinStack) GetMin() int {
    if this.top == nil {
        return math.MinInt32
    }
    min := this.top.value
    temp := this.top
    for temp != nil {
        if temp.value < min {
            min = temp.value
        }
        temp = temp.prev
    }
    return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
