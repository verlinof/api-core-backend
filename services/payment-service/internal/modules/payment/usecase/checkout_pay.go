package usecase

import (
	"log"
	pkg_midtrans "payment-service/pkg/helper/midtrans"

	"github.com/gofiber/fiber/v2"
)

func PaymentNotificationHandler(ctx *fiber.Ctx) error {
	// 1. Parse JSON body dari Midtrans
	var notificationPayload map[string]interface{}
	if err := ctx.BodyParser(&notificationPayload); err != nil {
		log.Printf("[Webhook] Error parsing body: %v\n", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	// 2. Dapatkan OrderID dari payload
	orderID, exists := notificationPayload["order_id"].(string)
	if !exists {
		log.Println("[Webhook] Invalid payload: order_id not found")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// 3. ✨ PRAKTIK TERBAIK: Verifikasi status transaksi ke Midtrans
	// Ini JAUH LEBIH AMAN daripada hanya memvalidasi signature_key.
	// Kita "bertanya balik" ke Midtrans: "Hei, order_id ini statusnya beneran apa?"
	transactionStatusResp, err := pkg_midtrans.CoreApiClient.CheckTransaction(orderID)
	if err != nil {
		log.Printf("[Webhook] Error checking transaction status: %v\n", err)
		// Balas 500 agar Midtrans coba kirim ulang nanti
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 4. Proses status transaksi
	// Di sini kamu akan melakukan UPDATE di database-mu
	status := transactionStatusResp.TransactionStatus
	fraudStatus := transactionStatusResp.FraudStatus

	log.Printf("[Webhook] Received notification for OrderID %s. Status: %s, FraudStatus: %s\n",
		orderID, status, fraudStatus)

	if status == "capture" {
		if fraudStatus == "accept" {
			// Pembayaran berhasil DAN aman
			log.Printf("Payment for OrderID %s is successful and accepted.\n", orderID)
			// TODO: Update status order di DB kamu menjadi "PAID" atau "SUCCESS"
			// TODO: Kirim email konfirmasi ke user
			// TODO: Jalankan logika bisnis lain (misal: kurangi stok)
		} else if fraudStatus == "challenge" {
			// Pembayaran berhasil, tapi ditandai "challenge" oleh Midtrans
			log.Printf("Payment for OrderID %s is challenged by FDS.\n", orderID)
			// TODO: Update status order di DB kamu menjadi "CHALLENGE"
			// Kamu bisa memproses manual atau biarkan
		}
	} else if status == "settlement" {
		// Pembayaran berhasil (untuk tipe pembayaran tertentu seperti VA)
		log.Printf("Payment for OrderID %s is settled.\n", orderID)
		// TODO: Update status order di DB kamu menjadi "PAID" atau "SUCCESS"
		// TODO: Kirim email konfirmasi ke user
		// TODO: Jalankan logika bisnis lain
	} else if status == "pending" {
		// User belum bayar
		log.Printf("Payment for OrderID %s is still pending.\n", orderID)
		// TODO: Update status order di DB kamu menjadi "PENDING"
	} else if status == "deny" || status == "expire" || status == "cancel" {
		// Pembayaran gagal, kadaluarsa, atau dibatalkan
		log.Printf("Payment for OrderID %s is denied/expired/canceled.\n", orderID)
		// TODO: Update status order di DB kamu menjadi "FAILED" atau "CANCELED"
	}

	// 5. Kirim balasan 200 OK
	// WAJIB, agar Midtrans tahu notifikasinya sudah diterima & berhenti mengirim ulang.
	return ctx.SendStatus(fiber.StatusOK)
}
