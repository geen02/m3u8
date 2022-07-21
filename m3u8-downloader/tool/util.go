package tool

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(joinPath ...string) (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	p := strings.Replace(dir, "\\", "/", -1)
	whole := filepath.Join(joinPath...)
	whole = filepath.Join(p, whole)
	return whole, nil
}

func ResolveURL(u *url.URL, p string) string {
	if strings.HasPrefix(p, "https://") || strings.HasPrefix(p, "http://") {
		return p
	}
	var baseURL string
	url_string := u.String()
	if strings.Index(p, "/") == 0 {
		baseURL = u.Scheme + "://" + u.Host
	} else {
		baseURL = url_string[0:strings.LastIndex(url_string, "/")]
	}
	
	idx := strings.LastIndex(url_string, "?")
	if idx == -1 {
		return baseURL + "/" + p
	} else {
		parameter := url_string[idx + 1:]
		if strings.LastIndex(p, "?") == -1 {
			return baseURL + "/" + p + "?" + parameter
		} else {
			return baseURL + "/" + p + "&" + parameter
		}
	}
	// baseURL/p (not include parameter)
	// baseURL/p?parameter (include parameter if p not contain ?)
	// baseURL/p&parameter (include parameter if p contain ?)
}

func DrawProgressBar(prefix string, proportion float32, width int, suffix ...string) {
	pos := int(proportion * float32(width))
	s := fmt.Sprintf("[%s] %s%*s %6.2f%% %s",
		prefix, strings.Repeat("■", pos), width-pos, "", proportion*100, strings.Join(suffix, ""))
	fmt.Print("\r" + s)
}
