# A Balatro Round 1 Calculator

Calculating how to best start the worst thing to happen to human productivity since the 1989 release of Belgian Techno anthem "Pump Up the Jam".

## Balatro

Balatro is a game based on poker. Without getting TOO DEEP into it, you make poker hands which are then scored based on how they rank as in regular poker.

A single "game" of Balatro (known as a run) is broken into levels known as "antes", each of which consists of three rounds known as "blinds". To advance, you have to clear each blind by scoring enough points based on the hands you play (you can play multiple hands per round, but have a small number of hands you can play per round).

After you beat a round, you're then given "money" based on how many hands you have left to play that round. This money is important as it gives you the ability to gain powers that make your hands higher scoring.

## Ante 1 Small Blind - the first round of the game

Your goal here is to score 300 points. However, to maximize the amount of money you can get from this blind, it's STRONGLY recommended to only play one hand (i.e. win in one shot).

For those seeking a refresher on poker hands, here's the list in descending order of value:

1. Straight Flush - five consecutive values from one suit
2. Four of a Kind - four cards of the same value 
3. Full House - three cards of one value plus a pair of a different value
4. Flush - five non-consecutive values from one suit
5. Straight - five consectuve values from multiple suits
6. Three of a Kind - three cards of the same value plus two other cards that aren't a pair
7. Two Pair - two cards of one value plus two cards of a different value and a fifth card of a third value
8. Pair - two cards of one value plus three other cards of three other values
9. High Card - five cards where none of the values match and there are at least two different suits

Each hand will score a different amount of points in a round based on two things: the value of the hand itself, and whichever cards are played; each card has a point value based on its number: 2-10 score their respective values of points, face cards (Jack, Queen, and King) also score 10 points, and Aces score 11. There are ways of adjusting how many points are scored, but it's HIGHLY unlikely you will have any of them going into round 1. Thus, I'm going to focus on the base values the game uses for scoring; I won't get too deep into what these are, but in short: for round 1, anything three of a kind or below will NOT one shot this round (as the "assorted other cards" are not typically factored into scoring, but would not clear even if they were), while anything four of a kind or above will clear BEFORE calculating your actual score based on hand value alone, so they're not involved in this. Additionally, These programs do not calculate for straights as the "math" on those is easy (it needs to be at least Queen high in order to clear). Therefore, we're left with just flushes and full houses.

## Full House
Let's start with the (relatively) easier of the two to figure out.

A full house is three of one card plus two of another. For example, if we had, say, three eights and two sixes, we would have (3x8) + (2x6) for a total of 36 points (24 + 12).

For our purposes, we'll just simply say that we need 35 points to pass (as the hand score accounts for quite a bit of the "work"). This is where the Go file comes in; originally a thing that I did just for "fun", this program calculates how many different full houses will clear round 1 in one shot. The original version was just concerned about "values" (i.e. the face cards all being worth 10 points), but the version here is changed to reflect the actual cards themselves (to an extent because of types; there may be a way to distinguish between the four 10-point cards, but that's getting into the realm of "coding dark magic").

Long story short: there are 92 full houses you can play; this seems like a lot, but you have to remember that there are OVER 3,700 possible full houses. I won't list all of the full houses here, but the main jist of what I found was:

- The "worst" full house to beat round 1 in a single hand is 5s over 10s (three fives with a pair of 10s or face cards; this "x over y" language is common to poker)
- The value of the pair required lowers as the value of the three of a kind goes up (almost perfectly linearly)
- The hand won't clear if the three of a kind is 2-4, and will always clear if the three of a kind is aces (10/J/Q/K over 2 comes up ONE POINT SHORT)

## Flush
Here's where things get tricky - there are A LOT of different combinations (over 5,000 according to one site I found), and thanks to a 40 point requirement, a lot of them won't one shot (for example, A, 2, 3, 4, 6 comes up woefully short - and this hand has the highest scoring card!).

Unfortunately, there's no "easy" way to remember how to determine if your flush has enough value; you just have to do the math yourself. I won't comment on the mental math skills of the audience here; instead, I made a simple HTML form with some JavaScript magic (partially thanks to Copilot) which can figure out if your flush has enough value to pass round 1. Simply choose five (and only five) cards, and the page will tell you if your cards are good enough or not!

The most difficult part of this was trying to figure out how to detect straight flushes (as that makes the individual card values irrelevant), since a Straight Flush of values that add to <40 would register as not being good enough (including the "wraparound" of A, 2, 3, 4, 5). For various reasons, I was finding that either the program would think EVERY hand or NO hand was a straight flush.

P.S.: There IS a way to be able to play four card flushes, but you almost certainly won't have it this early in the game, and it only affects a very small subset of hands (any non-straight flushes with an A plus values 9 and higher) as the same effect works on straights as well (i.e. 10, J, Q, K becomes a straight, so if its suited it's a straight flush). Coincidentally, I should mention that you cannot clear round 1 with any four card straight as that requires 45 points, and the highest possible 4 card straight (A, K, Q, J) yields only 41.

## Other notes
If you were wondering about the future for this project, I personally don't see the point in getting any further than this; at that point, you might as well just play the darn game (after all, it was nominated for Game of the Year)! However, I will note that all of this ONLY works for one shotting round 1; there's too much potential variance to be "safely" able to calculate for round 2. I will note, however, that without any modifications (i.e. you go into round 2 with the same basic conditions as round 1), you won't be able to clear round 2 in one hand with ANY flush or full house (the most points you can get is 372, and you need 450).
