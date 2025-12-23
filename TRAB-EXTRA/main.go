package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"golang.org/x/crypto/sha3"
)

func DesafioA() {
	fmt.Println("--- Desafio A: Quebrando Resistência a Colisão ---")

	seen := make(map[string]string)
	var mu sync.Mutex
	var i int64 = 0
	found := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for w := 0; w < runtime.NumCPU(); w++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					currentI := atomic.AddInt64(&i, 1)
					input := fmt.Sprintf("seed_%d", currentI)

					d := sha3.NewShake128()
					d.Write([]byte(input))
					hashBytes := make([]byte, 4)
					d.Read(hashBytes)
					hash := hex.EncodeToString(hashBytes)

					mu.Lock()
					if val, exists := seen[hash]; exists {
						fmt.Printf("Colisão encontrada!\nInput 1: %s\nInput 2: %s\nHash: %s\n\n", val, input, hash)
						cancel()
						mu.Unlock()
						found <- true
						return
					}
					seen[hash] = input
					mu.Unlock()
				}
			}
		}()
	}
	<-found
}

func DesafioB(nomeCompleto string) {
	fmt.Println("--- Desafio B: Quebrando Segunda Pré-imagem ---")

	x1 := "Aluno: " + nomeCompleto
	d1 := sha3.NewShake128()
	d1.Write([]byte(x1))
	targetHash := make([]byte, 4)
	d1.Read(targetHash)
	targetHex := hex.EncodeToString(targetHash)

	fmt.Printf("Alvo (x1): %s | Hash Alvo: %s\n", x1, targetHex)

	var i int64 = 0
	found := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for w := 0; w < runtime.NumCPU(); w++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					currentI := atomic.AddInt64(&i, 1)
					x2 := fmt.Sprintf("nonce_%d", currentI)

					if x2 == x1 {
						continue
					}

					d2 := sha3.NewShake128()
					d2.Write([]byte(x2))
					h2 := make([]byte, 4)
					d2.Read(h2)

					if hex.EncodeToString(h2) == targetHex {
						fmt.Printf("Segunda pré-imagem encontrada!\nInput (x2): %s\nHash: %s\n\n", x2, targetHex)
						cancel()
						found <- true
						return
					}
				}
			}
		}()
	}
	<-found
}

// getSHAKE128 gera o hash SHAKE128 com o tamanho solicitado [cite: 5, 13]
func getSHAKE128(input string, outputSize int) []byte {
	hash := make([]byte, outputSize)
	d := sha3.NewShake128()
	d.Write([]byte(input))
	d.Read(hash)
	return hash
}

func desafioC(targetHex string) {
	fmt.Printf("--- Desafio C: Quebrando Pré-imagem para %s (34 bits) ---\n", targetHex)

	targetBytes, _ := hex.DecodeString(targetHex + "0")
	targetValue := uint64(targetBytes[0])<<32 | uint64(targetBytes[1])<<24 | uint64(targetBytes[2])<<16 | uint64(targetBytes[3])<<8 | uint64(targetBytes[4])

	mask := uint64(0xFFFFFFFFF0)
	targetValue &= mask

	numCPU := runtime.NumCPU()
	fmt.Printf("Utilizando %d núcleos do processador...\n", numCPU)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var foundCounter int32
	var totalAttempts uint64

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := uint64(workerID); ; j += uint64(numCPU) {
				select {
				case <-ctx.Done():
					return
				default:
					input := fmt.Sprintf("pass_%d", j)
					hash := getSHAKE128(input, 5)

					currentVal := uint64(hash[0])<<32 | uint64(hash[1])<<24 | uint64(hash[2])<<16 | uint64(hash[3])<<8 | uint64(hash[4])

					if (currentVal & mask) == targetValue {
						if atomic.CompareAndSwapInt32(&foundCounter, 0, 1) {
							fmt.Printf("\n[Worker %d] SUCESSO!\n", workerID)
							fmt.Printf("Input encontrado: %s\n", input)
							fmt.Printf("Hash (Hex): %s\n", strings.ToUpper(hex.EncodeToString(hash)))
							cancel() // Para todos os outros workers
						}
						return
					}

					if j%1000000 == 0 {
						atomic.AddUint64(&totalAttempts, 1000000)
					}
				}
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	DesafioA()
	DesafioB("Aluno: Luiz Guilherme Moreira Leite")
	desafioC("17675FC0")
}
