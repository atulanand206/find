package com.creations.games.think.data

import com.badlogic.gdx.graphics.Color
import com.creations.games.engine.values.Values
import java.util.*

abstract class Item {
	abstract val id: UUID
	abstract val name: String
	abstract val text: String
	abstract val specs: ViewSpecs
	abstract val img: ImageSpecs
}

data class Clue(
		override val id: UUID,
		override val name: String,
		override val text: String,
		override val specs: ViewSpecs,
		// Image assigned to the clue.
		override val img: ImageSpecs,
		// Preferably, use integral values;
		// Fraction possible when easily want to introduce a clue between some two.
		val rating: Float,
		// id of the puzzle image.
		val puzzleImgId: UUID,
		// Percent width in the puzzle image from bottom left
		val xP: Float,
		// Percent height in the puzzle image from bottom left
		val yP: Float) : Item()

data class Puzzle(
		override val id: UUID,
		override val name: String,
		override val text: String,
		override val specs: ViewSpecs,
		override val img: ImageSpecs,
		var note: String,
		val solution: String,
		val clues: Array<Clue>) : Item() {
	override fun equals(other: Any?): Boolean {
		if (this === other) return true
		if (javaClass != other?.javaClass) return false

		other as Puzzle

		if (id != other.id) return false
		if (name != other.name) return false
		if (text != other.text) return false
		if (specs != other.specs) return false
		if (img != other.img) return false
		if (note != other.note) return false
		if (solution != other.solution) return false
		if (!clues.contentEquals(other.clues)) return false

		return true
	}

	override fun hashCode(): Int {
		var result = id.hashCode()
		result = 31 * result + name.hashCode()
		result = 31 * result + text.hashCode()
		result = 31 * result + specs.hashCode()
		result = 31 * result + img.hashCode()
		result = 31 * result + note.hashCode()
		result = 31 * result + solution.hashCode()
		result = 31 * result + clues.contentHashCode()
		return result
	}
}

data class Theme(
		override val id: UUID,
		override val name: String,
		override val text: String,
		override val specs: ViewSpecs,
		override val img: ImageSpecs,
		val puzzles: Array<Puzzle>) : Item() {
	override fun equals(other: Any?): Boolean {
		if (this === other) return true
		if (javaClass != other?.javaClass) return false

		other as Theme

		if (id != other.id) return false
		if (name != other.name) return false
		if (text != other.text) return false
		if (specs != other.specs) return false
		if (img != other.img) return false
		if (!puzzles.contentEquals(other.puzzles)) return false

		return true
	}

	override fun hashCode(): Int {
		var result = id.hashCode()
		result = 31 * result + name.hashCode()
		result = 31 * result + text.hashCode()
		result = 31 * result + specs.hashCode()
		result = 31 * result + img.hashCode()
		result = 31 * result + puzzles.contentHashCode()
		return result
	}
}

val xClueDark = Appearance(Color.valueOf("#EA9999"), false)
val xClueLight = Appearance(Color.valueOf("#985555"), false)
val xPuzzleDark = Appearance(Color.valueOf("#E26C6C"), false)
val xPuzzleLight = Appearance(Color.valueOf("#580808"), false)
val xThemeDark = Appearance(Color.valueOf("#fcfcfc"), false)
val xThemeLight = Appearance(Color.valueOf("#3A2E2E"), false)
val xBackgroundDark = Color.valueOf("#fcfcfc")
val xBackgroundLight = Color.valueOf("#152259")

private fun backgroundColor(): Color = if (Values.isDark) xBackgroundDark else xBackgroundLight

val xClueViewSpecs = ViewSpecs(UUID.randomUUID(), xClueLight, xClueDark, 10f, 10f, "clue")
val xPuzzleViewSpecs = ViewSpecs(UUID.randomUUID(), xPuzzleLight, xPuzzleDark, 100f, 100f, "puzzle")
val xThemeViewSpecs = ViewSpecs(UUID.randomUUID(), xThemeLight, xThemeDark, 100f, 100f, "theme")

val xPuzzleImageId: UUID = UUID.randomUUID()
val xClueImageSpecs = ImageSpecs(UUID.randomUUID(), "clue-image-ref")
val xPuzzleImageSpecs = ImageSpecs(xPuzzleImageId, "puzzle-image-ref")
val xThemeImageSpecs = ImageSpecs(UUID.randomUUID(), "theme-image-ref")

val xClue = Clue(UUID.randomUUID(), "clue-name", "clue-text",
		xClueViewSpecs, xClueImageSpecs, 100f, xPuzzleImageId, 100f, 100f)
val xClues = arrayOf(xClue)

val xPuzzle = Puzzle(UUID.randomUUID(), "Converging", "The good old series \nis around for centuries \nand has stirred many minds.",
		xPuzzleViewSpecs, xPuzzleImageSpecs, "puzzle-note", "puzzle-solution", xClues)
val xPuzzles = arrayOf(xPuzzle, xPuzzle, xPuzzle)

val xTheme = Theme(UUID.randomUUID(), "THE FIBONACCI", "theme-text",
		xThemeViewSpecs, xThemeImageSpecs, xPuzzles)