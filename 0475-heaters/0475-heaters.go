func findRadius(houses []int, heaters []int) int {
    //sort.Ints(houses)
    sort.Ints(heaters)
    var max int
    for i:=0;i<len(houses);i++{
        var r1,r2 int
        left,right:=find(0,len(heaters)-1,houses[i],heaters)
        if left>=0{
        r1=houses[i]-heaters[left]
        }else{
            r1=-1
        }
        if right<len(heaters){
        r2=heaters[right]-houses[i]
        }else{
            r2=-1
        }
        var max1 int
        if r1==-1{
            max1=r2
        }else if r2==-1{
            max1=r1
        }else if r1>r2{
            max1=r2
        }else if r1<=r2{
            max1=r1
        }
        if max1>max{
            max=max1
        }
    }
    return max
}
func find(left int,right int,target int,heaters []int)(int,int){
    if left==right{
        if heaters[left]==target{
            return left,left
        }else if heaters[left]<target{
            return left,left+1
        }else if heaters[left]>target{
            return left-1,left
        }
    }else if left<right{
        mid:=(left+right)/2
        if heaters[mid]==target{
            return mid,mid
        }
        if heaters[mid]>target{
            return find(left,mid,target,heaters)
        }else{
            return find(mid+1,right,target,heaters)
        }
    }
    return -1,len(heaters)
}
