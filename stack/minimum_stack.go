package stack

type MinStack struct {
	Store    []int
	Minimums []int
}

func Constructor() MinStack {
	return MinStack{
		Store:    make([]int, 0),
		Minimums: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Store = append(this.Store, val)
	if len(this.Minimums) == 0 {
		this.Minimums = append(this.Minimums, val)
	} else {
		if this.Minimums[len(this.Minimums)-1] >= val {
			this.Minimums = append(this.Minimums, val)
		}
	}
}

func (this *MinStack) Pop() {
	deleted, newStore := this.Store[len(this.Store)-1], this.Store[:len(this.Store)-1]
	this.Store = newStore
	if deleted == this.Minimums[len(this.Minimums)-1] {
		this.Minimums = this.Minimums[:len(this.Minimums)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Store[len(this.Store)-1]
}

func (this *MinStack) GetMin() int {
	return this.Minimums[len(this.Minimums)-1]
}
