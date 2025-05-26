package main

import "fmt"
/*
Design a hit counter which counts the number of hits received in the past 5 minutes.

Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing). You may assume that the earliest timestamp starts at 1.

It is possible that several hits arrive roughly at the same time.

Example:

HitCounter counter = new HitCounter();

// hit at timestamp 1.
counter.hit(1);

// hit at timestamp 2.
counter.hit(2);

// hit at timestamp 3.
counter.hit(3);

// get hits at timestamp 4, should return 3.
counter.getHits(4);

// hit at timestamp 300.
counter.hit(300);

// get hits at timestamp 300, should return 4.
counter.getHits(300);

// get hits at timestamp 301, should return 3.
counter.getHits(301);
Follow up:
What if the number of hits per second could be very large? Does your design scale?
*/

type HitCounter struct {
	counts []int
}

func (h *HitCounter) Hit(timestamp int) {
	h.counts = append(h.counts, timestamp)

	for (len(h.counts) > 0) && (timestamp - 300 >= h.counts[0])  {
		h.counts = h.counts[1:]
	}
}

func (h *HitCounter) GetHits(timestamp int) int {
	for (len(h.counts) > 0) && (timestamp - 300 >= h.counts[0])  {
		h.counts = h.counts[1:]
	}

	return len(h.counts)
}

func main() {
	hc := &HitCounter{counts: make([]int, 0)}

	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	fmt.Println(hc.GetHits(4))
	hc.Hit(300)
	fmt.Println(hc.GetHits(300))
	fmt.Println(hc.GetHits(301))
}


/*
Усложнение (Yandex):

Есть последовательность запросов, упорядоченная по времени
Запросы бывают двух видов:
1. Пользователь сгенерировал событие (нажал на красную кнопку) в момент времени
2. Посчитать количество пользователей в момент времени , которые за последние
5 минут сгенерировали >= 1000 событий (нажали на красную кнопку >= 1000 раз).
Необходимо реализовать структуру данных, умеющую эффективно обрабатывать данные
запросы.
Ограничения по времени на обработку запроса: O(1) амортизированное.
*/
