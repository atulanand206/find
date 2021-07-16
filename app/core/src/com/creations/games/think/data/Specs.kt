package com.creations.games.think.data

import com.badlogic.gdx.graphics.Color
import java.util.*

data class Appearance(val color: Color, val transparent: Boolean)

data class ViewSpecs(
		val id: UUID,
		val light: Appearance,
		val dark: Appearance,
		// percent height from the main image; ranging 0 to 100
		val height: Float,
		// percent width from the main image; ranging 0 to 100
		val width: Float,
		val key: String = "")

data class ImageSpecs(val id: UUID, val ref: String)