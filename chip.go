package main

const pad = "        " // 8 espacios de margen izquierdo

// ChipFace devuelve la cara de Chip con color según la emoción.
// Emociones: feliz, explicando, listo, error, hablando
func ChipFace(emotion string) string {
	switch emotion {
	case "feliz":
		return Yellow + pad + "▛▀▀▀▀▀▀▀▀▀▀▀▜" + Reset + "\n" +
			Yellow + pad + "▌  ◉     ◉  ▐" + Reset + "\n" +
			Yellow + pad + "▌     ◡     ▐" + Reset + "\n" +
			Yellow + pad + "▙▄▄▄▄▄▄▄▄▄▄▄▟" + Reset
	case "hablando":
		return Yellow + pad + "▛▀▀▀▀▀▀▀▀▀▀▀▜" + Reset + "\n" +
			Yellow + pad + "▌  ◉     ◉  ▐" + Reset + "\n" +
			Yellow + pad + "▌     ◡     ▐" + Reset + "\n" +
			Yellow + pad + "▙▄▄▄▄▄▄▇▄▄▄▄▟" + Reset
	case "explicando":
		return Yellow + pad + "▛▀▀▀▀▀▀▀▀▀▀▀▜" + Reset + "\n" +
			Yellow + pad + "▌  ◉     ◉  ▐" + Reset + "\n" +
			Yellow + pad + "▌     ¿     ▐" + Reset + "\n" +
			Yellow + pad + "▙▄▄▄▄▄▄▄▄▄▄▄▟" + Reset
	case "listo":
		return Yellow + pad + "▛▀▀▀▀▀▀▀▀▀▀▀▜" + Reset + "\n" +
			Yellow + pad + "▌  ^     ^  ▐" + Reset + "\n" +
			Yellow + pad + "▌     ▔     ▐" + Reset + "\n" +
			Yellow + pad + "▙▄▄▄▄▄▄▄▄▄▄▄▟" + Reset
	case "error":
		return Yellow + pad + "▛▀▀▀▀▀▀▀▀▀▀▀▜" + Reset + "\n" +
			Yellow + pad + "▌  >     <  ▐" + Reset + "\n" +
			Yellow + pad + "▌     ▀     ▐" + Reset + "\n" +
			Yellow + pad + "▙▄▄▄▄▄▄▄▄▄▄▄▟" + Reset
	default:
		// Si no reconocen la emoción, feliz por defecto
		return ChipFace("feliz")
	}
}
