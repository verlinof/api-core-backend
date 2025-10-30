package pkg_midtrans

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var (
	SnapClient    snap.Client
	CoreApiClient coreapi.Client
)

func Init() {
	// 1. Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 2. Baca environment dari .env
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	envStr := os.Getenv("MIDTRANS_ENV")

	var env midtrans.EnvironmentType
	if envStr == "production" {
		env = midtrans.Production
	} else {
		env = midtrans.Sandbox
	}

	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY is not set in .env file")
	}

	// 3. Inisialisasi Snap Client
	SnapClient.New(serverKey, env)

	// 4. Inisialisasi CoreAPI Client
	CoreApiClient.New(serverKey, env)

	log.Println("Midtrans Clients Initialized. Environment:", serverKey)
}
