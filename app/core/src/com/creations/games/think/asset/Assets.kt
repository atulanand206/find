package com.creations.games.think.asset

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.scenes.scene2d.ui.Label
import com.creations.games.engine.asset.GameAssetManager
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.widgets.TextLabel

class Assets(private val gam: GameAssetManager) {
	lateinit var circle: Texture
	lateinit var rect: Texture

	init {
		//load assets
		gam.addAssetsToLoad(FileNames.circlePath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.squarePath, Texture::class.java)

		//load fonts
		FontSize.values().forEach { gam.addFontToLoad(FileNames.fontPath, it.value) }

		gam.addOnLoadListener {
			//automatically added when assets are loaded
			circle = gam.get(FileNames.circlePath, Texture::class.java)
			rect = gam.get(FileNames.squarePath, Texture::class.java)
		}
	}

	/**
	 * creates a new label with the provided text, color and size
	 */
	fun createLabel(
			text: String,
			size: FontSize = FontSize.F12,
			color: Color = Color.WHITE
	) = TextLabel(text, Label.LabelStyle(gam.getFont(FileNames.fontPath, size.value), color))

	fun addLabel(
			text: String,
			size: FontSize,
			color: Color,
			x: Float,
			y: Float,
			align: Int): GameObject {
		val label = createLabel(text, size, color)
		label.setPosition(x, y, align)
		return label
	}
}