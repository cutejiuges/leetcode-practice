package LC1600

/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/7 下午10:56
 * @FilePath: throne_inheritance
 * @Description:
 */

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{king: kingName, edges: map[string][]string{}, dead: map[string]bool{}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.edges[parentName] = append(this.edges[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var preOrder func(string)
	var ans []string
	preOrder = func(name string) {
		if !this.dead[name] {
			ans = append(ans, name)
		}
		for _, child := range this.edges[name] {
			preOrder(child)
		}
	}
	preOrder(this.king)
	return ans
}
