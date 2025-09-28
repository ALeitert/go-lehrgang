package foundations

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

func goroutines() {
	// We can run functions concurrently with the `go` command.

	start := time.Now()
	print := func(id int) {
		for i := 0; time.Since(start) < 100*time.Millisecond; i++ {
			fmt.Println("goroutine:", id, "  iter:", i)
		}
	}

	for id := range 3 {
		// Starts a new goroutine with its own stack.
		go print(id)
	}

	time.Sleep(time.Second)
}

func channels() {
	// Channels allow to communicate between goroutines.

	// They send/receive objects of a specified type.
	var numberChannel chan int

	// They can also send empty tokens:
	var tokenChannel chan struct{}

	// Channel-variables are pointers. Their zero-value is `nil`.
	fmt.Println(numberChannel, tokenChannel) // nil nil

	// We can create them with `make`.
	numberChannel = make(chan int, 5)  // They have a fixed capacity ...
	tokenChannel = make(chan struct{}) // ... which can be 0.

	// ---

	// Read and write with `<-`.
	numberChannel <- 123
	fmt.Println(<-numberChannel) // 123

	go func() {
		// If a channel is full, a goroutine blocks until it can write.
		// For 0-capacity channels: until there is a reader.
		tokenChannel <- struct{}{}
		fmt.Println("done writing token")
	}()

	go func() {
		time.Sleep(time.Second)
		<-tokenChannel

		// If a channel is empty, reading will block until a value is written.
		tokenChannel <- struct{}{}
		fmt.Println("done reading token")
	}()

	time.Sleep(time.Second)
	<-tokenChannel

	// ---

	// We can close channels.
	// Once closed, they remain closed.
	close(tokenChannel)

	// Reding from a closed chanel is always successful.
	// All waiting goroutines will unblock.
	<-tokenChannel

	// When reading, we can check if a channel is closed with the ok-pattern.
	_, ok := <-tokenChannel
	if !ok {
		fmt.Println("channel is closed")
	}

	// ---

	for i := range cap(numberChannel) {
		numberChannel <- i
	}
	close(numberChannel)

	// We can continue reading remaining objects from buffered channels even if
	// they are closed.
	for {
		number, ok := <-numberChannel
		fmt.Println(number, ok)
		// 0 true
		// 1 true
		// 2 true
		// 3 true
		// 4 true
		// 0 false

		if !ok {
			break
		}
	}

	// ---

	// We can iterate over channels with a for-loop.
	numberChannel = make(chan int, 5)
	for i := range 5 {
		numberChannel <- i + 1
	}

	for num := range numberChannel {
		fmt.Println(num)
	}
}

func selectStatement() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for i := 0; ; i++ {
			sleepTime := rand.IntN(400)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			chan1 <- i
		}
	}()

	go func() {
		for i := 0; ; i++ {
			sleepTime := rand.IntN(100)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			chan2 <- i
		}
	}()

	timer := time.NewTimer(5 * time.Second)

	// The select statement allows to simultaneously wait for multiple chanel
	// operations.
	for done := false; !done; {
		select {
		case v := <-chan1:
			fmt.Println("chan1:", v)

		case v := <-chan2:
			fmt.Println("chan2:", v)

		case <-timer.C:
			fmt.Println("done")
			done = true
		}
	}
}

func contexts() {
	// Contexts form a tree.
	// If we cancel one context, all its children are cancelled too.

	// Root contexts: both are empty contexts.
	rootCtx := context.Background()
	_ = context.TODO()

	// We can construct children with various features:

	// Calling `cancel` indicates that we want to stop.
	cancelCtx, cancel := context.WithCancel(rootCtx)
	defer cancel() // <- Always make sure to can when no longer needed.

	// If you want to be more descriptive.
	_, cancelWithCause := context.WithCancelCause(rootCtx)
	cancelWithCause(errors.New("reason for cancel"))

	// Stop at a particular time. (Same for .WithDeadlineCause.)
	_, deadlineCancel := context.WithDeadline(rootCtx, time.Date(2025, time.December, 24, 0, 0, 0, 0, time.UTC))
	defer deadlineCancel()

	// Stop after a certain amount of time passed.
	// Great, for example, for API or similar calls over network.
	timeoutCtx, timeoutCancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer timeoutCancel()

	//
	// ---
	//

	// Start a worker which we cancel later.
	// Context should be first argument.
	go func(ctx context.Context) {
		defer fmt.Println("Background worker (1) stopped")
		for i, done := 0, false; !done; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			fmt.Println("Background worker (1) at step", i)
		}
	}(cancelCtx)

	time.Sleep(time.Second)
	cancel()

	// Start a worker which we cancel later.
	// Context should be first argument.
	go func(ctx context.Context) {
		defer fmt.Println("Background worker (2) stopped")
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			fmt.Println("Background worker (2) at step", i)
		}
	}(timeoutCtx)
	<-timeoutCtx.Done()

	// Indicates why an error was cancelled.
	fmt.Println(timeoutCtx.Err())

	//
	// ---
	//

	// We can store key-value pairs in contexts. Can be any type.
	kvCtx := context.WithValue(rootCtx, "key", "value")
	fmt.Println("key:", kvCtx.Value("key"))
}

func waitGroup() {
	// Problem: we have multiple goroutines and want to wait for all to finish.

	// A wait group has an internal counter.
	wg := sync.WaitGroup{}

	wg.Add(1) // Increases the number of events to wait for.
	go func() {
		defer wg.Done() // Indicates that one task is done (decreases the counter by 1).

		for i := range 20 {
			sleepTime := rand.IntN(400)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			fmt.Println("1:", i)
		}
	}()

	// New in 1.25: wg.Go()
	wg.Go(func() {
		for i := range 20 {
			sleepTime := rand.IntN(100)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			fmt.Println("2:", i)
		}
	})

	wg.Wait() // Blocks until internal counter reaches 0.
	fmt.Println("done")
}

func syncPackage() {
	// Useful things in sync:

	// Ensure that an operation was performed exactly once.
	_ = sync.Once{}

	// A threadsave hash table (without generics)
	_ = sync.Map{}

	// Mutex (i.e. a lock) and Read/Write Mutex.
	_ = sync.Mutex{}
	_ = sync.RWMutex{}

	// Keeps a group of objects around. Provides them if requested or generates
	// new instances. Unfortunately no limit.
	_ = sync.Pool{}

	// Atomics
	_ = atomic.Int64{} // Bool, [U]Int(32|64), Pointer, uintptr, Value(*)
}
