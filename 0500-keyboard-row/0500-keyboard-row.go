import "strings"
func findWords(words []string) []string {
    if len(words) == 0{
        return []string{}
    }
    result := []string{}
    dic := make(map[int][]string)
    dic[1] = []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"}
    dic[2] = []string{"A", "S", "D", "F", "G", "H", "J", "K", "L"}
    dic[3] = []string{"Z", "X", "C", "V", "B", "N", "M"}
    for _, word := range words{
        for i:=1;i<4;i+=1{
            dict := dic[i]
            glb_flag :=1
            for _, c := range word{
                flag:=0
                for _, v := range dict {
                    if strings.ToLower(v) == strings.ToLower(string(c)){
                        flag =1
                        break
                    }
                }
                glb_flag = glb_flag * flag   
            }
            if glb_flag == 1{
               result = append(result, word)       
               break
            }
               
            
        }
        
    }
    
    return result
}
