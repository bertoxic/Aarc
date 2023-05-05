package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
)

func GenerateVerificationCode() (string, time.Time) {
	b := make([]byte, 8)
	n, err := rand.Read(b)
	log.Println(n)
	if err != nil {
		log.Println(err)
	}
	expiry := time.Now().Add(30 * time.Minute)
//	log.Println(base64.URLEncoding.EncodeToString(b))

	return base64.URLEncoding.EncodeToString(b), expiry
}

type PcProperties struct {
 ScreenSize int
 Cpu int
 StorageSize int
 RamSize int
 Gen int
 Resolution int
 Price int
}

func NewPC ()*PcProperties{
	return &PcProperties{
	
	}
}

func (m *PcProperties)EstimatePrice() int {
	b:= 0
	sum := func (a ...int) (n int) {
		
		for i :=0; i<len(a); i ++  {
			b = b +	a[i]
			log.Println(b)
		}
		log.Println(b)
		return b
	}
	sum (m.screenSizeP(), m.cpuP(),m.RamSizeP(),m.GenP(),m.ResolutionP(), m.StorageSizeP())
	a:=&sum
	log.Println(*a)
	return b

}

func (p *PcProperties)screenSizeP()(price int){
		switch p.ScreenSize {
		case 15: return 45000
		case 14: return 35000
		case 13: return 30000
		case 12: return 20000

		default: return 20
		}
}
func (p *PcProperties)StorageSizeP()(price int){
		switch p.StorageSize {
		case 1024: return 70000
		case 512: return 50000
		case 256: return 30000
		case 128: return 20000

		default: return 100
		}
}


func (p *PcProperties)cpuP()(price int){
		switch p.Cpu {
		case 3: return 20000
		case 5: return 36000
		case 7: return 52000
		case 9: return 97000

		default: return 5
		}
}


func (p *PcProperties)RamSizeP()(price int){
		switch p.RamSize {
		case 4: return 10000
		case 8: return 16000
		case 12: return 32000
		case 16: return 44000
		case 32: return 60000

		default: return 10
		}
}


func (p *PcProperties)GenP()(price int){
		switch p.Gen {
		case 4: return 10000
		case 5: return 20000
		case 6: return 30000
		case 7: return 35000
		case 8: return 40000
		case 9: return 60000
		case 10: return 65000
		case 11: return 70000
		case 12: return 90000

		default: return 10
		}
}


func (p *PcProperties)ResolutionP()(price int){
		switch p.Resolution {
		case 720: return 20000
		case 14: return 20000
		case 1080: return 30000
		case 4000: return 50000

		default: return 7
		}
}


