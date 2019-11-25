package conf

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestConfig_Int64(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	v := c.Int64("config.int64")
	expected := int64(9223372036854775807)
	if v != expected {
		t.Fatalf("'%v' != '%v'", v, expected)
	}
}

// Test not working, see https://github.com/knadh/koanf/issues/6
//func TestConfig_Int64s(t *testing.T) {
//	c,_ := FromYaml("testdata/config.yaml")
//	vs := c.Int64s("config.int64s")
//	fmt.Printf("%#v", vs)
//	expected := []int64{9223372036854775805,9223372036854775806,9223372036854775807}
//	for i, e := range expected {
//		v := vs[i]
//		if v != e {
//			t.Fatalf("'%v' != '%v'", v, expected)
//		}
//	}
//}

func TestConfig_Int64s2(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Int64s("config.int64s2")
	expected := []int64{1, 2, 3}
	for i, e := range expected {
		v := vs[i]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

// Test not working, see https://github.com/knadh/koanf/issues/6
//func TestConfig_Int64map(t *testing.T) {
//	c,_ := FromYaml("testdata/config.yaml")
//	vs := c.Int64Map("config.int64map")
//	fmt.Printf("%#v", vs)
//	expected := map[string]int64{"a": 9223372036854775805, "b": 9223372036854775806, "c": 9223372036854775807}
//	for k, e := range expected {
//		v := vs[k]
//		if v != e {
//			t.Fatalf("'%v' != '%v'", v, expected)
//		}
//	}
//}

func TestConfig_Int64map2(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Int64Map("config.int64map2")
	expected := map[string]int64{"a": 1, "b": 2, "c": 3}
	for k, e := range expected {
		v := vs[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_Int(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	v := c.Int("config.int")
	expected := 1
	if v != expected {
		t.Fatalf("'%v' != '%v'", v, expected)
	}
}

func TestConfig_Ints(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Ints("config.ints")
	expected := []int{1, 2, 3}
	for i, e := range expected {
		v := vs[i]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_Intmap(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.IntMap("config.intmap")
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, e := range expected {
		v := vs[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_Float64(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	v := c.Float64("config.float64")
	expected := 1.0
	if v != expected {
		t.Fatalf("'%v' != '%v'", v, expected)
	}
}

func TestConfig_Float64s(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Float64s("config.float64s")
	expected := []float64{1.0, 1.1, 1.2}
	for i, e := range expected {
		v := vs[i]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_Float64Map(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Float64Map("config.float64map")
	expected := map[string]float64{"a": 1.0, "b": 1.1, "c": 1.2}
	for k, e := range expected {
		v := vs[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_String(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	v := c.String("config.string")
	expected := "foo"
	if v != expected {
		t.Fatalf("'%v' != '%v'", v, expected)
	}
}

func TestConfig_Strings(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Strings("config.strings")
	expected := []string{"foo", "bar"}
	for i, e := range expected {
		v := vs[i]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_StringMap(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.StringMap("config.stringmap")
	expected := map[string]string{"foo": "foo", "bar": "bar"}
	for k, e := range expected {
		v := vs[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_Bool(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	v := c.Bool("config.true")
	if v != true {
		t.Fatalf("'%v' != '%v'", v, true)
	}
	v = c.Bool("config.false")
	if v != false {
		t.Fatalf("'%v' != '%v'", v, false)
	}
}

func TestConfig_Bools(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.Bools("config.bools")
	expected := []bool{true, false}
	for i, e := range expected {
		v := vs[i]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_BoolMap(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	vs := c.BoolMap("config.boolmap")
	expected := map[string]bool{"true": true, "false": false}
	for k, e := range expected {
		v := vs[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestNewConfig(t *testing.T) {
	c := NewConfig("_")
	if c.Delimiter() != "_" {
		t.Fatalf("'%v' != '%v'", c.Delimiter(), "_")
	}
}

func TestConfig_Exists(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	if !c.Exists("config.int") {
		t.Fatalf("config.int key not existing")
	}
	if c.Exists("config.nonexisting") {
		t.Fatalf("config.nonexisting key existing")
	}
}

func TestConfig_Get(t *testing.T) {
	c, _ := FromYaml("testdata/config.yaml")
	if c.Get("config.int") != 1 {
		t.Fatalf("'%v' != '%v'", c.Get("config.int"), 1)
	}
	if c.Get("config.nonexisting") != nil {
		t.Fatalf("'%v' != '%v'", c.Get("config.nonexisting"), nil)
	}
}

func TestConfig_Keys(t *testing.T) {
	c, _ := FromJson("testdata/keys.json")
	vs := c.Keys()
	expected := []string{"1.1.1.1.1.1", "1.1.2", "2"}
	for k, e := range expected {
		v := vs[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_KeyMap(t *testing.T) {
	c, _ := FromJson("testdata/keys.json")
	keys := c.KeyMap()
	expected := map[string][]string{"1": {"1"}, "1.1.1": {"1", "1.1"}, "1.1.1.1.1.1": {"1", "1.1", "1.1.1"}, "1.1.2": {"1", "1.2"}, "2": {"2"}}
	for ek1, ev1 := range expected {
		v1, ok := keys[ek1]
		if !ok {
			t.Fatalf("Non existing expected key '%v' in '%#v'", ek1, keys)
		}
		for ek2, ev2 := range ev1 {
			v2 := v1[ek2]
			if v2 != ev2 {
				t.Fatalf("'%v' != '%v'", v2, ev2)
			}
		}
	}
}

func TestConfig_Unmarshal(t *testing.T) {

	type Name struct {
		First string
		Last  string
	}

	type Person struct {
		Name  Name
		Title string
		Age   int
		Ids   []int
	}

	var expected = Person{
		Name: Name{
			First: "Andreas",
			Last:  "Kogelbauer",
		},
		Title: "Scatman",
		Age:   34,
		Ids:   []int{1, 2, 3},
	}

	var p Person
	c, _ := FromJson("testdata/unmarshal.json")
	_ = c.Unmarshal("", &p)

	if p.Name != expected.Name {
		t.Fatalf("'%#v' != '%#v'", p.Name, expected.Name)
	}

	if p.Title != expected.Title {
		t.Fatalf("'%#v' != '%#v'", p.Name, expected.Name)
	}

	if p.Age != expected.Age {
		t.Fatalf("'%#v' != '%#v'", p.Name, expected.Name)
	}

	for k, e := range expected.Ids {
		v := p.Ids[k]
		if v != e {
			t.Fatalf("'%v' != '%v'", v, expected)
		}
	}
}

func TestConfig_Watch(t *testing.T) {

	type C struct {
		Config int
	}

	var updateFile = func(config int) {
		c := C{
			Config: config,
		}
		data, _ := json.Marshal(c)
		_ = ioutil.WriteFile("testdata/watch.json", data, 0644)

	}

	ch := make(chan int)
	updateFile(1)
	c, _ := FromJson("testdata/watch.json")
	err := c.WatchFile("testdata/watch.json", func() {
		config := c.Int("Config")
		ch <- config
	})
	if err != nil {
		t.Fatal(err)
	}
	updateFile(2)
	config := <-ch
	if config != 2 {
		t.Fatal("onReaload never called")
	}
}
