package init

import (
	"github.com/joho/godotenv"
	lg "quote/pkg/utils"
)

func Init() {

	err := godotenv.Load("../../.env")

	if err != nil {
		lg.Errl.Fatal("Error loading .env file")
	}
}
