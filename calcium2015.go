package solution

// you can also use imports, for example:
//import "fmt"
import "math"
import "sort"

// helper sort interface imp. to do map sort(2d sllice sort)
type ByDepth [][]int

func (s ByDepth) Len() int {
    return len(s)
}
func (s ByDepth) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByDepth) Less(i, j int) bool {
    return s[i][1] > s[j][1]
}

// helper: find max value in slice
func maxInSlice(slice []int) int {
    max := slice[0]
    for _, value := range slice {
        if value > max {
            max = value
        }
    }
    return max
}

// find min no. of cameras needed to achive diameter
func minCamsNeeded(nodes [][]int, treeDepth []int, root int, diameter int) int {
    
    if(len(nodes[root])==0){
        treeDepth[root] = 0
        return 0
    } 
    
    children := make([]int, len(nodes[root]))
    copy(children , nodes[root])
    
    cams := 0
    for _, child := range children {
        cams += minCamsNeeded(nodes, treeDepth, child, diameter)
    }
    
    childrenDepths := make([][]int, len(children))
    for i, child := range children {
        childrenDepths[i] = make([]int, 2)
        childrenDepths[i][0] = child
        childrenDepths[i][1] = treeDepth[child]
    }
    
    sort.Sort(ByDepth(childrenDepths))

    var idxToRm int  
    for i := 0; i < len(childrenDepths)-1; i++ {
        if((childrenDepths[i][1] + childrenDepths[i+1][1] +2) > diameter){
            idxToRm = posInSlice(children, childrenDepths[i][0])
            children = append(children[:idxToRm], children[idxToRm+1:]...)
            cams += 1   
        } 
    }
    
    if((childrenDepths[len(childrenDepths)-1][1] + 1) > diameter){
        //remove
        idxToRm = posInSlice(children, childrenDepths[len(childrenDepths)-1][0])
        children = append(children[:idxToRm], children[idxToRm+1:]...)
        cams += 1   
    } 
    
    if len(children) == 0 {
        treeDepth[root] = 0
    } else {
        depths := make([]int, 0)
        for _, child := range children {
            depths = append(depths,treeDepth[child])
        }
        treeDepth[root] = maxInSlice(depths) + 1
    }
    
    return cams
}

// helper: find min value in slice
func minInSlice(slice []int) int {
    min := slice[0]
    for _, value := range slice {
        if value < min {
            min = value
        }
    }
    return min
}

// helper: find position of a value in a slice
func  posInSlice(slice []int, value int) int {
    for p, v := range slice {
        if (v == value) {
            return p
        }
    }
    return -1
}

func Solution(A []int, B []int, K int) int {
    // write your code in Go 1.4
    /*steps
    1. change to adjacency list rep. & make it a tree
    2. minimize diameter
    */
    
    // change graph rep. to to adj. list 
    n := len(A)
    nodes := make([][]int , n+1)
    
    // adjacency list rep. of input graph
    for i:=0; i<n; i++  { 
        nodes[A[i]] = append(nodes[A[i]], B[i])
        nodes[B[i]] = append(nodes[B[i]], A[i])
    }
       
    /* remove back references from the adj. list
    *  and treat it as a tree 
    */
    q := make([]int, 1)
    var parent, childIdx int
    for len(q) > 0{  
        parent, q = q[0], q[1:]
        for _, node := range nodes[parent]{
            childIdx = posInSlice(nodes[node], parent)
            nodes[node] = append(nodes[node][:childIdx], nodes[node][childIdx+1:]...)
            q = append(q,node)
        }
        
    }
    
    treeDepth := make([]int, n+1)
    first := 0
    last := int(math.Min(float64(900), float64(n)))
    candidates := make([]int, 0)
    cams := 0; diameter := 0 
    
    // minimize diameter
    for first <= last {
        diameter = int((first + last)/2)
        cams = minCamsNeeded(nodes, treeDepth, 0, diameter)
        if(cams > K){
            first = diameter+1
        } else {
            candidates = append(candidates, diameter)
            last = diameter-1
        }
    } 
    
    return minInSlice(candidates)
}
