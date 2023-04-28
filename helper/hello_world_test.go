package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("zul")
	// t.Fail() gagal lanjut eksekusi, t.FailNow() = gagal dn berhenti exekusi 
	if result != "Hello zul1" {
		//panic("result is not pass")
		t.Fail()
	}

	fmt.Println("test hello world")
} 

func TestHelloWorldMahar(t *testing.T) {
	result := HelloWorld("mahar")

	if result != "Hello mahar" {
		panic("result is not pass")
		t.Fail()
	}

	fmt.Println("hello world mahar")
} 


func TestHelloWordlAssert(t *testing.T) {
	result := HelloWorld("zul");

	assert.Equal(t, "Hello zul", result, "result must by Hello zul")

	//requiere.Equal(t, "Hello zul", result, "result must by Hello zul")vt.FailNow() github.com/stretchr/testify/require

	fmt.Println("test hello world assert")
}

func TestMain(m *testing.M) {
	fmt.Println("code before test")
	m.Run()
	fmt.Println("code after unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("zul", func(t *testing.T){
		result := HelloWorld("zul");

		require.Equal(t, "Hello zul", result, "Result Must by Hello zul")
	})

	t.Run("Mahar", func(t *testing.T){
		result := HelloWorld("mahar");

		require.Equal(t, "Hello mahar", result, "Result Must by Hello zul")
	})
}

func TestHelloWorldTable(t *testing.T) {
	test :=[] struct {
		name  string
		request string
		expected string
	} {
		{
			name: "zul",
			request: "zul",
			expected: "Hello zul",
		},
		{
			name: "mahar",
			request: "mahar",
			expected: "Hello mahar",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.name);
			require.Equal(t, test.expected, result, test.expected)
		})
	}
}