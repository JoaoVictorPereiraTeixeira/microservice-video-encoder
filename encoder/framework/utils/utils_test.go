package utils_test

import (
	"testing"

	"github.com.br/JoaoVictorPereiraTeixeira/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJosn(t *testing.T) {
	json := `{
		"id":"1234562322",
		"file_path":"convite.mp4",
		"status":"pending"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `john`
	err = utils.IsJson(json)
	require.Error(t, err)
}
