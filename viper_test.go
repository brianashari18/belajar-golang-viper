package golang_viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	config := viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "brian", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "brian", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, 3306, config.GetInt("database.port"))

}

func TestENVFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "brian", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "brian", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
