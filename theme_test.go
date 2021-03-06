package pterm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTheme_WithDescriptionMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithDescriptionMessageStyle(s)

	assert.Equal(t, s, p2.DescriptionMessageStyle)
}

func TestTheme_WithDescriptionPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithDescriptionPrefixStyle(s)

	assert.Equal(t, s, p2.DescriptionPrefixStyle)
}

func TestTheme_WithErrorMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithErrorMessageStyle(s)

	assert.Equal(t, s, p2.ErrorMessageStyle)
}

func TestTheme_WithErrorPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithErrorPrefixStyle(s)

	assert.Equal(t, s, p2.ErrorPrefixStyle)
}

func TestTheme_WithFatalMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithFatalMessageStyle(s)

	assert.Equal(t, s, p2.FatalMessageStyle)
}

func TestTheme_WithFatalPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithFatalPrefixStyle(s)

	assert.Equal(t, s, p2.FatalPrefixStyle)
}

func TestTheme_WithHighlightStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithHighlightStyle(s)

	assert.Equal(t, s, p2.HighlightStyle)
}

func TestTheme_WithInfoMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithInfoMessageStyle(s)

	assert.Equal(t, s, p2.InfoMessageStyle)
}

func TestTheme_WithInfoPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithInfoPrefixStyle(s)

	assert.Equal(t, s, p2.InfoPrefixStyle)
}

func TestTheme_WithPrimaryStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithPrimaryStyle(s)

	assert.Equal(t, s, p2.PrimaryStyle)
}

func TestTheme_WithSecondaryStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithSecondaryStyle(s)

	assert.Equal(t, s, p2.SecondaryStyle)
}

func TestTheme_WithSuccessMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithSuccessMessageStyle(s)

	assert.Equal(t, s, p2.SuccessMessageStyle)
}

func TestTheme_WithSuccessPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithSuccessPrefixStyle(s)

	assert.Equal(t, s, p2.SuccessPrefixStyle)
}

func TestTheme_WithWarningMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithWarningMessageStyle(s)

	assert.Equal(t, s, p2.WarningMessageStyle)
}

func TestTheme_WithWarningPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithWarningPrefixStyle(s)

	assert.Equal(t, s, p2.WarningPrefixStyle)
}

func TestTheme_WithBulletListBulletStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithBulletListBulletStyle(s)

	assert.Equal(t, s, p2.BulletListBulletStyle)
}

func TestTheme_WithBulletListTextStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithBulletListTextStyle(s)

	assert.Equal(t, s, p2.BulletListTextStyle)
}

func TestTheme_WithLetterStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithLetterStyle(s)

	assert.Equal(t, s, p2.LetterStyle)
}

func TestTheme_WithDebugMessageStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithDebugMessageStyle(s)

	assert.Equal(t, s, p2.DebugMessageStyle)
}

func TestTheme_WithDebugPrefixStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithDebugPrefixStyle(s)

	assert.Equal(t, s, p2.DebugPrefixStyle)
}

func TestTheme_WithTreeStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithTreeStyle(s)

	assert.Equal(t, s, p2.TreeStyle)
}

func TestTheme_WithTreeTextStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithTreeTextStyle(s)

	assert.Equal(t, s, p2.TreeTextStyle)
}

func TestTheme_WithBoxStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithBoxStyle(s)

	assert.Equal(t, s, p2.BoxStyle)
}

func TestTheme_WithBoxTextStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithBoxTextStyle(s)

	assert.Equal(t, s, p2.BoxTextStyle)
}

func TestTheme_WithBarLabelStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithBarLabelStyle(s)

	assert.Equal(t, s, p2.BarLabelStyle)
}

func TestTheme_WithBarStyle(t *testing.T) {
	s := Style{FgRed, BgBlue, Bold}
	p := Theme{}
	p2 := p.WithBarStyle(s)

	assert.Equal(t, s, p2.BarStyle)
}
