package cmd

import (
	"testing"
)

func TestUICmd_ExecutesWithoutError(t *testing.T) {
	// O comando deve executar sem pânico ou erro
	rootCmd.SetArgs([]string{"ui", "--testmode"})

	// Por padrão, RunQuickUI roda o Bubble Tea — então o teste só verifica se não falha
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}
