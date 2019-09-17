package ping

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"plathome-backend/controller"
	"plathome-backend/models"
	"time"
)

var status = map[string]string{}

func pingAndWriteDB(ip string, db *controller.Database) {
	status[ip] = "connecting"
	p := fastping.NewPinger()

	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		log.Fatal(err)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		m := models.Device{}
		m.IP = addr.String()
		db.First(&m)
		m.State = "alive"
		status[ip] = "ok"
		db.Update(&m)
		log.Print(fmt.Sprintf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt))
	}
	p.OnIdle = func() {
		if status[ip] != "ok" {
			status[ip] = "timeout"
			m := models.Device{}
			m.IP = ip
			db.First(&m)
			m.State = "timeout"
			db.Update(&m)
			log.Println(ip + " Timeouted")
		}
		log.Println(ip + " finished")
	}
	err = p.Run()
	if err != nil {
		m := models.Device{}
		m.IP = ip
		db.First(&m)
		m.State = err.Error()
		db.Update(&m)
		log.Fatal(err)
	}
}

func pingAndWriteDBAll(db *controller.Database) {
	ds := db.FindAll()
	for _, d := range *ds {
		pingAndWriteDB(d.IP, db)
	}
}

func StartPingTask(db *controller.Database) {
	log.Println("waiting for 3 minutes")
	time.Sleep(1 * time.Minute)
	log.Println("waiting for 2 minutes")
	time.Sleep(1 * time.Minute)
	log.Println("waiting for 1 minutes")
	time.Sleep(1 * time.Minute)
	log.Println("ping manager started")
	for {
		pingAndWriteDBAll(db)
		time.Sleep(1 * time.Minute)
	}

}
