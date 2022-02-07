# shcaptcha
A golang wrapper around hcaptcha service to bypass captcha challenges using [hcaptcha accessibility cookie](https://www.hcaptcha.com/accessibility).

Requirements
------------
You need an accessibility cookie which you can get by following these steps:
- Go to https://www.hcaptcha.com/accessibility
- Sign up with a valid email.
- Check your inbox to verify your email.
- You'll be redirected to hcaptcha dashboard and the accessibility cookie will be set.

Cookies are only valid for 24 hours and you might get rate limited if you abuse too much (which you probably will).

Basic usage
-------

```golang
package main

import (
	"fmt"
	"log"
	"github.com/Grapphy/shcaptcha"
)

func main() {
	var accessibility_cookie string = "Some accessibility cookie"
	var shcpt *shcaptcha.Client = shcaptcha.NewClient(accessibility_cookie)
	shcpt.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0")

	var sitekey string = "3ceb8624-1970-4e6b-91d5-70317b70b651"
	var url string = "https://2captcha.com/demo/hcaptcha?difficulty=always-on"

	solution, err := shcpt.BypassCaptcha(sitekey, url)
	if err != nil {
		log.Fatal(err)
	}
    
	fmt.Println(solution)
}
```

Using proxies
-------------

```golang
package main

import (
	"fmt"
	"log"
	"github.com/Grapphy/shcaptcha"
)

func main() {
	var accessibility_cookie string = "Some accessibility cookie"
	var shcpt *shcaptcha.Client = shcaptcha.NewClient(accessibility_cookie)
	shcpt.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0")
    
	proxy := "http://host:port"
	shcpt.SetProxy(proxy)
    
	var sitekey string = "3ceb8624-1970-4e6b-91d5-70317b70b651"
	var url string = "https://2captcha.com/demo/hcaptcha?difficulty=always-on"
    
	solution, err := shcpt.BypassCaptcha(sitekey, url)
	if err != nil {
		log.Fatal(err)
	}
    
	fmt.Println(solution)
}
```

License
-------
This repository is under [MIT license](https://mit-license.org/).
