package conf

import (
	"bufio"
	"io"
	"os"
	"strings"
	//"fmt"
)

const mid = "=="

type Config struct {
	m map[string]string
	s string
}

//config example as follow
/*[app]
*#应用运行模式,我们采用了gin框架,目前支持debug/release/test三种
*mode=debug
*#应用的名称,以后扩展,用做应用标识，便于分布式计算
*name=xxx
*#应用部署的访问协议,支持http/https两种
*protocal=http
*#应用域名
*domain=localhost:888
*#静态资源所在的服务器地址,便于动静态分离
*asset=localhost:888/asset
*#请求contextpath
*ctxpath=
*#服务器绑定的地址
*addr=localhost
*#端口
*port=8888
*[session]
*#sessionID标识字符串,对标PHP的SESSIONID,java的JSESSIONID
*name=restgo_session_id
*#session过期时间以秒为单位,0表示访问结束时过期
*timelive=3600
 */

//Unix Style Parse Config File

func (c *Config) Init(path string) {
	c.m = make(map[string]string)
	f, err := os.Open(path) //try to open file
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {

		b, _, err := r.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))

		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")

		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.s = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.s) == 0 {
			continue
		}

		i := strings.Index(s, "=")
		if i < 0 {
			continue
		}

		key := strings.TrimSpace(s[:i])
		if len(key) == 0 {
			continue
		}

		value := strings.TrimSpace(s[i+1:])

		if len(value) == 0 {
			continue
		}

		//mark
		pos := strings.Index(value, "\t#")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, "#")
		if pos > -1 {
			value = value[0:pos]
		}
		pos = strings.Index(value, "\t//")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, "//")
		if pos > -1 {
			value = value[0:pos]
		}

		if len(value) == 0 {
			continue
		}

		k := c.s + mid + key
		c.m[k] = strings.TrimSpace(value)
	}
}

func (c *Config) Read(node, key string) string {
	key = node + mid + key
	v, err := c.m[key]
	if !err {
		return ""
	}
	return v
}
