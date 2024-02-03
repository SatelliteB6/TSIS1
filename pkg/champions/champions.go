package champions

type Response struct {
	Champions []Champion `json:"champions"`
}

type Champion struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Spells []Spell `json:"spells"`
}

type Spell struct {
	SpellId int    `json:"spell_id"`
	Name    string `json:"name"`
}

func prepareChampions() []Champion {
	var champions []Champion

	// Example champions with spells
	champions = append(champions, Champion{
		Id:   1,
		Name: "Ahri",
		Spells: []Spell{
			{SpellId: 1, Name: "Orb of Deception"},
			{SpellId: 2, Name: "Fox-Fire"},
			{SpellId: 3, Name: "Charm"},
			{SpellId: 4, Name: "Spirit Rush"},
		},
	})

	champions = append(champions, Champion{
		Id:   2,
		Name: "Ezreal",
		Spells: []Spell{
			{SpellId: 1, Name: "Mystic Shot"},
			{SpellId: 2, Name: "Essence Flux"},
			{SpellId: 3, Name: "Arcane Shift"},
			{SpellId: 4, Name: "Trueshot Barrage"},
		},
	})

	champions = append(champions, Champion{
		Id:   3,
		Name: "Darius",
		Spells: []Spell{
			{SpellId: 1, Name: "Decimate"},
			{SpellId: 2, Name: "Crippling Strike"},
			{SpellId: 3, Name: "Apprehend"},
			{SpellId: 4, Name: "Noxian Guillotine"},
		},
	})

	champions = append(champions, Champion{
		Id:   4,
		Name: "Jinx",
		Spells: []Spell{
			{SpellId: 1, Name: "Switcheroo!"},
			{SpellId: 2, Name: "Zap!"},
			{SpellId: 3, Name: "Flame Chompers!"},
			{SpellId: 4, Name: "Super Mega Death Rocket!"},
		},
	})

	champions = append(champions, Champion{
		Id:   5,
		Name: "Leona",
		Spells: []Spell{
			{SpellId: 1, Name: "Shield of Daybreak"},
			{SpellId: 2, Name: "Eclipse"},
			{SpellId: 3, Name: "Zenith Blade"},
			{SpellId: 4, Name: "Solar Flare"},
		},
	})

	champions = append(champions, Champion{
		Id:   6,
		Name: "Katarina",
		Spells: []Spell{
			{SpellId: 1, Name: "Bouncing Blades"},
			{SpellId: 2, Name: "Sinister Steel"},
			{SpellId: 3, Name: "Shunpo"},
			{SpellId: 4, Name: "Death Lotus"},
		},
	})

	champions = append(champions, Champion{
		Id:   7,
		Name: "Yasuo",
		Spells: []Spell{
			{SpellId: 1, Name: "Steel Tempest"},
			{SpellId: 2, Name: "Wind Wall"},
			{SpellId: 3, Name: "Sweeping Blade"},
			{SpellId: 4, Name: "Last Breath"},
		},
	})

	champions = append(champions, Champion{
		Id:   8,
		Name: "Lux",
		Spells: []Spell{
			{SpellId: 1, Name: "Light Binding"},
			{SpellId: 2, Name: "Prismatic Barrier"},
			{SpellId: 3, Name: "Lucent Singularity"},
			{SpellId: 4, Name: "Final Spark"},
		},
	})

	champions = append(champions, Champion{
		Id:   9,
		Name: "Lee Sin",
		Spells: []Spell{
			{SpellId: 1, Name: "Sonic Wave"},
			{SpellId: 2, Name: "Safeguard"},
			{SpellId: 3, Name: "Tempest"},
			{SpellId: 4, Name: "Dragon's Rage"},
		},
	})

	return champions
}
