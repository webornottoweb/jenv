## jenv - package for making your json files configurable with .env files

---

### Usage

- create json file which will be your data source

```
{
    "ID": "123",
    "Name": "Some name",
    "Boolval": "Some bool"
}
```
> Notice how fields values are wrapped with quotes `"` - each field value for parsing with current package should be wrapped

- create a struct which will be a target for your json file parsing

```
import (
	"encoding/json"

	lj "github.com/webornottoweb/jenv/pkg/json"
)

type User struct {
	ID      lj.JEnvInt
	Name    lj.JEnvString
	Boolval lj.JEnvBool
}
```
> Notice, how local json package implementation was described with `lj`

- start parsing with Go `encoding/json` package

```
func main() {
	file, _ := ioutil.ReadFile("data.json")

	var user User

	err := json.Unmarshal([]byte(file), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```
> Result: `{123 Some name false}`

- if u want to use env value, just wrap json value with `env[]`

```
{
    "ID": "123",
    "Name": "env[USER_NAME]",
    "Boolval": "env[USER_BOOL]"
}
```
> Notice that you should have .env file with provided `USER_NAME` and `USER_BOOL` variables

> Result: `{123 name from env true}`

---

### Tasks

- implement raw array values parsing (strings/integers/booleans)
- tests cover