package stacks

type CustomStack struct {
    Size int
    Items []int
    Idx int
}


func ConstructorStack(maxSize int) CustomStack {
    return CustomStack{Size:maxSize,Items:make([]int,maxSize),Idx:-1}
}


func (this *CustomStack) Push(x int)  {
    if this.Idx < (this.Size-1){
        this.Idx++
        this.Items[this.Idx] = x
    }
}


func (this *CustomStack) Pop() int {
    x := -1
    if this.Idx > -1{
        x = this.Items[this.Idx]
        this.Idx--
    }
    return x
}


func (this *CustomStack) Increment(k int, val int)  {
    i := 0
    for i < k && i < len(this.Items){
        this.Items[i] = this.Items[i]+val
        i++
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */