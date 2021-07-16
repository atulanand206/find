package com.creations.games.think.data

import java.util.*

data class Color(val rgba: IntArray, val hsla: IntArray, val hex: String, val transparent: Boolean) {
	override fun equals(other: Any?): Boolean {
		if (this === other) return true
		if (javaClass != other?.javaClass) return false

		other as Color

		if (!rgba.contentEquals(other.rgba)) return false
		if (!hsla.contentEquals(other.hsla)) return false
		if (hex != other.hex) return false
		if (transparent != other.transparent) return false

		return true
	}

	override fun hashCode(): Int {
		var result = rgba.contentHashCode()
		result = 31 * result + hsla.contentHashCode()
		result = 31 * result + hex.hashCode()
		result = 31 * result + transparent.hashCode()
		return result
	}
}

data class ViewSpecs(val id: UUID, val light: Color, val dark: Color, val height: Int, val width: Int, val key: String = "")

data class ImageSpecs(val id: UUID, val ref: String)