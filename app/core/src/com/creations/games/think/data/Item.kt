package com.creations.games.think.data

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
		val puzzles: Array<Puzzle>): Item() {
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
