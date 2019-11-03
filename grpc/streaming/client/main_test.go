package main

import (
	"log"
	"sync"
	"testing"
)

func BenchmarkPushStream100x100(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		if _, err = pushStream(client, 100, 100); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkUnaryPush100x100(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		if _, err = push(client, 100, 100); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkPushStream10000x1(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		if _, err = pushStream(client, 10000, 1); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkUnaryPush10000x1(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		if _, err = push(client, 10000, 1); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkPushStream1000x1000(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		if _, err = pushStream(client, 1000, 1000); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkUnaryPush1000x1000(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		if _, err = push(client, 1000, 1000); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkPushStream1000x1000Parallel2(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 2
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = pushStream(client, 1000/nworkers, 1000); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkUnaryPush1000x1000Parallel2(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 2
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = push(client, 1000/nworkers, 1000); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkPushStream1000x1000Parallel10(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 10
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = pushStream(client, 1000/nworkers, 1000); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkUnaryPush1000x1000Parallel10(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 10
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = push(client, 1000/nworkers, 1000); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkPushStream10000x1Parallel10(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 10
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = pushStream(client, 10000/nworkers, 1); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkUnaryPush10000x1Parallel10(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 2
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = push(client, 10000/nworkers, 1); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkPushStream100x100Parallel10(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 10
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = pushStream(client, 100/nworkers, 100); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkUnaryPush100x100Parallel10(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	nworkers := 2
	for i := 1; i <= b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(nworkers)
		for j := 1; j <= nworkers; j++ {
			go func() {
				if _, err = push(client, 100/nworkers, 100); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
