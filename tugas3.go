package main

import "fmt"

func main()  {
  var buah = []string{"Apel","Jeruk","Anggur","Melon"}
  buah =append(buah,"Pepaya")
  fmt.Println("Jumlah Element :", len(buah))
  fmt.Println("Isi Element : ", buah)

for i:=0; i<5; i++{
  fmt.Println("Element Ke - :", i,(buah[i]))
}

  // var i = 0
  // for {
  //   fmt.Println("Element Ke - : ", i, (buah[i]))
  //   i++
  //   if i == len(buah) {
  //     break
  //   }
  // }
}
