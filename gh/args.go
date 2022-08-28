package gh

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

var graphqlArgs = []string{"api", "graphql", "-f"}

func CreateLabelApiArgs(owner string, repo string, labelName string) []string {
	rand.Seed(time.Now().UnixNano())
	rgb := []float64{rand.Float64() * 255, rand.Float64() * 255, rand.Float64() * 255}
	colorCode := colorful.LinearRgb(rgb[0], rgb[1], rgb[2])
	baseArgs := []string{
		"api",
		"--method",
		"POST",
		fmt.Sprintf("/repos/%s/%s/labels", owner, repo),
		"-f",
		fmt.Sprintf(`name=%s`, labelName),
		"-f",
		fmt.Sprintf(`color=%s`, colorCode.Hex()[1:]),
	}

	return baseArgs
}

func GetLabelApiArgs(owner string, repo string, labelName string) []string {
	baseArgs := []string{
		"api",
		"--method",
		"GET",
		fmt.Sprintf("/repos/%s/%s/labels/%s", owner, repo, labelName),
	}

	return baseArgs
}
