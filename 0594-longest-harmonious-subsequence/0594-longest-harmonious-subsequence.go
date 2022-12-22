func findLHS(nums []int) int {
    god := make(map[int]int);
    for _, val := range nums {
        god[val] += 1;
    }
    
    ans := 0;
    for key, _ := range god {
        _, ok := god[key-1]
        if !ok {
            continue;
        }
        sum := god[key-1] + god[key];
        if sum > ans {
            ans = sum;
        }
    }
    
    return ans;
}