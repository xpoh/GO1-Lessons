package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

type Config struct {
	Port       string
	Name       string
	Url        string
	jsonConfig string
}

func readJsonFromFile(fileName string) (Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	c := Config{}
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	fmt.Println(c)
	return c, err
}

func ReadConfig() (Config, error) {
	var err = error(nil)
	var c Config

	var port = flag.String("port", "80", "port number")
	var name = flag.String("name", "user", "username")
	var urr = flag.String("url", "http://localhost", "dest address")
	var jsonConfig = flag.String("file", "", "json config file")

	flag.Parse()

	if *jsonConfig != "" {
		c, err = readJsonFromFile(*jsonConfig)
		if err != nil {
			panic(err)
		}
	} else {
		u, err := url.ParseRequestURI(*urr)
		if err != nil {
			panic(err)
		}
		c = Config{
			Port: *port,
			Name: *name,
			Url:  u.String(),
		}
		fmt.Println(json.Marshal(c))
	}

	return c, err
}
