
# Concurrent Book Reservation System

This document explains the concurrency approach used in the Library Management System.

## Goroutines

We use Goroutines to process multiple reservation requests simultaneously. When a member requests to reserve a book, a new Goroutine is created to handle the reservation process. This allows the system to handle multiple requests without blocking.

## Channels

We use a channel to queue incoming reservation requests. The `ReserveBook` method sends a `Reservation` object to the `ReservationChannel`. A separate Goroutine, the `ReservationWorker`, listens on this channel for incoming reservations.

## Mutexes

We use a `sync.Mutex` to prevent race conditions when updating book availability. Each book has its own mutex. Before a book's status is updated, the mutex is locked. After the update, the mutex is unlocked. This ensures that only one Goroutine can modify a book's status at a time.

## Auto-Cancellation of Reservations

If a reserved book is not borrowed within 5 seconds, it is automatically unreserved. This is handled by a timer-based Goroutine. When a reservation is made, a new Goroutine is started with a 5-second timer. If the timer expires and the book has not been borrowed, the book's status is set back to "Available".
