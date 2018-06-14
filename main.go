package main

import (
	"fmt"
	"github.com/mmaelzer/cam"
	"io/ioutil"
	"log"
)

const (
	path = "http://192.168.0.165:8081/?action=stream"
)

func main() {
	/*f, err := os.Open("/dev/input/event10")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	for {
		f.Read(b)
		sec := binary.LittleEndian.Uint64(b[0:8])
		usec := binary.LittleEndian.Uint64(b[8:16])
		t := time.Unix(int64(sec), int64(usec))
		fmt.Println(t)
		var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, code, value)
	}*/

	/*
		jar, _ := cookiejar.New(nil)
		client := &http.Client{
			Timeout: time.Second * 5,
			Jar:     jar,
		}
		req, err := http.NewRequest("GET", path, nil)
		catch(err)

		res, err := client.Do(req)
		catch(err)

		_, param, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
		catch(err)

		bounds := strings.Trim(param["boundary"], "-")
		log.Println(bounds)
		r := multipart.NewReader(res.Body, bounds)
		i := 0
		for {
			i++
			p, err := r.NextPart()
			if err == io.EOF || err != nil {
				log.Fatal(err)
			}


				spew, err := ioutil.ReadAll(p)
				catch(err)

				//fmt.Printf("%#x\n\n\n", spew)

				//make image from it
				ioutil.WriteFile(fmt.Sprintf("Image_%d.jpg", i), spew, os.ModePerm)


			//convert to go image
			bb, err := ioutil.ReadAll(p)
			img, err := jpeg.Decode(bytes.NewReader(bb))
			//img, _, err := image.Decode(p)
			catch(err)

			//create file first
			f, err := os.Create(fmt.Sprintf("image_%d.jpg", i))
			catch(err)

			jpeg.Encode(f, img, &jpeg.Options{Quality: 50})
			f.Close()
		}*/

	camera := cam.Camera{
		URL: path,
	}

	frames, err := camera.Subscribe()
	if err != nil {
		panic(err)
	}

	for frame := range frames {
		filename := fmt.Sprintf("%d.jpeg", frame.Timestamp.UnixNano())
		ioutil.WriteFile(filename, frame.Bytes, 0644)
	}
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
