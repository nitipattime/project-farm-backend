package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/farms", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1) // ถ้าไม่มี page ให้ default เป็น 1
		_ = c.QueryInt("limit", 10)
		_ = c.Query("search", "")
		response := fiber.Map{}

		if page == 1 {
			response = fiber.Map{
				"data": fiber.Map{
					"list": []fiber.Map{
						{
							"id":           1,
							"farm_name":    "ฟาร์มแม่สอด",
							"house_count":  4,
							"province":     "ตาก",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณสมชาย ใจดี",
						},
						{
							"id":           2,
							"farm_name":    "ฟาร์มไก่ทอง",
							"house_count":  6,
							"province":     "นครราชสีมา",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณสายใจ บุญมาก",
						},
						{
							"id":           3,
							"farm_name":    "ไก่ดีฟาร์ม",
							"house_count":  3,
							"province":     "เชียงใหม่",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณมนตรี ทองสุข",
						},
						{
							"id":           4,
							"farm_name":    "เจริญฟาร์ม",
							"house_count":  5,
							"province":     "อุดรธานี",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณอารีย์ บัวแก้ว",
						},
						{
							"id":           5,
							"farm_name":    "สุวรรณฟาร์ม",
							"house_count":  2,
							"province":     "สุพรรณบุรี",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณณรงค์ อ่อนน้อม",
						},
						{
							"id":           6,
							"farm_name":    "ฟาร์มบ้านสวน",
							"house_count":  7,
							"province":     "นครปฐม",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณวิภา กลางใจ",
						},
						{
							"id":           7,
							"farm_name":    "ไทยฟาร์ม",
							"house_count":  8,
							"province":     "ขอนแก่น",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณเจษฎา ตั้งตรง",
						},
						{
							"id":           8,
							"farm_name":    "ไทยฟาร์ม",
							"house_count":  9,
							"province":     "ขอนแก่น",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณเจษฎา ตั้งตรง",
						},
						{
							"id":           9,
							"farm_name":    "ไทยฟาร์ม",
							"house_count":  10,
							"province":     "ขอนแก่น",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณเจษฎา ตั้งตรง",
						},
					},
					"page":       1,
					"limit":      9,
					"total":      15,
					"totalPages": 2,
					"summary": fiber.Map{
						"totalFarms":   2,
						"totalHouses":  2,
						"totalChicken": 25000,
					},
				},
			}
		} else if page == 2 {
			response = fiber.Map{
				"data": fiber.Map{
					"list": []fiber.Map{
						{
							"id":           10,
							"farm_name":    "ฟาร์มแม่สอด",
							"house_count":  10,
							"province":     "ตาก",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณสมชาย ใจดี",
						},
						{
							"id":           11,
							"farm_name":    "ฟาร์มไก่ทอง",
							"house_count":  11,
							"province":     "นครราชสีมา",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณสายใจ บุญมาก",
						},
						{
							"id":           12,
							"farm_name":    "ไก่ดีฟาร์ม",
							"house_count":  12,
							"province":     "เชียงใหม่",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณมนตรี ทองสุข",
						},
						{
							"id":           13,
							"farm_name":    "เจริญฟาร์ม",
							"house_count":  13,
							"province":     "อุดรธานี",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณอารีย์ บัวแก้ว",
						},
						{
							"id":           14,
							"farm_name":    "สุวรรณฟาร์ม",
							"house_count":  14,
							"province":     "สุพรรณบุรี",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณณรงค์ อ่อนน้อม",
						},
						{
							"id":           15,
							"farm_name":    "ฟาร์มบ้านสวน",
							"house_count":  15,
							"province":     "นครปฐม",
							"district":     "กรุงเทพ",
							"subdistrict":  "เขตพระนคร",
							"contact_name": "คุณวิภา กลางใจ",
						},
					},
					"page":       2,
					"limit":      9,
					"total":      15,
					"totalPages": 2,
					"summary": fiber.Map{
						"totalFarms":   2,
						"totalHouses":  2,
						"totalChicken": 25000,
					},
				},
			}
		}

		return c.JSON(response)
	})

	// ✅ POST /farms - mock create farm
	app.Post("/farms", func(c *fiber.Ctx) error {

		type FarmRequest struct {
			Search string `json:"search"`
			Page   int    `json:"page"`
			Limit  int    `json:"limit"`
		}

		var req FarmRequest

		// parse body
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// ไม่ตรวจ body ไม่ต้อง validate — mock อย่างเดียว
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Farm created successfully (mock)",
		})
	})

	// ✅ GET /districts
	app.Get("/districts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{"id": 1, "name_th": "เขตพระนคร"},
				{"id": 2, "name_th": "เขตดุสิต"},
			},
		})
	})

	// ✅ GET /subdistricts
	app.Get("/subdistricts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{"id": 1, "name_th": "เขตพระนคร", "zip_code": 18377},
				{"id": 2, "name_th": "เขตดุสิต", "zip_code": 11000},
			},
		})
	})

	// ✅ GET /address
	app.Get("/address", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{
					"subdistrict_id": 1,
					"subdistrict":    "เขตพระนคร",
					"district_id":    18377,
					"district":       "เขตพระนคร",
					"province_id":    1,
					"province":       "เขตพระนคร",
				},
				{
					"subdistrict_id": 2,
					"subdistrict":    "เขตพระนคร",
					"district_id":    18377,
					"district":       "เขตพระนคร",
					"province_id":    2,
					"province":       "เขตพระนคร",
				},
			},
		})
	})

	// ✅ GET /provinces
	app.Get("/provinces", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{"id": 1, "name_th": "กรุงเทพ"},
				{"id": 2, "name_th": "นนทบุรี"},
			},
		})
	})

	// ❗❗❗❗

	// ✅ GET /houses
	app.Get("/houses", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{
					"id":        1,
					"name":      "โรงเรือน A1",
					"manager":   "คุณสมปอง พันธุ์ดี",
					"avgWeight": 1800,
					"duration":  35,
					"status":    "กำลังเพาะเลี้ยง",
					"startDate": "01/06/25",
					"endDate":   "05/07/25",
				},
				{
					"id":        2,
					"name":      "โรงเรือน B2",
					"manager":   "คุณสายฝน ใจเย็น",
					"avgWeight": 1650,
					"duration":  28,
					"status":    "รอเพาะเลี้ยง",
					"startDate": "10/06/25",
					"endDate":   "08/07/25",
				},
				{
					"id":        3,
					"name":      "โรงเรือน C3",
					"manager":   "คุณมนัส แสงทอง",
					"avgWeight": 1725,
					"duration":  32,
					"status":    "เพาะเลี้ยงเสร็จสิ้น",
					"startDate": "05/05/25",
					"endDate":   "06/06/25",
				},
				{
					"id":        4,
					"name":      "โรงเรือน D4",
					"manager":   "คุณจินตนา เมฆขาว",
					"avgWeight": 1600,
					"duration":  30,
					"status":    "กำลังเพาะเลี้ยง",
					"startDate": "15/06/25",
					"endDate":   "15/07/25",
				},
			},
		})
	})

	// ✅ GET /house-details
	app.Get("/house-details", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{
					"id":        1,
					"manager":   "คุณสมชาย ใจดี",
					"avgWeight": 1850,
					"duration":  30,
					"status":    "กำลังเพาะเลี้ยง",
				},
				{
					"id":        2,
					"manager":   "คุณสายฝน บุญเลิศ",
					"avgWeight": 1780,
					"duration":  28,
					"status":    "รอเพาะเลี้ยง",
				},
			},
		})
	})

	// ✅ POST /houses
	app.Post("/houses", func(c *fiber.Ctx) error {
		// ไม่ตรวจสอบข้อมูลที่ส่งมา — mock response
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "สร้างโรงเรือนสำเร็จ (mock)",
			"id":      999, // mock id กลับ
		})
	})

	// ✅ POST /house-details
	app.Post("/house-details", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "สร้างรายละเอียดโรงเรือนสำเร็จ (mock)",
			"id":      888,
		})
	})

	app.Listen(":3000")
}
