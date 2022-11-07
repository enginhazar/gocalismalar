// Example get-active-window reads the _NET_ACTIVE_WINDOW property of the root
// window and uses the result (a window id) to get the name of the window.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Tarih struct {
	time.Time
}
type sicil struct {
	Sicilno     int32
	TcKimlikNo  int64
	Adi         string
	Soyadi      string
	Adres       string
	DogumTarihi Tarih
}

func main() {

	rabbitMqConsume()

}

func rabbitMqConsume() {
	conn, err := amqp.Dial("amqp://admin:123456@localhost/")

	//amqp://guest:guest@localhost:5672/
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()

	// Kuyruk tanımlandı
	_, err = ch.QueueDeclare("Sicil", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	msgs, err := ch.Consume("Sicil", "", true, false, false, false, nil)

	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)
	go func() {
		layout := `"2006-01-02T15:04:05"`
		for d := range msgs {
			var gosicil sicil
			jsonveri := []byte(d.Body)
			err := json.Unmarshal(jsonveri, &gosicil)
			if err != nil {
				log.Fatalln(err)
				log.Fatalln("Fatal Error")

			}
			//fmt.Println(gosicil.sicilno)
			fmt.Printf("Sicilno      : %d\n", gosicil.Sicilno)
			fmt.Printf("TcKimlikNo   : %+v\n", gosicil.TcKimlikNo)
			fmt.Printf("Adı          : %s\n", gosicil.Adi)
			fmt.Printf("Soyadi       : %s\n", gosicil.Soyadi)
			fmt.Printf("Adres        : %s\n", gosicil.Adres)
			fmt.Printf("Dogum Tarihi : %s\n", gosicil.DogumTarihi.Format(layout))

			//	fmt.Println("---------------------\n")

		}
	}()

	log.Printf("Kuyruk Dinleniyor....")
	<-forever

}

// UnmarshalJSON / Unmarshall (decode) işlemi sırasında Tarih alanı formatı C# 'tan yazılan formata uygun şekilde parse edilmesi
// / için yazıldı
func (t *Tarih) UnmarshalJSON(b []byte) (err error) {
	layout := `"2006-01-02T15:04:05"`

	date, err := time.Parse(layout, string(b))
	if err != nil {
		log.Println(err)
		return err
	}
	t.Time = date
	return
}

//func main() {
//	pakets.NotifyTest()
//	//pakets.JsonMap()
//}

//var mt sync.Mutex
//
//func main() {
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//	var bakiye float64 = 100
//	fmt.Printf("İlk bakiye: %.2f\n", bakiye)
//
//	go paracek(&bakiye, &wg, 20)
//	go parayatir(&bakiye, &wg, 60)
//	wg.Wait()
//}
//
//func parayatir(bakiye *float64, wg *sync.WaitGroup, yatirilanTutar float64) {
//
//	mt.Lock()
//	*bakiye += yatirilanTutar
//	fmt.Printf("Bakiye : %.2f\n", *bakiye)
//	mt.Unlock()
//	fmt.Println("Para Yatırma İşlemi")
//	wg.Done()
//
//}
//
//func paracek(bakiye *float64, wg *sync.WaitGroup, cekilecekMiktar float64) {
//
//	mt.Lock()
//
//	*bakiye -= cekilecekMiktar
//	fmt.Printf("Bakiye : %.2f\n", *bakiye)
//
//	mt.Unlock()
//
//	fmt.Printf("Para Çekme İşlemi")
//	wg.Done()
//
//}

//
//// waiting group
//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//
//	go fonksiyon1(&wg)
//	go fonksiyon2(&wg)
//
//	fmt.Println("Merhaba Dünya!")
//
//	wg.Wait()
//
//	//waitgroup tamamlandığında ekrana yazı bastıralım.
//	fmt.Println("WaitGroup'lar tamamlandı.")
//}
//
//func fonksiyon2(wg *sync.WaitGroup) {
//	time.Sleep(time.Second + 2)
//	fmt.Println("2.fonksiyon")
//	wg.Done()
//}
//
//func fonksiyon1(wg *sync.WaitGroup) {
//	time.Sleep(time.Second + 2)
//	fmt.Println("1.fonksiyon")
//	wg.Done()
//}

//func main() {
//
//	k := make(chan string)
//	go func() {
//		fmt.Println("1")
//		time.Sleep(time.Second + 4)
//		fmt.Println("2")
//		k <- "engin"
//		fmt.Println(k)
//	}()
//	<-k
//	fmt.Println("disarıda")
//	fmt.Println(k)
//	//eszamanlilik()
//
//	//fmt.Println("ilk yazım")
//	//go yaz("eş zamanlı print")
//	//
//	//time.Sleep(time.Second + 2)
//
//}

//func yaz(s string) {
//	fmt.Println(s)
//}

//	func main() {
//		fmt.Println("engin")
//		fmt.Println("defne")
//
//		sicil1 := sicil{ad: "engin", soyad: "hazar", tckimlikno: 37396047616}
//
//		fmt.Println(sicil1)
//
//		for i := 1; i < 100; i++ {
//
//			fmt.Println(i)
//			switch {
//			case i%3 == 0 && i%5 == 0:
//				print("3 ve 5 bölündü")
//			case i%3 == 0:
//				fmt.Println("3 'e bölündü")
//			case i%5 == 0:
//				fmt.Println("5 'e bölündü")
//
//			}
//			//
//
//			//if i%3 == 0 && i%5 == 0 {
//			//	fmt.Println(i)
//			//	fmt.Println("Fizz Buzz")
//			//}
//			//if i%3 == 0 {
//			//	fmt.Println("fizz")
//			//}
//			//if i%5 == 0 {
//			//	fmt.Println("buzz")
//			//}
//		}
//	}
//type sicil struct {
//	ad         string
//	soyad      string
//	tckimlikno int
//}
