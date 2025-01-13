package main

import (
	"io"
	"log"
	"net"
)

// echo — это функция-обработчик, просто отражающая полученные данные
// func echo(conn net.Conn) {
// 	defer conn.Close()
// 	// Создаем буфер для хранения полученных данных
// 	b := make([]byte, 512)
// 	for {
// 		// Получаем данные через conn.Read в буфер
// 		size, err := conn.Read(b[0:])
// 		if err == io.EOF {
// 			log.Println("Client disconnected")
// 			break
// 		}
// 		if err != nil {
// 			log.Println("Unexpected error")
// 			break
// 		}
// 		log.Printf("Received %d bytes: %s\n", size, string(b))
// 		// Отправляем данные через conn.Write
// 		log.Println("Writing data")
// 		if _, err := conn.Write(b[0:size]); err != nil {
// 			log.Fatalln("Unable to write data")
// 		}
// 	}
// }

// func echo(conn net.Conn) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)
// 	s, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatalln("Unable to read data")
// 	}
// 	log.Printf("Read %d bytes: %s", len(s), s)

// 	log.Println("Writing data")
// 	writer := bufio.NewWriter(conn)

// 	if _, err := writer.WriteString(s); err != nil {
// 		log.Fatalln("Unable to write data")
// 	}
// 	writer.Flush()
// }

func echo(conn net.Conn) {
	defer conn.Close()
	// Копируем данные из io.Reader в io.Writer через io.Copy()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	// Привязываемся к TCP-порту 20080 во всех интерфейсах
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Ожидаем соединения и при его установке создаем net.Conn
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Обрабатываем соединение, используя горутины для многопоточности
		go echo(conn)
	}
}
