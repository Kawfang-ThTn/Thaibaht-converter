# 1. Base Image
FROM golang:1.22-alpine

# 2. ตั้งค่า Working directory ใน Container
WORKDIR /app

# 3. คัดลอกและติดตั้ง Dependencies
COPY go.mod go.sum ./
RUN go mod download

# 4. คัดลอกโค้ดทั้งหมด
COPY . .

# 5. Build โปรแกรม (สั่ง build จากไฟล์ main.go ที่อยู่ด้านนอก)
RUN go build -o main main.go

# 6. รันคำสั่งเมื่อ Container เริ่มทำงาน
CMD ["./main"]