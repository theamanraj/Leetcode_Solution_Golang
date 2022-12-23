func nextGreaterElement(n int) int {
   if n<12 { return -1}
   
   nums :=[]byte(strconv.Itoa(n))
   p, k := -1, -1
   for i:=len(nums)-2;i>=0;i--{ if nums[i]<nums[i+1]{ p=i;break} } // look out the pivoit
   if p==-1{ return -1}
   // look out k, which is the one bigger than nums[p]
   for i:=len(nums)-1;i>p;i--{ if nums[i]>nums[p]{  k=i;break }   }
   
   nums[p],nums[k]=nums[k],nums[p] // switch
   // reverse from p+1 till end.  before: desc, after asc
   reverse(nums[p+1:])           
   if t, err := strconv.Atoi(string(nums)); err!=nil || t>math.MaxInt32 {
       return -1
   }else{
       return t
   }
}

func reverse( nums []byte) {
   n:=len(nums)
   for i:=0;i<len(nums)/2;i++{ nums[i], nums[n-1-i]= nums[n-1-i], nums[i]}        
} 