package main

import (
   "net"
   "fmt"
   "strconv"
   "time"
)

func main() {

   for i := 22010; i < 22011; i++ {
      port := strconv.FormatInt(int64(i), 10)
      conn, err := net.DialTimeout("tcp", "target1.bowneconsulting.com:" + port, 1*time.Second)
      fmt.Println("Scanning Port",i)
      if err == nil {
         fmt.Println("Port",i, "open")
       	 Buffer := make([]byte, 1024) // String
		 numBytesRead, err := conn.Read(Buffer)
		 if err == nil {
            fmt.Printf("Banner: %s\n", Buffer[0:numBytesRead])
            fmt.Printf("Banner: %s\n", Buffer[40:46])
            newPort := string(Buffer[41:46])
	    conn2, err := net.DialTimeout("tcp", "target1.bowneconsulting.com:" + newPort, 1*time.Second)
	    if err == nil{
	    Buffer2 := make([]byte, 1024) // String
	    numBytesRead2, err := conn2.Read(Buffer2)
		if err == nil {
		    fmt.Printf("Banner: %s\n", Buffer2[0:numBytesRead2])
			         conn.Close()

			}
		 } else { fmt.Println("Error: ",err) }
		} else { fmt.Println("Error: ",err)}
         conn.Close()

      } else {

        fmt.Println("Error: ",err)
      }
    }
}
