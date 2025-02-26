package LC1472

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/26 下午11:24
 * @FilePath: browser_history
 * @Description: 模拟实现浏览器访问记录
 */

type BrowserHistory struct {
	urls []string
	idx  int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		urls: []string{homepage},
		idx:  0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	for len(this.urls) > this.idx+1 {
		this.urls = this.urls[:this.idx+1]
	}
	this.urls = append(this.urls, url)
	this.idx++
}

func (this *BrowserHistory) Back(steps int) string {
	this.idx = max(0, this.idx-steps)
	return this.urls[this.idx]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.idx = min(len(this.urls)-1, this.idx+steps)
	return this.urls[this.idx]
}
