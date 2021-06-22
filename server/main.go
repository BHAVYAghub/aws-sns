package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/endpoint",handler1)
	http.ListenAndServe(":8888",nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {

	m:=map[string]string{}
	b,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println("^^^",b)
	err=json.Unmarshal(b,&m)
	if err!=nil{
		log.Println(err)
		return
	}



	m2 :=map[string]interface{}{}
	err=json.Unmarshal([]byte(m["Message"]),&m2)
	log.Println("@@m2",m2)
	for k,v:=range m{
		log.Println("qwerty$$$$",k," ",v)
	}
	log.Println("qwerty%%%%%%%%%%%%%%%",m)

}

