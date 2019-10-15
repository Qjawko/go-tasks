package piscine

import "fmt"

/*

Yes, there does exist such a way.
First you select a digit d from {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.
Then you select a digit e from ({0, 1, 2, 3, 4, 5, 6, 7, 8, 9}-d).
Then you select a digit f from (({0, 1, 2, 3, 4, 5, 6, 7, 8, 9}-d)-e).
You first select 0 for d, then 1, and so on until you get to 7.
And, you always select the least digit first for e and f also, with the additional condition that d < e < f.
List out the first sequence, 012, 013, 014, 015, 016, 017, 018, 019.
Then list all the other numbers beneath them with the condition that for all numbers e and f,
and with d held constant, the digits for e and f follow the natural number sequence down the column.
Partition each set of sequences by d.
The column rule only applies within each partition.
(this description might come as incomplete or could use some revision).

*/

func PrintComb() {
	var i, j, k = 0, 0, 0
	for i = 0; i <= 9; i++ {
		for j = 1; j <= 9; j++ {
			for k = 2; k <= 9; k++ {
				if i < j && j < k {
					fmt.Printf("%d%d%d", i, j, k)

					if i == 7 && j == 8 && k == 9 {
						continue
					}
					fmt.Print(", ")
				}
			}
		}
	}
}
