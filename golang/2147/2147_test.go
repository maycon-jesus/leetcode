package main

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name     string
		corridor string
		want     int
	}{
		{
			name:     "Exemplo 1 - SSPPSPS",
			corridor: "SSPPSPS",
			want:     3,
		},
		{
			name:     "Exemplo 2 - PPSPSP",
			corridor: "PPSPSP",
			want:     1,
		},
		{
			name:     "Exemplo 3 - apenas uma cadeira",
			corridor: "S",
			want:     0,
		},
		{
			name:     "Corredor vazio",
			corridor: "",
			want:     0,
		},
		{
			name:     "Apenas plantas",
			corridor: "PPPP",
			want:     0,
		},
		{
			name:     "Número ímpar de cadeiras",
			corridor: "SSS",
			want:     0,
		},
		{
			name:     "Exatamente duas cadeiras",
			corridor: "SS",
			want:     1,
		},
		{
			name:     "Duas cadeiras com plantas no meio",
			corridor: "SPPPS",
			want:     1,
		},
		{
			name:     "Quatro cadeiras sem plantas entre divisões",
			corridor: "SSSS",
			want:     1,
		},
		{
			name:     "Quatro cadeiras com uma planta entre divisões",
			corridor: "SSPSS",
			want:     2,
		},
		{
			name:     "Quatro cadeiras com duas plantas entre divisões",
			corridor: "SSPPSS",
			want:     3,
		},
		{
			name:     "Seis cadeiras com múltiplas plantas",
			corridor: "SSPPSPSPS",
			want:     0,
		},
		{
			name:     "Seis cadeiras em sequência",
			corridor: "SSSSSS",
			want:     1,
		},
		{
			name:     "Plantas no início",
			corridor: "PPSS",
			want:     1,
		},
		{
			name:     "Plantas no final",
			corridor: "SSPP",
			want:     1,
		},
		{
			name:     "Plantas no início e no final",
			corridor: "PPSSPP",
			want:     1,
		},
		{
			name:     "Caso com múltiplas divisões",
			corridor: "SSPSSPSS",
			want:     4,
		},
		{
			name:     "Caso com três plantas entre divisões",
			corridor: "SSPPPSS",
			want:     4,
		},
		{
			name:     "Oito cadeiras com plantas variadas",
			corridor: "SSPSPSPPSS",
			want:     6,
		},
		{
			name:     "Caso complexo com múltiplas combinações",
			corridor: "SSPPSPPSSPPSSS",
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfWays(tt.corridor)
			if got != tt.want {
				t.Errorf("numberOfWays(%q) = %d, want %d", tt.corridor, got, tt.want)
			}
		})
	}
}

func TestNumberOfWaysEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		corridor string
		want     int
	}{
		{
			name:     "String com tamanho 1 - cadeira",
			corridor: "S",
			want:     0,
		},
		{
			name:     "String com tamanho 1 - planta",
			corridor: "P",
			want:     0,
		},
		{
			name:     "Três cadeiras consecutivas",
			corridor: "SSS",
			want:     0,
		},
		{
			name:     "Cinco cadeiras consecutivas",
			corridor: "SSSSS",
			want:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfWays(tt.corridor)
			if got != tt.want {
				t.Errorf("numberOfWays(%q) = %d, want %d", tt.corridor, got, tt.want)
			}
		})
	}
}

func TestNumberOfWaysMultipleChoices(t *testing.T) {
	tests := []struct {
		name     string
		corridor string
		want     int
	}{
		{
			name:     "Uma escolha entre divisões",
			corridor: "SSPSS",
			want:     2,
		},
		{
			name:     "Duas escolhas entre divisões",
			corridor: "SSPPSS",
			want:     3,
		},
		{
			name:     "Três escolhas entre divisões",
			corridor: "SSPPPSS",
			want:     4,
		},
		{
			name:     "Múltiplas divisões com uma escolha cada",
			corridor: "SSPSPSPS",
			want:     0,
		},
		{
			name:     "Duas divisões com duas e uma escolhas",
			corridor: "SSPPSPS",
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfWays(tt.corridor)
			if got != tt.want {
				t.Errorf("numberOfWays(%q) = %d, want %d", tt.corridor, got, tt.want)
			}
		})
	}
}

func BenchmarkNumberOfWays(b *testing.B) {
	corridor := "SSPPSPSPPSSPPSSS"
	for i := 0; i < b.N; i++ {
		numberOfWays(corridor)
	}
}

func BenchmarkNumberOfWaysLong(b *testing.B) {
	// Criar um corredor longo com padrão repetitivo
	corridor := ""
	for i := 0; i < 1000; i++ {
		corridor += "SSPPS"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numberOfWays(corridor)
	}
}
