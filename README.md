# balatro_calculator

Calculating how to beat the worst thing to happen to human productivity since the 1989 release of Belgian Techno anthem "Pump Up the Jam".

## Balatro

Balatro is a game based on poker. Without getting TOO DEEP into it, you make poker hands which are then scored based on how they rank as in regular poker.

A single "game" of Balatro (known as a run) is broken into levels known as "antes", each of which consists of three rounds known as "blinds". To advance, you have to clear each blind by scoring enough points based on the hands you play (you can play multiple hands per round).

## Ante 1 Small Blind

Your goal in this ante is to score 300 chips. However, to maximize the amount of money you can get from this blind, it's STRONGLY recommended to only play one hand (i.e. one shot).

Anything three of a kind or below will NOT clear this blind, while anything four of a kind or above will clear BEFORE calculating chips, so they're not involved in this. Additionally, straights are not here as the "math" on those is easy (it needs to be at least Queen high in order to clear). Therefore, we're left with just flushes and full houses.

## Full House
Let's start with the (relatively) easier of the two to figure out.

A full house is three of one card plus two of another.

For our purposes, we'll ignore most of how scoring in Balatro works and just simply say that we need 35 "points" to pass. 

This is where the go file comes in. Originally a thing that I did just for "fun", this program calculates how many different full houses will clear round 1 in one shot. Long story short: there are 92 full houses you can play; this seems like a lot, but you have to remember that there are OVER 3,700 possible full houses.

## Flush
Here's where things get tricky.

A Flush is five cards from one of the four suits in a standard deck of cards (Spades, Hearts, Clubs, and Diamonds). This means that there are A LOT of different combinations (over 5,000 according to one site I found), and thanks to a 40 point requirement, a lot of them won't one shot.