package models

import "fmt"

type TarotCard string

func IsValidTarotCard(card TarotCard) bool {
	allCards := GetAllTarotCards()
	fmt.Println(card)
	for _, validCard := range allCards {
		if validCard == string(card) {
			return true
		}
	}
	return false
}

const (
	// Major Arcana
	TheFool          TarotCard = "The Fool"
	TheMagician      TarotCard = "The Magician"
	TheHighPriestess TarotCard = "The High Priestess"
	TheEmpress       TarotCard = "The Empress"
	TheEmperor       TarotCard = "The Emperor"
	TheHierophant    TarotCard = "The Hierophant"
	TheLovers        TarotCard = "The Lovers"
	TheChariot       TarotCard = "The Chariot"
	Strength         TarotCard = "Strength"
	TheHermit        TarotCard = "The Hermit"
	WheelOfFortune   TarotCard = "Wheel of Fortune"
	Justice          TarotCard = "Justice"
	TheHangedMan     TarotCard = "The Hanged Man"
	Death            TarotCard = "Death"
	Temperance       TarotCard = "Temperance"
	TheDevil         TarotCard = "The Devil"
	TheTower         TarotCard = "The Tower"
	TheStar          TarotCard = "The Star"
	TheMoon          TarotCard = "The Moon"
	TheSun           TarotCard = "The Sun"
	Judgement        TarotCard = "Judgement"
	TheWorld         TarotCard = "The World"

	// Minor Arcana - Wands
	AceOfWands    TarotCard = "Ace of Wands"
	TwoOfWands    TarotCard = "Two of Wands"
	ThreeOfWands  TarotCard = "Three of Wands"
	FourOfWands   TarotCard = "Four of Wands"
	FiveOfWands   TarotCard = "Five of Wands"
	SixOfWands    TarotCard = "Six of Wands"
	SevenOfWands  TarotCard = "Seven of Wands"
	EightOfWands  TarotCard = "Eight of Wands"
	NineOfWands   TarotCard = "Nine of Wands"
	TenOfWands    TarotCard = "Ten of Wands"
	PageOfWands   TarotCard = "Page of Wands"
	KnightOfWands TarotCard = "Knight of Wands"
	QueenOfWands  TarotCard = "Queen of Wands"
	KingOfWands   TarotCard = "King of Wands"

	// Minor Arcana - Cups
	AceOfCups    TarotCard = "Ace of Cups"
	TwoOfCups    TarotCard = "Two of Cups"
	ThreeOfCups  TarotCard = "Three of Cups"
	FourOfCups   TarotCard = "Four of Cups"
	FiveOfCups   TarotCard = "Five of Cups"
	SixOfCups    TarotCard = "Six of Cups"
	SevenOfCups  TarotCard = "Seven of Cups"
	EightOfCups  TarotCard = "Eight of Cups"
	NineOfCups   TarotCard = "Nine of Cups"
	TenOfCups    TarotCard = "Ten of Cups"
	PageOfCups   TarotCard = "Page of Cups"
	KnightOfCups TarotCard = "Knight of Cups"
	QueenOfCups  TarotCard = "Queen of Cups"
	KingOfCups   TarotCard = "King of Cups"

	// Minor Arcana - Swords
	AceOfSwords    TarotCard = "Ace of Swords"
	TwoOfSwords    TarotCard = "Two of Swords"
	ThreeOfSwords  TarotCard = "Three of Swords"
	FourOfSwords   TarotCard = "Four of Swords"
	FiveOfSwords   TarotCard = "Five of Swords"
	SixOfSwords    TarotCard = "Six of Swords"
	SevenOfSwords  TarotCard = "Seven of Swords"
	EightOfSwords  TarotCard = "Eight of Swords"
	NineOfSwords   TarotCard = "Nine of Swords"
	TenOfSwords    TarotCard = "Ten of Swords"
	PageOfSwords   TarotCard = "Page of Swords"
	KnightOfSwords TarotCard = "Knight of Swords"
	QueenOfSwords  TarotCard = "Queen of Swords"
	KingOfSwords   TarotCard = "King of Swords"

	// Minor Arcana - Pentacles
	AceOfPentacles    TarotCard = "Ace of Pentacles"
	TwoOfPentacles    TarotCard = "Two of Pentacles"
	ThreeOfPentacles  TarotCard = "Three of Pentacles"
	FourOfPentacles   TarotCard = "Four of Pentacles"
	FiveOfPentacles   TarotCard = "Five of Pentacles"
	SixOfPentacles    TarotCard = "Six of Pentacles"
	SevenOfPentacles  TarotCard = "Seven of Pentacles"
	EightOfPentacles  TarotCard = "Eight of Pentacles"
	NineOfPentacles   TarotCard = "Nine of Pentacles"
	TenOfPentacles    TarotCard = "Ten of Pentacles"
	PageOfPentacles   TarotCard = "Page of Pentacles"
	KnightOfPentacles TarotCard = "Knight of Pentacles"
	QueenOfPentacles  TarotCard = "Queen of Pentacles"
	KingOfPentacles   TarotCard = "King of Pentacles"
)

// GetAllTarotCards returns all tarot card names as a slice of strings
func GetAllTarotCards() []string {
	return []string{
		// Major Arcana
		string(TheFool),
		string(TheMagician),
		string(TheHighPriestess),
		string(TheEmpress),
		string(TheEmperor),
		string(TheHierophant),
		string(TheLovers),
		string(TheChariot),
		string(Strength),
		string(TheHermit),
		string(WheelOfFortune),
		string(Justice),
		string(TheHangedMan),
		string(Death),
		string(Temperance),
		string(TheDevil),
		string(TheTower),
		string(TheStar),
		string(TheMoon),
		string(TheSun),
		string(Judgement),
		string(TheWorld),

		// Minor Arcana - Wands
		string(AceOfWands),
		string(TwoOfWands),
		string(ThreeOfWands),
		string(FourOfWands),
		string(FiveOfWands),
		string(SixOfWands),
		string(SevenOfWands),
		string(EightOfWands),
		string(NineOfWands),
		string(TenOfWands),
		string(PageOfWands),
		string(KnightOfWands),
		string(QueenOfWands),
		string(KingOfWands),

		// Minor Arcana - Cups
		string(AceOfCups),
		string(TwoOfCups),
		string(ThreeOfCups),
		string(FourOfCups),
		string(FiveOfCups),
		string(SixOfCups),
		string(SevenOfCups),
		string(EightOfCups),
		string(NineOfCups),
		string(TenOfCups),
		string(PageOfCups),
		string(KnightOfCups),
		string(QueenOfCups),
		string(KingOfCups),

		// Minor Arcana - Swords
		string(AceOfSwords),
		string(TwoOfSwords),
		string(ThreeOfSwords),
		string(FourOfSwords),
		string(FiveOfSwords),
		string(SixOfSwords),
		string(SevenOfSwords),
		string(EightOfSwords),
		string(NineOfSwords),
		string(TenOfSwords),
		string(PageOfSwords),
		string(KnightOfSwords),
		string(QueenOfSwords),
		string(KingOfSwords),

		// Minor Arcana - Pentacles
		string(AceOfPentacles),
		string(TwoOfPentacles),
		string(ThreeOfPentacles),
		string(FourOfPentacles),
		string(FiveOfPentacles),
		string(SixOfPentacles),
		string(SevenOfPentacles),
		string(EightOfPentacles),
		string(NineOfPentacles),
		string(TenOfPentacles),
		string(PageOfPentacles),
		string(KnightOfPentacles),
		string(QueenOfPentacles),
		string(KingOfPentacles),
	}
}
